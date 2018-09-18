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

// UserServices Define the method scope
var _globalUserServices UserServices

// ItemServices Define the method scope
var _globalItemServices ItemServices

// Router is generated do not edit.
func Router() http.Handler {
	router := mux.NewRouter()

	// Registered routing GET /user/{userID}
	router.Path("/user/{userID}").
		Methods("GET").
		HandlerFunc(_globalUserServices._operationGet)

	// Registered routing DELETE /user/{userID}
	router.Path("/user/{userID}").
		Methods("DELETE").
		HandlerFunc(_globalUserServices._operationDelete)

	// Registered routing PUT /user
	router.Path("/user").
		Methods("PUT").
		HandlerFunc(_globalUserServices._operationCreate)

	// Registered routing POST /item/{itemID}
	router.Path("/item/{itemID}").
		Methods("POST").
		HandlerFunc(_globalItemServices._operationUpdateItem)

	// Registered routing GET /item
	router.Path("/item").
		Methods("GET").
		HandlerFunc(_globalItemServices._operationListItem)

	// Registered routing GET /item/{itemID}
	router.Path("/item/{itemID}").
		Methods("GET").
		HandlerFunc(_globalItemServices._operationGetItem)

	// Registered routing DELETE /item/{itemID}
	router.Path("/item/{itemID}").
		Methods("DELETE").
		HandlerFunc(_globalItemServices._operationDeleteItem)

	// Registered routing PUT /item
	router.Path("/item").
		Methods("PUT").
		HandlerFunc(_globalItemServices._operationCreateItem)

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

// _requestBodyUser Parsing the body for of user
func _requestBodyUser(r *http.Request) (user *User, err error) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		r.Body.Close()
		json.Unmarshal(body, &user)
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

// _requestPathUserId Parsing the path for of userID
func _requestPathUserId(r *http.Request) (userID int, err error) {
	var _userID = mux.Vars(r)["userID"]
	if i, err := strconv.ParseInt(_userID, 0, 0); err == nil {
		userID = int(i)
	}

	return
}

// _operationGet Is the route of Get
func (s UserServices) _operationGet(w http.ResponseWriter, r *http.Request) {

	// Parsing userID.
	_userID, err := _requestPathUserId(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Call Get.
	_user, _err := s.Get(_userID)

	// Response code 200 OK for user.
	if _user != nil {
		data, err := json.Marshal(_user)
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

	data, err := json.Marshal(_user)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(data)

	return
}

// _operationDelete Is the route of Delete
func (s UserServices) _operationDelete(w http.ResponseWriter, r *http.Request) {

	// Parsing userID.
	_userID, err := _requestPathUserId(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Call Delete.
	_err := s.Delete(_userID)

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)

		return
	}

	w.WriteHeader(204)
	w.Write(nil)

	return
}

// _operationCreate Is the route of Create
func (s UserServices) _operationCreate(w http.ResponseWriter, r *http.Request) {

	// Parsing user.
	_user, err := _requestBodyUser(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Call Create.
	_userID, _err := s.Create(_user)

	// Response code 200 OK for userID.
	if _userID != 0 {
		data, err := json.Marshal(_userID)
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

	data, err := json.Marshal(_userID)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(data)

	return
}

// _operationUpdateItem Is the route of UpdateItem
func (s ItemServices) _operationUpdateItem(w http.ResponseWriter, r *http.Request) {

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

	// Parsing item.
	_item, err := _requestBodyItem(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Call UpdateItem.
	_err := s.UpdateItem(_userID, _itemID, _item)

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)

		return
	}

	w.WriteHeader(204)
	w.Write(nil)

	return
}

// _operationListItem Is the route of ListItem
func (s ItemServices) _operationListItem(w http.ResponseWriter, r *http.Request) {

	// Parsing userID.
	_userID, err := _requestQueryUserId(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Call ListItem.
	_items, _err := s.ListItem(_userID)

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

// _operationGetItem Is the route of GetItem
func (s ItemServices) _operationGetItem(w http.ResponseWriter, r *http.Request) {

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
	_item, _err := s.GetItem(_userID, _itemID)

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

// _operationDeleteItem Is the route of DeleteItem
func (s ItemServices) _operationDeleteItem(w http.ResponseWriter, r *http.Request) {

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

	// Call DeleteItem.
	_err := s.DeleteItem(_userID, _itemID)

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)

		return
	}

	w.WriteHeader(204)
	w.Write(nil)

	return
}

// _operationCreateItem Is the route of CreateItem
func (s ItemServices) _operationCreateItem(w http.ResponseWriter, r *http.Request) {

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
	_itemID, _err := s.CreateItem(_userID, _item)

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
