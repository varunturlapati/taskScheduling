package main

import (
	ts "TaskScheduling/TaskScheduling"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	serverJsonPtr := flag.String("servers", "", "Provide path to the JSON containing server info")
	taskJsonPtr := flag.String("tasks", "", "Provide path to the JSON containing task info")
	schedPolicyPtr := flag.String("policy", "RR", "Specify RR, SRR, SZZ, case insensitive")
	flag.Parse()

	var servers ts.Servers
	var tasks ts.Tasks
	var policy ts.Policy
	var ok bool
	var policyMap map[string]ts.Policy = map[string]ts.Policy{
		"RR": ts.RoundRobin,
		"SRR": ts.SortAndOneDirection,
		"SZZ": ts.SortAndZigZag,
	}

	srvStr, err := ioutil.ReadFile(*serverJsonPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading file under %s resulted in error: %v", *serverJsonPtr, err)
		os.Exit(1)
	}
	tskStr, err := ioutil.ReadFile(*taskJsonPtr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Reading file under %s resulted in error: %v", *taskJsonPtr, err)
		os.Exit(1)
	}

	err = json.Unmarshal(srvStr, &servers)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse the json file to get servers because of: %v", err)
		os.Exit(1)
	}
	err = json.Unmarshal(tskStr, &tasks)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse the json file to get tasks because of: %v", err)
		os.Exit(1)
	}

	if policy, ok = policyMap[strings.ToUpper(*schedPolicyPtr)]; ok == false {
		fmt.Fprintf(os.Stderr, "Policy %s is not supported", *schedPolicyPtr)
		os.Exit(1)
	}

	ts.DistributeTasks(servers, tasks, policy)
	jsonBytes, err := json.Marshal(servers)
	if err != nil {
		fmt.Printf("Error generating output file allocation.json")
	}
	ioutil.WriteFile("allocation.json", jsonBytes, os.FileMode(777))
}
