// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PersistentHistoryTransaction] class.
var persistentHistoryTransactionClass = _PersistentHistoryTransactionClass{objc.GetClass("NSPersistentHistoryTransaction")}

type _PersistentHistoryTransactionClass struct {
	class objc.Class
}

// A set of changes in the persistent history based on a context save or batch operation. [Full Topic]
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

// Requests an entity description using the provided context for the managed object type affected by the transaction. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/entityDescription(with:)
func (pc _PersistentHistoryTransactionClass) EntityDescriptionWithContext(context unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("entityDescriptionWithContext:"), context)
	return rv
}
// Obtains a notification for use in merging the transaction’s changes into a managed object context. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryTransaction/objectIDNotification()
func (p_ PersistentHistoryTransaction) ObjectIDNotification() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("objectIDNotification"))
	return rv
}


