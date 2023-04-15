package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/table"
)

type TrueCaller struct {
	nomorCode   string
	countryCode string
}

func (tc *TrueCaller) Start() {
	color.New(color.BgHiYellow, color.FgBlack).Println(" GO TRUECALLER API ")
	fmt.Println(color.HiBlueString("======================================"))
	fmt.Println(color.YellowString("Creator : xzhndvs"))
	fmt.Println(color.YellowString("Instagram : @xzhndvs"))
	fmt.Println(color.YellowString("WhatsApp : 081281524356"))
	fmt.Println(color.HiBlueString("======================================"))
}

func (tc *TrueCaller) FetchData() (string, string, string, string, string, string, string) {
	url := fmt.Sprintf("https://xzhndvs.vercel.app/api/truecaller?nomorCode=%s&countryCode=%s", tc.nomorCode, tc.countryCode)
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	scanner := bufio.NewScanner(res.Body)
	scanner.Scan()
	jsonStr := scanner.Text()

	var m map[string]interface{}
	err = json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		panic(err)
	}

	data := m["data"].(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})
	var name string
	if m["data"] != nil {
		data := m["data"].(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})
		if data["name"] != nil {
			name = data["name"].(string)
		} else {
			name = "Not found"
		}
	}

	var image string
	if data["image"] != nil {
		image = data["image"].(string)
	} else {
		image = "No image found"
	}

	phone := data["phones"].([]interface{})[0].(map[string]interface{})
	provider := phone["carrier"].(string)
	dialingCode := int(phone["dialingCode"].(float64))
	numberType := phone["numberType"].(string)
	country := phone["countryCode"].(string)
	nationalFormat := phone["nationalFormat"].(string)

	return name, image, provider, strconv.Itoa(dialingCode), numberType, country, nationalFormat
}

func main() {
	tc := TrueCaller{}

	tc.Start()

	for {
		fmt.Print(color.RedString("Enter nomorCode : "))
		fmt.Scanln(&tc.nomorCode)
		fmt.Print(color.RedString("Enter countryCode : "))
		fmt.Scanln(&tc.countryCode)

		name, image, provider, dialingCodeStr, numberType, country, nationalFormat := tc.FetchData()
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendRow([]interface{}{"Name", name})
		t.AppendRow([]interface{}{"Provider", provider})
		t.AppendRow([]interface{}{"National Format", nationalFormat})
		t.AppendRow([]interface{}{"Dialing Code", dialingCodeStr})
		t.AppendRow([]interface{}{"Number Type", numberType})
		t.AppendRow([]interface{}{"Country", country})
		t.AppendRow([]interface{}{"Image Profile", image})
		t.Render()

		fmt.Println()
	}
}
