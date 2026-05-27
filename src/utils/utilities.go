package utilities

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"

	"go-vrf/src/model"
)

func GenerateUUIDs(quantidade int) []string {
	uuids := make([]string, quantidade)
	for i := range uuids {
		uuids[i] = uuid.New().String()
	}
	return uuids
}

type JsonData = map[string]model.EdgeCluster

func ReadT0Json(name string) (JsonData, error) {
	if err := ValidateFileName(name); err != nil {
		return JsonData{}, err
	}

	file, err := os.Open(name + ".json")
	if err != nil {
		return JsonData{}, err
	}
	defer func() { _ = file.Close() }()

	var data JsonData
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return JsonData{}, err
	}

	return data, nil
}

func SaveToFile(name string, result any) error {
	if err := ValidateFileName(name); err != nil {
		return err
	}
	filename := fmt.Sprintf("%s.json", name)

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o600)
	if err != nil {
		return fmt.Errorf("error opening file: %w", err)
	}
	defer func() { _ = f.Close() }()

	tierName := map[string]any{name: result}

	jsonData, err := json.Marshal(tierName)
	if err != nil {
		return fmt.Errorf("error marshalling JSON data: %w", err)
	}

	if _, err := f.Write(jsonData); err != nil {
		return fmt.Errorf("error writing JSON data to file: %w", err)
	}

	fmt.Println("JSON data saved to:", filename)
	return nil
}

type JsonT1Data = map[string]model.Organizations

func ReadT1Json(name string) (JsonT1Data, error) {
	if err := ValidateFileName(name); err != nil {
		return JsonT1Data{}, err
	}

	file, err := os.Open(name + ".json")
	if err != nil {
		return JsonT1Data{}, err
	}
	defer func() { _ = file.Close() }()

	var data JsonT1Data
	if err := json.NewDecoder(file).Decode(&data); err != nil {
		return JsonT1Data{}, err
	}

	return data, nil
}
