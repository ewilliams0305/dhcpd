package dhcpd

import (
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"golang.org/x/term"
)

/********************************************************
*
* MAIN APPLICATION VIEW MODEL
*
*********************************************************/

type DhcpdViewModel struct {
	baner BannerModel
	list  list.Model
	err   error
	help  HelpModel
	size  tea.WindowSizeMsg
}

func InitialModel() *DhcpdViewModel {

	w, h, _ := term.GetSize(int(os.Stdout.Fd()))

	app = &DhcpdViewModel{
		baner: NewBanner("CONFIGURE YOUR DHCP SERVER", BannerNormalState, w),
		list:  NewMainMenu(),
		help:  NewHelpModel(),
		size: tea.WindowSizeMsg{
			Width:  w,
			Height: h,
		},
	}
	return app
}

// Init implements tea.Model.
func (*DhcpdViewModel) Init() tea.Cmd {
	return tea.Batch(tick, readConfiguration)
}

// Update implements tea.Model.
func (m *DhcpdViewModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		}

	case tickMsg:
		w, h, _ := term.GetSize(int(os.Stdout.Fd()))
		if w != m.size.Width || h != m.size.Height {
			m.updateSize(w, h)
		}
		return m, tea.Batch(tick, func() tea.Msg { return tea.WindowSizeMsg{Width: w, Height: h} })

	case tea.WindowSizeMsg:
		m.updateSize(msg.Width, msg.Height)
		return m, nil

	case error:

	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View implements tea.Model.
func (m *DhcpdViewModel) View() string {
	s := m.baner.View()
	s += docStyle.Render(m.list.View())
	s += "\n\n\n"
	s += m.help.renderHelpInfo()
	return s
}

/********************************************************
*
* INITIALIZE THE APP WITH FLAGS
*
*********************************************************/

func Run() {

	m := InitialModel()
	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("DHCPD CLI failed to start, there's been an error: %v", err)
		os.Exit(1)
	}
}

/********************************************************
*
* APPLICATION POLLING LOOP
*
*********************************************************/

// Pointless type to trigger the update function
type tickMsg int

// Updates the entire view if the size changed
func (m *DhcpdViewModel) updateSize(w, h int) {
	m.size.Width = w
	m.size.Height = h
}

// Sends a message back to the update function to start the tick over again.
func tick() tea.Msg {
	time.Sleep(time.Second + 1)
	return tickMsg(1)
}
