package main

import (
        "fmt"
        "net"
        "os"
        "io/ioutil"
        "bufio"
        "strings"
        "time"
)

type Bot struct {
        server      string
        port        string
        name        string
        channel     string
        autoMsg     string
        conn        net.Conn
}

/*
Creates our new Bot
*/
func NewBot() *Bot {
        return &Bot {
                server:     "irc.chat.twitch.tv",
                port:       "6667",
                name:       "trofiebot",
                channel:    "#3ygun",
                autoMsg:    "Hit me with some pasta",
                conn:       nil,
        }
}

/*
Connects to the chatroom
*/
func (bot *Bot) Connect() {
    var err error
    fmt.Printf("Connecting to %s channel\n", bot.channel)
    bot.conn, err = net.Dial("tcp", bot.server+":"+bot.port)
    if err != nil {
        fmt.Printf("Cannot connect to channel, retrying")
        bot.Connect()
    }
    fmt.Printf("Connected to IRC server %s\n", bot.server)
}

/*
Sends a message to general chat
*/
func (bot *Bot) Message(message string) {
	if message == "" {
		return
	}
    fmt.Println("printing message: %s", message)
	fmt.Fprintf(bot.conn, "PRIVMSG "+bot.channel+" :"+message+"\r\n")
}

func (bot *Bot) ConsoleInput() {
	reader := bufio.NewReader(os.Stdin)
	for {
		text, _ := reader.ReadString('\n')
		if text == "/quit" {
			bot.conn.Close()
			os.Exit(0)
		}
		if text != "" {
			bot.Message(text)
		}
	}
}

/*
Sends timeout duration to bot.Ban
*/
func (bot *Bot) Timeout(user, reason, duration) {
    if duration == 0 {
        return
    }
    go bot.Ban(user, reason, duration)
}

/*
Ban's the user from twitch chat
If duration sent is < 0, permanently ban
*/
func (bot *Bot) Ban(user, reason, duration) {
    if duration == -1 {
        msg := fmt.Fprintf(".ban %s for %s", user, reason)
    }
    msg := fmt.Fprintf(".timeout %s for %s", user, reason)
    time.Sleep(1 * time.Second)
    bot.Message(msg)
}

func main() {
    ircbot := NewBot()
    go ircbot.ConsoleInput()
    ircbot.Connect()

    pass1, err := ioutil.ReadFile("twitch_pass.txt")
    fmt.Printf("$$$$ %s\n", string(pass1))
	pass := strings.Replace(string(pass1), "\n", "", 0)
	if err != nil {
		fmt.Println("Error reading from twitch_pass.txt.  Maybe it isn't created?")
		os.Exit(1)
	}

    fmt.Fprintf(ircbot.conn, "PASS %s\r\n", pass)
    fmt.Fprintf(ircbot.conn, "NICK %s\r\n", ircbot.name)
	fmt.Fprintf(ircbot.conn, "JOIN %s\r\n", ircbot.channel)

    fmt.Printf("Inserted information to server...\n")
	fmt.Printf("If you don't see the stream chat it probably means the Twitch oAuth password is wrong\n")
	fmt.Printf("Channel: " + ircbot.channel + "\n")
	defer ircbot.conn.Close()

    for {
        ircbot.Message(ircbot.autoMsg)
        time.Sleep(30 * time.Second)
    }
}
