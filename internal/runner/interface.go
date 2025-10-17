// Package runner
package runner

import "golang.org/x/crypto/ssh"

type CommandRunner interface {
	LoadCommand(path string) (map[string]string, error)
	RunCommand(client *ssh.Client, cmd string) (string, error)
}

type SSHConnector interface {
	GetSigner() (ssh.Signer, error)
	NewConfig(signer ssh.Signer) (*ssh.ClientConfig, error)
	Connect(cfg *ssh.ClientConfig) (*ssh.Client, error)
}
