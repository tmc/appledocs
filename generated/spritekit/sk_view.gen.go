// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SKView] class.
var (
	sKViewClass     _SKViewClass
	sKViewClassOnce sync.Once
)

func getSKViewClass() _SKViewClass {
	sKViewClassOnce.Do(func() {
		sKViewClass = _SKViewClass{objc.GetClass("SKView")}
	})
	return sKViewClass
}

type _SKViewClass struct {
	class objc.Class
}

// An interface definition for the [SKView] class.
type ISKView interface {
	IView
	ConvertPointFromScene(point unsafe.Pointer, scene unsafe.Pointer) unsafe.Pointer
	ConvertPointToScene(point unsafe.Pointer, scene unsafe.Pointer) unsafe.Pointer
	PresentScene(scene unsafe.Pointer)
	PresentSceneTransition(scene unsafe.Pointer, transition unsafe.Pointer)
	TextureFromNode(node unsafe.Pointer) unsafe.Pointer
	TextureFromNodeCrop(node unsafe.Pointer, crop unsafe.Pointer) unsafe.Pointer
}

// A view subclass that renders a SpriteKit scene. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView
type SKView struct {
	View
}

// SKViewFrom constructs a [SKView] from an unsafe.Pointer.
//
// A view subclass that renders a SpriteKit scene.
func SKViewFrom(ptr unsafe.Pointer) SKView {
	return SKView{
		View: ViewFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SKViewClass) Alloc() SKView {
	rv := objc.Send[SKView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SKViewClass) New() SKView {
	rv := objc.Send[SKView](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SKView) Init() SKView {
	rv := objc.Send[SKView](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SKView) Autorelease() SKView {
	rv := objc.Send[SKView](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKView creates a new SKView instance.
func NewSKView() SKView {
	return getSKViewClass().New()
}


// Converts a point from scene coordinates to view coordinates. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/convert(_:from:)
func (s_ SKView) ConvertPointFromScene(point unsafe.Pointer, scene unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("convertPoint:fromScene:"), point, scene)
	return rv
}
// Converts a point from view coordinates to scene coordinates. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/convert(_:to:)
func (s_ SKView) ConvertPointToScene(point unsafe.Pointer, scene unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("convertPoint:toScene:"), point, scene)
	return rv
}
// Presents a scene. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/presentScene(_:)
func (s_ SKView) PresentScene(scene unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("presentScene:"), scene)
}
// Transitions from the current scene to a new scene. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/presentScene(_:transition:)
func (s_ SKView) PresentSceneTransition(scene unsafe.Pointer, transition unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("presentScene:transition:"), scene, transition)
}
// Renders the contents of a node tree and returns the rendered image as a texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/texture(from:)
func (s_ SKView) TextureFromNode(node unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("textureFromNode:"), node)
	return rv
}
// Renders a portion of a node’s contents and returns the rendered image as a texture. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/SpriteKit/SKView/texture(from:crop:)
func (s_ SKView) TextureFromNodeCrop(node unsafe.Pointer, crop unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("textureFromNode:crop:"), node, crop)
	return rv
}


