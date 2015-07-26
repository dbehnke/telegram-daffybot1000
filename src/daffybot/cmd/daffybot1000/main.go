package main

import (
	"log"
	"strings"
	"sync"
	"time"

	"github.com/Syfaro/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("YourAPI")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	var allchats struct {
		sync.Mutex
		chats map[int]string
	}

	allchats.chats = make(map[int]string)
	log.Printf("Authorized on account %s", bot.Self.UserName)

	go func() {
		for {
			log.Println("Brodcast Loop SLEEP 60 seconds")
			time.Sleep(60 * time.Second)
			log.Println("Broadcast Loop AWAKE")
			func() {
				allchats.Lock()
				defer allchats.Unlock()
				log.Printf("Broadcast Loop allchats size = %d", len(allchats.chats))
				for id, username := range allchats.chats {
					log.Printf("Broadcast Loop sending message to %d - %s", id, username)
					msg := tgbotapi.NewMessage(id, "Broadcast Test")
					bot.SendMessage(msg)
				}
			}()
		}
	}()

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.UpdatesChan(u)

	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		if strings.Contains(update.Message.Text, "/subscribe") {
			func() {
				allchats.Lock()
				defer allchats.Unlock()
				allchats.chats[update.Message.Chat.ID] = update.Message.From.UserName
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "subscribed")
				bot.SendMessage(msg)
				log.Printf("[%s] subscribed", update.Message.From.UserName)
			}()
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			msg.ReplyToMessageID = update.Message.MessageID

			bot.SendMessage(msg)
		}
	}
}
