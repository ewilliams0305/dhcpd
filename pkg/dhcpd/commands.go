package dhcpd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// Configures the dhcp scope for the dhcp server
// The dhcp server will restart when this command is invoked
func configureScope(subnet SubnetDeclaration) tea.Cmd {

	return func() tea.Msg {
		return true
	}
}

// Saves a dhcp lease as as static host
// DHCP leases should be stored outside of the dhcp scopes start/end DUH
func reserveLease(lease Lease) tea.Cmd {

	return func() tea.Msg {
		return true
	}
}

func readConfiguration() tea.Msg {
	//return DhcpConfig{}
	return fmt.Errorf("TESTING FAILED BANNER RETURNED FROM COMMAND")
}

func viewCurrentLeases() tea.Msg {
	return make([]Lease, 0)
}

func openJournal() tea.Cmd {

	c := exec.Command("journalctl", "-u", "virtualcontrol.service", "-f")
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return journalClosedMessage{err}
	})
}
func stopService() tea.Cmd {

	c := exec.Command("systemctl", "stop", "virtualcontrol.service")
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return serviceClosedMessage{err}
	})
}
func startService() tea.Cmd {

	c := exec.Command("systemctl", "start", "virtualcontrol.service")
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return serviceClosedMessage{err}
	})
}
func restartService() tea.Cmd {

	c := exec.Command("systemctl", "restart", "virtualcontrol.service")
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return serviceClosedMessage{err}
	})
}
