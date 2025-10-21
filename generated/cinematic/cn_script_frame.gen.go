// Code generated from Apple documentation for Cinematic. DO NOT EDIT.

package cinematic

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CNScriptFrame] class.
var (
	CNScriptFrameClass     _CNScriptFrameClass
	CNScriptFrameClassOnce sync.Once
)

func getCNScriptFrameClass() _CNScriptFrameClass {
	CNScriptFrameClassOnce.Do(func() {
		CNScriptFrameClass = _CNScriptFrameClass{objc.GetClass("CNScriptFrame")}
	})
	return CNScriptFrameClass
}

type _CNScriptFrameClass struct {
	class objc.Class
}

// An interface definition for the [CNScriptFrame] class.
type ICNScriptFrame interface {
	objectivec.IObject
	DetectionForID(detectionID unsafe.Pointer) unsafe.Pointer
}

// An object that represents what to focus on, and where to focus, in a given movie frame.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame
type CNScriptFrame struct {
	objectivec.Object
}

// CNScriptFrameFrom constructs a [CNScriptFrame] from an unsafe.Pointer.
//
// An object that represents what to focus on, and where to focus, in a given movie frame.
func CNScriptFrameFrom(ptr unsafe.Pointer) CNScriptFrame {
	return CNScriptFrame{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CNScriptFrameClass) Alloc() CNScriptFrame {
	rv := objc.Send[CNScriptFrame](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CNScriptFrameClass) New() CNScriptFrame {
	rv := objc.Send[CNScriptFrame](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNScriptFrame) Init() CNScriptFrame {
	rv := objc.Send[CNScriptFrame](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNScriptFrame) Autorelease() CNScriptFrame {
	rv := objc.Send[CNScriptFrame](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNScriptFrame creates a new CNScriptFrame instance.
func NewCNScriptFrame() CNScriptFrame {
	return getCNScriptFrameClass().New()
}


// The detection in the frame with the given detection ID, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/detectionForID:
func (c_ CNScriptFrame) DetectionForID(detectionID unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("detectionForID:"), detectionID)
	return rv
}

// All detections for the Cinematic movie.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/allDetections
func (c_ CNScriptFrame) AllDetections() []CNDetection {
	rv := objc.Send[[]CNDetection](c_.ID, objc.Sel("allDetections"))
	return rv
}

// What to focus on in a given frame of the movie.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/focusDetection
func (c_ CNScriptFrame) FocusDetection() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("focusDetection"))
	return rv
}

// Where to focus in a given frame of the movie.
//
// [Full Topic]: https://developer.apple.com/documentation/Cinematic/CNScriptFrame/focusDisparity
func (c_ CNScriptFrame) FocusDisparity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("focusDisparity"))
	return rv
}




