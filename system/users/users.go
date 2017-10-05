package users

import "os/user"

func GetUsername() string {
	u, e := user.Current()
	if e != nil {
		return ""
	}
	return u.Username
}
