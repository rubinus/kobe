package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"kobe/api"
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

func (c *KobeClient) CreateProject(name string, source string) (*api.Project, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := api.CreateProjectRequest{
		Name:   name,
		Source: source,
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

func (c KobeClient) RunPlaybook(project, playbook string, inventory api.Inventory) (*api.Result, error) {
	conn, err := c.createConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := &api.RunPlaybookRequest{
		Project:   project,
		Playbook:  playbook,
		Inventory: &inventory,
	}
	req, err := client.RunPlaybook(context.Background(), request)
	if err != nil {
		return nil, err
	}
	return req.Result, nil
}

func (c *KobeClient) WatchRunPlaybook(taskId string, writer io.Writer) error {
	conn, err := c.createConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	req := &api.WatchPlaybookRequest{
		TaskId: taskId,
	}
	server, err := client.WatchRunPlaybook(context.Background(), req)
	if err != nil {
		return err
	}
	for {
		msg, err := server.Recv()
		if err != nil {
			return err
		}
		_, err = writer.Write(msg.Stream)
		if err != nil || err == io.EOF {
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
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
