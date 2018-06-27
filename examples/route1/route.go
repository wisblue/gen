// Code generated; DO NOT EDIT.

//
package route1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	// Registered routing POST /item/{itemID}
	router.Path("/item/{itemID}").
		Methods("POST").
		HandlerFunc(_UpdateItem)

	// Registered routing GET /item/
	router.Path("/item/").
		Methods("GET").
		HandlerFunc(_ListItem)

	// Registered routing GET /item/{itemID}
	router.Path("/item/{itemID}").
		Methods("GET").
		HandlerFunc(_GetItem)

	// Registered routing DELETE /item/{itemID}
	router.Path("/item/{itemID}").
		Methods("DELETE").
		HandlerFunc(_DeleteItem)

	// Registered routing PUT /item/
	router.Path("/item/").
		Methods("PUT").
		HandlerFunc(_CreateItem)

	return router
}

// _UpdateItem Is the route of UpdateItem
func _UpdateItem(w http.ResponseWriter, r *http.Request) {

	// Parsing the query for userID.
	var _userID int
	if i, err := strconv.ParseInt(r.URL.Query().Get("userID"), 0, 0); err == nil {
		_userID = int(i)
	}

	// Parsing the path for itemID.
	var _itemID int
	if i, err := strconv.ParseInt(mux.Vars(r)["itemID"], 0, 0); err == nil {
		_itemID = int(i)
	}

	// Parsing the body for item.
	var _item *Item
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		r.Body.Close()
		json.Unmarshal(body, &_item)
	}

	// Call UpdateItem.
	_err := UpdateItem(_userID, _itemID, _item)

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	w.Write(nil)
	return
}

// _ListItem Is the route of ListItem
func _ListItem(w http.ResponseWriter, r *http.Request) {

	// Parsing the query for userID.
	var _userID int
	if i, err := strconv.ParseInt(r.URL.Query().Get("userID"), 0, 0); err == nil {
		_userID = int(i)
	}

	// Call ListItem.
	_items, _err := ListItem(_userID)

	// Response code 200 OK for items.
	if _items != nil {
		data, err := json.Marshal(_items)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(data)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	w.Write(nil)
	return
}

// _GetItem Is the route of GetItem
func _GetItem(w http.ResponseWriter, r *http.Request) {

	// Parsing the query for userID.
	var _userID int
	if i, err := strconv.ParseInt(r.URL.Query().Get("userID"), 0, 0); err == nil {
		_userID = int(i)
	}

	// Parsing the path for itemID.
	var _itemID int
	if i, err := strconv.ParseInt(mux.Vars(r)["itemID"], 0, 0); err == nil {
		_itemID = int(i)
	}

	// Call GetItem.
	_item, _err := GetItem(_userID, _itemID)

	// Response code 200 OK for item.
	if _item != nil {
		data, err := json.Marshal(_item)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(data)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	w.Write(nil)
	return
}

// _DeleteItem Is the route of DeleteItem
func _DeleteItem(w http.ResponseWriter, r *http.Request) {

	// Parsing the query for userID.
	var _userID int
	if i, err := strconv.ParseInt(r.URL.Query().Get("userID"), 0, 0); err == nil {
		_userID = int(i)
	}

	// Parsing the path for itemID.
	var _itemID int
	if i, err := strconv.ParseInt(mux.Vars(r)["itemID"], 0, 0); err == nil {
		_itemID = int(i)
	}

	// Call DeleteItem.
	_err := DeleteItem(_userID, _itemID)

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	w.Write(nil)
	return
}

// _CreateItem Is the route of CreateItem
func _CreateItem(w http.ResponseWriter, r *http.Request) {

	// Parsing the query for userID.
	var _userID int
	if i, err := strconv.ParseInt(r.URL.Query().Get("userID"), 0, 0); err == nil {
		_userID = int(i)
	}

	// Parsing the body for item.
	var _item *Item
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		r.Body.Close()
		json.Unmarshal(body, &_item)
	}

	// Call CreateItem.
	_itemID, _err := CreateItem(_userID, _item)

	// Response code 200 OK for itemID.
	if _itemID != 0 {
		data, err := json.Marshal(_itemID)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Write(data)
		return
	}

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	w.Write(nil)
	return
}
