// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCMatMulDescriptor */


/* debug [class_header]: Header for MLCMatMulDescriptor */
// The class instance for the [CMatMulDescriptor] class.
var (
	CMatMulDescriptorClass     _CMatMulDescriptorClass
	CMatMulDescriptorClassOnce sync.Once
)

func getCMatMulDescriptorClass() _CMatMulDescriptorClass {
	CMatMulDescriptorClassOnce.Do(func() {
		CMatMulDescriptorClass = _CMatMulDescriptorClass{objc.GetClass("MLCMatMulDescriptor")}
	})
	return CMatMulDescriptorClass
}

type _CMatMulDescriptorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CMatMulDescriptor */
// An interface definition for the [CMatMulDescriptor] class.
type ICMatMulDescriptor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CMatMulDescriptor */
	// properties:
	Alpha() float32
	TransposesX() bool
	TransposesY() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CMatMulDescriptor */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CMatMulDescriptor */
// Alloc allocates a new instance without initialization.
func (cc _CMatMulDescriptorClass) Alloc() CMatMulDescriptor {
	rv := objc.Send[CMatMulDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CMatMulDescriptorClass) New() CMatMulDescriptor {
	rv := objc.Send[CMatMulDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CMatMulDescriptor) Init() CMatMulDescriptor {
	rv := objc.Send[CMatMulDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CMatMulDescriptor) Autorelease() CMatMulDescriptor {
	rv := objc.Send[CMatMulDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCMatMulDescriptor creates a new CMatMulDescriptor instance.
func NewCMatMulDescriptor() CMatMulDescriptor {
	return getCMatMulDescriptorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CMatMulDescriptor */
// A configuration object you use to create a matrix multiplication layer.


// A configuration object you use to create a matrix multiplication layer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulDescriptor
type CMatMulDescriptor struct {
	objectivec.Object
}

// CMatMulDescriptorFrom constructs a [CMatMulDescriptor] from an unsafe.Pointer.
//
// A configuration object you use to create a matrix multiplication layer.
func CMatMulDescriptorFrom(ptr unsafe.Pointer) CMatMulDescriptor {
	return CMatMulDescriptor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CMatMulDescriptor */

// Creates a batched matrix multiplication descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulDescriptor/init()
func NewCMatMulDescriptor() CMatMulDescriptor {
	rv := objc.Send[CMatMulDescriptor](objc.ID(getCMatMulDescriptorClass().class), objc.Sel("descriptor"))
	return rv
}/* debug [class_init_methods/constructor]: NewCMatMulDescriptor */


// Creates a batched matrix multiplication descriptor with the alpha value and transpose options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulDescriptor/init(alpha:transposesX:transposesY:)
func NewCMatMulDescriptorWithAlphaTransposesXTransposesY(alpha float32, transposesX bool, transposesY bool) CMatMulDescriptor {
	rv := objc.Send[CMatMulDescriptor](objc.ID(getCMatMulDescriptorClass().class), objc.Sel("descriptorWithAlpha:transposesX:transposesY:"), alpha, transposesX, transposesY)
	return rv
}/* debug [class_init_methods/constructor]: NewCMatMulDescriptorWithAlphaTransposesXTransposesY */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CMatMulDescriptor */

// Creates a batched matrix multiplication descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulDescriptor/init()
func (cc _CMatMulDescriptorClass) Descriptor() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptor"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Descriptor) */


// Creates a batched matrix multiplication descriptor with the alpha value and transpose options you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulDescriptor/init(alpha:transposesX:transposesY:)
func (cc _CMatMulDescriptorClass) DescriptorWithAlphaTransposesXTransposesY(alpha float32, transposesX bool, transposesY bool) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("descriptorWithAlpha:transposesX:transposesY:"), alpha, transposesX, transposesY)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DescriptorWithAlphaTransposesXTransposesY) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CMatMulDescriptor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CMatMulDescriptor */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CMatMulDescriptor */

// A scalar value you specify to scale the result in C = alpha x A x B.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulDescriptor/alpha
func (c_ CMatMulDescriptor) Alpha() float32 {
	rv := objc.Send[float32](c_.ID, objc.Sel("alpha"))
	return rv
}/* debug [instance_properties/getter]: alpha */


// A Boolean that specifies whether you choose to transpose the last two dimensions of x.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulDescriptor/transposesX
func (c_ CMatMulDescriptor) TransposesX() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("transposesX"))
	return rv
}/* debug [instance_properties/getter]: transposesX */


// A Boolean that specifies whether you choose to transpose the last two dimensions of y.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCMatMulDescriptor/transposesY
func (c_ CMatMulDescriptor) TransposesY() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("transposesY"))
	return rv
}/* debug [instance_properties/getter]: transposesY */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCMatMulDescriptor */


