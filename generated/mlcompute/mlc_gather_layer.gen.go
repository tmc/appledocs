// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MLCGatherLayer */


/* debug [class_header]: Header for MLCGatherLayer */
// The class instance for the [CGatherLayer] class.
var (
	CGatherLayerClass     _CGatherLayerClass
	CGatherLayerClassOnce sync.Once
)

func getCGatherLayerClass() _CGatherLayerClass {
	CGatherLayerClassOnce.Do(func() {
		CGatherLayerClass = _CGatherLayerClass{objc.GetClass("MLCGatherLayer")}
	})
	return CGatherLayerClass
}

type _CGatherLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CGatherLayer */
// An interface definition for the [CGatherLayer] class.
type ICGatherLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CGatherLayer */
	// properties:
	Dimension() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CGatherLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CGatherLayer */
// Alloc allocates a new instance without initialization.
func (cc _CGatherLayerClass) Alloc() CGatherLayer {
	rv := objc.Send[CGatherLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CGatherLayerClass) New() CGatherLayer {
	rv := objc.Send[CGatherLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CGatherLayer) Init() CGatherLayer {
	rv := objc.Send[CGatherLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CGatherLayer) Autorelease() CGatherLayer {
	rv := objc.Send[CGatherLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCGatherLayer creates a new CGatherLayer instance.
func NewCGatherLayer() CGatherLayer {
	return getCGatherLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CGatherLayer */
// A layer that fetches data at the locations you specify.


// A layer that fetches data at the locations you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGatherLayer
type CGatherLayer struct {
	CLayer
}

// CGatherLayerFrom constructs a [CGatherLayer] from an unsafe.Pointer.
//
// A layer that fetches data at the locations you specify.
func CGatherLayerFrom(ptr unsafe.Pointer) CGatherLayer {
	return CGatherLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CGatherLayer */

// Creates a gather layer with the dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGatherLayer/init(dimension:)
func NewCGatherLayerWithDimension(dimension uint) CGatherLayer {
	rv := objc.Send[CGatherLayer](objc.ID(getCGatherLayerClass().class), objc.Sel("layerWithDimension:"), dimension)
	return rv
}/* debug [class_init_methods/constructor]: NewCGatherLayerWithDimension */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CGatherLayer */

// Creates a gather layer with the dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGatherLayer/init(dimension:)
func (cc _CGatherLayerClass) LayerWithDimension(dimension uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithDimension:"), dimension)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithDimension) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CGatherLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CGatherLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CGatherLayer */

// The dimension to index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCGatherLayer/dimension
func (c_ CGatherLayer) Dimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dimension"))
	return rv
}/* debug [instance_properties/getter]: dimension */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCGatherLayer */


