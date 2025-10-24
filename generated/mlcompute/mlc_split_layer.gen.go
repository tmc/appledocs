// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCSplitLayer */


/* debug [class_header]: Header for MLCSplitLayer */
// The class instance for the [CSplitLayer] class.
var (
	CSplitLayerClass     _CSplitLayerClass
	CSplitLayerClassOnce sync.Once
)

func getCSplitLayerClass() _CSplitLayerClass {
	CSplitLayerClassOnce.Do(func() {
		CSplitLayerClass = _CSplitLayerClass{objc.GetClass("MLCSplitLayer")}
	})
	return CSplitLayerClass
}

type _CSplitLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CSplitLayer */
// An interface definition for the [CSplitLayer] class.
type ICSplitLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CSplitLayer */
	// properties:
	Dimension() uint
	SplitCount() uint
	SplitSectionLengths() []foundation.Number
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CSplitLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CSplitLayer */
// Alloc allocates a new instance without initialization.
func (cc _CSplitLayerClass) Alloc() CSplitLayer {
	rv := objc.Send[CSplitLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CSplitLayerClass) New() CSplitLayer {
	rv := objc.Send[CSplitLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CSplitLayer) Init() CSplitLayer {
	rv := objc.Send[CSplitLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CSplitLayer) Autorelease() CSplitLayer {
	rv := objc.Send[CSplitLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCSplitLayer creates a new CSplitLayer instance.
func NewCSplitLayer() CSplitLayer {
	return getCSplitLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CSplitLayer */
// A layer that splits a tensor value into a list of subtensors.


// A layer that splits a tensor value into a list of subtensors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSplitLayer
type CSplitLayer struct {
	CLayer
}

// CSplitLayerFrom constructs a [CSplitLayer] from an unsafe.Pointer.
//
// A layer that splits a tensor value into a list of subtensors.
func CSplitLayerFrom(ptr unsafe.Pointer) CSplitLayer {
	return CSplitLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CSplitLayer */

// Creates a split layer with the number of splits and dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSplitLayer/init(splitCount:dimension:)
func NewCSplitLayerWithSplitCountDimension(splitCount uint, dimension uint) CSplitLayer {
	rv := objc.Send[CSplitLayer](objc.ID(getCSplitLayerClass().class), objc.Sel("layerWithSplitCount:dimension:"), splitCount, dimension)
	return rv
}/* debug [class_init_methods/constructor]: NewCSplitLayerWithSplitCountDimension */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CSplitLayer */

// Creates a split layer with the number of splits and dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSplitLayer/init(splitCount:dimension:)
func (cc _CSplitLayerClass) LayerWithSplitCountDimension(splitCount uint, dimension uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithSplitCount:dimension:"), splitCount, dimension)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithSplitCountDimension) */


// Creates a split layer with the lengths of each split section and dimension you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSplitLayer/layerWithSplitSectionLengths:dimension:
func (cc _CSplitLayerClass) LayerWithSplitSectionLengthsDimension(splitSectionLengths []foundation.Number, dimension uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithSplitSectionLengths:dimension:"), splitSectionLengths, dimension)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithSplitSectionLengthsDimension) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CSplitLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CSplitLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CSplitLayer */

// The dimension or axis along which to split the tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSplitLayer/dimension
func (c_ CSplitLayer) Dimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("dimension"))
	return rv
}/* debug [instance_properties/getter]: dimension */


// The number of splits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSplitLayer/splitCount
func (c_ CSplitLayer) SplitCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("splitCount"))
	return rv
}/* debug [instance_properties/getter]: splitCount */


// An array that contains the lengths of each split section.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCSplitLayer/splitSectionLengths-32abw
func (c_ CSplitLayer) SplitSectionLengths() []foundation.Number {
	rv := objc.Send[[]foundation.Number](c_.ID, objc.Sel("splitSectionLengths"))
	return rv
}/* debug [instance_properties/getter]: splitSectionLengths */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCSplitLayer */


