// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Box] class.
var (
	boxClass     _BoxClass
	boxClassOnce sync.Once
)

func getBoxClass() _BoxClass {
	boxClassOnce.Do(func() {
		boxClass = _BoxClass{objc.GetClass("NSBox")}
	})
	return boxClass
}

type _BoxClass struct {
	class objc.Class
}

// An interface definition for the [Box] class.
type IBox interface {
	IView
	SetFrameFromContentFrame(contentFrame unsafe.Pointer)
	SetTitleWithMnemonic(stringWithAmpersand string)
	SizeToFit()
}

// A stylized rectangular box with an optional title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox
type Box struct {
	View
}

// BoxFrom constructs a [Box] from an unsafe.Pointer.
//
// A stylized rectangular box with an optional title.
func BoxFrom(ptr unsafe.Pointer) Box {
	return Box{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (bc _BoxClass) Alloc() Box {
	rv := objc.Send[Box](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (bc _BoxClass) New() Box {
	rv := objc.Send[Box](objc.ID(bc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (b_ Box) Init() Box {
	rv := objc.Send[Box](b_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b_ Box) Autorelease() Box {
	rv := objc.Send[Box](b_.ID, objc.Sel("autorelease"))
	return rv
}

// NewBox creates a new Box instance.
func NewBox() Box {
	return getBoxClass().New()
}


// Places the receiver so its content view lies on the specified frame.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/setFrameFromContentFrame(_:)
func (b_ Box) SetFrameFromContentFrame(contentFrame unsafe.Pointer) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setFrameFromContentFrame:"), contentFrame)
}

// Sets the title of the receiver with a character denoted as an access key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/setTitleWithMnemonic:
func (b_ Box) SetTitleWithMnemonic(stringWithAmpersand string) {
	objc.Send[objc.ID](b_.ID, objc.Sel("setTitleWithMnemonic:"), objc.String(stringWithAmpersand))
}

// Resizes and moves the receiver’s content view so it just encloses its subviews.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/sizeToFit()
func (b_ Box) SizeToFit() {
	objc.Send[objc.ID](b_.ID, objc.Sel("sizeToFit"))
}



