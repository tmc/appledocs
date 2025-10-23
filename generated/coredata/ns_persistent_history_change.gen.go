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
	PersistentHistoryChangeClass     _PersistentHistoryChangeClass
	PersistentHistoryChangeClassOnce sync.Once
)

func getPersistentHistoryChangeClass() _PersistentHistoryChangeClass {
	PersistentHistoryChangeClassOnce.Do(func() {
		PersistentHistoryChangeClass = _PersistentHistoryChangeClass{objc.GetClass("NSPersistentHistoryChange")}
	})
	return PersistentHistoryChangeClass
}

type _PersistentHistoryChangeClass struct {
	class objc.Class
}

// An interface definition for the [PersistentHistoryChange] class.
type IPersistentHistoryChange interface {
	objectivec.IObject
	// properties:
	ChangeID() unsafe.Pointer
	Tombstone() objc.ID
	UpdatedProperties() unsafe.Pointer
	ChangeType() PersistentHistoryChangeType /* not a class type */
	SetChangeType(value PersistentHistoryChangeType /* not a class type */)
	ChangedObjectID() IManagedObjectID
	SetChangedObjectID(value IManagedObjectID)
	Transaction() IPersistentHistoryTransaction
	SetTransaction(value IPersistentHistoryTransaction)
	// methods:
}

// A change representing the insertion, update, or deletion of a managed object in the persistent store.


// A change representing the insertion, update, or deletion of a managed object in the persistent store.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/changeID
func (p_ PersistentHistoryChange) ChangeID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("changeID"))
	return rv
}


// A dictionary of attributes marked for preservation after deletion, and their values when deleted.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/tombstone
func (p_ PersistentHistoryChange) Tombstone() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("tombstone"))
	return rv
}


// The set of properties that were updated on the managed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryChange/updatedProperties
func (p_ PersistentHistoryChange) UpdatedProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("updatedProperties"))
	return rv
}


// The type of change to the managed object in the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/changetype
func (p_ PersistentHistoryChange) ChangeType() PersistentHistoryChangeType /* not a class type */ {
	rv := objc.Send[PersistentHistoryChangeType](p_.ID, objc.Sel("changeType"))
	return rv
}


// The type of change to the managed object in the persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/changetype
func (p_ PersistentHistoryChange) SetChangeType(value PersistentHistoryChangeType /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setChangeType:"), value)
}


// The identifier of the managed object that changed. (swift)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/changedobjectid
func (p_ PersistentHistoryChange) ChangedObjectID() IManagedObjectID {
	rv := objc.Send[ManagedObjectID](p_.ID, objc.Sel("changedObjectID"))
	return rv
}


// The identifier of the managed object that changed. (swift)
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/changedobjectid
func (p_ PersistentHistoryChange) SetChangedObjectID(value IManagedObjectID) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setChangedObjectID:"), value)
}


// The persistent history transaction containing this change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/transaction
func (p_ PersistentHistoryChange) Transaction() IPersistentHistoryTransaction {
	rv := objc.Send[PersistentHistoryTransaction](p_.ID, objc.Sel("transaction"))
	return rv
}


// The persistent history transaction containing this change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorychange/transaction
func (p_ PersistentHistoryChange) SetTransaction(value IPersistentHistoryTransaction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransaction:"), value)
}



