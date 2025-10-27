// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetVariant] class.
var (
	AssetVariantClass     _AssetVariantClass
	AssetVariantClassOnce sync.Once
)

func getAssetVariantClass() _AssetVariantClass {
	AssetVariantClassOnce.Do(func() {
		AssetVariantClass = _AssetVariantClass{objc.GetClass("AVAssetVariant")}
	})
	return AssetVariantClass
}

type _AssetVariantClass struct {
	class objc.Class
}





// An interface definition for the [AssetVariant] class.
type IAssetVariant interface {
	objectivec.IObject
	

	// properties:
	AudioAttributes() IAVAssetVariantAudioAttributes
	AverageBitRate() float64
	PeakBitRate() float64
	URL() foundation.foundation.INSURL
	VideoAttributes() IAVAssetVariantVideoAttributes


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetVariantClass) Alloc() AssetVariant {
	rv := objc.Send[AssetVariant](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetVariantClass) New() AssetVariant {
	rv := objc.Send[AssetVariant](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetVariant) Init() AssetVariant {
	rv := objc.Send[AssetVariant](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetVariant) Autorelease() AssetVariant {
	rv := objc.Send[AssetVariant](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetVariant creates a new AssetVariant instance.
func NewAssetVariant() AssetVariant {
	return getAssetVariantClass().New()
}





// An object that represents a bit rate variant.


// An object that represents a bit rate variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant
type AssetVariant struct {
	objectivec.Object
}

// AssetVariantFrom constructs a [AssetVariant] from an unsafe.Pointer.
//
// An object that represents a bit rate variant.
func AssetVariantFrom(ptr unsafe.Pointer) AssetVariant {
	return AssetVariant{objectivec.Object{objc.ID(ptr)}}
}

























// The audio rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/audioAttributes-swift.property
func (a_ AssetVariant) AudioAttributes() IAVAssetVariantAudioAttributes {
	rv := objc.Send[AssetVariantAudioAttributes](a_.ID, objc.Sel("audioAttributes"))
	return rv
}


// The average bit rate for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/averageBitRate-7bnsq
func (a_ AssetVariant) AverageBitRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("averageBitRate"))
	return rv
}


// The peak bit rate for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/peakBitRate-38p2b
func (a_ AssetVariant) PeakBitRate() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("peakBitRate"))
	return rv
}


// Provides URL to media playlist corresponding to variant
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/url
func (a_ AssetVariant) URL() foundation.foundation.INSURL {
	rv := objc.Send[foundation.NSURL](a_.ID, objc.Sel("URL"))
	return rv
}


// The video rendition attributes for the variant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetVariant/videoAttributes-swift.property
func (a_ AssetVariant) VideoAttributes() IAVAssetVariantVideoAttributes {
	rv := objc.Send[AssetVariantVideoAttributes](a_.ID, objc.Sel("videoAttributes"))
	return rv
}








