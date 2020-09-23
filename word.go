package main

import (
	"fmt"
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"log"
	"os"
	"os/signal"
	"sync"
)

type WordCounterConfig struct {
	hours int
}

type contactIndex struct {
	room string
	name string
}

var mux sync.Mutex

func WordCounter(config WordCounterConfig) *wechaty.Plugin {
	bot := wechaty.NewPlugin()
	bot.OnMessage(func(context *wechaty.Context, message *user.Message) {
		if message.Type() == schemas.MessageTypeText {
			index := contactIndex{
				room: message.Room().ID(),
				name: message.From().Name(),
			}
			count := len([]rune(message.Text()))
			fmt.Println("OnMessage", index, count)

			mux.Lock()

			mux.Unlock()
		}
	})
	return bot
}

func main() {
	var bot = wechaty.NewWechaty()

	bot.OnScan(func(context *wechaty.Context, qrCode string, status schemas.ScanStatus, data string) {
		fmt.Printf("Scan QR Code to login: %v\nhttps://wechaty.github.io/qrcode/%s\n", status, qrCode)
	}).OnLogin(func(context *wechaty.Context, user *user.ContactSelf) {
		fmt.Printf("User %s logined\n", user.Name())
	}).OnLogout(func(context *wechaty.Context, user *user.ContactSelf, reason string) {
		fmt.Printf("User %s logouted: %s\n", user, reason)
	})

	bot.Use(WordCounter(WordCounterConfig{hours: 6}))

	var err = bot.Start()
	if err != nil {
		panic(err)
	}

	var quitSig = make(chan os.Signal)
	signal.Notify(quitSig, os.Interrupt, os.Kill)

	select {
	case <-quitSig:
		log.Fatal("exit.by.signal")
	}
}
