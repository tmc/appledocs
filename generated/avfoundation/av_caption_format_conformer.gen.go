// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptionFormatConformer */


/* debug [class_header]: Header for AVCaptionFormatConformer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptionFormatConformer */
// An interface definition for the [CaptionFormatConformer] class.
type ICaptionFormatConformer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptionFormatConformer */
	// properties:
	ConformsCaptionsToTimeRange() bool
	SetConformsCaptionsToTimeRange(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptionFormatConformer */
	// methods:
	ConformedCaptionForCaptionError(caption IAVCaption, outError objectivec.IObject) ICaption
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptionFormatConformer */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptionFormatConformer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptionFormatConformer */

// Creates a new object with format conversion settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/init(conversionSettings:)
func NewCaptionFormatConformerWithConversionSettings(conversionSettings foundation.IDictionary) CaptionFormatConformer {
	instance := getCaptionFormatConformerClass().Alloc()
	rv := objc.Send[CaptionFormatConformer](instance.ID, objc.Sel("initWithConversionSettings:"), conversionSettings)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCaptionFormatConformerWithConversionSettings */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptionFormatConformer */

// A class method that creates a new object with format conversion settings.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/captionFormatConformerWithConversionSettings:
func (cc _CaptionFormatConformerClass) CaptionFormatConformerWithConversionSettings(conversionSettings foundation.IDictionary) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("captionFormatConformerWithConversionSettings:"), conversionSettings)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CaptionFormatConformerWithConversionSettings) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptionFormatConformer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptionFormatConformer */

// Creates a caption that conforms to a specific format.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/conformedCaption(for:)
func (c_ CaptionFormatConformer) ConformedCaptionForCaptionError(caption IAVCaption, outError objectivec.IObject) ICaption {
	rv := objc.Send[Caption](c_.ID, objc.Sel("conformedCaptionForCaption:error:"), caption, outError)
	return rv
}/* debug [instance_methods/method]: ConformedCaptionForCaptionError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptionFormatConformer */

// A Boolean value that indicates whether to conform the time range of a canonical caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/conformsCaptionsToTimeRange
func (c_ CaptionFormatConformer) ConformsCaptionsToTimeRange() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("conformsCaptionsToTimeRange"))
	return rv
}/* debug [instance_properties/getter]: conformsCaptionsToTimeRange */


// A Boolean value that indicates whether to conform the time range of a canonical caption.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionFormatConformer/conformsCaptionsToTimeRange
func (c_ CaptionFormatConformer) SetConformsCaptionsToTimeRange(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setConformsCaptionsToTimeRange:"), value)
}/* debug [instance_properties/setter]: conformsCaptionsToTimeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptionFormatConformer */


