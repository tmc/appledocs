// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCMultiheadAttentionDescriptor */


/* debug [class_header]: Header for MLCMultiheadAttentionDescriptor */
// The class instance for the [CMultiheadAttentionDescriptor] class.
var (
	CMultiheadAttentionDescriptorClass     _CMultiheadAttentionDescriptorClass
	CMultiheadAttentionDescriptorClassOnce sync.Once
)

func getCMultiheadAttentionDescriptorClass() _CMultiheadAttentionDescriptorClass {
	CMultiheadAttentionDescriptorClassOnce.Do(func() {
		CMultiheadAttentionDescriptorClass = _CMultiheadAttentionDescriptorClass{objc.GetClass("MLCMultiheadAttentionDescriptor")}
	})
	return CMultiheadAttentionDescriptorClass
}

type _CMultiheadAttentionDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CMultiheadAttentionDescriptor */
// An interface definition for the [CMultiheadAttentionDescriptor] class.
type ICMultiheadAttentionDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CMultiheadAttentionDescriptor */
	// properties:
	AddsZeroAttention() bool
	Dropout() float32
	HasAttentionBiases() bool
	HasBiases() bool
	HeadCount() uint
	KeyDimension() uint
	ModelDimension() uint
	ValueDimension() uint
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CMultiheadAttentionDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CMultiheadAttentionDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CMultiheadAttentionDescriptorClass) Alloc() CMultiheadAttentionDescriptor {
	rv := objc.Send[CMultiheadAttentionDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CMultiheadAttentionDescriptorClass) New() CMultiheadAttentionDescriptor {
	rv := objc.Send[CMultiheadAttentionDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CMultiheadAttentionDescriptor) Init() CMultiheadAttentionDescriptor {
	rv := objc.Send[CMultiheadAttentionDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CMultiheadAttentionDescriptor) Autorelease() CMultiheadAttentionDescriptor {
	rv := objc.Send[CMultiheadAttentionDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCMultiheadAttentionDescriptor creates a new CMultiheadAttentionDescriptor instance.
func NewCMultiheadAttentionDescriptor() CMultiheadAttentionDescriptor {
	return getCMultiheadAttentionDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CMultiheadAttentionDescriptor */
// A configuration object you use to create a multi-head attention layer.


// A configuration object you use to create a multi-head attention layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor
type CMultiheadAttentionDescriptor struct {
	objectivec.Object
}

// CMultiheadAttentionDescriptorFrom constructs a [CMultiheadAttentionDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create a multi-head attention layer.
func CMultiheadAttentionDescriptorFrom(ptr unsafe.Pointer) CMultiheadAttentionDescriptor {
	return CMultiheadAttentionDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CMultiheadAttentionDescriptor */

// Creates a multi-head attention descriptor with the model dimension and number of parallel attention heads you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/init(modelDimension:headCount:)
func NewCMultiheadAttentionDescriptorWithModelDimensionHeadCount(modelDimension uint, headCount uint) CMultiheadAttentionDescriptor {
	rv := objc.Send[CMultiheadAttentionDescriptor](objc.ID(getCMultiheadAttentionDescriptorClass().class), objc.Sel("descriptorWithModelDimension:headCount:"), modelDimension, headCount)
	return rv
}/* debug [class_init_methods/constructor]: NewCMultiheadAttentionDescriptorWithModelDimensionHeadCount */


// Creates a multi-head attention descriptor with the dimensions, number of attention heads, dropout rate, and bias and padding options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/init(modelDimension:keyDimension:valueDimension:headCount:dropout:hasBiases:hasAttentionBiases:addsZeroAttention:)
func NewCMultiheadAttentionDescriptorWithModelDimensionKeyDimensionValueDimensionHeadCountDropoutHasBiasesHasAttentionBiasesAddsZeroAttention(modelDimension uint, keyDimension uint, valueDimension uint, headCount uint, dropout float32, hasBiases bool, hasAttentionBiases bool, addsZeroAttention bool) CMultiheadAttentionDescriptor {
	rv := objc.Send[CMultiheadAttentionDescriptor](objc.ID(getCMultiheadAttentionDescriptorClass().class), objc.Sel("descriptorWithModelDimension:keyDimension:valueDimension:headCount:dropout:hasBiases:hasAttentionBiases:addsZeroAttention:"), modelDimension, keyDimension, valueDimension, headCount, dropout, hasBiases, hasAttentionBiases, addsZeroAttention)
	return rv
}/* debug [class_init_methods/constructor]: NewCMultiheadAttentionDescriptorWithModelDimensionKeyDimensionValueDimensionHeadCountDropoutHasBiasesHasAttentionBiasesAddsZeroAttention */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CMultiheadAttentionDescriptor */

// Creates a multi-head attention descriptor with the model dimension and number of parallel attention heads you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/init(modelDimension:headCount:)
func (cc _CMultiheadAttentionDescriptorClass) DescriptorWithModelDimensionHeadCount(modelDimension uint, headCount uint) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithModelDimension:headCount:"), modelDimension, headCount)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithModelDimensionHeadCount) */


// Creates a multi-head attention descriptor with the dimensions, number of attention heads, dropout rate, and bias and padding options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/init(modelDimension:keyDimension:valueDimension:headCount:dropout:hasBiases:hasAttentionBiases:addsZeroAttention:)
func (cc _CMultiheadAttentionDescriptorClass) DescriptorWithModelDimensionKeyDimensionValueDimensionHeadCountDropoutHasBiasesHasAttentionBiasesAddsZeroAttention(modelDimension uint, keyDimension uint, valueDimension uint, headCount uint, dropout float32, hasBiases bool, hasAttentionBiases bool, addsZeroAttention bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithModelDimension:keyDimension:valueDimension:headCount:dropout:hasBiases:hasAttentionBiases:addsZeroAttention:"), modelDimension, keyDimension, valueDimension, headCount, dropout, hasBiases, hasAttentionBiases, addsZeroAttention)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithModelDimensionKeyDimensionValueDimensionHeadCountDropoutHasBiasesHasAttentionBiasesAddsZeroAttention) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CMultiheadAttentionDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CMultiheadAttentionDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CMultiheadAttentionDescriptor */

// A Boolean that specifies whether you add a row of zeros to projected key and value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/addsZeroAttention
func (c_ CMultiheadAttentionDescriptor) AddsZeroAttention() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("addsZeroAttention"))
	return rv
}/* debug [instance_properties/getter]: addsZeroAttention */


// The dropout rate you apply to the output projection weights.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/dropout
func (c_ CMultiheadAttentionDescriptor) Dropout() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("dropout"))
	return rv
}/* debug [instance_properties/getter]: dropout */


// A Boolean that specifies whether you add an array of biases to key and value, respectively.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/hasAttentionBiases
func (c_ CMultiheadAttentionDescriptor) HasAttentionBiases() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasAttentionBiases"))
	return rv
}/* debug [instance_properties/getter]: hasAttentionBiases */


// A Boolean that specifies whether you add a bias to query, key, value, and output projections.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/hasBiases
func (c_ CMultiheadAttentionDescriptor) HasBiases() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasBiases"))
	return rv
}/* debug [instance_properties/getter]: hasBiases */


// The number of parallel attention heads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/headCount
func (c_ CMultiheadAttentionDescriptor) HeadCount() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("headCount"))
	return rv
}/* debug [instance_properties/getter]: headCount */


// The total dimension of key space, which must be divisible by the number of heads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/keyDimension
func (c_ CMultiheadAttentionDescriptor) KeyDimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("keyDimension"))
	return rv
}/* debug [instance_properties/getter]: keyDimension */


// The model or embedding dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/modelDimension
func (c_ CMultiheadAttentionDescriptor) ModelDimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("modelDimension"))
	return rv
}/* debug [instance_properties/getter]: modelDimension */


// The total dimension of value space, which must be divisible by the number of heads.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMultiheadAttentionDescriptor/valueDimension
func (c_ CMultiheadAttentionDescriptor) ValueDimension() uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("valueDimension"))
	return rv
}/* debug [instance_properties/getter]: valueDimension */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCMultiheadAttentionDescriptor */


