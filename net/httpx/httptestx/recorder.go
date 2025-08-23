package httptestx

import (
	"encoding/json"
	"net/http/httptest"
)

type JSONResponseRecorder[T any] struct {
	*httptest.ResponseRecorder
}

func NewJSONResponseRecorder[T any]() *JSONResponseRecorder[T] {
	return &JSONResponseRecorder[T]{
		ResponseRecorder: httptest.NewRecorder(),
	}
}

func (r JSONResponseRecorder[T]) Scan() (T, error) {
	var t T
	err := json.NewDecoder(r.Body).Decode(&t)
	return t, err
}

func (r JSONResponseRecorder[T]) MustScan() T {
	t, err := r.Scan()
	if err != nil {
		panic(err)
	}
	return t
}
