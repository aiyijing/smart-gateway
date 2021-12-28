package service

import (
	"context"
	"github.com/aiyijing/smart-gateway/internal/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type V2ray struct {
	name       string
	docker     *client.Client
	volumeName string
	configName string
}

func NewV2ray(name string, docker *client.Client) *V2ray {
	return &V2ray{
		name:       name,
		docker:     docker,
		volumeName: name,
		configName: "config.json",
	}
}

func (v *V2ray) Start() error {
	ctx := context.Background()
	if err := v.ensureVolumeExist(); err != nil {
		return err
	}

	config := &container.Config{
		Image: "v2ray/official",
	}
	hostConfig := &container.HostConfig{
		NetworkMode: "host",
		Binds:       []string{"v2ray:/etc/v2ray"},
	}

	c, err := v.docker.ContainerCreate(
		ctx,
		config,
		hostConfig,
		nil,
		nil,
		v.name,
	)

	err = v.docker.ContainerStart(ctx,
		c.ID,
		types.ContainerStartOptions{},
	)
	if err != nil {
		return err
	}
	return nil
}

func (v *V2ray) Stop() error {
	ctx := context.Background()
	err := util.StopContainer(ctx, v.name, v.docker)
	if err != nil {
		return err
	}
	return util.RemoveContainer(ctx, v.name, v.docker)
}

func (v *V2ray) Restart() error {
	_ = v.Stop()
	return v.Start()
}

func (v *V2ray) Describe() *Desc {
	return &Desc{
		Name:   v.name,
		Volume: v.volumeName,
		Config: v.configName,
	}
}

func (v *V2ray) ensureVolumeExist() error {
	ctx := context.Background()
	_, err := v.docker.VolumeInspect(ctx, v.volumeName)
	if !client.IsErrNotFound(err) {
		return err
	}
	_, err = v.docker.VolumeCreate(ctx, volume.VolumeCreateBody{
		Name: v.volumeName,
	})
	return err
}
