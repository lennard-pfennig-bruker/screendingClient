package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Action: func(cCtx *cli.Context) error {
			mode := "server"
			if cCtx.NArg() > 0 {
				mode = cCtx.Args().Get(0)
			}
			go base.Generate()
			switch mode {
			case "server":
				fmt.Println("server mode")
			case "display":
				fmt.Println("display locally")
				display()
			default:
				fmt.Println("no mode set.. defaulting to server")
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
