package retry

import (
	"errors"
	"log"
	"time"
)

type myFunc func(attempt int) (retry bool, err error)

const maxRetries = 10

var errMaxRetriesReached = errors.New("max retries reached")

func try(fn myFunc) error {
	var err error
	var retry bool
	attempt := 1
	for {
		retry, err = fn(attempt)
		if !retry || err == nil {
			break
		}

		attempt++
		if attempt > maxRetries {
			return errMaxRetriesReached
		}
	}

	return err
}

func doSomething() (int, error) {
	time.Sleep(1 * time.Second)
	return 42, nil
}

func main() {
	err := try(func(attempt int) (bool, error) {
		var errSomething error
		_, errSomething = doSomething()
		return attempt < 5, errSomething
	})

	if err != nil {
		log.Fatalln("error:", err)
	}
}
