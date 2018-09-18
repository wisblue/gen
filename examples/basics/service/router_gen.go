// Code generated; DO NOT EDIT.
// file router_gen.go

package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Router is generated do not edit.
func Router() http.Handler {
	router := mux.NewRouter()

	// Registered routing GET /item/{itemID}
	router.Path("/item/{itemID}").
		Methods("GET").
		HandlerFunc(_operationGetItem)

	// Registered routing PUT /item/
	router.Path("/item/").
		Methods("PUT").
		HandlerFunc(_operationCreateItem)

	return router
}

// _requestBodyItem Parsing the body for of item
func _requestBodyItem(r *http.Request) (item *Item, err error) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		r.Body.Close()
		json.Unmarshal(body, &item)
	}

	return
}

// _requestPathItemId Parsing the path for of itemID
func _requestPathItemId(r *http.Request) (itemID int, err error) {
	var _itemID = mux.Vars(r)["itemID"]
	if i, err := strconv.ParseInt(_itemID, 0, 0); err == nil {
		itemID = int(i)
	}

	return
}

// _requestQueryUserId Parsing the query for of userID
func _requestQueryUserId(r *http.Request) (userID int, err error) {
	var _userID = r.URL.Query().Get("userID")
	if i, err := strconv.ParseInt(_userID, 0, 0); err == nil {
		userID = int(i)
	}

	return
}

// _operationGetItem Is the route of GetItem
func _operationGetItem(w http.ResponseWriter, r *http.Request) {

	// Parsing userID.
	_userID, err := _requestQueryUserId(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Parsing itemID.
	_itemID, err := _requestPathItemId(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
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

// _operationCreateItem Is the route of CreateItem
func _operationCreateItem(w http.ResponseWriter, r *http.Request) {

	// Parsing userID.
	_userID, err := _requestQueryUserId(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Parsing item.
	_item, err := _requestBodyItem(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
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
