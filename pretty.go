package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var data interface{}

	dec := json.NewDecoder(os.Stdin)
	err := dec.Decode(&data)
	if err != nil {
		log.Fatalf("json.Decode failed: %v", err)
	}

	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "    ")
	err = enc.Encode(data)
	if err != nil {
		log.Fatalf("json.Encode failed: %v", err)
	}
}
