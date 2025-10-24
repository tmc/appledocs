// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetReaderSampleReferenceOutput] class.
var (
	AssetReaderSampleReferenceOutputClass     _AssetReaderSampleReferenceOutputClass
	AssetReaderSampleReferenceOutputClassOnce sync.Once
)

func getAssetReaderSampleReferenceOutputClass() _AssetReaderSampleReferenceOutputClass {
	AssetReaderSampleReferenceOutputClassOnce.Do(func() {
		AssetReaderSampleReferenceOutputClass = _AssetReaderSampleReferenceOutputClass{objc.GetClass("AVAssetReaderSampleReferenceOutput")}
	})
	return AssetReaderSampleReferenceOutputClass
}

type _AssetReaderSampleReferenceOutputClass struct {
	class objc.Class
}





// An interface definition for the [AssetReaderSampleReferenceOutput] class.
type IAssetReaderSampleReferenceOutput interface {
	IAssetReaderOutput
	

	// properties:
	Track() IAVAssetTrack
	AlwaysCopiesSampleData() bool
	SetAlwaysCopiesSampleData(value bool)
	KCMSampleBufferAttachmentKey_SampleReferenceByteOffset() foundation.String
	KCMSampleBufferAttachmentKey_SampleReferenceURL() foundation.String


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetReaderSampleReferenceOutputClass) Alloc() AssetReaderSampleReferenceOutput {
	rv := objc.Send[AssetReaderSampleReferenceOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetReaderSampleReferenceOutputClass) New() AssetReaderSampleReferenceOutput {
	rv := objc.Send[AssetReaderSampleReferenceOutput](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetReaderSampleReferenceOutput) Init() AssetReaderSampleReferenceOutput {
	rv := objc.Send[AssetReaderSampleReferenceOutput](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetReaderSampleReferenceOutput) Autorelease() AssetReaderSampleReferenceOutput {
	rv := objc.Send[AssetReaderSampleReferenceOutput](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetReaderSampleReferenceOutput creates a new AssetReaderSampleReferenceOutput instance.
func NewAssetReaderSampleReferenceOutput() AssetReaderSampleReferenceOutput {
	return getAssetReaderSampleReferenceOutputClass().New()
}





// An object that reads sample references from an asset track.
//
// Apps can extract information about the location of samples in a track — the file URL and offset — by adding an instance of this class to an asset reader. Read the and attachments on the extracted sample buffers to get the location of the sample data. You can also append sample buffers that you extract using this class to an instance to create movie tracks that aren’t self-contained and reference data in the original file instead. To write tracks that aren’t self-contained, use instances of that you configure to write files of type . Because this output doesn’t return sample data, it ignores the value of the property.


// An object that reads sample references from an asset track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput
type AssetReaderSampleReferenceOutput struct {
	AssetReaderOutput
}

// AssetReaderSampleReferenceOutputFrom constructs a [AssetReaderSampleReferenceOutput] from an unsafe.Pointer.
//
// An object that reads sample references from an asset track.
func AssetReaderSampleReferenceOutputFrom(ptr unsafe.Pointer) AssetReaderSampleReferenceOutput {
	return AssetReaderSampleReferenceOutput{
		AssetReaderOutput: AssetReaderOutputFrom(ptr),
	}
}






// Creates an object that supplies sample references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/init(track:)
func NewAssetReaderSampleReferenceOutputWithTrack(track IAVAssetTrack) AssetReaderSampleReferenceOutput {
	instance := getAssetReaderSampleReferenceOutputClass().Alloc()
	rv := objc.Send[AssetReaderSampleReferenceOutput](instance.ID, objc.Sel("initWithTrack:"), track)
	rv.Autorelease()
	return rv
}







// Returns a new object that supplies sample references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/assetReaderSampleReferenceOutputWithTrack:
func (ac _AssetReaderSampleReferenceOutputClass) AssetReaderSampleReferenceOutputWithTrack(track IAVAssetTrack) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetReaderSampleReferenceOutputWithTrack:"), track)
	return rv
}

















// The track from which the output reads sample references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/track
func (a_ AssetReaderSampleReferenceOutput) Track() IAVAssetTrack {
	rv := objc.Send[AssetTrack](a_.ID, objc.Sel("track"))
	return rv
}


// A Boolean value that indicates whether the output vends copied sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/alwayscopiessampledata
func (a_ AssetReaderSampleReferenceOutput) AlwaysCopiesSampleData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("alwaysCopiesSampleData"))
	return rv
}


// A Boolean value that indicates whether the output vends copied sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/alwayscopiessampledata
func (a_ AssetReaderSampleReferenceOutput) SetAlwaysCopiesSampleData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAlwaysCopiesSampleData:"), value)
}


// Indicates the byte offset at which the sample data begins (type `CFNumber`).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_SampleReferenceByteOffset
func (a_ AssetReaderSampleReferenceOutput) KCMSampleBufferAttachmentKey_SampleReferenceByteOffset() foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("kCMSampleBufferAttachmentKey_SampleReferenceByteOffset"))
	return rv
}


// Indicates the URL where the sample data is (type `CFURL`).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_SampleReferenceURL
func (a_ AssetReaderSampleReferenceOutput) KCMSampleBufferAttachmentKey_SampleReferenceURL() foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("kCMSampleBufferAttachmentKey_SampleReferenceURL"))
	return rv
}







