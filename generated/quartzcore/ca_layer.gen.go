// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coregraphics"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Layer] class.
var (
	LayerClass     _LayerClass
	LayerClassOnce sync.Once
)

func getLayerClass() _LayerClass {
	LayerClassOnce.Do(func() {
		LayerClass = _LayerClass{objc.GetClass("CALayer")}
	})
	return LayerClass
}

type _LayerClass struct {
	class objc.Class
}

// An interface definition for the [Layer] class.
type ILayer interface {
	objectivec.IObject
	ActionForKey(event string) objc.ID
	AddAnimationForKey(anim unsafe.Pointer, key string)
	AddConstraint(c unsafe.Pointer)
	AddSublayer(layer unsafe.Pointer)
	AffineTransform() coregraphics.CGAffineTransform
	AnimationForKey(key string) unsafe.Pointer
	AnimationKeys() []string
	ContainsPoint(p coregraphics.CGPoint) bool
	ContentsAreFlipped() bool
	ConvertRectFromLayer(r coregraphics.CGRect, l unsafe.Pointer) coregraphics.CGRect
	ConvertPointFromLayer(p coregraphics.CGPoint, l unsafe.Pointer) coregraphics.CGPoint
	ConvertPointToLayer(p coregraphics.CGPoint, l unsafe.Pointer) coregraphics.CGPoint
	ConvertRectToLayer(r coregraphics.CGRect, l unsafe.Pointer) coregraphics.CGRect
	ConvertTimeFromLayer(t unsafe.Pointer, l unsafe.Pointer) unsafe.Pointer
	ConvertTimeToLayer(t unsafe.Pointer, l unsafe.Pointer) unsafe.Pointer
	Display()
	DisplayIfNeeded()
	DrawInContext(ctx coregraphics.CGContextRef)
	HitTest(p coregraphics.CGPoint) unsafe.Pointer
	InsertSublayerAbove(layer unsafe.Pointer, sibling unsafe.Pointer)
	InsertSublayerAtIndex(layer unsafe.Pointer, idx unsafe.Pointer)
	InsertSublayerBelow(layer unsafe.Pointer, sibling unsafe.Pointer)
	LayoutIfNeeded()
	LayoutSublayers()
	ModelLayer() unsafe.Pointer
	NeedsDisplay() bool
	NeedsLayout() bool
	PreferredFrameSize() coregraphics.CGSize
	PresentationLayer() unsafe.Pointer
	RemoveAllAnimations()
	RemoveAnimationForKey(key string)
	RemoveFromSuperlayer()
	RenderInContext(ctx coregraphics.CGContextRef)
	ReplaceSublayerWith(oldLayer unsafe.Pointer, newLayer unsafe.Pointer)
	ResizeWithOldSuperlayerSize(size coregraphics.CGSize)
	ResizeSublayersWithOldSize(size coregraphics.CGSize)
	ScrollPoint(p coregraphics.CGPoint)
	ScrollRectToVisible(r coregraphics.CGRect)
	SetAffineTransform(m coregraphics.CGAffineTransform)
	SetNeedsDisplay()
	SetNeedsDisplayInRect(r coregraphics.CGRect)
	SetNeedsLayout()
	ShouldArchiveValueForKey(key string) bool
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




// Override to copy or initialize custom fields of the specified layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/init(layer:)
func NewLayerWithLayer(layer objc.ID) Layer {
	instance := getLayerClass().Alloc()
	rv := objc.Send[Layer](instance.ID, objc.Sel("initWithLayer:"), layer)
	rv.Autorelease()
	return rv
}



// Initializes a layer with a remote client ID.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/init(remoteClientId:)
func NewLayerWithRemoteClientId(client_id unsafe.Pointer) Layer {
	rv := objc.Send[Layer](objc.ID(getLayerClass().class), objc.Sel("layerWithRemoteClientId:"), client_id)
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerCurveExpansionFactor(_:)
func (lc _LayerClass) CornerCurveExpansionFactor(curve unsafe.Pointer) float64 {
	rv := objc.Send[float64](objc.ID(lc.class), objc.Sel("cornerCurveExpansionFactor:"), curve)
	return rv
}

// Returns the default action for the current class.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/defaultAction(forKey:)
func (lc _LayerClass) DefaultActionForKey(event string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("defaultActionForKey:"), objc.String(event))
	return rv
}

// Specifies the default value associated with the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/defaultValue(forKey:)
func (lc _LayerClass) DefaultValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("defaultValueForKey:"), objc.String(key))
	return rv
}

