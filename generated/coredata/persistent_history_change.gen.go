// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentHistoryChange] class.
var (
	persistentHistoryChangeClass     _PersistentHistoryChangeClass
	persistentHistoryChangeClassOnce sync.Once
)

func getPersistentHistoryChangeClass() _PersistentHistoryChangeClass {
	persistentHistoryChangeClassOnce.Do(func() {
		persistentHistoryChangeClass = _PersistentHistoryChangeClass{objc.GetClass("NSPersistentHistoryChange")}
	})
	return persistentHistoryChangeClass
}

type _PersistentHistoryChangeClass struct {
	class objc.Class
}

// An interface definition for the [PersistentHistoryChange] class.
type IPersistentHistoryChange interface {
	objectivec.IObject
}

// A change representing the insertion, update, or deletion of a managed object in the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange
type PersistentHistoryChange struct {
	objectivec.Object
}

// PersistentHistoryChangeFrom constructs a [PersistentHistoryChange] from an unsafe.Pointer.
//
// A change representing the insertion, update, or deletion of a managed object in the persistent store.
func PersistentHistoryChangeFrom(ptr unsafe.Pointer) PersistentHistoryChange {
	return PersistentHistoryChange{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentHistoryChangeClass) Alloc() PersistentHistoryChange {
	rv := objc.Send[PersistentHistoryChange](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentHistoryChangeClass) New() PersistentHistoryChange {
	rv := objc.Send[PersistentHistoryChange](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentHistoryChange) Init() PersistentHistoryChange {
	rv := objc.Send[PersistentHistoryChange](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentHistoryChange) Autorelease() PersistentHistoryChange {
	rv := objc.Send[PersistentHistoryChange](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentHistoryChange creates a new PersistentHistoryChange instance.
func NewPersistentHistoryChange() PersistentHistoryChange {
	return getPersistentHistoryChangeClass().New()
}


// The change’s numeric identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/changeID
func (p_ PersistentHistoryChange) ChangeID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("changeID"))
	return rv
}

// The type of change to the managed object in the persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/changeType
func (p_ PersistentHistoryChange) ChangeType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("changeType"))
	return rv
}

// The identifier of the managed object that changed. (swift) Declaration: @property(readonly, copy) NSManagedObjectID *changedObjectID; (objc) Availability: iOS: 11.0 — iPadOS: 11.0 — Mac Catalyst: 13.1 — macOS: 10.13 — tvOS: 11.0 — visionOS: 1.0 — watchOS: 4.0 (objc,swift) }
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/changedObjectID
func (p_ PersistentHistoryChange) ChangedObjectID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("changedObjectID"))
	return rv
}

// A dictionary of attributes marked for preservation after deletion, and their values when deleted.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/tombstone
func (p_ PersistentHistoryChange) Tombstone() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("tombstone"))
	return rv
}

// The persistent history transaction containing this change.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/transaction
func (p_ PersistentHistoryChange) Transaction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("transaction"))
	return rv
}

// The set of properties that were updated on the managed object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/updatedProperties
func (p_ PersistentHistoryChange) UpdatedProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("updatedProperties"))
	return rv
}



