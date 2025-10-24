// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterKeypadInput] class.
var (
	MTRClusterKeypadInputClass     _MTRClusterKeypadInputClass
	MTRClusterKeypadInputClassOnce sync.Once
)

func getMTRClusterKeypadInputClass() _MTRClusterKeypadInputClass {
	MTRClusterKeypadInputClassOnce.Do(func() {
		MTRClusterKeypadInputClass = _MTRClusterKeypadInputClass{objc.GetClass("MTRClusterKeypadInput")}
	})
	return MTRClusterKeypadInputClass
}

type _MTRClusterKeypadInputClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterKeypadInput] class.
type IMTRClusterKeypadInput interface {
	IMTRGenericCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterKeypadInput
type MTRClusterKeypadInput struct {
	MTRGenericCluster
}

// MTRClusterKeypadInputFrom constructs a [MTRClusterKeypadInput] from an unsafe.Pointer.
func MTRClusterKeypadInputFrom(ptr unsafe.Pointer) MTRClusterKeypadInput {
	return MTRClusterKeypadInput{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterKeypadInputClass) Alloc() MTRClusterKeypadInput {
	rv := objc.Send[MTRClusterKeypadInput](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterKeypadInputClass) New() MTRClusterKeypadInput {
	rv := objc.Send[MTRClusterKeypadInput](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterKeypadInput) Init() MTRClusterKeypadInput {
	rv := objc.Send[MTRClusterKeypadInput](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterKeypadInput) Autorelease() MTRClusterKeypadInput {
	rv := objc.Send[MTRClusterKeypadInput](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterKeypadInput creates a new MTRClusterKeypadInput instance.
func NewMTRClusterKeypadInput() MTRClusterKeypadInput {
	return getMTRClusterKeypadInputClass().New()
}
