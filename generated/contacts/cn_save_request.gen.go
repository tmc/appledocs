// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNSaveRequest] class.
var (
	CNSaveRequestClass     _CNSaveRequestClass
	CNSaveRequestClassOnce sync.Once
)

func getCNSaveRequestClass() _CNSaveRequestClass {
	CNSaveRequestClassOnce.Do(func() {
		CNSaveRequestClass = _CNSaveRequestClass{objc.GetClass("CNSaveRequest")}
	})
	return CNSaveRequestClass
}

type _CNSaveRequestClass struct {
	class objc.Class
}

// An interface definition for the [CNSaveRequest] class.
type ICNSaveRequest interface {
	objectivec.IObject
	CNErrorUserInfoAffectedRecordsKey() string
	ShouldRefetchContacts() bool
	SetShouldRefetchContacts(value bool)
	TransactionAuthor() string
	SetTransactionAuthor(value string)
}

// An object that collects the changes you want to save to the user’s contacts database.
//
// Create a new object for each save operation you want to make. You can batch multiple changes into one save request (note that these changes only apply to objects). In the case of overlapping changes in multiple or concurrent save requests, the last change wins. If you try to add an object (that is, a contact or a group) that already exists in the contact store, you receive the error and the array is updated to contain the object you tried to add. If you try to update or delete an object that is not present in the contact store, the save request does not perform the update or deletion, the error occurs, and the array is updated to contain the object you tried to update or delete. Do not access objects in the save request while that request is executing.


// An object that collects the changes you want to save to the user’s contacts database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNSaveRequest
type CNSaveRequest struct {
	objectivec.Object
}

// CNSaveRequestFrom constructs a [CNSaveRequest] from an unsafe.Pointer.
//
// An object that collects the changes you want to save to the user’s contacts database.
func CNSaveRequestFrom(ptr unsafe.Pointer) CNSaveRequest {
	return CNSaveRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNSaveRequestClass) Alloc() CNSaveRequest {
	rv := objc.Send[CNSaveRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNSaveRequestClass) New() CNSaveRequest {
	rv := objc.Send[CNSaveRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNSaveRequest) Init() CNSaveRequest {
	rv := objc.Send[CNSaveRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNSaveRequest) Autorelease() CNSaveRequest {
	rv := objc.Send[CNSaveRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNSaveRequest creates a new CNSaveRequest instance.
func NewCNSaveRequest() CNSaveRequest {
	return getCNSaveRequestClass().New()
}



// The contact, group, and container objects for which the error code applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnerroruserinfoaffectedrecordskey
func (c_ CNSaveRequest) CNErrorUserInfoAffectedRecordsKey() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CNErrorUserInfoAffectedRecordsKey"))
	return rv
}


// A Boolean value that indicates whether to refetch the added and updated contacts after the save request executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/shouldrefetchcontacts
func (c_ CNSaveRequest) ShouldRefetchContacts() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("shouldRefetchContacts"))
	return rv
}


// A Boolean value that indicates whether to refetch the added and updated contacts after the save request executes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/shouldrefetchcontacts
func (c_ CNSaveRequest) SetShouldRefetchContacts(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setShouldRefetchContacts:"), value)
}


// A string that identifies the author of the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/transactionauthor
func (c_ CNSaveRequest) TransactionAuthor() string {
	rv := objc.Send[string](c_.ID, objc.Sel("transactionAuthor"))
	return rv
}


// A string that identifies the author of the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnsaverequest/transactionauthor
func (c_ CNSaveRequest) SetTransactionAuthor(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTransactionAuthor:"), objc.String(value))
}



