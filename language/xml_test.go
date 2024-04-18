package gotutorial

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/AndiHappy/auxiliaryG/util"
	"log"
	"os"
	"testing"
)

type Notes struct {
	To      string `xml:"to"`
	From    string `xml:"from"`
	Heading string `xml:"heading"`
	Body    string `xml:"body"`
}

func TestXml(t *testing.T) {
	data, _ := os.ReadFile("notes.xml")
	note := &Notes{}
	_ = xml.Unmarshal([]byte(data), &note)
	fmt.Println(note.To)
	fmt.Println(note.From)
	fmt.Println(note.Heading)
	fmt.Println(note.Body)
	util.DeleteFile("notes.xml")
	bytes, _ := xml.Marshal(note)
	util.AppendFileContent(string(bytes), "notes.xml")
}

func TestXml2(t *testing.T) {
	note := &Notes{To: "Nicky",
		From:    "Rock",
		Heading: "Meeting",
		Body:    "Meeting at 5pm!",
	}
	file, _ := xml.MarshalIndent(note, "", " ")
	_ = os.WriteFile("notes1.xml", file, 0644)
}

type CatalogNodes struct {
	CatalogNodes []Catalog `json:"catlog_nodes"`
}

type Catalog struct {
	ProductId string `json: "product_id"`
	Quantity  int    `json: "quantity"`
}

func TestJson(t *testing.T) {
	type Salary struct {
		Basic, HRA, TA float64
	}
	type Employee struct {
		FirstName, LastName, Email string
		Age                        int
		MonthlySalary              []Salary
	}
	data := Employee{
		FirstName: "Mark",
		LastName:  "Jones",
		Email:     "mark@gmail.com",
		Age:       25,
		MonthlySalary: []Salary{
			Salary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			Salary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			Salary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	file, _ := json.MarshalIndent(data, "", " ")

	_ = os.WriteFile("test.json", file, 0644)
}

func TestCSV(t *testing.T) {
	rows := [][]string{
		{"Name", "City", "Language"},
		{"Pinky", "London", "Python"},
		{"Nicky", "Paris", "Golang"},
		{"Micky", "Tokyo", "Php"},
	}

	csvfile, err := os.Create("test.csv")
	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvfile)
	for _, row := range rows {
		_ = csvwriter.Write(row)
	}
	csvwriter.Flush()
	csvfile.Close()
}
