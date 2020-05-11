package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"kobe/api"
)

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

func (c KobeClient) RunPlaybook(project, playbook string, inventory api.Inventory, writer io.Writer, result *api.Result) error {
	conn, err := c.createConnection()
	if err != nil {
		return err
	}
	defer conn.Close()
	client := api.NewKobeApiClient(conn)
	request := &api.RunPlaybookRequest{
		Project:   project,
		Playbook:  playbook,
		Inventory: &inventory,
	}
	server, err := client.RunPlaybook(context.Background(), request)
	if err != nil {
		return err
	}
	for {
		msg, err := server.Recv()
		if err != nil {
			return err
		}
		result = msg.Result
		_, err = writer.Write(msg.Stream)
		if err != nil {
			return err
		}
		if result.Finished {
			break
		}
	}
	return nil
}

func (c *KobeClient) createConnection() (*grpc.ClientConn, error) {
	address := fmt.Sprintf("%s:%d", c.host, c.port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return conn, nil
}
