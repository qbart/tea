package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/wzshiming/ctc"
)

var (
	subdomain *string = flag.String("subdomain", "", "Specify subdomain to use")
	port      *uint   = flag.Uint("port", 3000, "Specify port to forward to")
)

func main() {
	flag.Parse()
	if *subdomain == "" {
		fmt.Print(ctc.ForegroundRed, "Empty subdomain", ctc.Reset, "\n")
		os.Exit(1)
	}
	fmt.Printf("Forwarding https://%s.tcp.sh -> http://localhost:%d\n", *subdomain, *port)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

LOOP:
	for {
		select {
		case <-ctx.Done():
			stop()
			break LOOP
		}
	}

	fmt.Println("Exiting..")
}
