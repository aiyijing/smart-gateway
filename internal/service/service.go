package service

import (
	"fmt"
	"github.com/aiyijing/smart-gateway/internal/util"
	"github.com/docker/docker/client"
	"path"
)

const (
	// Service Status
	Running = "Running"
	Stop    = "Stop"
	Unknown = "Unknown"
)

type Service interface {
	Stop() error
	Start() error
	Restart() error
	Describe() *Desc
}

type Desc struct {
	Name   string
	Volume string
	Config string
}

func (d *Desc) Equal(desc *Desc) bool {
	return d.Name == desc.Name &&
		d.Config == desc.Config &&
		d.Volume == desc.Volume
}

type Manager struct {
	service map[string]Service
	docker  *client.Client
}

func NewManager() *Manager {
	return &Manager{service: map[string]Service{}}
}
func (m *Manager) AddService(name string, svc Service) error {
	des := svc.Describe()
	for _, v := range m.service {
		if des.Equal(v.Describe()) {
			return fmt.Errorf("service %s has exists", des.Name)
		}
	}
	m.service[name] = svc
	return nil
}

func (m *Manager) StartService(name string) error {
	svc, ok := m.service[name]
	if !ok {
		return fmt.Errorf("service %s not exists", name)
	}
	return svc.Start()
}

func (m *Manager) StopService(name string) error {
	svc, ok := m.service[name]
	if !ok {
		return fmt.Errorf("service %s not exists", name)
	}
	return svc.Stop()
}

func (m *Manager) RestartService(name string) error {
	svc, ok := m.service[name]
	if !ok {
		return fmt.Errorf("service %s not exists", name)
	}
	return svc.Restart()
}

func (m *Manager) UpdateConfig(name string, data string) error {
	svc, ok := m.service[name]
	if !ok {
		return fmt.Errorf("service %s not exists", name)
	}
	return m.updateConfig(data, svc.Describe())
}
func (m *Manager) GetCurrentConfig(name string) (string, error) {
	svc, ok := m.service[name]
	if !ok {
		return "", fmt.Errorf("service %s not exists", name)
	}
	return m.getCurrentConfig(svc.Describe())
}

func (m *Manager) getCurrentConfig(desc *Desc) (string, error) {
	mount := "/var/lib/docker/volumes/"
	p := path.Join(mount, desc.Volume, "_data")
	filePath := path.Join(p, desc.Config)

	data, err := util.ReadAll(filePath)
	return string(data), err
}

func (m *Manager) updateConfig(data string, desc *Desc) error {
	mount := "/var/lib/docker/volumes/"
	p := path.Join(mount, desc.Volume, "_data")
	filePath := path.Join(p, desc.Config)
	return util.WriteOnCreate(filePath, []byte(data))
}
