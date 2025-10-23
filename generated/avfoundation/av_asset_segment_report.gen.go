// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AssetSegmentReport] class.
var (
	AssetSegmentReportClass     _AssetSegmentReportClass
	AssetSegmentReportClassOnce sync.Once
)

func getAssetSegmentReportClass() _AssetSegmentReportClass {
	AssetSegmentReportClassOnce.Do(func() {
		AssetSegmentReportClass = _AssetSegmentReportClass{objc.GetClass("AVAssetSegmentReport")}
	})
	return AssetSegmentReportClass
}

type _AssetSegmentReportClass struct {
	class objc.Class
}

// An interface definition for the [AssetSegmentReport] class.
type IAssetSegmentReport interface {
	objectivec.IObject
	SegmentType() unsafe.Pointer
	SetSegmentType(value unsafe.Pointer)
	TrackReports() unsafe.Pointer
	SetTrackReports(value unsafe.Pointer)
}

// An object that provides information about segment data.
//
// You receive a segment report through the delegate method.


// An object that provides information about segment data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReport
type AssetSegmentReport struct {
	objectivec.Object
}

// AssetSegmentReportFrom constructs a [AssetSegmentReport] from an unsafe.Pointer.
//
// An object that provides information about segment data.
func AssetSegmentReportFrom(ptr unsafe.Pointer) AssetSegmentReport {
	return AssetSegmentReport{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AssetSegmentReportClass) Alloc() AssetSegmentReport {
	rv := objc.Send[AssetSegmentReport](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AssetSegmentReportClass) New() AssetSegmentReport {
	rv := objc.Send[AssetSegmentReport](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AssetSegmentReport) Init() AssetSegmentReport {
	rv := objc.Send[AssetSegmentReport](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AssetSegmentReport) Autorelease() AssetSegmentReport {
	rv := objc.Send[AssetSegmentReport](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAssetSegmentReport creates a new AssetSegmentReport instance.
func NewAssetSegmentReport() AssetSegmentReport {
	return getAssetSegmentReportClass().New()
}



// The type of segment data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/segmenttype
func (a_ AssetSegmentReport) SegmentType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("segmentType"))
	return rv
}


// The type of segment data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/segmenttype
func (a_ AssetSegmentReport) SetSegmentType(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSegmentType:"), value)
}


// The reports for the segment’s track data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/trackreports
func (a_ AssetSegmentReport) TrackReports() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("trackReports"))
	return rv
}


// The reports for the segment’s track data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avassetsegmentreport/trackreports
func (a_ AssetSegmentReport) SetTrackReports(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTrackReports:"), value)
}



