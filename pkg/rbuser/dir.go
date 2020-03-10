package rbuser

import (
	"fmt"
	"os"
)

// CreateHome creates a user's home directory and makes them the owner.
func (user *User) CreateHome() error {
	if err := os.MkdirAll(user.HomeDirectory, os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s/.forward", user.HomeDirectory))
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(fmt.Sprintf("%s\n", user.Altmail)); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	return os.Chown(user.HomeDirectory, user.UIDNumber, user.GidNumber)
}

// CreateHome creates a user's web directory and makes them the owner.
func (user *User) CreateWebDir() error {
	folder := fmt.Sprintf("/webtree/%d/%s", []rune(user.UID)[0], user.UID)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	return os.Chown(folder, user.UIDNumber, user.GidNumber)
}

// LinkPublicHTML link's a user's web directory to public_html in their home directory.
func (user *User) LinkPublicHTML() error {
	return os.Symlink(fmt.Sprintf("/webtree/%d/%s", []rune(user.UID)[0], user.UID), fmt.Sprintf("%s/public_html", user.HomeDirectory))
}

// MigrateHome migrates a user's home directory and makes them the owner.
func (user *User) MigrateHome(newHome string) error {
	if err := os.Rename(user.HomeDirectory, newHome); err != nil {
		return err
	}
	user.HomeDirectory = newHome
	return user.LinkPublicHTML()
}

// DelWebDir deletes a user's web directory.
func (user *User) DelWebDir() error {
	return os.RemoveAll(fmt.Sprintf("/webtree/%d/%s", []rune(user.UID)[0], user.UID))
}

// DelHomeDir deletes a user's home directory.
func (user *User) DelHomeDir() error {
	return os.RemoveAll(user.HomeDirectory)
}

// DelExtraFiles deletes any leftover files owned by a user.
func (user *User) DelExtraFiles() error {
	for _, file := range []string{
		"/local/share/agreement/statedir/%s",
		"/var/mail/%s",
		"/var/spool/cron/crontabs/%s",
	} {
		if err := os.Remove(fmt.Sprintf(file, user.UID)); err != nil {
			return err
		}
	}
	return nil
}
