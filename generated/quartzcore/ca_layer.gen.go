// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CALayer */


/* debug [class_header]: Header for CALayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Layer */
// An interface definition for the [Layer] class.
type ILayer interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Layer */
	// properties:
	Actions() foundation.IDictionary
	SetActions(value foundation.IDictionary)
	AllowsEdgeAntialiasing() bool
	SetAllowsEdgeAntialiasing(value bool)
	AllowsGroupOpacity() bool
	SetAllowsGroupOpacity(value bool)
	AnchorPoint() corefoundation.CGPoint
	SetAnchorPoint(value corefoundation.CGPoint)
	AnchorPointZ() float64
	SetAnchorPointZ(value float64)
	AutoresizingMask() AutoresizingMask
	SetAutoresizingMask(value AutoresizingMask)
	BackgroundColor() ColorRef /* not a class type */
	SetBackgroundColor(value ColorRef /* not a class type */)
	BackgroundFilters() objc.IObject /* cross-framework: NSArray */
	SetBackgroundFilters(value objc.IObject /* cross-framework: NSArray */)
	BorderColor() ColorRef /* not a class type */
	SetBorderColor(value ColorRef /* not a class type */)
	BorderWidth() float64
	SetBorderWidth(value float64)
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)
	CompositingFilter() objc.ID
	SetCompositingFilter(value objc.ID)
	Constraints() []Constraint
	SetConstraints(value []Constraint)
	Contents() objc.ID
	SetContents(value objc.ID)
	ContentsCenter() corefoundation.CGRect
	SetContentsCenter(value corefoundation.CGRect)
	ContentsFormat() LayerContentsFormat /* typedef */
	SetContentsFormat(value LayerContentsFormat /* typedef */)
	ContentsGravity() LayerContentsGravity /* typedef */
	SetContentsGravity(value LayerContentsGravity /* typedef */)
	ContentsHeadroom() float64
	SetContentsHeadroom(value float64)
	ContentsRect() corefoundation.CGRect
	SetContentsRect(value corefoundation.CGRect)
	ContentsScale() float64
	SetContentsScale(value float64)
	CornerCurve() LayerCornerCurve /* typedef */
	SetCornerCurve(value LayerCornerCurve /* typedef */)
	CornerRadius() float64
	SetCornerRadius(value float64)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	DrawsAsynchronously() bool
	SetDrawsAsynchronously(value bool)
	EdgeAntialiasingMask() EdgeAntialiasingMask
	SetEdgeAntialiasingMask(value EdgeAntialiasingMask)
	Filters() objc.IObject /* cross-framework: NSArray */
	SetFilters(value objc.IObject /* cross-framework: NSArray */)
	Frame() corefoundation.CGRect
	SetFrame(value corefoundation.CGRect)
	DoubleSided() bool
	SetDoubleSided(value bool)
	GeometryFlipped() bool
	SetGeometryFlipped(value bool)
	Hidden() bool
	SetHidden(value bool)
	Opaque() bool
	SetOpaque(value bool)
	LayoutManager() unsafe.Pointer
	SetLayoutManager(value unsafe.Pointer)
	MagnificationFilter() LayerContentsFilter /* typedef */
	SetMagnificationFilter(value LayerContentsFilter /* typedef */)
	Mask() ILayer
	SetMask(value ILayer)
	MaskedCorners() CornerMask
	SetMaskedCorners(value CornerMask)
	MasksToBounds() bool
	SetMasksToBounds(value bool)
	MinificationFilter() LayerContentsFilter /* typedef */
	SetMinificationFilter(value LayerContentsFilter /* typedef */)
	MinificationFilterBias() float32
	SetMinificationFilterBias(value float32)
	Name() objc.IObject /* cross-framework: NSString */
	SetName(value objc.IObject /* cross-framework: NSString */)
	NeedsDisplayOnBoundsChange() bool
	SetNeedsDisplayOnBoundsChange(value bool)
	Opacity() float32
	SetOpacity(value float32)
	Position() corefoundation.CGPoint
	SetPosition(value corefoundation.CGPoint)
	PreferredDynamicRange() DynamicRange /* typedef */
	SetPreferredDynamicRange(value DynamicRange /* typedef */)
	RasterizationScale() float64
	SetRasterizationScale(value float64)
	ShadowColor() ColorRef /* not a class type */
	SetShadowColor(value ColorRef /* not a class type */)
	ShadowOffset() corefoundation.CGSize
	SetShadowOffset(value corefoundation.CGSize)
	ShadowOpacity() float32
	SetShadowOpacity(value float32)
	ShadowPath() PathRef /* not a class type */
	SetShadowPath(value PathRef /* not a class type */)
	ShadowRadius() float64
	SetShadowRadius(value float64)
	ShouldRasterize() bool
	SetShouldRasterize(value bool)
	Style() objc.IObject /* cross-framework: NSDictionary */
	SetStyle(value objc.IObject /* cross-framework: NSDictionary */)
	Sublayers() []Layer
	SetSublayers(value []Layer)
	SublayerTransform() objc.IObject /* cross-framework: CATransform3D */
	SetSublayerTransform(value objc.IObject /* cross-framework: CATransform3D */)
	Superlayer() ILayer
	ToneMapMode() ToneMapMode /* typedef */
	SetToneMapMode(value ToneMapMode /* typedef */)
	Transform() objc.IObject /* cross-framework: CATransform3D */
	SetTransform(value objc.IObject /* cross-framework: CATransform3D */)
	VisibleRect() corefoundation.CGRect
	WantsExtendedDynamicRangeContent() bool
	SetWantsExtendedDynamicRangeContent(value bool)
	ZPosition() float64
	SetZPosition(value float64)
	IsDoubleSided() bool
	SetIsDoubleSided(value bool)
	IsGeometryFlipped() bool
	SetIsGeometryFlipped(value bool)
	IsHidden() bool
	SetIsHidden(value bool)
	IsOpaque() bool
	SetIsOpaque(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Layer */
	// methods:
	ActionForKey(event objc.IObject /* cross-framework: NSString */) unsafe.Pointer
	AddAnimationForKey(anim IAnimation, key objc.IObject /* cross-framework: NSString */)
	AddConstraint(c IConstraint)
	AddSublayer(layer ILayer)
	AffineTransform() corefoundation.CGAffineTransform
	AnimationForKey(key objc.IObject /* cross-framework: NSString */) IAnimation
	AnimationKeys() []string
	ContainsPoint(p corefoundation.CGPoint) bool
	ContentsAreFlipped() bool
	ConvertRectFromLayer(r corefoundation.CGRect, l ILayer) corefoundation.CGRect
	ConvertPointFromLayer(p corefoundation.CGPoint, l ILayer) corefoundation.CGPoint
	ConvertPointToLayer(p corefoundation.CGPoint, l ILayer) corefoundation.CGPoint
	ConvertRectToLayer(r corefoundation.CGRect, l ILayer) corefoundation.CGRect
	ConvertTimeFromLayer(t float64, l ILayer) float64
	ConvertTimeToLayer(t float64, l ILayer) float64
	Display()
	DisplayIfNeeded()
	DrawInContext(ctx ContextRef /* not a class type */)
	HitTest(p corefoundation.CGPoint) ILayer
	InsertSublayerAbove(layer ILayer, sibling ILayer)
	InsertSublayerAtIndex(layer ILayer, idx objectivec.IObject)
	InsertSublayerBelow(layer ILayer, sibling ILayer)
	LayoutIfNeeded()
	LayoutSublayers()
	ModelLayer() objectivec.IObject
	NeedsDisplay() bool
	NeedsLayout() bool
	PreferredFrameSize() corefoundation.CGSize
	PresentationLayer() objectivec.IObject
	RemoveAllAnimations()
	RemoveAnimationForKey(key objc.IObject /* cross-framework: NSString */)
	RemoveFromSuperlayer()
	RenderInContext(ctx ContextRef /* not a class type */)
	ReplaceSublayerWith(oldLayer ILayer, newLayer ILayer)
	ResizeWithOldSuperlayerSize(size corefoundation.CGSize)
	ResizeSublayersWithOldSize(size corefoundation.CGSize)
	ScrollPoint(p corefoundation.CGPoint)
	ScrollRectToVisible(r corefoundation.CGRect)
	SetAffineTransform(m corefoundation.CGAffineTransform)
	SetNeedsDisplay()
	SetNeedsDisplayInRect(r corefoundation.CGRect)
	SetNeedsLayout()
	ShouldArchiveValueForKey(key objc.IObject /* cross-framework: NSString */) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Layer */
// Alloc allocates a new instance without initialization.
func (lc _LayerClass) Alloc() Layer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Layer */
// An object that manages image-based content and allows you to perform animations on that content.
//
// Layers are often used to provide the backing store for views but can also be used without a view to display content. A layer’s main job is to manage the visual content that you provide but the layer itself has visual attributes that can be set, such as a background color, border, and shadow. In addition to managing visual content, the layer also maintains information about the geometry of its content (such as its position, size, and transform) that is used to present that content onscreen. Modifying the properties of the layer is how you initiate animations on the layer’s content or geometry. A layer object encapsulates the duration and pacing of a layer and its animations by adopting the protocol, which defines the layer’s timing information. If the layer object was created by a view, the view typically assigns itself as the layer’s delegate automatically, and you should not change that relationship. For layers you create yourself, you can assign a object and use that object to provide the contents of the layer dynamically and perform other tasks. A layer may also have a layout manager object (assigned to the property) to manage the layout of subviews separately.


// An object that manages image-based content and allows you to perform animations on that content.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Layer */

// Override to copy or initialize custom fields of the specified layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/init(layer:)
func NewLayerWithLayer(layer objc.IObject) Layer {
	instance := getLayerClass().Alloc()
	rv := objc.Send[Layer](instance.ID, objc.Sel("initWithLayer:"), layer)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewLayerWithLayer */


// Initializes a layer with a remote client ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/init(remoteClientId:)
func NewLayerWithRemoteClientId(client_id uint32 /* not a class type */) Layer {
	rv := objc.Send[Layer](objc.ID(getLayerClass().class), objc.Sel("layerWithRemoteClientId:"), client_id)
	return rv
}/* debug [class_init_methods/constructor]: NewLayerWithRemoteClientId */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Layer */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerCurveExpansionFactor(_:)
func (lc _LayerClass) CornerCurveExpansionFactor(curve LayerCornerCurve /* typedef */) float64 {
	rv := objc.Send[float64](objc.ID(lc.class), objc.Sel("cornerCurveExpansionFactor:"), curve)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CornerCurveExpansionFactor) */


// Returns the default action for the current class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/defaultAction(forKey:)
func (lc _LayerClass) DefaultActionForKey(event objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(lc.class), objc.Sel("defaultActionForKey:"), event)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultActionForKey) */


// Specifies the default value associated with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/defaultValue(forKey:)
func (lc _LayerClass) DefaultValueForKey(key objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(lc.class), objc.Sel("defaultValueForKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultValueForKey) */


// Initializes a layer with a remote client ID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/init(remoteClientId:)
func (lc _LayerClass) LayerWithRemoteClientId(client_id uint32 /* not a class type */) ILayer {
	rv := objc.Send[Layer](objc.ID(lc.class), objc.Sel("layerWithRemoteClientId:"), client_id)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithRemoteClientId) */


// Creates and returns an instance of the layer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layer
func (lc _LayerClass) Layer() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(lc.class), objc.Sel("layer"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Layer) */


// Returns a Boolean indicating whether changes to the specified key require the layer to be redisplayed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplay(forKey:)
func (lc _LayerClass) NeedsDisplayForKey(key objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](objc.ID(lc.class), objc.Sel("needsDisplayForKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=NeedsDisplayForKey) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Layer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Layer */

// Returns the action object assigned to the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/action(forKey:)
func (l_ Layer) ActionForKey(event objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("actionForKey:"), event)
	return rv
}/* debug [instance_methods/method]: ActionForKey */


// Add the specified animation object to the layer’s render tree.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/add(_:forKey:)
func (l_ Layer) AddAnimationForKey(anim IAnimation, key objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addAnimation:forKey:"), anim, key)
}/* debug [instance_methods/method]: AddAnimationForKey */


// Adds the specified constraint to the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/addConstraint(_:)
func (l_ Layer) AddConstraint(c IConstraint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addConstraint:"), c)
}/* debug [instance_methods/method]: AddConstraint */


// Appends the layer to the layer’s list of sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/addSublayer(_:)
func (l_ Layer) AddSublayer(layer ILayer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("addSublayer:"), layer)
}/* debug [instance_methods/method]: AddSublayer */


// Returns an affine version of the layer’s transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/affineTransform()
func (l_ Layer) AffineTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](l_.ID, objc.Sel("affineTransform"))
	return rv
}/* debug [instance_methods/method]: AffineTransform */


// Returns the animation object with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/animation(forKey:)
func (l_ Layer) AnimationForKey(key objc.IObject /* cross-framework: NSString */) IAnimation {
	rv := objc.Send[Animation](l_.ID, objc.Sel("animationForKey:"), key)
	return rv
}/* debug [instance_methods/method]: AnimationForKey */


// Returns an array of strings that identify the animations currently attached to the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/animationKeys()
func (l_ Layer) AnimationKeys() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("animationKeys"))
	return rv
}/* debug [instance_methods/method]: AnimationKeys */


// Returns whether the receiver contains a specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contains(_:)
func (l_ Layer) ContainsPoint(p corefoundation.CGPoint) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("containsPoint:"), p)
	return rv
}/* debug [instance_methods/method]: ContainsPoint */


// Returns a Boolean indicating whether the layer content is implicitly flipped when rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsAreFlipped()
func (l_ Layer) ContentsAreFlipped() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("contentsAreFlipped"))
	return rv
}/* debug [instance_methods/method]: ContentsAreFlipped */


// Converts the rectangle from the specified layer’s coordinate system to the receiver’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:from:)-4kx9l
func (l_ Layer) ConvertRectFromLayer(r corefoundation.CGRect, l ILayer) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l_.ID, objc.Sel("convertRect:fromLayer:"), r, l)
	return rv
}/* debug [instance_methods/method]: ConvertRectFromLayer */


// Converts the point from the specified layer’s coordinate system to the receiver’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:from:)-8kl76
func (l_ Layer) ConvertPointFromLayer(p corefoundation.CGPoint, l ILayer) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](l_.ID, objc.Sel("convertPoint:fromLayer:"), p, l)
	return rv
}/* debug [instance_methods/method]: ConvertPointFromLayer */


// Converts the point from the receiver’s coordinate system to the specified layer’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:to:)-7dcke
func (l_ Layer) ConvertPointToLayer(p corefoundation.CGPoint, l ILayer) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](l_.ID, objc.Sel("convertPoint:toLayer:"), p, l)
	return rv
}/* debug [instance_methods/method]: ConvertPointToLayer */


