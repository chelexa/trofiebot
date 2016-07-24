package bot

import (
        "bufio"
        "fmt"
        "net"
        "net/textproto"
        "os"
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
                autoMsg:    "Let's play emotemon! BibleThump",
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
    fmt.Printf("before %s\n", bot.channel)
    if err != nil {
        fmt.Printf("Cannot connect to channel, retrying")
        bot.Connect()
    }
    fmt.Printf("Connected to IRC server %s\n", bot.server)
}

func (bot *Bot) Close() {
    bot.conn.Close()
    fmt.Printf("Closed connection from %s\n", bot.server)
}

func (bot *Bot) LogIn(pass string) {
    //join channel
    fmt.Fprintf(bot.conn, "PASS %s\r\n", pass)
    fmt.Fprintf(bot.conn, "NICK %s\r\n", bot.name)
    fmt.Fprintf(bot.conn, "JOIN %s\r\n", bot.channel)
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

func (bot *Bot) AutoMessage() {
    for {
        bot.Message("30 seconds has passed")
        time.Sleep(30 * time.Second)
    }
}

/*
Sends timeout duration to bot.Ban
*/
// func (bot *Bot) Timeout(user, reason, duration) {
//     if duration == 0 {
//         return
//     }
//     go bot.Ban(user, reason, duration)
// }

/*
Ban's the user from twitch chat
If duration sent is < 0, permanently ban
*/
// func (bot *Bot) Ban(user, reason, duration) {
//     if duration == -1 {
//         msg := fmt.Fprintf(".ban %s for %s", user, reason)
//     }
//     msg := fmt.Fprintf(".timeout %s for %s", user, reason)
//     time.Sleep(1 * time.Second)
//     bot.Message(msg)
// }

func (bot *Bot) HandleChat()  {

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
        } else if strings.Contains(line, ".tmi.twitch.tv PRIVMSG " + bot.channel) {
            userdata := strings.Split(line, ".tmi.twitch.tv PRIVMSG " + bot.channel)
            username := strings.Split(userdata[0], "@")
            usermessage := strings.Replace(userdata[1], " :", "", 1)
            fmt.Printf(username[1] + ": " + usermessage + "\n")
            if strings.Contains(usermessage, "Kappa") {
                bot.Message(username[1] + " caught a Kappa")
            }
        }
    }
}

/*
func main() {
    ircbot := NewBot()
    go ircbot.ConsoleInput()
    ircbot.Connect()

    parser := bufio.NewReader(ircbot.conn)
    proto := textproto.NewReader(parser)

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

    go ircbot.AutoMessage()

	defer ircbot.conn.Close()

    ircbot.Message("A wild Kappa has appeared with CP 1")

    for {
        line, err := proto.ReadLine()
        if err != nil {
            break
        }

        fmt.Printf("Read line %s \r\n", line)


        if strings.Contains(line, "PING") {
            pongResponse := strings.Split(line, "PING ")
            fmt.Fprintf(ircbot.conn, "PONG %s\r\n", pongResponse[1])
        } else if strings.Contains(line, ".tmi.twitch.tv PRIVMSG " + ircbot.channel) {

            userdata := strings.Split(line, ".tmi.twitch.tv PRIVMSG " + ircbot.channel)
            username := strings.Split(userdata[0], "@")
            usermessage := strings.Replace(userdata[1], " :", "", 1)
            fmt.Printf(username[1] + ": " + usermessage + "\n")

            if strings.Contains(usermessage, "Kappa") {
                ircbot.Message(username[1] + " caught a Kappa")
            }
        }
    }
}
*/
