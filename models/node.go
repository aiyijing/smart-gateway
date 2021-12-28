package models

import (
	"fmt"
)

type Node struct {
	Name     string                 `json:"name"`
	OutBound map[string]interface{} `json:"outBound"`
	Enable   bool                   `json:"enable"`
}

func GetAllNode() ([]Node, error) {
	var nodes []Node
	err := DB.Read("v2ray", "nodes", &nodes)
	return nodes, err
}

func AddNode(node Node) error {
	var nodes []Node
	err := DB.Read("v2ray", "nodes", &nodes)
	if err != nil {
		return err
	}

	for _, n := range nodes {
		if n.Name == node.Name {
			return fmt.Errorf("node %v already exists", node.Name)
		}
	}

	nodes = append(nodes, node)
	return DB.Write("v2ray", "nodes", nodes)
}

func UpdateNode(node Node) error {
	var nodes []Node
	var updated = false
	err := DB.Read("v2ray", "nodes", &nodes)
	if err != nil {
		return err
	}

	for i, n := range nodes {
		if n.Name == node.Name {
			nodes[i] = node
			updated = true
		}
	}

	if !updated {
		return fmt.Errorf("node %v not exists", node.Name)
	}
	return DB.Write("v2ray", "nodes", &nodes)
}

func DeleteNode(name string) error {
	var nodes []Node
	var deleted = false
	err := DB.Read("v2ray", "nodes", &nodes)
	if err != nil {
		return err
	}

	for i, n := range nodes {
		if n.Name == name {
			nodes = append(nodes[:i], nodes[i+1:]...)
			deleted = true
			break
		}
	}
	if !deleted {
		return fmt.Errorf("node %v not exists", name)
	}

	return DB.Write("v2ray", "nodes", &nodes)
}
