// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StagedMigrationManager] class.
var (
	stagedMigrationManagerClass     _StagedMigrationManagerClass
	stagedMigrationManagerClassOnce sync.Once
)

func getStagedMigrationManagerClass() _StagedMigrationManagerClass {
	stagedMigrationManagerClassOnce.Do(func() {
		stagedMigrationManagerClass = _StagedMigrationManagerClass{objc.GetClass("NSStagedMigrationManager")}
	})
	return stagedMigrationManagerClass
}

type _StagedMigrationManagerClass struct {
	class objc.Class
}

// An interface definition for the [StagedMigrationManager] class.
type IStagedMigrationManager interface {
	objectivec.IObject
}

// An object that handles the migration event loop and provides access to the migrating persistent store.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSStagedMigrationManager
type StagedMigrationManager struct {
	objectivec.Object
}

// StagedMigrationManagerFrom constructs a [StagedMigrationManager] from an unsafe.Pointer.
//
// An object that handles the migration event loop and provides access to the migrating persistent store.
func StagedMigrationManagerFrom(ptr unsafe.Pointer) StagedMigrationManager {
	return StagedMigrationManager{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StagedMigrationManagerClass) Alloc() StagedMigrationManager {
	rv := objc.Send[StagedMigrationManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StagedMigrationManagerClass) New() StagedMigrationManager {
	rv := objc.Send[StagedMigrationManager](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ StagedMigrationManager) Init() StagedMigrationManager {
	rv := objc.Send[StagedMigrationManager](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ StagedMigrationManager) Autorelease() StagedMigrationManager {
	rv := objc.Send[StagedMigrationManager](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStagedMigrationManager creates a new StagedMigrationManager instance.
func NewStagedMigrationManager() StagedMigrationManager {
	return getStagedMigrationManagerClass().New()
}




