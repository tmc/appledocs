// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVSemanticSegmentationMatte */


/* debug [class_header]: Header for AVSemanticSegmentationMatte */
// The class instance for the [SemanticSegmentationMatte] class.
var (
	SemanticSegmentationMatteClass     _SemanticSegmentationMatteClass
	SemanticSegmentationMatteClassOnce sync.Once
)

func getSemanticSegmentationMatteClass() _SemanticSegmentationMatteClass {
	SemanticSegmentationMatteClassOnce.Do(func() {
		SemanticSegmentationMatteClass = _SemanticSegmentationMatteClass{objc.GetClass("AVSemanticSegmentationMatte")}
	})
	return SemanticSegmentationMatteClass
}

type _SemanticSegmentationMatteClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SemanticSegmentationMatte */
// An interface definition for the [SemanticSegmentationMatte] class.
type ISemanticSegmentationMatte interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SemanticSegmentationMatte */
	// properties:
	MatteType() SemanticSegmentationMatteType /* typedef */
	MattingImage() PixelBufferRef /* not a class type */
	PixelFormatType() uint32 /* not a class type */
	KCVPixelFormatType_OneComponent8() uint32 /* not a class type */
	SetKCVPixelFormatType_OneComponent8(value uint32 /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SemanticSegmentationMatte */
	// methods:
	SemanticSegmentationMatteByApplyingExifOrientation(exifOrientation ImagePropertyOrientation /* not a class type */) objectivec.IObject
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType objc.IObject /* cross-framework: NSString */) foundation.Dictionary
	SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError(pixelBuffer PixelBufferRef /* not a class type */, outError objectivec.IObject) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SemanticSegmentationMatte */
