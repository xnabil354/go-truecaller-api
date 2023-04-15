package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
)

type TrueCaller struct {
	nomorCode   string
	countryCode string
}

func (tc *TrueCaller) Start() {
	color.New(color.BgHiYellow, color.FgBlack).Println(" TRUECALLER|API ")
	fmt.Println(color.HiBlueString("======================================"))
	fmt.Println(color.YellowString("Creator : xzhndvs"))
	fmt.Println(color.YellowString("Instagram : @xzhndvs"))
	fmt.Println(color.YellowString("WhatsApp : 081281524356"))
	fmt.Println(color.HiBlueString("======================================"))
}

func (tc *TrueCaller) FetchData() (string, string, string) {
	url := fmt.Sprintf("https://xzhndvs.vercel.app/api/truecaller?nomorCode=%s&countryCode=%s", tc.nomorCode, tc.countryCode)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	scanner.Scan()
	jsonStr := scanner.Text()

	var m map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := m["data"].(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})
	name := data["name"].(string)
	phone := data["phones"].([]interface{})[0].(map[string]interface{})
	provider := phone["carrier"].(string)
	nationalFormat := phone["nationalFormat"].(string)

	return name, provider, nationalFormat
}

func main() {
	tc := TrueCaller{}

	tc.Start()

	for {
		fmt.Print(color.RedString("Enter nomorCode : "))
		fmt.Scanln(&tc.nomorCode)
		fmt.Print(color.RedString("Enter countryCode : "))
		fmt.Scanln(&tc.countryCode)

		name, provider, nationalFormat := tc.FetchData()
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendRow([]interface{}{"Name", name})
		t.AppendRow([]interface{}{"Provider", provider})
		t.AppendRow([]interface{}{"National Format", nationalFormat})
		t.Render()

		fmt.Println()
	}
}
