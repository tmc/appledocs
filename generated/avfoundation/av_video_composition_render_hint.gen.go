// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [VideoCompositionRenderHint] class.
var (
	VideoCompositionRenderHintClass     _VideoCompositionRenderHintClass
	VideoCompositionRenderHintClassOnce sync.Once
)

func getVideoCompositionRenderHintClass() _VideoCompositionRenderHintClass {
	VideoCompositionRenderHintClassOnce.Do(func() {
		VideoCompositionRenderHintClass = _VideoCompositionRenderHintClass{objc.GetClass("AVVideoCompositionRenderHint")}
	})
	return VideoCompositionRenderHintClass
}

type _VideoCompositionRenderHintClass struct {
	class objc.Class
}





// An interface definition for the [VideoCompositionRenderHint] class.
type IVideoCompositionRenderHint interface {
	objectivec.IObject
	

	// properties:
	EndCompositionTime() objc.IObject /* cross-framework: Time */
	StartCompositionTime() objc.IObject /* cross-framework: Time */


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (vc _VideoCompositionRenderHintClass) Alloc() VideoCompositionRenderHint {
	rv := objc.Send[VideoCompositionRenderHint](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (vc _VideoCompositionRenderHintClass) New() VideoCompositionRenderHint {
	rv := objc.Send[VideoCompositionRenderHint](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ VideoCompositionRenderHint) Init() VideoCompositionRenderHint {
	rv := objc.Send[VideoCompositionRenderHint](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ VideoCompositionRenderHint) Autorelease() VideoCompositionRenderHint {
	rv := objc.Send[VideoCompositionRenderHint](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewVideoCompositionRenderHint creates a new VideoCompositionRenderHint instance.
func NewVideoCompositionRenderHint() VideoCompositionRenderHint {
	return getVideoCompositionRenderHintClass().New()
}





// Information about upcoming composition requests, such as composition start time and end time.


// Information about upcoming composition requests, such as composition start time and end time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint
type VideoCompositionRenderHint struct {
	objectivec.Object
}

// VideoCompositionRenderHintFrom constructs a [VideoCompositionRenderHint] from an unsafe.Pointer.
//
// Information about upcoming composition requests, such as composition start time and end time.
func VideoCompositionRenderHintFrom(ptr unsafe.Pointer) VideoCompositionRenderHint {
	return VideoCompositionRenderHint{objectivec.Object{objc.ID(ptr)}}
}

























// The end time of the upcoming composition requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint/endCompositionTime
func (v_ VideoCompositionRenderHint) EndCompositionTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](v_.ID, objc.Sel("endCompositionTime"))
	return rv
}


// The start time of the upcoming composition requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionRenderHint/startCompositionTime
func (v_ VideoCompositionRenderHint) StartCompositionTime() objc.IObject /* cross-framework: Time */ {
	rv := objc.Send[corevideo.Time](v_.ID, objc.Sel("startCompositionTime"))
	return rv
}








