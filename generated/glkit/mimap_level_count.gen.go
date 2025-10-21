// Code generated from Apple documentation for GLKit. DO NOT EDIT.

package glkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [mimapLevelCount] class.
var (
	MimapLevelCountClass     _mimapLevelCountClass
	MimapLevelCountClassOnce sync.Once
)

func getmimapLevelCountClass() _mimapLevelCountClass {
	MimapLevelCountClassOnce.Do(func() {
		MimapLevelCountClass = _mimapLevelCountClass{objc.GetClass("mimapLevelCount")}
	})
	return MimapLevelCountClass
}

type _mimapLevelCountClass struct {
	class objc.Class
}

// An interface definition for the [mimapLevelCount] class.
type ImimapLevelCount interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/GLKit/GLKTextureInfo/mimapLevelCount-c.ivar
type mimapLevelCount struct {
	objectivec.Object
}

// mimapLevelCountFrom constructs a [mimapLevelCount] from an unsafe.Pointer.
func mimapLevelCountFrom(ptr unsafe.Pointer) mimapLevelCount {
	return mimapLevelCount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _mimapLevelCountClass) Alloc() mimapLevelCount {
	rv := objc.Send[mimapLevelCount](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _mimapLevelCountClass) New() mimapLevelCount {
	rv := objc.Send[mimapLevelCount](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ mimapLevelCount) Init() mimapLevelCount {
	rv := objc.Send[mimapLevelCount](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ mimapLevelCount) Autorelease() mimapLevelCount {
	rv := objc.Send[mimapLevelCount](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewmimapLevelCount creates a new mimapLevelCount instance.
func NewmimapLevelCount() mimapLevelCount {
	return getmimapLevelCountClass().New()
}




