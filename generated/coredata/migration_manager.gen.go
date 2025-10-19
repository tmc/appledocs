// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MigrationManager] class.
var (
	migrationManagerClass     _MigrationManagerClass
	migrationManagerClassOnce sync.Once
)

func getMigrationManagerClass() _MigrationManagerClass {
	migrationManagerClassOnce.Do(func() {
		migrationManagerClass = _MigrationManagerClass{objc.GetClass("NSMigrationManager")}
	})
	return migrationManagerClass
}

type _MigrationManagerClass struct {
	class objc.Class
}

// An interface definition for the [MigrationManager] class.
type IMigrationManager interface {
	objectivec.IObject
	DestinationEntityForEntityMapping(mEntity unsafe.Pointer) unsafe.Pointer
}

// A migration manager instance that performs a migration of data from one persistent store to another using a given mapping model.
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

// Alloc allocates a new instance without initialization.
func (mc _MigrationManagerClass) Alloc() MigrationManager {
	rv := objc.Send[MigrationManager](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MigrationManagerClass) New() MigrationManager {
	rv := objc.Send[MigrationManager](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MigrationManager) Init() MigrationManager {
	rv := objc.Send[MigrationManager](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MigrationManager) Autorelease() MigrationManager {
	rv := objc.Send[MigrationManager](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMigrationManager creates a new MigrationManager instance.
func NewMigrationManager() MigrationManager {
	return getMigrationManagerClass().New()
}


// Returns the entity description for the destination entity of a given entity mapping.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationManager/destinationEntity(for:)
func (m_ MigrationManager) DestinationEntityForEntityMapping(mEntity unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("destinationEntityForEntityMapping:"), mEntity)
	return rv
}


