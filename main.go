package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/AlexsJones/shed/configuration"
	"github.com/abiosoft/ishell"
	"github.com/dimiro1/banner"
)

const b string = `
{{ .AnsiColor.Green }}_______________/\\\________________________________/\\\__
{{ .AnsiColor.Green }} ______________\/\\\_______________________________\/\\\__
{{ .AnsiColor.Green }}  ______________\/\\\_______________________________\/\\\__
{{ .AnsiColor.Green }}   __/\\\\\\\\\\_\/\\\_____________/\\\\\\\\_________\/\\\__
{{ .AnsiColor.Green }}    _\/\\\//////__\/\\\\\\\\\\____/\\\/////\\\___/\\\\\\\\\__
{{ .AnsiColor.Green }}     _\/\\\\\\\\\\_\/\\\/////\\\__/\\\\\\\\\\\___/\\\////\\\__
{{ .AnsiColor.Green }}      _\////////\\\_\/\\\___\/\\\_\//\\///////___\/\\\__\/\\\__
{{ .AnsiColor.Green }}       __/\\\\\\\\\\_\/\\\___\/\\\__\//\\\\\\\\\\_\//\\\\\\\/\\_
{{ .AnsiColor.Green }}        _\//////////__\///____\///____\//////////___\///////\//__
{{ .AnsiColor.Default }}
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
			mapConfiguration.Push(strings.Join(c.Args, " "))
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
	shell.AddCmd(&ishell.Cmd{
		Name: "save",
		Help: "Saves out a new ShedFile",
		Func: func(c *ishell.Context) {

			mapConfiguration.Save()
		},
	})
	shell.AddCmd(&ishell.Cmd{
		Name: "load",
		Help: "Loads a local ShedFile into a schedule",
		Func: func(c *ishell.Context) {

			mapConfiguration.Load()
		},
	})
	shell.Start()
}
