// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentHistoryTransaction] class.
var (
	PersistentHistoryTransactionClass     _PersistentHistoryTransactionClass
	PersistentHistoryTransactionClassOnce sync.Once
)

func getPersistentHistoryTransactionClass() _PersistentHistoryTransactionClass {
	PersistentHistoryTransactionClassOnce.Do(func() {
		PersistentHistoryTransactionClass = _PersistentHistoryTransactionClass{objc.GetClass("NSPersistentHistoryTransaction")}
	})
	return PersistentHistoryTransactionClass
}

type _PersistentHistoryTransactionClass struct {
	class objc.Class
}

// An interface definition for the [PersistentHistoryTransaction] class.
type IPersistentHistoryTransaction interface {
	objectivec.IObject
	ObjectIDNotification() foundation.Notification
	Author() string
	BundleID() string
	Changes() []PersistentHistoryChange
	ContextName() string
	ProcessID() string
	StoreID() string
	Timestamp() foundation.NSDate
	Token() NSPersistentHistoryToken
	TransactionNumber() unsafe.Pointer
}

// A set of changes in the persistent history based on a context save or batch operation.


// A set of changes in the persistent history based on a context save or batch operation.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/entityDescription(with:)
func (pc _PersistentHistoryTransactionClass) EntityDescriptionWithContext(context IManagedObjectContext) EntityDescription {
	rv := objc.Send[EntityDescription](objc.ID(pc.class), objc.Sel("entityDescriptionWithContext:"), context)
	return rv
}


// The entity description of the persistent history transaction entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/entityDescription
func (pc _PersistentHistoryTransactionClass) EntityDescription() NSEntityDescription {
	rv := objc.Send[NSEntityDescription](objc.ID(pc.class), objc.Sel("entityDescription"))
	return rv
}

// A fetch request that has the persistent history transaction as the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/fetchRequest
func (pc _PersistentHistoryTransactionClass) FetchRequest() NSFetchRequest {
	rv := objc.Send[NSFetchRequest](objc.ID(pc.class), objc.Sel("fetchRequest"))
	return rv
}

// Obtains a notification for use in merging the transaction’s changes into a managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/objectIDNotification()
func (p_ PersistentHistoryTransaction) ObjectIDNotification() foundation.Notification {
	rv := objc.Send[foundation.Notification](p_.ID, objc.Sel("objectIDNotification"))
	return rv
}


// A granular description of the context that made the persistent history change, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/author
func (p_ PersistentHistoryTransaction) Author() string {
	rv := objc.Send[string](p_.ID, objc.Sel("author"))
	return rv
}


// The originating bundle’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/bundleID
func (p_ PersistentHistoryTransaction) BundleID() string {
	rv := objc.Send[string](p_.ID, objc.Sel("bundleID"))
	return rv
}


// The array of persistent history changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/changes
func (p_ PersistentHistoryTransaction) Changes() []PersistentHistoryChange {
	rv := objc.Send[[]PersistentHistoryChange](p_.ID, objc.Sel("changes"))
	return rv
}


// The originating context’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/contextName
func (p_ PersistentHistoryTransaction) ContextName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("contextName"))
	return rv
}


// The entity description of the persistent history transaction entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/entityDescription
func (p_ PersistentHistoryTransaction) EntityDescription() NSEntityDescription {
	rv := objc.Send[NSEntityDescription](p_.ID, objc.Sel("entityDescription"))
	return rv
}


// A fetch request that has the persistent history transaction as the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/fetchRequest
func (p_ PersistentHistoryTransaction) FetchRequest() NSFetchRequest {
	rv := objc.Send[NSFetchRequest](p_.ID, objc.Sel("fetchRequest"))
	return rv
}


// The originating process’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/processID
func (p_ PersistentHistoryTransaction) ProcessID() string {
	rv := objc.Send[string](p_.ID, objc.Sel("processID"))
	return rv
}


// The originating store’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/storeID
func (p_ PersistentHistoryTransaction) StoreID() string {
	rv := objc.Send[string](p_.ID, objc.Sel("storeID"))
	return rv
}


// The date of the persistent history change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/timestamp
func (p_ PersistentHistoryTransaction) Timestamp() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("timestamp"))
	return rv
}


// The token that represents this transaction in the persistent history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/token
func (p_ PersistentHistoryTransaction) Token() NSPersistentHistoryToken {
	rv := objc.Send[NSPersistentHistoryToken](p_.ID, objc.Sel("token"))
	return rv
}


// The transaction’s numeric identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/transactionNumber
func (p_ PersistentHistoryTransaction) TransactionNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("transactionNumber"))
	return rv
}



