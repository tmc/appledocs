// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentHistoryChangeRequest] class.
var (
	PersistentHistoryChangeRequestClass     _PersistentHistoryChangeRequestClass
	PersistentHistoryChangeRequestClassOnce sync.Once
)

func getPersistentHistoryChangeRequestClass() _PersistentHistoryChangeRequestClass {
	PersistentHistoryChangeRequestClassOnce.Do(func() {
		PersistentHistoryChangeRequestClass = _PersistentHistoryChangeRequestClass{objc.GetClass("NSPersistentHistoryChangeRequest")}
	})
	return PersistentHistoryChangeRequestClass
}

type _PersistentHistoryChangeRequestClass struct {
	class objc.Class
}

// An interface definition for the [PersistentHistoryChangeRequest] class.
type IPersistentHistoryChangeRequest interface {
	IPersistentStoreRequest
}

// A request to fetch or purge persistent history.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest
type PersistentHistoryChangeRequest struct {
	PersistentStoreRequest
}

// PersistentHistoryChangeRequestFrom constructs a [PersistentHistoryChangeRequest] from an unsafe.Pointer.
//
// A request to fetch or purge persistent history.
func PersistentHistoryChangeRequestFrom(ptr unsafe.Pointer) PersistentHistoryChangeRequest {
	return PersistentHistoryChangeRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentHistoryChangeRequestClass) Alloc() PersistentHistoryChangeRequest {
	rv := objc.Send[PersistentHistoryChangeRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentHistoryChangeRequestClass) New() PersistentHistoryChangeRequest {
	rv := objc.Send[PersistentHistoryChangeRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentHistoryChangeRequest) Init() PersistentHistoryChangeRequest {
	rv := objc.Send[PersistentHistoryChangeRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentHistoryChangeRequest) Autorelease() PersistentHistoryChangeRequest {
	rv := objc.Send[PersistentHistoryChangeRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentHistoryChangeRequest creates a new PersistentHistoryChangeRequest instance.
func NewPersistentHistoryChangeRequest() PersistentHistoryChangeRequest {
	return getPersistentHistoryChangeRequestClass().New()
}


// Purges history older than that defined by a given token.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/deleteHistory(before:)-5kghb
func (pc _PersistentHistoryChangeRequestClass) DeleteHistoryBeforeToken(token unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("deleteHistoryBeforeToken:"), token)
	return rv
}

// Retrieves the request history after a given token.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(after:)-3rmfm
func (pc _PersistentHistoryChangeRequestClass) FetchHistoryAfterToken(token unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchHistoryAfterToken:"), token)
	return rv
}

// Retrieves history since a given date.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(after:)-qi5b
func (pc _PersistentHistoryChangeRequestClass) FetchHistoryAfterDate(date unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchHistoryAfterDate:"), date)
	return rv
}

// Retrieves history based on a fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(withFetch:)
func (pc _PersistentHistoryChangeRequestClass) FetchHistoryWithFetchRequest(fetchRequest unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchHistoryWithFetchRequest:"), fetchRequest)
	return rv
}

// The specified fetch request, when retrieving history.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchRequest
func (p_ PersistentHistoryChangeRequest) FetchRequest() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("fetchRequest"))
	return rv
}


// SetFetchRequest sets the value of the fetchRequest property.
// The specified fetch request, when retrieving history.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchRequest
func (p_ PersistentHistoryChangeRequest) SetFetchRequest(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setFetchRequest:"), value)
}
// The type of result that this request returns.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/resultType
func (p_ PersistentHistoryChangeRequest) ResultType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("resultType"))
	return rv
}


// SetResultType sets the value of the resultType property.
// The type of result that this request returns.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/resultType
func (p_ PersistentHistoryChangeRequest) SetResultType(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setResultType:"), value)
}


