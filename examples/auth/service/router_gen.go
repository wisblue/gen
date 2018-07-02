// Code generated; DO NOT EDIT.
// file router_gen.go

package service

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// ItemServices Define the method scope
var _globalItemServices ItemServices

// Router is generated do not edit.
func Router() http.Handler {
	router := mux.NewRouter()

	// Registered routing POST /item
	router.Path("/item").
		Methods("POST").
		HandlerFunc(_globalItemServices._operationUpdate)

	// Registered routing GET /item
	router.Path("/item").
		Methods("GET").
		HandlerFunc(_globalItemServices._operationList)

	// Registered routing PUT /item
	router.Path("/item").
		Methods("PUT").
		HandlerFunc(_globalItemServices._operationCreate)

	return router
}

// _requestQueryUuid Parsing the query for of uuid
func _requestQueryUuid(r *http.Request) (uuid *string, err error) {
	var _uuid = r.URL.Query().Get("uuid")
	uuid = &_uuid
	return
}

// _requestBodyItem Parsing the body for of item
func _requestBodyItem(r *http.Request) (item *Item, err error) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		r.Body.Close()
		json.Unmarshal(body, &item)
	}

	return
}

// _requestHeaderToken Parsing the header for of token
func _requestHeaderToken(r *http.Request) (token string, err error) {
	var _token = r.Header.Get("token")
	token = _token
	return
}

// _securityVerify Is the security of Verify
func _securityVerify(r *http.Request) (userID uint64, err error) {

	// Parsing token.
	_token, err := _requestHeaderToken(r)
	if err != nil {
		return
	}

	// Call Verify.
	return _globalItemServices.Verify(_token)

}

// _operationUpdate Is the route of Update
func (s ItemServices) _operationUpdate(w http.ResponseWriter, r *http.Request) {

	// Permission verification call Verify.
	_userID, err := _securityVerify(r)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	// Parsing uuid.
	_uuid, err := _requestQueryUuid(r)
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

	// Call Update.
	_err := s.Update(_userID, _uuid, _item)

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	w.WriteHeader(204)
	w.Write(nil)
	return
}

// _operationList Is the route of List
func (s ItemServices) _operationList(w http.ResponseWriter, r *http.Request) {

	// Permission verification undefined.
	var _token string
	// Parsing uuid.
	_uuid, err := _requestQueryUuid(r)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Call List.
	_items, _err := s.List(_token, _uuid)

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

	w.WriteHeader(204)
	w.Write(nil)
	return
}

// _operationCreate Is the route of Create
func (s ItemServices) _operationCreate(w http.ResponseWriter, r *http.Request) {

	// Permission verification call Verify.
	_userID, err := _securityVerify(r)
	if err != nil {
		http.Error(w, err.Error(), 403)
		return
	}

	// Parsing uuid.
	_uuid, err := _requestQueryUuid(r)
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

	// Call Create.
	_token, _err := s.Create(_userID, _uuid, _item)

	w.Header().Set("token", fmt.Sprint(_token))

	// Response code 400 Bad Request for err.
	if _err != nil {
		http.Error(w, _err.Error(), 400)
		return
	}

	w.WriteHeader(204)
	w.Write(nil)
	return
}
