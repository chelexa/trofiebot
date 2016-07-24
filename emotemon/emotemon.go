package emotemon

import (
    	"math/rand"
    	"time"
)

type Emotemon struct {
        emotemon string
        cp int
        CaptureCount int
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
    capture := rng.Intn(100) + 1

    // Set CP
    emotemon.cp = capture

    // Set capture count initially to the CP
    emotemon.CaptureCount = capture

    // Set the emotemon
    emotemon.emotemon = emotes[rng.Intn(5)]
}

/*
* Creates a new emotemon
*/
func New() *Emotemon {

	emotemon := &Emotemon {
                emotemon: "Kappa",
                cp: 1,
                CaptureCount: 1,
        }

    emotemon.generateEmotemon()

    return emotemon
}


