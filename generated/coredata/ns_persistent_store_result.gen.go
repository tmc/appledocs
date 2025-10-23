// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentStoreResult] class.
var (
	PersistentStoreResultClass     _PersistentStoreResultClass
	PersistentStoreResultClassOnce sync.Once
)

func getPersistentStoreResultClass() _PersistentStoreResultClass {
	PersistentStoreResultClassOnce.Do(func() {
		PersistentStoreResultClass = _PersistentStoreResultClass{objc.GetClass("NSPersistentStoreResult")}
	})
	return PersistentStoreResultClass
}

type _PersistentStoreResultClass struct {
	class objc.Class
}

// An interface definition for the [PersistentStoreResult] class.
type IPersistentStoreResult interface {
	objectivec.IObject
}

// The abstract base class for results returned from a persistent store coordinator.


// The abstract base class for results returned from a persistent store coordinator.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentStoreResult
type PersistentStoreResult struct {
	objectivec.Object
}

// PersistentStoreResultFrom constructs a [PersistentStoreResult] from an unsafe.Pointer.
//
// The abstract base class for results returned from a persistent store coordinator.
func PersistentStoreResultFrom(ptr unsafe.Pointer) PersistentStoreResult {
	return PersistentStoreResult{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentStoreResultClass) Alloc() PersistentStoreResult {
	rv := objc.Send[PersistentStoreResult](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentStoreResultClass) New() PersistentStoreResult {
	rv := objc.Send[PersistentStoreResult](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentStoreResult) Init() PersistentStoreResult {
	rv := objc.Send[PersistentStoreResult](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentStoreResult) Autorelease() PersistentStoreResult {
	rv := objc.Send[PersistentStoreResult](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentStoreResult creates a new PersistentStoreResult instance.
func NewPersistentStoreResult() PersistentStoreResult {
	return getPersistentStoreResultClass().New()
}




