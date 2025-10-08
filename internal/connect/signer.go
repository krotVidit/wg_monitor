// Package connect реализует подключение через ssh
package connect

import (
	"os"
	"wg-monitor/app/internal/domain"

	"golang.org/x/crypto/ssh"
)

func GetSigner(cfg *domain.Config) (ssh.Signer, error) {
	key, err := os.ReadFile(cfg.PathKey)
	if err != nil {
		return nil, err
	}

	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	return signer, nil
}
