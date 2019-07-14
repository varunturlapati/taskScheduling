package TaskScheduling

import (
	"fmt"
	"strings"
)

type Server struct {
	Name         string
	TaskList     Tasks
	TotalRuntime uint32
}

type Servers []*Server

func (s *Server) AssignTask(t *Task) error {
	if s.TaskList == nil {
		s.TaskList = make(Tasks, 0)
	}
	s.TaskList = append(s.TaskList, t)
	s.UpdateRuntime(t)
	return nil
}

func (s *Server) UpdateRuntime(t *Task) {
	s.TotalRuntime += t.Runtime
}

func (s *Server) ShowTasks() {
	if s == nil {
		fmt.Printf("Server doesn't exist")
	}
	fmt.Printf("Server %s has the following tasks:\n %s", s.Name, s.TaskList)
}

func (s *Server) String() string {
	return fmt.Sprintf("Server %s --> Tasks %s\n%s's total runtime = %d", s.Name, s.TaskList, s.Name, s.TotalRuntime)
}

func (s Servers) String() string {
	var strList []string
	for ind, elem := range(s) {
		strList = append(strList, fmt.Sprintf("Server #%d: %s", ind, elem))
	}
	return strings.Join(strList, "\n")
}

