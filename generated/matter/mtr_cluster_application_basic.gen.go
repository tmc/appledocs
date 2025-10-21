// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterApplicationBasic] class.
var (
	MTRClusterApplicationBasicClass     _MTRClusterApplicationBasicClass
	MTRClusterApplicationBasicClassOnce sync.Once
)

func getMTRClusterApplicationBasicClass() _MTRClusterApplicationBasicClass {
	MTRClusterApplicationBasicClassOnce.Do(func() {
		MTRClusterApplicationBasicClass = _MTRClusterApplicationBasicClass{objc.GetClass("MTRClusterApplicationBasic")}
	})
	return MTRClusterApplicationBasicClass
}

type _MTRClusterApplicationBasicClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterApplicationBasic] class.
type IMTRClusterApplicationBasic interface {
	IMTRGenericCluster
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterApplicationBasic
type MTRClusterApplicationBasic struct {
	MTRGenericCluster
}

// MTRClusterApplicationBasicFrom constructs a [MTRClusterApplicationBasic] from an unsafe.Pointer.
func MTRClusterApplicationBasicFrom(ptr unsafe.Pointer) MTRClusterApplicationBasic {
	return MTRClusterApplicationBasic{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterApplicationBasicClass) Alloc() MTRClusterApplicationBasic {
	rv := objc.Send[MTRClusterApplicationBasic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterApplicationBasicClass) New() MTRClusterApplicationBasic {
	rv := objc.Send[MTRClusterApplicationBasic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterApplicationBasic) Init() MTRClusterApplicationBasic {
	rv := objc.Send[MTRClusterApplicationBasic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterApplicationBasic) Autorelease() MTRClusterApplicationBasic {
	rv := objc.Send[MTRClusterApplicationBasic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterApplicationBasic creates a new MTRClusterApplicationBasic instance.
func NewMTRClusterApplicationBasic() MTRClusterApplicationBasic {
	return getMTRClusterApplicationBasicClass().New()
}




