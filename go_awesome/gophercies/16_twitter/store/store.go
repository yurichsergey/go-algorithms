package store

func Load(path string) ([]string, error)
func Save(path string, users []string) error
