package rbuser

// DisableUser disables a user's LDAP account.
func (rb *RbLdap) DisableUser(user User) error {
	user.LoginShell = noLoginShell
	return rb.Update(user)
}

// ExpireUser disables a user's LDAP account.
func (rb *RbLdap) ExpireUser(user User) error {
	user.LoginShell = expiredShell
	return rb.Update(user)
}

// EnableUser enables a user's LDAP account.
func (rb *RbLdap) RenableUser(user User) error {
	return rb.ResetShell(user)
}

// ResetShell resets a user's shell.
func (rb *RbLdap) ResetShell(user User) error {
	user.LoginShell = defaultShell
	return rb.Update(user)
}
