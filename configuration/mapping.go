package configuration

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

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
	m.tableMap = []string{"Step", "Resource locator", "Executed", "Successful"}
	return &m
}

//Clear ...
func (m *MapConfiguration) Clear() {
	m.maps = nil
}

//Push ...
func (m *MapConfiguration) Push(uri string) {
	i := item{Uri: uri}
	i.isExecuted(false)
	i.Success = "?"
	m.maps = append(m.maps, &i)
}

//List ...
func (m *MapConfiguration) List() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(m.tableMap)

	var data [][]string

	var inc = 1
	for _, current := range m.maps {
		if current.Uri != "" {
			data = append(data, []string{strconv.Itoa(inc), current.Uri,
				fmt.Sprintf(current.Executed), current.Success})
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
	c := exec.Command(i.Uri)
	file, err := ioutil.TempFile(os.TempDir(), "go-")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	c.Stdout = file
	c.Stderr = file
	i.Log = file.Name()
	err = c.Run()
	if err != nil {
		i.isSuccess(false)
		return
	}
	i.isSuccess(true)
}

//Logs ...
func (m *MapConfiguration) Logs(i int) {

	if len(m.maps) < i || i < 0 {
		fmt.Println("Index out of bounds")
		return
	}

	current := m.maps[i]
	if current.Log != "" {
		file, err := os.Open(current.Log)
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
		if current.Uri != "" {
			m.run(current)
			current.isExecuted(true)
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			data = append(data, []string{strconv.Itoa(inc), current.Uri,
				fmt.Sprintf(current.Executed), current.Success})
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

//Save ...
func (m *MapConfiguration) Save() {

	if _, err := os.Stat("Shedfile"); os.IsExist(err) {
		os.Remove("Shedfile")
	}

	f, err := os.Create("Shedfile")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	var b []byte
	for _, i := range m.maps {
		b, err = json.Marshal(i)
		if err != nil {
			fmt.Println(err)
			continue
		}
		jsonStr := string(b)
		if _, err = f.WriteString(jsonStr + "\n"); err != nil {
			panic(err)
		}
	}

	fmt.Println("Created new Shedfile...")
}

//Load ...
func (m *MapConfiguration) Load() {

	file, err := os.Open("Shedfile")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var count int
	for scanner.Scan() {
		line := scanner.Text()
		i := &item{}
		err = json.Unmarshal([]byte(line), &i)
		if err != nil {
			fmt.Println(err)
		}

		m.maps = append(m.maps, i)
		count++
	}

	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Loaded Shedfile with " + strconv.Itoa(count) + " steps")
}
