package rbuser

import (
	"fmt"
	"html/template"
	"os"
	"time"
)

// User is a user in RedBrick's LDAP.
type User struct {
	UID              string
	UserType         string
	ObjectClass      []string
	Newbie           bool   // New this year
	CN               string // Full name
	Altmail          string // Alternate email
	ID               int    // DCU ID number
	Course           string // DCU course code
	Year             int    // DCU course year number/code
	YearsPaid        int    // Number of years paid (integer)
	UpdatedBy        string // Username of user last to update
	Updated          time.Time
	CreatedBy        string // Username of user that created them
	Created          time.Time
	Birthday         time.Time
	UIDNumber        int
	GidNumber        int
	Gecos            string
	LoginShell       string
	HomeDirectory    string
	UserPassword     string   // Crypted password.
	Host             []string // List of hosts.
	ShadowLastChange int
}

// Vhost returns the user's Apache macro template.
func (u *User) Vhost() string {
	return fmt.Sprintf("use VHost /storage/webtree/%s/%s %s %s %s", string([]rune(u.UID)[0]), u.UID, u.UID, u.UserType, u.UID)
}

// PrettyPrint prints a user's information to stdout.
func (u *User) PrettyPrint() error {
	const output = `User Information
================
{{ with .UID }}uid: {{ . }}
{{ end }}{{ with .UserType }}usertype: {{ . }}
{{ end }}{{ with .ObjectClass }}objectClass: {{ . }}
{{ end }}{{ with .Newbie }}newbie: {{ . }}
{{ end }}{{ with .CN }}cn: {{ . }}
{{ end }}{{ with .Altmail }}altmail: {{ . }}
{{ end }}{{ with .ID }}id: {{ . }}
{{ end }}{{ with .Course }}course: {{ . }}
{{ end }}{{ with .Year }}year: {{ . }}
{{ end }}{{ with .YearsPaid }}yearsPaid: {{ . }}
{{ end }}{{ with .UpdatedBy }}updatedBy: {{ . }}
{{ end }}{{ with .Updated.Format "2006-01-02 15:04:05" }}updated: {{ . }}
{{ end }}{{ with .CreatedBy }}createdby: {{ . }}
{{ end }}{{ with .Created.Format "2006-01-02 15:04:05" }}created: {{ . }}
{{ end }}{{ with .Birthday.Format "2006-01-02 15:04:05" }}birthday: {{ . }}
{{ end }}{{ with .UIDNumber }}uidNumber: {{ . }}
{{ end }}{{ with .GidNumber }}gidNumber: {{ . }}
{{ end }}{{ with .Gecos }}gecos: {{ . }}
{{ end }}{{ with .LoginShell }}loginShell: {{ . }}
{{ end }}{{ with .HomeDirectory }}homeDirectory: {{ . }}
{{ end }}{{ with .UserPassword }}userPassword: {{ . }}
{{ end }}{{ with .Host }}host: {{ . }}
{{ end }}{{ with .ShadowLastChange }}shadowLastChange: {{ . }}{{ end }}
`

	t := template.Must(template.New("user").Parse(output))
	return t.Execute(os.Stdout, u)
}

// CreateHome creates a user's home directory and makes them the owner.
func (u *User) CreateHome() error {
	if err := os.MkdirAll(u.HomeDirectory, os.ModePerm); err != nil {
		return err
	}
	file, err := os.Create(fmt.Sprintf("%s/.forward", u.HomeDirectory))
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.WriteString(fmt.Sprintf("%s\n", u.Altmail)); err != nil {
		return err
	}
	if err := file.Sync(); err != nil {
		return err
	}
	return os.Chown(u.HomeDirectory, u.UIDNumber, u.GidNumber)
}

// CreateHome creates a user's web directory and makes them the owner.
func (u *User) CreateWebDir() error {
	folder := fmt.Sprintf("/webtree/%d/%s", []rune(u.UID)[0], u.UID)
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	return os.Chown(folder, u.UIDNumber, u.GidNumber)
}

// LinkPublicHTML link's a user's web directory to public_html in their home directory.
func (u *User) LinkPublicHTML() error {
	return os.Symlink(fmt.Sprintf("/webtree/%d/%s", []rune(u.UID)[0], u.UID), fmt.Sprintf("%s/public_html", u.HomeDirectory))
}

// MigrateHome migrates a user's home directory and makes them the owner.
func (u *User) MigrateHome(newHome string) error {
	if err := os.Rename(u.HomeDirectory, newHome); err != nil {
		return err
	}
	u.HomeDirectory = newHome
	return u.LinkPublicHTML()
}

// DelWebDir deletes a user's web directory.
func (u *User) DelWebDir() error {
	return os.RemoveAll(fmt.Sprintf("/webtree/%d/%s", []rune(u.UID)[0], u.UID))
}

// DelHomeDir deletes a user's home directory.
func (u *User) DelHomeDir() error {
	return os.RemoveAll(u.HomeDirectory)
}

// DelExtraFiles deletes any leftover files owned by a user.
func (u *User) DelExtraFiles() error {
	for _, file := range []string{
		"/local/share/agreement/statedir/%s",
		"/var/mail/%s",
		"/var/spool/cron/crontabs/%s",
	} {
		if err := os.Remove(fmt.Sprintf(file, u.UID)); err != nil {
			return err
		}
	}
	return nil
}