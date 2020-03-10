package rbldap

import (
	"github.com/redbrick/rb-ldap/pkg/rbuser"
	"github.com/urfave/cli"
)

// AlertUnpaidUsers sends a warning email to all users that have not paid.
func AlterUnpaidUsers(ctx *cli.Context) error {
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
	}
	p := newPrompt()
	if confirm, err := p.Confirm("Email All Unpaid Users"); !confirm || err != nil {
		return err
	}
	rb, err := rbuser.NewRbLdap(
		ctx.GlobalString("user"),
		ctx.GlobalString("password"),
		ctx.GlobalString("host"),
		ctx.GlobalInt("port"),
		ctx.GlobalString("smtp"),
	)
	if err != nil {
		return err
	}
	defer rb.Conn.Close()
	return rb.AlertUnpaidUsers()
}

// DisableUnpaidUsers disables all user accounts that have not paid.
func DisableUnpaidUsers(ctx *cli.Context) error {
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
	}
	p := newPrompt()
	if confirm, err := p.Confirm("Disable All Unpaid Users"); !confirm || err != nil {
		return err
	}
	rb, err := rbuser.NewRbLdap(
		ctx.GlobalString("user"),
		ctx.GlobalString("password"),
		ctx.GlobalString("host"),
		ctx.GlobalInt("port"),
		ctx.GlobalString("smtp"),
	)
	if err != nil {
		return err
	}
	defer rb.Conn.Close()
	admin, err := p.ReadUser("Disabled by")
	if err != nil {
		return err
	}
	return rb.DisableUnpaidUsers(admin)
}

// DeleteUnpaidUsers deletes all user accounts that have not paid.
func DeleteUnpaidUsers(ctx *cli.Context) error {
	if ctx.GlobalBool("dry-run") {
		return errNotImplemented
	}
	p := newPrompt()
	if confirm, err := p.Confirm("Delete All Unpaid Users, THIS CANNOT BE UNDONE"); !confirm || err != nil {
		return err
	}
	rb, err := rbuser.NewRbLdap(
		ctx.GlobalString("user"),
		ctx.GlobalString("password"),
		ctx.GlobalString("host"),
		ctx.GlobalInt("port"),
		ctx.GlobalString("smtp"),
	)
	if err != nil {
		return err
	}
	defer rb.Conn.Close()
	return rb.DeleteUnpaidUsers()
}