// Initializes a layer with a remote client ID.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/init(remoteClientId:)
func (lc _LayerClass) LayerWithRemoteClientId(client_id unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("layerWithRemoteClientId:"), client_id)
	return rv
}

// Creates and returns an instance of the layer object.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layer
func (lc _LayerClass) Layer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("layer"))
	return rv
}

// Returns a Boolean indicating whether changes to the specified key require the layer to be redisplayed.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplay(forKey:)
func (lc _LayerClass) NeedsDisplayForKey(key string) bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("needsDisplayForKey:"), objc.String(key))
	return rv
}

// Returns the action object assigned to the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/action(forKey:)
func (l_ Layer) ActionForKey(event string) objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("actionForKey:"), objc.String(event))
	return rv
}

// Add the specified animation object to the layer’s render tree.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/add(_:forKey:)
func (l_ Layer) AddAnimationForKey(anim unsafe.Pointer, key string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addAnimation:forKey:"), anim, objc.String(key))
}

// Adds the specified constraint to the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/addConstraint(_:)
func (l_ Layer) AddConstraint(c unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addConstraint:"), c)
}

// Appends the layer to the layer’s list of sublayers.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/addSublayer(_:)
func (l_ Layer) AddSublayer(layer unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addSublayer:"), layer)
}

// Returns an affine version of the layer’s transform.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/affineTransform()
func (l_ Layer) AffineTransform() coregraphics.CGAffineTransform {
	rv := objc.Send[coregraphics.CGAffineTransform](l_.ID, objc.Sel("affineTransform"))
	return rv
}

// Returns the animation object with the specified identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/animation(forKey:)
func (l_ Layer) AnimationForKey(key string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("animationForKey:"), objc.String(key))
	return rv
}

// Returns an array of strings that identify the animations currently attached to the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/animationKeys()
func (l_ Layer) AnimationKeys() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("animationKeys"))
	return rv
}

// Returns whether the receiver contains a specified point.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contains(_:)
func (l_ Layer) ContainsPoint(p coregraphics.CGPoint) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("containsPoint:"), p)
	return rv
}

// Returns a Boolean indicating whether the layer content is implicitly flipped when rendered.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsAreFlipped()
func (l_ Layer) ContentsAreFlipped() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("contentsAreFlipped"))
	return rv
}

// Converts the rectangle from the specified layer’s coordinate system to the receiver’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:from:)-4kx9l
func (l_ Layer) ConvertRectFromLayer(r coregraphics.CGRect, l unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("convertRect:fromLayer:"), r, l)
	return rv
}

// Converts the point from the specified layer’s coordinate system to the receiver’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:from:)-8kl76
func (l_ Layer) ConvertPointFromLayer(p coregraphics.CGPoint, l unsafe.Pointer) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](l_.ID, objc.Sel("convertPoint:fromLayer:"), p, l)
	return rv
}

// Converts the point from the receiver’s coordinate system to the specified layer’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:to:)-7dcke
func (l_ Layer) ConvertPointToLayer(p coregraphics.CGPoint, l unsafe.Pointer) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](l_.ID, objc.Sel("convertPoint:toLayer:"), p, l)
	return rv
}

// Converts the rectangle from the receiver’s coordinate system to the specified layer’s coordinate system.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:to:)-tly5
func (l_ Layer) ConvertRectToLayer(r coregraphics.CGRect, l unsafe.Pointer) coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("convertRect:toLayer:"), r, l)
	return rv
}

// Converts the time interval from the specified layer’s time space to the receiver’s time space.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convertTime(_:from:)
func (l_ Layer) ConvertTimeFromLayer(t unsafe.Pointer, l unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("convertTime:fromLayer:"), t, l)
	return rv
}

// Converts the time interval from the receiver’s time space to the specified layer’s time space
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convertTime(_:to:)
func (l_ Layer) ConvertTimeToLayer(t unsafe.Pointer, l unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("convertTime:toLayer:"), t, l)
	return rv
}

// Reloads the content of this layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/display()
func (l_ Layer) Display() {
	objc.Send[objc.ID](l_.ID, objc.Sel("display"))
}

