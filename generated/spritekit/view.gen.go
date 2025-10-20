// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [View] class.
var (
	viewClass     _ViewClass
	viewClassOnce sync.Once
)

func getViewClass() _ViewClass {
	viewClassOnce.Do(func() {
		viewClass = _ViewClass{objc.GetClass("SKView")}
	})
	return viewClass
}

type _ViewClass struct {
	class objc.Class
}

// An interface definition for the [View] class.
type IView interface {
	objectivec.IObject
	ConvertPointFromScene(point coregraphics.CGPoint, scene unsafe.Pointer) coregraphics.CGPoint
	ConvertPointToScene(point coregraphics.CGPoint, scene unsafe.Pointer) coregraphics.CGPoint
	PresentScene(scene unsafe.Pointer)
	PresentSceneTransition(scene unsafe.Pointer, transition unsafe.Pointer)
	TextureFromNode(node unsafe.Pointer) unsafe.Pointer
	TextureFromNodeCrop(node unsafe.Pointer, crop coregraphics.CGRect) unsafe.Pointer
}

// A view subclass that renders a SpriteKit scene.
//
// You present a scene by calling the view’s method. When a scene is presented by the view, it alternates between running its simulation (which animates the content) and rendering the content for display. You can pause the scene by setting the view’s property to .
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView
type View struct {
	objectivec.Object
}

// ViewFrom constructs a [View] from an unsafe.Pointer.
//
// A view subclass that renders a SpriteKit scene.
func ViewFrom(ptr unsafe.Pointer) View {
	return View{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (vc _ViewClass) Alloc() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (vc _ViewClass) New() View {
	rv := objc.Send[View](objc.ID(vc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (v_ View) Init() View {
	rv := objc.Send[View](v_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v_ View) Autorelease() View {
	rv := objc.Send[View](v_.ID, objc.Sel("autorelease"))
	return rv
}

// NewView creates a new View instance.
func NewView() View {
	return getViewClass().New()
}


// Converts a point from scene coordinates to view coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/convert(_:from:)
func (v_ View) ConvertPointFromScene(point coregraphics.CGPoint, scene unsafe.Pointer) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPoint:fromScene:"), point, scene)
	return rv
}

// Converts a point from view coordinates to scene coordinates.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/convert(_:to:)
func (v_ View) ConvertPointToScene(point coregraphics.CGPoint, scene unsafe.Pointer) coregraphics.CGPoint {
	rv := objc.Send[coregraphics.CGPoint](v_.ID, objc.Sel("convertPoint:toScene:"), point, scene)
	return rv
}

// Presents a scene.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/presentScene(_:)
func (v_ View) PresentScene(scene unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentScene:"), scene)
}

// Transitions from the current scene to a new scene.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/presentScene(_:transition:)
func (v_ View) PresentSceneTransition(scene unsafe.Pointer, transition unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("presentScene:transition:"), scene, transition)
}

// Renders the contents of a node tree and returns the rendered image as a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/texture(from:)
func (v_ View) TextureFromNode(node unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("textureFromNode:"), node)
	return rv
}

// Renders a portion of a node’s contents and returns the rendered image as a texture.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/texture(from:crop:)
func (v_ View) TextureFromNodeCrop(node unsafe.Pointer, crop coregraphics.CGRect) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("textureFromNode:crop:"), node, crop)
	return rv
}

// A Boolean property that indicates whether the view is rendered using transparency.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/allowsTransparency
func (v_ View) AllowsTransparency() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("allowsTransparency"))
	return rv
}


// SetAllowsTransparency sets the value of the allowsTransparency property.
// A Boolean property that indicates whether the view is rendered using transparency.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/allowsTransparency
func (v_ View) SetAllowsTransparency(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAllowsTransparency:"), value)
}
// A delegate that allows dynamic control of the view’s render rate.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/delegate
func (v_ View) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// A delegate that allows dynamic control of the view’s render rate.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/delegate
func (v_ View) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDelegate:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/disableDepthStencilBuffer
func (v_ View) DisableDepthStencilBuffer() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("disableDepthStencilBuffer"))
	return rv
}


// SetDisableDepthStencilBuffer sets the value of the disableDepthStencilBuffer property.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/disableDepthStencilBuffer
func (v_ View) SetDisableDepthStencilBuffer(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setDisableDepthStencilBuffer:"), value)
}
// The number of frames that must pass before the scene is called to update its contents.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/frameInterval
func (v_ View) FrameInterval() int {
	rv := objc.Send[int](v_.ID, objc.Sel("frameInterval"))
	return rv
}


// SetFrameInterval sets the value of the frameInterval property.
// The number of frames that must pass before the scene is called to update its contents.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/frameInterval
func (v_ View) SetFrameInterval(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setFrameInterval:"), value)
}
// A Boolean value that indicates whether parent-child and sibling relationships affect the rendering order of nodes in the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/ignoresSiblingOrder
func (v_ View) IgnoresSiblingOrder() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("ignoresSiblingOrder"))
	return rv
}


