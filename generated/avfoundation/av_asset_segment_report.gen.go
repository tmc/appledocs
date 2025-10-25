// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVAssetSegmentReport */


/* debug [class_header]: Header for AVAssetSegmentReport */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AssetSegmentReport */
// An interface definition for the [AssetSegmentReport] class.
type IAssetSegmentReport interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for AssetSegmentReport */
	// properties:
	SegmentType() AssetSegmentType
	TrackReports() []AssetSegmentTrackReport
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AssetSegmentReport */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AssetSegmentReport */
// Alloc allocates a new instance without initialization.
func (ac _AssetSegmentReportClass) Alloc() AssetSegmentReport {
	rv := objc.Send[AssetSegmentReport](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AssetSegmentReport */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AssetSegmentReport *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AssetSegmentReport */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AssetSegmentReport */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AssetSegmentReport */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AssetSegmentReport */

// The type of segment data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReport/segmentType
func (a_ AssetSegmentReport) SegmentType() AssetSegmentType {
	rv := objc.Send[AssetSegmentType](a_.ID, objc.Sel("segmentType"))
	return rv
}/* debug [instance_properties/getter]: segmentType */


// The reports for the segment’s track data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentReport/trackReports
func (a_ AssetSegmentReport) TrackReports() []AssetSegmentTrackReport {
	rv := objc.Send[[]AssetSegmentTrackReport](a_.ID, objc.Sel("trackReports"))
	return rv
}/* debug [instance_properties/getter]: trackReports */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVAssetSegmentReport */



