package rbuser

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// LdapConf provides a medium for connecting to a LDAP server.
type ldapConf struct {
	user     string
	password string
	host     string
	port     int
	Conn     *ldap.Conn
}

// Connect to the LDAP server.
func (conf *ldapConf) connect() error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", conf.host, conf.port))
	if err != nil {
		return err
	}
	conf.Conn = l
	return conf.Conn.Bind(conf.user, conf.password)
}
