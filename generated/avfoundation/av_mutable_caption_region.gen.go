// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVMutableCaptionRegion */


/* debug [class_header]: Header for AVMutableCaptionRegion */
// The class instance for the [MutableCaptionRegion] class.
var (
	MutableCaptionRegionClass     _MutableCaptionRegionClass
	MutableCaptionRegionClassOnce sync.Once
)

func getMutableCaptionRegionClass() _MutableCaptionRegionClass {
	MutableCaptionRegionClassOnce.Do(func() {
		MutableCaptionRegionClass = _MutableCaptionRegionClass{objc.GetClass("AVMutableCaptionRegion")}
	})
	return MutableCaptionRegionClass
}

type _MutableCaptionRegionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MutableCaptionRegion */
// An interface definition for the [MutableCaptionRegion] class.
type IMutableCaptionRegion interface {
	ICaptionRegion
	
/* debug [class_interface_properties]: Properties for MutableCaptionRegion */
	// properties:
	DisplayAlignment() CaptionRegionDisplayAlignment
	SetDisplayAlignment(value CaptionRegionDisplayAlignment)
	Origin() objc.IObject /* cross-framework: AVCaptionPoint */
	SetOrigin(value objc.IObject /* cross-framework: AVCaptionPoint */)
	Scroll() CaptionRegionScroll
	SetScroll(value CaptionRegionScroll)
	Size() objc.IObject /* cross-framework: AVCaptionSize */
	SetSize(value objc.IObject /* cross-framework: AVCaptionSize */)
	WritingMode() CaptionRegionWritingMode
	SetWritingMode(value CaptionRegionWritingMode)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MutableCaptionRegion */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MutableCaptionRegion */
// Alloc allocates a new instance without initialization.
func (mc _MutableCaptionRegionClass) Alloc() MutableCaptionRegion {
	rv := objc.Send[MutableCaptionRegion](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutableCaptionRegionClass) New() MutableCaptionRegion {
	rv := objc.Send[MutableCaptionRegion](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableCaptionRegion) Init() MutableCaptionRegion {
	rv := objc.Send[MutableCaptionRegion](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableCaptionRegion) Autorelease() MutableCaptionRegion {
	rv := objc.Send[MutableCaptionRegion](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableCaptionRegion creates a new MutableCaptionRegion instance.
func NewMutableCaptionRegion() MutableCaptionRegion {
	return getMutableCaptionRegionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MutableCaptionRegion */
// A mutable caption region subclass that you use to create new caption regions.


// A mutable caption region subclass that you use to create new caption regions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion
type MutableCaptionRegion struct {
	CaptionRegion
}

// MutableCaptionRegionFrom constructs a [MutableCaptionRegion] from an unsafe.Pointer.
//
// A mutable caption region subclass that you use to create new caption regions.
func MutableCaptionRegionFrom(ptr unsafe.Pointer) MutableCaptionRegion {
	return MutableCaptionRegion{
		CaptionRegion: CaptionRegionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MutableCaptionRegion */

// Creates a caption region that has an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/init(identifier:)
func NewMutableCaptionRegionWithIdentifier(identifier objc.IObject /* cross-framework: NSString */) MutableCaptionRegion {
	instance := getMutableCaptionRegionClass().Alloc()
	rv := objc.Send[MutableCaptionRegion](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMutableCaptionRegionWithIdentifier */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MutableCaptionRegion */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MutableCaptionRegion */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MutableCaptionRegion */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MutableCaptionRegion */

// The alignment of lines for the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/displayAlignment
func (m_ MutableCaptionRegion) DisplayAlignment() CaptionRegionDisplayAlignment {
	rv := objc.Send[CaptionRegionDisplayAlignment](m_.ID, objc.Sel("displayAlignment"))
	return rv
}/* debug [instance_properties/getter]: displayAlignment */


// The alignment of lines for the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/displayAlignment
func (m_ MutableCaptionRegion) SetDisplayAlignment(value CaptionRegionDisplayAlignment) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplayAlignment:"), value)
}/* debug [instance_properties/setter]: displayAlignment */


// The region’s top-left position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/origin
func (m_ MutableCaptionRegion) Origin() objc.IObject /* cross-framework: AVCaptionPoint */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("origin"))
	return rv
}/* debug [instance_properties/getter]: origin */


// The region’s top-left position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/origin
func (m_ MutableCaptionRegion) SetOrigin(value objc.IObject /* cross-framework: AVCaptionPoint */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOrigin:"), value)
}/* debug [instance_properties/setter]: origin */


// The scroll mode of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/scroll
func (m_ MutableCaptionRegion) Scroll() CaptionRegionScroll {
	rv := objc.Send[CaptionRegionScroll](m_.ID, objc.Sel("scroll"))
	return rv
}/* debug [instance_properties/getter]: scroll */


// The scroll mode of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/scroll
func (m_ MutableCaptionRegion) SetScroll(value CaptionRegionScroll) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScroll:"), value)
}/* debug [instance_properties/setter]: scroll */


// The height and width of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/size
func (m_ MutableCaptionRegion) Size() objc.IObject /* cross-framework: AVCaptionSize */ {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("size"))
	return rv
}/* debug [instance_properties/getter]: size */


// The height and width of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/size
func (m_ MutableCaptionRegion) SetSize(value objc.IObject /* cross-framework: AVCaptionSize */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}/* debug [instance_properties/setter]: size */


// The block and inline progression direction of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/writingMode
func (m_ MutableCaptionRegion) WritingMode() CaptionRegionWritingMode {
	rv := objc.Send[CaptionRegionWritingMode](m_.ID, objc.Sel("writingMode"))
	return rv
}/* debug [instance_properties/getter]: writingMode */


// The block and inline progression direction of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/writingMode
func (m_ MutableCaptionRegion) SetWritingMode(value CaptionRegionWritingMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWritingMode:"), value)
}/* debug [instance_properties/setter]: writingMode */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVMutableCaptionRegion */


