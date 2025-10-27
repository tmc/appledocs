// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [RenderedCaptionImage] class.
var (
	RenderedCaptionImageClass     _RenderedCaptionImageClass
	RenderedCaptionImageClassOnce sync.Once
)

func getRenderedCaptionImageClass() _RenderedCaptionImageClass {
	RenderedCaptionImageClassOnce.Do(func() {
		RenderedCaptionImageClass = _RenderedCaptionImageClass{objc.GetClass("AVRenderedCaptionImage")}
	})
	return RenderedCaptionImageClass
}

type _RenderedCaptionImageClass struct {
	class objc.Class
}





// An interface definition for the [RenderedCaptionImage] class.
type IRenderedCaptionImage interface {
	objectivec.IObject
	

	// properties:
	PixelBuffer() PixelBufferRef /* not a class type */
	Position() corefoundation.CGPoint
	ReadOnlyPixelBuffer() ReadOnlyPixelBuffer /* not a class type */
	SetReadOnlyPixelBuffer(value ReadOnlyPixelBuffer /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (rc _RenderedCaptionImageClass) Alloc() RenderedCaptionImage {
	rv := objc.Send[RenderedCaptionImage](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _RenderedCaptionImageClass) New() RenderedCaptionImage {
	rv := objc.Send[RenderedCaptionImage](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ RenderedCaptionImage) Init() RenderedCaptionImage {
	rv := objc.Send[RenderedCaptionImage](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ RenderedCaptionImage) Autorelease() RenderedCaptionImage {
	rv := objc.Send[RenderedCaptionImage](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewRenderedCaptionImage creates a new RenderedCaptionImage instance.
func NewRenderedCaptionImage() RenderedCaptionImage {
	return getRenderedCaptionImageClass().New()
}





// An object that provides a rendered pixel buffer and its position in pixels.


// An object that provides a rendered pixel buffer and its position in pixels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRenderedCaptionImage
type RenderedCaptionImage struct {
	objectivec.Object
}

// RenderedCaptionImageFrom constructs a [RenderedCaptionImage] from an unsafe.Pointer.
//
// An object that provides a rendered pixel buffer and its position in pixels.
func RenderedCaptionImageFrom(ptr unsafe.Pointer) RenderedCaptionImage {
	return RenderedCaptionImage{objectivec.Object{objc.ID(ptr)}}
}

























// An object that contains pixel data for the rendered caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRenderedCaptionImage/pixelBuffer
func (r_ RenderedCaptionImage) PixelBuffer() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](r_.ID, objc.Sel("pixelBuffer"))
	return rv
}


// A point that defines the position, in pixels, of the rendered caption image relative to the video frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVRenderedCaptionImage/position
func (r_ RenderedCaptionImage) Position() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](r_.ID, objc.Sel("position"))
	return rv
}


// A CVReadOnlyPixelBuffer that contains pixel data for the rendered caption
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avrenderedcaptionimage/readonlypixelbuffer
func (r_ RenderedCaptionImage) ReadOnlyPixelBuffer() ReadOnlyPixelBuffer /* not a class type */ {
	rv := objc.Send[ReadOnlyPixelBuffer](r_.ID, objc.Sel("readOnlyPixelBuffer"))
	return rv
}


// A CVReadOnlyPixelBuffer that contains pixel data for the rendered caption
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avrenderedcaptionimage/readonlypixelbuffer
func (r_ RenderedCaptionImage) SetReadOnlyPixelBuffer(value ReadOnlyPixelBuffer /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setReadOnlyPixelBuffer:"), value)
}








