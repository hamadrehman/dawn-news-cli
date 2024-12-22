package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/hamadrehman/dawn-news-cli/feed"
	"github.com/hamadrehman/dawn-news-cli/ui"
)

func main() {

	news_feed := feed.GetNews()
	newsList := ui.CreateNewsList(news_feed)

	m := ui.Model{
		List: newsList,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
