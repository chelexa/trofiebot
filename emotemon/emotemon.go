package emotemon

import (
    	import "math/rand"
)

type Emotemon struct {
        emotemon string
        cp int
        started bool
        bot Bot
}

func New(bot *Bot) *Emotemon {
    return &Emotemon {
                emotemon: "Kappa",
                cp: 1,
                started: false,
                bot: *bot,
        }
}

func ReceiveLine(user string, line string) {

}

func (emotemon *Emotemon) Start() {
	emotemon.generateEmotemon()
	
}

func (emotemon *Emotemon) generateEmotemon() {
	source := rand.NewSource(time.Now().UnixNano())
    rng := rand.New(source)
    emotemon.cp = rng.Intn(100)
}

