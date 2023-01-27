package main

import (
	"flag"
	"log"
	"os"
	"runtime"

	"github.com/plin2k/api-mocker/v2/config"
	"github.com/plin2k/api-mocker/v2/mocker"
)

func main() {

	if len(os.Args[0:]) < 1 {
		log.Fatalln("You must pass run subcommand")
	}

	switch os.Args[1] {
	case "run":
		run()
	default:
		log.Println("Unknown command")
	}
}

func run() {
	var (
		cfg      = &config.Mocker{}
		maxProcs int
	)

	runFlags := flag.NewFlagSet("run", flag.ExitOnError)
	runFlags.StringVar(&cfg.Host, "host", "127.0.0.1", "Host (default: 127.0.0.1)")
	runFlags.IntVar(&cfg.Port, "port", 8080, "Port (default: 8080)")
	runFlags.StringVar(&cfg.SrcPath, "src", "api-mocker.xml", "Source file (default: api-mocker.xml)")
	runFlags.IntVar(&maxProcs, "gomaxprocs", runtime.NumCPU(), "Set runtime.GOMAXPROCS")

	runtime.GOMAXPROCS(maxProcs)

	if runFlags.Parse(os.Args[2:]) != nil {
		log.Fatalln("can't parse arguments")
	}

	p := mocker.New(cfg)
	err := p.Process()
	if err != nil {
		log.Fatal(err)
	}

	if err = p.Run(); err != nil {
		log.Fatal(err)
	}
}
