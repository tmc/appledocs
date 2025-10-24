// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSMatrixLogSoftMax */


/* debug [class_header]: Header for MPSMatrixLogSoftMax */
// The class instance for the [MatrixLogSoftMax] class.
var (
	MatrixLogSoftMaxClass     _MatrixLogSoftMaxClass
	MatrixLogSoftMaxClassOnce sync.Once
)

func getMatrixLogSoftMaxClass() _MatrixLogSoftMaxClass {
	MatrixLogSoftMaxClassOnce.Do(func() {
		MatrixLogSoftMaxClass = _MatrixLogSoftMaxClass{objc.GetClass("MPSMatrixLogSoftMax")}
	})
	return MatrixLogSoftMaxClass
}

type _MatrixLogSoftMaxClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixLogSoftMax */
// An interface definition for the [MatrixLogSoftMax] class.
type IMatrixLogSoftMax interface {
	IMatrixSoftMax
	
/* debug [class_interface_properties]: Properties for MatrixLogSoftMax */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixLogSoftMax */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixLogSoftMax */
// Alloc allocates a new instance without initialization.
func (mc _MatrixLogSoftMaxClass) Alloc() MatrixLogSoftMax {
	rv := objc.Send[MatrixLogSoftMax](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixLogSoftMaxClass) New() MatrixLogSoftMax {
	rv := objc.Send[MatrixLogSoftMax](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixLogSoftMax) Init() MatrixLogSoftMax {
	rv := objc.Send[MatrixLogSoftMax](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixLogSoftMax) Autorelease() MatrixLogSoftMax {
	rv := objc.Send[MatrixLogSoftMax](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixLogSoftMax creates a new MatrixLogSoftMax instance.
func NewMatrixLogSoftMax() MatrixLogSoftMax {
	return getMatrixLogSoftMaxClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixLogSoftMax */
// A logarithmic softmax kernel that operates on matrices.


// A logarithmic softmax kernel that operates on matrices.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixLogSoftMax
type MatrixLogSoftMax struct {
	MatrixSoftMax
}

// MatrixLogSoftMaxFrom constructs a [MatrixLogSoftMax] from an unsafe.Pointer.
//
// A logarithmic softmax kernel that operates on matrices.
func MatrixLogSoftMaxFrom(ptr unsafe.Pointer) MatrixLogSoftMax {
	return MatrixLogSoftMax{
		MatrixSoftMax: MatrixSoftMaxFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixLogSoftMax *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixLogSoftMax */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixLogSoftMax */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixLogSoftMax */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixLogSoftMax */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixLogSoftMax */



