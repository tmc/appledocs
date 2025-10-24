// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRBaseClusterApplicationBasic] class.
var (
	MTRBaseClusterApplicationBasicClass     _MTRBaseClusterApplicationBasicClass
	MTRBaseClusterApplicationBasicClassOnce sync.Once
)

func getMTRBaseClusterApplicationBasicClass() _MTRBaseClusterApplicationBasicClass {
	MTRBaseClusterApplicationBasicClassOnce.Do(func() {
		MTRBaseClusterApplicationBasicClass = _MTRBaseClusterApplicationBasicClass{objc.GetClass("MTRBaseClusterApplicationBasic")}
	})
	return MTRBaseClusterApplicationBasicClass
}

type _MTRBaseClusterApplicationBasicClass struct {
	class objc.Class
}

// An interface definition for the [MTRBaseClusterApplicationBasic] class.
type IMTRBaseClusterApplicationBasic interface {
	IMTRGenericBaseCluster
	// properties:
	// methods:
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRBaseClusterApplicationBasic
type MTRBaseClusterApplicationBasic struct {
	MTRGenericBaseCluster
}

// MTRBaseClusterApplicationBasicFrom constructs a [MTRBaseClusterApplicationBasic] from an unsafe.Pointer.
func MTRBaseClusterApplicationBasicFrom(ptr unsafe.Pointer) MTRBaseClusterApplicationBasic {
	return MTRBaseClusterApplicationBasic{
		MTRGenericBaseCluster: MTRGenericBaseClusterFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRBaseClusterApplicationBasicClass) Alloc() MTRBaseClusterApplicationBasic {
	rv := objc.Send[MTRBaseClusterApplicationBasic](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRBaseClusterApplicationBasicClass) New() MTRBaseClusterApplicationBasic {
	rv := objc.Send[MTRBaseClusterApplicationBasic](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRBaseClusterApplicationBasic) Init() MTRBaseClusterApplicationBasic {
	rv := objc.Send[MTRBaseClusterApplicationBasic](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRBaseClusterApplicationBasic) Autorelease() MTRBaseClusterApplicationBasic {
	rv := objc.Send[MTRBaseClusterApplicationBasic](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRBaseClusterApplicationBasic creates a new MTRBaseClusterApplicationBasic instance.
func NewMTRBaseClusterApplicationBasic() MTRBaseClusterApplicationBasic {
	return getMTRBaseClusterApplicationBasicClass().New()
}
