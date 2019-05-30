package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "tcp_test_server"
	app.Usage = "Simple TCP server that is useful for testing purposes."
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "address, a",
			Usage: "The address to bind to",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		addr := ctx.String("address")

		if addr == "" {
			message := "The address argument is required: `tcp_test_server -a 0.0.0.0:9000`"
			// Exit with 65, EX_DATAERR, to indicate input data was incorrect
			return cli.NewExitError(message, 65)
		}

		server := NewServer(addr)

		var gracefulStop = make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)

		go func() {
			sig := <-gracefulStop
			log.Printf("Caught sig: %+v", sig)
			server.WriteSummary()
			log.Println("Server stopped")
			os.Exit(0)
		}()

		server.Listen()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
