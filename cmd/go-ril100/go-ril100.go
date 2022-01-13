package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	baseUrl = "https://v5.db.transport.rest/stations/"
)

type loc struct {
	Lat float64 `json:"latitude"`
	Lng float64 `json:"longitude"`
}

type addr struct {
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Street  string `json:"street"`
}

type rilObject struct {
	Kind     string `json:"type"`
	Name     string `json:"name"`
	Address  addr   `json:"address"`
	Location loc    `json:"location"`
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	p := ""
	if len(os.Args) > 2 {
		p = strings.Trim(strings.Trim(os.Args[2], " "), "-")
	}
	q := strings.Trim(string(os.Args[1]), " ")
	if strings.Trim(q, "-") == "h" || strings.Trim(q, "-") == "help" || p == "h" || p == "help" {
		usage()
		return
	}
	if q == "" {
		usage()
		return
	}
	res, err := query(q)
	if err != nil {
		fmt.Println(err)
		return
	}

	if res.Name == "" {
		fmt.Println("No Result")
		return
	}

	if p == "json" {
		if err != nil {
			fmt.Println(err)
			return
		}
		js, _ := json.Marshal(res)
		fmt.Println(string(js))
	} else {
		fmt.Printf("Name     : %s\n", res.Name)
		fmt.Printf("Type     : %s\n", res.Kind)
		fmt.Printf("Address  :\n\tStreet : %s\n\tCity   : %s %s\n", res.Address.Street, res.Address.Zipcode, res.Address.City)
		fmt.Printf("Location :\n\tLatitude  :%f\n\tLongitude :%f\n", res.Location.Lat, res.Location.Lng)
	}
	return
}

func usage() {
	fmt.Println()
	fmt.Println("Usage")
	fmt.Println("-----")
	fmt.Println(filepath.Base(os.Args[0]) + " <RIL100> [--json]")
	fmt.Println(filepath.Base(os.Args[0]) + " -h")
	fmt.Println()
}

func query(q string) (rilObject, error) {
	ro := rilObject{}
	resp, err := http.Get(baseUrl + q)
	if err != nil {
		return ro, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ro, err
	}
	err = json.Unmarshal([]byte(body), &ro)
	if err != nil {
		return ro, err
	}
	return ro, nil
}
