// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [TextLayer] class.
var (
	textLayerClass     _TextLayerClass
	textLayerClassOnce sync.Once
)

func getTextLayerClass() _TextLayerClass {
	textLayerClassOnce.Do(func() {
		textLayerClass = _TextLayerClass{objc.GetClass("CATextLayer")}
	})
	return textLayerClass
}

type _TextLayerClass struct {
	class objc.Class
}

// An interface definition for the [TextLayer] class.
type ITextLayer interface {
	ILayer
}

// A layer that provides simple text layout and rendering of plain or attributed strings.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATextLayer
type TextLayer struct {
	Layer
}

// TextLayerFrom constructs a [TextLayer] from an unsafe.Pointer.
//
// A layer that provides simple text layout and rendering of plain or attributed strings.
func TextLayerFrom(ptr unsafe.Pointer) TextLayer {
	return TextLayer{
		Layer: LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TextLayerClass) Alloc() TextLayer {
	rv := objc.Send[TextLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TextLayerClass) New() TextLayer {
	rv := objc.Send[TextLayer](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TextLayer) Init() TextLayer {
	rv := objc.Send[TextLayer](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TextLayer) Autorelease() TextLayer {
	rv := objc.Send[TextLayer](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTextLayer creates a new TextLayer instance.
func NewTextLayer() TextLayer {
	return getTextLayerClass().New()
}




