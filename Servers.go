package taskScheduling

import "fmt"

type Server struct {
	Name string
	TaskList Tasks
	TotalRuntime uint32
}

type Servers []Server

func (s *Server) String() string {
	return fmt.Sprintf("Server %s", (*s).Name)
}

func (s *Server) AssignTask(t *Task) error {
	if s.TaskList == nil {
		s.TaskList = make(Tasks, 0)
	}
	s.TaskList = append(s.TaskList, *t)
	return nil
}

func (s *Server) ShowTasks() {
	if s == nil {
		fmt.Errorf("Server doesn't exist")
	}
	fmt.Printf("Server %s has the following tasks:\n %s", (*s).Name, (*s).TaskList)
}


