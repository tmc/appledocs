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
	StagedMigrationManagerClass     _StagedMigrationManagerClass
	StagedMigrationManagerClassOnce sync.Once
)

func getStagedMigrationManagerClass() _StagedMigrationManagerClass {
	StagedMigrationManagerClassOnce.Do(func() {
		StagedMigrationManagerClass = _StagedMigrationManagerClass{objc.GetClass("NSStagedMigrationManager")}
	})
	return StagedMigrationManagerClass
}

type _StagedMigrationManagerClass struct {
	class objc.Class
}

// An interface definition for the [StagedMigrationManager] class.
type IStagedMigrationManager interface {
	objectivec.IObject
	// properties:
	Container() IPersistentContainer
	Stages() []MigrationStage /* primitive/slice/pointer. */
	NSPersistentStoreStagedMigrationManagerOptionKey() string /* primitive/slice/pointer. */
	// methods:
}

// An object that handles the migration event loop and provides access to the migrating persistent store.
//
// A staged migration manager contains the individual stages of a migration and applies those stages, in the order you specify, when that migration runs. The manager handles the migration’s event loop, and provides access to the migrating store through its property. Stages can be custom, which enables you to perform tasks immediately before and after a stage runs, or lightweight, which supplements custom stages with those that Core Data can invoke automatically because they’re already compatible with lightweight migrations. Use to include an instance of in your persistent store’s options dictionary, as the following example shows:


// An object that handles the migration event loop and provides access to the migrating persistent store.
//
// [Full Topic]
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



// The container that provides access to the migrating persistent store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSStagedMigrationManager/container
func (s_ StagedMigrationManager) Container() IPersistentContainer {
	rv := objc.Send[PersistentContainer](s_.ID, objc.Sel("container"))
	return rv
}


// The migration stages.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSStagedMigrationManager/stages
func (s_ StagedMigrationManager) Stages() []MigrationStage /* primitive/slice/pointer. */ {
	rv := objc.Send[[]MigrationStage](s_.ID, objc.Sel("stages"))
	return rv
}


// The key for specifying your staged migration manager.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nspersistentstorestagedmigrationmanageroptionkey
func (s_ StagedMigrationManager) NSPersistentStoreStagedMigrationManagerOptionKey() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](s_.ID, objc.Sel("NSPersistentStoreStagedMigrationManagerOptionKey"))
	return rv
}




