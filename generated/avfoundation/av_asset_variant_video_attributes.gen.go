// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetVariantVideoAttributes] class.
var (
	AssetVariantVideoAttributesClass     _AssetVariantVideoAttributesClass
	AssetVariantVideoAttributesClassOnce sync.Once
)

func getAssetVariantVideoAttributesClass() _AssetVariantVideoAttributesClass {
	AssetVariantVideoAttributesClassOnce.Do(func() {
		AssetVariantVideoAttributesClass = _AssetVariantVideoAttributesClass{objc.GetClass("AVAssetVariantVideoAttributes")}
	})
	return AssetVariantVideoAttributesClass
}

type _AssetVariantVideoAttributesClass struct {
	class objc.Class
}





// An interface definition for the [AssetVariantVideoAttributes] class.
type IAssetVariantVideoAttributes interface {
	objectivec.IObject
	

	// properties:
	CodecTypes() []foundation.Number
	NominalFrameRate() float64
	PresentationSize() corefoundation.CGSize
	VideoLayoutAttributes() []AssetVariantVideoLayoutAttributes
	VideoRange() VideoRange
	AudioAttributes() IAVAssetVariantAudioAttributes
	SetAudioAttributes(value IAVAssetVariantAudioAttributes)
	VideoAttributes() IAVAssetVariantVideoAttributes
	SetVideoAttributes(value IAVAssetVariantVideoAttributes)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetVariantVideoAttributesClass) Alloc() AssetVariantVideoAttributes {
	rv := objc.Send[AssetVariantVideoAttributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetVariantVideoAttributesClass) New() AssetVariantVideoAttributes {
	rv := objc.Send[AssetVariantVideoAttributes](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetVariantVideoAttributes) Init() AssetVariantVideoAttributes {
	rv := objc.Send[AssetVariantVideoAttributes](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetVariantVideoAttributes) Autorelease() AssetVariantVideoAttributes {
	rv := objc.Send[AssetVariantVideoAttributes](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetVariantVideoAttributes creates a new AssetVariantVideoAttributes instance.
func NewAssetVariantVideoAttributes() AssetVariantVideoAttributes {
	return getAssetVariantVideoAttributesClass().New()
}





// An object that defines the video attributes for an asset variant.


// An object that defines the video attributes for an asset variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class
type AssetVariantVideoAttributes struct {
	objectivec.Object
}

// AssetVariantVideoAttributesFrom constructs a [AssetVariantVideoAttributes] from an unsafe.Pointer.
//
// An object that defines the video attributes for an asset variant.
func AssetVariantVideoAttributesFrom(ptr unsafe.Pointer) AssetVariantVideoAttributes {
	return AssetVariantVideoAttributes{objectivec.Object{objc.ID(ptr)}}
}

























// The video sample codec types present in the variant’s renditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantVideoAttributes/codecTypes
func (a_ AssetVariantVideoAttributes) CodecTypes() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("codecTypes"))
	return rv
}


// The nominal frame rate of the variant’s renditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantVideoAttributes/nominalFrameRate
func (a_ AssetVariantVideoAttributes) NominalFrameRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("nominalFrameRate"))
	return rv
}


// The presentation size of the variant’s renditions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/presentationSize
func (a_ AssetVariantVideoAttributes) PresentationSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a_.ID, objc.Sel("presentationSize"))
	return rv
}


// Attributes that describe the layout of the video content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/videoLayoutAttributes
func (a_ AssetVariantVideoAttributes) VideoLayoutAttributes() []AssetVariantVideoLayoutAttributes {
	rv := objc.Send[[]AssetVariantVideoLayoutAttributes](a_.ID, objc.Sel("videoLayoutAttributes"))
	return rv
}


// The video range of the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/VideoAttributes-swift.class/videoRange
func (a_ AssetVariantVideoAttributes) VideoRange() VideoRange {
	rv := objc.Send[VideoRange](a_.ID, objc.Sel("videoRange"))
	return rv
}


// The audio rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.property
func (a_ AssetVariantVideoAttributes) AudioAttributes() IAVAssetVariantAudioAttributes {
	rv := objc.Send[AssetVariantAudioAttributes](a_.ID, objc.Sel("audioAttributes"))
	return rv
}


// The audio rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.property
func (a_ AssetVariantVideoAttributes) SetAudioAttributes(value IAVAssetVariantAudioAttributes) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioAttributes:"), value)
}


// The video rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.property
func (a_ AssetVariantVideoAttributes) VideoAttributes() IAVAssetVariantVideoAttributes {
	rv := objc.Send[AssetVariantVideoAttributes](a_.ID, objc.Sel("videoAttributes"))
	return rv
}


// The video rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.property
func (a_ AssetVariantVideoAttributes) SetVideoAttributes(value IAVAssetVariantVideoAttributes) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoAttributes:"), value)
}








