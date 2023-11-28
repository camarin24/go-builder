package services

import (
	"context"

	"github.com/camarin24/go-studio/types/models"
)

func (s *Service) GetProjects(ctx context.Context) ([]models.ProjectModel, error) {
	var projects []models.ProjectModel
	err := s.db.With(ctx).Select().OrderBy("created").All(&projects)
	return projects, err
}
