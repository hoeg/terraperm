package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hoeg/terraperm/module/policy"
)

func main() {
	var out string
	flag.StringVar(&out, "out", "-", "Location to output the policy, defaults to stdout.")
	flag.Parse()

	/*
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
	*/
	b, err := os.Open("../../test/test_data.txt")
	if err != nil {
		fmt.Printf("Failed to load trace: %v\n", err)
		return
	}

	p := policy.NewParser(b)
	reqs, err := p.Requests()
	if err != nil {
		fmt.Printf("Failed to parse requests: %v\n", err)
		return
	}

	stmts := policy.NewStatements()
	stmts.AddRequests(reqs)

	iamPolicy, err := policy.Print(stmts.List())
	if err != nil {
		fmt.Printf("Failed to print policy %v\n", err)
		return
	}
	fmt.Printf("%v\n", iamPolicy)
	fmt.Printf("Done\n")
}
