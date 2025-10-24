// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRDeviceControllerFactory */


/* debug [class_header]: Header for MTRDeviceControllerFactory */
// The class instance for the [MTRDeviceControllerFactory] class.
var (
	MTRDeviceControllerFactoryClass     _MTRDeviceControllerFactoryClass
	MTRDeviceControllerFactoryClassOnce sync.Once
)

func getMTRDeviceControllerFactoryClass() _MTRDeviceControllerFactoryClass {
	MTRDeviceControllerFactoryClassOnce.Do(func() {
		MTRDeviceControllerFactoryClass = _MTRDeviceControllerFactoryClass{objc.GetClass("MTRDeviceControllerFactory")}
	})
	return MTRDeviceControllerFactoryClass
}

type _MTRDeviceControllerFactoryClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRDeviceControllerFactory */
// An interface definition for the [MTRDeviceControllerFactory] class.
type IMTRDeviceControllerFactory interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRDeviceControllerFactory */
	// properties:
	Running() bool
	KnownFabrics() []MTRFabricInfo
	IsRunning() bool
	SetIsRunning(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRDeviceControllerFactory */
	// methods:
	CreateControllerOnExistingFabricError(startupParams IMTRDeviceControllerStartupParams, error_ unsafe.Pointer) IMTRDeviceController
	CreateControllerOnNewFabricError(startupParams IMTRDeviceControllerStartupParams, error_ unsafe.Pointer) IMTRDeviceController
	PreWarmCommissioningSession()
	StartControllerFactoryError(startupParams IMTRDeviceControllerFactoryParams, error_ unsafe.Pointer) bool
	StopControllerFactory()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRDeviceControllerFactory */
// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerFactoryClass) Alloc() MTRDeviceControllerFactory {
	rv := objc.Send[MTRDeviceControllerFactory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRDeviceControllerFactoryClass) New() MTRDeviceControllerFactory {
	rv := objc.Send[MTRDeviceControllerFactory](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRDeviceControllerFactory) Init() MTRDeviceControllerFactory {
	rv := objc.Send[MTRDeviceControllerFactory](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRDeviceControllerFactory) Autorelease() MTRDeviceControllerFactory {
	rv := objc.Send[MTRDeviceControllerFactory](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRDeviceControllerFactory creates a new MTRDeviceControllerFactory instance.
func NewMTRDeviceControllerFactory() MTRDeviceControllerFactory {
	return getMTRDeviceControllerFactoryClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRDeviceControllerFactory */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory
type MTRDeviceControllerFactory struct {
	objectivec.Object
}

// MTRDeviceControllerFactoryFrom constructs a [MTRDeviceControllerFactory] from an unsafe.Pointer.
func MTRDeviceControllerFactoryFrom(ptr unsafe.Pointer) MTRDeviceControllerFactory {
	return MTRDeviceControllerFactory{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRDeviceControllerFactory *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRDeviceControllerFactory */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/sharedInstance()
func (mc _MTRDeviceControllerFactoryClass) SharedInstance() MTRDeviceControllerFactory {
	rv := objc.Send[MTRDeviceControllerFactory](objc.ID(mc.class), objc.Sel("sharedInstance"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedInstance) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRDeviceControllerFactory */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRDeviceControllerFactory */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/createController(onExistingFabric:)
func (m_ MTRDeviceControllerFactory) CreateControllerOnExistingFabricError(startupParams IMTRDeviceControllerStartupParams, error_ unsafe.Pointer) IMTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("createControllerOnExistingFabric:error:"), startupParams, error_)
	return rv
}/* debug [instance_methods/method]: CreateControllerOnExistingFabricError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/createController(onNewFabric:)
func (m_ MTRDeviceControllerFactory) CreateControllerOnNewFabricError(startupParams IMTRDeviceControllerStartupParams, error_ unsafe.Pointer) IMTRDeviceController {
	rv := objc.Send[MTRDeviceController](m_.ID, objc.Sel("createControllerOnNewFabric:error:"), startupParams, error_)
	return rv
}/* debug [instance_methods/method]: CreateControllerOnNewFabricError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/preWarmCommissioningSession()
func (m_ MTRDeviceControllerFactory) PreWarmCommissioningSession() {
	objc.Send[objc.ID](m_.ID, objc.Sel("preWarmCommissioningSession"))
}/* debug [instance_methods/method]: PreWarmCommissioningSession */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/start(_:)
func (m_ MTRDeviceControllerFactory) StartControllerFactoryError(startupParams IMTRDeviceControllerFactoryParams, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("startControllerFactory:error:"), startupParams, error_)
	return rv
}/* debug [instance_methods/method]: StartControllerFactoryError */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/stop()
func (m_ MTRDeviceControllerFactory) StopControllerFactory() {
	objc.Send[objc.ID](m_.ID, objc.Sel("stopControllerFactory"))
}/* debug [instance_methods/method]: StopControllerFactory */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRDeviceControllerFactory */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/isRunning
func (m_ MTRDeviceControllerFactory) Running() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("running"))
	return rv
}/* debug [instance_properties/getter]: running */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/knownFabrics
func (m_ MTRDeviceControllerFactory) KnownFabrics() []MTRFabricInfo {
	rv := objc.Send[[]MTRFabricInfo](m_.ID, objc.Sel("knownFabrics"))
	return rv
}/* debug [instance_properties/getter]: knownFabrics */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerfactory/isrunning
func (m_ MTRDeviceControllerFactory) IsRunning() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("isRunning"))
	return rv
}/* debug [instance_properties/getter]: isRunning */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/matter/mtrdevicecontrollerfactory/isrunning
func (m_ MTRDeviceControllerFactory) SetIsRunning(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setIsRunning:"), value)
}/* debug [instance_properties/setter]: isRunning */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRDeviceControllerFactory */



