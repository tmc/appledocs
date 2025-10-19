// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MigrationManager] class.
var migrationManagerClass = _MigrationManagerClass{objc.GetClass("NSMigrationManager")}

type _MigrationManagerClass struct {
	class objc.Class
}

// A migration manager instance that performs a migration of data from one persistent store to another using a given mapping model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager

type MigrationManager struct {
	objectivec.Object
}

// MigrationManagerFrom constructs a [MigrationManager] from an unsafe.Pointer.
//
// A migration manager instance that performs a migration of data from one persistent store to another using a given mapping model.
func MigrationManagerFrom(ptr unsafe.Pointer) MigrationManager {
	return MigrationManager{objectivec.Object{objc.ID(ptr)}}
}

// Returns the entity description for the destination entity of a given entity mapping. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationEntity(for:)
func (m_ MigrationManager) DestinationEntityForEntityMapping(mEntity unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("destinationEntityForEntityMapping:"), mEntity)
	return rv
}