// Alloc allocates a new instance without initialization.
func (sc _SemanticSegmentationMatteClass) Alloc() SemanticSegmentationMatte {
	rv := objc.Send[SemanticSegmentationMatte](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SemanticSegmentationMatteClass) New() SemanticSegmentationMatte {
	rv := objc.Send[SemanticSegmentationMatte](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SemanticSegmentationMatte) Init() SemanticSegmentationMatte {
	rv := objc.Send[SemanticSegmentationMatte](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SemanticSegmentationMatte) Autorelease() SemanticSegmentationMatte {
	rv := objc.Send[SemanticSegmentationMatte](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSemanticSegmentationMatte creates a new SemanticSegmentationMatte instance.
func NewSemanticSegmentationMatte() SemanticSegmentationMatte {
	return getSemanticSegmentationMatteClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SemanticSegmentationMatte */
// An object that wraps a matting image for a particular semantic segmentation.
//
// The matting image stores its pixel data as objects in format. The image file contains the semantic segmentation matte as an auxiliary image, accessible using the ImageIO framework’s function.


// An object that wraps a matting image for a particular semantic segmentation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte
type SemanticSegmentationMatte struct {
	objectivec.Object
}

// SemanticSegmentationMatteFrom constructs a [SemanticSegmentationMatte] from an unsafe.Pointer.
//
// An object that wraps a matting image for a particular semantic segmentation.
func SemanticSegmentationMatteFrom(ptr unsafe.Pointer) SemanticSegmentationMatte {
	return SemanticSegmentationMatte{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SemanticSegmentationMatte */

// Returns a new semantic segmentation matte instance from auxiliary image information in an image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/init(fromImageSourceAuxiliaryDataType:dictionaryRepresentation:)
func NewSemanticSegmentationMatteFromImageSourceAuxiliaryDataTypeDictionaryRepresentationError(imageSourceAuxiliaryDataType StringRef /* not a class type */, imageSourceAuxiliaryDataInfoDictionary objc.IObject /* cross-framework: NSDictionary */, outError objectivec.IObject) SemanticSegmentationMatte {
	rv := objc.Send[SemanticSegmentationMatte](objc.ID(getSemanticSegmentationMatteClass().class), objc.Sel("semanticSegmentationMatteFromImageSourceAuxiliaryDataType:dictionaryRepresentation:error:"), imageSourceAuxiliaryDataType, imageSourceAuxiliaryDataInfoDictionary, outError)
	return rv
}/* debug [class_init_methods/constructor]: NewSemanticSegmentationMatteFromImageSourceAuxiliaryDataTypeDictionaryRepresentationError */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SemanticSegmentationMatte */

// Returns a new semantic segmentation matte instance from auxiliary image information in an image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/init(fromImageSourceAuxiliaryDataType:dictionaryRepresentation:)
func (sc _SemanticSegmentationMatteClass) SemanticSegmentationMatteFromImageSourceAuxiliaryDataTypeDictionaryRepresentationError(imageSourceAuxiliaryDataType StringRef /* not a class type */, imageSourceAuxiliaryDataInfoDictionary objc.IObject /* cross-framework: NSDictionary */, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(sc.class), objc.Sel("semanticSegmentationMatteFromImageSourceAuxiliaryDataType:dictionaryRepresentation:error:"), imageSourceAuxiliaryDataType, imageSourceAuxiliaryDataInfoDictionary, outError)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SemanticSegmentationMatteFromImageSourceAuxiliaryDataTypeDictionaryRepresentationError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SemanticSegmentationMatte */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SemanticSegmentationMatte */

// Returns a new semantic segmentation matte instance with the specified Exif orientation applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/applyingExifOrientation(_:)
func (s_ SemanticSegmentationMatte) SemanticSegmentationMatteByApplyingExifOrientation(exifOrientation ImagePropertyOrientation /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("semanticSegmentationMatteByApplyingExifOrientation:"), exifOrientation)
	return rv
}/* debug [instance_methods/method]: SemanticSegmentationMatteByApplyingExifOrientation */


// Returns a dictionary of primitive map information to use when writing an image file with a semantic segmentation matte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/dictionaryRepresentation(forAuxiliaryDataType:)
func (s_ SemanticSegmentationMatte) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType objc.IObject /* cross-framework: NSString */) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](s_.ID, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), outAuxDataType)
	return rv
}/* debug [instance_methods/method]: DictionaryRepresentationForAuxiliaryDataType */


// Returns a semantic segmentation matte instance that wraps the replacement pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/replacingSemanticSegmentationMatte(with:)
func (s_ SemanticSegmentationMatte) SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError(pixelBuffer PixelBufferRef /* not a class type */, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("semanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBuffer:error:"), pixelBuffer, outError)
	return rv
}/* debug [instance_methods/method]: SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SemanticSegmentationMatte */

// The semantic segmentation matte image type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/matteType-swift.property
func (s_ SemanticSegmentationMatte) MatteType() SemanticSegmentationMatteType /* typedef */ {
	rv := objc.Send[foundation.NSString](s_.ID, objc.Sel("matteType"))
	return rv
}/* debug [instance_properties/getter]: matteType */


// The semantic segmentation matte’s internal image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/mattingImage
func (s_ SemanticSegmentationMatte) MattingImage() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](s_.ID, objc.Sel("mattingImage"))
	return rv
}/* debug [instance_properties/getter]: mattingImage */


// The pixel format type for this object’s internal matting image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/pixelFormatType
func (s_ SemanticSegmentationMatte) PixelFormatType() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("pixelFormatType"))
	return rv
}/* debug [instance_properties/getter]: pixelFormatType */


// 8-bit one component, black is zero.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent8
func (s_ SemanticSegmentationMatte) KCVPixelFormatType_OneComponent8() uint32 /* not a class type */ {
	rv := objc.Send[uint32](s_.ID, objc.Sel("kCVPixelFormatType_OneComponent8"))
	return rv
}/* debug [instance_properties/getter]: kCVPixelFormatType_OneComponent8 */


// 8-bit one component, black is zero.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent8
func (s_ SemanticSegmentationMatte) SetKCVPixelFormatType_OneComponent8(value uint32 /* not a class type */) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKCVPixelFormatType_OneComponent8:"), value)
}/* debug [instance_properties/setter]: kCVPixelFormatType_OneComponent8 */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVSemanticSegmentationMatte */


