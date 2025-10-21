// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [MTRBindingClusterTargetStruct] class.
var (
	MTRBindingClusterTargetStructClass     _MTRBindingClusterTargetStructClass
	MTRBindingClusterTargetStructClassOnce sync.Once
)

func getMTRBindingClusterTargetStructClass() _MTRBindingClusterTargetStructClass {
	MTRBindingClusterTargetStructClassOnce.Do(func() {
		MTRBindingClusterTargetStructClass = _MTRBindingClusterTargetStructClass{objc.GetClass("MTRBindingClusterTargetStruct")}
	})
	return MTRBindingClusterTargetStructClass
}

type _MTRBindingClusterTargetStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRBindingClusterTargetStruct] class.
type IMTRBindingClusterTargetStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBindingClusterTargetStruct
type MTRBindingClusterTargetStruct struct {
	objectivec.Object
}

// MTRBindingClusterTargetStructFrom constructs a [MTRBindingClusterTargetStruct] from an unsafe.Pointer.
func MTRBindingClusterTargetStructFrom(ptr unsafe.Pointer) MTRBindingClusterTargetStruct {
	return MTRBindingClusterTargetStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBindingClusterTargetStructClass) Alloc() MTRBindingClusterTargetStruct {
	rv := objc.Send[MTRBindingClusterTargetStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBindingClusterTargetStructClass) New() MTRBindingClusterTargetStruct {
	rv := objc.Send[MTRBindingClusterTargetStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBindingClusterTargetStruct) Init() MTRBindingClusterTargetStruct {
	rv := objc.Send[MTRBindingClusterTargetStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBindingClusterTargetStruct) Autorelease() MTRBindingClusterTargetStruct {
	rv := objc.Send[MTRBindingClusterTargetStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBindingClusterTargetStruct creates a new MTRBindingClusterTargetStruct instance.
func NewMTRBindingClusterTargetStruct() MTRBindingClusterTargetStruct {
	return getMTRBindingClusterTargetStructClass().New()
}




