// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentCloudKitContainerEventResult] class.
var persistentCloudKitContainerEventResultClass = _PersistentCloudKitContainerEventResultClass{objc.GetClass("NSPersistentCloudKitContainerEventResult")}

type _PersistentCloudKitContainerEventResultClass struct {
	class objc.Class
}

// An interface definition for the [PersistentCloudKitContainerEventResult] class.
type IPersistentCloudKitContainerEventResult interface {
	IPersistentStoreResult
}

// The result of a request to fetch persistent CloudKit container events. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult

type PersistentCloudKitContainerEventResult struct {
	PersistentStoreResult
}

// PersistentCloudKitContainerEventResultFrom constructs a [PersistentCloudKitContainerEventResult] from an unsafe.Pointer.
//
// The result of a request to fetch persistent CloudKit container events.
func PersistentCloudKitContainerEventResultFrom(ptr unsafe.Pointer) PersistentCloudKitContainerEventResult {
	return PersistentCloudKitContainerEventResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (pc _PersistentCloudKitContainerEventResultClass) Alloc() PersistentCloudKitContainerEventResult {
	rv := objc.Send[PersistentCloudKitContainerEventResult](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PersistentCloudKitContainerEventResultClass) New() PersistentCloudKitContainerEventResult {
	rv := objc.Send[PersistentCloudKitContainerEventResult](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentCloudKitContainerEventResult) Init() PersistentCloudKitContainerEventResult {
	rv := objc.Send[PersistentCloudKitContainerEventResult](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentCloudKitContainerEventResult) Autorelease() PersistentCloudKitContainerEventResult {
	rv := objc.Send[PersistentCloudKitContainerEventResult](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentCloudKitContainerEventResult creates a new PersistentCloudKitContainerEventResult instance.
func NewPersistentCloudKitContainerEventResult() PersistentCloudKitContainerEventResult {
	return persistentCloudKitContainerEventResultClass.New()
}




