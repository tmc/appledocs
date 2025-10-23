// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SaveChangesRequest] class.
var (
	SaveChangesRequestClass     _SaveChangesRequestClass
	SaveChangesRequestClassOnce sync.Once
)

func getSaveChangesRequestClass() _SaveChangesRequestClass {
	SaveChangesRequestClassOnce.Do(func() {
		SaveChangesRequestClass = _SaveChangesRequestClass{objc.GetClass("NSSaveChangesRequest")}
	})
	return SaveChangesRequestClass
}

type _SaveChangesRequestClass struct {
	class objc.Class
}

// An interface definition for the [SaveChangesRequest] class.
type ISaveChangesRequest interface {
	IPersistentStoreRequest
	// properties:
	UpdatedObjects() unsafe.Pointer
	DeletedObjects() IManagedObject
	SetDeletedObjects(value IManagedObject)
	InsertedObjects() IManagedObject
	SetInsertedObjects(value IManagedObject)
	LockedObjects() IManagedObject
	SetLockedObjects(value IManagedObject)
	// methods:
}

// An encapsulation of a collection of changes to be made by an object store in response to a save operation on a managed object context.


// An encapsulation of a collection of changes to be made by an object store in response to a save operation on a managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest
type SaveChangesRequest struct {
	PersistentStoreRequest
}

// SaveChangesRequestFrom constructs a [SaveChangesRequest] from an unsafe.Pointer.
//
// An encapsulation of a collection of changes to be made by an object store in response to a save operation on a managed object context.
func SaveChangesRequestFrom(ptr unsafe.Pointer) SaveChangesRequest {
	return SaveChangesRequest{
		PersistentStoreRequest: PersistentStoreRequestFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SaveChangesRequestClass) Alloc() SaveChangesRequest {
	rv := objc.Send[SaveChangesRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SaveChangesRequestClass) New() SaveChangesRequest {
	rv := objc.Send[SaveChangesRequest](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SaveChangesRequest) Init() SaveChangesRequest {
	rv := objc.Send[SaveChangesRequest](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SaveChangesRequest) Autorelease() SaveChangesRequest {
	rv := objc.Send[SaveChangesRequest](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSaveChangesRequest creates a new SaveChangesRequest instance.
func NewSaveChangesRequest() SaveChangesRequest {
	return getSaveChangesRequestClass().New()
}



// The objects that were modified in the calling context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSSaveChangesRequest/updatedObjects
func (s_ SaveChangesRequest) UpdatedObjects() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("updatedObjects"))
	return rv
}


// The objects that were deleted in the calling context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/deletedobjects
func (s_ SaveChangesRequest) DeletedObjects() IManagedObject {
	rv := objc.Send[ManagedObject](s_.ID, objc.Sel("deletedObjects"))
	return rv
}


// The objects that were deleted in the calling context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/deletedobjects
func (s_ SaveChangesRequest) SetDeletedObjects(value IManagedObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDeletedObjects:"), value)
}


// The objects that were inserted into the calling context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/insertedobjects
func (s_ SaveChangesRequest) InsertedObjects() IManagedObject {
	rv := objc.Send[ManagedObject](s_.ID, objc.Sel("insertedObjects"))
	return rv
}


// The objects that were inserted into the calling context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/insertedobjects
func (s_ SaveChangesRequest) SetInsertedObjects(value IManagedObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInsertedObjects:"), value)
}


// The objects that were flagged for optimistic locking on the calling context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/lockedobjects
func (s_ SaveChangesRequest) LockedObjects() IManagedObject {
	rv := objc.Send[ManagedObject](s_.ID, objc.Sel("lockedObjects"))
	return rv
}


// The objects that were flagged for optimistic locking on the calling context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nssavechangesrequest/lockedobjects
func (s_ SaveChangesRequest) SetLockedObjects(value IManagedObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setLockedObjects:"), value)
}



