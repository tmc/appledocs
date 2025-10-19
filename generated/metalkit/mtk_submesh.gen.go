// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MTKSubmesh] class.
var (
	mTKSubmeshClass     _MTKSubmeshClass
	mTKSubmeshClassOnce sync.Once
)

func getMTKSubmeshClass() _MTKSubmeshClass {
	mTKSubmeshClassOnce.Do(func() {
		mTKSubmeshClass = _MTKSubmeshClass{objc.GetClass("MTKSubmesh")}
	})
	return mTKSubmeshClass
}

type _MTKSubmeshClass struct {
	class objc.Class
}

// An interface definition for the [MTKSubmesh] class.
type IMTKSubmesh interface {
	objectivec.IObject
}

// A container for the index data of a Model I/O submesh, suitable for use in a Metal app. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKSubmesh
type MTKSubmesh struct {
	objectivec.Object
}

// MTKSubmeshFrom constructs a [MTKSubmesh] from an unsafe.Pointer.
//
// A container for the index data of a Model I/O submesh, suitable for use in a Metal app.
func MTKSubmeshFrom(ptr unsafe.Pointer) MTKSubmesh {
	return MTKSubmesh{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTKSubmeshClass) Alloc() MTKSubmesh {
	rv := objc.Send[MTKSubmesh](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MTKSubmeshClass) New() MTKSubmesh {
	rv := objc.Send[MTKSubmesh](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTKSubmesh) Init() MTKSubmesh {
	rv := objc.Send[MTKSubmesh](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTKSubmesh) Autorelease() MTKSubmesh {
	rv := objc.Send[MTKSubmesh](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKSubmesh creates a new MTKSubmesh instance.
func NewMTKSubmesh() MTKSubmesh {
	return getMTKSubmeshClass().New()
}




