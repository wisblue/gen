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
	ItemID  int    `json:"item_id"` // ItemID is the Item ID
	Name    string `json:"name"`    // Name is the name
	Message string `json:"message"` // Name is the message
}

// ItemServices #path:"/item/"#
type ItemServices struct{}

// Verify #security:"apiKey"#
func (ItemServices) Verify(_token string /* #in:"header"# */) {
	Client = Client.
		SetHeader("token", fmt.Sprint(_token))

}

var Client = requests.NewClient().NewRequest()

// Create #route:"POST /"#
func (ItemServices) Update(_uuid *string, _item *Item) (_err error) {
	resp, err := Client.Clone().
		SetQuery("uuid", fmt.Sprint(_uuid)).
		SetJSON(_item).
		Post("/item")
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

// List #route:"GET /"#
func (ItemServices) List(_uuid *string) (_items []*Item, _err error) {
	resp, err := Client.Clone().
		SetQuery("uuid", fmt.Sprint(_uuid)).
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

// Create #route:"PUT /"#
func (ItemServices) Create(_uuid *string, _item *Item) (_token string /* #in:"header"# */, _err error) {
	resp, err := Client.Clone().
		SetQuery("uuid", fmt.Sprint(_uuid)).
		SetJSON(_item).
		Put("/item")
	if err != nil {
		return "", err
	}

	switch code := resp.StatusCode(); code {
	case 400:
		_err = fmt.Errorf(string(resp.Body()))
	default:
		err = fmt.Errorf("Undefined code %d %s", code, http.StatusText(code))
	}

	if err != nil {
		return "", err
	}

	return
}
