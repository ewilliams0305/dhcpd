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
		listOption{title: "Status", desc: "displays the current system status"},
		listOption{title: "Logs", desc: "view dhcp server logs and realtime information"},
	}

	return list.New(items, list.NewDefaultDelegate(), 100, 20)
}
