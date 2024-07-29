package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
	"strconv"
	"net/http"
)

func getPIDs() []string {
	cmd := exec.Command("nvidia-smi", "-i", os.Args[1], "--query-compute-apps=pid", "--format=csv")
	// time.Sleep(time.Second)
	output, _ := cmd.CombinedOutput()
	// fmt.Println(err)
	result := string(output)
	// fmt.Println(result)
	resultArr := strings.Split(result, "\n")

	wantToHelp := []string{}
	for _, pid := range resultArr {
		if pid == "pid" {
			continue
		}
		if len(pid) > 0 {
			wantToHelp = append(wantToHelp, pid)
		}
	}
	if len(wantToHelp) == 0 {
		fmt.Println("len==0")
		os.Exit(0)
	}

	return wantToHelp
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gpuid [second divider(default 3)]")
		os.Exit(1)
	}

	sleeptime := time.Second / 3
	if len(os.Args) == 3 {
		divider, err := strconv.Atoi(string(os.Args[2]))
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		sleeptime = time.Second / time.Duration(divider)
	}

	newPIDs := getPIDs()
	wantToHelp := newPIDs[0]
	fmt.Println("want to help: ", wantToHelp)
	fmt.Println("interval: ", sleeptime)

	notified := false

	for len(newPIDs) > 0 {
		newPIDs := getPIDs()
		for _, pid := range newPIDs {
			if pid != wantToHelp {
				fmt.Println("kill -11", pid)
				cmd := exec.Command("kill", "-11", pid)
				output, err := cmd.CombinedOutput()
				fmt.Println("kill -8", pid)
				cmd = exec.Command("kill", "-8", pid)
				output, err = cmd.CombinedOutput()
				fmt.Println("kill -6", pid)
				cmd = exec.Command("kill", "-6", pid)
				output, err = cmd.CombinedOutput()
				// fmt.Println("kill -9", pid)
				// cmd = exec.Command("kill", "-9", pid)
				// output, err = cmd.CombinedOutput()
				// time.Sleep(time.Millisecond * 10)
				fmt.Println(string(output), err)
				if !notified {
					notified = true
					go func() {
						response, err := http.Get("https://kelee-push.daisyorge.fun/v1?form=text&content=GPU" + os.Args[1] + " down")
						fmt.Println(response, err)
						time.Sleep(time.Second * 15)
						notified = false
					}()
				}
			}
			// fmt.Print(".")
			time.Sleep(time.Second / sleeptime)
		}
	}
}
