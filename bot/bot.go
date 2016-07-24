package bot

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"
	"strings"
	"time"
    "github.com/chelexa/trofiebot/emotemon"
)

/*
Bot is a ircbot object
*/
type Bot struct {
	server  string
	port    string
	name    string
	channel string
	conn    net.Conn
    game    *emotemon.EmotemonGame
}

/*
NewBot creates a new Bot with the default parameters
*/
func NewBot() *Bot {
	return &Bot{
		server:  "irc.chat.twitch.tv",
		port:    "6667",
		name:    "trofiebot",
		channel: "#3ygun",
		conn:    nil,
	}
}

/*
Connect to the chatroom
*/
func (bot *Bot) Connect() {
	var err error
	fmt.Printf("Connecting to %s channel\n", bot.channel)
	bot.conn, err = net.Dial("tcp", bot.server+":"+bot.port)
	fmt.Printf("before %s\n", bot.channel)
	if err != nil {
		fmt.Printf("Cannot connect to channel, retrying")
		bot.Connect()
	}
	fmt.Printf("Connected to IRC server %s\n", bot.server)
}

/*
Close the connection to the chatroom
*/
func (bot *Bot) Close() {
	bot.conn.Close()
	fmt.Printf("Closed connection from %s\n", bot.server)
}

/*
LogIn logs into the irc service and joins a channel
*/
func (bot *Bot) LogIn(pass string) {
	//join channel
	fmt.Fprintf(bot.conn, "PASS %s\r\n", pass)
	fmt.Fprintf(bot.conn, "NICK %s\r\n", bot.name)
	fmt.Fprintf(bot.conn, "JOIN %s\r\n", bot.channel)
}

/*
Message sends a string to the chat channel
*/
func (bot *Bot) Message(message string) {
	if message != "" {
    	fmt.Printf("printing message: %s\r\n", message)
    	fmt.Fprintf(bot.conn, "PRIVMSG "+bot.channel+" :"+message+"\r\n")
    }
}

/*
ConsoleInput allows for controll over bot actions
*/
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
AutoMessage prints a string to chat
*/
func (bot *Bot) AutoMessage() {
	for {
		bot.Message("30 seconds has passed")
		time.Sleep(30 * time.Second)
	}
}

func (bot *Bot) emotemonGame() {
    gameMessage := make(chan string)
    bot.game = emotemon.NewEmotemonGame(gameMessage)
    go bot.game.Start()
    //defer bot.game.Close()

    for {
        bot.Message(<- gameMessage)
    }

}

/*
handleChat parses and responds to chat
*/
func (bot *Bot) handleChat() {
	//Creates the chat reader
	proto := textproto.NewReader(bufio.NewReader(bot.conn))

	for {
		line, err := proto.ReadLine()
		if err != nil {
			break
		}

		fmt.Printf("Read line %s \r\n", line)

		if strings.Contains(line, "PING") {
			pongResponse := strings.Split(line, "PING ")
			bot.Message("PONG " + pongResponse[1] + "\r\n")
		} else if strings.Contains(line, ".tmi.twitch.tv PRIVMSG "+bot.channel) {
			userdata := strings.Split(line, ".tmi.twitch.tv PRIVMSG "+bot.channel)
			username := strings.Split(userdata[0], "@")
			usermessage := strings.Replace(userdata[1], " :", "", 1)
			fmt.Printf(username[1] + ": " + usermessage + "\r\n")
			if strings.Contains(usermessage, bot.game.CurrentEmotemon()) {
                bot.game.CaptureAttempt(username[1], 1)
			} else if strings.Contains(usermessage, "LIST"){
                bot.game.GetTrainerEmotemon(username[1])
            }
		}
	}
}

func (bot *Bot) Start() {
    go bot.emotemonGame()

    bot.handleChat()
}
