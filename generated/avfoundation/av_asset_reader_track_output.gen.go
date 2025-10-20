// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AssetReaderTrackOutput] class.
var (
	AssetReaderTrackOutputClass     _AssetReaderTrackOutputClass
	AssetReaderTrackOutputClassOnce sync.Once
)

func getAssetReaderTrackOutputClass() _AssetReaderTrackOutputClass {
	AssetReaderTrackOutputClassOnce.Do(func() {
		AssetReaderTrackOutputClass = _AssetReaderTrackOutputClass{objc.GetClass("AVAssetReaderTrackOutput")}
	})
	return AssetReaderTrackOutputClass
}

type _AssetReaderTrackOutputClass struct {
	class objc.Class
}

// An interface definition for the [AssetReaderTrackOutput] class.
type IAssetReaderTrackOutput interface {
	IAssetReaderOutput
}

// An object that reads media data from a single track of an asset.
//
// Read the media data of an asset track by adding a track output to an asset reader. You can read the media samples in their stored format, or you can convert them to an alternative format. A track output produces uncompressed output. For audio output settings, this means that must be . For video output settings, this means that the dictionary must contain values for uncompressed video output, as defined in . A track output doesn’t support the audio setting key or the following video settings keys: , , and . When constructing video output settings, the choice of pixel format affects the performance and quality of the decompression. For optimal performance when decompressing video, the requested pixel format should be one that the decoder supports natively to avoid unnecessary conversions. Below are some recommendations: For H.264, use or when you know the video is full range. In iOS, use for JPEG output. In macOS, is the preferred pixel format for video and generally provides the best performance when decoding. If you need to work in the RGB domain, use in iOS, and in macOS. ProRes-encoded media can contain up to 12 bits per channel. For ProRes-encoded sources that you wish to preserve more than 8 bits per channel during decompression, use one of the following pixel formats: , , , or . doesn’t support scaling with any of these high-bit-depth pixel formats. If you use the above pixel formats, don’t specify or in the dictionary. Only ProRes encoders support these pixel formats. ProRes 4444-encoded media can contain a mathematically lossless alpha channel. To preserve the alpha channel during decompression, use a pixel format with an alpha component such as or . To test whether your source contains an alpha channel, check that the track’s format description has a key with a value of .
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput
type AssetReaderTrackOutput struct {
	AssetReaderOutput
}

// AssetReaderTrackOutputFrom constructs a [AssetReaderTrackOutput] from an unsafe.Pointer.
//
// An object that reads media data from a single track of an asset.
func AssetReaderTrackOutputFrom(ptr unsafe.Pointer) AssetReaderTrackOutput {
	return AssetReaderTrackOutput{
		AssetReaderOutput: AssetReaderOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetReaderTrackOutputClass) Alloc() AssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetReaderTrackOutputClass) New() AssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReaderTrackOutput) Init() AssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReaderTrackOutput) Autorelease() AssetReaderTrackOutput {
	rv := objc.Send[AssetReaderTrackOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReaderTrackOutput creates a new AssetReaderTrackOutput instance.
func NewAssetReaderTrackOutput() AssetReaderTrackOutput {
	return getAssetReaderTrackOutputClass().New()
}




