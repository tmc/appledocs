// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSMatrixLogSoftMaxGradient */


/* debug [class_header]: Header for MPSMatrixLogSoftMaxGradient */
// The class instance for the [MatrixLogSoftMaxGradient] class.
var (
	MatrixLogSoftMaxGradientClass     _MatrixLogSoftMaxGradientClass
	MatrixLogSoftMaxGradientClassOnce sync.Once
)

func getMatrixLogSoftMaxGradientClass() _MatrixLogSoftMaxGradientClass {
	MatrixLogSoftMaxGradientClassOnce.Do(func() {
		MatrixLogSoftMaxGradientClass = _MatrixLogSoftMaxGradientClass{objc.GetClass("MPSMatrixLogSoftMaxGradient")}
	})
	return MatrixLogSoftMaxGradientClass
}

type _MatrixLogSoftMaxGradientClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixLogSoftMaxGradient */
// An interface definition for the [MatrixLogSoftMaxGradient] class.
type IMatrixLogSoftMaxGradient interface {
	IMatrixSoftMaxGradient
	
/* debug [class_interface_properties]: Properties for MatrixLogSoftMaxGradient */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixLogSoftMaxGradient */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixLogSoftMaxGradient */
// Alloc allocates a new instance without initialization.
func (mc _MatrixLogSoftMaxGradientClass) Alloc() MatrixLogSoftMaxGradient {
	rv := objc.Send[MatrixLogSoftMaxGradient](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixLogSoftMaxGradientClass) New() MatrixLogSoftMaxGradient {
	rv := objc.Send[MatrixLogSoftMaxGradient](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixLogSoftMaxGradient) Init() MatrixLogSoftMaxGradient {
	rv := objc.Send[MatrixLogSoftMaxGradient](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixLogSoftMaxGradient) Autorelease() MatrixLogSoftMaxGradient {
	rv := objc.Send[MatrixLogSoftMaxGradient](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixLogSoftMaxGradient creates a new MatrixLogSoftMaxGradient instance.
func NewMatrixLogSoftMaxGradient() MatrixLogSoftMaxGradient {
	return getMatrixLogSoftMaxGradientClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixLogSoftMaxGradient */
// A logarithmic gradient softmax kernel that operates on matrices.


// A logarithmic gradient softmax kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixLogSoftMaxGradient
type MatrixLogSoftMaxGradient struct {
	MatrixSoftMaxGradient
}

// MatrixLogSoftMaxGradientFrom constructs a [MatrixLogSoftMaxGradient] from an unsafe.Pointer.
//
// A logarithmic gradient softmax kernel that operates on matrices.
func MatrixLogSoftMaxGradientFrom(ptr unsafe.Pointer) MatrixLogSoftMaxGradient {
	return MatrixLogSoftMaxGradient{
		MatrixSoftMaxGradient: MatrixSoftMaxGradientFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixLogSoftMaxGradient *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixLogSoftMaxGradient */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixLogSoftMaxGradient */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixLogSoftMaxGradient */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixLogSoftMaxGradient */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixLogSoftMaxGradient */



