// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetVariantAudioRenditionSpecificAttributes] class.
var (
	AssetVariantAudioRenditionSpecificAttributesClass     _AssetVariantAudioRenditionSpecificAttributesClass
	AssetVariantAudioRenditionSpecificAttributesClassOnce sync.Once
)

func getAssetVariantAudioRenditionSpecificAttributesClass() _AssetVariantAudioRenditionSpecificAttributesClass {
	AssetVariantAudioRenditionSpecificAttributesClassOnce.Do(func() {
		AssetVariantAudioRenditionSpecificAttributesClass = _AssetVariantAudioRenditionSpecificAttributesClass{objc.GetClass("AVAssetVariantAudioRenditionSpecificAttributes")}
	})
	return AssetVariantAudioRenditionSpecificAttributesClass
}

type _AssetVariantAudioRenditionSpecificAttributesClass struct {
	class objc.Class
}





// An interface definition for the [AssetVariantAudioRenditionSpecificAttributes] class.
type IAssetVariantAudioRenditionSpecificAttributes interface {
	objectivec.IObject
	

	// properties:
	ChannelCount() int
	Binaural() bool
	Downmix() bool
	Immersive() bool
	IsBinaural() bool
	SetIsBinaural(value bool)
	IsDownmix() bool
	SetIsDownmix(value bool)
	IsImmersive() bool
	SetIsImmersive(value bool)
	FormatIDs() objectivec.IObject
	SetFormatIDs(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetVariantAudioRenditionSpecificAttributesClass) Alloc() AssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Send[AssetVariantAudioRenditionSpecificAttributes](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetVariantAudioRenditionSpecificAttributesClass) New() AssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Send[AssetVariantAudioRenditionSpecificAttributes](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetVariantAudioRenditionSpecificAttributes) Init() AssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Send[AssetVariantAudioRenditionSpecificAttributes](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetVariantAudioRenditionSpecificAttributes) Autorelease() AssetVariantAudioRenditionSpecificAttributes {
	rv := objc.Send[AssetVariantAudioRenditionSpecificAttributes](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetVariantAudioRenditionSpecificAttributes creates a new AssetVariantAudioRenditionSpecificAttributes instance.
func NewAssetVariantAudioRenditionSpecificAttributes() AssetVariantAudioRenditionSpecificAttributes {
	return getAssetVariantAudioRenditionSpecificAttributesClass().New()
}





// An object that represents attributes specific to a particular rendition.


// An object that represents attributes specific to a particular rendition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/RenditionSpecificAttributes
type AssetVariantAudioRenditionSpecificAttributes struct {
	objectivec.Object
}

// AssetVariantAudioRenditionSpecificAttributesFrom constructs a [AssetVariantAudioRenditionSpecificAttributes] from an unsafe.Pointer.
//
// An object that represents attributes specific to a particular rendition.
func AssetVariantAudioRenditionSpecificAttributesFrom(ptr unsafe.Pointer) AssetVariantAudioRenditionSpecificAttributes {
	return AssetVariantAudioRenditionSpecificAttributes{objectivec.Object{objc.ID(ptr)}}
}

























// The count of audio channels in the rendition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariantAudioRenditionSpecificAttributes/channelCount
func (a_ AssetVariantAudioRenditionSpecificAttributes) ChannelCount() int {
	rv := objc.Send[int](a_.ID, objc.Sel("channelCount"))
	return rv
}


// A Boolean value that indicates the variant is best suited for delivery to headphones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/RenditionSpecificAttributes/isBinaural
func (a_ AssetVariantAudioRenditionSpecificAttributes) Binaural() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("binaural"))
	return rv
}


// A Boolean value that indicates whether the variant is a downmix derivative of other media of greater channel count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/RenditionSpecificAttributes/isDownmix
func (a_ AssetVariantAudioRenditionSpecificAttributes) Downmix() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("downmix"))
	return rv
}


// A Boolean value that indicates whether this variant contains virtualized or otherwise preprocessed audio content suitable for various purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/AudioAttributes-swift.class/RenditionSpecificAttributes/isImmersive
func (a_ AssetVariantAudioRenditionSpecificAttributes) Immersive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("immersive"))
	return rv
}


// A Boolean value that indicates the variant is best suited for delivery to headphones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.class/renditionspecificattributes/isbinaural
func (a_ AssetVariantAudioRenditionSpecificAttributes) IsBinaural() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isBinaural"))
	return rv
}


// A Boolean value that indicates the variant is best suited for delivery to headphones.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.class/renditionspecificattributes/isbinaural
func (a_ AssetVariantAudioRenditionSpecificAttributes) SetIsBinaural(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsBinaural:"), value)
}


// A Boolean value that indicates whether the variant is a downmix derivative of other media of greater channel count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.class/renditionspecificattributes/isdownmix
func (a_ AssetVariantAudioRenditionSpecificAttributes) IsDownmix() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isDownmix"))
	return rv
}


// A Boolean value that indicates whether the variant is a downmix derivative of other media of greater channel count.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.class/renditionspecificattributes/isdownmix
func (a_ AssetVariantAudioRenditionSpecificAttributes) SetIsDownmix(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsDownmix:"), value)
}


// A Boolean value that indicates whether this variant contains virtualized or otherwise preprocessed audio content suitable for various purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.class/renditionspecificattributes/isimmersive
func (a_ AssetVariantAudioRenditionSpecificAttributes) IsImmersive() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isImmersive"))
	return rv
}


// A Boolean value that indicates whether this variant contains virtualized or otherwise preprocessed audio content suitable for various purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.class/renditionspecificattributes/isimmersive
func (a_ AssetVariantAudioRenditionSpecificAttributes) SetIsImmersive(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsImmersive:"), value)
}


// The audio formats of the renditions present in the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.class/formatids
func (a_ AssetVariantAudioRenditionSpecificAttributes) FormatIDs() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("formatIDs"))
	return rv
}


// The audio formats of the renditions present in the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetvariant/audioattributes-swift.class/formatids
func (a_ AssetVariantAudioRenditionSpecificAttributes) SetFormatIDs(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFormatIDs:"), value)
}








