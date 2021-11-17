package maps

import "errors"

type Dictionary map[string]string

func (d Dictionary) search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", errors.New("key not exist")
	}
	return value, nil
}

func (d Dictionary) add(key, value string) {
	d[key] = value
}