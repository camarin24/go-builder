package models

type ProjectModel struct {
	BaseModel
	Name string
}

func (p ProjectModel) TableName() string {
	return "projects"
}
