// Package command
package command

import (
	"log"

	"golang.org/x/crypto/ssh"
)

func RunCommand(client *ssh.Client, cmd string) (string, error) {
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer func() {
		if e := session.Close(); e != nil && e.Error() != "EOF" {
			log.Println("Ошибка закрытии сессии", e)
		}
	}()

	output, err := session.CombinedOutput(cmd)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
