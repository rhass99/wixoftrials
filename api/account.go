package api
// Every Entity should satisfy all db-struct interfaces
// 1 - dbEntityCreator - OK
// 2 - dbKeyCheckerCreator
// 2 - dbPutter
// 3 - dbSearcher

//// Check if user is logged in using the CheckLoggedIn function
//_, exists := ul.CheckLoggedIn(ctx)
//if exists == true {
//	return errors.New("User is logged in, kindly logout first")
//}
//// Check if there is a key or create a new key with email as ID

import (
	"context"
	"errors"
	"google.golang.org/appengine/datastore"
	"time"
)

type Account struct {
	Key *datastore.Key `json:"id" datastore:"-"`
	Email string `json:"email"`
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	Password string `json:"password"`
	Gender string `json:"gender"`
	CTime time.Time `json:"created"`
}

//Checks if the current user is logged in and returns App engine user struct, returns error if not
//func (ul *Account) CheckLoggedIn(ctx context.Context) (*user.User, bool){
//	aeuser := user.Current(ctx)
//	if aeuser == nil {
//		return nil, false
//	}
//	return aeuser, true
//}

// Check if user already exists in Database, returns it.
func (ul *Account) dbCheckEntity(ctx context.Context) (bool, error) {

	if ul.Email == "" {
		return true, errors.New("No Email to check")
	}
	return true, nil

	//q := datastore.NewQuery("Account").Filter("Email=", ul.Email).KeysOnly()
	//
	//iter := q.Run(ctx)
	//for {
	//	var newKey *datastore.Key
	//	key, err := iter.Next(newKey);
	//	if err == datastore.Done {
	//		break
	//	}
	//	if err != nil {
	//		return true, err
	//	}
	//	if key != nil {
	//		return true, nil
	//	}
	//}
	//return true, nil
}



	//
	//{
	//	return false, err
	//} else {
	//	if newAccount[0].Key == nil {
	//		return false, nil
	//	}
	//	return true, nil
	//}


func (ul *Account) dbCreate(ctx context.Context) (error) {
	// Create timestamp and add it to Account
	if newKey, err := datastore.Put(ctx, ul.Key, ul); err != nil {
		return err
	} else {
		ul.CTime = time.Now()
		ul.Key = newKey
		datastore.Get(ctx, ul.Key, nil)
		return nil
	}
}
