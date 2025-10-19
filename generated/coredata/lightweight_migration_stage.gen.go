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

// An interface definition for the [LightweightMigrationStage] class.
type ILightweightMigrationStage interface {
	IMigrationStage
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
// Alloc allocates a new instance without initialization.
func (lc _LightweightMigrationStageClass) Alloc() LightweightMigrationStage {
	rv := objc.Send[LightweightMigrationStage](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (lc _LightweightMigrationStageClass) New() LightweightMigrationStage {
	rv := objc.Send[LightweightMigrationStage](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LightweightMigrationStage) Init() LightweightMigrationStage {
	rv := objc.Send[LightweightMigrationStage](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LightweightMigrationStage) Autorelease() LightweightMigrationStage {
	rv := objc.Send[LightweightMigrationStage](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLightweightMigrationStage creates a new LightweightMigrationStage instance.
func NewLightweightMigrationStage() LightweightMigrationStage {
	return lightweightMigrationStageClass.New()
}




