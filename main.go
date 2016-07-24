package main

import (
        "flag"
        "fmt"
        "github.com/chelexa/trofiebot/bot"
        "io/ioutil"
        "os"
        "strings"
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
    fmt.Printf("$$$$ %s\n", string(pass1))
    pass := strings.Replace(string(pass1), "\n", "", 0)
    if err != nil {
        fmt.Println("Error reading from twitch_pass.txt.  Maybe it isn't created?")
        os.Exit(1)
    }

    ircbot.LogIn(pass)
    go ircbot.AutoMessage()
    ircbot.Message("A wild Kappa has appeared with CP 1")

    //run forever :)
    ircbot.HandleChat()
}
