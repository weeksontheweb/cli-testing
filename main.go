/* greet.go */
package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(c *cli.Context) {
		println("Hello friend!")
	}

	// global level flags
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
		cli.StringFlag{
			Name:  "f, file",
			Usage: "Specify an alternate fig file (default: fig.yml)",
		},
		cli.StringFlag{
			Name:  "p, project-name",
			Usage: "Specify an alternate project name (default: directory name)",
		},
	}

	// Commands
	app.Commands = []cli.Command{
		{
			Name: "build",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "no-cache",
					Usage: "Do not use cache when building the image.",
				},
			},
			Usage:  "Build or rebuild services",
			//Action: CmdBuild,
		},
		// etc...
		{
			Name: "run",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "d",
					Usage: "Detached mode: Run container in the background, print new container name.",
				},
				cli.BoolFlag{
					Name:  "T",
					Usage: "Disables psuedo-tty allocation. By default `fig run` allocates a TTY.",
				},
				cli.BoolFlag{
					Name:  "rm",
					Usage: "Remove container after run.  Ignored in detached mode.",
				},
				cli.BoolFlag{
					Name:  "no-deps",
					Usage: "Don't start linked services.",
				},
			},
			Usage:  "Run a one-off command",
			//Action: CmdRm,
		},

		{
			Name: "up",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "watch",
					Usage: "Watch build directory for changes and auto-rebuild/restart",
				},
				cli.BoolFlag{
					Name:  "d",
					Usage: "Detached mode: Run containers in the background, print new container names.",
				},
				cli.BoolFlag{
					Name:  "k,kill",
					Usage: "Kill instead of stop on terminal stignal",
				},
				cli.BoolFlag{
					Name:  "no-clean",
					Usage: "Don't remove containers after termination signal interrupt (CTRL+C)",
				},
				cli.BoolFlag{
					Name:  "no-deps",
					Usage: "Don't start linked services.",
				},
				cli.BoolFlag{
					Name:  "no-recreate",
					Usage: "If containers already exist, don't recreate them.",
				},
			},
			Usage:  "Create and start containers",
			Action: CmdUp,
		},
	}
	app.Run(os.Args)
}

func CmdUp() {

}