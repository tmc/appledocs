// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTRAttributePath] class.
var (
	MTRAttributePathClass     _MTRAttributePathClass
	MTRAttributePathClassOnce sync.Once
)

func getMTRAttributePathClass() _MTRAttributePathClass {
	MTRAttributePathClassOnce.Do(func() {
		MTRAttributePathClass = _MTRAttributePathClass{objc.GetClass("MTRAttributePath")}
	})
	return MTRAttributePathClass
}

type _MTRAttributePathClass struct {
	class objc.Class
}

// An interface definition for the [MTRAttributePath] class.
type IMTRAttributePath interface {
	IMTRClusterPath
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRAttributePath
type MTRAttributePath struct {
	MTRClusterPath
}

// MTRAttributePathFrom constructs a [MTRAttributePath] from an unsafe.Pointer.
func MTRAttributePathFrom(ptr unsafe.Pointer) MTRAttributePath {
	return MTRAttributePath{
		MTRClusterPath: MTRClusterPathFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRAttributePathClass) Alloc() MTRAttributePath {
	rv := objc.Send[MTRAttributePath](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTRAttributePathClass) New() MTRAttributePath {
	rv := objc.Send[MTRAttributePath](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRAttributePath) Init() MTRAttributePath {
	rv := objc.Send[MTRAttributePath](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRAttributePath) Autorelease() MTRAttributePath {
	rv := objc.Send[MTRAttributePath](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRAttributePath creates a new MTRAttributePath instance.
func NewMTRAttributePath() MTRAttributePath {
	return getMTRAttributePathClass().New()
}




