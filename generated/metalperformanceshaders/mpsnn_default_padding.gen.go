// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [DefaultPadding] class.
var (
	DefaultPaddingClass     _DefaultPaddingClass
	DefaultPaddingClassOnce sync.Once
)

func getDefaultPaddingClass() _DefaultPaddingClass {
	DefaultPaddingClassOnce.Do(func() {
		DefaultPaddingClass = _DefaultPaddingClass{objc.GetClass("MPSNNDefaultPadding")}
	})
	return DefaultPaddingClass
}

type _DefaultPaddingClass struct {
	class objc.Class
}





// An interface definition for the [DefaultPadding] class.
type IDefaultPadding interface {
	objectivec.IObject
	

	// properties:


	

	// methods:
	Label()


}





// Alloc allocates a new instance without initialization.
func (dc _DefaultPaddingClass) Alloc() DefaultPadding {
	rv := objc.Send[DefaultPadding](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _DefaultPaddingClass) New() DefaultPadding {
	rv := objc.Send[DefaultPadding](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ DefaultPadding) Init() DefaultPadding {
	rv := objc.Send[DefaultPadding](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ DefaultPadding) Autorelease() DefaultPadding {
	rv := objc.Send[DefaultPadding](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewDefaultPadding creates a new DefaultPadding instance.
func NewDefaultPadding() DefaultPadding {
	return getDefaultPaddingClass().New()
}





// A class that provides predefined padding policies for common tasks.


// A class that provides predefined padding policies for common tasks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNNDefaultPadding
type DefaultPadding struct {
	objectivec.Object
}

// DefaultPaddingFrom constructs a [DefaultPadding] from an unsafe.Pointer.
//
// A class that provides predefined padding policies for common tasks.
func DefaultPaddingFrom(ptr unsafe.Pointer) DefaultPadding {
	return DefaultPadding{objectivec.Object{objc.ID(ptr)}}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2867160-paddingwithmethod
func (dc _DefaultPaddingClass) PaddingWithMethod(method PaddingMethod) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("paddingWithMethod:"), method)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2867164-fortensorflowaveragepooling
func (dc _DefaultPaddingClass) ForTensorflowAveragePooling() {
	objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("forTensorflowAveragePooling"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2867164-paddingfortensorflowaveragepooli
func (dc _DefaultPaddingClass) PaddingForTensorflowAveragePooling() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("paddingForTensorflowAveragePooling"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2947962-fortensorflowaveragepoolingvalid
func (dc _DefaultPaddingClass) ForTensorflowAveragePoolingValidOnly() {
	objc.Send[objc.ID](objc.ID(dc.class), objc.Sel("forTensorflowAveragePoolingValidOnly"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2947962-paddingfortensorflowaveragepooli
func (dc _DefaultPaddingClass) PaddingForTensorflowAveragePoolingValidOnly() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(dc.class), objc.Sel("paddingForTensorflowAveragePoolingValidOnly"))
	return rv
}












// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndefaultpadding/2889871-label
func (d_ DefaultPadding) Label() {
	objc.Send[objc.ID](d_.ID, objc.Sel("label"))
}













