package utilities

import (
	"encoding/json"
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
