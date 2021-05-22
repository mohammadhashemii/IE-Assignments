package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type Task struct{
	Task		string		`json:"task"`
	Numbers		[]int		`json:"numbers"`
}

// IntNumbers :to check if the input numbers are integer or not
type IntNumbers struct {
	Numbers		[]int		`json:"numbers"`
}

type TaskResponse struct {
	Task			string		`json:"task,omitempty"`
	Numbers			[]int		`json:"numbers,omitempty"`
	Answer			float32		`json:"answer,omitempty"`
	Sorted			[]int		`json:"sorted,omitempty"`
	Code			int			`json:"code"`
	Message			string		`json:"message"`
}

type HistoryResponse struct {
	Size 		int 				`json:"size"`
	History 	[]TaskResponse		`json:"history"`
	Code		int					`json:"code"`
	Message 	string				`json:"message"`
}

// A function to return the mean of a slice of numbers
func (t *Task) mean() float32 {
	sum := 0
	for _, num := range t.Numbers {
		sum += num
	}
	return float32(sum) / float32(len(t.Numbers))
}

func BadEndPoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := TaskResponse{}
	resp.Code = http.StatusNotImplemented
	resp.Message = "Your Endpoint has not yet implemented!"

	w.WriteHeader(resp.Code)
	enc := json.NewEncoder(w)
	enc.SetIndent("", "	")
	enc.Encode(resp)
}

var history = HistoryResponse{}

func Calculate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
		case "POST" :
			body, _ := ioutil.ReadAll(r.Body)
			resp := TaskResponse{}

			intNumbers := IntNumbers{}
			notIntegerError := json.Unmarshal(body, &intNumbers)
			if notIntegerError != nil {

				resp.Code = http.StatusNotImplemented
				resp.Message = "You entered non-integer numbers!"
				w.WriteHeader(resp.Code)
				enc := json.NewEncoder(w)
				enc.SetIndent("", "	")
				enc.Encode(resp)
				break
			}

			t := Task{}
			// json.Unmarshal(body, &t)
			dec := json.NewDecoder(bytes.NewReader(body))
			dec.DisallowUnknownFields() // Force errors

			if err := dec.Decode(&t); err != nil {
				resp.Code = http.StatusNotImplemented
				resp.Message = "You entered unknown fields!"
				w.WriteHeader(resp.Code)
				enc := json.NewEncoder(w)
				enc.SetIndent("", "	")
				enc.Encode(resp)
				return
			}

			switch t.Task {
			case "mean":
				resp.Task = t.Task
				resp.Numbers = t.Numbers
				resp.Answer = t.mean()
				resp.Code = http.StatusOK
				resp.Message = "Task done successfully!"
				history.Size = history.Size + 1
				history.History = append(history.History, resp)

			case "sort":
				resp.Task = t.Task
				resp.Numbers = t.Numbers
				resp.Sorted = append([]int(nil), resp.Numbers...)
				sort.Ints(resp.Sorted)
				resp.Code = http.StatusOK
				resp.Message = "Task done successfully!"
				history.Size = history.Size + 1
				history.History = append(history.History, resp)
			default:	// if the user entered some un-supported tasks
				resp.Task = t.Task
				resp.Numbers = t.Numbers
				resp.Code = http.StatusNotImplemented
				resp.Message = "Undefined task! Only mean and sort are supported!"
			}



			w.WriteHeader(resp.Code)
			enc := json.NewEncoder(w)
			enc.SetIndent("", "	")
			enc.Encode(resp)
		default:
			resp := TaskResponse{}
			resp.Code = http.StatusMethodNotAllowed
			resp.Message = "This HTTP method is not allowed!"
			w.WriteHeader(resp.Code)
			enc := json.NewEncoder(w)
			enc.SetIndent("", "	")
			enc.Encode(resp)
	}

}

func History(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		history.Code = http.StatusOK
		history.Message = "History sent successfully!"
		w.WriteHeader(history.Code)
		enc := json.NewEncoder(w)
		enc.SetIndent("", "	")
		enc.Encode(history)
	default:
		resp := TaskResponse{}
		resp.Code = http.StatusMethodNotAllowed
		resp.Message = "This HTTP method is not allowed!"
		w.WriteHeader(resp.Code)
		enc := json.NewEncoder(w)
		enc.SetIndent("", "	")
		enc.Encode(resp)
	}

}


func main() {
	http.HandleFunc("/calculator", Calculate)
	http.HandleFunc("/history", History)
	http.HandleFunc("/", BadEndPoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}