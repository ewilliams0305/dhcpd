package dhcpd

import tea "github.com/charmbracelet/bubbletea"

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
	return DhcpConfig{}
}

func viewCurrentLeases() tea.Msg {
	return make([]Lease, 0)
}

func stopDhcpServer() tea.Msg {
	return false
}

func startDhcpServer() tea.Msg {
	return false
}

func restartDhcpServer() tea.Msg {
	return false
}
