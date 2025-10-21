// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterMediaInput] class.
var (
	MTRClusterMediaInputClass     _MTRClusterMediaInputClass
	MTRClusterMediaInputClassOnce sync.Once
)

func getMTRClusterMediaInputClass() _MTRClusterMediaInputClass {
	MTRClusterMediaInputClassOnce.Do(func() {
		MTRClusterMediaInputClass = _MTRClusterMediaInputClass{objc.GetClass("MTRClusterMediaInput")}
	})
	return MTRClusterMediaInputClass
}

type _MTRClusterMediaInputClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterMediaInput] class.
type IMTRClusterMediaInput interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterMediaInput
type MTRClusterMediaInput struct {
	MTRGenericCluster
}

// MTRClusterMediaInputFrom constructs a [MTRClusterMediaInput] from an unsafe.Pointer.
func MTRClusterMediaInputFrom(ptr unsafe.Pointer) MTRClusterMediaInput {
	return MTRClusterMediaInput{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterMediaInputClass) Alloc() MTRClusterMediaInput {
	rv := objc.Send[MTRClusterMediaInput](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterMediaInputClass) New() MTRClusterMediaInput {
	rv := objc.Send[MTRClusterMediaInput](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterMediaInput) Init() MTRClusterMediaInput {
	rv := objc.Send[MTRClusterMediaInput](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterMediaInput) Autorelease() MTRClusterMediaInput {
	rv := objc.Send[MTRClusterMediaInput](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterMediaInput creates a new MTRClusterMediaInput instance.
func NewMTRClusterMediaInput() MTRClusterMediaInput {
	return getMTRClusterMediaInputClass().New()
}




