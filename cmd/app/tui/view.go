package tui

import (
	"github.com/charmbracelet/lipgloss"
)

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.renderTabs(),
		m.renderContent(),
	)
}

func (m model) renderTabs() string {
	trafficStatsTab := tabInactiveStyle.Render("Traffic Stats")
	subscriptionTab := tabInactiveStyle.Render("Subsciption")

	switch m.activeTab {
	case tabTrafficStats:
		trafficStatsTab = tabActiveStyle.Render("Traffic Stats")
	case tabSubscription:
		subscriptionTab = tabActiveStyle.Render("Subsciption")
	}

	return tabBarStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Left, trafficStatsTab, subscriptionTab),
	)
}

func (m model) renderContent() string {
	switch m.activeTab {
	case tabTrafficStats:
		return boxStyle.Render(m.trafficTable.View())
	case tabSubscription:
		return boxStyle.Render(subscriptionDataHeader.Render("Subscription Data") + "\n" + m.subscriptionData)
	}
	return "unknown tab"
}
