// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class AVCaptionRendererScene */


/* debug [class_header]: Header for AVCaptionRendererScene */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CaptionRendererScene */
// An interface definition for the [CaptionRendererScene] class.
type ICaptionRendererScene interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CaptionRendererScene */
	// properties:
	HasActiveCaptions() bool
	NeedsPeriodicRefresh() bool
	TimeRange() TimeRange /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CaptionRendererScene */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CaptionRendererScene */
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CaptionRendererScene */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CaptionRendererScene *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CaptionRendererScene */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CaptionRendererScene */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CaptionRendererScene */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CaptionRendererScene */

// A Boolean value that indicates whether the scene contains one or more active captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene/hasActiveCaptions
func (c_ CaptionRendererScene) HasActiveCaptions() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasActiveCaptions"))
	return rv
}/* debug [instance_properties/getter]: hasActiveCaptions */


// A Boolean value that indicates whether the scene requires redrawing while your app progresses through the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene/needsPeriodicRefresh
func (c_ CaptionRendererScene) NeedsPeriodicRefresh() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("needsPeriodicRefresh"))
	return rv
}/* debug [instance_properties/getter]: needsPeriodicRefresh */


// The time range during which the system doesn’t modify the scene.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene/timeRange
func (c_ CaptionRendererScene) TimeRange() TimeRange /* not a class type */ {
	rv := objc.Send[TimeRange](c_.ID, objc.Sel("timeRange"))
	return rv
}/* debug [instance_properties/getter]: timeRange */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class AVCaptionRendererScene */



