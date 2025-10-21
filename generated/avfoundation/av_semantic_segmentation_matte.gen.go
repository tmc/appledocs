// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [SemanticSegmentationMatte] class.
type ISemanticSegmentationMatte interface {
	objectivec.IObject
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType appkit.string) foundation.Dictionary
}

// An object that wraps a matting image for a particular semantic segmentation.
//
// The matting image stores its pixel data as objects in format. The image file contains the semantic segmentation matte as an auxiliary image, accessible using the ImageIO framework’s function.
//
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

// Alloc allocates a new instance without initialization.
func (sc _SemanticSegmentationMatteClass) Alloc() SemanticSegmentationMatte {
	rv := objc.Send[SemanticSegmentationMatte](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns a dictionary of primitive map information to use when writing an image file with a semantic segmentation matte.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/dictionaryRepresentation(forAuxiliaryDataType:)
func (s_ SemanticSegmentationMatte) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType appkit.string) foundation.Dictionary {
	rv := objc.Send[foundation.Dictionary](s_.ID, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), outAuxDataType)
	return rv
}

// The semantic segmentation matte image type.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/matteType-swift.property
func (s_ SemanticSegmentationMatte) MatteType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("matteType"))
	return rv
}

// The pixel format type for this object’s internal matting image.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/pixelFormatType
func (s_ SemanticSegmentationMatte) PixelFormatType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("pixelFormatType"))
	return rv
}

// The semantic segmentation matte’s internal image.
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte/mattingimage
func (s_ SemanticSegmentationMatte) MattingImage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("mattingImage"))
	return rv
}


// SetMattingImage sets the value of the mattingImage property.
// The semantic segmentation matte’s internal image.

//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avsemanticsegmentationmatte/mattingimage
func (s_ SemanticSegmentationMatte) SetMattingImage(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMattingImage:"), value)
}

// 8-bit one component, black is zero.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent8
func (s_ SemanticSegmentationMatte) KCVPixelFormatType_OneComponent8() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("kCVPixelFormatType_OneComponent8"))
	return rv
}


// SetKCVPixelFormatType_OneComponent8 sets the value of the kCVPixelFormatType_OneComponent8 property.
// 8-bit one component, black is zero.

//
// [Full Topic]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent8
func (s_ SemanticSegmentationMatte) SetKCVPixelFormatType_OneComponent8(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setKCVPixelFormatType_OneComponent8:"), value)
}



