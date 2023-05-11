package client

import (
	"context"
	"fmt"
	"github.com/KubeOperator/kobe/api"
	"google.golang.org/grpc"
	"io"
)

func NewKobeClient(host string, port int) *KobeClient {
	return &KobeClient{
		host: host,
		port: port,
	}
}

type KobeClient struct {
	host string
	port int
}

func (c *KobeClient) CreateProject(name string, source string, inventByte []byte) (*api.Project, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := api.CreateProjectRequest{
		Name:      name,
		Source:    source,
		Inventory: inventByte,
	}
	resp, err := client.CreateProject(context.Background(), &request)
	if err != nil {
		return nil, err
	}
	return resp.Item, nil

}

func (c KobeClient) ListProject() ([]*api.Project, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := api.ListProjectRequest{}
	resp, err := client.ListProject(context.Background(), &request)
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

func (c KobeClient) RunPlaybook(project, playbook, tag string, inventory *api.Inventory) (*api.Result, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := &api.RunPlaybookRequest{
		Project:   project,
		Playbook:  playbook,
		Inventory: inventory,
		Tag:       tag,
	}
	req, err := client.RunPlaybook(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return req.Result, nil
}

func (c KobeClient) RunAdhoc(pattern, module, param string, inventory *api.Inventory) (*api.Result, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := &api.RunAdhocRequest{
		Inventory: inventory,
		Module:    module,
		Param:     param,
		Pattern:   pattern,
	}
	req, err := client.RunAdhoc(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return req.Result, nil
}

func (c *KobeClient) WatchRun(taskId string, writer io.Writer) error {
	conn, err := c.createConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	req := &api.WatchRequest{
		TaskId: taskId,
	}
	server, err := client.WatchResult(context.Background(), req)
	if err != nil {
		return err
	}
	for {
		msg, err := server.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		_, err = writer.Write(msg.Stream)
		if err != nil {
			break
		}
	}
	return nil
}

func (c *KobeClient) GetResult(taskId string) (*api.Result, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := api.GetResultRequest{
		TaskId: taskId,
	}
	resp, err := client.GetResult(context.Background(), &request)
	if err != nil {
		return nil, err
	}
	return resp.Item, nil
}

func (c *KobeClient) ListResult() ([]*api.Result, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := api.ListResultRequest{}
	resp, err := client.ListResult(context.Background(), &request)
	if err != nil {
		return nil, err
	}
	return resp.Items, nil
}

func (c *KobeClient) createConnection() (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", c.host, c.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(100*1024*1024)))
	if err != nil {
		return nil, err
	}
	return conn, nil
}
