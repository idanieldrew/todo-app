package database

import "fmt"

type Conf struct {
	ServerName string
	User             string
	Pass    string
	Database     string
}

func Connection(c Conf) string {
	connect := fmt.Sprintf("%s:%s@/%s", c.User, c.Pass, c.Database)
	return connect
}
