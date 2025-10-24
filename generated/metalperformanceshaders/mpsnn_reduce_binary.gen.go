// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [ReduceBinary] class.
type IReduceBinary interface {
	ICNNBinaryKernel
	

	// properties:
	SecondarySourceClipRect() Region get set /* not a class type */
	SetSecondarySourceClipRect(value Region get set /* not a class type */)
	PrimarySourceClipRect() Region get set /* not a class type */
	SetPrimarySourceClipRect(value Region get set /* not a class type */)
	PrimaryOffset() Offset get set /* not a class type */
	SetPrimaryOffset(value Offset get set /* not a class type */)
	SecondaryOffset() Offset get set /* not a class type */
	SetSecondaryOffset(value Offset get set /* not a class type */)


	

	// methods:


}





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

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942556-secondarysourcecliprect
func (r_ ReduceBinary) SecondarySourceClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("secondarySourceClipRect"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942556-secondarysourcecliprect
func (r_ ReduceBinary) SetSecondarySourceClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSecondarySourceClipRect:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942561-primarysourcecliprect
func (r_ ReduceBinary) PrimarySourceClipRect() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("primarySourceClipRect"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942561-primarysourcecliprect
func (r_ ReduceBinary) SetPrimarySourceClipRect(value Region get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPrimarySourceClipRect:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/3750640-primaryoffset
func (r_ ReduceBinary) PrimaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("primaryOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/3750640-primaryoffset
func (r_ ReduceBinary) SetPrimaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setPrimaryOffset:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/3750641-secondaryoffset
func (r_ ReduceBinary) SecondaryOffset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("secondaryOffset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/3750641-secondaryoffset
func (r_ ReduceBinary) SetSecondaryOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setSecondaryOffset:"), value)
}








