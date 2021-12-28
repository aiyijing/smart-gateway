package util

import (
	"context"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func StopContainer(ctx context.Context, name string, cli *client.Client) error {
	duration := time.Second * 5
	err := cli.ContainerStop(ctx, name, &duration)
	if client.IsErrNotFound(err) {
		return nil
	}
	return err
}

func RemoveContainer(ctx context.Context, name string, cli *client.Client) error {
	err := cli.ContainerRemove(ctx, name, types.ContainerRemoveOptions{})
	if client.IsErrNotFound(err) {
		return nil
	}
	return err
}
