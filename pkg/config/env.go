package config

import (
	"os"
	"sync"
	"bufio"
	"io"
	"strings"
	"path/filepath"
	"runtime"
)

type environment struct {
	arguments 	map[string]string
}

var (
	instance 	*environment
	once 		sync.Once
)

func getFilepath(filename string) string {
	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		panic("Cant determine the location of caller")
	}

	return filepath.Join(filepath.Dir(currentFile), filename)
}

func getInstance() *environment {
	once.Do(func() {
		// Find the current location of this package, which is also where .env is stored
		path := getFilepath(".env")

		file, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm) 
		if err != nil {
			file.Close()
			instance = &environment{}
			return
		}
		defer file.Close()		
		
		variables := make(map[string]string)

		rd := bufio.NewReader(file)
		for {
			line, err := rd.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				
				instance = &environment{}
				return
			}

			line = strings.TrimSpace(line)
			keyVal := strings.Split(line, "=")

			if len(keyVal) > 1 {
				variables[keyVal[0]] = string(keyVal[1])
			}	
		}

		instance = &environment{arguments: variables}
	})

	return instance
}

func Add(key, value string) {
	env := getInstance()
	env.arguments[key] = value
}

func Look(key string) (string, bool) {
	env := getInstance()

	value := env.arguments[key]
	if value == "" {
		return "", false
	}

	return value, true
}
