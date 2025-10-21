// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LightweightMigrationStage] class.
var (
	LightweightMigrationStageClass     _LightweightMigrationStageClass
	LightweightMigrationStageClassOnce sync.Once
)

func getLightweightMigrationStageClass() _LightweightMigrationStageClass {
	LightweightMigrationStageClassOnce.Do(func() {
		LightweightMigrationStageClass = _LightweightMigrationStageClass{objc.GetClass("NSLightweightMigrationStage")}
	})
	return LightweightMigrationStageClass
}

type _LightweightMigrationStageClass struct {
	class objc.Class
}

// An interface definition for the [LightweightMigrationStage] class.
type ILightweightMigrationStage interface {
	IMigrationStage
}

// An object that describes a series of models suitable for lightweight migration.
//
// Use when you have a series of models to migrate and those models are compatible with lightweight migrations. Instances of this class supplement your custom migration stages and help maintain a consistent stage order for the entire migration.
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getLightweightMigrationStageClass().New()
}




// Creates a lightweight migration stage with the specified version checksums.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSLightweightMigrationStage/initWithVersionChecksums:
func NewLightweightMigrationStageWithVersionChecksums(versionChecksums unsafe.Pointer) LightweightMigrationStage {
	instance := getLightweightMigrationStageClass().Alloc()
	rv := objc.Send[LightweightMigrationStage](instance.ID, objc.Sel("initWithVersionChecksums:"), versionChecksums)
	rv.Autorelease()
	return rv
}


// The array of version checksums.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSLightweightMigrationStage/versionChecksums
func (l_ LightweightMigrationStage) VersionChecksums() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("versionChecksums"))
	return rv
}