// Initiates the update process for a layer if it is currently marked as needing an update.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/displayIfNeeded()
func (l_ Layer) DisplayIfNeeded() {
	objc.Send[objc.ID](l_.ID, objc.Sel("displayIfNeeded"))
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

// Inserts the specified sublayer above a different sublayer that already belongs to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/insertSublayer(_:above:)
func (l_ Layer) InsertSublayerAbove(layer unsafe.Pointer, sibling unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("insertSublayer:above:"), layer, sibling)
}

// Inserts the specified layer into the receiver’s list of sublayers at the specified index.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/insertSublayer(_:at:)
func (l_ Layer) InsertSublayerAtIndex(layer unsafe.Pointer, idx unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("insertSublayer:atIndex:"), layer, idx)
}

// Inserts the specified sublayer below a different sublayer that already belongs to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/insertSublayer(_:below:)
func (l_ Layer) InsertSublayerBelow(layer unsafe.Pointer, sibling unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("insertSublayer:below:"), layer, sibling)
}

// Recalculate the receiver’s layout, if required.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutIfNeeded()
func (l_ Layer) LayoutIfNeeded() {
	objc.Send[objc.ID](l_.ID, objc.Sel("layoutIfNeeded"))
}

// Tells the layer to update its layout.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutSublayers()
func (l_ Layer) LayoutSublayers() {
	objc.Send[objc.ID](l_.ID, objc.Sel("layoutSublayers"))
}

// Returns the model layer object associated with the receiver, if any.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/model()
func (l_ Layer) ModelLayer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("modelLayer"))
	return rv
}

// Returns a Boolean indicating whether the layer has been marked as needing an update.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplay()
func (l_ Layer) NeedsDisplay() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("needsDisplay"))
	return rv
}

// Returns a Boolean indicating whether the layer has been marked as needing a layout update.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsLayout()
func (l_ Layer) NeedsLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("needsLayout"))
	return rv
}

// Returns the preferred size of the layer in the coordinate space of its superlayer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/preferredFrameSize()
func (l_ Layer) PreferredFrameSize() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](l_.ID, objc.Sel("preferredFrameSize"))
	return rv
}

// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/presentation()
func (l_ Layer) PresentationLayer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("presentationLayer"))
	return rv
}

// Remove all animations attached to the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/removeAllAnimations()
func (l_ Layer) RemoveAllAnimations() {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeAllAnimations"))
}

// Remove the animation object with the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/removeAnimation(forKey:)
func (l_ Layer) RemoveAnimationForKey(key string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeAnimationForKey:"), objc.String(key))
}

// Detaches the layer from its parent layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/removeFromSuperlayer()
func (l_ Layer) RemoveFromSuperlayer() {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeFromSuperlayer"))
}

// Renders the layer and its sublayers into the specified context.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/render(in:)
func (l_ Layer) RenderInContext(ctx coregraphics.CGContextRef) {
	objc.Send[objc.ID](l_.ID, objc.Sel("renderInContext:"), ctx)
}

// Replaces the specified sublayer with a different layer object.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/replaceSublayer(_:with:)
func (l_ Layer) ReplaceSublayerWith(oldLayer unsafe.Pointer, newLayer unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("replaceSublayer:with:"), oldLayer, newLayer)
}

// Informs the receiver that the size of its superlayer changed.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/resize(withOldSuperlayerSize:)
func (l_ Layer) ResizeWithOldSuperlayerSize(size coregraphics.CGSize) {
	objc.Send[objc.ID](l_.ID, objc.Sel("resizeWithOldSuperlayerSize:"), size)
}

// Informs the receiver’s sublayers that the receiver’s size has changed.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/resizeSublayers(withOldSize:)
func (l_ Layer) ResizeSublayersWithOldSize(size coregraphics.CGSize) {
	objc.Send[objc.ID](l_.ID, objc.Sel("resizeSublayersWithOldSize:"), size)
}

// Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified point lies at the origin of the scroll layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/scroll(_:)
func (l_ Layer) ScrollPoint(p coregraphics.CGPoint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("scrollPoint:"), p)
}

// Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified rectangle becomes visible.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/scrollRectToVisible(_:)
func (l_ Layer) ScrollRectToVisible(r coregraphics.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("scrollRectToVisible:"), r)
}

