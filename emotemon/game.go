package emotemon

import (
	//"fmt"
	"time"
)

type EmotemonGame struct {
	emotemon *Emotemon
	trainers map[string]*Emotemon
	output chan<- string
}

func NewEmotemonGame(output chan<- string) *EmotemonGame {
	return &EmotemonGame {
		trainers: make(map[string]*Emotemon),
		output: output,
	}
}

func (game *EmotemonGame) Start() {
	capturer := make(chan string)

	for {
		game.emotemon = NewEmotemon(capturer)

		game.output <- game.emotemon.Found()
		
		//wait for emotemon to be captured
		username := <- capturer
		game.output <- game.emotemon.GetEmote() + " was captured: " + username

		game.trainers[username] = game.emotemon

		time.Sleep(5 * time.Second)
	}

}

func (game *EmotemonGame) GetTrainerEmotemon(username string) {
	if emotemon, exists := game.trainers[username]; exists {
		game.output <- "Trainer " + username + " has " + emotemon.String()
	}
}

func (game *EmotemonGame) CurrentEmotemon() string {
	return game.emotemon.GetEmote()
}

func (game *EmotemonGame) CaptureAttempt(username string, attack int) {
	game.emotemon.CaptureAttempt(username, attack)
}