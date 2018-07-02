// The gen tool automatically generates routing code and openapi3 documents
// Gen can also generate client code
//
// Basic usage:
//   1. Install gen tool `go get -u -v github.com/wzshiming/gen/cmd/gen`
//   2. Add gen tool to $PATH
//   3. Execute it `gen run github.com/wzshiming/gen/examples/route3`
//   4. Open http://127.0.0.1:8080/swagger/?url=./openapi.json# with your browser.

//go:generate gen route github.com/wzshiming/gen/examples/route3
//go:generate gen openapi github.com/wzshiming/gen/examples/route3

package route3

import (
	"strconv"
)

// Item is a item
type Item struct {
	ItemID  int    `json:"item_id"` // ItemID is the Item ID
	Name    string `json:"name"`    // Name is the name
	Message string `json:"message"` // Name is the message
}

// ItemServices #path:"/item/"#
type ItemServices struct{}

// Verify #security:"apiKey"#
func (*ItemServices) Verify(token string /* #in:"cookie"# */) (userID uint64, err error) {
	return strconv.ParseUint(token, 0, 0)
}

// Create #route:"PUT /"#
func (*ItemServices) Create(userID uint64 /* #in:"security"# */, uuid *string, item *Item) (err error) {
	return nil
}

// List #route:"GET /"#
func (*ItemServices) List(token string /* #in:"security"# */, uuid *string) (items []*Item, err error) {
	return nil, nil
}
