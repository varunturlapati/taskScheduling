package TaskScheduling

import "fmt"

type Policy int

const (
	RoundRobin Policy = iota
	SortAndZigZag
	SortAndOneDirection
)

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

func RoundRobinImpl(servers Servers, tasks Tasks) {
	for _, task := range(tasks) {
		for _, server := range(servers) {
			server.AssignTask(&task)
		}
	}

}

func SortAndZigZagImpl(servers Servers, tasks Tasks) {
	fmt.Printf("Not impl")
}

func SortAndRoundRobinImpl(servers Servers, tasks Tasks) {
	fmt.Printf("Not impl")
}