// Converts the rectangle from the receiver’s coordinate system to the specified layer’s coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convert(_:to:)-tly5
func (l_ Layer) ConvertRectToLayer(r corefoundation.CGRect, l ILayer) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l_.ID, objc.Sel("convertRect:toLayer:"), r, l)
	return rv
}/* debug [instance_methods/method]: ConvertRectToLayer */


// Converts the time interval from the specified layer’s time space to the receiver’s time space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convertTime(_:from:)
func (l_ Layer) ConvertTimeFromLayer(t float64, l ILayer) float64 {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("convertTime:fromLayer:"), t, l)
	return rv
}/* debug [instance_methods/method]: ConvertTimeFromLayer */


// Converts the time interval from the receiver’s time space to the specified layer’s time space
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/convertTime(_:to:)
func (l_ Layer) ConvertTimeToLayer(t float64, l ILayer) float64 {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("convertTime:toLayer:"), t, l)
	return rv
}/* debug [instance_methods/method]: ConvertTimeToLayer */


// Reloads the content of this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/display()
func (l_ Layer) Display() {
	objc.Send[objc.ID](l_.ID, objc.Sel("display"))
}/* debug [instance_methods/method]: Display */


// Initiates the update process for a layer if it is currently marked as needing an update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/displayIfNeeded()
func (l_ Layer) DisplayIfNeeded() {
	objc.Send[objc.ID](l_.ID, objc.Sel("displayIfNeeded"))
}/* debug [instance_methods/method]: DisplayIfNeeded */


// Draws the layer’s content using the specified graphics context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/draw(in:)
func (l_ Layer) DrawInContext(ctx ContextRef /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("drawInContext:"), ctx)
}/* debug [instance_methods/method]: DrawInContext */


// Returns the farthest descendant of the receiver in the layer hierarchy (including itself) that contains the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/hitTest(_:)
func (l_ Layer) HitTest(p corefoundation.CGPoint) ILayer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("hitTest:"), p)
	return rv
}/* debug [instance_methods/method]: HitTest */


// Inserts the specified sublayer above a different sublayer that already belongs to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/insertSublayer(_:above:)
func (l_ Layer) InsertSublayerAbove(layer ILayer, sibling ILayer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("insertSublayer:above:"), layer, sibling)
}/* debug [instance_methods/method]: InsertSublayerAbove */


// Inserts the specified layer into the receiver’s list of sublayers at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/insertSublayer(_:at:)
func (l_ Layer) InsertSublayerAtIndex(layer ILayer, idx objectivec.IObject) {
	objc.Send[objc.ID](l_.ID, objc.Sel("insertSublayer:atIndex:"), layer, idx)
}/* debug [instance_methods/method]: InsertSublayerAtIndex */


// Inserts the specified sublayer below a different sublayer that already belongs to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/insertSublayer(_:below:)
func (l_ Layer) InsertSublayerBelow(layer ILayer, sibling ILayer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("insertSublayer:below:"), layer, sibling)
}/* debug [instance_methods/method]: InsertSublayerBelow */


// Recalculate the receiver’s layout, if required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutIfNeeded()
func (l_ Layer) LayoutIfNeeded() {
	objc.Send[objc.ID](l_.ID, objc.Sel("layoutIfNeeded"))
}/* debug [instance_methods/method]: LayoutIfNeeded */


// Tells the layer to update its layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutSublayers()
func (l_ Layer) LayoutSublayers() {
	objc.Send[objc.ID](l_.ID, objc.Sel("layoutSublayers"))
}/* debug [instance_methods/method]: LayoutSublayers */


