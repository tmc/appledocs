// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentCloudKitContainerEventResult] class.
var (
	PersistentCloudKitContainerEventResultClass     _PersistentCloudKitContainerEventResultClass
	PersistentCloudKitContainerEventResultClassOnce sync.Once
)

func getPersistentCloudKitContainerEventResultClass() _PersistentCloudKitContainerEventResultClass {
	PersistentCloudKitContainerEventResultClassOnce.Do(func() {
		PersistentCloudKitContainerEventResultClass = _PersistentCloudKitContainerEventResultClass{objc.GetClass("NSPersistentCloudKitContainerEventResult")}
	})
	return PersistentCloudKitContainerEventResultClass
}

type _PersistentCloudKitContainerEventResultClass struct {
	class objc.Class
}

// An interface definition for the [PersistentCloudKitContainerEventResult] class.
type IPersistentCloudKitContainerEventResult interface {
	IPersistentStoreResult
	Result() objc.ID
	ResultType() PersistentCloudKitContainerEventResultType
}

// The result of a request to fetch persistent CloudKit container events.


// The result of a request to fetch persistent CloudKit container events.
//
// [Full Topic]
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getPersistentCloudKitContainerEventResultClass().New()
}



// The result of the persistent CloudKit container event request, which the result type determines.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult/result
func (p_ PersistentCloudKitContainerEventResult) Result() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("result"))
	return rv
}


// The type of result that the CloudKit container event fetch request returns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult/resultType-swift.property
func (p_ PersistentCloudKitContainerEventResult) ResultType() PersistentCloudKitContainerEventResultType {
	rv := objc.Send[PersistentCloudKitContainerEventResultType](p_.ID, objc.Sel("resultType"))
	return rv
}



