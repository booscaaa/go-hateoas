package hateoas_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/booscaaa/go-hateoas/hateoas"
)

func TestHateoas(t *testing.T) {
	mockStruct := struct {
		ID      int
		Make    string
		Model   string
		Mileage int
		hateoas.Hateoas
	}{
		ID:      10,
		Make:    "Ford",
		Model:   "Taurus",
		Mileage: 200000,
	}

	hateoas.Generate(&mockStruct, "vehicle", "http://localhost:3000")

	out, err := json.Marshal(mockStruct)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
}
