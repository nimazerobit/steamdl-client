package tui

import (
	"fmt"
	"steam-lancache/internal/config"
	"steam-lancache/internal/stats"
	"strings"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	case tickTraffic:
		if m.activeTab == tabTrafficStats {

			// take snapshot from stats package
			snap := stats.Snapshot()

			if m.lastTraffic == nil {
				m.lastTraffic = snap
			}

			newRows := make([]table.Row, 0, len(snap))

			for _, service := range config.Domains.Order {
				bytes := snap[service]
				last := m.lastTraffic[service]

				delta := bytes - last
				speedMB := float64(delta) / 1_000_000
				totalGB := float64(bytes) / 1_000_000_000

				newRows = append(newRows, table.Row{
					strings.Title(service),
					fmt.Sprintf("%.2f GB", totalGB),
					fmt.Sprintf("%.2f MB/s", speedMB),
				})
			}

			m.lastTraffic = snap

			m.trafficTable.SetRows(newRows)
			m.trafficTable.SetHeight(len(newRows)*2 + 1)
		}
		cmds = append(cmds, trafficUpdateTick(1000)) // 1 second refresh

	case tea.KeyMsg:
		switch msg.String() {
		case "left", "shift+tab":
			m.activeTab = (m.activeTab - 1 + 2) % 2
		case "right", "tab":
			m.activeTab = (m.activeTab + 1) % 2
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}
