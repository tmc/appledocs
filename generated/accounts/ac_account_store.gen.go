// Code generated from Apple documentation for Accounts. DO NOT EDIT.

package accounts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ACAccountStore] class.
var (
	ACAccountStoreClass     _ACAccountStoreClass
	ACAccountStoreClassOnce sync.Once
)

func getACAccountStoreClass() _ACAccountStoreClass {
	ACAccountStoreClassOnce.Do(func() {
		ACAccountStoreClass = _ACAccountStoreClass{objc.GetClass("ACAccountStore")}
	})
	return ACAccountStoreClass
}

type _ACAccountStoreClass struct {
	class objc.Class
}

// An interface definition for the [ACAccountStore] class.
type IACAccountStore interface {
	objectivec.IObject
	Accounts() objc.ID
}

// The object you use to request, manage, and store the user’s account information.
//
// The class provides an interface for accessing, managing, and storing accounts. To create and retrieve accounts from the Accounts database, you must create an object. Each object belongs to a single account store object.


// The object you use to request, manage, and store the user’s account information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore
type ACAccountStore struct {
	objectivec.Object
}

// ACAccountStoreFrom constructs a [ACAccountStore] from an unsafe.Pointer.
//
// The object you use to request, manage, and store the user’s account information.
func ACAccountStoreFrom(ptr unsafe.Pointer) ACAccountStore {
	return ACAccountStore{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ACAccountStoreClass) Alloc() ACAccountStore {
	rv := objc.Send[ACAccountStore](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ACAccountStoreClass) New() ACAccountStore {
	rv := objc.Send[ACAccountStore](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ ACAccountStore) Init() ACAccountStore {
	rv := objc.Send[ACAccountStore](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ ACAccountStore) Autorelease() ACAccountStore {
	rv := objc.Send[ACAccountStore](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewACAccountStore creates a new ACAccountStore instance.
func NewACAccountStore() ACAccountStore {
	return getACAccountStoreClass().New()
}



// The accounts managed by this account store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Accounts/ACAccountStore/accounts
func (a_ ACAccountStore) Accounts() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("accounts"))
	return rv
}



