package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/turnage/graw/reddit"
)

type movieBot struct {
	bot reddit.Bot
}

func (r *movieBot) Post(p *reddit.Post) error {
	if strings.Contains(p.SelfText, "testing movie") {
		<-time.After(10 * time.Second)
		return r.bot.SendMessage(
			p.Author,
			fmt.Sprintf("Reminder: %s", p.Title),
			"You've Been reminded!",
		)
	}
	return nil
}

func main() {
	bot, err := reddit.NewBotFromAgentFile("/home/zach/goProjects/MovieFinder/src/config/moviefinder.agent", 0)

	if err != nil {
		fmt.Println("Failed to create bot handle: ", err)
		return

	}

	harvest, err := bot.Listing("/r/golang", "")
	if err != nil {
		fmt.Println("Failed to getch /r/golang: ", err)
		return
	}
	for _, post := range harvest.Posts[:5] {
		fmt.Printf("[%s] posted [%s]\n", post.Author, post.Title)
	}
}
