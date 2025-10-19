// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentHistoryChangeRequest] class.
var persistentHistoryChangeRequestClass = _PersistentHistoryChangeRequestClass{objc.GetClass("NSPersistentHistoryChangeRequest")}

type _PersistentHistoryChangeRequestClass struct {
	class objc.Class
}

// An interface definition for the [PersistentHistoryChangeRequest] class.
type IPersistentHistoryChangeRequest interface {
	IPersistentStoreRequest
}

// A request to fetch or purge persistent history. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return persistentHistoryChangeRequestClass.New()
}


// Purges history older than that defined by a given token. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/deleteHistory(before:)-5kghb
func (pc _PersistentHistoryChangeRequestClass) DeleteHistoryBeforeToken(token unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("deleteHistoryBeforeToken:"), token)
	return rv
}
// Retrieves the request history after a given token. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(after:)-3rmfm
func (pc _PersistentHistoryChangeRequestClass) FetchHistoryAfterToken(token unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchHistoryAfterToken:"), token)
	return rv
}
// Retrieves history since a given date. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(after:)-qi5b
func (pc _PersistentHistoryChangeRequestClass) FetchHistoryAfterDate(date unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchHistoryAfterDate:"), date)
	return rv
}
// Retrieves history based on a fetch request. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChangeRequest/fetchHistory(withFetch:)
func (pc _PersistentHistoryChangeRequestClass) FetchHistoryWithFetchRequest(fetchRequest unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchHistoryWithFetchRequest:"), fetchRequest)
	return rv
}