// Sets the layer’s transform to the specified affine transform.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setAffineTransform(_:)
func (l_ Layer) SetAffineTransform(m coregraphics.CGAffineTransform) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAffineTransform:"), m)
}

// Marks the layer’s contents as needing to be updated.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsDisplay()
func (l_ Layer) SetNeedsDisplay() {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsDisplay"))
}

// Marks the region within the specified rectangle as needing to be updated.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsDisplay(_:)
func (l_ Layer) SetNeedsDisplayInRect(r coregraphics.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsDisplayInRect:"), r)
}

// Invalidates the layer’s layout and marks it as needing an update.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsLayout()
func (l_ Layer) SetNeedsLayout() {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsLayout"))
}

// Returns a Boolean indicating whether the value of the specified key should be archived.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shouldArchiveValue(forKey:)
func (l_ Layer) ShouldArchiveValueForKey(key string) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("shouldArchiveValueForKey:"), objc.String(key))
	return rv
}

// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isdoublesided
func (l_ Layer) IsDoubleSided() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isDoubleSided"))
	return rv
}


// SetIsDoubleSided sets the value of the isDoubleSided property.
// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isdoublesided
func (l_ Layer) SetIsDoubleSided(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsDoubleSided:"), value)
}

// A Boolean value indicating whether the layer contains completely opaque content.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isopaque
func (l_ Layer) IsOpaque() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isOpaque"))
	return rv
}


// SetIsOpaque sets the value of the isOpaque property.
// A Boolean value indicating whether the layer contains completely opaque content.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isopaque
func (l_ Layer) SetIsOpaque(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsOpaque:"), value)
}

// A Boolean indicating whether the layer is displayed. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/ishidden
func (l_ Layer) IsHidden() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isHidden"))
	return rv
}


// SetIsHidden sets the value of the isHidden property.
// A Boolean indicating whether the layer is displayed. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/ishidden
func (l_ Layer) SetIsHidden(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsHidden:"), value)
}

// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isgeometryflipped
func (l_ Layer) IsGeometryFlipped() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isGeometryFlipped"))
	return rv
}


// SetIsGeometryFlipped sets the value of the isGeometryFlipped property.
// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.

//
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isgeometryflipped
func (l_ Layer) SetIsGeometryFlipped(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsGeometryFlipped:"), value)
}

// A dictionary containing layer actions.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/actions
func (l_ Layer) Actions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("actions"))
	return rv
}


// SetActions sets the value of the actions property.
// A dictionary containing layer actions.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/actions
func (l_ Layer) SetActions(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setActions:"), value)
}

// A Boolean indicating whether the layer is allowed to perform edge antialiasing.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsEdgeAntialiasing
func (l_ Layer) AllowsEdgeAntialiasing() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("allowsEdgeAntialiasing"))
	return rv
}


// SetAllowsEdgeAntialiasing sets the value of the allowsEdgeAntialiasing property.
// A Boolean indicating whether the layer is allowed to perform edge antialiasing.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsEdgeAntialiasing
func (l_ Layer) SetAllowsEdgeAntialiasing(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAllowsEdgeAntialiasing:"), value)
}

// A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsGroupOpacity
func (l_ Layer) AllowsGroupOpacity() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("allowsGroupOpacity"))
	return rv
}


// SetAllowsGroupOpacity sets the value of the allowsGroupOpacity property.
// A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsGroupOpacity
func (l_ Layer) SetAllowsGroupOpacity(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAllowsGroupOpacity:"), value)
}

// Defines the anchor point of the layer’s bounds rectangle. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPoint
func (l_ Layer) AnchorPoint() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](l_.ID, objc.Sel("anchorPoint"))
	return rv
}


// SetAnchorPoint sets the value of the anchorPoint property.
// Defines the anchor point of the layer’s bounds rectangle. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPoint
func (l_ Layer) SetAnchorPoint(value coregraphics.CGPoint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAnchorPoint:"), value)
}

// The anchor point for the layer’s position along the z axis. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPointZ
func (l_ Layer) AnchorPointZ() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("anchorPointZ"))
	return rv
}


// SetAnchorPointZ sets the value of the anchorPointZ property.
// The anchor point for the layer’s position along the z axis. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPointZ
func (l_ Layer) SetAnchorPointZ(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAnchorPointZ:"), value)
}

