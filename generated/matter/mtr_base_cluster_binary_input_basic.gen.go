// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterBinaryInputBasic] class.
var (
	MTRBaseClusterBinaryInputBasicClass     _MTRBaseClusterBinaryInputBasicClass
	MTRBaseClusterBinaryInputBasicClassOnce sync.Once
)

func getMTRBaseClusterBinaryInputBasicClass() _MTRBaseClusterBinaryInputBasicClass {
	MTRBaseClusterBinaryInputBasicClassOnce.Do(func() {
		MTRBaseClusterBinaryInputBasicClass = _MTRBaseClusterBinaryInputBasicClass{objc.GetClass("MTRBaseClusterBinaryInputBasic")}
	})
	return MTRBaseClusterBinaryInputBasicClass
}

type _MTRBaseClusterBinaryInputBasicClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterBinaryInputBasic] class.
type IMTRBaseClusterBinaryInputBasic interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterBinaryInputBasic
type MTRBaseClusterBinaryInputBasic struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterBinaryInputBasicFrom constructs a [MTRBaseClusterBinaryInputBasic] from an unsafe.Pointer.
func MTRBaseClusterBinaryInputBasicFrom(ptr unsafe.Pointer) MTRBaseClusterBinaryInputBasic {
	return MTRBaseClusterBinaryInputBasic{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterBinaryInputBasicClass) Alloc() MTRBaseClusterBinaryInputBasic {
	rv := objc.Send[MTRBaseClusterBinaryInputBasic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterBinaryInputBasicClass) New() MTRBaseClusterBinaryInputBasic {
	rv := objc.Send[MTRBaseClusterBinaryInputBasic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterBinaryInputBasic) Init() MTRBaseClusterBinaryInputBasic {
	rv := objc.Send[MTRBaseClusterBinaryInputBasic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterBinaryInputBasic) Autorelease() MTRBaseClusterBinaryInputBasic {
	rv := objc.Send[MTRBaseClusterBinaryInputBasic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterBinaryInputBasic creates a new MTRBaseClusterBinaryInputBasic instance.
func NewMTRBaseClusterBinaryInputBasic() MTRBaseClusterBinaryInputBasic {
	return getMTRBaseClusterBinaryInputBasicClass().New()
}
