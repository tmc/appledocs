// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptionRegion] class.
var (
	CaptionRegionClass     _CaptionRegionClass
	CaptionRegionClassOnce sync.Once
)

func getCaptionRegionClass() _CaptionRegionClass {
	CaptionRegionClassOnce.Do(func() {
		CaptionRegionClass = _CaptionRegionClass{objc.GetClass("AVCaptionRegion")}
	})
	return CaptionRegionClass
}

type _CaptionRegionClass struct {
	class objc.Class
}





// An interface definition for the [CaptionRegion] class.
type ICaptionRegion interface {
	objectivec.IObject
	

	// properties:
	DisplayAlignment() CaptionRegionDisplayAlignment
	Identifier() objc.IObject /* cross-framework: NSString */
	Origin() objc.IObject /* cross-framework: AVCaptionPoint */
	Scroll() CaptionRegionScroll
	Size() objc.IObject /* cross-framework: AVCaptionSize */
	WritingMode() CaptionRegionWritingMode


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptionRegionClass) Alloc() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionRegionClass) New() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionRegion) Init() CaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionRegion) Autorelease() CaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionRegion creates a new CaptionRegion instance.
func NewCaptionRegion() CaptionRegion {
	return getCaptionRegionClass().New()
}





// An object that represents the region in which the system presents a caption.
//
// The framework defines four regions, and doesn’t support configuring region settings.


// An object that represents the region in which the system presents a caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion
type CaptionRegion struct {
	objectivec.Object
}

// CaptionRegionFrom constructs a [CaptionRegion] from an unsafe.Pointer.
//
// An object that represents the region in which the system presents a caption.
func CaptionRegionFrom(ptr unsafe.Pointer) CaptionRegion {
	return CaptionRegion{objectivec.Object{objc.ID(ptr)}}
}















// The bottom region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTBottom
func (cc _CaptionRegionClass) AppleITTBottomRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("appleITTBottomRegion"))
	return rv
}

// The left region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTLeft
func (cc _CaptionRegionClass) AppleITTLeftRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("appleITTLeftRegion"))
	return rv
}

// The right region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTRight
func (cc _CaptionRegionClass) AppleITTRightRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("appleITTRightRegion"))
	return rv
}

// The top region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTTop
func (cc _CaptionRegionClass) AppleITTTopRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("appleITTTopRegion"))
	return rv
}

// The bottom caption region for SubRip Text (SRT) format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/subRipTextBottom
func (cc _CaptionRegionClass) SubRipTextBottomRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("subRipTextBottomRegion"))
	return rv
}






// Encodes the region using the specified encoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/encode(with:)
func (c_ CaptionRegion) EncodeWithCoder(encoder foundation.Coder) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeWithCoder:"), encoder)
}


// Returns a Boolean value that indicates whether an object equals another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/isEqual(_:)
func (c_ CaptionRegion) IsEqual(object objc.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqual:"), object)
	return rv
}







// The bottom region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTBottom
func (c_ CaptionRegion) AppleITTBottomRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("appleITTBottomRegion"))
	return rv
}


// The left region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTLeft
func (c_ CaptionRegion) AppleITTLeftRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("appleITTLeftRegion"))
	return rv
}


// The right region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTRight
func (c_ CaptionRegion) AppleITTRightRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("appleITTRightRegion"))
	return rv
}


// The top region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTTop
func (c_ CaptionRegion) AppleITTTopRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("appleITTTopRegion"))
	return rv
}


// The alignment of lines for the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/displayAlignment-swift.property
func (c_ CaptionRegion) DisplayAlignment() CaptionRegionDisplayAlignment {
	rv := objc.Send[CaptionRegionDisplayAlignment](c_.ID, objc.Sel("displayAlignment"))
	return rv
}


// A string that identifies the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/identifier
func (c_ CaptionRegion) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}


// The region’s top-left position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/origin
func (c_ CaptionRegion) Origin() objc.IObject /* cross-framework: AVCaptionPoint */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("origin"))
	return rv
}


// The scroll mode of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/scroll-swift.property
func (c_ CaptionRegion) Scroll() CaptionRegionScroll {
	rv := objc.Send[CaptionRegionScroll](c_.ID, objc.Sel("scroll"))
	return rv
}


// The height and width of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/size
func (c_ CaptionRegion) Size() objc.IObject /* cross-framework: AVCaptionSize */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("size"))
	return rv
}


// The bottom caption region for SubRip Text (SRT) format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/subRipTextBottom
func (c_ CaptionRegion) SubRipTextBottomRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("subRipTextBottomRegion"))
	return rv
}


// The block and inline progression direction of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/writingMode-swift.property
func (c_ CaptionRegion) WritingMode() CaptionRegionWritingMode {
	rv := objc.Send[CaptionRegionWritingMode](c_.ID, objc.Sel("writingMode"))
	return rv
}








