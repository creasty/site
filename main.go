package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/creasty/site/utils"
)

func main() {
	devMode := flag.Bool("dev", false, "enable development mode")
	flag.Parse()

	runtime.GOMAXPROCS(runtime.NumCPU())
	utils.LoadConfig()

	run(*devMode)
}

func run(devMode bool) {
	if devMode {
		runDevServer()
	} else {
		if err := serveApi(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

func runDevServer() {
	cmd := exec.Command(
		"gin",
		"-g",
		"-p", strconv.Itoa(int(utils.Config.DevPort)),
		"-a", strconv.Itoa(int(utils.Config.Port)),
		"-b", os.Args[0],
		"r",
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		fmt.Println("You'll need to have gin installed: https://github.com/codegangsta/gin")
		os.Exit(1)
	}
}
