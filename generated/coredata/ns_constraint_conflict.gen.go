// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ConstraintConflict] class.
var (
	ConstraintConflictClass     _ConstraintConflictClass
	ConstraintConflictClassOnce sync.Once
)

func getConstraintConflictClass() _ConstraintConflictClass {
	ConstraintConflictClassOnce.Do(func() {
		ConstraintConflictClass = _ConstraintConflictClass{objc.GetClass("NSConstraintConflict")}
	})
	return ConstraintConflictClass
}

type _ConstraintConflictClass struct {
	class objc.Class
}

// An interface definition for the [ConstraintConflict] class.
type IConstraintConflict interface {
	objectivec.IObject
}

// An encapsulation of conflicts that occur during an attempt to save a managed object.
//
// A constraint conflict occurs when your data model is using unique constraints and one or more managed objects are violating that constraint. When this error occurs, the error instance can be interrogated to determine which instance of is violating the constraint and which property on the instance is in violation.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict
type ConstraintConflict struct {
	objectivec.Object
}

// ConstraintConflictFrom constructs a [ConstraintConflict] from an unsafe.Pointer.
//
// An encapsulation of conflicts that occur during an attempt to save a managed object.
func ConstraintConflictFrom(ptr unsafe.Pointer) ConstraintConflict {
	return ConstraintConflict{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ConstraintConflictClass) Alloc() ConstraintConflict {
	rv := objc.Send[ConstraintConflict](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ConstraintConflictClass) New() ConstraintConflict {
	rv := objc.Send[ConstraintConflict](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ConstraintConflict) Init() ConstraintConflict {
	rv := objc.Send[ConstraintConflict](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ConstraintConflict) Autorelease() ConstraintConflict {
	rv := objc.Send[ConstraintConflict](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewConstraintConflict creates a new ConstraintConflict instance.
func NewConstraintConflict() ConstraintConflict {
	return getConstraintConflictClass().New()
}


// Initializes a constraint conflict.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/init(constraint:database:databaseSnapshot:conflicting:conflictingSnapshots:)
func NewConstraintConflictWithConstraintDatabaseObjectDatabaseSnapshotConflictingObjectsConflictingSnapshots(contraint unsafe.Pointer, databaseObject unsafe.Pointer, databaseSnapshot objc.ID, conflictingObjects unsafe.Pointer, conflictingSnapshots objc.ID) ConstraintConflict {
	instance := getConstraintConflictClass().Alloc()
	rv := objc.Send[ConstraintConflict](instance.ID, objc.Sel("initWithConstraint:databaseObject:databaseSnapshot:conflictingObjects:conflictingSnapshots:"), contraint, databaseObject, databaseSnapshot, conflictingObjects, conflictingSnapshots)
	rv.Autorelease()
	return rv
}


// The managed objects that are in conflict.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/conflictingObjects
func (c_ ConstraintConflict) ConflictingObjects() []ManagedObject {
	rv := objc.Send[[]ManagedObject](c_.ID, objc.Sel("conflictingObjects"))
	return rv
}

// The original property values of objects in violation of the constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/conflictingSnapshots
func (c_ ConstraintConflict) ConflictingSnapshots() []unsafe.Pointer {
	rv := objc.Send[[]unsafe.Pointer](c_.ID, objc.Sel("conflictingSnapshots"))
	return rv
}

// The constraint that has been violated.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/constraint
func (c_ ConstraintConflict) Constraint() []string {
	rv := objc.Send[[]string](c_.ID, objc.Sel("constraint"))
	return rv
}

// The values that the conflicting objects had when the conflict was created.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/constraintValues
func (c_ ConstraintConflict) ConstraintValues() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("constraintValues"))
	return rv
}

// The object whose database row is using constraint values.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/databaseObject
func (c_ ConstraintConflict) DatabaseObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("databaseObject"))
	return rv
}

// The values currently stored in the database.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/databaseSnapshot
func (c_ ConstraintConflict) DatabaseSnapshot() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("databaseSnapshot"))
	return rv
}


