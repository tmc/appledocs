// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [PersistentStoreRequest] class.
var (
	PersistentStoreRequestClass     _PersistentStoreRequestClass
	PersistentStoreRequestClassOnce sync.Once
)

func getPersistentStoreRequestClass() _PersistentStoreRequestClass {
	PersistentStoreRequestClassOnce.Do(func() {
		PersistentStoreRequestClass = _PersistentStoreRequestClass{objc.GetClass("NSPersistentStoreRequest")}
	})
	return PersistentStoreRequestClass
}

type _PersistentStoreRequestClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStoreRequest] class.
type IPersistentStoreRequest interface {
	objectivec.IObject
}

// Criteria used to retrieve data from or save data to a persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequest
type PersistentStoreRequest struct {
	objectivec.Object
}

// PersistentStoreRequestFrom constructs a [PersistentStoreRequest] from an unsafe.Pointer.
//
// Criteria used to retrieve data from or save data to a persistent store.
func PersistentStoreRequestFrom(ptr unsafe.Pointer) PersistentStoreRequest {
	return PersistentStoreRequest{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentStoreRequestClass) Alloc() PersistentStoreRequest {
	rv := objc.Send[PersistentStoreRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentStoreRequestClass) New() PersistentStoreRequest {
	rv := objc.Send[PersistentStoreRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentStoreRequest) Init() PersistentStoreRequest {
	rv := objc.Send[PersistentStoreRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentStoreRequest) Autorelease() PersistentStoreRequest {
	rv := objc.Send[PersistentStoreRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentStoreRequest creates a new PersistentStoreRequest instance.
func NewPersistentStoreRequest() PersistentStoreRequest {
	return getPersistentStoreRequestClass().New()
}


// The stores the request should be sent to.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequest/affectedStores
func (p_ PersistentStoreRequest) AffectedStores() []PersistentStore {
	rv := objc.Send[[]PersistentStore](p_.ID, objc.Sel("affectedStores"))
	return rv
}


// SetAffectedStores sets the value of the affectedStores property.
// The stores the request should be sent to.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequest/affectedStores
func (p_ PersistentStoreRequest) SetAffectedStores(value []PersistentStore) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](p_.ID, objc.Sel("setAffectedStores:"), nsArray)
}
// The type of the fetch request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreRequest/requestType
func (p_ PersistentStoreRequest) RequestType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("requestType"))
	return rv
}



