package main

import (
	"fmt"
	"github.com/slack-go/slack"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"time"
)

// Глобальный рандомизатор
var RND = rand.New(rand.NewSource(time.Now().Unix()))

// Эмоджи слака
type SlackEmoji string

const (
	SlackEmojiVictory           = SlackEmoji(":v:")
	SlackEmojiHand              = SlackEmoji(":hand:")
	SlackEmojiOkHand            = SlackEmoji(":ok_hand:")
	SlackEmojiSpockHand         = SlackEmoji(":spock-hand:")
	SlackEmojiWave              = SlackEmoji(":wave:")
	SlackEmojiHandshake         = SlackEmoji(":handshake:")
	SlackEmojiFingersCrossed    = SlackEmoji(":hand_with_index_and_middle_fingers_crossed:")
	SlackEmojiOpenHands         = SlackEmoji(":open_hands:")
	SlackEmojiCallMeHand        = SlackEmoji(":call_me_hand:")
	SlackEmojiILoveYouHand      = SlackEmoji(":i_love_you_hand_sign:")
	SlackEmojiRaisedHandSplayed = SlackEmoji(":raised_hand_with_fingers_splayed:")

	SlackEmojiClock1  = SlackEmoji(":clock1:")
	SlackEmojiClock2  = SlackEmoji(":clock2:")
	SlackEmojiClock3  = SlackEmoji(":clock3:")
	SlackEmojiClock4  = SlackEmoji(":clock4:")
	SlackEmojiClock5  = SlackEmoji(":clock5:")
	SlackEmojiClock6  = SlackEmoji(":clock6:")
	SlackEmojiClock7  = SlackEmoji(":clock7:")
	SlackEmojiClock8  = SlackEmoji(":clock8:")
	SlackEmojiClock9  = SlackEmoji(":clock9:")
	SlackEmojiClock10 = SlackEmoji(":clock10:")
	SlackEmojiClock11 = SlackEmoji(":clock11:")
	SlackEmojiClock12 = SlackEmoji(":clock12:")
)

// Статус в слаке со случайным эмоджи и текстом
type SlackStatusCombo struct {
	// При установке будет выбран одно из эмоджи
	Emoji []SlackEmoji
	// При установке будет выбран один из текстов
	Texts []string
}

func main() {
	api := slack.New("xoxp-*")
	go RunSwitching(api)

	WaitForCancel()

	// Сброс статуса при выходе
	_ = SetStatus(api, "", "")
}

// Цикл смены комбинаций
func RunSwitching(api *slack.Client) {
	for {
		statusText, statusEmoji := GenerateCombo()
		if err := SetStatus(api, statusText, statusEmoji); err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 6)
	}
}

// Устанавливает статус и эмоджи
func SetStatus(api *slack.Client, text string, emoji SlackEmoji) error {
	for {
		fmt.Printf("[%s] New status: %s - %s\n", time.Now().Format(time.Stamp), emoji, text)
		err := api.SetUserCustomStatus(text, string(emoji), 0)
		if err != nil {
			if !strings.HasPrefix(err.Error(), "slack rate limit exceeded") {
				return err
			}
			fmt.Printf("Rate limit exceeded, waiting for 10 seconds ...\n")
			time.Sleep(time.Second * 10)
		} else {
			break
		}
	}
	return nil
}

// Генерирует новую комбинацию жеста и текста
func GenerateCombo() (string, SlackEmoji) {
	// Комбинации жестов и текстов
	combos := []SlackStatusCombo{
		//{
		//	Emoji: []SlackEmoji{SlackEmojiCallMeHand},
		//	Texts: []string{"Call Me Maybe"},
		//},
		//{
		//	Emoji: []SlackEmoji{SlackEmojiOkHand},
		//	Texts: []string{"All Is OK", "Fine", "Great"},
		//},
		{
			Emoji: []SlackEmoji{
				SlackEmojiCallMeHand,
				SlackEmojiOkHand,
				SlackEmojiVictory,
				SlackEmojiHand,
				SlackEmojiSpockHand,
				SlackEmojiWave,
				SlackEmojiHandshake,
				SlackEmojiFingersCrossed,
				SlackEmojiOpenHands,
				SlackEmojiILoveYouHand,
				SlackEmojiRaisedHandSplayed,
			},
			Texts: []string{"Status automated"},
		},
	}

	combo := combos[RND.Int()%len(combos)]
	emoji := combo.Emoji[RND.Int()%len(combo.Emoji)]
	text := ""
	if len(combo.Texts) > 0 {
		text = combo.Texts[RND.Int()%len(combo.Texts)]
	}

	return text, emoji
}

func WaitForCancel() {
	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
