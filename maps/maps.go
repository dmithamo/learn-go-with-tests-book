package maps

import "errors"

type Dict struct {
	Words map[string]string
}

var NotFoundErr = errors.New("word not found")
var WordExistsErr = errors.New("word already exists")

func (d *Dict) Search(key string) (string, error) {
	def, ok := d.Words[key]
	if !ok {
		return "", NotFoundErr
	}

	return def, nil
}

func (d *Dict) Add(word, defn string) error {
	_, err := d.Search(word)
	if err == nil {
		return WordExistsErr
	}

	d.Words[word] = defn
	return nil
}
