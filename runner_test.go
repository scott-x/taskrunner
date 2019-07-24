/*
* @Author: apple
* @Date:   2019-07-10 06:14:58
* @Last Modified by:   scottxiong
* @Last Modified time: 2019-07-24 17:55:03
 */

package taskrunner

import (
	"log"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	d := func(dc DataChan) error {
		for i := 0; i < 30; i++ {
			dc <- i
			log.Printf("Dispatcher sent: %v", i)
		}
		return nil
	}
	e := func(dc DataChan) error {
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

	runner := NewRunner(30, false, d, e)
	go runner.StartAll()
	time.Sleep(1 * time.Second)
}
