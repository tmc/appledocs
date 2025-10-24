// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corevideo"
)

/* debug [class.gen.go]: Generating class CNCompositionInfo */


/* debug [class_header]: Header for CNCompositionInfo */
// The class instance for the [CNCompositionInfo] class.
var (
	CNCompositionInfoClass     _CNCompositionInfoClass
	CNCompositionInfoClassOnce sync.Once
)

func getCNCompositionInfoClass() _CNCompositionInfoClass {
	CNCompositionInfoClassOnce.Do(func() {
		CNCompositionInfoClass = _CNCompositionInfoClass{objc.GetClass("CNCompositionInfo")}
	})
	return CNCompositionInfoClass
}

type _CNCompositionInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNCompositionInfo */
// An interface definition for the [CNCompositionInfo] class.
type ICNCompositionInfo interface {
	ICNAssetInfo
	
/* debug [class_interface_properties]: Properties for CNCompositionInfo */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNCompositionInfo */
	// methods:
	InsertTimeRangeOfCinematicAssetInfoAtTimeError(timeRange TimeRange /* not a class type */, assetInfo ICNAssetInfo, startTime objc.IObject /* cross-framework: Time */, outError unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNCompositionInfo */
// Alloc allocates a new instance without initialization.
func (cc _CNCompositionInfoClass) Alloc() CNCompositionInfo {
	rv := objc.Send[CNCompositionInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNCompositionInfoClass) New() CNCompositionInfo {
	rv := objc.Send[CNCompositionInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNCompositionInfo) Init() CNCompositionInfo {
	rv := objc.Send[CNCompositionInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNCompositionInfo) Autorelease() CNCompositionInfo {
	rv := objc.Send[CNCompositionInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNCompositionInfo creates a new CNCompositionInfo instance.
func NewCNCompositionInfo() CNCompositionInfo {
	return getCNCompositionInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNCompositionInfo */
// An object that enables you to add the appropriate number of tracks for a Cinematic asset.


// An object that enables you to add the appropriate number of tracks for a Cinematic asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCompositionInfo-vzoh
type CNCompositionInfo struct {
	CNAssetInfo
}

// CNCompositionInfoFrom constructs a [CNCompositionInfo] from an unsafe.Pointer.
//
// An object that enables you to add the appropriate number of tracks for a Cinematic asset.
func CNCompositionInfoFrom(ptr unsafe.Pointer) CNCompositionInfo {
	return CNCompositionInfo{
		CNAssetInfo: CNAssetInfoFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNCompositionInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNCompositionInfo */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNCompositionInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNCompositionInfo */

// Inserts a time range of Cinematic source asset into the corresponding tracks of a composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCompositionInfo-vzoh/insertTimeRange:ofCinematicAssetInfo:atTime:error:
func (c_ CNCompositionInfo) InsertTimeRangeOfCinematicAssetInfoAtTimeError(timeRange TimeRange /* not a class type */, assetInfo ICNAssetInfo, startTime objc.IObject /* cross-framework: Time */, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("insertTimeRange:ofCinematicAssetInfo:atTime:error:"), timeRange, assetInfo, startTime, outError)
	return rv
}/* debug [instance_methods/method]: InsertTimeRangeOfCinematicAssetInfoAtTimeError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNCompositionInfo */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNCompositionInfo */



