// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LightweightMigrationStage] class.
var lightweightMigrationStageClass = _LightweightMigrationStageClass{objc.GetClass("NSLightweightMigrationStage")}

type _LightweightMigrationStageClass struct {
	class objc.Class
}

// An object that describes a series of models suitable for lightweight migration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSLightweightMigrationStage

type LightweightMigrationStage struct {
	MigrationStage
}

// LightweightMigrationStageFrom constructs a [LightweightMigrationStage] from an unsafe.Pointer.
//
// An object that describes a series of models suitable for lightweight migration.
func LightweightMigrationStageFrom(ptr unsafe.Pointer) LightweightMigrationStage {
	return LightweightMigrationStage{
		MigrationStage: MigrationStageFrom(ptr),
	}
}



