// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
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
	PrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool
	SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool)
	SpatialVideoCaptureEnabled() bool
	SetSpatialVideoCaptureEnabled(value bool)
	SpatialVideoCaptureSupported() bool
	Metadata() []MetadataItem
	SetMetadata(value []MetadataItem)
	MovieFragmentInterval() objectivec.IObject
	SetMovieFragmentInterval(value objectivec.IObject)
	PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
	PrimaryConstituentDeviceSwitchingBehaviorForRecording() CapturePrimaryConstituentDeviceSwitchingBehavior
	IsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool
	SetIsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool)
	IsSpatialVideoCaptureEnabled() bool
	SetIsSpatialVideoCaptureEnabled(value bool)
	IsSpatialVideoCaptureSupported() bool
	SetIsSpatialVideoCaptureSupported(value bool)


	

	// methods:
	OutputSettingsForConnection(connection IAVCaptureConnection) foundation.IDictionary
	SetOutputSettingsForConnection(outputSettings foundation.IDictionary, connection IAVCaptureConnection)
	SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions)


}





// Alloc allocates a new instance without initialization.
func (cc _CaptureMovieFileOutputClass) Alloc() CaptureMovieFileOutput {
	rv := objc.Send[CaptureMovieFileOutput](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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





















// Returns the settings the output uses to encode media from the specified connection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/outputSettings(for:)
func (c_ CaptureMovieFileOutput) OutputSettingsForConnection(connection IAVCaptureConnection) foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](c_.ID, objc.Sel("outputSettingsForConnection:"), connection)
	return rv
}


// Sets the options the output uses to encode media from the given connection while recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/setOutputSettings(_:for:)
func (c_ CaptureMovieFileOutput) SetOutputSettingsForConnection(outputSettings foundation.IDictionary, connection IAVCaptureConnection) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOutputSettings:forConnection:"), outputSettings, connection)
}


// Sets the camera switching behavior to use during recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/setPrimaryConstituentDeviceSwitchingBehaviorForRecording(_:restrictedSwitchingBehaviorConditions:)
func (c_ CaptureMovieFileOutput) SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions(switchingBehavior CapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryConstituentDeviceSwitchingBehaviorForRecording:restrictedSwitchingBehaviorConditions:"), switchingBehavior, restrictedSwitchingBehaviorConditions)
}







// A Boolean value that indicates whether to restrict constituent device switching behavior during recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/isPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled
func (c_ CaptureMovieFileOutput) PrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("primaryConstituentDeviceSwitchingBehaviorForRecordingEnabled"))
	return rv
}


// A Boolean value that indicates whether to restrict constituent device switching behavior during recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/isPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled
func (c_ CaptureMovieFileOutput) SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled:"), value)
}


// A Boolean value that indicates whether a movie file output captures spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/isSpatialVideoCaptureEnabled
func (c_ CaptureMovieFileOutput) SpatialVideoCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("spatialVideoCaptureEnabled"))
	return rv
}


// A Boolean value that indicates whether a movie file output captures spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/isSpatialVideoCaptureEnabled
func (c_ CaptureMovieFileOutput) SetSpatialVideoCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSpatialVideoCaptureEnabled:"), value)
}


// A Boolean value that indicates whether a movie file output supports capturing spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/isSpatialVideoCaptureSupported
func (c_ CaptureMovieFileOutput) SpatialVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("spatialVideoCaptureSupported"))
	return rv
}


// The metadata for the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/metadata
func (c_ CaptureMovieFileOutput) Metadata() []MetadataItem {
	rv := objc.Send[[]MetadataItem](c_.ID, objc.Sel("metadata"))
	return rv
}


// The metadata for the output file.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/metadata
func (c_ CaptureMovieFileOutput) SetMetadata(value []MetadataItem) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setMetadata:"), nsArray)
}


// The number of seconds of output that are written per fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/movieFragmentInterval
func (c_ CaptureMovieFileOutput) MovieFragmentInterval() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("movieFragmentInterval"))
	return rv
}


// The number of seconds of output that are written per fragment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/movieFragmentInterval
func (c_ CaptureMovieFileOutput) SetMovieFragmentInterval(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setMovieFragmentInterval:"), value)
}


// The conditions during which camera switching may occur while recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/primaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording
func (c_ CaptureMovieFileOutput) PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording() CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Send[CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c_.ID, objc.Sel("primaryConstituentDeviceRestrictedSwitchingBehaviorConditionsForRecording"))
	return rv
}


// The camera switching behavior to use for recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMovieFileOutput/primaryConstituentDeviceSwitchingBehaviorForRecording
func (c_ CaptureMovieFileOutput) PrimaryConstituentDeviceSwitchingBehaviorForRecording() CapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Send[CapturePrimaryConstituentDeviceSwitchingBehavior](c_.ID, objc.Sel("primaryConstituentDeviceSwitchingBehaviorForRecording"))
	return rv
}


// A Boolean value that indicates whether to restrict constituent device switching behavior during recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isprimaryconstituentdeviceswitchingbehaviorforrecordingenabled
func (c_ CaptureMovieFileOutput) IsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled"))
	return rv
}


// A Boolean value that indicates whether to restrict constituent device switching behavior during recording.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isprimaryconstituentdeviceswitchingbehaviorforrecordingenabled
func (c_ CaptureMovieFileOutput) SetIsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsPrimaryConstituentDeviceSwitchingBehaviorForRecordingEnabled:"), value)
}


// A Boolean value that indicates whether a movie file output captures spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isspatialvideocaptureenabled
func (c_ CaptureMovieFileOutput) IsSpatialVideoCaptureEnabled() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSpatialVideoCaptureEnabled"))
	return rv
}


// A Boolean value that indicates whether a movie file output captures spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isspatialvideocaptureenabled
func (c_ CaptureMovieFileOutput) SetIsSpatialVideoCaptureEnabled(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSpatialVideoCaptureEnabled:"), value)
}


// A Boolean value that indicates whether a movie file output supports capturing spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isspatialvideocapturesupported
func (c_ CaptureMovieFileOutput) IsSpatialVideoCaptureSupported() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isSpatialVideoCaptureSupported"))
	return rv
}


// A Boolean value that indicates whether a movie file output supports capturing spatial videos.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturemoviefileoutput/isspatialvideocapturesupported
func (c_ CaptureMovieFileOutput) SetIsSpatialVideoCaptureSupported(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsSpatialVideoCaptureSupported:"), value)
}







