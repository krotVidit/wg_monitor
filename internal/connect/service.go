// Package connect
package connect

import (
	"wg-monitor/app/internal/domain"

	"golang.org/x/crypto/ssh"
)

type SSHService struct {
	Cfg *domain.Config
}

func (s *SSHService) GetSigner() (ssh.Signer, error) {
	return GetSigner(s.Cfg)
}

func (s *SSHService) NewConfig(signer ssh.Signer) (*ssh.ClientConfig, error) {
	return NewSSHConfig(signer, s.Cfg)
}

func (s *SSHService) Connect(cfg *ssh.ClientConfig) (*ssh.Client, error) {
	return ConnectSSH(s.Cfg, cfg)
}
