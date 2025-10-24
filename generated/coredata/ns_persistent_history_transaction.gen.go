// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
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
	// properties:
	Author() objc.IObject   /* cross-framework: NSString */
	BundleID() objc.IObject /* cross-framework: NSString */
	Changes() []IPersistentHistoryChange
	ContextName() objc.IObject /* cross-framework: NSString */
	ProcessID() objc.IObject   /* cross-framework: NSString */
	StoreID() objc.IObject     /* cross-framework: NSString */
	Timestamp() objc.IObject   /* cross-framework: Date */
	SetTimestamp(value objc.IObject /* cross-framework: Date */)
	Token() IPersistentHistoryToken
	SetToken(value IPersistentHistoryToken)
	TransactionNumber() unsafe.Pointer
	SetTransactionNumber(value unsafe.Pointer)
	// methods:
	ObjectIDNotification() objc.IObject /* cross-framework: Notification */
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
func (pc _PersistentHistoryTransactionClass) EntityDescriptionWithContext(context IManagedObjectContext) IEntityDescription {
	rv := objc.Send[EntityDescription](objc.ID(pc.class), objc.Sel("entityDescriptionWithContext:"), context)
	return rv
}

// The entity description of the persistent history transaction entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/entityDescription
func (pc _PersistentHistoryTransactionClass) EntityDescription() IEntityDescription {
	rv := objc.Send[EntityDescription](objc.ID(pc.class), objc.Sel("entityDescription"))
	return rv
}

// A fetch request that has the persistent history transaction as the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/fetchRequest
func (pc _PersistentHistoryTransactionClass) FetchRequest() IFetchRequest {
	rv := objc.Send[FetchRequest](objc.ID(pc.class), objc.Sel("fetchRequest"))
	return rv
}

// Obtains a notification for use in merging the transaction’s changes into a managed object context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/objectIDNotification()
func (p_ PersistentHistoryTransaction) ObjectIDNotification() objc.IObject /* cross-framework: Notification */ {
	rv := objc.Send[foundation.Notification](p_.ID, objc.Sel("objectIDNotification"))
	return rv
}

// A granular description of the context that made the persistent history change, if available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/author
func (p_ PersistentHistoryTransaction) Author() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("author"))
	return rv
}

// The originating bundle’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/bundleID
func (p_ PersistentHistoryTransaction) BundleID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("bundleID"))
	return rv
}

// The array of persistent history changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/changes
func (p_ PersistentHistoryTransaction) Changes() []IPersistentHistoryChange {
	rv := objc.Send[[]PersistentHistoryChange](p_.ID, objc.Sel("changes"))
	return rv
}

// The originating context’s name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/contextName
func (p_ PersistentHistoryTransaction) ContextName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("contextName"))
	return rv
}

// The entity description of the persistent history transaction entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/entityDescription
func (p_ PersistentHistoryTransaction) EntityDescription() IEntityDescription {
	rv := objc.Send[EntityDescription](p_.ID, objc.Sel("entityDescription"))
	return rv
}

// A fetch request that has the persistent history transaction as the entity.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/fetchRequest
func (p_ PersistentHistoryTransaction) FetchRequest() IFetchRequest {
	rv := objc.Send[FetchRequest](p_.ID, objc.Sel("fetchRequest"))
	return rv
}

// The originating process’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/processID
func (p_ PersistentHistoryTransaction) ProcessID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("processID"))
	return rv
}

// The originating store’s identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/storeID
func (p_ PersistentHistoryTransaction) StoreID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("storeID"))
	return rv
}

// The date of the persistent history change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/timestamp
func (p_ PersistentHistoryTransaction) Timestamp() objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](p_.ID, objc.Sel("timestamp"))
	return rv
}

// The date of the persistent history change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/timestamp
func (p_ PersistentHistoryTransaction) SetTimestamp(value objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTimestamp:"), value)
}

// The token that represents this transaction in the persistent history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/token
func (p_ PersistentHistoryTransaction) Token() IPersistentHistoryToken {
	rv := objc.Send[PersistentHistoryToken](p_.ID, objc.Sel("token"))
	return rv
}

// The token that represents this transaction in the persistent history.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/token
func (p_ PersistentHistoryTransaction) SetToken(value IPersistentHistoryToken) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setToken:"), value)
}

// The transaction’s numeric identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/transactionnumber
func (p_ PersistentHistoryTransaction) TransactionNumber() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("transactionNumber"))
	return rv
}

// The transaction’s numeric identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistenthistorytransaction/transactionnumber
func (p_ PersistentHistoryTransaction) SetTransactionNumber(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransactionNumber:"), value)
}
