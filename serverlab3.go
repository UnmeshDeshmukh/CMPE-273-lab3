package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	//"net/url"
)

type Inputkeyvaluepair struct {
	Key   int    `json:"key"`
	Value string `json:"value"`
}

type Response struct {
	Keyvaluearray []Inputkeyvaluepair `json:"response"`
}

var counter int

var Response1 Response
var Response2 Response
var Response3 Response

func gett(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//	id := p.ByName("id")

	// Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload

	uj, _ := json.MarshalIndent(Response1, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response received is as follows %s", uj)
}

func getHandler3000(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, _ := strconv.Atoi(p.ByName("id"))
	inputkeyvaluepair := Inputkeyvaluepair{}
	i := 0
	for i = 0; i < len(Response1.Keyvaluearray); i++ {
		if Response1.Keyvaluearray[i].Key == id {

			inputkeyvaluepair = Response1.Keyvaluearray[i]

			break
		}
	}
	var getResponse Response
	if i != len(Response1.Keyvaluearray) {
		

		getResponse.Keyvaluearray = append(getResponse.Keyvaluearray, inputkeyvaluepair)

	}
	// Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload

	uj, _ := json.MarshalIndent(getResponse, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response received is as follows %s", uj)
}

func gett2(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//	id := p.ByName("id")

	// Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload

	uj, _ := json.MarshalIndent(Response2, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response received is as follows %s", uj)
}

func getHandler3001(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id, _ := strconv.Atoi(p.ByName("id"))
	inputkeyvaluepair := Inputkeyvaluepair{}
	i := 0
	for i = 0; i < len(Response2.Keyvaluearray); i++ {
		if Response2.Keyvaluearray[i].Key == id {

			inputkeyvaluepair = Response2.Keyvaluearray[i]

			break
		}
	}
	var getResponse Response
	if i != len(Response2.Keyvaluearray) {
		

		getResponse.Keyvaluearray = append(getResponse.Keyvaluearray, inputkeyvaluepair)

	}
	// Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload

	uj, _ := json.MarshalIndent(getResponse, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response received is as follows %s", uj)
}

func gett3(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	//	id := p.ByName("id")

	// Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload

	uj, _ := json.MarshalIndent(Response3, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response received is as follows %s", uj)
}

func getHandler3002(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	inputkeyvaluepair := Inputkeyvaluepair{}
	i := 0
	for i = 0; i < len(Response3.Keyvaluearray); i++ {
		if Response3.Keyvaluearray[i].Key == id {

			inputkeyvaluepair = Response3.Keyvaluearray[i]

			break
		}
	}
	var getResponse Response
	if i != len(Response3.Keyvaluearray) {
		

		getResponse.Keyvaluearray = append(getResponse.Keyvaluearray, inputkeyvaluepair)

	}
	// Marshal provided interface into JSON structure

	// Write content-type, statuscode, payload

	uj, _ := json.MarshalIndent(getResponse, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response received is as follows %s", uj)
}

func putt(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	Id, _ := strconv.Atoi(id)
	value := p.ByName("value")
	//	counter += 1
	//fmt.Println("The counter is ----", counter)
	temp := Inputkeyvaluepair{}
	temp.Key = Id
	temp.Value = value
	//	response := Response{}
	Response1.Keyvaluearray = append(Response1.Keyvaluearray, temp)
	uj, _ := json.MarshalIndent(Response1, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response received is as follows %s", uj)

}

func putt2(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	Id, _ := strconv.Atoi(id)
	value := p.ByName("value")

	//counter += 1
	//fmt.Println("The counter is ----", counter)
	temp := Inputkeyvaluepair{}
	temp.Key = Id
	temp.Value = value
	//	response := Response{}
	Response2.Keyvaluearray = append(Response2.Keyvaluearray, temp)
	uj, _ := json.MarshalIndent(Response2, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response received is as follows %s", uj)

}
func putt3(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	Id, _ := strconv.Atoi(id)
	value := p.ByName("value")
	//counter += 1
	//fmt.Println("The counter is ----", counter)
	temp := Inputkeyvaluepair{}
	temp.Key = Id
	temp.Value = value
	//	response := Response{}
	Response3.Keyvaluearray = append(Response3.Keyvaluearray, temp)
	uj, _ := json.MarshalIndent(Response3, "", "\t")
	//fmt.Println(uj)
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, "The JSON response received is as follows %s", uj)

}

func main() {
	counter = 1

	r1 := httprouter.New()
	r2 := httprouter.New()
	r3 := httprouter.New()
	r1.GET("/keys", gett)
	//r.POST("/keys", postt)
	r2.GET("/keys", gett2)
	r3.GET("/keys", gett3)

	r1.GET("/keys/:id", getHandler3000)
	//r.POST("/keys", postt)
	r2.GET("/keys/:id", getHandler3001)
	r3.GET("/keys/:id", getHandler3002)
	r1.PUT("/keys/:id/:value", putt)

	r2.PUT("/keys/:id/:value", putt2)

	r3.PUT("/keys/:id/:value", putt3)
	go func() {
		server3 := http.Server{
			Addr:    "0.0.0.0:3000",
			Handler: r1,
		}
		server3.ListenAndServe()
	}()

	go func() {
		server1 := http.Server{
			Addr:    "0.0.0.0:3001",
			Handler: r2,
		}
		server1.ListenAndServe()
	}()
	server2 := http.Server{
		Addr:    "0.0.0.0:3002",
		Handler: r3,
	}
	server2.ListenAndServe()
}
