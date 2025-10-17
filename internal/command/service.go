// Package command
package command

import "golang.org/x/crypto/ssh"

type CommandService struct{}

func (c *CommandService) LoadCommand(path string) (map[string]string, error) {
	return LoadCommand(path)
}

func (c *CommandService) RunCommand(client *ssh.Client, cmd string) (string, error) {
	return RunCommand(client, cmd)
}
