package dhcpd

import (
	"fmt"
	"os/exec"

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

func viewLogs() tea.Cmd {
	c := exec.Command("journalctl", "-u", "dhcpd.service", "-f")
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return err
	})
}

// The dhcp service command is a tea.Cmd used to manipulate the state of the linux dhcpd.service
func dhcpdServiceCmd(a Action) tea.Cmd {
	c := exec.Command("systemctl", a.String(), "dhcpd.service")
	return tea.ExecProcess(c, func(err error) tea.Msg {
		return a.Result(err)
	})
}

func serviceStatusCmd() tea.Msg {
	status, err := queryStatus()
	if err != nil {
		return err
	}

	return status
}
