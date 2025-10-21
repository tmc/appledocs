// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBasic] class.
var (
	MTRClusterBasicClass     _MTRClusterBasicClass
	MTRClusterBasicClassOnce sync.Once
)

func getMTRClusterBasicClass() _MTRClusterBasicClass {
	MTRClusterBasicClassOnce.Do(func() {
		MTRClusterBasicClass = _MTRClusterBasicClass{objc.GetClass("MTRClusterBasic")}
	})
	return MTRClusterBasicClass
}

type _MTRClusterBasicClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBasic] class.
type IMTRClusterBasic interface {
	IMTRClusterBasicInformation
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBasic
type MTRClusterBasic struct {
	MTRClusterBasicInformation
}

// MTRClusterBasicFrom constructs a [MTRClusterBasic] from an unsafe.Pointer.
func MTRClusterBasicFrom(ptr unsafe.Pointer) MTRClusterBasic {
	return MTRClusterBasic{
		MTRClusterBasicInformation: MTRClusterBasicInformationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBasicClass) Alloc() MTRClusterBasic {
	rv := objc.Send[MTRClusterBasic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBasicClass) New() MTRClusterBasic {
	rv := objc.Send[MTRClusterBasic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBasic) Init() MTRClusterBasic {
	rv := objc.Send[MTRClusterBasic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBasic) Autorelease() MTRClusterBasic {
	rv := objc.Send[MTRClusterBasic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBasic creates a new MTRClusterBasic instance.
func NewMTRClusterBasic() MTRClusterBasic {
	return getMTRClusterBasicClass().New()
}




