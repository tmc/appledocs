// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRClusterBinaryInputBasic] class.
var (
	MTRClusterBinaryInputBasicClass     _MTRClusterBinaryInputBasicClass
	MTRClusterBinaryInputBasicClassOnce sync.Once
)

func getMTRClusterBinaryInputBasicClass() _MTRClusterBinaryInputBasicClass {
	MTRClusterBinaryInputBasicClassOnce.Do(func() {
		MTRClusterBinaryInputBasicClass = _MTRClusterBinaryInputBasicClass{objc.GetClass("MTRClusterBinaryInputBasic")}
	})
	return MTRClusterBinaryInputBasicClass
}

type _MTRClusterBinaryInputBasicClass struct {
	class objc.Class
}

// An interface definition for the [MTRClusterBinaryInputBasic] class.
type IMTRClusterBinaryInputBasic interface {
	IMTRGenericCluster
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRClusterBinaryInputBasic
type MTRClusterBinaryInputBasic struct {
	MTRGenericCluster
}

// MTRClusterBinaryInputBasicFrom constructs a [MTRClusterBinaryInputBasic] from an unsafe.Pointer.
func MTRClusterBinaryInputBasicFrom(ptr unsafe.Pointer) MTRClusterBinaryInputBasic {
	return MTRClusterBinaryInputBasic{
		MTRGenericCluster: MTRGenericClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRClusterBinaryInputBasicClass) Alloc() MTRClusterBinaryInputBasic {
	rv := objc.Send[MTRClusterBinaryInputBasic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRClusterBinaryInputBasicClass) New() MTRClusterBinaryInputBasic {
	rv := objc.Send[MTRClusterBinaryInputBasic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRClusterBinaryInputBasic) Init() MTRClusterBinaryInputBasic {
	rv := objc.Send[MTRClusterBinaryInputBasic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRClusterBinaryInputBasic) Autorelease() MTRClusterBinaryInputBasic {
	rv := objc.Send[MTRClusterBinaryInputBasic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRClusterBinaryInputBasic creates a new MTRClusterBinaryInputBasic instance.
func NewMTRClusterBinaryInputBasic() MTRClusterBinaryInputBasic {
	return getMTRClusterBinaryInputBasicClass().New()
}




