// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptionRendererScene] class.
var (
	CaptionRendererSceneClass     _CaptionRendererSceneClass
	CaptionRendererSceneClassOnce sync.Once
)

func getCaptionRendererSceneClass() _CaptionRendererSceneClass {
	CaptionRendererSceneClassOnce.Do(func() {
		CaptionRendererSceneClass = _CaptionRendererSceneClass{objc.GetClass("AVCaptionRendererScene")}
	})
	return CaptionRendererSceneClass
}

type _CaptionRendererSceneClass struct {
	class objc.Class
}





// An interface definition for the [CaptionRendererScene] class.
type ICaptionRendererScene interface {
	objectivec.IObject
	

	// properties:
	HasActiveCaptions() bool
	NeedsPeriodicRefresh() bool
	TimeRange() TimeRange /* not a class type */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CaptionRendererSceneClass) Alloc() CaptionRendererScene {
	rv := objc.Send[CaptionRendererScene](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionRendererSceneClass) New() CaptionRendererScene {
	rv := objc.Send[CaptionRendererScene](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionRendererScene) Init() CaptionRendererScene {
	rv := objc.Send[CaptionRendererScene](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionRendererScene) Autorelease() CaptionRendererScene {
	rv := objc.Send[CaptionRendererScene](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionRendererScene creates a new CaptionRendererScene instance.
func NewCaptionRendererScene() CaptionRendererScene {
	return getCaptionRendererSceneClass().New()
}





// An object that holds a time range and an associated state which indicates when the renderer draws output.
//
// To render a scene, the object considers state like the existence of captions and regions, their temporal overlaps, and whether captions use animation effects. Your app can request time ranges where visual differences exist and use these time ranges to optimize drawing performance, like drawing once per scene. Alternatively, it can ignore scenes, and instead call repeatedly, but this may have additional performance impact.


// An object that holds a time range and an associated state which indicates when the renderer draws output.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene
type CaptionRendererScene struct {
	objectivec.Object
}

// CaptionRendererSceneFrom constructs a [CaptionRendererScene] from an unsafe.Pointer.
//
// An object that holds a time range and an associated state which indicates when the renderer draws output.
func CaptionRendererSceneFrom(ptr unsafe.Pointer) CaptionRendererScene {
	return CaptionRendererScene{objectivec.Object{objc.ID(ptr)}}
}

























// A Boolean value that indicates whether the scene contains one or more active captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene/hasActiveCaptions
func (c_ CaptionRendererScene) HasActiveCaptions() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasActiveCaptions"))
	return rv
}


// A Boolean value that indicates whether the scene requires redrawing while your app progresses through the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene/needsPeriodicRefresh
func (c_ CaptionRendererScene) NeedsPeriodicRefresh() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("needsPeriodicRefresh"))
	return rv
}


// The time range during which the system doesn’t modify the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene/timeRange
func (c_ CaptionRendererScene) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRange"))
	return rv
}








