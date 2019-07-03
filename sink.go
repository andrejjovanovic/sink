package main

import (
	"fmt"
	"os"
	"log"
	"github.com/urfave/cli"
)


func main() {
	app := cli.NewApp()
	
	app.Name = "Sink"
	app.Version = "0.0.1"
	app.Usage = "Continious sync across devices!"
	app.Copyright = "(c) 2019 Andrej Jovanovic"

	app.Commands = []cli.Command{
		{
		  Name:     "add",
		  Usage: 	"Add 'SOURCE' or 'DESTINATION'",
		  Flags: []cli.Flag{
			cli.StringFlag {
				Name : "source, s",
			},
			cli.StringFlag {
				Name : "destination, d",
			},
		  },

		  Action: func(c *cli.Context) error {
			  if c.String("source") != "" {
				fmt.Fprintf(c.App.Writer, "Source set to: %v\n", c.String("source"))
			  } else if c.String("destination") != "" {
				fmt.Fprintf(c.App.Writer, "Destination set to: %v\n", c.String("destination"))
			  }
			  return nil
		  },
		},
		{
		  Name:     "remove",
		},
	}

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "print, p",
			Value: "Djeste barbari",
			Usage: "Just print a message",
		  },
	}
	
	err := app.Run(os.Args)
  	if err != nil {
   		log.Fatal(err)
	}
	
}