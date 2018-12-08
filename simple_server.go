package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/tealeg/xlsx"
)

type Kos struct {
	Nama     string `json:"nama"`
	Location int    `json:"location"`
	Room     [4]int `json:"rom"`
}

func main() {
	port := 9090

	http.HandleFunc("/", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/template", test_template)
	http.HandleFunc("/api", api_test)
	http.HandleFunc("/api_excel", api_excel)

	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Go Index")
}

func home(w http.ResponseWriter, r *http.Request) {
	var x [5]float64
	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83
	fmt.Println(x)
}

func test_template(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("template/template.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	template.ExecuteTemplate(w, "template", nil)
	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func api_test(w http.ResponseWriter, r *http.Request) {
	data := Kos{"Kos Mantan", 20, [4]int{1, 0, 0, 1}}
	json_data, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == "POST" {
		fmt.Println("Test Api")
	} else if r.Method == "GET" {
		fmt.Println("Get API")
		w.Header().Set("Content-Type", "application/json")
		w.Write(json_data)
		fmt.Println(json_data)
	}
}

func api_excel(w http.ResponseWriter, r *http.Request) {
	xlFile, err := xlsx.OpenFile("/home/rhyanz46/Documents/test.xlsx")
	if err != nil {
		fmt.Println("error gan")
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\n", text)
			}
		}
	}
	// fmt.Println(exl_data)
}
