package dns

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"steam-lancache/internal/config"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DNSList struct {
	Name        string `json:"name"`
	PersianName string `json:"persian_name"`
	IP          string `json:"ip"`
}

type dnsModel struct {
	servers []DNSList
	cursor  int
}

func (m dnsModel) Init() tea.Cmd {
	return nil
}

func (m dnsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.cursor = 0
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.servers)-1 {
				m.cursor++
			}
		case "enter":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m dnsModel) View() string {
	var s string

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("212"))

	s += titleStyle.Render("[?] Select a DNS Server") + "\n\n"

	for i, server := range m.servers {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		itemStyle := lipgloss.NewStyle()
		if m.cursor == i {
			itemStyle = itemStyle.
				Foreground(lipgloss.Color("46")).
				Bold(true)
		}

		item := fmt.Sprintf("%s %d. %s (%s)", cursor, i+1, server.Name, server.IP)
		s += itemStyle.Render(item) + "\n"
	}

	s += "\n" + lipgloss.NewStyle().Faint(true).Render("[↑/↓ or k/j to navigate, Enter to select, q to select first server]")

	return s
}

func FetchDNSServers() ([]DNSList, error) {
	resp, err := http.Get(config.DNSListAPI)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch DNS list: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch DNS list (status: %d)", resp.StatusCode)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read DNS response body: %w", err)
	}

	var servers []DNSList
	if err := json.Unmarshal(bodyBytes, &servers); err != nil {
		return nil, fmt.Errorf("failed to parse DNS JSON response: %w", err)
	}

	return servers, nil
}

func DNSSelection(servers []DNSList) DNSList {
	model := dnsModel{
		servers: servers,
		cursor:  0,
	}

	p := tea.NewProgram(model)
	finalModel, err := p.Run()
	if err != nil {
		// fallback to first server on error
		return servers[0]
	}

	m := finalModel.(dnsModel)
	if m.cursor >= 0 && m.cursor < len(servers) {
		return servers[m.cursor]
	}
	return servers[0]
}
