// Package connect реализует подключение через ssh
package connect

import (
	"fmt"
	"wg-monitor/app/internal/domain"

	"golang.org/x/crypto/ssh"
)

func NewSSHConfig(signer ssh.Signer, cfg *domain.Config) *ssh.ClientConfig {
	sshConfig := &ssh.ClientConfig{
		User: cfg.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	return sshConfig
}

func ConnectSSH(cfg *domain.Config, sshConfig *ssh.ClientConfig) (*ssh.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, err
	}
	return client, nil
}
