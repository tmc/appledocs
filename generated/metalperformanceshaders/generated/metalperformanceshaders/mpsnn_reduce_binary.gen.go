// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class MPSNNReduceBinary */


/* debug [class_header]: Header for MPSNNReduceBinary */
// The class instance for the [ReduceBinary] class.
var (
	ReduceBinaryClass     _ReduceBinaryClass
	ReduceBinaryClassOnce sync.Once
)

func getReduceBinaryClass() _ReduceBinaryClass {
	ReduceBinaryClassOnce.Do(func() {
		ReduceBinaryClass = _ReduceBinaryClass{objc.GetClass("MPSNNReduceBinary")}
	})
	return ReduceBinaryClass
}

type _ReduceBinaryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ReduceBinary */
// An interface definition for the [ReduceBinary] class.
type IReduceBinary interface {
	ICNNBinaryKernel
	
/* debug [class_interface_properties]: Properties for ReduceBinary */
	// properties:
	SecondarySourceClipRect() Region get set /* not a class type */
	SetSecondarySourceClipRect(value Region get set /* not a class type */)
	PrimarySourceClipRect() Region get set /* not a class type */
	SetPrimarySourceClipRect(value Region get set /* not a class type */)
	PrimaryOffset() Offset get set /* not a class type */
	SetPrimaryOffset(value Offset get set /* not a class type */)
	SecondaryOffset() Offset get set /* not a class type */
	SetSecondaryOffset(value Offset get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ReduceBinary */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ReduceBinary */
// Alloc allocates a new instance without initialization.
func (rc _ReduceBinaryClass) Alloc() ReduceBinary {
	rv := objc.Send[ReduceBinary](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReduceBinaryClass) New() ReduceBinary {
	rv := objc.Send[ReduceBinary](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReduceBinary) Init() ReduceBinary {
	rv := objc.Send[ReduceBinary](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReduceBinary) Autorelease() ReduceBinary {
	rv := objc.Send[ReduceBinary](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReduceBinary creates a new ReduceBinary instance.
func NewReduceBinary() ReduceBinary {
	return getReduceBinaryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ReduceBinary */
// The base class for binary reduction filters.


// The base class for binary reduction filters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNReduceBinary
type ReduceBinary struct {
	CNNBinaryKernel
}

// ReduceBinaryFrom constructs a [ReduceBinary] from an unsafe.Pointer.
//
// The base class for binary reduction filters.
func ReduceBinaryFrom(ptr unsafe.Pointer) ReduceBinary {
	return ReduceBinary{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ReduceBinary *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ReduceBinary */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ReduceBinary */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ReduceBinary */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ReduceBinary */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942556-secondarysourcecliprect
func (r_ ReduceBinary) SecondarySourceClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("secondarySourceClipRect"))
	return rv
}/* debug [instance_properties/getter]: secondarySourceClipRect */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942556-secondarysourcecliprect
func (r_ ReduceBinary) SetSecondarySourceClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSecondarySourceClipRect:"), value)
}/* debug [instance_properties/setter]: secondarySourceClipRect */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942561-primarysourcecliprect
func (r_ ReduceBinary) PrimarySourceClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("primarySourceClipRect"))
	return rv
}/* debug [instance_properties/getter]: primarySourceClipRect */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942561-primarysourcecliprect
func (r_ ReduceBinary) SetPrimarySourceClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPrimarySourceClipRect:"), value)
}/* debug [instance_properties/setter]: primarySourceClipRect */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/3750640-primaryoffset
func (r_ ReduceBinary) PrimaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("primaryOffset"))
	return rv
}/* debug [instance_properties/getter]: primaryOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/3750640-primaryoffset
func (r_ ReduceBinary) SetPrimaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPrimaryOffset:"), value)
}/* debug [instance_properties/setter]: primaryOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/3750641-secondaryoffset
func (r_ ReduceBinary) SecondaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("secondaryOffset"))
	return rv
}/* debug [instance_properties/getter]: secondaryOffset */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/3750641-secondaryoffset
func (r_ ReduceBinary) SetSecondaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSecondaryOffset:"), value)
}/* debug [instance_properties/setter]: secondaryOffset */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNNReduceBinary */



