package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/hamadrehman/dawn-news-cli/feed"
)

var docStyle = lipgloss.NewStyle()

type Model struct {
	List list.Model
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.List.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.List, cmd = m.List.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return docStyle.Render(m.List.View())
}

type Item struct {
	title, desc string
}

func (i Item) Title() string       { return i.title }
func (i Item) Description() string { return i.desc }
func (i Item) FilterValue() string { return i.title }

func (m Model) Init() tea.Cmd { return nil }

func CreateNewsList(news_feed feed.Feed) list.Model {
	items := make([]list.Item, len(news_feed.Channel.Items))

	d := list.NewDefaultDelegate()
	newsList := list.New(items, d, 0, 0)
	newsList.Title = "DAWN NEWS"

	for idx, newsstr := range news_feed.Channel.Items {
		items[idx] = Item{
			title: strings.TrimSpace(newsstr.Title),
			desc:  "",
		}
	}
	return newsList

}
