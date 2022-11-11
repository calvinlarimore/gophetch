package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	red     = "\033[31m"
	cyan    = "\033[36m"
	magenta = "\033[35m"
	reset   = "\033[0m"

	bold = "\033[1m"

	format = "\033[36m\033[1m%10s\033[0m - %s%s\033[0m"
)

func main() {
	fmt.Print("      - \033[36m\033[1mGophetch\033[0m -\n\n")

	username := make(chan string, 1)
	go getUser(username)

	os := make(chan string, 1)
	go getOS(os)

	cpu := make(chan string, 1)
	go getCPU(cpu)

	memory := make(chan string, 1)
	go getMemory(memory)

	uptime := make(chan string, 1)
	go getUptime(uptime)

	shell := make(chan string, 1)
	go getShell(shell)

	desktop := make(chan string, 1)
	go getDesktop(desktop)

	fmt.Println(<-username)
	fmt.Println(<-os)
	fmt.Println(<-cpu)
	fmt.Println(<-memory)
	fmt.Println(<-uptime)
	fmt.Println(<-shell)
	fmt.Println(<-desktop)

	fmt.Println("\n         - \033[36m\033[1m-\033[0m -")
}

func formatLine(f string, v string) string {
	return fmt.Sprintf(format, f, magenta, v)
}

func formatError(f string) string {
	return fmt.Sprintf(format, f, red, "Error!")
}

func getUser(c chan string) {
	u, err := user.Current()
	if err != nil {
		c <- formatError("User")
		return
	}

	h, err := os.Hostname()
	if err != nil {
		c <- formatError("User")
		return
	}

	c <- formatLine("User", fmt.Sprintf("%s@%s", u.Username, h))
}

func getOS(c chan string) {
	h, err := host.Info()
	if err != nil {
		c <- formatError("OS")
		return
	}

	title := cases.Title(language.Und)

	c <- formatLine("OS", fmt.Sprintf("%s (%s)", title.String(h.Platform), title.String(h.OS)))
}

func getCPU(c chan string) {
	i, err := cpu.Info()

	if err != nil {
		c <- formatError("CPU")
		return
	}

	c <- formatLine("CPU", fmt.Sprintf("%dx %s", len(i), i[0].ModelName))
}

func getMemory(c chan string) {
	m, err := mem.VirtualMemory()

	if err != nil {
		c <- formatError("Memory")
		return
	}

	formatMemory := func(b uint64) string {
		const unit = 1000

		if b < unit {
			return fmt.Sprintf("%d B", b)
		}

		div, exp := int64(unit), 0
		for n := b / unit; n >= unit; n /= unit {
			div *= unit
			exp++
		}

		return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
	}

	c <- formatLine("Memory", fmt.Sprintf("%s / %s (%s free)", formatMemory(m.Used), formatMemory(m.Total), formatMemory(m.Free)))
}

func getUptime(c chan string) {
	u, err := host.Uptime()
	if err != nil {
		c <- formatError("Uptime")
		return
	}

	c <- formatLine("Uptime", fmt.Sprintf("%d Days %02d:%02d:%02d", u/(24*60*60), u/(60*60)%24, u/60%60, (u%60)%60))
}

func getShell(c chan string) {
	s, ok := os.LookupEnv("SHELL")

	if !ok {
		c <- formatLine("Shell", "/usr/bin/sh or Unknown")
		return
	}

	c <- formatLine("Shell", s)
}

func getDesktop(c chan string) {
	d, ok := os.LookupEnv("XDG_CURRENT_DESKTOP")

	if !ok {
		c <- formatLine("WM/DE", "None/Unknown")
		return
	}

	c <- formatLine("WM/DE", d)
}
