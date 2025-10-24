// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [PortraitEffectsMatte] class.
var (
	PortraitEffectsMatteClass     _PortraitEffectsMatteClass
	PortraitEffectsMatteClassOnce sync.Once
)

func getPortraitEffectsMatteClass() _PortraitEffectsMatteClass {
	PortraitEffectsMatteClassOnce.Do(func() {
		PortraitEffectsMatteClass = _PortraitEffectsMatteClass{objc.GetClass("AVPortraitEffectsMatte")}
	})
	return PortraitEffectsMatteClass
}

type _PortraitEffectsMatteClass struct {
	class objc.Class
}





// An interface definition for the [PortraitEffectsMatte] class.
type IPortraitEffectsMatte interface {
	objectivec.IObject
	

	// properties:
	MattingImage() PixelBufferRef /* not a class type */
	PixelFormatType() uint32 /* not a class type */


	

	// methods:
	PortraitEffectsMatteByApplyingExifOrientation(exifOrientation ImagePropertyOrientation /* not a class type */) objectivec.IObject
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType objc.IObject /* cross-framework: NSString */) foundation.Dictionary
	PortraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBufferError(pixelBuffer PixelBufferRef /* not a class type */, outError objectivec.IObject) objectivec.IObject


}





// Alloc allocates a new instance without initialization.
func (pc _PortraitEffectsMatteClass) Alloc() PortraitEffectsMatte {
	rv := objc.Send[PortraitEffectsMatte](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PortraitEffectsMatteClass) New() PortraitEffectsMatte {
	rv := objc.Send[PortraitEffectsMatte](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PortraitEffectsMatte) Init() PortraitEffectsMatte {
	rv := objc.Send[PortraitEffectsMatte](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PortraitEffectsMatte) Autorelease() PortraitEffectsMatte {
	rv := objc.Send[PortraitEffectsMatte](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPortraitEffectsMatte creates a new PortraitEffectsMatte instance.
func NewPortraitEffectsMatte() PortraitEffectsMatte {
	return getPortraitEffectsMatteClass().New()
}





// An auxiliary image used to separate foreground from background with high resolution.
//
// Before iOS 11, the iPhone camera software used depth maps to render a shallow depth of field (the effect) into still images taken in Portrait Mode before discarding the maps. Because the effect was part of the photo, you couldn’t access the maps separately, as metadata, for photos taken by devices running iOS 10 or earlier. Starting in iOS 11, apps accessing the photo library can use images containing embedded auxiliary depth maps to render creative depth effects, such as forced perspective, or image projection from 2D to 3D space. These depth maps are low-resolution compared to the full-resolution RGB image. As such, the depth effects you can render are limited by the resolution and accuracy of the maps. Fine detail, such as hair, is challenging to preserve faithfully at the resolution of these depth maps. Starting in iOS 12, the portrait effects matte helps achieve this fine-grain level of detail. Using the auxiliary matte image, you can improve the quality of rendered portrait effects, such as Natural Light, Studio Light, Contour Light, Stage Light, and Stage Light Mono. Unlike the depth map, the portrait effects matte isn’t intended to faithfully preserve all gradations of depth in the scene. It’s a depth-guided, people-focused segmentation mask generated from a proprietary Apple neural network trained to detect people. It separates an individual in the foreground from whatever is in the background, with greater detail and clarity than with the depth map alone. It achieves this clarity in part because the matte image has higher resolution than the depth map. For information about capturing the portrait effects matte, see . To learn how to extract a portrait effects matte from photos previously captured in portrait mode on a device running iOS 12, see .


// An auxiliary image used to separate foreground from background with high resolution.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte
type PortraitEffectsMatte struct {
	objectivec.Object
}

// PortraitEffectsMatteFrom constructs a [PortraitEffectsMatte] from an unsafe.Pointer.
//
// An auxiliary image used to separate foreground from background with high resolution.
func PortraitEffectsMatteFrom(ptr unsafe.Pointer) PortraitEffectsMatte {
	return PortraitEffectsMatte{objectivec.Object{objc.ID(ptr)}}
}






// Initializes a portrait effects matte instance from auxiliary image information in an image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/init(fromDictionaryRepresentation:)
func NewPortraitEffectsMatteFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary objc.IObject /* cross-framework: NSDictionary */, outError objectivec.IObject) PortraitEffectsMatte {
	rv := objc.Send[PortraitEffectsMatte](objc.ID(getPortraitEffectsMatteClass().class), objc.Sel("portraitEffectsMatteFromDictionaryRepresentation:error:"), imageSourceAuxDataInfoDictionary, outError)
	return rv
}







// Initializes a portrait effects matte instance from auxiliary image information in an image file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/init(fromDictionaryRepresentation:)
func (pc _PortraitEffectsMatteClass) PortraitEffectsMatteFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary objc.IObject /* cross-framework: NSDictionary */, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("portraitEffectsMatteFromDictionaryRepresentation:error:"), imageSourceAuxDataInfoDictionary, outError)
	return rv
}












// Returns a derivative portrait effects matte after applying the specified EXIF orientation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/applyingExifOrientation(_:)
func (p_ PortraitEffectsMatte) PortraitEffectsMatteByApplyingExifOrientation(exifOrientation ImagePropertyOrientation /* not a class type */) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("portraitEffectsMatteByApplyingExifOrientation:"), exifOrientation)
	return rv
}


// A dictionary of primitive map information used for writing an image file with a portrait effects matte.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/dictionaryRepresentation(forAuxiliaryDataType:)
func (p_ PortraitEffectsMatte) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType objc.IObject /* cross-framework: NSString */) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](p_.ID, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), outAuxDataType)
	return rv
}


// Returns a portrait effects matte by wrapping the replacement pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/replacingPortraitEffectsMatte(with:)
func (p_ PortraitEffectsMatte) PortraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBufferError(pixelBuffer PixelBufferRef /* not a class type */, outError objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](p_.ID, objc.Sel("portraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBuffer:error:"), pixelBuffer, outError)
	return rv
}







// The portrait effects matte’s internal image, formatted as a pixel buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/mattingImage
func (p_ PortraitEffectsMatte) MattingImage() PixelBufferRef /* not a class type */ {
	rv := objc.Send[PixelBufferRef](p_.ID, objc.Sel("mattingImage"))
	return rv
}


// The pixel format type of this portrait effects matte’s internal image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/pixelFormatType
func (p_ PortraitEffectsMatte) PixelFormatType() uint32 /* not a class type */ {
	rv := objc.Send[uint32](p_.ID, objc.Sel("pixelFormatType"))
	return rv
}