// A bitmask defining how the layer is resized when the bounds of its superlayer changes.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/autoresizingMask
func (l_ Layer) AutoresizingMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("autoresizingMask"))
	return rv
}


// SetAutoresizingMask sets the value of the autoresizingMask property.
// A bitmask defining how the layer is resized when the bounds of its superlayer changes.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/autoresizingMask
func (l_ Layer) SetAutoresizingMask(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAutoresizingMask:"), value)
}

// The background color of the receiver. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundColor
func (l_ Layer) BackgroundColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](l_.ID, objc.Sel("backgroundColor"))
	return rv
}


// SetBackgroundColor sets the value of the backgroundColor property.
// The background color of the receiver. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundColor
func (l_ Layer) SetBackgroundColor(value coregraphics.CGColorRef) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBackgroundColor:"), value)
}

// An array of Core Image filters to apply to the content immediately behind the layer. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundFilters
func (l_ Layer) BackgroundFilters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("backgroundFilters"))
	return rv
}


// SetBackgroundFilters sets the value of the backgroundFilters property.
// An array of Core Image filters to apply to the content immediately behind the layer. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundFilters
func (l_ Layer) SetBackgroundFilters(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBackgroundFilters:"), value)
}

// The color of the layer’s border. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/borderColor
func (l_ Layer) BorderColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](l_.ID, objc.Sel("borderColor"))
	return rv
}


// SetBorderColor sets the value of the borderColor property.
// The color of the layer’s border. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/borderColor
func (l_ Layer) SetBorderColor(value coregraphics.CGColorRef) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBorderColor:"), value)
}

// The width of the layer’s border. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/borderWidth
func (l_ Layer) BorderWidth() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("borderWidth"))
	return rv
}


// SetBorderWidth sets the value of the borderWidth property.
// The width of the layer’s border. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/borderWidth
func (l_ Layer) SetBorderWidth(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBorderWidth:"), value)
}

// The layer’s bounds rectangle. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/bounds
func (l_ Layer) Bounds() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("bounds"))
	return rv
}


// SetBounds sets the value of the bounds property.
// The layer’s bounds rectangle. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/bounds
func (l_ Layer) SetBounds(value coregraphics.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBounds:"), value)
}

// A CoreImage filter used to composite the layer and the content behind it. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/compositingFilter
func (l_ Layer) CompositingFilter() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("compositingFilter"))
	return rv
}


// SetCompositingFilter sets the value of the compositingFilter property.
// A CoreImage filter used to composite the layer and the content behind it. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/compositingFilter
func (l_ Layer) SetCompositingFilter(value objc.ID) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCompositingFilter:"), value)
}

// The constraints used to position current layer’s sublayers.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/constraints
func (l_ Layer) Constraints() []Constraint {
	rv := objc.Send[[]Constraint](l_.ID, objc.Sel("constraints"))
	return rv
}


// SetConstraints sets the value of the constraints property.
// The constraints used to position current layer’s sublayers.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/constraints
func (l_ Layer) SetConstraints(value []Constraint) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](l_.ID, objc.Sel("setConstraints:"), nsArray)
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

// The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsCenter
func (l_ Layer) ContentsCenter() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("contentsCenter"))
	return rv
}


// SetContentsCenter sets the value of the contentsCenter property.
// The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsCenter
func (l_ Layer) SetContentsCenter(value coregraphics.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsCenter:"), value)
}

// A hint for the desired storage format of the layer contents.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsFormat
func (l_ Layer) ContentsFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("contentsFormat"))
	return rv
}


// SetContentsFormat sets the value of the contentsFormat property.
// A hint for the desired storage format of the layer contents.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsFormat
func (l_ Layer) SetContentsFormat(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsFormat:"), value)
}

// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (l_ Layer) ContentsGravity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("contentsGravity"))
	return rv
}


// SetContentsGravity sets the value of the contentsGravity property.
// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (l_ Layer) SetContentsGravity(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsGravity:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsHeadroom
func (l_ Layer) ContentsHeadroom() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("contentsHeadroom"))
	return rv
}


// SetContentsHeadroom sets the value of the contentsHeadroom property.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsHeadroom
func (l_ Layer) SetContentsHeadroom(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsHeadroom:"), value)
}

// The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsRect
func (l_ Layer) ContentsRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("contentsRect"))
	return rv
}


