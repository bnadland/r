package main

import (
	"fmt"
	"github.com/gizak/termui"
	"github.com/jzelinskie/geddit"
)

func isActiveCursor(c int, i int) string {
	if c == i {
		return "*"
	}
	return " "
}

func showSubreddit(subredditName string) error {
	r := geddit.NewSession("r by /u/bnadland")

	submissions, err := r.SubredditSubmissions(subredditName, geddit.HotSubmissions, geddit.ListingOptions{})
	if err != nil {
		return err
	}

	isActive := true

	cursor := 3
	for isActive {

		entries := []string{}
		for i, submission := range submissions {
			entries = append(entries, fmt.Sprintf("%s %s", isActiveCursor(cursor, i), submission.Title))
		}

		ls := termui.NewList()
		ls.Items = entries
		ls.ItemFgColor = termui.ColorDefault
		ls.Border.Label = fmt.Sprintf("Subreddit: %s", subredditName)
		ls.Height = termui.TermHeight()
		ls.Width = termui.TermWidth()
		ls.Y = 0
		termui.Render(ls)
		event := <-termui.EventCh()
		if event.Type == termui.EventKey {
			switch event.Key {
			case termui.KeyArrowLeft:
				isActive = false
			case termui.KeyArrowDown:
				cursor = cursor + 1
				if cursor > len(submissions) {
					cursor = len(submissions)
				}
			case termui.KeyArrowUp:
				cursor = cursor - 1
				if cursor < 0 {
					cursor = 0
				}
			}
		}
	}
	return nil
}

func main() {
	err := termui.Init()
	if err != nil {
		panic(err)
	}
	defer termui.Close()

	showSubreddit("golang")
}
