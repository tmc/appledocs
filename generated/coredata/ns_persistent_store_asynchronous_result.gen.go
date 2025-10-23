// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PersistentStoreAsynchronousResult] class.
var (
	PersistentStoreAsynchronousResultClass     _PersistentStoreAsynchronousResultClass
	PersistentStoreAsynchronousResultClassOnce sync.Once
)

func getPersistentStoreAsynchronousResultClass() _PersistentStoreAsynchronousResultClass {
	PersistentStoreAsynchronousResultClassOnce.Do(func() {
		PersistentStoreAsynchronousResultClass = _PersistentStoreAsynchronousResultClass{objc.GetClass("NSPersistentStoreAsynchronousResult")}
	})
	return PersistentStoreAsynchronousResultClass
}

type _PersistentStoreAsynchronousResultClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStoreAsynchronousResult] class.
type IPersistentStoreAsynchronousResult interface {
	IPersistentStoreResult
	// properties:
	ManagedObjectContext() IManagedObjectContext
	OperationError() Error /* not a class type */
	Progress() Progress /* not a class type */
	// methods:
	Cancel()
}

// A concrete class used to represent the results of an asynchronous request.


// A concrete class used to represent the results of an asynchronous request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult
type PersistentStoreAsynchronousResult struct {
	PersistentStoreResult
}

// PersistentStoreAsynchronousResultFrom constructs a [PersistentStoreAsynchronousResult] from an unsafe.Pointer.
//
// A concrete class used to represent the results of an asynchronous request.
func PersistentStoreAsynchronousResultFrom(ptr unsafe.Pointer) PersistentStoreAsynchronousResult {
	return PersistentStoreAsynchronousResult{
		PersistentStoreResult: PersistentStoreResultFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentStoreAsynchronousResultClass) Alloc() PersistentStoreAsynchronousResult {
	rv := objc.Send[PersistentStoreAsynchronousResult](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentStoreAsynchronousResultClass) New() PersistentStoreAsynchronousResult {
	rv := objc.Send[PersistentStoreAsynchronousResult](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentStoreAsynchronousResult) Init() PersistentStoreAsynchronousResult {
	rv := objc.Send[PersistentStoreAsynchronousResult](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentStoreAsynchronousResult) Autorelease() PersistentStoreAsynchronousResult {
	rv := objc.Send[PersistentStoreAsynchronousResult](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentStoreAsynchronousResult creates a new PersistentStoreAsynchronousResult instance.
func NewPersistentStoreAsynchronousResult() PersistentStoreAsynchronousResult {
	return getPersistentStoreAsynchronousResultClass().New()
}



// Cancels the asynchronous fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult/cancel()
func (p_ PersistentStoreAsynchronousResult) Cancel() {
	objc.Send[objc.ID](p_.ID, objc.Sel("cancel"))
}


// The managed object context for the result.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult/managedObjectContext
func (p_ PersistentStoreAsynchronousResult) ManagedObjectContext() IManagedObjectContext {
	rv := objc.Send[ManagedObjectContext](p_.ID, objc.Sel("managedObjectContext"))
	return rv
}


// An error that contains details if the asynchronous fetch request fails.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult/operationError
func (p_ PersistentStoreAsynchronousResult) OperationError() Error /* not a class type */ {
	rv := objc.Send[Error](p_.ID, objc.Sel("operationError"))
	return rv
}


// An object that reports progress for the asynchronous fetch request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult/progress
func (p_ PersistentStoreAsynchronousResult) Progress() Progress /* not a class type */ {
	rv := objc.Send[Progress](p_.ID, objc.Sel("progress"))
	return rv
}