// SetContentsRect sets the value of the contentsRect property.
// The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsRect
func (l_ Layer) SetContentsRect(value coregraphics.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsRect:"), value)
}

// The scale factor applied to the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
func (l_ Layer) ContentsScale() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("contentsScale"))
	return rv
}


// SetContentsScale sets the value of the contentsScale property.
// The scale factor applied to the layer.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
func (l_ Layer) SetContentsScale(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsScale:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerCurve
func (l_ Layer) CornerCurve() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("cornerCurve"))
	return rv
}


// SetCornerCurve sets the value of the cornerCurve property.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerCurve
func (l_ Layer) SetCornerCurve(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCornerCurve:"), value)
}

// The radius to use when drawing rounded corners for the layer’s background. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerRadius
func (l_ Layer) CornerRadius() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("cornerRadius"))
	return rv
}


// SetCornerRadius sets the value of the cornerRadius property.
// The radius to use when drawing rounded corners for the layer’s background. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerRadius
func (l_ Layer) SetCornerRadius(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCornerRadius:"), value)
}

// The layer’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/delegate
func (l_ Layer) Delegate() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The layer’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/delegate
func (l_ Layer) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/drawsAsynchronously
func (l_ Layer) DrawsAsynchronously() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("drawsAsynchronously"))
	return rv
}


// SetDrawsAsynchronously sets the value of the drawsAsynchronously property.
// A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/drawsAsynchronously
func (l_ Layer) SetDrawsAsynchronously(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDrawsAsynchronously:"), value)
}

// A bitmask defining how the edges of the receiver are rasterized.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/edgeAntialiasingMask
func (l_ Layer) EdgeAntialiasingMask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("edgeAntialiasingMask"))
	return rv
}


// SetEdgeAntialiasingMask sets the value of the edgeAntialiasingMask property.
// A bitmask defining how the edges of the receiver are rasterized.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/edgeAntialiasingMask
func (l_ Layer) SetEdgeAntialiasingMask(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEdgeAntialiasingMask:"), value)
}

// An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/filters
func (l_ Layer) Filters() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("filters"))
	return rv
}


// SetFilters sets the value of the filters property.
// An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/filters
func (l_ Layer) SetFilters(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setFilters:"), value)
}

// The layer’s frame rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/frame
func (l_ Layer) Frame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("frame"))
	return rv
}


// SetFrame sets the value of the frame property.
// The layer’s frame rectangle.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/frame
func (l_ Layer) SetFrame(value coregraphics.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setFrame:"), value)
}

// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isDoubleSided
func (l_ Layer) DoubleSided() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("doubleSided"))
	return rv
}


// SetDoubleSided sets the value of the doubleSided property.
// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isDoubleSided
func (l_ Layer) SetDoubleSided(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDoubleSided:"), value)
}

// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isGeometryFlipped
func (l_ Layer) GeometryFlipped() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("geometryFlipped"))
	return rv
}


// SetGeometryFlipped sets the value of the geometryFlipped property.
// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isGeometryFlipped
func (l_ Layer) SetGeometryFlipped(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGeometryFlipped:"), value)
}

// A Boolean indicating whether the layer is displayed. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isHidden
func (l_ Layer) Hidden() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("hidden"))
	return rv
}


// SetHidden sets the value of the hidden property.
// A Boolean indicating whether the layer is displayed. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isHidden
func (l_ Layer) SetHidden(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHidden:"), value)
}

// A Boolean value indicating whether the layer contains completely opaque content.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isOpaque
func (l_ Layer) Opaque() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("opaque"))
	return rv
}


// SetOpaque sets the value of the opaque property.
// A Boolean value indicating whether the layer contains completely opaque content.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isOpaque
func (l_ Layer) SetOpaque(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOpaque:"), value)
}

// The object responsible for laying out the layer’s sublayers.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutManager
func (l_ Layer) LayoutManager() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("layoutManager"))
	return rv
}


// SetLayoutManager sets the value of the layoutManager property.
// The object responsible for laying out the layer’s sublayers.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutManager
func (l_ Layer) SetLayoutManager(value objc.ID) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLayoutManager:"), value)
}

// The filter used when increasing the size of the content.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/magnificationFilter
func (l_ Layer) MagnificationFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("magnificationFilter"))
	return rv
}


