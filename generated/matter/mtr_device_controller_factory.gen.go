// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [MTRDeviceControllerFactory] class.
type IMTRDeviceControllerFactory interface {
	objectivec.IObject
	CreateControllerOnExistingFabricError(startupParams unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	CreateControllerOnNewFabricError(startupParams unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer
	StartControllerFactoryError(startupParams unsafe.Pointer, error_ unsafe.Pointer) bool
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory
type MTRDeviceControllerFactory struct {
	objectivec.Object
}

// MTRDeviceControllerFactoryFrom constructs a [MTRDeviceControllerFactory] from an unsafe.Pointer.
func MTRDeviceControllerFactoryFrom(ptr unsafe.Pointer) MTRDeviceControllerFactory {
	return MTRDeviceControllerFactory{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MTRDeviceControllerFactoryClass) Alloc() MTRDeviceControllerFactory {
	rv := objc.Send[MTRDeviceControllerFactory](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/createController(onExistingFabric:)
func (m_ MTRDeviceControllerFactory) CreateControllerOnExistingFabricError(startupParams unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("createControllerOnExistingFabric:error:"), startupParams, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/createController(onNewFabric:)
func (m_ MTRDeviceControllerFactory) CreateControllerOnNewFabricError(startupParams unsafe.Pointer, error_ unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("createControllerOnNewFabric:error:"), startupParams, error_)
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRDeviceControllerFactory/start(_:)
func (m_ MTRDeviceControllerFactory) StartControllerFactoryError(startupParams unsafe.Pointer, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("startControllerFactory:error:"), startupParams, error_)
	return rv
}



