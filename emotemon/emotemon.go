package emotemon

import (
    	"math/rand"
    	"time"
    	"strconv"
)

type Emotemon struct {
        emote string
        cp int
        health int
        capturer chan<- string
}

/*
* Generates random values for an emotemon
* Ensures that CaptureCount == CP and CP > 0
*/
func (emotemon *Emotemon) generateEmotemon() {

	// static array of several emotes.
	//TODO: get emotes from Twitch.tv API
	emotes := [5]string{
		"Kappa",
		"BibleThump",
		"BabyRage",
		"KappaPride",
		"BrainSlug",
	}

	// Creating a seed for the RNG
	source := rand.NewSource(time.Now().UnixNano())
	// The RNG
    rng := rand.New(source)

    // Randomly generating the capture value for this emotemon (ensure the value is not 0)
    capture := rng.Intn(5) + 1

    // Set CP
    emotemon.cp = capture

    // Set capture count initially to the CP
    emotemon.health = capture

    // Set the emote
    emotemon.emote = emotes[rng.Intn(5)]
}

/*
* Creates a new emotemon
*/
func NewEmotemon(capturer chan<- string) *Emotemon {

	emotemon := &Emotemon {
                emote: "Kappa",
                cp: 1,
                health: 1,
                capturer: capturer,
        }

    emotemon.generateEmotemon()

    return emotemon
}

func (emotemon *Emotemon) CaptureAttempt(username string, attack int) {
	emotemon.health -= attack
	if emotemon.health <= 0 {
		emotemon.capturer <- username
	}
}

func (emotemon *Emotemon) GetEmote() string {
	return emotemon.emote;
}

func (emotemon *Emotemon) Found() string {
	cp := strconv.Itoa(emotemon.cp)
	return "A wild " + emotemon.emote + 
	" has appeared with CP " + cp
}

func (emotemon *Emotemon) String() string {
	return "Emote: " + emotemon.emote + "  CP: " + strconv.Itoa(emotemon.cp)
}

