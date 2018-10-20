package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var attackType = [...]string{
	"Type1",
	"Type2",
	"Type3",
	"Type4",
}

func main() {

	fmt.Println("Hello World")

	client := &http.Client{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for {
		source1 := 11 + r.Int63n(181)
		source2 := r.Int63n(256)
		source3 := r.Int63n(256)
		source4 := r.Int63n(256)

		source := fmt.Sprintf("%v.%v.%v.%v", source1, source2, source3, source4)

		dest1 := 11 + r.Int63n(181)
		dest2 := r.Int63n(256)
		dest3 := r.Int63n(256)
		dest4 := r.Int63n(256)

		dest := fmt.Sprintf("%v.%v.%v.%v", dest1, dest2, dest3, dest4)

		selectedType := attackType[rand.Intn(len(attackType))]

		body := map[string]string{
			"srcIP":      source,
			"dstIP":      dest,
			"attackType": selectedType,
		}

		jsonBody, err := json.Marshal(body)

		if err != nil {
			fmt.Println(err.Error())
			panic(err)
		}

		_, err = client.Post("http://localhost:3000", "application/json", bytes.NewReader(jsonBody))
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Duration(1+r.Int63n(9)) * time.Second)
	}
}
