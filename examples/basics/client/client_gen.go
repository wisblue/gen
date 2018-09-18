// Code generated; DO NOT EDIT.
// file ../client/client_gen.go

package main

import (
	"encoding/json"
	"fmt"
	"github.com/wzshiming/requests"
	"net/http"
)

// Item is a item
type Item struct {
	ItemID  int    // ItemID is the Item ID
	Name    string // Name is the name
	Message string // Name is the message
}

var Client = requests.NewClient().NewRequest()

// GetItem #route:"GET /item/{itemID}"#
func GetItem(_userID int, _itemID int) (_item *Item, _err error) {
	resp, err := Client.Clone().
		SetQuery("userID", fmt.Sprint(_userID)).
		SetPath("itemID", fmt.Sprint(_itemID)).
		get("/item/{itemID}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_item)
	case 400:
		_err = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return nil, err
	}

	return
}

// CreateItem #route:"PUT /item/"#
func CreateItem(_userID int, _item *Item) (_itemID int, _err error) {
	resp, err := Client.Clone().
		SetQuery("userID", fmt.Sprint(_userID)).
		SetJSON(_item).
		put("/item/")
	if err != nil {
		return 0, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_itemID)
	case 400:
		_err = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return 0, err
	}

	return
}
