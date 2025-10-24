// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
)

// The class instance for the [TiledLayer] class.
var (
	TiledLayerClass     _TiledLayerClass
	TiledLayerClassOnce sync.Once
)

func getTiledLayerClass() _TiledLayerClass {
	TiledLayerClassOnce.Do(func() {
		TiledLayerClass = _TiledLayerClass{objc.GetClass("CATiledLayer")}
	})
	return TiledLayerClass
}

type _TiledLayerClass struct {
	class objc.Class
}

// An interface definition for the [TiledLayer] class.
type ITiledLayer interface {
	ILayer
	// properties:
	TileSize() objc.IObject /* cross-framework: Size */
	SetTileSize(value objc.IObject /* cross-framework: Size */)
	Contents() unsafe.Pointer
	SetContents(value unsafe.Pointer)
	LevelsOfDetail() int
	SetLevelsOfDetail(value int)
	LevelsOfDetailBias() int
	SetLevelsOfDetailBias(value int)
	// methods:
}

// A layer that provides a way to asynchronously provide tiles of the layer’s content, potentially cached at multiple levels of detail.
//
// As more data is required by the renderer, the layer’s method is called on one or more background threads to supply the drawing operations to fill in one tile of data. The clip bounds and current transformation matrix (CTM) of the drawing context can be used to determine the bounds and resolution of the tile being requested. Regions of the layer may be invalidated using the method however the update will be asynchronous. While the next display update will most likely not contain the updated content, a future update will.


// A layer that provides a way to asynchronously provide tiles of the layer’s content, potentially cached at multiple levels of detail.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer
type TiledLayer struct {
	Layer
}

// TiledLayerFrom constructs a [TiledLayer] from an unsafe.Pointer.
//
// A layer that provides a way to asynchronously provide tiles of the layer’s content, potentially cached at multiple levels of detail.
func TiledLayerFrom(ptr unsafe.Pointer) TiledLayer {
	return TiledLayer{
		Layer: LayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TiledLayerClass) Alloc() TiledLayer {
	rv := objc.Send[TiledLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (tc _TiledLayerClass) New() TiledLayer {
	rv := objc.Send[TiledLayer](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TiledLayer) Init() TiledLayer {
	rv := objc.Send[TiledLayer](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TiledLayer) Autorelease() TiledLayer {
	rv := objc.Send[TiledLayer](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTiledLayer creates a new TiledLayer instance.
func NewTiledLayer() TiledLayer {
	return getTiledLayerClass().New()
}



// The time, in seconds, that newly added images take to “fade-in” to the rendered representation of the tiled layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/fadeDuration()
func (tc _TiledLayerClass) FadeDuration() float64 {
	rv := objc.Send[TimeInterval](objc.ID(tc.class), objc.Sel("fadeDuration"))
	return rv
}


// The maximum size of each tile used to create the layer’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/tileSize
func (t_ TiledLayer) TileSize() objc.IObject /* cross-framework: Size */ {
	rv := objc.Send[corefoundation.Size](t_.ID, objc.Sel("tileSize"))
	return rv
}


// The maximum size of each tile used to create the layer’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/tileSize
func (t_ TiledLayer) SetTileSize(value objc.IObject /* cross-framework: Size */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTileSize:"), value)
}


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/contents
func (t_ TiledLayer) Contents() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("contents"))
	return rv
}


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/contents
func (t_ TiledLayer) SetContents(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContents:"), value)
}


// The number of levels of detail maintained by this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/levelsofdetail
func (t_ TiledLayer) LevelsOfDetail() int {
	rv := objc.Send[int](t_.ID, objc.Sel("levelsOfDetail"))
	return rv
}


// The number of levels of detail maintained by this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/levelsofdetail
func (t_ TiledLayer) SetLevelsOfDetail(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLevelsOfDetail:"), value)
}


// The number of magnified levels of detail for this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/levelsofdetailbias
func (t_ TiledLayer) LevelsOfDetailBias() int {
	rv := objc.Send[int](t_.ID, objc.Sel("levelsOfDetailBias"))
	return rv
}


// The number of magnified levels of detail for this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/catiledlayer/levelsofdetailbias
func (t_ TiledLayer) SetLevelsOfDetailBias(value int) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLevelsOfDetailBias:"), value)
}



