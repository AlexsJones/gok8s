package configuration

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/AlexsJones/shed/crypto"
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

//Count number of items in list
func (m *MapConfiguration) Count() int {
	if m.maps == nil {
		return 0
	}
	return len(m.maps)
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

	parts := strings.Fields(i.Uri)
	if len(parts) < 1 {
		fmt.Println("Nothing to run!")
		return
	}

	log.Printf("Executing: %s", parts)
	c := exec.Command(parts[0], parts[1:]...)
	file, err := ioutil.TempFile(os.TempDir(), "go-")

	if err != nil {
		panic(err)
	}
	defer file.Close()
	c.Stdout = file
	c.Stderr = file
	i.Log = file.Name()
	err = c.Start()
	if err != nil {
		i.isSuccess(false)
		fmt.Println(err.Error())
		return
	}
	c.Wait()
	if err != nil {
		i.isSuccess(false)
		fmt.Println(err.Error())
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

//SaveOptions ...
type SaveOptions struct {
	Encrypted  bool
	Passphrase []byte
}

//Save ...
func (m *MapConfiguration) Save(options *SaveOptions) {

	b, err := json.Marshal(m.maps)
	if err != nil {
		fmt.Println(err)
	}
	var saveText string
	if options.Encrypted {
		var o [32]byte
		copy(options.Passphrase[:32], o[:])
		saveText = string(crypto.EncryptText([]byte(b), o))
	} else {
		saveText = string(b)
	}

	if _, err = os.Stat("Shedfile"); os.IsExist(err) {
		os.Remove("Shedfile")
	}

	f, err := os.Create("Shedfile")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	fmt.Printf("Save size %d\n", len(saveText))
	if _, err = f.WriteString(saveText); err != nil {
		panic(err)
	}

	fmt.Println("Created new Shedfile...")
}

//Load ...
func (m *MapConfiguration) Load() {

	b, err := ioutil.ReadFile("Shedfile")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Load size %d\n", len(string(b)))

	contentType := http.DetectContentType(b)

	if strings.Compare("text/plain; charset=utf-8", contentType) == 0 {
		var maps []*item
		err = json.Unmarshal(b, &maps)
		if err != nil {
			fmt.Println(err)
			return
		}
		m.maps = maps
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter unlock phrase followed by RETURN key:")
		text, _ := reader.ReadString('\n')
		res, err := crypto.HashPassword(text)
		if err != nil {
			fmt.Println(err)
		}
		var o [32]byte
		copy(res[:32], o[:])
		decrypted, worked := crypto.DecryptText(b, o)
		if worked {
			var maps []*item
			err = json.Unmarshal(decrypted, &maps)
			if err != nil {
				fmt.Println(err)
				return
			}
			m.maps = maps
		} else {
			fmt.Println("Key did not match")
		}
	}

	fmt.Println("Loaded Shedfile with " + strconv.Itoa(len(m.maps)) + " steps")
}
