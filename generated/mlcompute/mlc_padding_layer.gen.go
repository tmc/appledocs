// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class MLCPaddingLayer */


/* debug [class_header]: Header for MLCPaddingLayer */
// The class instance for the [CPaddingLayer] class.
var (
	CPaddingLayerClass     _CPaddingLayerClass
	CPaddingLayerClassOnce sync.Once
)

func getCPaddingLayerClass() _CPaddingLayerClass {
	CPaddingLayerClassOnce.Do(func() {
		CPaddingLayerClass = _CPaddingLayerClass{objc.GetClass("MLCPaddingLayer")}
	})
	return CPaddingLayerClass
}

type _CPaddingLayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CPaddingLayer */
// An interface definition for the [CPaddingLayer] class.
type ICPaddingLayer interface {
	ICLayer
	
/* debug [class_interface_properties]: Properties for CPaddingLayer */
	// properties:
	ConstantValue() float32
	PaddingBottom() uint
	PaddingLeft() uint
	PaddingRight() uint
	PaddingTop() uint
	PaddingType() CPaddingType
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CPaddingLayer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CPaddingLayer */
// Alloc allocates a new instance without initialization.
func (cc _CPaddingLayerClass) Alloc() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CPaddingLayerClass) New() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CPaddingLayer) Init() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CPaddingLayer) Autorelease() CPaddingLayer {
	rv := objc.Send[CPaddingLayer](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCPaddingLayer creates a new CPaddingLayer instance.
func NewCPaddingLayer() CPaddingLayer {
	return getCPaddingLayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CPaddingLayer */
// A layer that pads a tensor with the padding sizes you specify.


// A layer that pads a tensor with the padding sizes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer
type CPaddingLayer struct {
	CLayer
}

// CPaddingLayerFrom constructs a [CPaddingLayer] from an unsafe.Pointer.
//
// A layer that pads a tensor with the padding sizes you specify.
func CPaddingLayerFrom(ptr unsafe.Pointer) CPaddingLayer {
	return CPaddingLayer{
		CLayer: CLayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CPaddingLayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CPaddingLayer */

// Creates a padding layer with the constant padding sizes and constant valu you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/layerWithConstantPadding:constantValue:
func (cc _CPaddingLayerClass) LayerWithConstantPaddingConstantValue(padding []foundation.Number, constantValue float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithConstantPadding:constantValue:"), padding, constantValue)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithConstantPaddingConstantValue) */


// Creates a padding layer with the reflection padding sizes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/layerWithReflectionPadding:
func (cc _CPaddingLayerClass) LayerWithReflectionPadding(padding []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithReflectionPadding:"), padding)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithReflectionPadding) */


// Creates a padding layer with the symmetric padding sizes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/layerWithSymmetricPadding:
func (cc _CPaddingLayerClass) LayerWithSymmetricPadding(padding []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithSymmetricPadding:"), padding)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithSymmetricPadding) */


// Creates a padding layer with the zero padding sizes you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/layerWithZeroPadding:
func (cc _CPaddingLayerClass) LayerWithZeroPadding(padding []foundation.Number) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("layerWithZeroPadding:"), padding)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LayerWithZeroPadding) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CPaddingLayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CPaddingLayer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CPaddingLayer */

// The constant value you use if padding type is constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/constantValue
func (c_ CPaddingLayer) ConstantValue() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("constantValue"))
	return rv
}/* debug [instance_properties/getter]: constantValue */


// The bottom padding size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/paddingBottom
func (c_ CPaddingLayer) PaddingBottom() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("paddingBottom"))
	return rv
}/* debug [instance_properties/getter]: paddingBottom */


// The left padding size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/paddingLeft
func (c_ CPaddingLayer) PaddingLeft() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("paddingLeft"))
	return rv
}/* debug [instance_properties/getter]: paddingLeft */


// The right padding size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/paddingRight
func (c_ CPaddingLayer) PaddingRight() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("paddingRight"))
	return rv
}/* debug [instance_properties/getter]: paddingRight */


// The top padding size.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/paddingTop
func (c_ CPaddingLayer) PaddingTop() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("paddingTop"))
	return rv
}/* debug [instance_properties/getter]: paddingTop */


// The padding type.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCPaddingLayer/paddingType
func (c_ CPaddingLayer) PaddingType() CPaddingType {
	rv := objc.Send[CPaddingType](c_.ID, objc.Sel("paddingType"))
	return rv
}/* debug [instance_properties/getter]: paddingType */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCPaddingLayer */



