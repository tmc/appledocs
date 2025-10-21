// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CNFetchRequest] class.
var (
	CNFetchRequestClass     _CNFetchRequestClass
	CNFetchRequestClassOnce sync.Once
)

func getCNFetchRequestClass() _CNFetchRequestClass {
	CNFetchRequestClassOnce.Do(func() {
		CNFetchRequestClass = _CNFetchRequestClass{objc.GetClass("CNFetchRequest")}
	})
	return CNFetchRequestClass
}

type _CNFetchRequestClass struct {
	class objc.Class
}

// An interface definition for the [CNFetchRequest] class.
type ICNFetchRequest interface {
	objectivec.IObject
}

// The base class for contact fetch requests.
//
// To fetch contacts, use .
//
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNFetchRequest
type CNFetchRequest struct {
	objectivec.Object
}

// CNFetchRequestFrom constructs a [CNFetchRequest] from an unsafe.Pointer.
//
// The base class for contact fetch requests.
func CNFetchRequestFrom(ptr unsafe.Pointer) CNFetchRequest {
	return CNFetchRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNFetchRequestClass) Alloc() CNFetchRequest {
	rv := objc.Send[CNFetchRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNFetchRequestClass) New() CNFetchRequest {
	rv := objc.Send[CNFetchRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNFetchRequest) Init() CNFetchRequest {
	rv := objc.Send[CNFetchRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNFetchRequest) Autorelease() CNFetchRequest {
	rv := objc.Send[CNFetchRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNFetchRequest creates a new CNFetchRequest instance.
func NewCNFetchRequest() CNFetchRequest {
	return getCNFetchRequestClass().New()
}




