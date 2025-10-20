// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MigrationStage] class.
var (
	migrationStageClass     _MigrationStageClass
	migrationStageClassOnce sync.Once
)

func getMigrationStageClass() _MigrationStageClass {
	migrationStageClassOnce.Do(func() {
		migrationStageClass = _MigrationStageClass{objc.GetClass("NSMigrationStage")}
	})
	return migrationStageClass
}

type _MigrationStageClass struct {
	class objc.Class
}

// An interface definition for the [MigrationStage] class.
type IMigrationStage interface {
	objectivec.IObject
}

// An abstract base class for describing an individual stage of a migration.
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

// Alloc allocates a new instance without initialization.
func (mc _MigrationStageClass) Alloc() MigrationStage {
	rv := objc.Send[MigrationStage](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MigrationStageClass) New() MigrationStage {
	rv := objc.Send[MigrationStage](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MigrationStage) Init() MigrationStage {
	rv := objc.Send[MigrationStage](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MigrationStage) Autorelease() MigrationStage {
	rv := objc.Send[MigrationStage](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMigrationStage creates a new MigrationStage instance.
func NewMigrationStage() MigrationStage {
	return getMigrationStageClass().New()
}


// The textual description of the migration stage’s purpose.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationStage/label
func (m_ MigrationStage) Label() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("label"))
	return rv
}

// SetLabel sets the value of the label property.
// The textual description of the migration stage’s purpose.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationStage/label
func (m_ MigrationStage) SetLabel(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}


