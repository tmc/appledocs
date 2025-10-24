// Code generated from Apple documentation for MLCompute. DO NOT EDIT.

package mlcompute

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MLCTensorOptimizerDeviceData */


/* debug [class_header]: Header for MLCTensorOptimizerDeviceData */
// The class instance for the [CTensorOptimizerDeviceData] class.
var (
	CTensorOptimizerDeviceDataClass     _CTensorOptimizerDeviceDataClass
	CTensorOptimizerDeviceDataClassOnce sync.Once
)

func getCTensorOptimizerDeviceDataClass() _CTensorOptimizerDeviceDataClass {
	CTensorOptimizerDeviceDataClassOnce.Do(func() {
		CTensorOptimizerDeviceDataClass = _CTensorOptimizerDeviceDataClass{objc.GetClass("MLCTensorOptimizerDeviceData")}
	})
	return CTensorOptimizerDeviceDataClass
}

type _CTensorOptimizerDeviceDataClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CTensorOptimizerDeviceData */
// An interface definition for the [CTensorOptimizerDeviceData] class.
type ICTensorOptimizerDeviceData interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CTensorOptimizerDeviceData */
	// properties:
	Data() foundation.Data
	SetData(value foundation.Data)
	Descriptor() IMLCTensorDescriptor
	SetDescriptor(value IMLCTensorDescriptor)
	Device() IMLCDevice
	SetDevice(value IMLCDevice)
	HasValidNumerics() bool
	SetHasValidNumerics(value bool)
	Label() objc.IObject /* cross-framework: NSString */
	SetLabel(value objc.IObject /* cross-framework: NSString */)
	OptimizerData() IMLCTensorData
	SetOptimizerData(value IMLCTensorData)
	OptimizerDeviceData() IMLCTensorOptimizerDeviceData
	SetOptimizerDeviceData(value IMLCTensorOptimizerDeviceData)
	TensorID() int
	SetTensorID(value int)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CTensorOptimizerDeviceData */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CTensorOptimizerDeviceData */
