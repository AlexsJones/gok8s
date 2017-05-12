package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/AlexsJones/shed/configuration"
	"github.com/abiosoft/ishell"
	"github.com/dimiro1/banner"
)

const b string = `
                                                                    dddddddd
                hhhhhhh                                             d::::::d
                h:::::h                                             d::::::d
                h:::::h                                             d::::::d
                h:::::h                                             d:::::d
    ssssssssss   h::::h hhhhh           eeeeeeeeeeee        ddddddddd:::::d
  ss::::::::::s  h::::hh:::::hhh      ee::::::::::::ee    dd::::::::::::::d
ss:::::::::::::s h::::::::::::::hh   e::::::eeeee:::::ee d::::::::::::::::d
s::::::ssss:::::sh:::::::hhh::::::h e::::::e     e:::::ed:::::::ddddd:::::d
 s:::::s  ssssss h::::::h   h::::::he:::::::eeeee::::::ed::::::d    d:::::d
   s::::::s      h:::::h     h:::::he:::::::::::::::::e d:::::d     d:::::d
      s::::::s   h:::::h     h:::::he::::::eeeeeeeeeee  d:::::d     d:::::d
ssssss   s:::::s h:::::h     h:::::he:::::::e           d:::::d     d:::::d
s:::::ssss::::::sh:::::h     h:::::he::::::::e          d::::::ddddd::::::dd
s::::::::::::::s h:::::h     h:::::h e::::::::eeeeeeee   d:::::::::::::::::d
 s:::::::::::ss  h:::::h     h:::::h  ee:::::::::::::e    d:::::::::ddd::::d
  sssssssssss    hhhhhhh     hhhhhhh    eeeeeeeeeeeeee     ddddddddd   ddddd
`

func main() {
	banner.Init(os.Stdout, true, true, bytes.NewBufferString(b))

	mapConfiguration := configuration.NewMapConfiguration()

	shell := ishell.New()

	shell.AddCmd(&ishell.Cmd{
		Name: "push",
		Help: "push a k8s config-map path",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 || c.Args[0] == "" {
				fmt.Println("Requires PATH or URL")
				return
			}
			c.Println("Pushing -> " + c.Args[0])
			mapConfiguration.Push(c.Args[0])
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "clear",
		Help: "clear the current stack",
		Func: func(c *ishell.Context) {

			mapConfiguration.Clear()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "retry",
		Help: "retry a certain action based on index",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 || c.Args[0] == "" {
				fmt.Println("Requires an index number")
				return
			}
			c.Println("Reapplying...")

			i, err := strconv.Atoi(c.Args[0])
			if err != nil {
				fmt.Println("Unable to parse int")
				return
			}
			mapConfiguration.Retry(i)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "list",
		Help: "list execution order",
		Func: func(c *ishell.Context) {

			mapConfiguration.List()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "logs",
		Help: "logs from an execution",
		Func: func(c *ishell.Context) {
			if len(c.Args) < 1 || c.Args[0] == "" {
				fmt.Println("Requires an index number")
				return
			}
			i, err := strconv.Atoi(c.Args[0])
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			mapConfiguration.Logs(i - 1)
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "run",
		Help: "Starts running k8s config-map paths",
		Func: func(c *ishell.Context) {

			mapConfiguration.Run()
		},
	})
	shell.Start()
}
