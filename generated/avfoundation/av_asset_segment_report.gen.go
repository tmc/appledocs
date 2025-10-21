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
}

// An object that provides information about segment data.
//
// You receive a segment report through the delegate method.
//
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




