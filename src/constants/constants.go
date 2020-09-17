package constants

import (
	"time"

	"github.com/hirochachacha/go-smb2"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

//Connection Connection Struct
type Connection struct {
	SSH            *ssh.Client
	SFTP           *sftp.Client
	SMB            *smb2.Session
	LastConnection time.Time
}

//ActiveConnections Active Connections
var ActiveConnections map[string]Connection

//CloseAllConnections CloseAllConnections
func CloseAllConnections(obj Connection) {
	if obj.SSH != nil {
		obj.SSH.Close()
	}

	if obj.SFTP != nil {
		obj.SFTP.Close()
	}

	if obj.SMB != nil {
		obj.SMB.Logoff()
	}
}