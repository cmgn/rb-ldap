package rbuser

// DisableUser disables a user's LDAP account.
func (rb *RbLdap) DisableUser(user RbUser) error {
	user.LoginShell = noLoginShell
	return rb.Update(user)
}

// ExpireUser disables a user's LDAP account.
func (rb *RbLdap) ExpireUser(user RbUser) error {
	user.LoginShell = expiredShell
	return rb.Update(user)
}

// EnableUser enables a user's LDAP account.
func (rb *RbLdap) RenableUser(user RbUser) error {
	return rb.ResetShell(user)
}

// ResetShell resets a user's shell.
func (rb *RbLdap) ResetShell(user RbUser) error {
	user.LoginShell = defaultShell
	return rb.Update(user)
}
