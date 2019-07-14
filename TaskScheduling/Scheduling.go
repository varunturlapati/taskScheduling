package TaskScheduling

import (
	"fmt"
)

type Policy int

const (
	RoundRobin Policy = iota
	SortAndZigZag
	SortAndOneDirection
)
var err error

func DistributeTasks(S Servers, T Tasks, policy Policy) {
	switch policy {
	case RoundRobin:
		RoundRobinImpl(S, T)

	case SortAndZigZag:
		SortAndZigZagImpl(S, T)

	case SortAndOneDirection:
		SortAndRoundRobinImpl(S, T)
	}
}

//Given servers and tasks, tasks are assigned to servers in a round robin fashion
func RoundRobinImpl(servers Servers, tasks Tasks) {
	var srvInd = 0
	for _, task := range(tasks) {
		err = servers[srvInd].AssignTask(task)
		if err != nil {
			fmt.Printf("Could not assign Task %s to Server %s", task, servers[srvInd])
		}
		srvInd = (srvInd + 1) % len(servers)
	}
	fmt.Printf("The allocation is: %s", servers)
}

//Given servers and tasks, tasks are sorted and then assigned to servers in a round robin fashion but also changing
//direction when one end of the servers is reached, and so on
func SortAndZigZagImpl(servers Servers, tasks Tasks) {
	var srvInd = 0
	tasks.Sort()
	var seq = zigzagGenerator(len(servers))
	fmt.Println(seq)
	for tInd, task := range(tasks) {
		srvInd = seq[tInd % len(seq)]
		fmt.Printf("srvInd = %d\n", srvInd)
		err = servers[srvInd].AssignTask(task)
		if err != nil {
			fmt.Printf("Could not assign Task %s to Server %s", task, servers[srvInd])
		}
	}
	fmt.Printf("The allocation is:\n%s\n", servers)
}

//Given servers and tasks, tasks are first sorted (ascending) and then assigned to servers in a round robin fashion
func SortAndRoundRobinImpl(servers Servers, tasks Tasks) {
	var srvInd = 0
	tasks.Sort()
	for _, task := range(tasks) {
		err = servers[srvInd].AssignTask(task)
		if err != nil {
			fmt.Printf("Could not assign Task %s to Server %s", task, servers[srvInd])
		}
		srvInd = (srvInd + 1) % len(servers)
	}
	fmt.Printf("The allocation is: %s", servers)
}

//Helper to generate zigzag order because % operator returns -ve ints and so wraparound is complex to implement
func zigzagGenerator(lim int) []int {
	var sequence []int
	for i := 0; i < lim; i++ {
		sequence = append(sequence, i)
	}
	for i := lim-1; i >= 0; i-- {
		sequence = append(sequence, i)
	}
	return sequence
}