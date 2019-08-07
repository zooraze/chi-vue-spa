package main

import (
	"time"
	"net/http"
	"encoding/json"
)

type Routes struct {
	staticDir    string
	apiKey			 string
	apiSecret		 string
}

// Route: Data
type Data struct {
	Query      string   `json:"query"`
	Datetime   string   `json:"datetime"`
}

// /api/data
func (r Routes) apiDataRoute(resp http.ResponseWriter, req *http.Request) {	
	var data Data

	// TODO(zooraze): actually grab data contents
	data.Query = "This is not the data you're looking for..."
	// TODO(zooraze): the time should not update every five seconds
	datetime := time.Now().Format("Mon Aug 5 10:17:08 EST 2019")
	data.Datetime = datetime 
 
	// Reformat data as JSON
	result, err := json.Marshal(data)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	resp.Write(result)
}
