// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [AssetSegmentTrackReport] class.
var (
	AssetSegmentTrackReportClass     _AssetSegmentTrackReportClass
	AssetSegmentTrackReportClassOnce sync.Once
)

func getAssetSegmentTrackReportClass() _AssetSegmentTrackReportClass {
	AssetSegmentTrackReportClassOnce.Do(func() {
		AssetSegmentTrackReportClass = _AssetSegmentTrackReportClass{objc.GetClass("AVAssetSegmentTrackReport")}
	})
	return AssetSegmentTrackReportClass
}

type _AssetSegmentTrackReportClass struct {
	class objc.Class
}





// An interface definition for the [AssetSegmentTrackReport] class.
type IAssetSegmentTrackReport interface {
	objectivec.IObject
	

	// properties:
	Duration() objc.IObject /* cross-framework: Time */
	EarliestPresentationTimeStamp() objc.IObject /* cross-framework: Time */
	FirstVideoSampleInformation() IAVAssetSegmentReportSampleInformation
	MediaType() MediaType /* typedef */
	TrackID() PersistentTrackID /* not a class type */
	SegmentType() AssetSegmentType
	SetSegmentType(value AssetSegmentType)
	TrackReports() IAVAssetSegmentTrackReport
	SetTrackReports(value IAVAssetSegmentTrackReport)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (ac _AssetSegmentTrackReportClass) Alloc() AssetSegmentTrackReport {
	rv := objc.Send[AssetSegmentTrackReport](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AssetSegmentTrackReportClass) New() AssetSegmentTrackReport {
	rv := objc.Send[AssetSegmentTrackReport](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetSegmentTrackReport) Init() AssetSegmentTrackReport {
	rv := objc.Send[AssetSegmentTrackReport](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetSegmentTrackReport) Autorelease() AssetSegmentTrackReport {
	rv := objc.Send[AssetSegmentTrackReport](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetSegmentTrackReport creates a new AssetSegmentTrackReport instance.
func NewAssetSegmentTrackReport() AssetSegmentTrackReport {
	return getAssetSegmentTrackReportClass().New()
}





// An object that provides information on a track in segment data.


// An object that provides information on a track in segment data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport
type AssetSegmentTrackReport struct {
	objectivec.Object
}

// AssetSegmentTrackReportFrom constructs a [AssetSegmentTrackReport] from an unsafe.Pointer.
//
// An object that provides information on a track in segment data.
func AssetSegmentTrackReportFrom(ptr unsafe.Pointer) AssetSegmentTrackReport {
	return AssetSegmentTrackReport{objectivec.Object{objc.ID(ptr)}}
}

























// The duration of a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/duration
func (a_ AssetSegmentTrackReport) Duration() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("duration"))
	return rv
}


// The earliest presentation timestamp (PTS) for this track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/earliestPresentationTimeStamp
func (a_ AssetSegmentTrackReport) EarliestPresentationTimeStamp() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](a_.ID, objc.Sel("earliestPresentationTimeStamp"))
	return rv
}


// Information about the first video sample in a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/firstVideoSampleInformation
func (a_ AssetSegmentTrackReport) FirstVideoSampleInformation() IAVAssetSegmentReportSampleInformation {
	rv := objc.Send[AssetSegmentReportSampleInformation](a_.ID, objc.Sel("firstVideoSampleInformation"))
	return rv
}


// The type of media a track contains.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/mediaType
func (a_ AssetSegmentTrackReport) MediaType() MediaType /* typedef */ {
	rv := objc.Send[foundation.NSString](a_.ID, objc.Sel("mediaType"))
	return rv
}


// A persistent unique identifier for a track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentTrackReport/trackID
func (a_ AssetSegmentTrackReport) TrackID() PersistentTrackID /* not a class type */ {
	rv := objc.Send[PersistentTrackID](a_.ID, objc.Sel("trackID"))
	return rv
}


// The type of segment data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/segmenttype
func (a_ AssetSegmentTrackReport) SegmentType() AssetSegmentType {
	rv := objc.Send[AssetSegmentType](a_.ID, objc.Sel("segmentType"))
	return rv
}


// The type of segment data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/segmenttype
func (a_ AssetSegmentTrackReport) SetSegmentType(value AssetSegmentType) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentType:"), value)
}


// The reports for the segment’s track data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/trackreports
func (a_ AssetSegmentTrackReport) TrackReports() IAVAssetSegmentTrackReport {
	rv := objc.Send[AssetSegmentTrackReport](a_.ID, objc.Sel("trackReports"))
	return rv
}


// The reports for the segment’s track data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/trackreports
func (a_ AssetSegmentTrackReport) SetTrackReports(value IAVAssetSegmentTrackReport) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTrackReports:"), value)
}








