// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MigrationStage] class.
var (
	MigrationStageClass     _MigrationStageClass
	MigrationStageClassOnce sync.Once
)

func getMigrationStageClass() _MigrationStageClass {
	MigrationStageClassOnce.Do(func() {
		MigrationStageClass = _MigrationStageClass{objc.GetClass("NSMigrationStage")}
	})
	return MigrationStageClass
}

type _MigrationStageClass struct {
	class objc.Class
}

// An interface definition for the [MigrationStage] class.
type IMigrationStage interface {
	objectivec.IObject
	// properties:
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	// methods:
}

// An abstract base class for describing an individual stage of a migration.


// An abstract base class for describing an individual stage of a migration.
//
// [Full Topic]
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
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationStage/label
func (m_ MigrationStage) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("label"))
	return rv
}


// The textual description of the migration stage’s purpose.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSMigrationStage/label
func (m_ MigrationStage) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setLabel:"), value)
}



