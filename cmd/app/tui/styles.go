package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// ---------- UI styles ----------
var (
	// Global
	tabActiveStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#ffffff")).
			Background(lipgloss.Color("#0051ffff")).
			Padding(0, 2)

	tabInactiveStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#aaa")).
				Background(lipgloss.Color("#333")).
				Padding(0, 2)

	tabBarStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#242424ff")).
			Padding(0, 5).
			Align(lipgloss.Center).
			Width(52)

	boxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			BorderForeground(lipgloss.Color("#0051ffff")).
			Width(50)

	// Traffic Stats
	trafficStatsHeaderStyle = lipgloss.NewStyle().
				Background(lipgloss.Color("#0051ffff")).
				Align(lipgloss.Center).
				MarginBottom(1).
				Bold(true)

	trafficStatsCellStyle = lipgloss.NewStyle().
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("240")).
				BorderBottom(true).
				Bold(false)

	// Subscription Data
	subscriptionDataHeader = lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FAFAFA")).
				Width(48).
				Align(lipgloss.Center)
)
