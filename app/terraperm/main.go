package main

import (
	"flag"
	"fmt"
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
	t := terraform.Tracer{exe}
	trace, err := t.makeTrace()
	if err != nil {
		fmt.Printf("Failed to create trace: %v\n", err)
	}
	/*
		- output filename as argument
		- wrap terraform binary
		- set TF_LOG=DEBUG
		- run terraform apply -auto-approve and record stderr
		- run terraform destroy -auto-approve
	 	- parse output to create policy
	*/
	fmt.Printf("Done\n")
}