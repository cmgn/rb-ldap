package rbuser

import (
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

// ResetPasswd generates a new password for a user and emails it to them.
func (rb *RbLdap) ResetPasswd(uid string) error {
	passwordModifyRequest := ldap.NewPasswordModifyRequest(fmt.Sprintf("uid=%s,ou=ldap,o=redbrick", uid), "", passwd(12))
	passwordModifyResponse, err := rb.Conn.PasswordModify(passwordModifyRequest)
	if err != nil {
		return err
	}
	user, err := rb.SearchUser(fmt.Sprintf("(&(uid=%s))", uid))
	if err != nil {
		return err
	}
	user.UserPassword = passwordModifyResponse.GeneratedPassword
	return rb.mailAccountUpdate(user)
}
