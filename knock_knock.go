package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	ip := "IP_TO_PING"
	for {
		out, _ := exec.Command("ping", ip, "-c 5", "-i 1").Output()
		// fmt.Println(string(out))
		if strings.Contains(string(out), "Request timeout") || strings.Contains(string(out), "Destination Host Unreachable") {
			fmt.Println("TANGO DOWN")
		} else {
			fmt.Println("KNOCK KNOCK")
			say := exec.Command("say", "knock knock")
			say.Run()
		}
	}
}
