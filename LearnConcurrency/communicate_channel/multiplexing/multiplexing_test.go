package multiplexing

import (
	"fmt"
	"github.com/naraycitra/learn-golang/LearnConcurrency/communicate_channel/ping_pong"
	"testing"
)

func TestFanIn(t *testing.T) {
	// testcase:
	ping := ping_pong.PingGen("ping")
	beep := ping_pong.PingGen("beep")
	clin := ping_pong.PingGen("clin")

	out := FanIn(ping, beep, clin)
	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}
	fmt.Println("aaaa")
}

func TestFanIn2(t *testing.T) {
	stopPing := make(chan bool)
	ping := ping_pong.PingGenWithQuit("ping", stopPing)
	out := FanIn(ping)
	// manual send stop
	for i := 0; i < 10; i++ {
		fmt.Println(<-out)
	}
	stopPing <- true
	fmt.Println("aaaaa")
}
