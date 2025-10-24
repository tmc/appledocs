// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MPSMatrixRandomMTGP32 */


/* debug [class_header]: Header for MPSMatrixRandomMTGP32 */
// The class instance for the [MatrixRandomMTGP32] class.
var (
	MatrixRandomMTGP32Class     _MatrixRandomMTGP32Class
	MatrixRandomMTGP32ClassOnce sync.Once
)

func getMatrixRandomMTGP32Class() _MatrixRandomMTGP32Class {
	MatrixRandomMTGP32ClassOnce.Do(func() {
		MatrixRandomMTGP32Class = _MatrixRandomMTGP32Class{objc.GetClass("MPSMatrixRandomMTGP32")}
	})
	return MatrixRandomMTGP32Class
}

type _MatrixRandomMTGP32Class struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatrixRandomMTGP32 */
// An interface definition for the [MatrixRandomMTGP32] class.
type IMatrixRandomMTGP32 interface {
	IMatrixRandom
	
/* debug [class_interface_properties]: Properties for MatrixRandomMTGP32 */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatrixRandomMTGP32 */
	// methods:
	SynchronizeState()
	SynchronizeStateOnCommandBuffer(commandBuffer unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatrixRandomMTGP32 */
// Alloc allocates a new instance without initialization.
func (mc _MatrixRandomMTGP32Class) Alloc() MatrixRandomMTGP32 {
	rv := objc.Send[MatrixRandomMTGP32](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatrixRandomMTGP32Class) New() MatrixRandomMTGP32 {
	rv := objc.Send[MatrixRandomMTGP32](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatrixRandomMTGP32) Init() MatrixRandomMTGP32 {
	rv := objc.Send[MatrixRandomMTGP32](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatrixRandomMTGP32) Autorelease() MatrixRandomMTGP32 {
	rv := objc.Send[MatrixRandomMTGP32](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatrixRandomMTGP32 creates a new MatrixRandomMTGP32 instance.
func NewMatrixRandomMTGP32() MatrixRandomMTGP32 {
	return getMatrixRandomMTGP32Class().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatrixRandomMTGP32 */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSMatrixRandomMTGP32
type MatrixRandomMTGP32 struct {
	MatrixRandom
}

// MatrixRandomMTGP32From constructs a [MatrixRandomMTGP32] from an unsafe.Pointer.
func MatrixRandomMTGP32From(ptr unsafe.Pointer) MatrixRandomMTGP32 {
	return MatrixRandomMTGP32{
		MatrixRandom: MatrixRandomFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatrixRandomMTGP32 */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242864-initwithcoder
func NewMatrixRandomMTGP32WithCoderDevice(aDecoder Coder /* not a class type */, device unsafe.Pointer) MatrixRandomMTGP32 {
	instance := getMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[MatrixRandomMTGP32](instance.ID, objc.Sel("initWithCoder:device:"), aDecoder, device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixRandomMTGP32WithCoderDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242865-initwithdevice
func NewMatrixRandomMTGP32WithDevice(device unsafe.Pointer) MatrixRandomMTGP32 {
	instance := getMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[MatrixRandomMTGP32](instance.ID, objc.Sel("initWithDevice:"), device)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixRandomMTGP32WithDevice */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242866-initwithdevice
func NewMatrixRandomMTGP32WithDeviceDestinationDataTypeSeed(device unsafe.Pointer, destinationDataType DataType, seed uint) MatrixRandomMTGP32 {
	instance := getMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[MatrixRandomMTGP32](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:"), device, destinationDataType, seed)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixRandomMTGP32WithDeviceDestinationDataTypeSeed */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242867-initwithdevice
func NewMatrixRandomMTGP32WithDeviceDestinationDataTypeSeedDistributionDescriptor(device unsafe.Pointer, destinationDataType DataType, seed uint, distributionDescriptor IMatrixRandomDistributionDescriptor) MatrixRandomMTGP32 {
	instance := getMatrixRandomMTGP32Class().Alloc()
	rv := objc.Send[MatrixRandomMTGP32](instance.ID, objc.Sel("initWithDevice:destinationDataType:seed:distributionDescriptor:"), device, destinationDataType, seed, distributionDescriptor)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMatrixRandomMTGP32WithDeviceDestinationDataTypeSeedDistributionDescriptor */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatrixRandomMTGP32 */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatrixRandomMTGP32 */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatrixRandomMTGP32 */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242868-synchronizestate
func (m_ MatrixRandomMTGP32) SynchronizeState() {
	objc.Send[objc.ID](m_.ID, objc.Sel("synchronizeState"))
}/* debug [instance_methods/method]: SynchronizeState */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandommtgp32/3242868-synchronizestateoncommandbuffer
func (m_ MatrixRandomMTGP32) SynchronizeStateOnCommandBuffer(commandBuffer unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("synchronizeStateOnCommandBuffer:"), commandBuffer)
}/* debug [instance_methods/method]: SynchronizeStateOnCommandBuffer */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatrixRandomMTGP32 */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MPSMatrixRandomMTGP32 */


