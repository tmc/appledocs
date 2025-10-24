// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptionFormatConformer] class.
var (
	CaptionFormatConformerClass     _CaptionFormatConformerClass
	CaptionFormatConformerClassOnce sync.Once
)

func getCaptionFormatConformerClass() _CaptionFormatConformerClass {
	CaptionFormatConformerClassOnce.Do(func() {
		CaptionFormatConformerClass = _CaptionFormatConformerClass{objc.GetClass("AVCaptionFormatConformer")}
	})
	return CaptionFormatConformerClass
}

type _CaptionFormatConformerClass struct {
	class objc.Class
}





// An interface definition for the [CaptionFormatConformer] class.
type ICaptionFormatConformer interface {
	objectivec.IObject
	

	// properties:
	ConformsCaptionsToTimeRange() bool
	SetConformsCaptionsToTimeRange(value bool)


	

	// methods:
	ConformedCaptionForCaptionError(caption IAVCaption, outError objectivec.IObject) ICaption


}





// Alloc allocates a new instance without initialization.
func (cc _CaptionFormatConformerClass) Alloc() CaptionFormatConformer {
	rv := objc.Send[CaptionFormatConformer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionFormatConformerClass) New() CaptionFormatConformer {
	rv := objc.Send[CaptionFormatConformer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionFormatConformer) Init() CaptionFormatConformer {
	rv := objc.Send[CaptionFormatConformer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionFormatConformer) Autorelease() CaptionFormatConformer {
	rv := objc.Send[CaptionFormatConformer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionFormatConformer creates a new CaptionFormatConformer instance.
func NewCaptionFormatConformer() CaptionFormatConformer {
	return getCaptionFormatConformerClass().New()
}





// An object that converts a canonical caption to a specific format.


// An object that converts a canonical caption to a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer
type CaptionFormatConformer struct {
	objectivec.Object
}

// CaptionFormatConformerFrom constructs a [CaptionFormatConformer] from an unsafe.Pointer.
//
// An object that converts a canonical caption to a specific format.
func CaptionFormatConformerFrom(ptr unsafe.Pointer) CaptionFormatConformer {
	return CaptionFormatConformer{objectivec.Object{objc.ID(ptr)}}
}






// Creates a new object with format conversion settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/init(conversionSettings:)
func NewCaptionFormatConformerWithConversionSettings(conversionSettings foundation.IDictionary) CaptionFormatConformer {
	instance := getCaptionFormatConformerClass().Alloc()
	rv := objc.Send[CaptionFormatConformer](instance.ID, objc.Sel("initWithConversionSettings:"), conversionSettings)
	rv.Autorelease()
	return rv
}







// A class method that creates a new object with format conversion settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/captionFormatConformerWithConversionSettings:
func (cc _CaptionFormatConformerClass) CaptionFormatConformerWithConversionSettings(conversionSettings foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("captionFormatConformerWithConversionSettings:"), conversionSettings)
	return rv
}












// Creates a caption that conforms to a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/conformedCaption(for:)
func (c_ CaptionFormatConformer) ConformedCaptionForCaptionError(caption IAVCaption, outError objectivec.IObject) ICaption {
	rv := objc.Send[Caption](c_.ID, objc.Sel("conformedCaptionForCaption:error:"), caption, outError)
	return rv
}







// A Boolean value that indicates whether to conform the time range of a canonical caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/conformsCaptionsToTimeRange
func (c_ CaptionFormatConformer) ConformsCaptionsToTimeRange() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("conformsCaptionsToTimeRange"))
	return rv
}


// A Boolean value that indicates whether to conform the time range of a canonical caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/conformsCaptionsToTimeRange
func (c_ CaptionFormatConformer) SetConformsCaptionsToTimeRange(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConformsCaptionsToTimeRange:"), value)
}







