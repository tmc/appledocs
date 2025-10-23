// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNAssetSpatialAudioInfo] class.
var (
	CNAssetSpatialAudioInfoClass     _CNAssetSpatialAudioInfoClass
	CNAssetSpatialAudioInfoClassOnce sync.Once
)

func getCNAssetSpatialAudioInfoClass() _CNAssetSpatialAudioInfoClass {
	CNAssetSpatialAudioInfoClassOnce.Do(func() {
		CNAssetSpatialAudioInfoClass = _CNAssetSpatialAudioInfoClass{objc.GetClass("CNAssetSpatialAudioInfo")}
	})
	return CNAssetSpatialAudioInfoClass
}

type _CNAssetSpatialAudioInfoClass struct {
	class objc.Class
}

// An interface definition for the [CNAssetSpatialAudioInfo] class.
type ICNAssetSpatialAudioInfo interface {
	objectivec.IObject
	// properties:
	SpatialAudioMixMetadata() foundation.objc.IObject /* cross-framework: NSData */
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5
type CNAssetSpatialAudioInfo struct {
	objectivec.Object
}

// CNAssetSpatialAudioInfoFrom constructs a [CNAssetSpatialAudioInfo] from an unsafe.Pointer.
func CNAssetSpatialAudioInfoFrom(ptr unsafe.Pointer) CNAssetSpatialAudioInfo {
	return CNAssetSpatialAudioInfo{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNAssetSpatialAudioInfoClass) Alloc() CNAssetSpatialAudioInfo {
	rv := objc.Send[CNAssetSpatialAudioInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNAssetSpatialAudioInfoClass) New() CNAssetSpatialAudioInfo {
	rv := objc.Send[CNAssetSpatialAudioInfo](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNAssetSpatialAudioInfo) Init() CNAssetSpatialAudioInfo {
	rv := objc.Send[CNAssetSpatialAudioInfo](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNAssetSpatialAudioInfo) Autorelease() CNAssetSpatialAudioInfo {
	rv := objc.Send[CNAssetSpatialAudioInfo](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNAssetSpatialAudioInfo creates a new CNAssetSpatialAudioInfo instance.
func NewCNAssetSpatialAudioInfo() CNAssetSpatialAudioInfo {
	return getCNAssetSpatialAudioInfoClass().New()
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/loadFromAsset:completionHandler:
func (cc _CNAssetSpatialAudioInfoClass) LoadFromAssetCompletionHandler(asset objc.IObject /* cross-framework Asset */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(cc.class), objc.Sel("loadFromAsset:completionHandler:"), asset, completionHandler)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNAssetSpatialAudioInfo-57yc5/spatialAudioMixMetadata
func (c_ CNAssetSpatialAudioInfo) SpatialAudioMixMetadata() foundation.objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](c_.ID, objc.Sel("spatialAudioMixMetadata"))
	return rv
}



