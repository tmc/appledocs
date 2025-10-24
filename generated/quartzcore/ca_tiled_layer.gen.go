// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CATiledLayer */


/* debug [class_header]: Header for CATiledLayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TiledLayer */
// An interface definition for the [TiledLayer] class.
type ITiledLayer interface {
	ILayer
	
/* debug [class_interface_properties]: Properties for TiledLayer */
	// properties:
	LevelsOfDetail() uintptr /* not a class type */
	SetLevelsOfDetail(value uintptr /* not a class type */)
	LevelsOfDetailBias() uintptr /* not a class type */
	SetLevelsOfDetailBias(value uintptr /* not a class type */)
	TileSize() corefoundation.CGSize
	SetTileSize(value corefoundation.CGSize)
	Contents() objectivec.IObject
	SetContents(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TiledLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TiledLayer */
// Alloc allocates a new instance without initialization.
func (tc _TiledLayerClass) Alloc() TiledLayer {
	rv := objc.Send[TiledLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TiledLayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TiledLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TiledLayer */

// The time, in seconds, that newly added images take to “fade-in” to the rendered representation of the tiled layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/fadeDuration()
func (tc _TiledLayerClass) FadeDuration() float64 {
	rv := objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("fadeDuration"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FadeDuration) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TiledLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TiledLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TiledLayer */

// The number of levels of detail maintained by this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/levelsOfDetail
func (t_ TiledLayer) LevelsOfDetail() uintptr /* not a class type */ {
	rv := objc.Send[uintptr](t_.ID, objc.Sel("levelsOfDetail"))
	return rv
}/* debug [instance_properties/getter]: levelsOfDetail */


// The number of levels of detail maintained by this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/levelsOfDetail
func (t_ TiledLayer) SetLevelsOfDetail(value uintptr /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLevelsOfDetail:"), value)
}/* debug [instance_properties/setter]: levelsOfDetail */


// The number of magnified levels of detail for this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/levelsOfDetailBias
func (t_ TiledLayer) LevelsOfDetailBias() uintptr /* not a class type */ {
	rv := objc.Send[uintptr](t_.ID, objc.Sel("levelsOfDetailBias"))
	return rv
}/* debug [instance_properties/getter]: levelsOfDetailBias */


// The number of magnified levels of detail for this layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/levelsOfDetailBias
func (t_ TiledLayer) SetLevelsOfDetailBias(value uintptr /* not a class type */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLevelsOfDetailBias:"), value)
}/* debug [instance_properties/setter]: levelsOfDetailBias */


// The maximum size of each tile used to create the layer’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/tileSize
func (t_ TiledLayer) TileSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t_.ID, objc.Sel("tileSize"))
	return rv
}/* debug [instance_properties/getter]: tileSize */


// The maximum size of each tile used to create the layer’s content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/tileSize
func (t_ TiledLayer) SetTileSize(value corefoundation.CGSize) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTileSize:"), value)
}/* debug [instance_properties/setter]: tileSize */


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/contents
func (t_ TiledLayer) Contents() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](t_.ID, objc.Sel("contents"))
	return rv
}/* debug [instance_properties/getter]: contents */


// An object that provides the contents of the layer. Animatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/calayer/contents
func (t_ TiledLayer) SetContents(value objectivec.IObject) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setContents:"), value)
}/* debug [instance_properties/setter]: contents */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CATiledLayer */



