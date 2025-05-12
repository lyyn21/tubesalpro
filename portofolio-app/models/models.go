package models

type Project struct {
	ID          int
	Name        string
	Description string
	Skills      []string
	Year        int
}

var Projects []Project

func AddProject(p Project) {
	Projects = append(Projects, p)
}

func GetAllProjects() []Project {
	return Projects
}
