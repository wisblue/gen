// The gen tool automatically generates routing code and openapi3 documents
// Gen can also generate client code
// 
// Basic usage:
//   1. Install gen tool `go get -u -v github.com/wzshiming/gen/cmd/gen`
//   2. Add gen tool to $PATH
//   3. Execute it `gen run -p github.com/wzshiming/gen/examples/route1`
//   4. Open http://127.0.0.1:8080/swagger/?url=./openapi.json# with your browser.

//go:generate gen route -p github.com/wzshiming/gen/examples/route1 -o route.go

package route1

import (
	"fmt"
)

// Item is a item
type Item struct {
	ItemID  int    // ItemID is the Item ID
	Name    string // Name is the name
	Message string // Name is the message
}

var size int
var mapitems = map[int][]*Item{}

// CreateItem #route:"PUT /item/"#
func CreateItem(userID int, item *Item) (itemID int, err error) {
	size++
	item.ItemID = size
	mapitems[userID] = append(mapitems[userID], item)
	return size, nil
}

// ListItem #route:"GET /item/"#
func ListItem(userID int) (items []*Item, err error) {
	return mapitems[userID], nil
}

// DeleteItem #route:"DELETE /item/{itemID}"#
func DeleteItem(userID int, itemID int) (err error) {
	for i, v := range mapitems[userID] {
		if v.ItemID == itemID {
			mapitems[userID] = append(mapitems[userID][:i], mapitems[userID][i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("It doesn't exist")
}

// GetItem #route:"GET /item/{itemID}"#
func GetItem(userID int, itemID int) (item *Item, err error) {
	for i, v := range mapitems[userID] {
		if v.ItemID == itemID {
			return mapitems[userID][i], nil
		}
	}
	return nil, fmt.Errorf("It doesn't exist")
}

// UpdateItem #route:"POST /item/{itemID}"#
func UpdateItem(userID int, itemID int, item *Item) (err error) {
	for i, v := range mapitems[userID] {
		if v.ItemID == itemID {
			item.ItemID = itemID
			*mapitems[userID][i] = *item
			return nil
		}
	}
	return fmt.Errorf("It doesn't exist")
}
