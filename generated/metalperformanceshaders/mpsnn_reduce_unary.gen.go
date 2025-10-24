// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





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





// An interface definition for the [ReduceUnary] class.
type IReduceUnary interface {
	ICNNKernel
	

	// properties:
	ClipRectSource() Region get set /* not a class type */
	SetClipRectSource(value Region get set /* not a class type */)
	Offset() Offset get set /* not a class type */
	SetOffset(value Offset get set /* not a class type */)


	

	// methods:


}





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

























// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/2942547-cliprectsource
func (r_ ReduceUnary) ClipRectSource() Region get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("clipRectSource"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/2942547-cliprectsource
func (r_ ReduceUnary) SetClipRectSource(value Region get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setClipRectSource:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/3750642-offset
func (r_ ReduceUnary) Offset() Offset get set /* not a class type */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("offset"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/3750642-offset
func (r_ ReduceUnary) SetOffset(value Offset get set /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setOffset:"), value)
}








