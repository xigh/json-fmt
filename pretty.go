package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	dec := json.NewDecoder(os.Stdin)
	var data interface{}
	err := dec.Decode(&data)
	if err != nil {
		log.Fatalf("json.Decode failed")
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	err = enc.Encode(data)
	if err != nil {
		log.Fatalf("json.Encode failed")
	}
}
