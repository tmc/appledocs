// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CNNLossDataDescriptor] class.
var (
	CNNLossDataDescriptorClass     _CNNLossDataDescriptorClass
	CNNLossDataDescriptorClassOnce sync.Once
)

func getCNNLossDataDescriptorClass() _CNNLossDataDescriptorClass {
	CNNLossDataDescriptorClassOnce.Do(func() {
		CNNLossDataDescriptorClass = _CNNLossDataDescriptorClass{objc.GetClass("MPSCNNLossDataDescriptor")}
	})
	return CNNLossDataDescriptorClass
}

type _CNNLossDataDescriptorClass struct {
	class objc.Class
}





// An interface definition for the [CNNLossDataDescriptor] class.
type ICNNLossDataDescriptor interface {
	objectivec.IObject
	

	// properties:
	Layout() DataLayout get /* not a class type */
	SetLayout(value DataLayout get /* not a class type */)
	BytesPerImage() objectivec.IObject
	SetBytesPerImage(value objectivec.IObject)
	Size() Size get /* not a class type */
	SetSize(value Size get /* not a class type */)
	BytesPerRow() objectivec.IObject
	SetBytesPerRow(value objectivec.IObject)


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CNNLossDataDescriptorClass) Alloc() CNNLossDataDescriptor {
	rv := objc.Send[CNNLossDataDescriptor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNNLossDataDescriptorClass) New() CNNLossDataDescriptor {
	rv := objc.Send[CNNLossDataDescriptor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNNLossDataDescriptor) Init() CNNLossDataDescriptor {
	rv := objc.Send[CNNLossDataDescriptor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNNLossDataDescriptor) Autorelease() CNNLossDataDescriptor {
	rv := objc.Send[CNNLossDataDescriptor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNNLossDataDescriptor creates a new CNNLossDataDescriptor instance.
func NewCNNLossDataDescriptor() CNNLossDataDescriptor {
	return getCNNLossDataDescriptorClass().New()
}





// An object that specifies properties used by a loss data descriptor.


// An object that specifies properties used by a loss data descriptor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCNNLossDataDescriptor
type CNNLossDataDescriptor struct {
	objectivec.Object
}

// CNNLossDataDescriptorFrom constructs a [CNNLossDataDescriptor] from an unsafe.Pointer.
//
// An object that specifies properties used by a loss data descriptor.
func CNNLossDataDescriptorFrom(ptr unsafe.Pointer) CNNLossDataDescriptor {
	return CNNLossDataDescriptor{objectivec.Object{objc.ID(ptr)}}
}










// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951840-cnnlossdatadescriptorwithdata
func (cc _CNNLossDataDescriptorClass) CnnLossDataDescriptorWithDataLayoutSize(data foundation.Data, layout DataLayout, size metal.IMTLSize) ICNNLossDataDescriptor {
	rv := objc.Send[CNNLossDataDescriptor](objc.ID(cc.class), objc.Sel("cnnLossDataDescriptorWithData:layout:size:"), data, layout, size)
	return rv
}

















// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951842-layout
func (c_ CNNLossDataDescriptor) Layout() DataLayout get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("layout"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951842-layout
func (c_ CNNLossDataDescriptor) SetLayout(value DataLayout get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLayout:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951847-bytesperimage
func (c_ CNNLossDataDescriptor) BytesPerImage() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("bytesPerImage"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951847-bytesperimage
func (c_ CNNLossDataDescriptor) SetBytesPerImage(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBytesPerImage:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951848-size
func (c_ CNNLossDataDescriptor) Size() Size get /* not a class type */ {
	rv := objc.Send[objc.ID](c_.ID, objc.Sel("size"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951848-size
func (c_ CNNLossDataDescriptor) SetSize(value Size get /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setSize:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951849-bytesperrow
func (c_ CNNLossDataDescriptor) BytesPerRow() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](c_.ID, objc.Sel("bytesPerRow"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdatadescriptor/2951849-bytesperrow
func (c_ CNNLossDataDescriptor) SetBytesPerRow(value objectivec.IObject) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setBytesPerRow:"), value)
}








