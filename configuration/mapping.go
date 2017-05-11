package configuration

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

type item struct {
	uri       string
	validated bool
	executed  string //✓ ✗
}

func (i *item) Executed(b bool) {

	switch b {
	case true:

		i.executed = "✓"
	case false:

		i.executed = "✗"
	}
}

//MapConfiguration ...
type MapConfiguration struct {
	maps     []*item
	tableMap []string
}

//NewMapConfiguration ...
func NewMapConfiguration() *MapConfiguration {
	m := MapConfiguration{}
	m.tableMap = []string{"Step", "Resource locator", "Validated", "Executed"}
	return &m
}

//Clear ...
func (m *MapConfiguration) Clear() {
	m.maps = nil
}

//Push ...
func (m *MapConfiguration) Push(uri string) {
	i := item{uri: uri, validated: false}
	i.Executed(false)
	m.maps = append(m.maps, &i)
}

//Pull ...
func (m *MapConfiguration) Pull() string {
	return ""
}

//List ...
func (m *MapConfiguration) List() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(m.tableMap)

	var data [][]string

	var inc = 1
	for _, current := range m.maps {
		if current.uri != "" {
			data = append(data, []string{strconv.Itoa(inc), current.uri,
				fmt.Sprint(current.validated), fmt.Sprintf(current.executed)})
			inc++
		}
	}
	for _, current := range data {
		table.Append(current)
	}

	if len(data) >= 1 {
		table.Render()
	} else {
		fmt.Println("Nothing scheduled...")
	}
}

func (m *MapConfiguration) run(i *item) {

}

//Retry ...
func (m *MapConfiguration) Retry(i int) {

	if len(m.maps) < i || i < 0 {
		fmt.Println("Index out of bounds")
		return
	}

	m.run(m.maps[i])
}

//Run ...
func (m *MapConfiguration) Run() {

	var data [][]string
	var inc = 1

	for _, current := range m.maps {
		if current.uri != "" {

			m.run(current)
			current.Executed(true)
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			data = append(data, []string{strconv.Itoa(inc), current.uri,
				fmt.Sprint(current.validated), fmt.Sprintf(current.executed)})
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Step", "Resource locator", "Validated", "Executed"})
			for _, v := range data {
				table.Append(v)
			}
			table.Render()
		}
	}
}