// Returns the model layer object associated with the receiver, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/model()
func (l_ Layer) ModelLayer() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("modelLayer"))
	return rv
}/* debug [instance_methods/method]: ModelLayer */


// Returns a Boolean indicating whether the layer has been marked as needing an update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplay()
func (l_ Layer) NeedsDisplay() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("needsDisplay"))
	return rv
}/* debug [instance_methods/method]: NeedsDisplay */


// Returns a Boolean indicating whether the layer has been marked as needing a layout update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsLayout()
func (l_ Layer) NeedsLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("needsLayout"))
	return rv
}/* debug [instance_methods/method]: NeedsLayout */


// Returns the preferred size of the layer in the coordinate space of its superlayer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/preferredFrameSize()
func (l_ Layer) PreferredFrameSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](l_.ID, objc.Sel("preferredFrameSize"))
	return rv
}/* debug [instance_methods/method]: PreferredFrameSize */


// Returns a copy of the presentation layer object that represents the state of the layer as it currently appears onscreen.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/presentation()
func (l_ Layer) PresentationLayer() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](l_.ID, objc.Sel("presentationLayer"))
	return rv
}/* debug [instance_methods/method]: PresentationLayer */


// Remove all animations attached to the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/removeAllAnimations()
func (l_ Layer) RemoveAllAnimations() {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeAllAnimations"))
}/* debug [instance_methods/method]: RemoveAllAnimations */


// Remove the animation object with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/removeAnimation(forKey:)
func (l_ Layer) RemoveAnimationForKey(key objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeAnimationForKey:"), key)
}/* debug [instance_methods/method]: RemoveAnimationForKey */


// Detaches the layer from its parent layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/removeFromSuperlayer()
func (l_ Layer) RemoveFromSuperlayer() {
	objc.Send[objc.ID](l_.ID, objc.Sel("removeFromSuperlayer"))
}/* debug [instance_methods/method]: RemoveFromSuperlayer */


// Renders the layer and its sublayers into the specified context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/render(in:)
func (l_ Layer) RenderInContext(ctx ContextRef /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("renderInContext:"), ctx)
}/* debug [instance_methods/method]: RenderInContext */


// Replaces the specified sublayer with a different layer object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/replaceSublayer(_:with:)
func (l_ Layer) ReplaceSublayerWith(oldLayer ILayer, newLayer ILayer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("replaceSublayer:with:"), oldLayer, newLayer)
}/* debug [instance_methods/method]: ReplaceSublayerWith */


// Informs the receiver that the size of its superlayer changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/resize(withOldSuperlayerSize:)
func (l_ Layer) ResizeWithOldSuperlayerSize(size corefoundation.CGSize) {
	objc.Send[objc.ID](l_.ID, objc.Sel("resizeWithOldSuperlayerSize:"), size)
}/* debug [instance_methods/method]: ResizeWithOldSuperlayerSize */


// Informs the receiver’s sublayers that the receiver’s size has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/resizeSublayers(withOldSize:)
func (l_ Layer) ResizeSublayersWithOldSize(size corefoundation.CGSize) {
	objc.Send[objc.ID](l_.ID, objc.Sel("resizeSublayersWithOldSize:"), size)
}/* debug [instance_methods/method]: ResizeSublayersWithOldSize */


// Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified point lies at the origin of the scroll layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/scroll(_:)
func (l_ Layer) ScrollPoint(p corefoundation.CGPoint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("scrollPoint:"), p)
}/* debug [instance_methods/method]: ScrollPoint */


// Initiates a scroll in the layer’s closest ancestor scroll layer so that the specified rectangle becomes visible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/scrollRectToVisible(_:)
func (l_ Layer) ScrollRectToVisible(r corefoundation.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("scrollRectToVisible:"), r)
}/* debug [instance_methods/method]: ScrollRectToVisible */


// Sets the layer’s transform to the specified affine transform.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setAffineTransform(_:)
func (l_ Layer) SetAffineTransform(m corefoundation.CGAffineTransform) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAffineTransform:"), m)
}/* debug [instance_methods/method]: SetAffineTransform */


// Marks the layer’s contents as needing to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsDisplay()
func (l_ Layer) SetNeedsDisplay() {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsDisplay"))
}/* debug [instance_methods/method]: SetNeedsDisplay */


// Marks the region within the specified rectangle as needing to be updated.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsDisplay(_:)
func (l_ Layer) SetNeedsDisplayInRect(r corefoundation.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsDisplayInRect:"), r)
}/* debug [instance_methods/method]: SetNeedsDisplayInRect */


// Invalidates the layer’s layout and marks it as needing an update.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/setNeedsLayout()
func (l_ Layer) SetNeedsLayout() {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsLayout"))
}/* debug [instance_methods/method]: SetNeedsLayout */


// Returns a Boolean indicating whether the value of the specified key should be archived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shouldArchiveValue(forKey:)
func (l_ Layer) ShouldArchiveValueForKey(key objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("shouldArchiveValueForKey:"), key)
	return rv
}/* debug [instance_methods/method]: ShouldArchiveValueForKey */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Layer */

// A dictionary containing layer actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/actions
func (l_ Layer) Actions() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](l_.ID, objc.Sel("actions"))
	return rv
}/* debug [instance_properties/getter]: actions */


// A dictionary containing layer actions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/actions
func (l_ Layer) SetActions(value foundation.IDictionary) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setActions:"), value)
}/* debug [instance_properties/setter]: actions */


// A Boolean indicating whether the layer is allowed to perform edge antialiasing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsEdgeAntialiasing
func (l_ Layer) AllowsEdgeAntialiasing() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("allowsEdgeAntialiasing"))
	return rv
}/* debug [instance_properties/getter]: allowsEdgeAntialiasing */


// A Boolean indicating whether the layer is allowed to perform edge antialiasing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsEdgeAntialiasing
func (l_ Layer) SetAllowsEdgeAntialiasing(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAllowsEdgeAntialiasing:"), value)
}/* debug [instance_properties/setter]: allowsEdgeAntialiasing */


