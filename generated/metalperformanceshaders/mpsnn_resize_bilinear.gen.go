// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ResizeBilinear] class.
var (
	ResizeBilinearClass     _ResizeBilinearClass
	ResizeBilinearClassOnce sync.Once
)

func getResizeBilinearClass() _ResizeBilinearClass {
	ResizeBilinearClassOnce.Do(func() {
		ResizeBilinearClass = _ResizeBilinearClass{objc.GetClass("MPSNNResizeBilinear")}
	})
	return ResizeBilinearClass
}

type _ResizeBilinearClass struct {
	class objc.Class
}

// An interface definition for the [ResizeBilinear] class.
type IResizeBilinear interface {
	objectivec.IObject
}

// A bilinear resizing filter.
//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear
type ResizeBilinear struct {
	objectivec.Object
}

// ResizeBilinearFrom constructs a [ResizeBilinear] from an unsafe.Pointer.
//
// A bilinear resizing filter.
func ResizeBilinearFrom(ptr unsafe.Pointer) ResizeBilinear {
	return ResizeBilinear{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (rc _ResizeBilinearClass) Alloc() ResizeBilinear {
	rv := objc.Send[ResizeBilinear](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (rc _ResizeBilinearClass) New() ResizeBilinear {
	rv := objc.Send[ResizeBilinear](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ResizeBilinear) Init() ResizeBilinear {
	rv := objc.Send[ResizeBilinear](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ResizeBilinear) Autorelease() ResizeBilinear {
	rv := objc.Send[ResizeBilinear](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewResizeBilinear creates a new ResizeBilinear instance.
func NewResizeBilinear() ResizeBilinear {
	return getResizeBilinearClass().New()
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/init(coder:device:)
func NewResizeBilinearWithCoderDevice(aDecoder unsafe.Pointer, device objc.ID) ResizeBilinear {
	instance := getResizeBilinearClass().Alloc()
	rv := objc.Send[ResizeBilinear](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/alignCorners
func (r_ ResizeBilinear) AlignCorners() bool {
	rv := objc.Send[bool](r_.ID, objc.Sel("alignCorners"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/resizeHeight
func (r_ ResizeBilinear) ResizeHeight() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("resizeHeight"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNResizeBilinear/resizeWidth
func (r_ ResizeBilinear) ResizeWidth() uint {
	rv := objc.Send[uint](r_.ID, objc.Sel("resizeWidth"))
	return rv
}
