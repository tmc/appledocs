// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptionRegion */


/* debug [class_header]: Header for AVCaptionRegion */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptionRegion */
// An interface definition for the [CaptionRegion] class.
type ICaptionRegion interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptionRegion */
	// properties:
	DisplayAlignment() CaptionRegionDisplayAlignment
	Identifier() objc.IObject /* cross-framework: NSString */
	Origin() objc.IObject /* cross-framework: AVCaptionPoint */
	Scroll() CaptionRegionScroll
	Size() objc.IObject /* cross-framework: AVCaptionSize */
	WritingMode() CaptionRegionWritingMode
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptionRegion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptionRegion */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptionRegion */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptionRegion *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptionRegion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptionRegion */

// The bottom region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTBottom
func (cc _CaptionRegionClass) AppleITTBottomRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("appleITTBottomRegion"))
	return rv
}/* debug [class_properties_class/property]: appleITTBottomRegion */

// The left region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTLeft
func (cc _CaptionRegionClass) AppleITTLeftRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("appleITTLeftRegion"))
	return rv
}/* debug [class_properties_class/property]: appleITTLeftRegion */

// The right region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTRight
func (cc _CaptionRegionClass) AppleITTRightRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("appleITTRightRegion"))
	return rv
}/* debug [class_properties_class/property]: appleITTRightRegion */

// The top region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTTop
func (cc _CaptionRegionClass) AppleITTTopRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("appleITTTopRegion"))
	return rv
}/* debug [class_properties_class/property]: appleITTTopRegion */

// The bottom caption region for SubRip Text (SRT) format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/subRipTextBottom
func (cc _CaptionRegionClass) SubRipTextBottomRegion() CaptionRegion {
	rv := objc.Send[CaptionRegion](objc.ID(cc.class), objc.Sel("subRipTextBottomRegion"))
	return rv
}/* debug [class_properties_class/property]: subRipTextBottomRegion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptionRegion */

// Encodes the region using the specified encoder.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/encode(with:)
func (c_ CaptionRegion) EncodeWithCoder(encoder foundation.Coder) {
	objc.Send[objc.ID](c_.ID, objc.Sel("encodeWithCoder:"), encoder)
}/* debug [instance_methods/method]: EncodeWithCoder */


// Returns a Boolean value that indicates whether an object equals another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/isEqual(_:)
func (c_ CaptionRegion) IsEqual(object objc.IObject) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqual:"), object)
	return rv
}/* debug [instance_methods/method]: IsEqual */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptionRegion */

// The bottom region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTBottom
func (c_ CaptionRegion) AppleITTBottomRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("appleITTBottomRegion"))
	return rv
}/* debug [instance_properties/getter]: appleITTBottomRegion */


// The left region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTLeft
func (c_ CaptionRegion) AppleITTLeftRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("appleITTLeftRegion"))
	return rv
}/* debug [instance_properties/getter]: appleITTLeftRegion */


// The right region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTRight
func (c_ CaptionRegion) AppleITTRightRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("appleITTRightRegion"))
	return rv
}/* debug [instance_properties/getter]: appleITTRightRegion */


// The top region for iTT format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/appleITTTop
func (c_ CaptionRegion) AppleITTTopRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("appleITTTopRegion"))
	return rv
}/* debug [instance_properties/getter]: appleITTTopRegion */


// The alignment of lines for the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/displayAlignment-swift.property
func (c_ CaptionRegion) DisplayAlignment() CaptionRegionDisplayAlignment {
	rv := objc.Send[CaptionRegionDisplayAlignment](c_.ID, objc.Sel("displayAlignment"))
	return rv
}/* debug [instance_properties/getter]: displayAlignment */


// A string that identifies the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/identifier
func (c_ CaptionRegion) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// The region’s top-left position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/origin
func (c_ CaptionRegion) Origin() objc.IObject /* cross-framework: AVCaptionPoint */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("origin"))
	return rv
}/* debug [instance_properties/getter]: origin */


// The scroll mode of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/scroll-swift.property
func (c_ CaptionRegion) Scroll() CaptionRegionScroll {
	rv := objc.Send[CaptionRegionScroll](c_.ID, objc.Sel("scroll"))
	return rv
}/* debug [instance_properties/getter]: scroll */


// The height and width of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/size
func (c_ CaptionRegion) Size() objc.IObject /* cross-framework: AVCaptionSize */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The bottom caption region for SubRip Text (SRT) format captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/subRipTextBottom
func (c_ CaptionRegion) SubRipTextBottomRegion() IAVCaptionRegion {
	rv := objc.Send[CaptionRegion](c_.ID, objc.Sel("subRipTextBottomRegion"))
	return rv
}/* debug [instance_properties/getter]: subRipTextBottomRegion */


// The block and inline progression direction of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/writingMode-swift.property
func (c_ CaptionRegion) WritingMode() CaptionRegionWritingMode {
	rv := objc.Send[CaptionRegionWritingMode](c_.ID, objc.Sel("writingMode"))
	return rv
}/* debug [instance_properties/getter]: writingMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptionRegion */



