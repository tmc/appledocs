// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBinding] class.
var (
	MTRClusterBindingClass     _MTRClusterBindingClass
	MTRClusterBindingClassOnce sync.Once
)

func getMTRClusterBindingClass() _MTRClusterBindingClass {
	MTRClusterBindingClassOnce.Do(func() {
		MTRClusterBindingClass = _MTRClusterBindingClass{objc.GetClass("MTRClusterBinding")}
	})
	return MTRClusterBindingClass
}

type _MTRClusterBindingClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBinding] class.
type IMTRClusterBinding interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBinding
type MTRClusterBinding struct {
	MTRGenericCluster
}

// MTRClusterBindingFrom constructs a [MTRClusterBinding] from an unsafe.Pointer.
func MTRClusterBindingFrom(ptr unsafe.Pointer) MTRClusterBinding {
	return MTRClusterBinding{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBindingClass) Alloc() MTRClusterBinding {
	rv := objc.Send[MTRClusterBinding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBindingClass) New() MTRClusterBinding {
	rv := objc.Send[MTRClusterBinding](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBinding) Init() MTRClusterBinding {
	rv := objc.Send[MTRClusterBinding](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBinding) Autorelease() MTRClusterBinding {
	rv := objc.Send[MTRClusterBinding](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBinding creates a new MTRClusterBinding instance.
func NewMTRClusterBinding() MTRClusterBinding {
	return getMTRClusterBindingClass().New()
}
