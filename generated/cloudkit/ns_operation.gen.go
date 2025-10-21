// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Operation] class.
var (
	OperationClass     _OperationClass
	OperationClassOnce sync.Once
)

func getOperationClass() _OperationClass {
	OperationClassOnce.Do(func() {
		OperationClass = _OperationClass{objc.GetClass("NSOperation")}
	})
	return OperationClass
}

type _OperationClass struct {
	class objc.Class
}

// An interface definition for the [Operation] class.
type IOperation interface {
	objectivec.IObject
}

// A parent class referenced by other CloudKit classes.
type Operation struct {
	objectivec.Object
}

// OperationFrom constructs a [Operation] from an unsafe.Pointer.
//
// A parent class referenced by other CloudKit classes.
func OperationFrom(ptr unsafe.Pointer) Operation {
	return Operation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OperationClass) Alloc() Operation {
	rv := objc.Send[Operation](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OperationClass) New() Operation {
	rv := objc.Send[Operation](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Operation) Init() Operation {
	rv := objc.Send[Operation](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Operation) Autorelease() Operation {
	rv := objc.Send[Operation](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOperation creates a new Operation instance.
func NewOperation() Operation {
	return getOperationClass().New()
}




