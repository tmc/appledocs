// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CNCompositionInfo] class.
type ICNCompositionInfo interface {
	ICNAssetInfo
	InsertTimeRangeOfCinematicAssetInfoAtTimeError(timeRange unsafe.Pointer, assetInfo ICNAssetInfo, startTime unsafe.Pointer, outError unsafe.Pointer) bool
}

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

// Alloc allocates a new instance without initialization.
func (cc _CNCompositionInfoClass) Alloc() CNCompositionInfo {
	rv := objc.Send[CNCompositionInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Inserts a time range of Cinematic source asset into the corresponding tracks of a composition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNCompositionInfo-vzoh/insertTimeRange:ofCinematicAssetInfo:atTime:error:

func (c_ CNCompositionInfo) InsertTimeRangeOfCinematicAssetInfoAtTimeError(timeRange unsafe.Pointer, assetInfo ICNAssetInfo, startTime unsafe.Pointer, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("insertTimeRange:ofCinematicAssetInfo:atTime:error:"), timeRange, assetInfo, startTime, outError)
	return rv
}



