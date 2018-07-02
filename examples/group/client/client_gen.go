// Code generated; DO NOT EDIT.
// file ../client/client_gen.go

package client

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

// ItemServices #path:"/item/"#
type ItemServices struct{}

// User is a user
type User struct {
	UserID int    // UserID is the User ID
	Name   string // Name is the name
}

// UserServices #path:"/user/"#
type UserServices struct{}

var Client = requests.NewClient().NewRequest()

// GetUser #route:"GET /{userID}"#
func (UserServices) Get(_userID int) (_user *User, _err error) {
	resp, err := Client.Clone().
		SetPath("userID", fmt.Sprint(_userID)).
		Get("/user/{userID}")
	if err != nil {
		return nil, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_user)
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

// DeleteUser #route:"DELETE /{userID}"#
func (UserServices) Delete(_userID int) (_err error) {
	resp, err := Client.Clone().
		SetPath("userID", fmt.Sprint(_userID)).
		Delete("/user/{userID}")
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

// CreateUser #route:"PUT /"#
func (UserServices) Create(_user *User) (_userID int, _err error) {
	resp, err := Client.Clone().
		SetJSON(_user).
		Put("/user")
	if err != nil {
		return 0, err
	}

	switch code := resp.StatusCode(); code {
	case 200:
		err = json.Unmarshal(resp.Body(), &_userID)
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

// UpdateItem #route:"POST /{itemID}"#
func (ItemServices) UpdateItem(_userID int, _itemID int, _item *Item) (_err error) {
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

// ListItem #route:"GET /"#
func (ItemServices) ListItem(_userID int) (_items []*Item, _err error) {
	resp, err := Client.Clone().
		SetQuery("userID", fmt.Sprint(_userID)).
		Get("/item")
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

// GetItem #route:"GET /{itemID}"#
func (ItemServices) GetItem(_userID int, _itemID int) (_item *Item, _err error) {
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

// DeleteItem #route:"DELETE /{itemID}"#
func (ItemServices) DeleteItem(_userID int, _itemID int) (_err error) {
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

// CreateItem #route:"PUT /"#
func (ItemServices) CreateItem(_userID int, _item *Item) (_itemID int, _err error) {
	resp, err := Client.Clone().
		SetQuery("userID", fmt.Sprint(_userID)).
		SetJSON(_item).
		Put("/item")
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
