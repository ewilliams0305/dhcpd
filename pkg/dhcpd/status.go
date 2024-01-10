package dhcpd

import (
	"os/exec"
	"strings"
)

type ServiceStatus struct {
	enabled bool
	running bool
}

func queryStatus() (*ServiceStatus, error) {

	cmd := exec.Command("systemctl", "status", "dhcpd.service")
	output, err := cmd.Output()
	if err != nil {
		serviceStatus.enabled = false
		serviceStatus.running = false
		return nil, err
	}

	data := string(output)

	if strings.Contains(data, "active (running)") {
		serviceStatus.running = true
	}

	if strings.Contains(data, "inactive (dead)") {
		serviceStatus.running = false
	}
	if strings.Contains(data, "dhcpd.service; enabled;") {
		serviceStatus.enabled = true
	}
	return serviceStatus, nil
}

func (s *ServiceStatus) View() string {
	if s == nil {
		return ""
	}

	if s.enabled && s.running {
		return "Service enabled and running"
	}

	if s.enabled && !s.running {
		return "Service enabled but stopped"
	}

	if !s.enabled && s.running {
		return "Service is disabled but running"
	}

	return "Service is disabled and stopped"
}
