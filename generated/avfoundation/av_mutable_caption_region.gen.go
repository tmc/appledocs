// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





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





// An interface definition for the [MutableCaptionRegion] class.
type IMutableCaptionRegion interface {
	ICaptionRegion
	

	// properties:
	DisplayAlignment() CaptionRegionDisplayAlignment
	SetDisplayAlignment(value CaptionRegionDisplayAlignment)
	Origin() AVCaptionPoint
	SetOrigin(value AVCaptionPoint)
	Scroll() CaptionRegionScroll
	SetScroll(value CaptionRegionScroll)
	Size() AVCaptionSize
	SetSize(value AVCaptionSize)
	WritingMode() CaptionRegionWritingMode
	SetWritingMode(value CaptionRegionWritingMode)


	

	// methods:


}





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






// Creates a caption region that has an identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/init(identifier:)
func NewMutableCaptionRegionWithIdentifier(identifier foundation.foundation.INSString) MutableCaptionRegion {
	instance := getMutableCaptionRegionClass().Alloc()
	rv := objc.Send[MutableCaptionRegion](instance.ID, objc.Sel("initWithIdentifier:"), identifier)
	rv.Autorelease()
	return rv
}






















// The alignment of lines for the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/displayAlignment
func (m_ MutableCaptionRegion) DisplayAlignment() CaptionRegionDisplayAlignment {
	rv := objc.Send[CaptionRegionDisplayAlignment](m_.ID, objc.Sel("displayAlignment"))
	return rv
}


// The alignment of lines for the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/displayAlignment
func (m_ MutableCaptionRegion) SetDisplayAlignment(value CaptionRegionDisplayAlignment) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisplayAlignment:"), value)
}


// The region’s top-left position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/origin
func (m_ MutableCaptionRegion) Origin() AVCaptionPoint {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("origin"))
	return rv
}


// The region’s top-left position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/origin
func (m_ MutableCaptionRegion) SetOrigin(value AVCaptionPoint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setOrigin:"), value)
}


// The scroll mode of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/scroll
func (m_ MutableCaptionRegion) Scroll() CaptionRegionScroll {
	rv := objc.Send[CaptionRegionScroll](m_.ID, objc.Sel("scroll"))
	return rv
}


// The scroll mode of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/scroll
func (m_ MutableCaptionRegion) SetScroll(value CaptionRegionScroll) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setScroll:"), value)
}


// The height and width of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/size
func (m_ MutableCaptionRegion) Size() AVCaptionSize {
	rv := objc.Send[objc.ID](m_.ID, objc.Sel("size"))
	return rv
}


// The height and width of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/size
func (m_ MutableCaptionRegion) SetSize(value AVCaptionSize) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSize:"), value)
}


// The block and inline progression direction of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/writingMode
func (m_ MutableCaptionRegion) WritingMode() CaptionRegionWritingMode {
	rv := objc.Send[CaptionRegionWritingMode](m_.ID, objc.Sel("writingMode"))
	return rv
}


// The block and inline progression direction of the region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMutableCaptionRegion/writingMode
func (m_ MutableCaptionRegion) SetWritingMode(value CaptionRegionWritingMode) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setWritingMode:"), value)
}







