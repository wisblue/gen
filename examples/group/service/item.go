
package service

import (
	"fmt"
)

// Item is a item
type Item struct {
	ItemID  int    // ItemID is the Item ID
	Name    string // Name is the name
	Message string // Name is the message
}

var sizeitems int
var mapitems = map[int][]*Item{}

// ItemServices #path:"/item/"#
type ItemServices struct{}

// CreateItem #route:"PUT /"#
func (*ItemServices) CreateItem(userID int, item *Item) (itemID int, err error) {
	sizeitems++
	item.ItemID = sizeitems
	mapitems[userID] = append(mapitems[userID], item)
	return sizeitems, nil
}

// ListItem #route:"GET /"#
func (ItemServices) ListItem(userID int) (items []*Item, err error) {
	return mapitems[userID], nil
}

// DeleteItem #route:"DELETE /{itemID}"#
func (ItemServices) DeleteItem(userID int, itemID int) (err error) {
	for i, v := range mapitems[userID] {
		if v.ItemID == itemID {
			mapitems[userID] = append(mapitems[userID][:i], mapitems[userID][i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("It doesn't exist")
}

// GetItem #route:"GET /{itemID}"#
func (ItemServices) GetItem(userID int, itemID int) (item *Item, err error) {
	for i, v := range mapitems[userID] {
		if v.ItemID == itemID {
			return mapitems[userID][i], nil
		}
	}
	return nil, fmt.Errorf("It doesn't exist")
}

// UpdateItem #route:"POST /{itemID}"#
func (ItemServices) UpdateItem(userID int, itemID int, item *Item) (err error) {
	for i, v := range mapitems[userID] {
		if v.ItemID == itemID {
			item.ItemID = itemID
			*mapitems[userID][i] = *item
			return nil
		}
	}
	return fmt.Errorf("It doesn't exist")
}
