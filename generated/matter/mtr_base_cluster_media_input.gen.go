// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterMediaInput] class.
var (
	MTRBaseClusterMediaInputClass     _MTRBaseClusterMediaInputClass
	MTRBaseClusterMediaInputClassOnce sync.Once
)

func getMTRBaseClusterMediaInputClass() _MTRBaseClusterMediaInputClass {
	MTRBaseClusterMediaInputClassOnce.Do(func() {
		MTRBaseClusterMediaInputClass = _MTRBaseClusterMediaInputClass{objc.GetClass("MTRBaseClusterMediaInput")}
	})
	return MTRBaseClusterMediaInputClass
}

type _MTRBaseClusterMediaInputClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterMediaInput] class.
type IMTRBaseClusterMediaInput interface {
	IMTRGenericBaseCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterMediaInput
type MTRBaseClusterMediaInput struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterMediaInputFrom constructs a [MTRBaseClusterMediaInput] from an unsafe.Pointer.
func MTRBaseClusterMediaInputFrom(ptr unsafe.Pointer) MTRBaseClusterMediaInput {
	return MTRBaseClusterMediaInput{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterMediaInputClass) Alloc() MTRBaseClusterMediaInput {
	rv := objc.Send[MTRBaseClusterMediaInput](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterMediaInputClass) New() MTRBaseClusterMediaInput {
	rv := objc.Send[MTRBaseClusterMediaInput](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterMediaInput) Init() MTRBaseClusterMediaInput {
	rv := objc.Send[MTRBaseClusterMediaInput](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterMediaInput) Autorelease() MTRBaseClusterMediaInput {
	rv := objc.Send[MTRBaseClusterMediaInput](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterMediaInput creates a new MTRBaseClusterMediaInput instance.
func NewMTRBaseClusterMediaInput() MTRBaseClusterMediaInput {
	return getMTRBaseClusterMediaInputClass().New()
}




