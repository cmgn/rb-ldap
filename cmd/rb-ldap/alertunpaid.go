package main

import (
	"github.com/redbrick/rb-ldap/internal/pkg/rbldap"
	"github.com/urfave/cli"
)

var (
	alertUnpaid = cli.Command{
		Action:   rbldap.AlterUnpaidUsers,
		Category: "Batch Commands",
		Name:     "alert-unpaid",
		Usage:    "Alert all unpaid users that their accounts will be disabled",
	}

	deleteUnpaid = cli.Command{
		Action:   rbldap.DeleteUnpaidUsers,
		Category: "Batch Commands",
		Name:     "delete-unpaid",
		Usage:    "Delete all unpaid users accounts that are outside their grace period",
	}

	disableUnpaid = cli.Command{
		Action:   rbldap.DisableUnpaidUsers,
		Category: "Batch Commands",
		Name:     "disable-unpaid",
		Usage:    "Diable all unpaid users accounts",
	}
)
