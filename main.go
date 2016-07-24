package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/chelexa/trofiebot/bot"
)

func main() {

	//Parse command line arguments
	flag.Parse()
	command := flag.Arg(0)

	//TODO: Other functionality, i.e "help", etc.
	switch command {
	case "run":
		runBot()
	}
}

func runBot() {
	ircbot := bot.NewBot()
	go ircbot.ConsoleInput()
	ircbot.Connect()
	defer ircbot.Close()

	//authenticating w/ twitch auth token
	pass1, err := ioutil.ReadFile("twitch_pass.txt")
	if err != nil {
		fmt.Println("Error reading from twitch_pass.txt.  Maybe it isn't created?")
		os.Exit(1)
	}
	pass := strings.Replace(string(pass1), "\n", "", 0)
	fmt.Printf("The password used is: %s\r\n", string(pass))

	ircbot.LogIn(pass)
	go ircbot.AutoMessage()

	//run forever :)
	ircbot.Start()
}
