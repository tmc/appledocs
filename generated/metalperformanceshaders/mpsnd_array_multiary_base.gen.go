// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSNDArrayMultiaryBase */


/* debug [class_header]: Header for MPSNDArrayMultiaryBase */
// The class instance for the [NDArrayMultiaryBase] class.
var (
	NDArrayMultiaryBaseClass     _NDArrayMultiaryBaseClass
	NDArrayMultiaryBaseClassOnce sync.Once
)

func getNDArrayMultiaryBaseClass() _NDArrayMultiaryBaseClass {
	NDArrayMultiaryBaseClassOnce.Do(func() {
		NDArrayMultiaryBaseClass = _NDArrayMultiaryBaseClass{objc.GetClass("MPSNDArrayMultiaryBase")}
	})
	return NDArrayMultiaryBaseClass
}

type _NDArrayMultiaryBaseClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NDArrayMultiaryBase */
// An interface definition for the [NDArrayMultiaryBase] class.
type INDArrayMultiaryBase interface {
	IKernel
	
/* debug [class_interface_properties]: Properties for NDArrayMultiaryBase */
	// properties:
	DestinationArrayAllocator() NDArrayAllocator get set /* not a class type */
	SetDestinationArrayAllocator(value NDArrayAllocator get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NDArrayMultiaryBase */
	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
	DestinationArrayDescriptor()
	DestinationArrayDescriptorForSourceArraysSourceState(sources unsafe.Pointer, state IState) INDArrayDescriptor
	Encode()
	ResultState()
	ResultStateForSourceArraysSourceStatesDestinationArray(sourceArrays unsafe.Pointer, sourceStates unsafe.Pointer, destinationArray INDArray) IState
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NDArrayMultiaryBase */
// Alloc allocates a new instance without initialization.
func (nc _NDArrayMultiaryBaseClass) Alloc() NDArrayMultiaryBase {
	rv := objc.Send[NDArrayMultiaryBase](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NDArrayMultiaryBaseClass) New() NDArrayMultiaryBase {
	rv := objc.Send[NDArrayMultiaryBase](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NDArrayMultiaryBase) Init() NDArrayMultiaryBase {
	rv := objc.Send[NDArrayMultiaryBase](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NDArrayMultiaryBase) Autorelease() NDArrayMultiaryBase {
	rv := objc.Send[NDArrayMultiaryBase](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNDArrayMultiaryBase creates a new NDArrayMultiaryBase instance.
func NewNDArrayMultiaryBase() NDArrayMultiaryBase {
	return getNDArrayMultiaryBaseClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NDArrayMultiaryBase */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayMultiaryBase
type NDArrayMultiaryBase struct {
	Kernel
}

// NDArrayMultiaryBaseFrom constructs a [NDArrayMultiaryBase] from an unsafe.Pointer.
func NDArrayMultiaryBaseFrom(ptr unsafe.Pointer) NDArrayMultiaryBase {
	return NDArrayMultiaryBase{
		Kernel: KernelFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NDArrayMultiaryBase */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131740-initwithcoder
func NewNDArrayMultiaryBaseWithCoderDevice(coder foundation.Coder, device unsafe.Pointer) NDArrayMultiaryBase {
	instance := getNDArrayMultiaryBaseClass().Alloc()
	rv := objc.Send[NDArrayMultiaryBase](instance.ID, objc.Sel("initWithCoder:device:"), coder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayMultiaryBaseWithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131741-initwithdevice
func NewNDArrayMultiaryBaseWithDeviceSourceCount(device unsafe.Pointer, count uint) NDArrayMultiaryBase {
	instance := getNDArrayMultiaryBaseClass().Alloc()
	rv := objc.Send[NDArrayMultiaryBase](instance.ID, objc.Sel("initWithDevice:sourceCount:"), device, count)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNDArrayMultiaryBaseWithDeviceSourceCount */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NDArrayMultiaryBase */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NDArrayMultiaryBase */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NDArrayMultiaryBase */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone
func (n_ NDArrayMultiaryBase) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131736-destinationarraydescriptor
func (n_ NDArrayMultiaryBase) DestinationArrayDescriptor() {
	objc.Send[objc.ID](n_.ID, objc.Sel("destinationArrayDescriptor"))
}/* debug [instance_methods/method]: DestinationArrayDescriptor */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131736-destinationarraydescriptorforsou
func (n_ NDArrayMultiaryBase) DestinationArrayDescriptorForSourceArraysSourceState(sources unsafe.Pointer, state IState) INDArrayDescriptor {
	rv := objc.Send[NDArrayDescriptor](n_.ID, objc.Sel("destinationArrayDescriptorForSourceArrays:sourceState:"), sources, state)
	return rv
}/* debug [instance_methods/method]: DestinationArrayDescriptorForSourceArraysSourceState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131739-encode
func (n_ NDArrayMultiaryBase) Encode() {
	objc.Send[objc.ID](n_.ID, objc.Sel("encode"))
}/* debug [instance_methods/method]: Encode */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131739-encodewithcoder
func (n_ NDArrayMultiaryBase) EncodeWithCoder(coder foundation.Coder) {
	objc.Send[objc.ID](n_.ID, objc.Sel("encodeWithCoder:"), coder)
}/* debug [instance_methods/method]: EncodeWithCoder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3143521-resultstate
func (n_ NDArrayMultiaryBase) ResultState() {
	objc.Send[objc.ID](n_.ID, objc.Sel("resultState"))
}/* debug [instance_methods/method]: ResultState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3143521-resultstateforsourcearrays
func (n_ NDArrayMultiaryBase) ResultStateForSourceArraysSourceStatesDestinationArray(sourceArrays unsafe.Pointer, sourceStates unsafe.Pointer, destinationArray INDArray) IState {
	rv := objc.Send[State](n_.ID, objc.Sel("resultStateForSourceArrays:sourceStates:destinationArray:"), sourceArrays, sourceStates, destinationArray)
	return rv
}/* debug [instance_methods/method]: ResultStateForSourceArraysSourceStatesDestinationArray */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NDArrayMultiaryBase */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131735-destinationarrayallocator
func (n_ NDArrayMultiaryBase) DestinationArrayAllocator() NDArrayAllocator get set /* not a class type */ {
	rv := objc.Send[objc.ID](n_.ID, objc.Sel("destinationArrayAllocator"))
	return rv
}/* debug [instance_properties/getter]: destinationArrayAllocator */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131735-destinationarrayallocator
func (n_ NDArrayMultiaryBase) SetDestinationArrayAllocator(value NDArrayAllocator get set /* not a class type */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setDestinationArrayAllocator:"), value)
}/* debug [instance_properties/setter]: destinationArrayAllocator */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSNDArrayMultiaryBase */


