package util

import (
	"io/ioutil"
	"os"
	"path"
)

func ReadAll(p string) ([]byte, error) {
	f, err := os.OpenFile(p, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteOnCreate(p string, data []byte) error {
	parentFile := path.Dir(p)
	err := os.MkdirAll(parentFile, os.ModePerm)
	if err != nil {
		return err
	}
	f, err := os.OpenFile(p, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err != nil {
		return err
	}
	_ = f.Close()
	return nil
}
