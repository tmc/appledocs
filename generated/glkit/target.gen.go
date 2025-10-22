// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [target] class.
var (
	TargetClass     _targetClass
	TargetClassOnce sync.Once
)

func gettargetClass() _targetClass {
	TargetClassOnce.Do(func() {
		TargetClass = _targetClass{objc.GetClass("target")}
	})
	return TargetClass
}

type _targetClass struct {
	class objc.Class
}

// An interface definition for the [target] class.
type Itarget interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/target-c.ivar

type target struct {
	objectivec.Object
}

// targetFrom constructs a [target] from an unsafe.Pointer.
func targetFrom(ptr unsafe.Pointer) target {
	return target{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (tc _targetClass) Alloc() target {
	rv := objc.Send[target](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _targetClass) New() target {
	rv := objc.Send[target](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ target) Init() target {
	rv := objc.Send[target](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ target) Autorelease() target {
	rv := objc.Send[target](t_.ID, objc.Sel("autorelease"))
	return rv
}

// Newtarget creates a new target instance.
func Newtarget() target {
	return gettargetClass().New()
}




