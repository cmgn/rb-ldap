package rbuser

import gomail "gopkg.in/gomail.v2"

// RbLdap provides a medium for communicating with RedBrick's LDAP server.
type RbLdap struct {
	*ldapConf
	Mail *gomail.Dialer
}

// NewRbLdap creates a connection to RedBrick's LDAP server.
func NewRbLdap(user, password, host string, port int, smtp string) (*RbLdap, error) {
	rb := RbLdap{
		&ldapConf{
			user:     user,
			password: password,
			host:     host,
			port:     port,
		},
		&gomail.Dialer{
			Host: smtp,
			Port: 587,
		},
	}
	return &rb, rb.connect()
}
