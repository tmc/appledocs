// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [MonitorConfiguration] class.
var (
	MonitorConfigurationClass     _MonitorConfigurationClass
	MonitorConfigurationClassOnce sync.Once
)

func getMonitorConfigurationClass() _MonitorConfigurationClass {
	MonitorConfigurationClassOnce.Do(func() {
		MonitorConfigurationClass = _MonitorConfigurationClass{objc.GetClass("CLMonitorConfiguration")}
	})
	return MonitorConfigurationClass
}

type _MonitorConfigurationClass struct {
	class objc.Class
}





// An interface definition for the [MonitorConfiguration] class.
type IMonitorConfiguration interface {
	objectivec.IObject
	

	// properties:
	EventHandler() unsafe.Pointer
	Name() foundation.foundation.INSString
	Queue() objectivec.IObject


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (mc _MonitorConfigurationClass) Alloc() MonitorConfiguration {
	rv := objc.Send[MonitorConfiguration](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MonitorConfigurationClass) New() MonitorConfiguration {
	rv := objc.Send[MonitorConfiguration](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MonitorConfiguration) Init() MonitorConfiguration {
	rv := objc.Send[MonitorConfiguration](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MonitorConfiguration) Autorelease() MonitorConfiguration {
	rv := objc.Send[MonitorConfiguration](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMonitorConfiguration creates a new MonitorConfiguration instance.
func NewMonitorConfiguration() MonitorConfiguration {
	return getMonitorConfigurationClass().New()
}





// An object for configuring a location monitor instance.


// An object for configuring a location monitor instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration
type MonitorConfiguration struct {
	objectivec.Object
}

// MonitorConfigurationFrom constructs a [MonitorConfiguration] from an unsafe.Pointer.
//
// An object for configuring a location monitor instance.
func MonitorConfigurationFrom(ptr unsafe.Pointer) MonitorConfiguration {
	return MonitorConfiguration{objectivec.Object{objc.ID(ptr)}}
}










// Creates a location monitor instance with the name, dispatch queue, and event handler you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration/configWithMonitorName:queue:eventHandler:
func (mc _MonitorConfigurationClass) ConfigWithMonitorNameQueueEventHandler(name foundation.foundation.INSString, queue objectivec.IObject, eventHandler unsafe.Pointer) IMonitorConfiguration {
	rv := objc.Send[MonitorConfiguration](objc.ID(mc.class), objc.Sel("configWithMonitorName:queue:eventHandler:"), name, queue, eventHandler)
	return rv
}

















// The block the framework calls as the event handler for the location monitor instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration/eventHandler
func (m_ MonitorConfiguration) EventHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("eventHandler"))
	return rv
}


// The name of the monitor instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration/name
func (m_ MonitorConfiguration) Name() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("name"))
	return rv
}


// The dispatch queue to bind the instance of a location monitor to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreLocation/CLMonitorConfiguration/queue
func (m_ MonitorConfiguration) Queue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("queue"))
	return rv
}








