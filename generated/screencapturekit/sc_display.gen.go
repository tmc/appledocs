// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Display] class.
var (
	DisplayClass     _DisplayClass
	DisplayClassOnce sync.Once
)

func getDisplayClass() _DisplayClass {
	DisplayClassOnce.Do(func() {
		DisplayClass = _DisplayClass{objc.GetClass("SCDisplay")}
	})
	return DisplayClass
}

type _DisplayClass struct {
	class objc.Class
}

// An interface definition for the [Display] class.
type IDisplay interface {
	objectivec.IObject
	DisplayID() unsafe.Pointer
	Frame() coregraphics.CGRect
	Height() int
	SetHeight(value int)
	Width() int
	SetWidth(value int)
}

// An instance that represents a display device.
//
// A display object represents a physical display connected to a Mac. Query the display to retrieve its unique identifier and onscreen coordinates. Retrieve the available displays from an instance of . Select a display to capture and use it to create an instance of . Apply the filter to an instance of to limit its output to content matching your criteria.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay
type Display struct {
	objectivec.Object
}

// DisplayFrom constructs a [Display] from an unsafe.Pointer.
//
// An instance that represents a display device.
func DisplayFrom(ptr unsafe.Pointer) Display {
	return Display{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (dc _DisplayClass) Alloc() Display {
	rv := objc.Send[Display](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (dc _DisplayClass) New() Display {
	rv := objc.Send[Display](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ Display) Init() Display {
	rv := objc.Send[Display](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ Display) Autorelease() Display {
	rv := objc.Send[Display](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDisplay creates a new Display instance.
func NewDisplay() Display {
	return getDisplayClass().New()
}


// The Core Graphics display identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/displayID
func (d_ Display) DisplayID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](d_.ID, objc.Sel("displayID"))
	return rv
}

// The frame of the display.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenCaptureKit/SCDisplay/frame
func (d_ Display) Frame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](d_.ID, objc.Sel("frame"))
	return rv
}

// The height of the display in points.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scdisplay/height
func (d_ Display) Height() int {
	rv := objc.Send[int](d_.ID, objc.Sel("height"))
	return rv
}


// SetHeight sets the value of the height property.
// The height of the display in points.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scdisplay/height
func (d_ Display) SetHeight(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setHeight:"), value)
}

// The width of the display in points.
//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scdisplay/width
func (d_ Display) Width() int {
	rv := objc.Send[int](d_.ID, objc.Sel("width"))
	return rv
}


// SetWidth sets the value of the width property.
// The width of the display in points.

//
// [Full Topic]: https://developer.apple.com/documentation/screencapturekit/scdisplay/width
func (d_ Display) SetWidth(value int) {
	objc.Send[objc.ID](d_.ID, objc.Sel("setWidth:"), value)
}



