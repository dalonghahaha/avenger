package tests

import (
	"fmt"
	"testing"

	"github.com/dalonghahaha/avenger/tools/ssh"
)

func TestSSH(t *testing.T) {
	addr := "114.215.106.133:22"
	username := "root"
	keyfile := "/Users/dengjialong/.ssh/id_rsa"
	client, err := ssh.DialWithKey(addr, username, keyfile)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()
	out, err := client.Cmd("ls /temp/").Output()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(out))
}
