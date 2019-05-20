package main

import (
	"log"
	"os"

	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "tcp_terminator"
	app.Usage = "Simple TCP server that is useful for testing purposes."

	app.Commands = []cli.Command{
		{
			Name:      "start",
			Usage:     "Start the TCP server and listen for messages",
			ArgsUsage: "[addr]",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "file, f",
					Usage: "A file to write messages to",
				},
			},
			Action: func(ctx *cli.Context) error {
				addr := ctx.Args().Get(0)

				if addr == "" {
					message := "The address argument is required: `tcp_terminator 0.0.0.0:9000`"
					// Exit with 65, EX_DATAERR, to indicate input data was incorrect
					return cli.NewExitError(message, 65)
				}

				file := ctx.String("file")
				server := NewServer(addr, file)
				server.Listen()

				if server.File != nil {
					log.Println("Closing file")
					server.File.Close()
				}

				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
