package maps

import "errors"

type Dict struct {
	Words map[string]string
}

type NotFoundErr string

var NOT_FOUND_ERR NotFoundErr
var WORD_EXISTS_ERR = errors.New("word already exists")

func (e NotFoundErr) Error() string {
	return "word not found"
}

func (d *Dict) Search(key string) (string, error) {
	def, ok := d.Words[key]
	if !ok {
		return "", NOT_FOUND_ERR
	}

	return def, nil
}

func (d *Dict) Add(word, defn string) error {
	_,err := d.Search(word)
	if err == nil {
		return WORD_EXISTS_ERR
	}

	d.Words[word] = defn
	return nil
}
