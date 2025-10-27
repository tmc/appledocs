// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetSegmentReportSampleInformation] class.
var (
	AssetSegmentReportSampleInformationClass     _AssetSegmentReportSampleInformationClass
	AssetSegmentReportSampleInformationClassOnce sync.Once
)

func getAssetSegmentReportSampleInformationClass() _AssetSegmentReportSampleInformationClass {
	AssetSegmentReportSampleInformationClassOnce.Do(func() {
		AssetSegmentReportSampleInformationClass = _AssetSegmentReportSampleInformationClass{objc.GetClass("AVAssetSegmentReportSampleInformation")}
	})
	return AssetSegmentReportSampleInformationClass
}

type _AssetSegmentReportSampleInformationClass struct {
	class objc.Class
}





// An interface definition for the [AssetSegmentReportSampleInformation] class.
type IAssetSegmentReportSampleInformation interface {
	objectivec.IObject
	

	// properties:
	IsSyncSample() bool
	Length() int
	Offset() int
	PresentationTimeStamp() objectivec.IObject
	Duration() objectivec.IObject
	SetDuration(value objectivec.IObject)
	EarliestPresentationTimeStamp() objectivec.IObject
	SetEarliestPresentationTimeStamp(value objectivec.IObject)
	FirstVideoSampleInformation() IAVAssetSegmentReportSampleInformation
	SetFirstVideoSampleInformation(value IAVAssetSegmentReportSampleInformation)
	MediaType() MediaType
	SetMediaType(value MediaType)
	TrackID() PersistentTrackID /* not a class type */
	SetTrackID(value PersistentTrackID /* not a class type */)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetSegmentReportSampleInformationClass) Alloc() AssetSegmentReportSampleInformation {
	rv := objc.Send[AssetSegmentReportSampleInformation](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetSegmentReportSampleInformationClass) New() AssetSegmentReportSampleInformation {
	rv := objc.Send[AssetSegmentReportSampleInformation](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetSegmentReportSampleInformation) Init() AssetSegmentReportSampleInformation {
	rv := objc.Send[AssetSegmentReportSampleInformation](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetSegmentReportSampleInformation) Autorelease() AssetSegmentReportSampleInformation {
	rv := objc.Send[AssetSegmentReportSampleInformation](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetSegmentReportSampleInformation creates a new AssetSegmentReportSampleInformation instance.
func NewAssetSegmentReportSampleInformation() AssetSegmentReportSampleInformation {
	return getAssetSegmentReportSampleInformationClass().New()
}





// An object that provides information about sample data in a track.


// An object that provides information about sample data in a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation
type AssetSegmentReportSampleInformation struct {
	objectivec.Object
}

// AssetSegmentReportSampleInformationFrom constructs a [AssetSegmentReportSampleInformation] from an unsafe.Pointer.
//
// An object that provides information about sample data in a track.
func AssetSegmentReportSampleInformationFrom(ptr unsafe.Pointer) AssetSegmentReportSampleInformation {
	return AssetSegmentReportSampleInformation{objectivec.Object{objc.ID(ptr)}}
}

























// A Boolean value that indicates whether the sample is a key frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation/isSyncSample
func (a_ AssetSegmentReportSampleInformation) IsSyncSample() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isSyncSample"))
	return rv
}


// The length of the sample data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation/length
func (a_ AssetSegmentReportSampleInformation) Length() int {
	rv := objc.Send[int](a_.ID, objc.Sel("length"))
	return rv
}


// The offset of a sample in the segment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation/offset
func (a_ AssetSegmentReportSampleInformation) Offset() int {
	rv := objc.Send[int](a_.ID, objc.Sel("offset"))
	return rv
}


// The presentation timestamp (PTS) of a sample.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReportSampleInformation/presentationTimeStamp
func (a_ AssetSegmentReportSampleInformation) PresentationTimeStamp() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("presentationTimeStamp"))
	return rv
}


// The duration of a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/duration
func (a_ AssetSegmentReportSampleInformation) Duration() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("duration"))
	return rv
}


// The duration of a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/duration
func (a_ AssetSegmentReportSampleInformation) SetDuration(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDuration:"), value)
}


// The earliest presentation timestamp (PTS) for this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/earliestpresentationtimestamp
func (a_ AssetSegmentReportSampleInformation) EarliestPresentationTimeStamp() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("earliestPresentationTimeStamp"))
	return rv
}


// The earliest presentation timestamp (PTS) for this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/earliestpresentationtimestamp
func (a_ AssetSegmentReportSampleInformation) SetEarliestPresentationTimeStamp(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setEarliestPresentationTimeStamp:"), value)
}


// Information about the first video sample in a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/firstvideosampleinformation
func (a_ AssetSegmentReportSampleInformation) FirstVideoSampleInformation() IAVAssetSegmentReportSampleInformation {
	rv := objc.Send[AssetSegmentReportSampleInformation](a_.ID, objc.Sel("firstVideoSampleInformation"))
	return rv
}


// Information about the first video sample in a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/firstvideosampleinformation
func (a_ AssetSegmentReportSampleInformation) SetFirstVideoSampleInformation(value IAVAssetSegmentReportSampleInformation) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFirstVideoSampleInformation:"), value)
}


// The type of media a track contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/mediatype
func (a_ AssetSegmentReportSampleInformation) MediaType() MediaType {
	rv := objc.Send[MediaType](a_.ID, objc.Sel("mediaType"))
	return rv
}


// The type of media a track contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/mediatype
func (a_ AssetSegmentReportSampleInformation) SetMediaType(value MediaType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setMediaType:"), value)
}


// A persistent unique identifier for a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/trackid
func (a_ AssetSegmentReportSampleInformation) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](a_.ID, objc.Sel("trackID"))
	return rv
}


// A persistent unique identifier for a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmenttrackreport/trackid
func (a_ AssetSegmentReportSampleInformation) SetTrackID(value PersistentTrackID /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTrackID:"), value)
}








