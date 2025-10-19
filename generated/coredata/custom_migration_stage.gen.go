// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CustomMigrationStage] class.
var (
	customMigrationStageClass     _CustomMigrationStageClass
	customMigrationStageClassOnce sync.Once
)

func getCustomMigrationStageClass() _CustomMigrationStageClass {
	customMigrationStageClassOnce.Do(func() {
		customMigrationStageClass = _CustomMigrationStageClass{objc.GetClass("NSCustomMigrationStage")}
	})
	return customMigrationStageClass
}

type _CustomMigrationStageClass struct {
	class objc.Class
}

// An interface definition for the [CustomMigrationStage] class.
type ICustomMigrationStage interface {
	IMigrationStage
}

// An object that enables you to participate in the migration between two versions of the same model.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage
type CustomMigrationStage struct {
	MigrationStage
}

// CustomMigrationStageFrom constructs a [CustomMigrationStage] from an unsafe.Pointer.
//
// An object that enables you to participate in the migration between two versions of the same model.
func CustomMigrationStageFrom(ptr unsafe.Pointer) CustomMigrationStage {
	return CustomMigrationStage{
		MigrationStage: MigrationStageFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CustomMigrationStageClass) Alloc() CustomMigrationStage {
	rv := objc.Send[CustomMigrationStage](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CustomMigrationStageClass) New() CustomMigrationStage {
	rv := objc.Send[CustomMigrationStage](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CustomMigrationStage) Init() CustomMigrationStage {
	rv := objc.Send[CustomMigrationStage](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CustomMigrationStage) Autorelease() CustomMigrationStage {
	rv := objc.Send[CustomMigrationStage](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCustomMigrationStage creates a new CustomMigrationStage instance.
func NewCustomMigrationStage() CustomMigrationStage {
	return getCustomMigrationStageClass().New()
}


// Creates a custom migration stage with the specified source and destination model references.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSCustomMigrationStage/initWithCurrentModelReference:nextModelReference:
func NewCustomMigrationStageWithCurrentModelReferenceNextModelReference(currentModel unsafe.Pointer, nextModel unsafe.Pointer) CustomMigrationStage {
	instance := getCustomMigrationStageClass().Alloc()
	rv := objc.Send[CustomMigrationStage](instance.ID, objc.Sel("initWithCurrentModelReference:nextModelReference:"), currentModel, nextModel)
	rv.Autorelease()
	return rv
}



