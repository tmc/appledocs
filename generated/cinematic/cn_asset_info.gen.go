// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/avfoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNAssetInfo */


/* debug [class_header]: Header for CNAssetInfo */
// The class instance for the [CNAssetInfo] class.
var (
	CNAssetInfoClass     _CNAssetInfoClass
	CNAssetInfoClassOnce sync.Once
)

func getCNAssetInfoClass() _CNAssetInfoClass {
	CNAssetInfoClassOnce.Do(func() {
		CNAssetInfoClass = _CNAssetInfoClass{objc.GetClass("CNAssetInfo")}
	})
	return CNAssetInfoClass
}

type _CNAssetInfoClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNAssetInfo */
// An interface definition for the [CNAssetInfo] class.
type ICNAssetInfo interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNAssetInfo */
	// properties:
	AllCinematicTracks() []avfoundation.AssetTrack
	Asset() avfoundation.Asset
	CinematicDisparityTrack() avfoundation.AssetTrack
	CinematicMetadataTrack() avfoundation.AssetTrack
	CinematicVideoTrack() avfoundation.AssetTrack
	FrameTimingTrack() avfoundation.AssetTrack
	NaturalSize() corefoundation.CGSize
	PreferredSize() corefoundation.CGSize
	PreferredTransform() corefoundation.CGAffineTransform
	SampleDataTrackIDs() []foundation.Number
	TimeRange() TimeRange /* not a class type */
	VideoCompositionTrackIDs() []foundation.Number
	VideoCompositionTracks() []avfoundation.AssetTrack
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNAssetInfo */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNAssetInfo */
// Alloc allocates a new instance without initialization.
func (cc _CNAssetInfoClass) Alloc() CNAssetInfo {
	rv := objc.Send[CNAssetInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNAssetInfoClass) New() CNAssetInfo {
	rv := objc.Send[CNAssetInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNAssetInfo) Init() CNAssetInfo {
	rv := objc.Send[CNAssetInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNAssetInfo) Autorelease() CNAssetInfo {
	rv := objc.Send[CNAssetInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNAssetInfo creates a new CNAssetInfo instance.
func NewCNAssetInfo() CNAssetInfo {
	return getCNAssetInfoClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNAssetInfo */
// An object that provides Cinematic-specific information about an asset, including its tracks.


// An object that provides Cinematic-specific information about an asset, including its tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t
type CNAssetInfo struct {
	objectivec.Object
}

// CNAssetInfoFrom constructs a [CNAssetInfo] from an unsafe.Pointer.
//
// An object that provides Cinematic-specific information about an asset, including its tracks.
func CNAssetInfoFrom(ptr unsafe.Pointer) CNAssetInfo {
	return CNAssetInfo{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNAssetInfo *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNAssetInfo */

// Determines if the asset is Cinematic.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/checkIfCinematic:completionHandler:
func (cc _CNAssetInfoClass) CheckIfCinematicCompletionHandler(asset avfoundation.Asset, completionHandler bool) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("checkIfCinematic:completionHandler:"), asset, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CheckIfCinematicCompletionHandler) */


// Loads the Cinematic asset information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/loadFromAsset:completionHandler:
func (cc _CNAssetInfoClass) LoadFromAssetCompletionHandler(asset avfoundation.Asset, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadFromAsset:completionHandler:"), asset, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadFromAssetCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNAssetInfo */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNAssetInfo */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNAssetInfo */

// An array of the Cinematic asset tracks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/allCinematicTracks
func (c_ CNAssetInfo) AllCinematicTracks() []avfoundation.AssetTrack {
	rv := objc.Send[[]avfoundation.AssetTrack](c_.ID, objc.Sel("allCinematicTracks"))
	return rv
}/* debug [instance_properties/getter]: allCinematicTracks */


// The original Cinematic source asset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/asset
func (c_ CNAssetInfo) Asset() avfoundation.Asset {
	rv := objc.Send[avfoundation.Asset](c_.ID, objc.Sel("asset"))
	return rv
}/* debug [instance_properties/getter]: asset */


// The Cinematic disparity track.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/cinematicDisparityTrack
func (c_ CNAssetInfo) CinematicDisparityTrack() avfoundation.AssetTrack {
	rv := objc.Send[avfoundation.AssetTrack](c_.ID, objc.Sel("cinematicDisparityTrack"))
	return rv
}/* debug [instance_properties/getter]: cinematicDisparityTrack */


// The Cinematic metadata track used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/cinematicMetadataTrack
func (c_ CNAssetInfo) CinematicMetadataTrack() avfoundation.AssetTrack {
	rv := objc.Send[avfoundation.AssetTrack](c_.ID, objc.Sel("cinematicMetadataTrack"))
	return rv
}/* debug [instance_properties/getter]: cinematicMetadataTrack */


// Track used for Cinematic video.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/cinematicVideoTrack
func (c_ CNAssetInfo) CinematicVideoTrack() avfoundation.AssetTrack {
	rv := objc.Send[avfoundation.AssetTrack](c_.ID, objc.Sel("cinematicVideoTrack"))
	return rv
}/* debug [instance_properties/getter]: cinematicVideoTrack */


// The track used for Cinematic frame timing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/frameTimingTrack
func (c_ CNAssetInfo) FrameTimingTrack() avfoundation.AssetTrack {
	rv := objc.Send[avfoundation.AssetTrack](c_.ID, objc.Sel("frameTimingTrack"))
	return rv
}/* debug [instance_properties/getter]: frameTimingTrack */


// The video size if rendered at its natural size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/naturalSize
func (c_ CNAssetInfo) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("naturalSize"))
	return rv
}/* debug [instance_properties/getter]: naturalSize */


// The video size if rendered at its natural size with the preferred transform applied.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/preferredSize
func (c_ CNAssetInfo) PreferredSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c_.ID, objc.Sel("preferredSize"))
	return rv
}/* debug [instance_properties/getter]: preferredSize */


// The preferred transform of the rendered image for display purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/preferredTransform
func (c_ CNAssetInfo) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](c_.ID, objc.Sel("preferredTransform"))
	return rv
}/* debug [instance_properties/getter]: preferredTransform */


// The source metadata track IDs required to implement the video composition instruction protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/sampleDataTrackIDs
func (c_ CNAssetInfo) SampleDataTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("sampleDataTrackIDs"))
	return rv
}/* debug [instance_properties/getter]: sampleDataTrackIDs */


// The time range over which all Cinematic tracks are valid.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/timeRange
func (c_ CNAssetInfo) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */


// Source video track IDs required to implement the video composition instruction protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/videoCompositionTrackIDs
func (c_ CNAssetInfo) VideoCompositionTrackIDs() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("videoCompositionTrackIDs"))
	return rv
}/* debug [instance_properties/getter]: videoCompositionTrackIDs */


// Tracks required to construct the video composition output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetInfo-8ja4t/videoCompositionTracks
func (c_ CNAssetInfo) VideoCompositionTracks() []avfoundation.AssetTrack {
	rv := objc.Send[[]avfoundation.AssetTrack](c_.ID, objc.Sel("videoCompositionTracks"))
	return rv
}/* debug [instance_properties/getter]: videoCompositionTracks */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNAssetInfo */



