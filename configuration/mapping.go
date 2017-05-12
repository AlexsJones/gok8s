package configuration

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strconv"

	"github.com/AlexsJones/shed/util"
	"github.com/olekukonko/tablewriter"
)

//MapConfiguration ...
type MapConfiguration struct {
	maps     []*item
	tableMap []string
}

//NewMapConfiguration ...
func NewMapConfiguration() *MapConfiguration {
	m := MapConfiguration{}
	m.tableMap = []string{"Step", "Resource locator", "Validated", "Executed", "Successful"}
	return &m
}

//Clear ...
func (m *MapConfiguration) Clear() {
	m.maps = nil
}

//Push ...
func (m *MapConfiguration) Push(uri string) {
	i := item{uri: uri}
	retb, _ := util.Exists(uri)
	_, err := url.ParseRequestURI(uri)
	if (err == nil) || (retb == true) {
		i.Validated(false)
	} else {
		i.Validated(true)
	}
	i.Executed(false)
	i.success = "?"
	m.maps = append(m.maps, &i)
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
				fmt.Sprint(current.validated), fmt.Sprintf(current.executed), current.success})
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
	c := exec.Command(i.uri)
	file, err := ioutil.TempFile(os.TempDir(), "go-")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	c.Stdout = file
	c.Stderr = file
	i.log = file.Name()
	err = c.Run()
	if err != nil {
		i.Success(false)
		return
	}
	i.Success(true)
}

//Logs ...
func (m *MapConfiguration) Logs(i int) {

	if len(m.maps) < i || i < 0 {
		fmt.Println("Index out of bounds")
		return
	}

	current := m.maps[i]
	if current.log != "" {
		file, err := os.Open(current.log)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}
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
				fmt.Sprint(current.validated), fmt.Sprintf(current.executed), current.success})
			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader(m.tableMap)
			for _, v := range data {
				table.Append(v)
			}
			table.Render()
			inc++
		}
	}
}
