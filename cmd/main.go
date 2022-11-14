package main

import (
	"fmt"
	"strings"

	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
)

type movieBot struct {
	bot reddit.Bot
}

func (r *movieBot) Post(p *reddit.Post) error {
	if strings.Contains(p.Title, "movie") {
		var comms []*reddit.Comment
		comms = p.Replies
		for _, comm := range comms {
			fmt.Print(comm)
		}
		fmt.Printf("[%s] posted [%s]\n", p.Author, p.Title)
		return nil
	} else {
		return nil
	}
}

func main() {
	bot, err := reddit.NewBotFromAgentFile("../config/moviefinder.agent", 0)

	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return
	}

	cfg := graw.Config{Subreddits: []string{"askreddit"}}
	handler := &movieBot{bot: bot}
	if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
		fmt.Println("Failed to start graw run: ", err)
	} else {
		fmt.Println("graw run failed: ", wait())
	}

}
