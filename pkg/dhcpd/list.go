package dhcpd

import "github.com/charmbracelet/bubbles/list"

type listOption struct {
	title, desc string
}

func (i listOption) Title() string       { return i.title }
func (i listOption) Description() string { return i.desc }
func (i listOption) FilterValue() string { return i.title }

func NewMainMenu() list.Model {
	items := []list.Item{
		listOption{title: "Subnet", desc: "configure the dhcp server subnet"},
		listOption{title: "Reservation", desc: "create or edit dhcp reservations"},
		listOption{title: "Status", desc: "displays the current dhcpd service"},
		listOption{title: "Start Server", desc: "starts the dhcpd service"},
		listOption{title: "Stop Server", desc: "stops the dhcpd service"},
		listOption{title: "Restart Server", desc: "restarts the dhcpd service"},
		listOption{title: "Enable DHCP", desc: "enables the dhcpd service"},
		listOption{title: "Disable DHCP", desc: "disables the dhcpd service"},
	}

	list := list.New(items, list.NewDefaultDelegate(), 100, 22)
	list.Title = "Select an option below to conifgure your dhcpd service"
	return list
}