// A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsGroupOpacity
func (l_ Layer) AllowsGroupOpacity() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("allowsGroupOpacity"))
	return rv
}/* debug [instance_properties/getter]: allowsGroupOpacity */


// A Boolean indicating whether the layer is allowed to composite itself as a group separate from its parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/allowsGroupOpacity
func (l_ Layer) SetAllowsGroupOpacity(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAllowsGroupOpacity:"), value)
}/* debug [instance_properties/setter]: allowsGroupOpacity */


// Defines the anchor point of the layer’s bounds rectangle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPoint
func (l_ Layer) AnchorPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](l_.ID, objc.Sel("anchorPoint"))
	return rv
}/* debug [instance_properties/getter]: anchorPoint */


// Defines the anchor point of the layer’s bounds rectangle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPoint
func (l_ Layer) SetAnchorPoint(value corefoundation.CGPoint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAnchorPoint:"), value)
}/* debug [instance_properties/setter]: anchorPoint */


// The anchor point for the layer’s position along the z axis. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPointZ
func (l_ Layer) AnchorPointZ() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("anchorPointZ"))
	return rv
}/* debug [instance_properties/getter]: anchorPointZ */


// The anchor point for the layer’s position along the z axis. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/anchorPointZ
func (l_ Layer) SetAnchorPointZ(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAnchorPointZ:"), value)
}/* debug [instance_properties/setter]: anchorPointZ */


// A bitmask defining how the layer is resized when the bounds of its superlayer changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/autoresizingMask
func (l_ Layer) AutoresizingMask() AutoresizingMask {
	rv := objc.Send[AutoresizingMask](l_.ID, objc.Sel("autoresizingMask"))
	return rv
}/* debug [instance_properties/getter]: autoresizingMask */


// A bitmask defining how the layer is resized when the bounds of its superlayer changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/autoresizingMask
func (l_ Layer) SetAutoresizingMask(value AutoresizingMask) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAutoresizingMask:"), value)
}/* debug [instance_properties/setter]: autoresizingMask */


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundColor
func (l_ Layer) BackgroundColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](l_.ID, objc.Sel("backgroundColor"))
	return rv
}/* debug [instance_properties/getter]: backgroundColor */


// The background color of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundColor
func (l_ Layer) SetBackgroundColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBackgroundColor:"), value)
}/* debug [instance_properties/setter]: backgroundColor */


// An array of Core Image filters to apply to the content immediately behind the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundFilters
func (l_ Layer) BackgroundFilters() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](l_.ID, objc.Sel("backgroundFilters"))
	return rv
}/* debug [instance_properties/getter]: backgroundFilters */


// An array of Core Image filters to apply to the content immediately behind the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/backgroundFilters
func (l_ Layer) SetBackgroundFilters(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBackgroundFilters:"), value)
}/* debug [instance_properties/setter]: backgroundFilters */


// The color of the layer’s border. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/borderColor
func (l_ Layer) BorderColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](l_.ID, objc.Sel("borderColor"))
	return rv
}/* debug [instance_properties/getter]: borderColor */


// The color of the layer’s border. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/borderColor
func (l_ Layer) SetBorderColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBorderColor:"), value)
}/* debug [instance_properties/setter]: borderColor */


// The width of the layer’s border. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/borderWidth
func (l_ Layer) BorderWidth() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("borderWidth"))
	return rv
}/* debug [instance_properties/getter]: borderWidth */


// The width of the layer’s border. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/borderWidth
func (l_ Layer) SetBorderWidth(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBorderWidth:"), value)
}/* debug [instance_properties/setter]: borderWidth */


// The layer’s bounds rectangle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/bounds
func (l_ Layer) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l_.ID, objc.Sel("bounds"))
	return rv
}/* debug [instance_properties/getter]: bounds */


// The layer’s bounds rectangle. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/bounds
func (l_ Layer) SetBounds(value corefoundation.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBounds:"), value)
}/* debug [instance_properties/setter]: bounds */


// A CoreImage filter used to composite the layer and the content behind it. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/compositingFilter
func (l_ Layer) CompositingFilter() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("compositingFilter"))
	return rv
}/* debug [instance_properties/getter]: compositingFilter */


// A CoreImage filter used to composite the layer and the content behind it. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/compositingFilter
func (l_ Layer) SetCompositingFilter(value objc.ID) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCompositingFilter:"), value)
}/* debug [instance_properties/setter]: compositingFilter */


// The constraints used to position current layer’s sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/constraints
func (l_ Layer) Constraints() []Constraint {
	rv := objc.Send[[]Constraint](l_.ID, objc.Sel("constraints"))
	return rv
}/* debug [instance_properties/getter]: constraints */


// The constraints used to position current layer’s sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/constraints
func (l_ Layer) SetConstraints(value []Constraint) {
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
}/* debug [instance_properties/setter]: constraints */


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (l_ Layer) Contents() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (l_ Layer) SetContents(value objc.ID) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */


// The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsCenter
func (l_ Layer) ContentsCenter() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l_.ID, objc.Sel("contentsCenter"))
	return rv
}/* debug [instance_properties/getter]: contentsCenter */


// The rectangle that defines how the layer contents are scaled if the layer’s contents are resized. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsCenter
func (l_ Layer) SetContentsCenter(value corefoundation.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsCenter:"), value)
}/* debug [instance_properties/setter]: contentsCenter */


// A hint for the desired storage format of the layer contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsFormat
func (l_ Layer) ContentsFormat() LayerContentsFormat /* typedef */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("contentsFormat"))
	return rv
}/* debug [instance_properties/getter]: contentsFormat */


// A hint for the desired storage format of the layer contents.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsFormat
func (l_ Layer) SetContentsFormat(value LayerContentsFormat /* typedef */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsFormat:"), value)
}/* debug [instance_properties/setter]: contentsFormat */


// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (l_ Layer) ContentsGravity() LayerContentsGravity /* typedef */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("contentsGravity"))
	return rv
}/* debug [instance_properties/getter]: contentsGravity */


// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (l_ Layer) SetContentsGravity(value LayerContentsGravity /* typedef */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsGravity:"), value)
}/* debug [instance_properties/setter]: contentsGravity */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsHeadroom
func (l_ Layer) ContentsHeadroom() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("contentsHeadroom"))
	return rv
}/* debug [instance_properties/getter]: contentsHeadroom */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsHeadroom
func (l_ Layer) SetContentsHeadroom(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsHeadroom:"), value)
}/* debug [instance_properties/setter]: contentsHeadroom */


// The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsRect
func (l_ Layer) ContentsRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l_.ID, objc.Sel("contentsRect"))
	return rv
}/* debug [instance_properties/getter]: contentsRect */


// The rectangle, in the unit coordinate space, that defines the portion of the layer’s contents that should be used. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsRect
func (l_ Layer) SetContentsRect(value corefoundation.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsRect:"), value)
}/* debug [instance_properties/setter]: contentsRect */


// The scale factor applied to the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
func (l_ Layer) ContentsScale() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("contentsScale"))
	return rv
}/* debug [instance_properties/getter]: contentsScale */


// The scale factor applied to the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsScale
func (l_ Layer) SetContentsScale(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setContentsScale:"), value)
}/* debug [instance_properties/setter]: contentsScale */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerCurve
func (l_ Layer) CornerCurve() LayerCornerCurve /* typedef */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("cornerCurve"))
	return rv
}/* debug [instance_properties/getter]: cornerCurve */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerCurve
func (l_ Layer) SetCornerCurve(value LayerCornerCurve /* typedef */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCornerCurve:"), value)
}/* debug [instance_properties/setter]: cornerCurve */


// The radius to use when drawing rounded corners for the layer’s background. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerRadius
func (l_ Layer) CornerRadius() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("cornerRadius"))
	return rv
}/* debug [instance_properties/getter]: cornerRadius */


// The radius to use when drawing rounded corners for the layer’s background. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/cornerRadius
func (l_ Layer) SetCornerRadius(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCornerRadius:"), value)
}/* debug [instance_properties/setter]: cornerRadius */


// The layer’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/delegate
func (l_ Layer) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The layer’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/delegate
func (l_ Layer) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/drawsAsynchronously
func (l_ Layer) DrawsAsynchronously() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("drawsAsynchronously"))
	return rv
}/* debug [instance_properties/getter]: drawsAsynchronously */


// A Boolean indicating whether drawing commands are deferred and processed asynchronously in a background thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/drawsAsynchronously
func (l_ Layer) SetDrawsAsynchronously(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDrawsAsynchronously:"), value)
}/* debug [instance_properties/setter]: drawsAsynchronously */


// A bitmask defining how the edges of the receiver are rasterized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/edgeAntialiasingMask
func (l_ Layer) EdgeAntialiasingMask() EdgeAntialiasingMask {
	rv := objc.Send[EdgeAntialiasingMask](l_.ID, objc.Sel("edgeAntialiasingMask"))
	return rv
}/* debug [instance_properties/getter]: edgeAntialiasingMask */


// A bitmask defining how the edges of the receiver are rasterized.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/edgeAntialiasingMask
func (l_ Layer) SetEdgeAntialiasingMask(value EdgeAntialiasingMask) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setEdgeAntialiasingMask:"), value)
}/* debug [instance_properties/setter]: edgeAntialiasingMask */


// An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/filters
func (l_ Layer) Filters() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](l_.ID, objc.Sel("filters"))
	return rv
}/* debug [instance_properties/getter]: filters */


// An array of Core Image filters to apply to the contents of the layer and its sublayers. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/filters
func (l_ Layer) SetFilters(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setFilters:"), value)
}/* debug [instance_properties/setter]: filters */


// The layer’s frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/frame
func (l_ Layer) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l_.ID, objc.Sel("frame"))
	return rv
}/* debug [instance_properties/getter]: frame */


// The layer’s frame rectangle.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/frame
func (l_ Layer) SetFrame(value corefoundation.CGRect) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setFrame:"), value)
}/* debug [instance_properties/setter]: frame */


// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isDoubleSided
func (l_ Layer) DoubleSided() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("doubleSided"))
	return rv
}/* debug [instance_properties/getter]: doubleSided */


// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isDoubleSided
func (l_ Layer) SetDoubleSided(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDoubleSided:"), value)
}/* debug [instance_properties/setter]: doubleSided */


// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isGeometryFlipped
func (l_ Layer) GeometryFlipped() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("geometryFlipped"))
	return rv
}/* debug [instance_properties/getter]: geometryFlipped */


// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isGeometryFlipped
func (l_ Layer) SetGeometryFlipped(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setGeometryFlipped:"), value)
}/* debug [instance_properties/setter]: geometryFlipped */


// A Boolean indicating whether the layer is displayed. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isHidden
func (l_ Layer) Hidden() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("hidden"))
	return rv
}/* debug [instance_properties/getter]: hidden */


// A Boolean indicating whether the layer is displayed. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isHidden
func (l_ Layer) SetHidden(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setHidden:"), value)
}/* debug [instance_properties/setter]: hidden */


// A Boolean value indicating whether the layer contains completely opaque content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isOpaque
func (l_ Layer) Opaque() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("opaque"))
	return rv
}/* debug [instance_properties/getter]: opaque */


// A Boolean value indicating whether the layer contains completely opaque content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/isOpaque
func (l_ Layer) SetOpaque(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOpaque:"), value)
}/* debug [instance_properties/setter]: opaque */


// The object responsible for laying out the layer’s sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutManager
func (l_ Layer) LayoutManager() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("layoutManager"))
	return rv
}/* debug [instance_properties/getter]: layoutManager */


// The object responsible for laying out the layer’s sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/layoutManager
func (l_ Layer) SetLayoutManager(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLayoutManager:"), value)
}/* debug [instance_properties/setter]: layoutManager */


// The filter used when increasing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/magnificationFilter
func (l_ Layer) MagnificationFilter() LayerContentsFilter /* typedef */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("magnificationFilter"))
	return rv
}/* debug [instance_properties/getter]: magnificationFilter */


// The filter used when increasing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/magnificationFilter
func (l_ Layer) SetMagnificationFilter(value LayerContentsFilter /* typedef */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMagnificationFilter:"), value)
}/* debug [instance_properties/setter]: magnificationFilter */


