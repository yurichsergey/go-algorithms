package story

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Story map[string]Arc

type Option struct {
	Text string `json:"text,omitempty"`
	Arc  string `json:"arc,omitempty"`
}

type Arc struct {
	Title   string   `json:"title,omitempty"`
	Story   []string `json:"story,omitempty"`
	Options []Option `json:"options,omitempty"`
}

func FromJSON(r io.Reader) (Story, error) {
	var s Story
	if err := json.NewDecoder(r).Decode(&s); err != nil {
		return nil, err
	}
	return s, nil
}

func FromJSONUnmarshal(r io.Reader) (Story, error) {
	var s Story
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}
	return s, nil
}

func FromFilePath(filepath string) Story {
	f, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}(f)

	storyStruct, errJSON := FromJSON(f)
	if errJSON != nil {
		log.Fatal(errJSON)
	}
	return storyStruct
}
