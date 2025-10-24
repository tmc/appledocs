// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBasic] class.
var (
	MTRBaseClusterBasicClass     _MTRBaseClusterBasicClass
	MTRBaseClusterBasicClassOnce sync.Once
)

func getMTRBaseClusterBasicClass() _MTRBaseClusterBasicClass {
	MTRBaseClusterBasicClassOnce.Do(func() {
		MTRBaseClusterBasicClass = _MTRBaseClusterBasicClass{objc.GetClass("MTRBaseClusterBasic")}
	})
	return MTRBaseClusterBasicClass
}

type _MTRBaseClusterBasicClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBasic] class.
type IMTRBaseClusterBasic interface {
	IMTRBaseClusterBasicInformation
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBasic
type MTRBaseClusterBasic struct {
	MTRBaseClusterBasicInformation
}

// MTRBaseClusterBasicFrom constructs a [MTRBaseClusterBasic] from an unsafe.Pointer.
func MTRBaseClusterBasicFrom(ptr unsafe.Pointer) MTRBaseClusterBasic {
	return MTRBaseClusterBasic{
		MTRBaseClusterBasicInformation: MTRBaseClusterBasicInformationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBasicClass) Alloc() MTRBaseClusterBasic {
	rv := objc.Send[MTRBaseClusterBasic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBasicClass) New() MTRBaseClusterBasic {
	rv := objc.Send[MTRBaseClusterBasic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBasic) Init() MTRBaseClusterBasic {
	rv := objc.Send[MTRBaseClusterBasic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBasic) Autorelease() MTRBaseClusterBasic {
	rv := objc.Send[MTRBaseClusterBasic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBasic creates a new MTRBaseClusterBasic instance.
func NewMTRBaseClusterBasic() MTRBaseClusterBasic {
	return getMTRBaseClusterBasicClass().New()
}
