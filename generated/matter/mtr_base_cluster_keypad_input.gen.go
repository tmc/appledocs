// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterKeypadInput] class.
var (
	MTRBaseClusterKeypadInputClass     _MTRBaseClusterKeypadInputClass
	MTRBaseClusterKeypadInputClassOnce sync.Once
)

func getMTRBaseClusterKeypadInputClass() _MTRBaseClusterKeypadInputClass {
	MTRBaseClusterKeypadInputClassOnce.Do(func() {
		MTRBaseClusterKeypadInputClass = _MTRBaseClusterKeypadInputClass{objc.GetClass("MTRBaseClusterKeypadInput")}
	})
	return MTRBaseClusterKeypadInputClass
}

type _MTRBaseClusterKeypadInputClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterKeypadInput] class.
type IMTRBaseClusterKeypadInput interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterKeypadInput
type MTRBaseClusterKeypadInput struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterKeypadInputFrom constructs a [MTRBaseClusterKeypadInput] from an unsafe.Pointer.
func MTRBaseClusterKeypadInputFrom(ptr unsafe.Pointer) MTRBaseClusterKeypadInput {
	return MTRBaseClusterKeypadInput{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterKeypadInputClass) Alloc() MTRBaseClusterKeypadInput {
	rv := objc.Send[MTRBaseClusterKeypadInput](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterKeypadInputClass) New() MTRBaseClusterKeypadInput {
	rv := objc.Send[MTRBaseClusterKeypadInput](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterKeypadInput) Init() MTRBaseClusterKeypadInput {
	rv := objc.Send[MTRBaseClusterKeypadInput](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterKeypadInput) Autorelease() MTRBaseClusterKeypadInput {
	rv := objc.Send[MTRBaseClusterKeypadInput](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterKeypadInput creates a new MTRBaseClusterKeypadInput instance.
func NewMTRBaseClusterKeypadInput() MTRBaseClusterKeypadInput {
	return getMTRBaseClusterKeypadInputClass().New()
}
