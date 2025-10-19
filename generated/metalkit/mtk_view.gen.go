// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MTKView] class.
var mTKViewClass = _MTKViewClass{objc.GetClass("MTKView")}

type _MTKViewClass struct {
	class objc.Class
}

// An interface definition for the [MTKView] class.
type IMTKView interface {
	IView
	Draw()
	ReleaseDrawables()
}

// A specialized view that creates, configures, and displays Metal objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView

type MTKView struct {
	View
}

// MTKViewFrom constructs a [MTKView] from an unsafe.Pointer.
//
// A specialized view that creates, configures, and displays Metal objects.
func MTKViewFrom(ptr unsafe.Pointer) MTKView {
	return MTKView{
		View: ViewFrom(ptr),
	}
}
// Alloc allocates a new instance without initialization.
func (mc _MTKViewClass) Alloc() MTKView {
	rv := objc.Send[MTKView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MTKViewClass) New() MTKView {
	rv := objc.Send[MTKView](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTKView) Init() MTKView {
	rv := objc.Send[MTKView](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTKView) Autorelease() MTKView {
	rv := objc.Send[MTKView](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKView creates a new MTKView instance.
func NewMTKView() MTKView {
	return mTKViewClass.New()
}


// Initializes a view with the specified frame rectangle and Metal device. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/init(frame:device:)
func NewMTKViewWithFrameDevice(frameRect unsafe.Pointer, device unsafe.Pointer) MTKView {
	instance := mTKViewClass.Alloc()
	rv := objc.Send[MTKView](instance.ID, objc.Sel("initWithFrame:device:"), frameRect, device)
	rv.Autorelease()
	return rv
}
// Initializes a view from data in a given unarchiver. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/init(coder:)
func NewMTKViewWithCoder(coder unsafe.Pointer) MTKView {
	instance := mTKViewClass.Alloc()
	rv := objc.Send[MTKView](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Redraws the view’s contents immediately. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/draw()
func (m_ MTKView) Draw() {
	objc.Send[objc.ID](m_.ID, objc.Sel("draw"))
}
// Releases the and objects. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/MetalKit/MTKView/releaseDrawables()
func (m_ MTKView) ReleaseDrawables() {
	objc.Send[objc.ID](m_.ID, objc.Sel("releaseDrawables"))
}

