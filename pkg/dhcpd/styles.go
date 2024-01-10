package dhcpd

import "github.com/charmbracelet/lipgloss"

const (
	PrimaryColor string = "#3F51B5"
	PrimaryLight string = "#C5CAE9"
	PrimaryDark  string = "#001F5F"
	AccentColor  string = "#00796B"
	ErrorColor   string = "#8A0B29"
)

var (
	docStyle = lipgloss.NewStyle().Margin(1, 2)

	bannerNormalStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(PrimaryLight)).
				Background(lipgloss.Color(PrimaryDark)).
				PaddingTop(1).
				PaddingLeft(1).
				MarginBottom(1).
				Align(lipgloss.Center)

	bannerErrorStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color(PrimaryLight)).
				Background(lipgloss.Color(ErrorColor)).
				PaddingTop(1).
				PaddingLeft(1).
				MarginBottom(1).
				Align(lipgloss.Center)
)
