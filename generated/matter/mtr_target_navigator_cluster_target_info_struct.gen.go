// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTRTargetNavigatorClusterTargetInfoStruct] class.
var (
	MTRTargetNavigatorClusterTargetInfoStructClass     _MTRTargetNavigatorClusterTargetInfoStructClass
	MTRTargetNavigatorClusterTargetInfoStructClassOnce sync.Once
)

func getMTRTargetNavigatorClusterTargetInfoStructClass() _MTRTargetNavigatorClusterTargetInfoStructClass {
	MTRTargetNavigatorClusterTargetInfoStructClassOnce.Do(func() {
		MTRTargetNavigatorClusterTargetInfoStructClass = _MTRTargetNavigatorClusterTargetInfoStructClass{objc.GetClass("MTRTargetNavigatorClusterTargetInfoStruct")}
	})
	return MTRTargetNavigatorClusterTargetInfoStructClass
}

type _MTRTargetNavigatorClusterTargetInfoStructClass struct {
	class objc.Class
}

// An interface definition for the [MTRTargetNavigatorClusterTargetInfoStruct] class.
type IMTRTargetNavigatorClusterTargetInfoStruct interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRTargetNavigatorClusterTargetInfoStruct
type MTRTargetNavigatorClusterTargetInfoStruct struct {
	objectivec.Object
}

// MTRTargetNavigatorClusterTargetInfoStructFrom constructs a [MTRTargetNavigatorClusterTargetInfoStruct] from an unsafe.Pointer.
func MTRTargetNavigatorClusterTargetInfoStructFrom(ptr unsafe.Pointer) MTRTargetNavigatorClusterTargetInfoStruct {
	return MTRTargetNavigatorClusterTargetInfoStruct{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRTargetNavigatorClusterTargetInfoStructClass) Alloc() MTRTargetNavigatorClusterTargetInfoStruct {
	rv := objc.Send[MTRTargetNavigatorClusterTargetInfoStruct](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRTargetNavigatorClusterTargetInfoStructClass) New() MTRTargetNavigatorClusterTargetInfoStruct {
	rv := objc.Send[MTRTargetNavigatorClusterTargetInfoStruct](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRTargetNavigatorClusterTargetInfoStruct) Init() MTRTargetNavigatorClusterTargetInfoStruct {
	rv := objc.Send[MTRTargetNavigatorClusterTargetInfoStruct](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRTargetNavigatorClusterTargetInfoStruct) Autorelease() MTRTargetNavigatorClusterTargetInfoStruct {
	rv := objc.Send[MTRTargetNavigatorClusterTargetInfoStruct](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRTargetNavigatorClusterTargetInfoStruct creates a new MTRTargetNavigatorClusterTargetInfoStruct instance.
func NewMTRTargetNavigatorClusterTargetInfoStruct() MTRTargetNavigatorClusterTargetInfoStruct {
	return getMTRTargetNavigatorClusterTargetInfoStructClass().New()
}




