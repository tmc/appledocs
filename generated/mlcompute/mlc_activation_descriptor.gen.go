// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCActivationDescriptor */


/* debug [class_header]: Header for MLCActivationDescriptor */
// The class instance for the [CActivationDescriptor] class.
var (
	CActivationDescriptorClass     _CActivationDescriptorClass
	CActivationDescriptorClassOnce sync.Once
)

func getCActivationDescriptorClass() _CActivationDescriptorClass {
	CActivationDescriptorClassOnce.Do(func() {
		CActivationDescriptorClass = _CActivationDescriptorClass{objc.GetClass("MLCActivationDescriptor")}
	})
	return CActivationDescriptorClass
}

type _CActivationDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CActivationDescriptor */
// An interface definition for the [CActivationDescriptor] class.
type ICActivationDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CActivationDescriptor */
	// properties:
	A() float32
	ActivationType() CActivationType
	B() float32
	C() float32
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CActivationDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CActivationDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CActivationDescriptorClass) Alloc() CActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CActivationDescriptorClass) New() CActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CActivationDescriptor) Init() CActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CActivationDescriptor) Autorelease() CActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCActivationDescriptor creates a new CActivationDescriptor instance.
func NewCActivationDescriptor() CActivationDescriptor {
	return getCActivationDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CActivationDescriptor */
// A configuration object you use to create an activation layer.
//
// The framework provides the following activation descriptor initializers to create the associated descriptors:


// A configuration object you use to create an activation layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor
type CActivationDescriptor struct {
	objectivec.Object
}

// CActivationDescriptorFrom constructs a [CActivationDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create an activation layer.
func CActivationDescriptorFrom(ptr unsafe.Pointer) CActivationDescriptor {
	return CActivationDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CActivationDescriptor */

// Creates an activation descriptor with the activation type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/init(type:)
func NewCActivationDescriptorWithType(activationType CActivationType) CActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](objc.ID(getCActivationDescriptorClass().class), objc.Sel("descriptorWithType:"), activationType)
	return rv
}/* debug [class_init_methods/constructor]: NewCActivationDescriptorWithType */


// Creates an activation descriptor with the activation type and parameter a that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/init(type:a:)
func NewCActivationDescriptorWithTypeA(activationType CActivationType, a float32) CActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](objc.ID(getCActivationDescriptorClass().class), objc.Sel("descriptorWithType:a:"), activationType, a)
	return rv
}/* debug [class_init_methods/constructor]: NewCActivationDescriptorWithTypeA */


// Creates an activation descriptor with the activation type and parameters a and b that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/init(type:a:b:)
func NewCActivationDescriptorWithTypeAB(activationType CActivationType, a float32, b float32) CActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](objc.ID(getCActivationDescriptorClass().class), objc.Sel("descriptorWithType:a:b:"), activationType, a, b)
	return rv
}/* debug [class_init_methods/constructor]: NewCActivationDescriptorWithTypeAB */


// Creates an activation descriptor with the activation type and parameters a, b, and c that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/init(type:a:b:c:)
func NewCActivationDescriptorWithTypeABC(activationType CActivationType, a float32, b float32, c float32) CActivationDescriptor {
	rv := objc.Send[CActivationDescriptor](objc.ID(getCActivationDescriptorClass().class), objc.Sel("descriptorWithType:a:b:c:"), activationType, a, b, c)
	return rv
}/* debug [class_init_methods/constructor]: NewCActivationDescriptorWithTypeABC */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CActivationDescriptor */

// Creates an activation descriptor with the activation type you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/init(type:)
func (cc _CActivationDescriptorClass) DescriptorWithType(activationType CActivationType) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithType:"), activationType)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithType) */


// Creates an activation descriptor with the activation type and parameter a that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/init(type:a:)
func (cc _CActivationDescriptorClass) DescriptorWithTypeA(activationType CActivationType, a float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithType:a:"), activationType, a)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithTypeA) */


// Creates an activation descriptor with the activation type and parameters a and b that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/init(type:a:b:)
func (cc _CActivationDescriptorClass) DescriptorWithTypeAB(activationType CActivationType, a float32, b float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithType:a:b:"), activationType, a, b)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithTypeAB) */


// Creates an activation descriptor with the activation type and parameters a, b, and c that you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/init(type:a:b:c:)
func (cc _CActivationDescriptorClass) DescriptorWithTypeABC(activationType CActivationType, a float32, b float32, c float32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithType:a:b:c:"), activationType, a, b, c)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithTypeABC) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CActivationDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CActivationDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CActivationDescriptor */

// The parameter a to the activation function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/a
func (c_ CActivationDescriptor) A() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("a"))
	return rv
}/* debug [instance_properties/getter]: a */


// The type of activation function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/activationType
func (c_ CActivationDescriptor) ActivationType() CActivationType {
	rv := objc.Send[CActivationType](c_.ID, objc.Sel("activationType"))
	return rv
}/* debug [instance_properties/getter]: activationType */


// The parameter b to the activation function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/b
func (c_ CActivationDescriptor) B() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("b"))
	return rv
}/* debug [instance_properties/getter]: b */


// The parameter c to the activation function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCActivationDescriptor/c
func (c_ CActivationDescriptor) C() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("c"))
	return rv
}/* debug [instance_properties/getter]: c */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCActivationDescriptor */


