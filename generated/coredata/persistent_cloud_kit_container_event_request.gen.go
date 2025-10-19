// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentCloudKitContainerEventRequest] class.
var persistentCloudKitContainerEventRequestClass = _PersistentCloudKitContainerEventRequestClass{objc.GetClass("NSPersistentCloudKitContainerEventRequest")}

type _PersistentCloudKitContainerEventRequestClass struct {
	class objc.Class
}

// A request to fetch setup, import, or export events in a persistent CloudKit container. [Full Topic]
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

// Creates a fetch request for events that occur after a specified event from a persistent CloudKit container. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchEvents(after:)-3yfp
func (pc _PersistentCloudKitContainerEventRequestClass) FetchEventsAfterEvent(event unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchEventsAfterEvent:"), event)
	return rv
}
// Creates a fetch request for events after a specified date from a persistent CloudKit container. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchEvents(after:)-5izg7
func (pc _PersistentCloudKitContainerEventRequestClass) FetchEventsAfterDate(date unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchEventsAfterDate:"), date)
	return rv
}
// Creates a fetch request for events that match a specified fetch request from a persistent CloudKit container. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchEvents(matchingFetch:)
func (pc _PersistentCloudKitContainerEventRequestClass) FetchEventsMatchingFetchRequest(fetchRequest unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchEventsMatchingFetchRequest:"), fetchRequest)
	return rv
}
// Creates a fetch request for all events in a persistent CloudKit container. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventRequest/fetchForEvents()
func (pc _PersistentCloudKitContainerEventRequestClass) FetchRequestForEvents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("fetchRequestForEvents"))
	return rv
}


