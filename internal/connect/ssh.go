// Package connect реализует подключение через ssh
package connect

import (
	"fmt"

	"wg-monitor/app/internal/domain"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

func NewSSHConfig(signer ssh.Signer, cfg *domain.Config) (*ssh.ClientConfig, error) {
	kh, err := knownhosts.New(cfg.KnowHost)
	if err != nil {
		return nil, fmt.Errorf("не удалось загрузить known_hosts: %w", err)
	}

	return &ssh.ClientConfig{
		User: cfg.User,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: kh,
	}, nil
}

func ConnectSSH(cfg *domain.Config, sshConfig *ssh.ClientConfig) (*ssh.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	client, err := ssh.Dial("tcp", addr, sshConfig)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения к %s: %w", addr, err)
	}
	return client, nil
}
