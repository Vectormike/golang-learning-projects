package main

import (
	"fmt"
	"github.com/gen2brain/beeep"
	"github.com/olebedev/when"
	"github.com/olebedev/when/rules/common"
	"github.com/olebedev/when/rules/en"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	markName  = "CLI-REMINDER-TOOL"
	markValue = "1"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please provide a reminder message and time")
		os.Exit(1)
	}

	now := time.Now()
	w := when.New(nil)
	w.Add(en.All...)
	w.Add(common.All...)

	t, err := w.Parse(os.Args[2], now)
	if err != nil {
		fmt.Println("Error parsing time", err)
		os.Exit(1)
	}

	if t == nil {
		fmt.Println("No time passed")
		os.Exit(1)
	}

	if now.After(t.Time) {
		fmt.Println("Time is in the past")
		os.Exit(1)
	}

	diff := t.Time.Sub(now)
	if os.Getenv(markName) == markValue {
		time.Sleep(diff)
		err = beeep.Notify("Reminder", strings.Join(os.Args[1:], " "), "")
		if err != nil {
			fmt.Println("Error showing notification", err)
			os.Exit(1)
		}
	} else {
		cmd := exec.Command(os.Args[0], os.Args[1:]...)
		cmd.Env = append(os.Environ(), fmt.Sprintf("%s=%s", markName, markValue))
		err = cmd.Start()
		if err != nil {
			fmt.Println("Error starting process", err)
			os.Exit(1)
		}
	}

}