// Alloc allocates a new instance without initialization.
func (cc _CTensorOptimizerDeviceDataClass) Alloc() CTensorOptimizerDeviceData {
	rv := objc.Send[CTensorOptimizerDeviceData](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CTensorOptimizerDeviceDataClass) New() CTensorOptimizerDeviceData {
	rv := objc.Send[CTensorOptimizerDeviceData](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CTensorOptimizerDeviceData) Init() CTensorOptimizerDeviceData {
	rv := objc.Send[CTensorOptimizerDeviceData](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CTensorOptimizerDeviceData) Autorelease() CTensorOptimizerDeviceData {
	rv := objc.Send[CTensorOptimizerDeviceData](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCTensorOptimizerDeviceData creates a new CTensorOptimizerDeviceData instance.
func NewCTensorOptimizerDeviceData() CTensorOptimizerDeviceData {
	return getCTensorOptimizerDeviceDataClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CTensorOptimizerDeviceData */
// An encapsulation of the device memory associated with a tensor that an optimizer uses.


// An encapsulation of the device memory associated with a tensor that an optimizer uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MLCompute/MLCTensorOptimizerDeviceData
type CTensorOptimizerDeviceData struct {
	objectivec.Object
}

// CTensorOptimizerDeviceDataFrom constructs a [CTensorOptimizerDeviceData] from an unsafe.Pointer.
//
// An encapsulation of the device memory associated with a tensor that an optimizer uses.
func CTensorOptimizerDeviceDataFrom(ptr unsafe.Pointer) CTensorOptimizerDeviceData {
	return CTensorOptimizerDeviceData{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CTensorOptimizerDeviceData *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CTensorOptimizerDeviceData */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CTensorOptimizerDeviceData */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CTensorOptimizerDeviceData */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CTensorOptimizerDeviceData */

// The tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/data
func (c_ CTensorOptimizerDeviceData) Data() foundation.Data {
	rv := objc.Send[foundation.Data](c_.ID, objc.Sel("data"))
	return rv
}/* debug [instance_properties/getter]: data */


// The tensor data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/data
func (c_ CTensorOptimizerDeviceData) SetData(value foundation.Data) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setData:"), value)
}/* debug [instance_properties/setter]: data */


// The configuration object you use to create a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/descriptor
func (c_ CTensorOptimizerDeviceData) Descriptor() IMLCTensorDescriptor {
	rv := objc.Send[CTensorDescriptor](c_.ID, objc.Sel("descriptor"))
	return rv
}/* debug [instance_properties/getter]: descriptor */


// The configuration object you use to create a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/descriptor
func (c_ CTensorOptimizerDeviceData) SetDescriptor(value IMLCTensorDescriptor) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDescriptor:"), value)
}/* debug [instance_properties/setter]: descriptor */


// The device associated with this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/device
func (c_ CTensorOptimizerDeviceData) Device() IMLCDevice {
	rv := objc.Send[CDevice](c_.ID, objc.Sel("device"))
	return rv
}/* debug [instance_properties/getter]: device */


// The device associated with this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/device
func (c_ CTensorOptimizerDeviceData) SetDevice(value IMLCDevice) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDevice:"), value)
}/* debug [instance_properties/setter]: device */


// A Boolean that indicates whether a tensor contains NaN or INF values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/hasvalidnumerics
func (c_ CTensorOptimizerDeviceData) HasValidNumerics() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasValidNumerics"))
	return rv
}/* debug [instance_properties/getter]: hasValidNumerics */


// A Boolean that indicates whether a tensor contains NaN or INF values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/hasvalidnumerics
func (c_ CTensorOptimizerDeviceData) SetHasValidNumerics(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setHasValidNumerics:"), value)
}/* debug [instance_properties/setter]: hasValidNumerics */


// A string that identifes this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/label
func (c_ CTensorOptimizerDeviceData) Label() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](c_.ID, objc.Sel("label"))
	return rv
}/* debug [instance_properties/getter]: label */


// A string that identifes this tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/label
func (c_ CTensorOptimizerDeviceData) SetLabel(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setLabel:"), value)
}/* debug [instance_properties/setter]: label */


// An array that contains optimizer buffers you specify when you create a tensor parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/optimizerdata
func (c_ CTensorOptimizerDeviceData) OptimizerData() IMLCTensorData {
	rv := objc.Send[CTensorData](c_.ID, objc.Sel("optimizerData"))
	return rv
}/* debug [instance_properties/getter]: optimizerData */


// An array that contains optimizer buffers you specify when you create a tensor parameter.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/optimizerdata
func (c_ CTensorOptimizerDeviceData) SetOptimizerData(value IMLCTensorData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptimizerData:"), value)
}/* debug [instance_properties/setter]: optimizerData */


// An array that contains the device optimizer buffers you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/optimizerdevicedata
func (c_ CTensorOptimizerDeviceData) OptimizerDeviceData() IMLCTensorOptimizerDeviceData {
	rv := objc.Send[CTensorOptimizerDeviceData](c_.ID, objc.Sel("optimizerDeviceData"))
	return rv
}/* debug [instance_properties/getter]: optimizerDeviceData */


// An array that contains the device optimizer buffers you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/optimizerdevicedata
func (c_ CTensorOptimizerDeviceData) SetOptimizerDeviceData(value IMLCTensorOptimizerDeviceData) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOptimizerDeviceData:"), value)
}/* debug [instance_properties/setter]: optimizerDeviceData */


// A number that uniquely identifies the tensor, which the framework assigns when it creates a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/tensorid
func (c_ CTensorOptimizerDeviceData) TensorID() int {
	rv := objc.Send[int](c_.ID, objc.Sel("tensorID"))
	return rv
}/* debug [instance_properties/getter]: tensorID */


// A number that uniquely identifies the tensor, which the framework assigns when it creates a tensor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/mlcompute/mlctensor/tensorid
func (c_ CTensorOptimizerDeviceData) SetTensorID(value int) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setTensorID:"), value)
}/* debug [instance_properties/setter]: tensorID */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MLCTensorOptimizerDeviceData */



