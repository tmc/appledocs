// Code generated from Apple documentation for MetricKit. DO NOT EDIT.

package metrickit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MXCallStackTree] class.
var (
	MXCallStackTreeClass     _MXCallStackTreeClass
	MXCallStackTreeClassOnce sync.Once
)

func getMXCallStackTreeClass() _MXCallStackTreeClass {
	MXCallStackTreeClassOnce.Do(func() {
		MXCallStackTreeClass = _MXCallStackTreeClass{objc.GetClass("MXCallStackTree")}
	})
	return MXCallStackTreeClass
}

type _MXCallStackTreeClass struct {
	class objc.Class
}

// An interface definition for the [MXCallStackTree] class.
type IMXCallStackTree interface {
	objectivec.IObject
	JSONRepresentation() unsafe.Pointer
}

// An object representing the call stack for an exception.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCallStackTree
type MXCallStackTree struct {
	objectivec.Object
}

// MXCallStackTreeFrom constructs a [MXCallStackTree] from an unsafe.Pointer.
//
// An object representing the call stack for an exception.
func MXCallStackTreeFrom(ptr unsafe.Pointer) MXCallStackTree {
	return MXCallStackTree{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MXCallStackTreeClass) Alloc() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MXCallStackTreeClass) New() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MXCallStackTree) Init() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MXCallStackTree) Autorelease() MXCallStackTree {
	rv := objc.Send[MXCallStackTree](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMXCallStackTree creates a new MXCallStackTree instance.
func NewMXCallStackTree() MXCallStackTree {
	return getMXCallStackTreeClass().New()
}


// Returns the contents of the stack tree in JSON format.
//
// [Full Topic]: https://developer.apple.com/documentation/MetricKit/MXCallStackTree/jsonRepresentation()
func (m_ MXCallStackTree) JSONRepresentation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("JSONRepresentation"))
	return rv
}



