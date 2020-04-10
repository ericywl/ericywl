package returnfuncs

import (
	"log"
	"time"
)

// startTimer returns the stopTimer function
func startTimer(name string) func() {
	t := time.Now()
	log.Println(name, "started")
	return func() {
		d := time.Now().Sub(t)
		log.Println(name, "took", d)
	}
}

func funkyFunc() {
	stopTimer := startTimer("funkyFunc")
	defer stopTimer()

	time.Sleep(1 * time.Second)
}
