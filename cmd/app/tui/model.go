package tui

import (
	"fmt"
	"time"

	"steam-lancache/internal/api"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

type tickTraffic time.Time
type tab int

const (
	tabTrafficStats tab = iota
	tabSubscription
	tabLogs
)

type model struct {
	activeTab tab

	trafficTable table.Model
	lastTraffic  map[string]int64

	subscriptionData string
}

func UIModel(subscriptionInfo api.SubscriptionDetails) model {
	// ----- Traffic -----
	columns := []table.Column{
		{Title: "Service Name", Width: 16},
		{Title: "Usage", Width: 12},
		{Title: "Speed", Width: 12},
	}
	trafficTable := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
		table.WithHeight(30),
	)
	trafficTable.SetStyles(table.Styles{
		Cell:   trafficStatsCellStyle,
		Header: trafficStatsHeaderStyle,
	})

	// ----- Subscription -----
	subData := fmt.Sprintf("Subscription ID: %d\n", subscriptionInfo.SubscriptionID)
	subData += fmt.Sprintf("User IP: %s\n", subscriptionInfo.UserIP)
	subData += fmt.Sprintf("Ends In: %s\n", subscriptionInfo.End)
	subData += fmt.Sprintf("Status: %s\n", subscriptionInfo.Status)
	subData += fmt.Sprintf("Upstream IP: %s\n", subscriptionInfo.UpstreamIP)

	return model{
		activeTab:        tabTrafficStats,
		trafficTable:     trafficTable,
		subscriptionData: subData,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(trafficUpdateTick(500))
}

func trafficUpdateTick(tickRate int) tea.Cmd {
	return tea.Tick(time.Duration(tickRate)*time.Millisecond, func(t time.Time) tea.Msg {
		return tickTraffic(t)
	})
}
