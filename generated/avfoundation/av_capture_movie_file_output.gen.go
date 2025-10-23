// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CaptureMovieFileOutput] class.
var (
	CaptureMovieFileOutputClass     _CaptureMovieFileOutputClass
	CaptureMovieFileOutputClassOnce sync.Once
)

func getCaptureMovieFileOutputClass() _CaptureMovieFileOutputClass {
	CaptureMovieFileOutputClassOnce.Do(func() {
		CaptureMovieFileOutputClass = _CaptureMovieFileOutputClass{objc.GetClass("AVCaptureMovieFileOutput")}
	})
	return CaptureMovieFileOutputClass
}

type _CaptureMovieFileOutputClass struct {
	class objc.Class
}

// An interface definition for the [CaptureMovieFileOutput] class.
type ICaptureMovieFileOutput interface {
	ICaptureFileOutput
	// properties:
	AvailableVideoCodecTypes() VideoCodecType /* not a class type */
	SetAvailableVideoCodecTypes(value VideoCodecType /* not a class type */)
	IsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool /* primitive/slice/pointer */
	SetIsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool /* primitive/slice/pointer */)
	IsSpatialVideoCaptureEnabled() bool /* primitive/slice/pointer */
	SetIsSpatialVideoCaptureEnabled(value bool /* primitive/slice/pointer */)
	IsSpatialVideoCaptureSupported() bool /* primitive/slice/pointer */
	SetIsSpatialVideoCaptureSupported(value bool /* primitive/slice/pointer */)
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	MovieFragmentInterval() Time /* not a class type */
	SetMovieFragmentInterval(value Time /* not a class type */)
	PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording() unsafe.Pointer
	SetPrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording(value unsafe.Pointer)
	PrimaryConstituentDeviceSwitchingBehaviorForRecording() unsafe.Pointer
	SetPrimaryConstituentDeviceSwitchingBehaviorForRecording(value unsafe.Pointer)
	// methods:
}

// A capture output that records video and audio to a QuickTime movie file.
//
// A movie file output provides a complete file recording interface for writing media data to QuickTime movie files. It includes the ability to configure QuickTime-specific options, including writing metadata collections to each file, specify media encoding options for each track, and specify the interval at which it writes movie fragments.


// A capture output that records video and audio to a QuickTime movie file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput
type CaptureMovieFileOutput struct {
	CaptureFileOutput
}

// CaptureMovieFileOutputFrom constructs a [CaptureMovieFileOutput] from an unsafe.Pointer.
//
// A capture output that records video and audio to a QuickTime movie file.
func CaptureMovieFileOutputFrom(ptr unsafe.Pointer) CaptureMovieFileOutput {
	return CaptureMovieFileOutput{
		CaptureFileOutput: CaptureFileOutputFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CaptureMovieFileOutputClass) Alloc() CaptureMovieFileOutput {
	rv := objc.Send[CaptureMovieFileOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CaptureMovieFileOutputClass) New() CaptureMovieFileOutput {
	rv := objc.Send[CaptureMovieFileOutput](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptureMovieFileOutput) Init() CaptureMovieFileOutput {
	rv := objc.Send[CaptureMovieFileOutput](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptureMovieFileOutput) Autorelease() CaptureMovieFileOutput {
	rv := objc.Send[CaptureMovieFileOutput](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptureMovieFileOutput creates a new CaptureMovieFileOutput instance.
func NewCaptureMovieFileOutput() CaptureMovieFileOutput {
	return getCaptureMovieFileOutputClass().New()
}



// The video codecs types the output supports for recording movie files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/availablevideocodectypes
func (c_ CaptureMovieFileOutput) AvailableVideoCodecTypes() VideoCodecType /* not a class type */ {
	rv := objc.Send[VideoCodecType](c_.ID, objc.Sel("availableVideoCodecTypes"))
	return rv
}


// The video codecs types the output supports for recording movie files.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/availablevideocodectypes
func (c_ CaptureMovieFileOutput) SetAvailableVideoCodecTypes(value VideoCodecType /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setAvailableVideoCodecTypes:"), value)
}


// A Boolean value that indicates whether to restrict constituent device switching behavior during recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isprimaryconstituentdeviceswitchingbehaviorforrecordingenabled
func (c_ CaptureMovieFileOutput) IsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled"))
	return rv
}


// A Boolean value that indicates whether to restrict constituent device switching behavior during recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isprimaryconstituentdeviceswitchingbehaviorforrecordingenabled
func (c_ CaptureMovieFileOutput) SetIsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled:"), value)
}


// A Boolean value that indicates whether a movie file output captures spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isspatialvideocaptureenabled
func (c_ CaptureMovieFileOutput) IsSpatialVideoCaptureEnabled() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSpatialVideoCaptureEnabled"))
	return rv
}


// A Boolean value that indicates whether a movie file output captures spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isspatialvideocaptureenabled
func (c_ CaptureMovieFileOutput) SetIsSpatialVideoCaptureEnabled(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSpatialVideoCaptureEnabled:"), value)
}


// A Boolean value that indicates whether a movie file output supports capturing spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isspatialvideocapturesupported
func (c_ CaptureMovieFileOutput) IsSpatialVideoCaptureSupported() bool /* primitive/slice/pointer */ {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSpatialVideoCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether a movie file output supports capturing spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isspatialvideocapturesupported
func (c_ CaptureMovieFileOutput) SetIsSpatialVideoCaptureSupported(value bool /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSpatialVideoCaptureSupported:"), value)
}


// The metadata for the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/metadata
func (c_ CaptureMovieFileOutput) Metadata() IAVMetadataItem {
	rv := objc.Send[MetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata for the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/metadata
func (c_ CaptureMovieFileOutput) SetMetadata(value IAVMetadataItem) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), value)
}


// The number of seconds of output that are written per fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/moviefragmentinterval
func (c_ CaptureMovieFileOutput) MovieFragmentInterval() Time /* not a class type */ {
	rv := objc.Send[Time](c_.ID, objc.Sel("movieFragmentInterval"))
	return rv
}


// The number of seconds of output that are written per fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/moviefragmentinterval
func (c_ CaptureMovieFileOutput) SetMovieFragmentInterval(value Time /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMovieFragmentInterval:"), value)
}


// The conditions during which camera switching may occur while recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/primaryconstituentdevicerestrictedswitchingbehaviorconditionsforrecording
func (c_ CaptureMovieFileOutput) PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("primaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording"))
	return rv
}


// The conditions during which camera switching may occur while recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/primaryconstituentdevicerestrictedswitchingbehaviorconditionsforrecording
func (c_ CaptureMovieFileOutput) SetPrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording:"), value)
}


// The camera switching behavior to use for recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/primaryconstituentdeviceswitchingbehaviorforrecording
func (c_ CaptureMovieFileOutput) PrimaryConstituentDeviceSwitchingBehaviorForRecording() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("primaryConstituentDeviceSwitchingBehaviorForRecording"))
	return rv
}


// The camera switching behavior to use for recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/primaryconstituentdeviceswitchingbehaviorforrecording
func (c_ CaptureMovieFileOutput) SetPrimaryConstituentDeviceSwitchingBehaviorForRecording(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryConstituentDeviceSwitchingBehaviorForRecording:"), value)
}



