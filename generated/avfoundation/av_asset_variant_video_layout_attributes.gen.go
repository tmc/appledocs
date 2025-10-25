// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetVariantVideoLayoutAttributes */


/* debug [class_header]: Header for AVAssetVariantVideoLayoutAttributes */
// The class instance for the [AssetVariantVideoLayoutAttributes] class.
var (
	AssetVariantVideoLayoutAttributesClass     _AssetVariantVideoLayoutAttributesClass
	AssetVariantVideoLayoutAttributesClassOnce sync.Once
)

func getAssetVariantVideoLayoutAttributesClass() _AssetVariantVideoLayoutAttributesClass {
	AssetVariantVideoLayoutAttributesClassOnce.Do(func() {
		AssetVariantVideoLayoutAttributesClass = _AssetVariantVideoLayoutAttributesClass{objc.GetClass("AVAssetVariantVideoLayoutAttributes")}
	})
	return AssetVariantVideoLayoutAttributesClass
}

type _AssetVariantVideoLayoutAttributesClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetVariantVideoLayoutAttributes */
// An interface definition for the [AssetVariantVideoLayoutAttributes] class.
type IAssetVariantVideoLayoutAttributes interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetVariantVideoLayoutAttributes */
	// properties:
	ProjectionType() ProjectionType /* not a class type */
	StereoViewComponents() StereoViewComponents /* not a class type */
	CodecTypes() VideoCodecType /* typedef */
	SetCodecTypes(value VideoCodecType /* typedef */)
	NominalFrameRate() float64
	SetNominalFrameRate(value float64)
	PresentationSize() corefoundation.CGSize
	SetPresentationSize(value corefoundation.CGSize)
	VideoLayoutAttributes() IAVAssetVariantVideoLayoutAttributes
	SetVideoLayoutAttributes(value IAVAssetVariantVideoLayoutAttributes)
	VideoRange() VideoRange /* typedef */
	SetVideoRange(value VideoRange /* typedef */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetVariantVideoLayoutAttributes */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetVariantVideoLayoutAttributes */
// Alloc allocates a new instance without initialization.
func (ac _AssetVariantVideoLayoutAttributesClass) Alloc() AssetVariantVideoLayoutAttributes {
	rv := objc.Send[AssetVariantVideoLayoutAttributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetVariantVideoLayoutAttributesClass) New() AssetVariantVideoLayoutAttributes {
	rv := objc.Send[AssetVariantVideoLayoutAttributes](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetVariantVideoLayoutAttributes) Init() AssetVariantVideoLayoutAttributes {
	rv := objc.Send[AssetVariantVideoLayoutAttributes](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetVariantVideoLayoutAttributes) Autorelease() AssetVariantVideoLayoutAttributes {
	rv := objc.Send[AssetVariantVideoLayoutAttributes](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetVariantVideoLayoutAttributes creates a new AssetVariantVideoLayoutAttributes instance.
func NewAssetVariantVideoLayoutAttributes() AssetVariantVideoLayoutAttributes {
	return getAssetVariantVideoLayoutAttributesClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetVariantVideoLayoutAttributes */
// Attributes that describe the layout of video content.


// Attributes that describe the layout of video content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/LayoutAttributes
type AssetVariantVideoLayoutAttributes struct {
	objectivec.Object
}

// AssetVariantVideoLayoutAttributesFrom constructs a [AssetVariantVideoLayoutAttributes] from an unsafe.Pointer.
//
// Attributes that describe the layout of video content.
func AssetVariantVideoLayoutAttributesFrom(ptr unsafe.Pointer) AssetVariantVideoLayoutAttributes {
	return AssetVariantVideoLayoutAttributes{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetVariantVideoLayoutAttributes *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetVariantVideoLayoutAttributes */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetVariantVideoLayoutAttributes */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetVariantVideoLayoutAttributes */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetVariantVideoLayoutAttributes */

// Describes the video projection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/LayoutAttributes/projectionType
func (a_ AssetVariantVideoLayoutAttributes) ProjectionType() ProjectionType /* not a class type */ {
	rv := objc.Send[ProjectionType](a_.ID, objc.Sel("projectionType"))
	return rv
}/* debug [instance_properties/getter]: projectionType */


// Attributes that describe the video’s stereo components.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/LayoutAttributes/stereoViewComponents
func (a_ AssetVariantVideoLayoutAttributes) StereoViewComponents() StereoViewComponents /* not a class type */ {
	rv := objc.Send[StereoViewComponents](a_.ID, objc.Sel("stereoViewComponents"))
	return rv
}/* debug [instance_properties/getter]: stereoViewComponents */


// The video sample codec types present in the variant’s renditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/codectypes
func (a_ AssetVariantVideoLayoutAttributes) CodecTypes() VideoCodecType /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("codecTypes"))
	return rv
}/* debug [instance_properties/getter]: codecTypes */


// The video sample codec types present in the variant’s renditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/codectypes
func (a_ AssetVariantVideoLayoutAttributes) SetCodecTypes(value VideoCodecType /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCodecTypes:"), value)
}/* debug [instance_properties/setter]: codecTypes */


// The nominal frame rate of the variant’s renditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/nominalframerate
func (a_ AssetVariantVideoLayoutAttributes) NominalFrameRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("nominalFrameRate"))
	return rv
}/* debug [instance_properties/getter]: nominalFrameRate */


// The nominal frame rate of the variant’s renditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/nominalframerate
func (a_ AssetVariantVideoLayoutAttributes) SetNominalFrameRate(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setNominalFrameRate:"), value)
}/* debug [instance_properties/setter]: nominalFrameRate */


// The presentation size of the variant’s renditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/presentationsize
func (a_ AssetVariantVideoLayoutAttributes) PresentationSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("presentationSize"))
	return rv
}/* debug [instance_properties/getter]: presentationSize */


// The presentation size of the variant’s renditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/presentationsize
func (a_ AssetVariantVideoLayoutAttributes) SetPresentationSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPresentationSize:"), value)
}/* debug [instance_properties/setter]: presentationSize */


// Attributes that describe the layout of the video content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/videolayoutattributes
func (a_ AssetVariantVideoLayoutAttributes) VideoLayoutAttributes() IAVAssetVariantVideoLayoutAttributes {
	rv := objc.Send[AssetVariantVideoLayoutAttributes](a_.ID, objc.Sel("videoLayoutAttributes"))
	return rv
}/* debug [instance_properties/getter]: videoLayoutAttributes */


// Attributes that describe the layout of the video content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/videolayoutattributes
func (a_ AssetVariantVideoLayoutAttributes) SetVideoLayoutAttributes(value IAVAssetVariantVideoLayoutAttributes) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoLayoutAttributes:"), value)
}/* debug [instance_properties/setter]: videoLayoutAttributes */


// The video range of the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/videorange
func (a_ AssetVariantVideoLayoutAttributes) VideoRange() VideoRange /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("videoRange"))
	return rv
}/* debug [instance_properties/getter]: videoRange */


// The video range of the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.class/videorange
func (a_ AssetVariantVideoLayoutAttributes) SetVideoRange(value VideoRange /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoRange:"), value)
}/* debug [instance_properties/setter]: videoRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetVariantVideoLayoutAttributes */



