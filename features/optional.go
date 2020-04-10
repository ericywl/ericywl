package features

import (
	"encoding/json"
	"errors"
	"io"
)

type valid interface {
	OK() error
}

type person struct {
	Name string
}

func (p person) OK() error {
	if p.Name == "" {
		return errors.New("name required")
	}

	return nil
}

func decode(r io.Reader, v interface{}) error {
	err := json.NewDecoder(r).Decode(v)
	if err != nil {
		return err
	}

	obj, ok := v.(valid)
	if !ok {
		return nil
	}

	err = obj.OK()
	if err != nil {
		return err
	}

	return nil
}
