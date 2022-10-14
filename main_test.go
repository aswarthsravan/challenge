package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

var (
	data1 = data{Name: "test1", Lat: 2.4, Lon: 2.5, ApparentT: 15.3}
	data2 = data{Name: "test2", Lat: 2.6, Lon: 2.7, ApparentT: 15.2}
	data3 = data{Name: "test3", Lat: 2.8, Lon: 2.9, ApparentT: 14.9}
	data4 = data{Name: "test4", Lat: 3.1, Lon: 3.2, ApparentT: 15.0}
)

//TestExtractSortedData unit tests basic functionality of extractSortedData
func TestExtractSortedData(t *testing.T) {
	inputData := []data{data1, data2, data3}
	input := input{Observations: observations{Data: inputData}}
	output := extractSortedData(input)
	if len(output.Data) != 2 || output.Data[0] != data2 || output.Data[1] != data1 {
		t.Errorf("extractSortedData is wrong")
	}
}

//TestExtractSortedDataEmpty unit tests that all <= 15.0 input gives empty output
func TestExtractSortedDataEmpty(t *testing.T) {
	inputData := []data{data3, data4}
	input := input{Observations: observations{Data: inputData}}
	output := extractSortedData(input)
	if len(output.Data) != 0 {
		t.Errorf("extractSortedData is wrong")
	}
}

//TestEndToEndPass tests the output end to end for given json
func TestEndToEndPass(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(processDataHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}

/*
//TestEndToEndFail tests the output end to end for a failed json
func TestEndToEndFail(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(processDataHandler)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusServiceUnavailable {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusServiceUnavailable)
	}

}
*/
