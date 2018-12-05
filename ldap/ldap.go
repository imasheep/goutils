// ldap
package ldap

import (
	"errors"
	"fmt"

	"github.com/toolkits/ldap"
)

type LdapInstance struct {
	Addr         string
	BaseDn       string
	BindDn       string
	BindPassword string
	UserField    string
	Attributes   []string
}

var Ldap *LdapInstance

func Init(ldapInstance LdapInstance) {
	Ldap = &ldapInstance
	return
}

func New() (l *LdapInstance) {
	l = Ldap
	return
}

func (this *LdapInstance) bind() (conn *ldap.Conn, err error) {

	fmt.Println(this.Addr)
	conn, err = ldap.Dial("tcp", this.Addr)
	if err != nil {
		return
	}

	err = conn.Bind(this.BindDn, this.BindPassword)
	if err != nil {
		return
	}
	return
}

func (this *LdapInstance) Auth(user string, password string) (ok bool, err error) {
	conn, err := this.bind()
	if err != nil {
		return false, err
	}
	userAttrs, err := this.Get(user)
	if err != nil {
		return false, err
	}
	err = conn.Bind(userAttrs["dn"], password)
	if err != nil {
		return false, err
	}
	return true, err
}

func (this *LdapInstance) Get(user string) (userAttrs map[string]string, err error) {

	userAttrs = make(map[string]string)

	conn, err := ldap.Dial("tcp", this.Addr)
	if err != nil {
		return
	}

	err = conn.Bind(this.BindDn, this.BindPassword)
	if err != nil {
		return
	}

	filter := "(" + this.UserField + "=" + user + ")"

	search := ldap.NewSearchRequest(
		this.BaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		nil,
		nil)

	searchResult, err := conn.Search(search)
	if err != nil {
		return userAttrs, fmt.Errorf("ldap search fail: %s", err.Error())
	}

	if len(searchResult.Entries) == 0 {
		return nil, errors.New("No Such User")
	}
	for _, key := range this.Attributes {
		userAttrs[key] = searchResult.Entries[0].GetAttributeValue(key)
	}
	userAttrs["dn"] = searchResult.Entries[0].DN
	return
}
