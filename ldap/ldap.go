// ldap
package ldap

import (
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

func (this *LdapInstance) Search(user string, password string) (ok bool, attrs map[string]string, err error) {

	attrs = make(map[string]string)
	filter := "(" + this.UserField + "=" + user + ")"
	conn, err := ldap.Dial("tcp", this.Addr)

	if err != nil {
		return false, attrs, fmt.Errorf("dial ldap fail: %s", err.Error())
	}
	defer conn.Close()

	if this.BindDn != "" {
		err = conn.Bind(this.BindDn, this.BindPassword)
	}
	if err != nil {
		return false, attrs, fmt.Errorf("ldap Bind fail: %s", err.Error())
	}

	search := ldap.NewSearchRequest(
		this.BaseDn,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,
		nil,
		nil)

	sr, err := conn.Search(search)
	if err != nil {
		return false, attrs, fmt.Errorf("ldap search fail: %s", err.Error())
	}

	defer func() {
		if err := recover(); err != nil {
			ok = false
		}
	}()
	err = conn.Bind(sr.Entries[0].DN, password)
	if err != nil {
		return false, attrs, err
	}

	for _, key := range this.Attributes {
		attrs[key] = sr.Entries[0].GetAttributeValue(key)
	}

	return true, attrs, err
}
