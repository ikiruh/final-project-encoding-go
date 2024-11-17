package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}
	defer yamlFile.Close()
	var data models.DockerCompose
	f, _ := os.ReadFile(j.FileInput)
	err = json.Unmarshal(f, &data)
	if err != nil {
		return err
	}
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	_, err = yamlFile.Write(yamlData)
	if err != nil {
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}
	defer jsonFile.Close()
	var data models.DockerCompose
	f, _ := os.ReadFile(y.FileInput)
	err = yaml.Unmarshal(f, &data)
	if err != nil {
		return err
	}
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		return err
	}
	return nil
}
