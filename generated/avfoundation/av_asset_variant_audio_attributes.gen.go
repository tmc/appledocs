// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetVariantAudioAttributes] class.
var (
	AssetVariantAudioAttributesClass     _AssetVariantAudioAttributesClass
	AssetVariantAudioAttributesClassOnce sync.Once
)

func getAssetVariantAudioAttributesClass() _AssetVariantAudioAttributesClass {
	AssetVariantAudioAttributesClassOnce.Do(func() {
		AssetVariantAudioAttributesClass = _AssetVariantAudioAttributesClass{objc.GetClass("AVAssetVariantAudioAttributes")}
	})
	return AssetVariantAudioAttributesClass
}

type _AssetVariantAudioAttributesClass struct {
	class objc.Class
}





// An interface definition for the [AssetVariantAudioAttributes] class.
type IAssetVariantAudioAttributes interface {
	objectivec.IObject
	

	// properties:
	FormatIDs() []foundation.Number
	AudioAttributes() IAVAssetVariantAudioAttributes
	SetAudioAttributes(value IAVAssetVariantAudioAttributes)
	VideoAttributes() IAVAssetVariantVideoAttributes
	SetVideoAttributes(value IAVAssetVariantVideoAttributes)


	

	// methods:
	RenditionSpecificAttributesForMediaOption(mediaSelectionOption IAVMediaSelectionOption) IAssetVariantAudioRenditionSpecificAttributes


}





// Alloc allocates a new instance without initialization.
func (ac _AssetVariantAudioAttributesClass) Alloc() AssetVariantAudioAttributes {
	rv := objc.Send[AssetVariantAudioAttributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetVariantAudioAttributesClass) New() AssetVariantAudioAttributes {
	rv := objc.Send[AssetVariantAudioAttributes](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetVariantAudioAttributes) Init() AssetVariantAudioAttributes {
	rv := objc.Send[AssetVariantAudioAttributes](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetVariantAudioAttributes) Autorelease() AssetVariantAudioAttributes {
	rv := objc.Send[AssetVariantAudioAttributes](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetVariantAudioAttributes creates a new AssetVariantAudioAttributes instance.
func NewAssetVariantAudioAttributes() AssetVariantAudioAttributes {
	return getAssetVariantAudioAttributesClass().New()
}





// An object that defines the audio attributes for an asset variant.


// An object that defines the audio attributes for an asset variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class
type AssetVariantAudioAttributes struct {
	objectivec.Object
}

// AssetVariantAudioAttributesFrom constructs a [AssetVariantAudioAttributes] from an unsafe.Pointer.
//
// An object that defines the audio attributes for an asset variant.
func AssetVariantAudioAttributesFrom(ptr unsafe.Pointer) AssetVariantAudioAttributes {
	return AssetVariantAudioAttributes{objectivec.Object{objc.ID(ptr)}}
}




















// Returns specific attributes for the media option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/renditionSpecificAttributes(for:)
func (a_ AssetVariantAudioAttributes) RenditionSpecificAttributesForMediaOption(mediaSelectionOption IAVMediaSelectionOption) IAssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Send[AssetVariantAudioRenditionSpecificAttributes](a_.ID, objc.Sel("renditionSpecificAttributesForMediaOption:"), mediaSelectionOption)
	return rv
}







// The audio formats of the renditions present in the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantAudioAttributes/formatIDs
func (a_ AssetVariantAudioAttributes) FormatIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("formatIDs"))
	return rv
}


// The audio rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.property
func (a_ AssetVariantAudioAttributes) AudioAttributes() IAVAssetVariantAudioAttributes {
	rv := objc.Send[AssetVariantAudioAttributes](a_.ID, objc.Sel("audioAttributes"))
	return rv
}


// The audio rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.property
func (a_ AssetVariantAudioAttributes) SetAudioAttributes(value IAVAssetVariantAudioAttributes) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAudioAttributes:"), value)
}


// The video rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.property
func (a_ AssetVariantAudioAttributes) VideoAttributes() IAVAssetVariantVideoAttributes {
	rv := objc.Send[AssetVariantVideoAttributes](a_.ID, objc.Sel("videoAttributes"))
	return rv
}


// The video rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/videoattributes-swift.property
func (a_ AssetVariantAudioAttributes) SetVideoAttributes(value IAVAssetVariantVideoAttributes) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setVideoAttributes:"), value)
}








