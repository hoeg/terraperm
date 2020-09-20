package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/hoeg/terraperm/module/terraform"
)

func main() {
	var out string
	flag.StringVar(&out, "out", "-", "Location to output the policy, defaults to stdout.")
	flag.Parse()

	exe, err := terraform.NewExecutor()
	if err != nil {
		fmt.Printf("Failed to create executor: %v\n", err)
		return
	}
	b := bytes.NewBuffer(nil)
	t := terraform.NewTracer(exe)

	err = t.MakeTrace(b)
	if err != nil {
		fmt.Printf("Failed to create trace: %v\n", err)
	}

	data, err := ioutil.ReadAll(b)
	if err != nil {
		fmt.Printf("Falied to consume trace: %v\n", err)
	}

	fmt.Printf("Trace: %s\n\n", string(data))
	fmt.Printf("Done\n")
}
