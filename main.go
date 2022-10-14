package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

//input data
type input struct {
	Observations observations `json:"observations"`
}

//observations key
type observations struct {
	Data []data `json:"data"`
}

//data key
type data struct {
	Name      string  `json:"name"`
	Lat       float64 `json:"lat"`
	Lon       float64 `json:"lon"`
	ApparentT float64 `json:"apparent_t"`
}

//output
type output struct {
	Data []data `json:"data"`
}

//error message
type errMsg struct {
	Msg string `json:"error"`
}

//extract data from input if input is correct.
func extractSortedData(input input) output {
	result := make([]data, 0)
	for _, v := range input.Observations.Data {
		if v.ApparentT > 15 {
			result = append(result, v)
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ApparentT < result[j].ApparentT
	})
	output := output{Data: result}
	return output
}

//process data path handler
func processDataHandler(w http.ResponseWriter, req *http.Request) {
	/*
		resp, err := http.Get("http://www.bom.gov.au/fwo/IDN60801/IDN6080195765.json")
		if err != nil {
			log.Println(err)
			http.Error(w, string(errMessageOutput), http.StatusServiceUnavailable)
			return
		}
		body, err := ioutil.ReadAll(resp.Body)
	*/
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	errMessage := errMsg{Msg: "Error Connecting to BOM."}
	errMessageOutput, _ := json.MarshalIndent(errMessage, "", "  ")
	body, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Println(err)
		http.Error(w, string(errMessageOutput), http.StatusServiceUnavailable)
		return
	}
	var input input

	err = json.Unmarshal(body, &input)
	if err != nil {
		log.Println(err)
		http.Error(w, string(errMessageOutput), http.StatusServiceUnavailable)
		return
	}
	output := extractSortedData(input)
	res, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		log.Println(err)
		http.Error(w, string(errMessageOutput), http.StatusServiceUnavailable)
		return
	}
	fmt.Fprintf(w, string(res))
}

//main function
func main() {
	http.HandleFunc("/", processDataHandler)
	port := os.Getenv("PORT")
	log.Printf("Starting server ")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Printf("Error in starting server > %v", err)
	}
}
