package kobe

import (
	"context"
	"kobe/api"
)

type Kobe struct{}

func (k Kobe) CreateProject(context.Context, *api.CreateProjectRequest) (*api.CreateProjectResponse, error) {
	return nil, nil
}
func (k Kobe) ListProject(context.Context, *api.ListProjectRequest) (*api.ListProjectResponse, error) {
	return &api.ListProjectResponse{}, nil
}
func (k Kobe) RunPlaybook(context.Context, *api.RunPlaybookRequest) (*api.RunPlaybookResponse, error) {
	return &api.RunPlaybookResponse{}, nil
}
func (k Kobe) GetResult(context.Context, *api.GetResultRequest) (*api.GetResultRequest, error) {
	return &api.GetResultRequest{}, nil
}
