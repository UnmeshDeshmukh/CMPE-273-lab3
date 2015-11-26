package main

import (
	"fmt"
	"hash/crc32"
	"net/http"
	"bytes"
	"io/ioutil"
	"strconv"
)

type InputStructure struct {
	Keys   int    `json:"key"`
	Value string `json:"value"`
}

var keys []InputStructure
var Server0 = "3000"
var Server1 = "3001"
var Server2 = "3002"
var servernumber = 3

type HashStructure struct {
	circular []StructNode
}

type StructNode struct {
	hashValue uint32
	nameOfNode  string
}

func (varCirclular *HashStructure) Serveradder(SERVERNAME string) {

	var nodeofstruct StructNode
	nodeofstruct.hashValue = hashing(SERVERNAME)
	nodeofstruct.nameOfNode = SERVERNAME
	varCirclular.circular = append(varCirclular.circular, nodeofstruct)

}


func (varCirclular *HashStructure) adder(inputkeystructure InputStructure) {

	for i := 0; i < servernumber; i++ {

		if varCirclular.circular[i].hashValue > hashing(strconv.Itoa(inputkeystructure.Keys)) {
			//fmt.Println("The key is stored in Server --", varCirclular.circular[i])
			putt(varCirclular.circular[i].nameOfNode, inputkeystructure)
			return
		}
	}
	
	putt(varCirclular.circular[0].nameOfNode, inputkeystructure)
}

func (varCirclular *HashStructure) gett(inputkeystructure InputStructure) {

	for i := 0; i < servernumber; i++ {

		if varCirclular.circular[i].hashValue > hashing(strconv.Itoa(inputkeystructure.Keys)) {
			fmt.Println("The key is stored in Server --", varCirclular.circular[i])
			gett(varCirclular.circular[i].nameOfNode, inputkeystructure)
			return
		}
	}
	fmt.Println("The key is stored in Server--", varCirclular.circular[0])
	gett(varCirclular.circular[0].nameOfNode, inputkeystructure)

}



func prepareInput() {
	//var temp string
	for i := 0; i < 10; i++ {
		keys = append(keys, InputStructure{})
		keys[i].Keys = i + 1
		//keys[i].Value = temp
	}
	keys[0].Value = "a"
	keys[1].Value = "b"
	keys[2].Value = "c"
	keys[3].Value = "d"
	keys[4].Value = "e"
	keys[5].Value = "f"
	keys[6].Value = "g"
	keys[7].Value = "h"
	keys[8].Value = "i"
	keys[9].Value = "j"

	//fmt.Println("The keys are---", keys)
}

func hashing(toBeHashed string) uint32 {
	hash := crc32.ChecksumIEEE
	hashed := hash([]byte(toBeHashed))
	//king := math.Mod(float64(hashing), 3)
	//fmt.Println(toBeHashed, "hashed to", hashing)
	return hashed
}




func gett(SERVERNAME string, inputkeystructure InputStructure) {

	stringUrl := "http://localhost:" + SERVERNAME + "/keys/" + strconv.Itoa(inputkeystructure.Keys)
	//fmt.Println("GET", stringUrl)
	res, _ := http.Get(stringUrl)
	//fmt.Println(Url.String())
	
	defer res.Body.Close()
	bodycontent, _ := ioutil.ReadAll(res.Body)
	fmt.Printf("%s\n", string(bodycontent))
}



func putt(SERVERNAME string, inputkeystructure InputStructure) {
	client := &http.Client{}
	stringUrl := "http://localhost:" + SERVERNAME + "/keys/" + strconv.Itoa(inputkeystructure.Keys) + "/" + inputkeystructure.Value


	//stringUrl:=makePUT_URL(hash,keys[i])
	jsonStr := []byte(`{}`)
	r1, _ := http.NewRequest("PUT", stringUrl, bytes.NewBuffer(jsonStr))

	r1.Header.Set("Content-Type", "application/json")
	r1.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp1, _ := client.Do(r1)
	
	fmt.Println(resp1.StatusCode)
	defer resp1.Body.Close()
}
func main() {
	prepareInput()
	var constanthash HashStructure
	constanthash.Serveradder(Server0)
	constanthash.Serveradder(Server1)
	constanthash.Serveradder(Server2)
	for i:=0;i<10;i++{
	constanthash.adder(keys[i])	
	}
	for i:=0;i<10;i++{
	constanthash.gett(keys[i])
	}	
}

