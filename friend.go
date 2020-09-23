package main

import (
	"fmt"
	"github.com/wechaty/go-wechaty/wechaty"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"github.com/wechaty/go-wechaty/wechaty/user"
	"log"
	"os"
	"os/signal"
	"strings"
)

// FriendshipAcceptedConfig provides settings to accept friendship automatically.
// disable greeting by setting greeting to ""
// accept all invitations if keyword is set to ""
type FriendshipAcceptedConfig struct {
	// disable greeting by setting greeting to ""
	greeting string
	// accept all invitations if keyword is set to ""
	keyword string
}

func FriendshipAccepter(config FriendshipAcceptedConfig) *wechaty.Plugin {

	log.Println("plugin")

	bot := wechaty.NewPlugin()
	bot.OnFriendship(func(context *wechaty.Context, friendship *user.Friendship) {

		log.Println("request hello:", friendship.Hello())

		switch friendship.Type() {
		case schemas.FriendshipTypeReceive:
			helloMessage := friendship.Hello()
			if config.keyword != "" && !strings.Contains(helloMessage, config.keyword) {

				fmt.Println("not accept")

				break
			}
			err := friendship.Accept()
			if err != nil {
				log.Println("accept friendship error ", err)
			}
		case schemas.FriendshipTypeConfirm:
			// do greeting
			if config.greeting != "" {
				contact := friendship.Contact()
				_, err := contact.Say(config.greeting)
				if err != nil {
					log.Println("greeting error ", err)
				}
			}
		case schemas.FriendshipTypeVerify:
			// This is for when we send a message to others, but they did not accept us as a friend.
			break
		default:
			log.Printf("error: friendshipType unknown %v \n", friendship)
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
