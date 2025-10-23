// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	DatabaseSnapshot() foundation.IDictionary
	ConflictingObjects() IManagedObject
	SetConflictingObjects(value IManagedObject)
	ConflictingSnapshots() unsafe.Pointer
	SetConflictingSnapshots(value unsafe.Pointer)
	Constraint() string
	SetConstraint(value string)
	ConstraintValues() string
	SetConstraintValues(value string)
	DatabaseObject() IManagedObject
	SetDatabaseObject(value IManagedObject)
}

// An encapsulation of conflicts that occur during an attempt to save a managed object.
//
// A constraint conflict occurs when your data model is using unique constraints and one or more managed objects are violating that constraint. When this error occurs, the error instance can be interrogated to determine which instance of is violating the constraint and which property on the instance is in violation.


// An encapsulation of conflicts that occur during an attempt to save a managed object.
//
// [Full Topic]
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



// The values currently stored in the database.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreData/NSConstraintConflict/databaseSnapshot
func (c_ ConstraintConflict) DatabaseSnapshot() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("databaseSnapshot"))
	return rv
}


// The managed objects that are in conflict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/conflictingobjects
func (c_ ConstraintConflict) ConflictingObjects() IManagedObject {
	rv := objc.Send[ManagedObject](c_.ID, objc.Sel("conflictingObjects"))
	return rv
}


// The managed objects that are in conflict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/conflictingobjects
func (c_ ConstraintConflict) SetConflictingObjects(value IManagedObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConflictingObjects:"), value)
}


// The original property values of objects in violation of the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/conflictingsnapshots
func (c_ ConstraintConflict) ConflictingSnapshots() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("conflictingSnapshots"))
	return rv
}


// The original property values of objects in violation of the constraint.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/conflictingsnapshots
func (c_ ConstraintConflict) SetConflictingSnapshots(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConflictingSnapshots:"), value)
}


// The constraint that has been violated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/constraint
func (c_ ConstraintConflict) Constraint() string {
	rv := objc.Send[string](c_.ID, objc.Sel("constraint"))
	return rv
}


// The constraint that has been violated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/constraint
func (c_ ConstraintConflict) SetConstraint(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstraint:"), objc.String(value))
}


// The values that the conflicting objects had when the conflict was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/constraintvalues
func (c_ ConstraintConflict) ConstraintValues() string {
	rv := objc.Send[string](c_.ID, objc.Sel("constraintValues"))
	return rv
}


// The values that the conflicting objects had when the conflict was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/constraintvalues
func (c_ ConstraintConflict) SetConstraintValues(value string) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConstraintValues:"), objc.String(value))
}


// The object whose database row is using constraint values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/databaseobject
func (c_ ConstraintConflict) DatabaseObject() IManagedObject {
	rv := objc.Send[ManagedObject](c_.ID, objc.Sel("databaseObject"))
	return rv
}


// The object whose database row is using constraint values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsconstraintconflict/databaseobject
func (c_ ConstraintConflict) SetDatabaseObject(value IManagedObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDatabaseObject:"), value)
}



