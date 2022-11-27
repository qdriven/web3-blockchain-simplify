package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func ReadFile(filepath string) ([]byte, error) {
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0666)
	if err != nil {
		return nil, fmt.Errorf("OpenFile %s error %s", filepath, err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(fmt.Sprintf("File %s close error %s", filepath, err))
		}
	}()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll %s error %s", filepath, err)
	}
	return data, nil
}

func FullPath(workspace, dir string, fileName string) string {
	if !path.IsAbs(workspace) {
		panic("path should be absolute path")
	}
	return path.Join(workspace, dir, fileName)
}
