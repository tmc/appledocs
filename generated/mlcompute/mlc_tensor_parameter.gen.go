// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCTensorParameter */


/* debug [class_header]: Header for MLCTensorParameter */
// The class instance for the [CTensorParameter] class.
var (
	CTensorParameterClass     _CTensorParameterClass
	CTensorParameterClassOnce sync.Once
)

func getCTensorParameterClass() _CTensorParameterClass {
	CTensorParameterClassOnce.Do(func() {
		CTensorParameterClass = _CTensorParameterClass{objc.GetClass("MLCTensorParameter")}
	})
	return CTensorParameterClass
}

type _CTensorParameterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CTensorParameter */
// An interface definition for the [CTensorParameter] class.
type ICTensorParameter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CTensorParameter */
	// properties:
	IsUpdatable() bool
	SetIsUpdatable(value bool)
	Tensor() IMLCTensor
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CTensorParameter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CTensorParameter */
// Alloc allocates a new instance without initialization.
func (cc _CTensorParameterClass) Alloc() CTensorParameter {
	rv := objc.Send[CTensorParameter](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CTensorParameterClass) New() CTensorParameter {
	rv := objc.Send[CTensorParameter](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTensorParameter) Init() CTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTensorParameter) Autorelease() CTensorParameter {
	rv := objc.Send[CTensorParameter](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTensorParameter creates a new CTensorParameter instance.
func NewCTensorParameter() CTensorParameter {
	return getCTensorParameterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CTensorParameter */
// A tensor parameter object.
//
// Use a tensor parameter to describe input tensors that the optimizer updates during training.


// A tensor parameter object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorParameter
type CTensorParameter struct {
	objectivec.Object
}

// CTensorParameterFrom constructs a [CTensorParameter] from an unsafe.Pointer.
//
// A tensor parameter object.
func CTensorParameterFrom(ptr unsafe.Pointer) CTensorParameter {
	return CTensorParameter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CTensorParameter */

// Creates a tensor parameter with the tensor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorParameter/init(tensor:)
func NewCTensorParameterWithTensor(tensor IMLCTensor) CTensorParameter {
	rv := objc.Send[CTensorParameter](objc.ID(getCTensorParameterClass().class), objc.Sel("parameterWithTensor:"), tensor)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorParameterWithTensor */


// Creates a tensor parameter with the tensor and optimizer data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorParameter/init(tensor:optimizerData:)
func NewCTensorParameterWithTensorOptimizerData(tensor IMLCTensor, optimizerData []CTensorData) CTensorParameter {
	rv := objc.Send[CTensorParameter](objc.ID(getCTensorParameterClass().class), objc.Sel("parameterWithTensor:optimizerData:"), tensor, optimizerData)
	return rv
}/* debug [class_init_methods/constructor]: NewCTensorParameterWithTensorOptimizerData */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CTensorParameter */

// Creates a tensor parameter with the tensor you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorParameter/init(tensor:)
func (cc _CTensorParameterClass) ParameterWithTensor(tensor IMLCTensor) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("parameterWithTensor:"), tensor)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ParameterWithTensor) */


// Creates a tensor parameter with the tensor and optimizer data you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorParameter/init(tensor:optimizerData:)
func (cc _CTensorParameterClass) ParameterWithTensorOptimizerData(tensor IMLCTensor, optimizerData []CTensorData) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("parameterWithTensor:optimizerData:"), tensor, optimizerData)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ParameterWithTensorOptimizerData) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CTensorParameter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CTensorParameter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CTensorParameter */

// A Boolean that indicates whether this tensor parameter is updatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorParameter/isUpdatable
func (c_ CTensorParameter) IsUpdatable() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isUpdatable"))
	return rv
}/* debug [instance_properties/getter]: isUpdatable */


// A Boolean that indicates whether this tensor parameter is updatable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorParameter/isUpdatable
func (c_ CTensorParameter) SetIsUpdatable(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsUpdatable:"), value)
}/* debug [instance_properties/setter]: isUpdatable */


// The underlying tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorParameter/tensor
func (c_ CTensorParameter) Tensor() IMLCTensor {
	rv := objc.Send[CTensor](c_.ID, objc.Sel("tensor"))
	return rv
}/* debug [instance_properties/getter]: tensor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCTensorParameter */


