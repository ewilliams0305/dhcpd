package dhcpd 

// Configures the dhcp scope for the dhcp server
// The dhcp server will restart when this command is invoked
func configureScope(scope DhcpScope) tea.Cmd {

  return func () tea.Msg {

  }
}

// Saves a dhcp lease as as static host
// DHCP leases should be stored outside of the dhcp scopes start/end DUH
func reserveLease(lease DhcpLease) tea.Cmd {

  return func () tea.Msg {

  }
}

func viewCurrentLeases() tea.Msg {
  return make(DhcpLease, 0)
}

func stopDhcpServer() tea.Msg {

}

func startDhcpServer() tea.Msg {

}

func restartDhcpServer() tea.Msg {

}

