package main

import (
	"fmt"
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"log"
	"os"
	"os/signal"
)

type RoomInviterConfig struct {
	password string
	room     string
	welcome  string
	rule     string
	repeat   string
}

func RoomInviter(config RoomInviterConfig) *wechaty.Plugin {
	bot := wechaty.NewPlugin()
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

	bot.Use(FriendshipAccepter(FriendshipAcceptedConfig{
		greeting: "hii",
		keyword:  "owo",
	}))

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
