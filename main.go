package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/AlexsJones/gok8s/configuration"
	"github.com/abiosoft/ishell"
	"github.com/dimiro1/banner"
)

const b string = `
                     /$$        /$$$$$$
                    | $$       /$$__  $$
  /$$$$$$   /$$$$$$ | $$   /$$| $$  \ $$  /$$$$$$$
 /$$__  $$ /$$__  $$| $$  /$$/|  $$$$$$/ /$$_____/
| $$  \ $$| $$  \ $$| $$$$$$/  >$$__  $$|  $$$$$$
| $$  | $$| $$  | $$| $$_  $$ | $$  \ $$ \____  $$
|  $$$$$$$|  $$$$$$/| $$ \  $$|  $$$$$$/ /$$$$$$$/
 \____  $$ \______/ |__/  \__/ \______/ |_______/
 /$$  \ $$
|  $$$$$$/
 \______/
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
		Name: "run",
		Help: "Starts running k8s config-map paths",
		Func: func(c *ishell.Context) {

			mapConfiguration.Run()
		},
	})
	shell.Start()
}
