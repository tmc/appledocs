// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSKernel */


/* debug [class_header]: Header for MPSKernel */
// The class instance for the [Kernel] class.
var (
	KernelClass     _KernelClass
	KernelClassOnce sync.Once
)

func getKernelClass() _KernelClass {
	KernelClassOnce.Do(func() {
		KernelClass = _KernelClass{objc.GetClass("MPSKernel")}
	})
	return KernelClass
}

type _KernelClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Kernel */
// An interface definition for the [Kernel] class.
type IKernel interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Kernel */
	// properties:
	Label() objectivec.IObject
	SetLabel(value objectivec.IObject)
	Device() Device get /* not a class type */
	SetDevice(value Device get /* not a class type */)
	Options() KernelOptions get set /* not a class type */
	SetOptions(value KernelOptions get set /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Kernel */
	// methods:
	CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Kernel */
// Alloc allocates a new instance without initialization.
func (kc _KernelClass) Alloc() Kernel {
	rv := objc.Send[Kernel](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (kc _KernelClass) New() Kernel {
	rv := objc.Send[Kernel](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ Kernel) Init() Kernel {
	rv := objc.Send[Kernel](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ Kernel) Autorelease() Kernel {
	rv := objc.Send[Kernel](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKernel creates a new Kernel instance.
func NewKernel() Kernel {
	return getKernelClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Kernel */
// A standard interface for Metal Performance Shaders kernels.
//
// You should not use the class directly. Instead, a number of subclasses are available that define specific high-performance data-parallel operations. The basic sequence for applying a kernel to an image is as follows: Initialize a kernel corresponding to the operation you wish to perform: Encode the kernel into a command buffer. Encoding the kernel merely encodes the operation into a command buffer. It does not modify any pixels, yet. All kernel state has been copied to the command buffer. Kernels may be reused. If the texture was previously operated on by another command encoder (e.g. a render command encoder), you should call the method on the other encoder before encoding the filter. Some kernels work in place, even in situations where Metal might not normally allow in-place operation on textures. If in-place operation is desired, you may attempt to call the method. If the operation cannot be completed in place, then will be returned and you will have to create a new result texture and try again. To make an in-place image filter reliable, pass a fallback block to the method to create a new texture to write to in the event that a filter cannot operate in place. You may repeat step 2 to encode more kernels, as desired. 3. After encoding any additional work to the command buffer using other encoders, submit the command buffer to your command queue, using:


// A standard interface for Metal Performance Shaders kernels.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSKernel
type Kernel struct {
	objectivec.Object
}

// KernelFrom constructs a [Kernel] from an unsafe.Pointer.
//
// A standard interface for Metal Performance Shaders kernels.
func KernelFrom(ptr unsafe.Pointer) Kernel {
	return Kernel{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Kernel */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/2875161-initwithcoder
func NewKernelWithCoder(aDecoder Coder /* not a class type */) Kernel {
	instance := getKernelClass().Alloc()
	rv := objc.Send[Kernel](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewKernelWithCoder */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/2867190-initwithcoder
func NewKernelWithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) Kernel {
	instance := getKernelClass().Alloc()
	rv := objc.Send[Kernel](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewKernelWithCoderDevice */


// Initializes a new kernel object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice
func NewKernelWithDevice(device unsafe.Pointer) Kernel {
	instance := getKernelClass().Alloc()
	rv := objc.Send[Kernel](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewKernelWithDevice */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Kernel */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Kernel */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Kernel */

// Makes a copy of this kernel object for a new device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone
func (k_ Kernel) CopyWithZoneDevice(zone Zone /* not a class type */, device unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](k_.ID, objc.Sel("copyWithZone:device:"), zone, device)
	return rv
}/* debug [instance_methods/method]: CopyWithZoneDevice */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Kernel */

// The string that identifies the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618803-label
func (k_ Kernel) Label() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](k_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// The string that identifies the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618803-label
func (k_ Kernel) SetLabel(value objectivec.IObject) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// The device on which the kernel will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618824-device
func (k_ Kernel) Device() Device get /* not a class type */ {
	rv := objc.Send[objc.ID](k_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The device on which the kernel will be used.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618824-device
func (k_ Kernel) SetDevice(value Device get /* not a class type */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// The set of options used to run the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618889-options
func (k_ Kernel) Options() KernelOptions get set /* not a class type */ {
	rv := objc.Send[objc.ID](k_.ID, objc.Sel("options"))
	return rv
}/* debug [instance_properties/getter]: options */


// The set of options used to run the kernel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618889-options
func (k_ Kernel) SetOptions(value KernelOptions get set /* not a class type */) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setOptions:"), value)
}/* debug [instance_properties/setter]: options */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSKernel */


