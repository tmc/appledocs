// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CaptionRenderer] class.
var (
	CaptionRendererClass     _CaptionRendererClass
	CaptionRendererClassOnce sync.Once
)

func getCaptionRendererClass() _CaptionRendererClass {
	CaptionRendererClassOnce.Do(func() {
		CaptionRendererClass = _CaptionRendererClass{objc.GetClass("AVCaptionRenderer")}
	})
	return CaptionRendererClass
}

type _CaptionRendererClass struct {
	class objc.Class
}





// An interface definition for the [CaptionRenderer] class.
type ICaptionRenderer interface {
	objectivec.IObject
	

	// properties:
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	Captions() []Caption
	SetCaptions(value []Caption)


	

	// methods:
	CaptionSceneChangesInRange(consideredTimeRange objectivec.IObject) []CaptionRendererScene
	RenderInContextForTime(ctx ContextRef /* not a class type */, time objectivec.IObject)


}





// Alloc allocates a new instance without initialization.
func (cc _CaptionRendererClass) Alloc() CaptionRenderer {
	rv := objc.Send[CaptionRenderer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CaptionRendererClass) New() CaptionRenderer {
	rv := objc.Send[CaptionRenderer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CaptionRenderer) Init() CaptionRenderer {
	rv := objc.Send[CaptionRenderer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CaptionRenderer) Autorelease() CaptionRenderer {
	rv := objc.Send[CaptionRenderer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCaptionRenderer creates a new CaptionRenderer instance.
func NewCaptionRenderer() CaptionRenderer {
	return getCaptionRendererClass().New()
}





// An object that renders captions for display at a particular time.
//
// This object renders a caption scene for a given time from a collection of captions. If there aren’t any captions to display at the specified time, the renderer draws an empty flood fill with a zero alpha or a color.


// An object that renders captions for display at a particular time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer
type CaptionRenderer struct {
	objectivec.Object
}

// CaptionRendererFrom constructs a [CaptionRenderer] from an unsafe.Pointer.
//
// An object that renders captions for display at a particular time.
func CaptionRendererFrom(ptr unsafe.Pointer) CaptionRenderer {
	return CaptionRenderer{objectivec.Object{objc.ID(ptr)}}
}




















// Determine render time ranges within an enclosing time range to account for visual changes among captions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/captionSceneChanges(in:)
func (c_ CaptionRenderer) CaptionSceneChangesInRange(consideredTimeRange objectivec.IObject) []CaptionRendererScene {
	rv := objc.Send[[]CaptionRendererScene](c_.ID, objc.Sel("captionSceneChangesInRange:"), consideredTimeRange)
	return rv
}


// Draw the captions for the time you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/render(in:for:)
func (c_ CaptionRenderer) RenderInContextForTime(ctx ContextRef /* not a class type */, time objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("renderInContext:forTime:"), ctx, time)
}







// The drawing bounds of caption scenes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/bounds
func (c_ CaptionRenderer) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c_.ID, objc.Sel("bounds"))
	return rv
}


// The drawing bounds of caption scenes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/bounds
func (c_ CaptionRenderer) SetBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBounds:"), value)
}


// The captions to render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/captions
func (c_ CaptionRenderer) Captions() []Caption {
	rv := objc.Send[[]Caption](c_.ID, objc.Sel("captions"))
	return rv
}


// The captions to render.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/captions
func (c_ CaptionRenderer) SetCaptions(value []Caption) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](c_.ID, objc.Sel("setCaptions:"), nsArray)
}








