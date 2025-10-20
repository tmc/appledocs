// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [Layer] class.
var (
	layerClass     _LayerClass
	layerClassOnce sync.Once
)

func getLayerClass() _LayerClass {
	layerClassOnce.Do(func() {
		layerClass = _LayerClass{objc.GetClass("CALayer")}
	})
	return layerClass
}

type _LayerClass struct {
	class objc.Class
}

// An interface definition for the [Layer] class.
type ILayer interface {
	objectivec.IObject
	DrawInContext(ctx coregraphics.CGContextRef)
	HitTest(p coregraphics.CGPoint) unsafe.Pointer
	SetNeedsDisplay()
}

// An object that manages image-based content and allows you to perform animations on that content.
//
// Layers are often used to provide the backing store for views but can also be used without a view to display content. A layer’s main job is to manage the visual content that you provide but the layer itself has visual attributes that can be set, such as a background color, border, and shadow. In addition to managing visual content, the layer also maintains information about the geometry of its content (such as its position, size, and transform) that is used to present that content onscreen. Modifying the properties of the layer is how you initiate animations on the layer’s content or geometry. A layer object encapsulates the duration and pacing of a layer and its animations by adopting the protocol, which defines the layer’s timing information. If the layer object was created by a view, the view typically assigns itself as the layer’s delegate automatically, and you should not change that relationship. For layers you create yourself, you can assign a object and use that object to provide the contents of the layer dynamically and perform other tasks. A layer may also have a layout manager object (assigned to the property) to manage the layout of subviews separately.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer
type Layer struct {
	objectivec.Object
}

// LayerFrom constructs a [Layer] from an unsafe.Pointer.
//
// An object that manages image-based content and allows you to perform animations on that content.
func LayerFrom(ptr unsafe.Pointer) Layer {
	return Layer{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayerClass) Alloc() Layer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayerClass) New() Layer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ Layer) Init() Layer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ Layer) Autorelease() Layer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayer creates a new Layer instance.
func NewLayer() Layer {
	return getLayerClass().New()
}


// Draws the layer’s content using the specified graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/draw(in:)
func (l_ Layer) DrawInContext(ctx coregraphics.CGContextRef) {
	objc.Send[objc.ID](l_.ID, objc.Sel("drawInContext:"), ctx)
}

// Returns the farthest descendant of the receiver in the layer hierarchy (including itself) that contains the specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/hitTest(_:)
func (l_ Layer) HitTest(p coregraphics.CGPoint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("hitTest:"), p)
	return rv
}

// Marks the layer’s contents as needing to be updated.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsDisplay()
func (l_ Layer) SetNeedsDisplay() {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsDisplay"))
}

// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (l_ Layer) Contents() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("contents"))
	return rv
}

// SetContents sets the value of the contents property.
// An object that provides the contents of the layer. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (l_ Layer) SetContents(value objc.ID) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContents:"), value)
}
// The layer’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/delegate
func (l_ Layer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("delegate"))
	return rv
}

// SetDelegate sets the value of the delegate property.
// The layer’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/delegate
func (l_ Layer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelegate:"), value)
}
// The object responsible for laying out the layer’s sublayers.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutManager
func (l_ Layer) LayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("layoutManager"))
	return rv
}

// SetLayoutManager sets the value of the layoutManager property.
// The object responsible for laying out the layer’s sublayers.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutManager
func (l_ Layer) SetLayoutManager(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLayoutManager:"), value)
}
// The opacity of the receiver. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/opacity
func (l_ Layer) Opacity() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("opacity"))
	return rv
}

// SetOpacity sets the value of the opacity property.
// The opacity of the receiver. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/opacity
func (l_ Layer) SetOpacity(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOpacity:"), value)
}
// The transform applied to the layer’s contents. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/transform
func (l_ Layer) Transform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("transform"))
	return rv
}

// SetTransform sets the value of the transform property.
// The transform applied to the layer’s contents. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/transform
func (l_ Layer) SetTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTransform:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsDynamicContentScaling
func (l_ Layer) WantsDynamicContentScaling() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("wantsDynamicContentScaling"))
	return rv
}

// SetWantsDynamicContentScaling sets the value of the wantsDynamicContentScaling property.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsDynamicContentScaling
func (l_ Layer) SetWantsDynamicContentScaling(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWantsDynamicContentScaling:"), value)
}


