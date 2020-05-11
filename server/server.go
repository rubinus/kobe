package server

import (
	"context"
	"kobe/api"
	"time"
)

type Kobe struct{}

func (k Kobe) CreateProject(ctx context.Context, req *api.CreateProjectRequest) (*api.CreateProjectResponse, error) {
	pm := ProjectManager{}

	p, err := pm.CreateProject(req.Name, req.Source)
	if err != nil {
		return nil, err
	}
	resp := &api.CreateProjectResponse{
		Item: p,
	}
	return resp, nil
}
func (k Kobe) ListProject(ctx context.Context, req *api.ListProjectRequest) (*api.ListProjectResponse, error) {
	pm := ProjectManager{}
	ps, err := pm.SearchProjects()
	if err != nil {
		return nil, err
	}
	resp := &api.ListProjectResponse{
		Items: ps,
	}
	return resp, nil
}

func (k Kobe) GetInventory(ctx context.Context, req *api.GetInventoryRequest) (*api.GetInventoryResponse, error) {
	item := InventoryCache.Get(req.Id)
	resp := &api.GetInventoryResponse{
		Item: item,
	}
	return resp, nil
}

func (k Kobe) RunPlaybook(req *api.RunPlaybookRequest, server api.KobeApi_RunPlaybookServer) error {
	rm := RunnerManager{}
	runner, _ := rm.CreatePlaybookRunner(req.Project, req.Playbook, req.Inventory)
	ch := make(chan []byte)
	var result = api.Result{
		StartTime: time.Now().String(),
		EndTime:   nil,
		Message:   "",
		Success:   false,
		Content:   "",
	}
	runner.Run(ch, &result)
	go func() {
		for buf := range ch {
			_ = server.Send(&api.WatchStream{
				Stream: buf,
				Result: &result,
			})
		}
	}()
	return nil
}
