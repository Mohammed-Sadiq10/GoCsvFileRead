package json

import (
	"encoding/json"
	"os"
)

func JsonFileOutput(jsonFilePath string, emp interface{}) error {
	file, err := os.Create(jsonFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(emp); err != nil {
		return err
	}

	return nil
}
