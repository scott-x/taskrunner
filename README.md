# taskrunner
taskrunner for golang

### useage

```golang
package main

import (
	"github.com/scott-x/taskrunner"
	"log"
	"time"
)

func main() {
	d := func(dc taskrunner.DataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %v", i)
		}
		return nil
	}
	e := func(dc taskrunner.DataChan) error {
	forloop:
		for {
			select {
			case d := <-dc:
				log.Printf("Excutor received: %v", d)
			default:
				break forloop
			}

		}
		return nil
	}

	runner := taskrunner.NewRunner(30, false, d, e)
	go runner.StartAll()
	time.Sleep(1 * time.Second)
}

```
