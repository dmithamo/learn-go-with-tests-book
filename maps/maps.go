package maps

type Dict struct {
	Words map[string]string
}

type NotFoundErr string

var NOT_FOUND_ERR NotFoundErr

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
