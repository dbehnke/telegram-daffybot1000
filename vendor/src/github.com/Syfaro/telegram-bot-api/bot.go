// Package tgbotapi has bindings for interacting with the Telegram Bot API.
package tgbotapi

// BotAPI has methods for interacting with all of Telegram's Bot API endpoints.
type BotAPI struct {
	Token   string      `json:"token"`
	Debug   bool        `json:"debug"`
	Self    User        `json:"-"`
	Updates chan Update `json:"-"`
}

// NewBotAPI creates a new BotAPI instance.
// Requires a token, provided by @BotFather on Telegram
func NewBotAPI(token string) (*BotAPI, error) {
	bot := &BotAPI{
		Token: token,
	}

	self, err := bot.GetMe()
	if err != nil {
		return &BotAPI{}, err
	}

	bot.Self = self

	return bot, nil
}
