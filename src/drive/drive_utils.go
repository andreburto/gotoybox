package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
)

const mapFileName string = "file_map.json"

func GetDataPath() string {
	var directory string = "data"

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", path, directory)
}

func GenerateUUID() string {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return id.String()
}

func LoadJsonToMap(o *os.File) map[string]string {
	var fileMap map[string]string
	decoder := json.NewDecoder(o)
	err := decoder.Decode(&fileMap)
	if err != nil {
		panic(err)
	}
	return fileMap
}

func GetIdFromFilePath(filePath string) string {
	var fileMap map[string]string = LoadFileMap()
	return fileMap[filePath]
}

func LoadFileMap() map[string]string {
	var path string = GetDataPath()

	cin, err := os.Open(fmt.Sprintf("%s/%s", path, mapFileName))
	if err != nil {
		panic(err)
	}
	defer cin.Close()

	return LoadJsonToMap(cin)
}

func LoadFileFromDisk(filePath string) string {
	var path string = GetDataPath()

	cin, err := os.Open(fmt.Sprintf("%s/%s", path, GetIdFromFilePath(filePath)))
	if err != nil {
		panic(err)
	}
	defer cin.Close()

	content, err := io.ReadAll(cin)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func AddFileToMap(id string, filePath string) {
	var path string = GetDataPath()

	cin, err := os.OpenFile(fmt.Sprintf("%s/%s", path, mapFileName), os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer cin.Close()

	var fileMap map[string]string = LoadJsonToMap(cin)
	fileMap[filePath] = id

	cin.Seek(0, 0) // Reset the file pointer to the beginning
	encoder := json.NewEncoder(cin)
	err = encoder.Encode(fileMap)
	if err != nil {
		panic(err)
	}
}

func SaveFileToDisk(fileContent string, filePath string) {
	var id string = GenerateUUID()
	var dataPath string = GetDataPath()

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		err = os.MkdirAll(dataPath, 0755)
		if err != nil {
			panic(err)
		}
	}

	destPath := fmt.Sprintf("%s/%s", dataPath, id)
	err := os.WriteFile(destPath, []byte(fileContent), 0644)
	if err != nil {
		panic(err)
	}
	fmt.Printf("File saved to %s\n", dataPath)
	fmt.Printf("filePath: %s\n", filePath)

	AddFileToMap(id, filePath)
}