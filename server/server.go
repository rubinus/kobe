package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"
	"kobe/api"
	"time"
)

type Kobe struct {
	taskCache      *cache.Cache
	inventoryCache *cache.Cache
}

func NewKobe() *Kobe {
	return &Kobe{
		taskCache:      cache.New(24*time.Hour, 5*time.Minute),
		inventoryCache: cache.New(10*time.Minute, 5*time.Minute),
	}
}

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
	item, _ := k.inventoryCache.Get(req.Id)
	resp := &api.GetInventoryResponse{
		Item: item.(*api.Inventory),
	}
	return resp, nil
}

func (k Kobe) RunPlaybook(req *api.RunPlaybookRequest, server api.KobeApi_RunPlaybookServer) error {
	rm := RunnerManager{
		inventoryCache: k.inventoryCache,
	}
	runner, _ := rm.CreatePlaybookRunner(req.Project, req.Playbook, req.Inventory)
	ch := make(chan []byte)
	id := uuid.NewV4().String()
	result := api.Result{
		Id:        id,
		StartTime: time.Now().String(),
		EndTime:   "",
		Message:   "",
		Success:   false,
		Finished:  false,
		Content:   "",
	}
	taskId := uuid.NewV4().String()
	go func() {
		k.taskCache.Set(taskId, &result, cache.DefaultExpiration)
		fmt.Println("taskId" + taskId)
		runner.Run(ch, &result)
		k.taskCache.Set(taskId, &result, cache.DefaultExpiration)
	}()
	for buf := range ch {
		fmt.Print(string(buf))
		_ = server.Send(&api.WatchStream{
			Stream: buf,
			Result: &result,
		})
	}
	return nil
}

func (k Kobe) SaveResult(ctx context.Context, req *api.SaveResultRequest) (*api.SaveResultResponse, error) {
	k.taskCache.Set(req.Item.Id, req.Item, cache.DefaultExpiration)
	return &api.SaveResultResponse{}, nil
}

func (k Kobe) GetResult(ctx context.Context, req *api.GetResultRequest) (*api.GetResultResponse, error) {
	id := req.GetTaskId()
	result, found := k.taskCache.Get(id)
	if !found {
		return nil, errors.New(fmt.Sprintf("can not find task: %s result", id))
	}
	return &api.GetResultResponse{Item: result.(*api.Result)}, nil
}