// SetIgnoresSiblingOrder sets the value of the ignoresSiblingOrder property.
// A Boolean value that indicates whether parent-child and sibling relationships affect the rendering order of nodes in the scene.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/ignoresSiblingOrder
func (v_ View) SetIgnoresSiblingOrder(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setIgnoresSiblingOrder:"), value)
}
// A Boolean value that indicates whether the content is rendered asynchronously.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/isAsynchronous
func (v_ View) Asynchronous() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("asynchronous"))
	return rv
}


// SetAsynchronous sets the value of the asynchronous property.
// A Boolean value that indicates whether the content is rendered asynchronously.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/isAsynchronous
func (v_ View) SetAsynchronous(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setAsynchronous:"), value)
}
// A Boolean value that indicates whether the view’s scene animations are paused.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/isPaused
func (v_ View) Paused() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("paused"))
	return rv
}


// SetPaused sets the value of the paused property.
// A Boolean value that indicates whether the view’s scene animations are paused.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/isPaused
func (v_ View) SetPaused(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPaused:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/preferredFrameRate
func (v_ View) PreferredFrameRate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("preferredFrameRate"))
	return rv
}


// SetPreferredFrameRate sets the value of the preferredFrameRate property.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/preferredFrameRate
func (v_ View) SetPreferredFrameRate(value unsafe.Pointer) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreferredFrameRate:"), value)
}
// The animation frame rate that the view uses to render its scene.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/preferredFramesPerSecond
func (v_ View) PreferredFramesPerSecond() int {
	rv := objc.Send[int](v_.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}


// SetPreferredFramesPerSecond sets the value of the preferredFramesPerSecond property.
// The animation frame rate that the view uses to render its scene.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/preferredFramesPerSecond
func (v_ View) SetPreferredFramesPerSecond(value int) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setPreferredFramesPerSecond:"), value)
}
// The scene currently presented by this view.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/scene
func (v_ View) Scene() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v_.ID, objc.Sel("scene"))
	return rv
}

// A Boolean value that indicates whether the view automatically culls non-visible nodes from the rendering tree.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/shouldCullNonVisibleNodes
func (v_ View) ShouldCullNonVisibleNodes() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("shouldCullNonVisibleNodes"))
	return rv
}


// SetShouldCullNonVisibleNodes sets the value of the shouldCullNonVisibleNodes property.
// A Boolean value that indicates whether the view automatically culls non-visible nodes from the rendering tree.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/shouldCullNonVisibleNodes
func (v_ View) SetShouldCullNonVisibleNodes(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShouldCullNonVisibleNodes:"), value)
}
// A Boolean value that indicates whether the view displays the number of drawing passes it needed to render the view.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsDrawCount
func (v_ View) ShowsDrawCount() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsDrawCount"))
	return rv
}


// SetShowsDrawCount sets the value of the showsDrawCount property.
// A Boolean value that indicates whether the view displays the number of drawing passes it needed to render the view.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsDrawCount
func (v_ View) SetShowsDrawCount(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShowsDrawCount:"), value)
}
// A Boolean value that indicates whether the view displays a frame rate indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsFPS
func (v_ View) ShowsFPS() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsFPS"))
	return rv
}


// SetShowsFPS sets the value of the showsFPS property.
// A Boolean value that indicates whether the view displays a frame rate indicator.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsFPS
func (v_ View) SetShowsFPS(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShowsFPS:"), value)
}
// A Boolean value that indicates whether the view displays information about physics fields in the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsFields
func (v_ View) ShowsFields() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsFields"))
	return rv
}


// SetShowsFields sets the value of the showsFields property.
// A Boolean value that indicates whether the view displays information about physics fields in the scene.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsFields
func (v_ View) SetShowsFields(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShowsFields:"), value)
}
// A Boolean value that indicates whether the view displays an overlay that shows physics bodies that are visible in the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsNodeCount
func (v_ View) ShowsNodeCount() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsNodeCount"))
	return rv
}


// SetShowsNodeCount sets the value of the showsNodeCount property.
// A Boolean value that indicates whether the view displays an overlay that shows physics bodies that are visible in the scene.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsNodeCount
func (v_ View) SetShowsNodeCount(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShowsNodeCount:"), value)
}
// A Boolean value that indicates whether the view displays physics-related debugging information.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsPhysics
func (v_ View) ShowsPhysics() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsPhysics"))
	return rv
}


// SetShowsPhysics sets the value of the showsPhysics property.
// A Boolean value that indicates whether the view displays physics-related debugging information.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsPhysics
func (v_ View) SetShowsPhysics(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShowsPhysics:"), value)
}
// A Boolean value that indicates whether the view displays the number of rectangles used to render the scene.
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsQuadCount
func (v_ View) ShowsQuadCount() bool {
	rv := objc.Send[bool](v_.ID, objc.Sel("showsQuadCount"))
	return rv
}


// SetShowsQuadCount sets the value of the showsQuadCount property.
// A Boolean value that indicates whether the view displays the number of rectangles used to render the scene.

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/showsQuadCount
func (v_ View) SetShowsQuadCount(value bool) {
	objc.Send[objc.ID](v_.ID, objc.Sel("setShowsQuadCount:"), value)
}


