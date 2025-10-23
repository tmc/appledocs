// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ColorSampler] class.
var (
	ColorSamplerClass     _ColorSamplerClass
	ColorSamplerClassOnce sync.Once
)

func getColorSamplerClass() _ColorSamplerClass {
	ColorSamplerClassOnce.Do(func() {
		ColorSamplerClass = _ColorSamplerClass{objc.GetClass("NSColorSampler")}
	})
	return ColorSamplerClass
}

type _ColorSamplerClass struct {
	class objc.Class
}

// An interface definition for the [ColorSampler] class.
type IColorSampler interface {
	objectivec.IObject
	ShowSamplerWithSelectionHandler(selectionHandler unsafe.Pointer)
}

// An object that displays the system’s color-sampling interface and returns the selected color to your app.
//
// Create an object when you want the user to select a color based on existing onscreen colors. When you call the method, AppKit shows the system’s color sampler interface and reports the selected color back to the provided block.


// An object that displays the system’s color-sampling interface and returns the selected color to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSampler
type ColorSampler struct {
	objectivec.Object
}

// ColorSamplerFrom constructs a [ColorSampler] from an unsafe.Pointer.
//
// An object that displays the system’s color-sampling interface and returns the selected color to your app.
func ColorSamplerFrom(ptr unsafe.Pointer) ColorSampler {
	return ColorSampler{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _ColorSamplerClass) Alloc() ColorSampler {
	rv := objc.Send[ColorSampler](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _ColorSamplerClass) New() ColorSampler {
	rv := objc.Send[ColorSampler](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ColorSampler) Init() ColorSampler {
	rv := objc.Send[ColorSampler](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ColorSampler) Autorelease() ColorSampler {
	rv := objc.Send[ColorSampler](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewColorSampler creates a new ColorSampler instance.
func NewColorSampler() ColorSampler {
	return getColorSamplerClass().New()
}



// Displays the system color-sampling interface asynchronously and reports the selected color back to your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSampler/show(selectionHandler:)
func (c_ ColorSampler) ShowSamplerWithSelectionHandler(selectionHandler unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("showSamplerWithSelectionHandler:"), selectionHandler)
}



