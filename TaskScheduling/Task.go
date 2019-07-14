package TaskScheduling

import (
	"fmt"
	"strings"
)

type Task struct {
	Name string
	Runtime uint32
}

type Tasks []*Task

func (t Tasks) Sort() {
	i, j := 0, 0
	var tmp *Task
	for i = 0; i < len(t)-1; i++ {
		for j = 0; j < len(t)-1-i; j++ {
			if t[j].Runtime > t[j+1].Runtime {
				tmp = t[j]
				t[j] = t[j+1]
				t[j+1] = tmp
			}
		}
	}
}

func (t *Task) String() string {
	return fmt.Sprintf("%s with Runtime %d", t.Name, t.Runtime)
}

func (ts *Tasks) String() string {
	var strList []string
	for ind, elem := range(*ts) {
		strList = append(strList, fmt.Sprintf("Task #%d: %s", ind, elem))
	}
	return strings.Join(strList, "\n")
}