// An optional layer whose alpha channel is used to mask the layer’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/mask
func (l_ Layer) Mask() ILayer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("mask"))
	return rv
}/* debug [instance_properties/getter]: mask */


// An optional layer whose alpha channel is used to mask the layer’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/mask
func (l_ Layer) SetMask(value ILayer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMask:"), value)
}/* debug [instance_properties/setter]: mask */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/maskedCorners
func (l_ Layer) MaskedCorners() CornerMask {
	rv := objc.Send[CornerMask](l_.ID, objc.Sel("maskedCorners"))
	return rv
}/* debug [instance_properties/getter]: maskedCorners */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/maskedCorners
func (l_ Layer) SetMaskedCorners(value CornerMask) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMaskedCorners:"), value)
}/* debug [instance_properties/setter]: maskedCorners */


// A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/masksToBounds
func (l_ Layer) MasksToBounds() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("masksToBounds"))
	return rv
}/* debug [instance_properties/getter]: masksToBounds */


// A Boolean indicating whether sublayers are clipped to the layer’s bounds. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/masksToBounds
func (l_ Layer) SetMasksToBounds(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMasksToBounds:"), value)
}/* debug [instance_properties/setter]: masksToBounds */


// The filter used when reducing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilter
func (l_ Layer) MinificationFilter() LayerContentsFilter /* typedef */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("minificationFilter"))
	return rv
}/* debug [instance_properties/getter]: minificationFilter */


// The filter used when reducing the size of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilter
func (l_ Layer) SetMinificationFilter(value LayerContentsFilter /* typedef */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMinificationFilter:"), value)
}/* debug [instance_properties/setter]: minificationFilter */


// The bias factor used by the minification filter to determine the levels of detail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilterBias
func (l_ Layer) MinificationFilterBias() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("minificationFilterBias"))
	return rv
}/* debug [instance_properties/getter]: minificationFilterBias */


// The bias factor used by the minification filter to determine the levels of detail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/minificationFilterBias
func (l_ Layer) SetMinificationFilterBias(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setMinificationFilterBias:"), value)
}/* debug [instance_properties/setter]: minificationFilterBias */


// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/name
func (l_ Layer) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// The name of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/name
func (l_ Layer) SetName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setName:"), value)
}/* debug [instance_properties/setter]: name */


// A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplayOnBoundsChange
func (l_ Layer) NeedsDisplayOnBoundsChange() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("needsDisplayOnBoundsChange"))
	return rv
}/* debug [instance_properties/getter]: needsDisplayOnBoundsChange */


// A Boolean indicating whether the layer contents must be updated when its bounds rectangle changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/needsDisplayOnBoundsChange
func (l_ Layer) SetNeedsDisplayOnBoundsChange(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setNeedsDisplayOnBoundsChange:"), value)
}/* debug [instance_properties/setter]: needsDisplayOnBoundsChange */


// The opacity of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/opacity
func (l_ Layer) Opacity() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("opacity"))
	return rv
}/* debug [instance_properties/getter]: opacity */


// The opacity of the receiver. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/opacity
func (l_ Layer) SetOpacity(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setOpacity:"), value)
}/* debug [instance_properties/setter]: opacity */


// The layer’s position in its superlayer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/position
func (l_ Layer) Position() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](l_.ID, objc.Sel("position"))
	return rv
}/* debug [instance_properties/getter]: position */


// The layer’s position in its superlayer’s coordinate space. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/position
func (l_ Layer) SetPosition(value corefoundation.CGPoint) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPosition:"), value)
}/* debug [instance_properties/setter]: position */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/preferredDynamicRange
func (l_ Layer) PreferredDynamicRange() DynamicRange /* typedef */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("preferredDynamicRange"))
	return rv
}/* debug [instance_properties/getter]: preferredDynamicRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/preferredDynamicRange
func (l_ Layer) SetPreferredDynamicRange(value DynamicRange /* typedef */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setPreferredDynamicRange:"), value)
}/* debug [instance_properties/setter]: preferredDynamicRange */


// The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/rasterizationScale
func (l_ Layer) RasterizationScale() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("rasterizationScale"))
	return rv
}/* debug [instance_properties/getter]: rasterizationScale */


// The scale at which to rasterize content, relative to the coordinate space of the layer. Animatable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/rasterizationScale
func (l_ Layer) SetRasterizationScale(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setRasterizationScale:"), value)
}/* debug [instance_properties/setter]: rasterizationScale */


// The color of the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowColor
func (l_ Layer) ShadowColor() ColorRef /* not a class type */ {
	rv := objc.Send[ColorRef](l_.ID, objc.Sel("shadowColor"))
	return rv
}/* debug [instance_properties/getter]: shadowColor */


// The color of the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowColor
func (l_ Layer) SetShadowColor(value ColorRef /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowColor:"), value)
}/* debug [instance_properties/setter]: shadowColor */


// The offset (in points) of the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOffset
func (l_ Layer) ShadowOffset() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](l_.ID, objc.Sel("shadowOffset"))
	return rv
}/* debug [instance_properties/getter]: shadowOffset */


// The offset (in points) of the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOffset
func (l_ Layer) SetShadowOffset(value corefoundation.CGSize) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowOffset:"), value)
}/* debug [instance_properties/setter]: shadowOffset */


// The opacity of the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOpacity
func (l_ Layer) ShadowOpacity() float32 {
	rv := objc.Send[float32](l_.ID, objc.Sel("shadowOpacity"))
	return rv
}/* debug [instance_properties/getter]: shadowOpacity */


// The opacity of the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowOpacity
func (l_ Layer) SetShadowOpacity(value float32) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowOpacity:"), value)
}/* debug [instance_properties/setter]: shadowOpacity */


// The shape of the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowPath
func (l_ Layer) ShadowPath() PathRef /* not a class type */ {
	rv := objc.Send[PathRef](l_.ID, objc.Sel("shadowPath"))
	return rv
}/* debug [instance_properties/getter]: shadowPath */


// The shape of the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowPath
func (l_ Layer) SetShadowPath(value PathRef /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowPath:"), value)
}/* debug [instance_properties/setter]: shadowPath */


// The blur radius (in points) used to render the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowRadius
func (l_ Layer) ShadowRadius() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("shadowRadius"))
	return rv
}/* debug [instance_properties/getter]: shadowRadius */


