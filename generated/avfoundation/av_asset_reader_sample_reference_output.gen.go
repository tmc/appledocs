// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetReaderSampleReferenceOutput */


/* debug [class_header]: Header for AVAssetReaderSampleReferenceOutput */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetReaderSampleReferenceOutput */
// An interface definition for the [AssetReaderSampleReferenceOutput] class.
type IAssetReaderSampleReferenceOutput interface {
	IAssetReaderOutput
	
/* debug [class_interface_properties]: Properties for AssetReaderSampleReferenceOutput */
	// properties:
	Track() IAVAssetTrack
	AlwaysCopiesSampleData() bool
	SetAlwaysCopiesSampleData(value bool)
	KCMSampleBufferAttachmentKey_SampleReferenceByteOffset() foundation.String
	KCMSampleBufferAttachmentKey_SampleReferenceURL() foundation.String
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetReaderSampleReferenceOutput */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetReaderSampleReferenceOutput */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetReaderSampleReferenceOutput */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetReaderSampleReferenceOutput */

// Creates an object that supplies sample references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/init(track:)
func NewAssetReaderSampleReferenceOutputWithTrack(track IAVAssetTrack) AssetReaderSampleReferenceOutput {
	instance := getAssetReaderSampleReferenceOutputClass().Alloc()
	rv := objc.Send[AssetReaderSampleReferenceOutput](instance.ID, objc.Sel("initWithTrack:"), track)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAssetReaderSampleReferenceOutputWithTrack */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetReaderSampleReferenceOutput */

// Returns a new object that supplies sample references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/assetReaderSampleReferenceOutputWithTrack:
func (ac _AssetReaderSampleReferenceOutputClass) AssetReaderSampleReferenceOutputWithTrack(track IAVAssetTrack) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("assetReaderSampleReferenceOutputWithTrack:"), track)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AssetReaderSampleReferenceOutputWithTrack) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetReaderSampleReferenceOutput */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetReaderSampleReferenceOutput */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetReaderSampleReferenceOutput */

// The track from which the output reads sample references.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderSampleReferenceOutput/track
func (a_ AssetReaderSampleReferenceOutput) Track() IAVAssetTrack {
	rv := objc.Send[AssetTrack](a_.ID, objc.Sel("track"))
	return rv
}/* debug [instance_properties/getter]: track */


// A Boolean value that indicates whether the output vends copied sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/alwayscopiessampledata
func (a_ AssetReaderSampleReferenceOutput) AlwaysCopiesSampleData() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("alwaysCopiesSampleData"))
	return rv
}/* debug [instance_properties/getter]: alwaysCopiesSampleData */


// A Boolean value that indicates whether the output vends copied sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetreaderoutput/alwayscopiessampledata
func (a_ AssetReaderSampleReferenceOutput) SetAlwaysCopiesSampleData(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAlwaysCopiesSampleData:"), value)
}/* debug [instance_properties/setter]: alwaysCopiesSampleData */


// Indicates the byte offset at which the sample data begins (type `CFNumber`).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_SampleReferenceByteOffset
func (a_ AssetReaderSampleReferenceOutput) KCMSampleBufferAttachmentKey_SampleReferenceByteOffset() foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("kCMSampleBufferAttachmentKey_SampleReferenceByteOffset"))
	return rv
}/* debug [instance_properties/getter]: kCMSampleBufferAttachmentKey_SampleReferenceByteOffset */


// Indicates the URL where the sample data is (type `CFURL`).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_SampleReferenceURL
func (a_ AssetReaderSampleReferenceOutput) KCMSampleBufferAttachmentKey_SampleReferenceURL() foundation.String {
	rv := objc.Send[foundation.String](a_.ID, objc.Sel("kCMSampleBufferAttachmentKey_SampleReferenceURL"))
	return rv
}/* debug [instance_properties/getter]: kCMSampleBufferAttachmentKey_SampleReferenceURL */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetReaderSampleReferenceOutput */


