package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mchcli "github.com/c-m-hunt/myclubhouse/cli"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:  "set-config",
				Usage: "Set your access token and subdomain",
				Action: func(c *cli.Context) error {
					args := c.Args()
					if len(args) != 2 {
						fmt.Println("Provide only two variables - access token and subdomain")
					} else {
						err := mchcli.SetConfig(args[0], args[1])
						if err != nil {
							fmt.Println("Config could not be saved")
							fmt.Println(err)
						}
					}
					return nil
				},
			},
			{
				Name:    "exp-mem",
				Aliases: []string{"ex"},
				Usage:   "Lists members with membership expiring in the next two weeks",
				Action: func(c *cli.Context) error {
					cfg := mchcli.Config{}
					cfg.EnsureConfig()
					mchcli.DisplayExpiringMembers(&cfg, time.Now().AddDate(0, 0, 14))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
