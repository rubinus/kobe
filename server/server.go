package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"kobe/api"
	"kobe/pkg/constant"
	"path"
	"time"
)

type Kobe struct {
	taskCache      *cache.Cache
	inventoryCache *cache.Cache
	chCache        *cache.Cache
}

func NewKobe() *Kobe {
	return &Kobe{
		taskCache:      cache.New(24*time.Hour, 5*time.Minute),
		chCache:        cache.New(24*time.Hour, 5*time.Minute),
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

func (k Kobe) WatchRunPlaybook(req *api.WatchPlaybookRequest, server api.KobeApi_WatchRunPlaybookServer) error {
	ch, found := k.chCache.Get(req.TaskId)
	if !found {
		return errors.New(fmt.Sprintf("can not find task: %s", req.TaskId))
	}
	val, ok := ch.(chan []byte)
	if !ok {
		return errors.New(fmt.Sprintf("invalid cache"))
	}
	for buf := range val {
		_ = server.Send(&api.WatchStream{
			Stream: buf,
		})
	}
	return nil
}

func (k Kobe) RunPlaybook(ctx context.Context, req *api.RunPlaybookRequest) (*api.RunPlaybookResult, error) {
	rm := RunnerManager{
		inventoryCache: k.inventoryCache,
	}
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
		Project:   req.Project,
	}
	k.taskCache.Set(result.Id, &result, cache.DefaultExpiration)
	k.chCache.Set(result.Id, ch, cache.DefaultExpiration)
	k.inventoryCache.Set(result.Id, req.Inventory, cache.DefaultExpiration)
	runner, _ := rm.CreatePlaybookRunner(req.Project, req.Playbook)
	go func() {
		runner.Run(ch, &result)
		result.Finished = true
		k.taskCache.Set(result.Id, &result, cache.DefaultExpiration)
	}()
	return &api.RunPlaybookResult{
		Result: &result,
	}, nil
}

func (k Kobe) GetResult(ctx context.Context, req *api.GetResultRequest) (*api.GetResultResponse, error) {
	id := req.GetTaskId()
	result, found := k.taskCache.Get(id)
	if !found {
		return nil, errors.New(fmt.Sprintf("can not find task: %s result", id))
	}
	val, ok := result.(*api.Result)
	if !ok {
		return nil, errors.New("invalid result type")
	}
	if val.Finished {
		bytes, err := ioutil.ReadFile(path.Join(constant.WorkDir, val.Project, val.Id, "result.json"))
		if err != nil {
			return nil, err
		}
		val.Content = string(bytes)
	}
	return &api.GetResultResponse{Item: val}, nil
}

func (k Kobe) ListResult(ctx context.Context, req *api.ListResultRequest) (*api.ListResultResponse, error) {
	var results []*api.Result
	resultMap := k.taskCache.Items()
	for taskId := range resultMap {
		item := resultMap[taskId].Object
		val, ok := item.(*api.Result)
		if !ok {
			continue
		}
		results = append(results, val)
	}
	return &api.ListResultResponse{
		Items: results,
	}, nil
}