// SetMagnificationFilter sets the value of the magnificationFilter property.
// The filter used when increasing the size of the content.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/magnificationFilter
func (l_ Layer) SetMagnificationFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMagnificationFilter:"), value)
}

// An optional layer whose alpha channel is used to mask the layer’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/mask
func (l_ Layer) Mask() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("mask"))
	return rv
}


// SetMask sets the value of the mask property.
// An optional layer whose alpha channel is used to mask the layer’s content.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/mask
func (l_ Layer) SetMask(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMask:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/maskedCorners
func (l_ Layer) MaskedCorners() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("maskedCorners"))
	return rv
}


// SetMaskedCorners sets the value of the maskedCorners property.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/maskedCorners
func (l_ Layer) SetMaskedCorners(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMaskedCorners:"), value)
}

// A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/masksToBounds
func (l_ Layer) MasksToBounds() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("masksToBounds"))
	return rv
}


// SetMasksToBounds sets the value of the masksToBounds property.
// A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/masksToBounds
func (l_ Layer) SetMasksToBounds(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMasksToBounds:"), value)
}

// The filter used when reducing the size of the content.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilter
func (l_ Layer) MinificationFilter() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("minificationFilter"))
	return rv
}


// SetMinificationFilter sets the value of the minificationFilter property.
// The filter used when reducing the size of the content.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilter
func (l_ Layer) SetMinificationFilter(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMinificationFilter:"), value)
}

// The bias factor used by the minification filter to determine the levels of detail.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilterBias
func (l_ Layer) MinificationFilterBias() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("minificationFilterBias"))
	return rv
}


// SetMinificationFilterBias sets the value of the minificationFilterBias property.
// The bias factor used by the minification filter to determine the levels of detail.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilterBias
func (l_ Layer) SetMinificationFilterBias(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMinificationFilterBias:"), value)
}

// The name of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/name
func (l_ Layer) Name() string {
	rv := objc.Send[string](l_.ID, objc.Sel("name"))
	return rv
}


// SetName sets the value of the name property.
// The name of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/name
func (l_ Layer) SetName(value string) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setName:"), objc.String(value))
}

// A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplayOnBoundsChange
func (l_ Layer) NeedsDisplayOnBoundsChange() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("needsDisplayOnBoundsChange"))
	return rv
}


// SetNeedsDisplayOnBoundsChange sets the value of the needsDisplayOnBoundsChange property.
// A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplayOnBoundsChange
func (l_ Layer) SetNeedsDisplayOnBoundsChange(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsDisplayOnBoundsChange:"), value)
}

// The opacity of the receiver. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/opacity
func (l_ Layer) Opacity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("opacity"))
	return rv
}


// SetOpacity sets the value of the opacity property.
// The opacity of the receiver. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/opacity
func (l_ Layer) SetOpacity(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOpacity:"), value)
}

// The layer’s position in its superlayer’s coordinate space. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/position
func (l_ Layer) Position() coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](l_.ID, objc.Sel("position"))
	return rv
}


// SetPosition sets the value of the position property.
// The layer’s position in its superlayer’s coordinate space. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/position
func (l_ Layer) SetPosition(value coregraphics.CGPoint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPosition:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/preferredDynamicRange
func (l_ Layer) PreferredDynamicRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("preferredDynamicRange"))
	return rv
}


// SetPreferredDynamicRange sets the value of the preferredDynamicRange property.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/preferredDynamicRange
func (l_ Layer) SetPreferredDynamicRange(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPreferredDynamicRange:"), value)
}

// The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/rasterizationScale
func (l_ Layer) RasterizationScale() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("rasterizationScale"))
	return rv
}


// SetRasterizationScale sets the value of the rasterizationScale property.
// The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/rasterizationScale
func (l_ Layer) SetRasterizationScale(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRasterizationScale:"), value)
}

// The color of the layer’s shadow. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowColor
func (l_ Layer) ShadowColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](l_.ID, objc.Sel("shadowColor"))
	return rv
}


// SetShadowColor sets the value of the shadowColor property.
// The color of the layer’s shadow. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowColor
func (l_ Layer) SetShadowColor(value coregraphics.CGColorRef) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowColor:"), value)
}

