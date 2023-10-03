package book

import (
	"errors"
	"net/http"
)

type Request struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	ISBN  string `json:"isbn"`
}

type Response struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Genre string `json:"genre"`
	ISBN  string `json:"isbn"`
}

func (s *Request) Bind(r *http.Request) error {
	if s.Name == "" {
		return errors.New("name: cannot be blank")
	}
	if s.Genre == "" {
		return errors.New("name: cannot be blank")
	}
	if s.ISBN == "" {
		return errors.New("name: cannot be blank")
	}
	return nil
}

func ParseFromEntity(data Entity) (res Response) {
	res = Response{
		ID:    data.ID,
		Name:  *data.Name,
		Genre: *data.Genre,
		ISBN:  *data.ISBN,
	}
	return
}

func ParseFromEntities(data []Entity) (res []Response) {
	res = make([]Response, 0)
	for _, i := range data {
		res = append(res, ParseFromEntity(i))
	}
	return
}
