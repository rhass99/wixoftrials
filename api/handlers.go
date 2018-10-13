package api

import (
	"encoding/json"
	"fmt"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/user"
	"io/ioutil"
	"net/http"
)

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	ctx := appengine.NewContext(r)
	u := user.Current(ctx)
	if u == nil {
		url, _ := user.LoginURL(ctx, "/")
		fmt.Fprintf(w, `<a href="%s">Sign in or register</a>`, url)
		return
	}
	url, _ := user.LogoutURL(ctx, "/")
	fmt.Fprintf(w, `Welcome, %s! (<a href="%s">sign out</a>)`, u, url)
}

func HandleRetrieveAll(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	var results []Account

	q := datastore.NewQuery("Account").Filter("Email =", " ")

	_, err := q.GetAll(ctx, &results)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}

	js, err := json.Marshal(results);
	if err != nil {
		log.Errorf(ctx, "%v", err)

	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func HandleAccountSignup(w http.ResponseWriter, r *http.Request) {
	// to create a new Account, I have to:
	// 1 - Create an empty account object
	ctx := appengine.NewContext(r)
	var account Account

	// 2 - Parse incoming JSON to the account
	if incomingJSON, err := ioutil.ReadAll(r.Body); err != nil {
		// Internal server error
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"userExist": true}`))
		return
	} else {
		json.Unmarshal([]byte(incomingJSON), &account)
	}

	// 3 - Check if user is logged in (yes => move on / no => prompt to signin)

	// 4 - Check if the Account exists in db (yes => return it/no => move on)
	keyExists, err := DBCheckEntity(&account, ctx);
	if err != nil {
		// Internal server error
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		if keyExists == true {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"userExist": true}`))
		} else {
			// 5 - Create the db Account struct (call dbCreate, mutates into full Account with key and timestamp)
			if err := DBCreate(&account, ctx); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}

	}


	//if js, err := json.Marshal(account); err != nil {
	//	log.Errorf(ctx,"Cannot parse JSON")
	//} else {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.Write(js)
	//}
}


















//func Handleincoming(w http.ResponseWriter, r *http.Request) {
//	ctx := appengine.NewContext(r)
//	received_json, _ := ioutil.ReadAll(r.Body)
//	jsonparser.ObjectEach(received_json, func(key []byte, value []byte, dataType jsonparser.ValueType, offset int) error {
//		var result = fmt.Sprintf("Key: '%s'\n Value: '%s'\n Type: %s\n", string(key), string(value), dataType)
//		w.Header().Set("content-type", "application/text")
//		w.Write([]byte(result))
//		log.Infof(ctx, result)
//		return nil
//	})
//	//json.Unmarshal([]byte(received_json), &arbitrary_json)
//	//js, err := json.Marshal(arbitrary_json)
//	//w.Header().Set("Content-Type", "application/json")
//	//w.Write(received_json)
//}
//
//func Handle(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "Hello, world!")
//
//}