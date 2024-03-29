package utilities

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/google/uuid"
	"go-vrf/src/model"
)

func GenerateUUIDs(quantidade int) []string {
	// Cria um slice para armazenar os UUIDs
	uuids := make([]string, quantidade)

	// Gera um novo UUID para cada elemento do slice
	for i := range uuids {
		uuids[i] = uuid.New().String()
	}

	// Retorna o slice de UUIDs
	return uuids
}

type JsonData = map[string]model.EdgeCluster

func ReadT0Json(filePath string) (JsonData, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return JsonData{}, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	var data JsonData
	reader := io.Reader(file)
	err = json.NewDecoder(reader).Decode(&data)
	if err != nil {
		return JsonData{}, err
	}

	return data, nil
}

func SaveToFile(name string, result interface{}) error {
	filename := fmt.Sprintf("%s.json", name)

	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// File doesn't exist, create it with appropriate permissions
		f, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("error creating file:", err)
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {

			}
		}(f)
	}

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	var tierName = make(map[string]interface{})
	tierName[name] = result

	jsonData, err := json.Marshal(tierName)
	if err != nil {
		return fmt.Errorf("error marshalling JSON data: %v", err)
	}

	_, err = f.Write(jsonData)
	if err != nil {
		return fmt.Errorf("error writing JSON data to file: %v", err)
	}

	fmt.Println("JSON data saved to:", filename)
	return nil
}
