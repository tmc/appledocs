// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBinding] class.
var (
	MTRBaseClusterBindingClass     _MTRBaseClusterBindingClass
	MTRBaseClusterBindingClassOnce sync.Once
)

func getMTRBaseClusterBindingClass() _MTRBaseClusterBindingClass {
	MTRBaseClusterBindingClassOnce.Do(func() {
		MTRBaseClusterBindingClass = _MTRBaseClusterBindingClass{objc.GetClass("MTRBaseClusterBinding")}
	})
	return MTRBaseClusterBindingClass
}

type _MTRBaseClusterBindingClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBinding] class.
type IMTRBaseClusterBinding interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBinding
type MTRBaseClusterBinding struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterBindingFrom constructs a [MTRBaseClusterBinding] from an unsafe.Pointer.
func MTRBaseClusterBindingFrom(ptr unsafe.Pointer) MTRBaseClusterBinding {
	return MTRBaseClusterBinding{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBindingClass) Alloc() MTRBaseClusterBinding {
	rv := objc.Send[MTRBaseClusterBinding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBindingClass) New() MTRBaseClusterBinding {
	rv := objc.Send[MTRBaseClusterBinding](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBinding) Init() MTRBaseClusterBinding {
	rv := objc.Send[MTRBaseClusterBinding](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBinding) Autorelease() MTRBaseClusterBinding {
	rv := objc.Send[MTRBaseClusterBinding](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBinding creates a new MTRBaseClusterBinding instance.
func NewMTRBaseClusterBinding() MTRBaseClusterBinding {
	return getMTRBaseClusterBindingClass().New()
}
