// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ReduceBinary] class.
var (
	ReduceBinaryClass     _ReduceBinaryClass
	ReduceBinaryClassOnce sync.Once
)

func getReduceBinaryClass() _ReduceBinaryClass {
	ReduceBinaryClassOnce.Do(func() {
		ReduceBinaryClass = _ReduceBinaryClass{objc.GetClass("MPSNNReduceBinary")}
	})
	return ReduceBinaryClass
}

type _ReduceBinaryClass struct {
	class objc.Class
}

// An interface definition for the [ReduceBinary] class.
type IReduceBinary interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other MetalPerformanceShaders classes.


// A parent class referenced by other MetalPerformanceShaders classes. [Full Topic]
type ReduceBinary struct {
	objectivec.Object
}

// ReduceBinaryFrom constructs a [ReduceBinary] from an unsafe.Pointer.
//
// A parent class referenced by other MetalPerformanceShaders classes.
func ReduceBinaryFrom(ptr unsafe.Pointer) ReduceBinary {
	return ReduceBinary{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _ReduceBinaryClass) Alloc() ReduceBinary {
	rv := objc.Send[ReduceBinary](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ReduceBinaryClass) New() ReduceBinary {
	rv := objc.Send[ReduceBinary](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceBinary) Init() ReduceBinary {
	rv := objc.Send[ReduceBinary](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceBinary) Autorelease() ReduceBinary {
	rv := objc.Send[ReduceBinary](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceBinary creates a new ReduceBinary instance.
func NewReduceBinary() ReduceBinary {
	return getReduceBinaryClass().New()
}




