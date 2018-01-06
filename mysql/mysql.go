//mysql.go
package mysql

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var RegisterModelWithPrefix = orm.RegisterModelWithPrefix

func New(flag string) (m orm.Ormer) {

	m = orm.NewOrm()
	m.Using(flag)
	return m

}

type MysqlInstance struct {
	Flag     string
	DbName   string
	Host     string
	Port     string
	User     string
	Password string
}

func regDatabase(mysqlInstance MysqlInstance) (err error) {

	instanceFlag := "RegMysqlInstance"

	flag := mysqlInstance.Flag
	name := mysqlInstance.DbName
	host := mysqlInstance.Host
	port := mysqlInstance.Port
	user := mysqlInstance.User
	password := mysqlInstance.Password

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8"

	err = orm.RegisterDataBase(flag, "mysql", dsn)
	if err != nil {
		return
	}

	orm.SetMaxIdleConns(flag, 30)
	orm.SetMaxOpenConns(flag, 100)

	fmt.Printf("%-20s: %-10s [ %s ]\n", instanceFlag, flag, dsn)

	return

}

func regDatabaseMulti(defaultFlag string, mysqlInstances []MysqlInstance) (err error) {

	for _, mysqlInstance := range mysqlInstances {

		err = regDatabase(mysqlInstance)
		if err != nil {
			return
		}

		if mysqlInstance.Flag == defaultFlag {

			mysqlInstance.Flag = "default"

			err = regDatabase(mysqlInstance)
			if err != nil {
				return
			}
		}
	}
	return
}

func Init(defaultFlag string, mysqlInstances []MysqlInstance) {

	orm.Debug = true
	regDatabaseMulti(defaultFlag, mysqlInstances)
	orm.RunSyncdb("default", false, true)

}

func (this *MysqlInstance) Init(flag string,

	dbName string,
	host string,
	port string,
	user string,
	password string) (err error) {

	this.Flag = flag
	this.DbName = dbName
	this.Host = host
	this.Port = port
	this.User = user
	this.Password = password

	return
}
