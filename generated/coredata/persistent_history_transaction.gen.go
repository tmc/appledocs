// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentHistoryTransaction] class.
var (
	persistentHistoryTransactionClass     _PersistentHistoryTransactionClass
	persistentHistoryTransactionClassOnce sync.Once
)

func getPersistentHistoryTransactionClass() _PersistentHistoryTransactionClass {
	persistentHistoryTransactionClassOnce.Do(func() {
		persistentHistoryTransactionClass = _PersistentHistoryTransactionClass{objc.GetClass("NSPersistentHistoryTransaction")}
	})
	return persistentHistoryTransactionClass
}

type _PersistentHistoryTransactionClass struct {
	class objc.Class
}

// An interface definition for the [PersistentHistoryTransaction] class.
type IPersistentHistoryTransaction interface {
	objectivec.IObject
	ObjectIDNotification() unsafe.Pointer
}

// A set of changes in the persistent history based on a context save or batch operation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction
type PersistentHistoryTransaction struct {
	objectivec.Object
}

// PersistentHistoryTransactionFrom constructs a [PersistentHistoryTransaction] from an unsafe.Pointer.
//
// A set of changes in the persistent history based on a context save or batch operation.
func PersistentHistoryTransactionFrom(ptr unsafe.Pointer) PersistentHistoryTransaction {
	return PersistentHistoryTransaction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PersistentHistoryTransactionClass) Alloc() PersistentHistoryTransaction {
	rv := objc.Send[PersistentHistoryTransaction](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PersistentHistoryTransactionClass) New() PersistentHistoryTransaction {
	rv := objc.Send[PersistentHistoryTransaction](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PersistentHistoryTransaction) Init() PersistentHistoryTransaction {
	rv := objc.Send[PersistentHistoryTransaction](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PersistentHistoryTransaction) Autorelease() PersistentHistoryTransaction {
	rv := objc.Send[PersistentHistoryTransaction](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPersistentHistoryTransaction creates a new PersistentHistoryTransaction instance.
func NewPersistentHistoryTransaction() PersistentHistoryTransaction {
	return getPersistentHistoryTransactionClass().New()
}


// Requests an entity description using the provided context for the managed object type affected by the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/entityDescription(with:)
func (pc _PersistentHistoryTransactionClass) EntityDescriptionWithContext(context unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("entityDescriptionWithContext:"), context)
	return rv
}
// Obtains a notification for use in merging the transaction’s changes into a managed object context.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/objectIDNotification()
func (p_ PersistentHistoryTransaction) ObjectIDNotification() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("objectIDNotification"))
	return rv
}


