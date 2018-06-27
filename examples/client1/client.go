// Code generated; DO NOT EDIT.

//
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

// UpdateItem #route:"POST /item/{itemID}"#
func UpdateItem(_userID int, _itemID int, _item *Item) (_err error) {
	resp, err := Client.Clone().
		SetQuery("userID", fmt.Sprint(_userID)).
		SetPath("itemID", fmt.Sprint(_itemID)).
		SetJSON(_item).
		Post("/item/{itemID}")

	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_err = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// ListItem #route:"GET /item/"#
func ListItem(_userID int) (_items []*Item, _err error) {
	resp, err := Client.Clone().
		SetQuery("userID", fmt.Sprint(_userID)).
		Get("/item/")

	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_items)
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

// GetItem #route:"GET /item/{itemID}"#
func GetItem(_userID int, _itemID int) (_item *Item, _err error) {
	resp, err := Client.Clone().
		SetQuery("userID", fmt.Sprint(_userID)).
		SetPath("itemID", fmt.Sprint(_itemID)).
		Get("/item/{itemID}")

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

// DeleteItem #route:"DELETE /item/{itemID}"#
func DeleteItem(_userID int, _itemID int) (_err error) {
	resp, err := Client.Clone().
		SetQuery("userID", fmt.Sprint(_userID)).
		SetPath("itemID", fmt.Sprint(_itemID)).
		Delete("/item/{itemID}")

	if err != nil {
		return err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_err = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return err
	}

	return
}

// CreateItem #route:"PUT /item/"#
func CreateItem(_userID int, _item *Item) (_itemID int, _err error) {
	resp, err := Client.Clone().
		SetQuery("userID", fmt.Sprint(_userID)).
		SetJSON(_item).
		Put("/item/")

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
