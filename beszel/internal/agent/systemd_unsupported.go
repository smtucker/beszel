//go:build !linux

package agent

import "beszel/internal/entities/system"

func getSystemdServices() []system.SystemdService {
	return nil
}
