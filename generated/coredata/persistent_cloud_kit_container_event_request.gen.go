// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentCloudKitContainerEventRequest] class.
var (
	persistentCloudKitContainerEventRequestClass     _PersistentCloudKitContainerEventRequestClass
	persistentCloudKitContainerEventRequestClassOnce sync.Once
)

func getPersistentCloudKitContainerEventRequestClass() _PersistentCloudKitContainerEventRequestClass {
	persistentCloudKitContainerEventRequestClassOnce.Do(func() {
		persistentCloudKitContainerEventRequestClass = _PersistentCloudKitContainerEventRequestClass{objc.GetClass("NSPersistentCloudKitContainerEventRequest")}
	})
	return persistentCloudKitContainerEventRequestClass
}

type _PersistentCloudKitContainerEventRequestClass struct {
	class objc.Class
}

// An interface definition for the [PersistentCloudKitContainerEventRequest] class.
type IPersistentCloudKitContainerEventRequest interface {
	IPersistentStoreRequest
}

// A request to fetch setup, import, or export events in a persistent CloudKit container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest
type PersistentCloudKitContainerEventRequest struct {
	PersistentStoreRequest
}

// PersistentCloudKitContainerEventRequestFrom constructs a [PersistentCloudKitContainerEventRequest] from an unsafe.Pointer.
//
// A request to fetch setup, import, or export events in a persistent CloudKit container.
func PersistentCloudKitContainerEventRequestFrom(ptr unsafe.Pointer) PersistentCloudKitContainerEventRequest {
	return PersistentCloudKitContainerEventRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentCloudKitContainerEventRequestClass) Alloc() PersistentCloudKitContainerEventRequest {
	rv := objc.Send[PersistentCloudKitContainerEventRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentCloudKitContainerEventRequestClass) New() PersistentCloudKitContainerEventRequest {
	rv := objc.Send[PersistentCloudKitContainerEventRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentCloudKitContainerEventRequest) Init() PersistentCloudKitContainerEventRequest {
	rv := objc.Send[PersistentCloudKitContainerEventRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentCloudKitContainerEventRequest) Autorelease() PersistentCloudKitContainerEventRequest {
	rv := objc.Send[PersistentCloudKitContainerEventRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentCloudKitContainerEventRequest creates a new PersistentCloudKitContainerEventRequest instance.
func NewPersistentCloudKitContainerEventRequest() PersistentCloudKitContainerEventRequest {
	return getPersistentCloudKitContainerEventRequestClass().New()
}


// Creates a fetch request for events that occur after a specified event from a persistent CloudKit container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchEvents(after:)-3yfp
func (pc _PersistentCloudKitContainerEventRequestClass) FetchEventsAfterEvent(event unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchEventsAfterEvent:"), event)
	return rv
}
// Creates a fetch request for events after a specified date from a persistent CloudKit container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchEvents(after:)-5izg7
func (pc _PersistentCloudKitContainerEventRequestClass) FetchEventsAfterDate(date unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchEventsAfterDate:"), date)
	return rv
}
// Creates a fetch request for events that match a specified fetch request from a persistent CloudKit container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchEvents(matchingFetch:)
func (pc _PersistentCloudKitContainerEventRequestClass) FetchEventsMatchingFetchRequest(fetchRequest unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchEventsMatchingFetchRequest:"), fetchRequest)
	return rv
}
// Creates a fetch request for all events in a persistent CloudKit container.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchForEvents()
func (pc _PersistentCloudKitContainerEventRequestClass) FetchRequestForEvents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchRequestForEvents"))
	return rv
}