// The blur radius (in points) used to render the layer’s shadow. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shadowRadius
func (l_ Layer) SetShadowRadius(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShadowRadius:"), value)
}/* debug [instance_properties/setter]: shadowRadius */


// A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shouldRasterize
func (l_ Layer) ShouldRasterize() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("shouldRasterize"))
	return rv
}/* debug [instance_properties/getter]: shouldRasterize */


// A Boolean that indicates whether the layer is rendered as a bitmap before compositing. Animatable
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/shouldRasterize
func (l_ Layer) SetShouldRasterize(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setShouldRasterize:"), value)
}/* debug [instance_properties/setter]: shouldRasterize */


// An optional dictionary used to store property values that aren’t explicitly defined by the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/style
func (l_ Layer) Style() objc.IObject /* cross-framework: NSDictionary */ {
	rv := objc.Send[foundation.NSDictionary](l_.ID, objc.Sel("style"))
	return rv
}/* debug [instance_properties/getter]: style */


// An optional dictionary used to store property values that aren’t explicitly defined by the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/style
func (l_ Layer) SetStyle(value objc.IObject /* cross-framework: NSDictionary */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setStyle:"), value)
}/* debug [instance_properties/setter]: style */


// An array containing the layer’s sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayers
func (l_ Layer) Sublayers() []Layer {
	rv := objc.Send[[]Layer](l_.ID, objc.Sel("sublayers"))
	return rv
}/* debug [instance_properties/getter]: sublayers */


// An array containing the layer’s sublayers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayers
func (l_ Layer) SetSublayers(value []Layer) {
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
}/* debug [instance_properties/setter]: sublayers */


// Specifies the transform to apply to sublayers when rendering. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayerTransform
func (l_ Layer) SublayerTransform() objc.IObject /* cross-framework: CATransform3D */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("sublayerTransform"))
	return rv
}/* debug [instance_properties/getter]: sublayerTransform */


// Specifies the transform to apply to sublayers when rendering. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/sublayerTransform
func (l_ Layer) SetSublayerTransform(value objc.IObject /* cross-framework: CATransform3D */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setSublayerTransform:"), value)
}/* debug [instance_properties/setter]: sublayerTransform */


// The superlayer of the layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/superlayer
func (l_ Layer) Superlayer() ILayer {
	rv := objc.Send[Layer](l_.ID, objc.Sel("superlayer"))
	return rv
}/* debug [instance_properties/getter]: superlayer */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/toneMapMode-swift.property
func (l_ Layer) ToneMapMode() ToneMapMode /* typedef */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("toneMapMode"))
	return rv
}/* debug [instance_properties/getter]: toneMapMode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/toneMapMode-swift.property
func (l_ Layer) SetToneMapMode(value ToneMapMode /* typedef */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setToneMapMode:"), value)
}/* debug [instance_properties/setter]: toneMapMode */


// The transform applied to the layer’s contents. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/transform
func (l_ Layer) Transform() objc.IObject /* cross-framework: CATransform3D */ {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("transform"))
	return rv
}/* debug [instance_properties/getter]: transform */


// The transform applied to the layer’s contents. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/transform
func (l_ Layer) SetTransform(value objc.IObject /* cross-framework: CATransform3D */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTransform:"), value)
}/* debug [instance_properties/setter]: transform */


// The visible region of the layer in its own coordinate space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/visibleRect
func (l_ Layer) VisibleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](l_.ID, objc.Sel("visibleRect"))
	return rv
}/* debug [instance_properties/getter]: visibleRect */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsExtendedDynamicRangeContent
func (l_ Layer) WantsExtendedDynamicRangeContent() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("wantsExtendedDynamicRangeContent"))
	return rv
}/* debug [instance_properties/getter]: wantsExtendedDynamicRangeContent */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/wantsExtendedDynamicRangeContent
func (l_ Layer) SetWantsExtendedDynamicRangeContent(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setWantsExtendedDynamicRangeContent:"), value)
}/* debug [instance_properties/setter]: wantsExtendedDynamicRangeContent */


// The layer’s position on the z axis. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/zPosition
func (l_ Layer) ZPosition() float64 {
	rv := objc.Send[float64](l_.ID, objc.Sel("zPosition"))
	return rv
}/* debug [instance_properties/getter]: zPosition */


// The layer’s position on the z axis. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CALayer/zPosition
func (l_ Layer) SetZPosition(value float64) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setZPosition:"), value)
}/* debug [instance_properties/setter]: zPosition */


// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isdoublesided
func (l_ Layer) IsDoubleSided() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isDoubleSided"))
	return rv
}/* debug [instance_properties/getter]: isDoubleSided */


// A Boolean indicating whether the layer displays its content when facing away from the viewer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isdoublesided
func (l_ Layer) SetIsDoubleSided(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsDoubleSided:"), value)
}/* debug [instance_properties/setter]: isDoubleSided */


// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isgeometryflipped
func (l_ Layer) IsGeometryFlipped() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isGeometryFlipped"))
	return rv
}/* debug [instance_properties/getter]: isGeometryFlipped */


// A Boolean that indicates whether the geometry of the layer and its sublayers is flipped vertically.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isgeometryflipped
func (l_ Layer) SetIsGeometryFlipped(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsGeometryFlipped:"), value)
}/* debug [instance_properties/setter]: isGeometryFlipped */


// A Boolean indicating whether the layer is displayed. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/ishidden
func (l_ Layer) IsHidden() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isHidden"))
	return rv
}/* debug [instance_properties/getter]: isHidden */


// A Boolean indicating whether the layer is displayed. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/ishidden
func (l_ Layer) SetIsHidden(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsHidden:"), value)
}/* debug [instance_properties/setter]: isHidden */


// A Boolean value indicating whether the layer contains completely opaque content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isopaque
func (l_ Layer) IsOpaque() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isOpaque"))
	return rv
}/* debug [instance_properties/getter]: isOpaque */


// A Boolean value indicating whether the layer contains completely opaque content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/isopaque
func (l_ Layer) SetIsOpaque(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsOpaque:"), value)
}/* debug [instance_properties/setter]: isOpaque */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CALayer */


