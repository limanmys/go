package connector

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"github.com/hirochachacha/go-smb2"
)

//OpenSMBConnection OpenSMBConnection
func OpenSMBConnection(ipAddress string, username string, password string) (*smb2.Session, error) {
	conn, err := net.Dial("tcp", ipAddress+":445")
	if err != nil {
		return nil, err
	}

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     username,
			Password: password,
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		return nil, err
	}
	return s, nil
}

//PutFileSMB PutFileSMB
func PutFileSMB(session *smb2.Session, localPath string, remotePath string) bool {
	fs, err := session.Mount("C$")
	if err != nil {
		panic(err)
	}
	defer fs.Umount()

	f, err := fs.Create(filepath.Base(remotePath))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	srcFile, err := os.Open(localPath)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	_, err = io.Copy(f, srcFile)
	if err != nil {
		return false
	}
	return true
}

//GetFileSMB GetFileSMB
func GetFileSMB(session *smb2.Session, localPath string, remotePath string) bool {
	fs, err := session.Mount("C$")
	if err != nil {
		panic(err)
	}
	defer fs.Umount()

	f, err := fs.Open(remotePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = os.Stat(localPath)
	if os.IsNotExist(err) {
		os.Create(localPath)
	}

	srcFile, err := os.OpenFile(localPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer srcFile.Close()

	_, err = io.Copy(srcFile, f)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}