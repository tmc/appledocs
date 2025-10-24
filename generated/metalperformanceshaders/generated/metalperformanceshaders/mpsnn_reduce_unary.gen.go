// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReduceUnary */


/* debug [class_header]: Header for MPSNNReduceUnary */
// The class instance for the [ReduceUnary] class.
var (
	ReduceUnaryClass     _ReduceUnaryClass
	ReduceUnaryClassOnce sync.Once
)

func getReduceUnaryClass() _ReduceUnaryClass {
	ReduceUnaryClassOnce.Do(func() {
		ReduceUnaryClass = _ReduceUnaryClass{objc.GetClass("MPSNNReduceUnary")}
	})
	return ReduceUnaryClass
}

type _ReduceUnaryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceUnary */
// An interface definition for the [ReduceUnary] class.
type IReduceUnary interface {
	ICNNKernel
	
/* debug [class_interface_properties]: Properties for ReduceUnary */
	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
	Offset() Offset get set /* not a class type */
	SetOffset(value Offset get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceUnary */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceUnary */
// Alloc allocates a new instance without initialization.
func (rc _ReduceUnaryClass) Alloc() ReduceUnary {
	rv := objc.Send[ReduceUnary](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceUnaryClass) New() ReduceUnary {
	rv := objc.Send[ReduceUnary](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceUnary) Init() ReduceUnary {
	rv := objc.Send[ReduceUnary](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceUnary) Autorelease() ReduceUnary {
	rv := objc.Send[ReduceUnary](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceUnary creates a new ReduceUnary instance.
func NewReduceUnary() ReduceUnary {
	return getReduceUnaryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceUnary */
// The base class for unary reduction filters.


// The base class for unary reduction filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceUnary
type ReduceUnary struct {
	CNNKernel
}

// ReduceUnaryFrom constructs a [ReduceUnary] from an unsafe.Pointer.
//
// The base class for unary reduction filters.
func ReduceUnaryFrom(ptr unsafe.Pointer) ReduceUnary {
	return ReduceUnary{
		CNNKernel: CNNKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceUnary *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceUnary */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceUnary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceUnary */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceUnary */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/2942547-cliprectsource
func (r_ ReduceUnary) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("clipRectSource"))
	return rv
}/* debug [instance_properties/getter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/2942547-cliprectsource
func (r_ ReduceUnary) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setClipRectSource:"), value)
}/* debug [instance_properties/setter]: clipRectSource */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/3750642-offset
func (r_ ReduceUnary) Offset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("offset"))
	return rv
}/* debug [instance_properties/getter]: offset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/3750642-offset
func (r_ ReduceUnary) SetOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOffset:"), value)
}/* debug [instance_properties/setter]: offset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceUnary */



