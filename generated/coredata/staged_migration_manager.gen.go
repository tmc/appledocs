// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [StagedMigrationManager] class.
var stagedMigrationManagerClass = _StagedMigrationManagerClass{objc.GetClass("NSStagedMigrationManager")}

type _StagedMigrationManagerClass struct {
	class objc.Class
}

// An object that handles the migration event loop and provides access to the migrating persistent store. [Full Topic]
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



