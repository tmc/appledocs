// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MigrationStage] class.
var migrationStageClass = _MigrationStageClass{objc.GetClass("NSMigrationStage")}

type _MigrationStageClass struct {
	class objc.Class
}

// An abstract base class for describing an individual stage of a migration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationStage

type MigrationStage struct {
	objectivec.Object
}

// MigrationStageFrom constructs a [MigrationStage] from an unsafe.Pointer.
//
// An abstract base class for describing an individual stage of a migration.
func MigrationStageFrom(ptr unsafe.Pointer) MigrationStage {
	return MigrationStage{objectivec.Object{objc.ID(ptr)}}
}



