package store

import (
	"encoding/json"
	"os"
)

func Load(path string) ([]string, error) {
	buf, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return []string{}, nil
	}
	if err != nil {
		return nil, err
	}

	var data []string
	err = json.Unmarshal(buf, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func Save(path string, users []string) error {
	buf, err := json.Marshal(users)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, buf, 0644)
	return err
}