// The offset (in points) of the layer’s shadow. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOffset
func (l_ Layer) ShadowOffset() coregraphics.CGSize {
	rv := objc.Send[coregraphics.CGSize](l_.ID, objc.Sel("shadowOffset"))
	return rv
}


// SetShadowOffset sets the value of the shadowOffset property.
// The offset (in points) of the layer’s shadow. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOffset
func (l_ Layer) SetShadowOffset(value coregraphics.CGSize) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowOffset:"), value)
}

// The opacity of the layer’s shadow. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOpacity
func (l_ Layer) ShadowOpacity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("shadowOpacity"))
	return rv
}


// SetShadowOpacity sets the value of the shadowOpacity property.
// The opacity of the layer’s shadow. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOpacity
func (l_ Layer) SetShadowOpacity(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowOpacity:"), value)
}

// The shape of the layer’s shadow. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowPath
func (l_ Layer) ShadowPath() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](l_.ID, objc.Sel("shadowPath"))
	return rv
}


// SetShadowPath sets the value of the shadowPath property.
// The shape of the layer’s shadow. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowPath
func (l_ Layer) SetShadowPath(value coregraphics.CGPathRef) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowPath:"), value)
}

// The blur radius (in points) used to render the layer’s shadow. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowRadius
func (l_ Layer) ShadowRadius() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("shadowRadius"))
	return rv
}


// SetShadowRadius sets the value of the shadowRadius property.
// The blur radius (in points) used to render the layer’s shadow. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowRadius
func (l_ Layer) SetShadowRadius(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowRadius:"), value)
}

// A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shouldRasterize
func (l_ Layer) ShouldRasterize() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("shouldRasterize"))
	return rv
}


// SetShouldRasterize sets the value of the shouldRasterize property.
// A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shouldRasterize
func (l_ Layer) SetShouldRasterize(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShouldRasterize:"), value)
}

// An optional dictionary used to store property values that aren’t explicitly defined by the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/style
func (l_ Layer) Style() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("style"))
	return rv
}


// SetStyle sets the value of the style property.
// An optional dictionary used to store property values that aren’t explicitly defined by the layer.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/style
func (l_ Layer) SetStyle(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setStyle:"), value)
}

// Specifies the transform to apply to sublayers when rendering. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayerTransform
func (l_ Layer) SublayerTransform() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("sublayerTransform"))
	return rv
}


// SetSublayerTransform sets the value of the sublayerTransform property.
// Specifies the transform to apply to sublayers when rendering. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayerTransform
func (l_ Layer) SetSublayerTransform(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSublayerTransform:"), value)
}

// An array containing the layer’s sublayers.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayers
func (l_ Layer) Sublayers() []Layer {
	rv := objc.Send[[]Layer](l_.ID, objc.Sel("sublayers"))
	return rv
}


// SetSublayers sets the value of the sublayers property.
// An array containing the layer’s sublayers.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayers
func (l_ Layer) SetSublayers(value []Layer) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](l_.ID, objc.Sel("setSublayers:"), nsArray)
}

// The superlayer of the layer.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/superlayer
func (l_ Layer) Superlayer() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("superlayer"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/toneMapMode-swift.property
func (l_ Layer) ToneMapMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("toneMapMode"))
	return rv
}


// SetToneMapMode sets the value of the toneMapMode property.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/toneMapMode-swift.property
func (l_ Layer) SetToneMapMode(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setToneMapMode:"), value)
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

// The visible region of the layer in its own coordinate space.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/visibleRect
func (l_ Layer) VisibleRect() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](l_.ID, objc.Sel("visibleRect"))
	return rv
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

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsExtendedDynamicRangeContent
func (l_ Layer) WantsExtendedDynamicRangeContent() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("wantsExtendedDynamicRangeContent"))
	return rv
}


// SetWantsExtendedDynamicRangeContent sets the value of the wantsExtendedDynamicRangeContent property.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsExtendedDynamicRangeContent
func (l_ Layer) SetWantsExtendedDynamicRangeContent(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWantsExtendedDynamicRangeContent:"), value)
}

// The layer’s position on the z axis. Animatable.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/zPosition
func (l_ Layer) ZPosition() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("zPosition"))
	return rv
}


// SetZPosition sets the value of the zPosition property.
// The layer’s position on the z axis. Animatable.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/zPosition
func (l_ Layer) SetZPosition(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setZPosition:"), value)
}


