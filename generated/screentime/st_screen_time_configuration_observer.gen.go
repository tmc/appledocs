// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class STScreenTimeConfigurationObserver */


/* debug [class_header]: Header for STScreenTimeConfigurationObserver */
// The class instance for the [STScreenTimeConfigurationObserver] class.
var (
	STScreenTimeConfigurationObserverClass     _STScreenTimeConfigurationObserverClass
	STScreenTimeConfigurationObserverClassOnce sync.Once
)

func getSTScreenTimeConfigurationObserverClass() _STScreenTimeConfigurationObserverClass {
	STScreenTimeConfigurationObserverClassOnce.Do(func() {
		STScreenTimeConfigurationObserverClass = _STScreenTimeConfigurationObserverClass{objc.GetClass("STScreenTimeConfigurationObserver")}
	})
	return STScreenTimeConfigurationObserverClass
}

type _STScreenTimeConfigurationObserverClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for STScreenTimeConfigurationObserver */
// An interface definition for the [STScreenTimeConfigurationObserver] class.
type ISTScreenTimeConfigurationObserver interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for STScreenTimeConfigurationObserver */
	// properties:
	Configuration() ISTScreenTimeConfiguration
	EnforcesChildRestrictions() bool
	SetEnforcesChildRestrictions(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for STScreenTimeConfigurationObserver */
	// methods:
	StartObserving()
	StopObserving()
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for STScreenTimeConfigurationObserver */
// Alloc allocates a new instance without initialization.
func (sc _STScreenTimeConfigurationObserverClass) Alloc() STScreenTimeConfigurationObserver {
	rv := objc.Send[STScreenTimeConfigurationObserver](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _STScreenTimeConfigurationObserverClass) New() STScreenTimeConfigurationObserver {
	rv := objc.Send[STScreenTimeConfigurationObserver](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ STScreenTimeConfigurationObserver) Init() STScreenTimeConfigurationObserver {
	rv := objc.Send[STScreenTimeConfigurationObserver](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ STScreenTimeConfigurationObserver) Autorelease() STScreenTimeConfigurationObserver {
	rv := objc.Send[STScreenTimeConfigurationObserver](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSTScreenTimeConfigurationObserver creates a new STScreenTimeConfigurationObserver instance.
func NewSTScreenTimeConfigurationObserver() STScreenTimeConfigurationObserver {
	return getSTScreenTimeConfigurationObserverClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for STScreenTimeConfigurationObserver */
// The object you use to observe changes to the current configuration.
//
// Use this class to start and stop observing the current configuration. For example, you can opt to disable private browsing in your web browser’s view controller when is .


// The object you use to observe changes to the current configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfigurationObserver
type STScreenTimeConfigurationObserver struct {
	objectivec.Object
}

// STScreenTimeConfigurationObserverFrom constructs a [STScreenTimeConfigurationObserver] from an unsafe.Pointer.
//
// The object you use to observe changes to the current configuration.
func STScreenTimeConfigurationObserverFrom(ptr unsafe.Pointer) STScreenTimeConfigurationObserver {
	return STScreenTimeConfigurationObserver{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for STScreenTimeConfigurationObserver */

// Creates a configuration observer that reports updates on the queue you specify.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfigurationObserver/init(updateQueue:)
func NewSTScreenTimeConfigurationObserverWithUpdateQueue(updateQueue unsafe.Pointer) STScreenTimeConfigurationObserver {
	instance := getSTScreenTimeConfigurationObserverClass().Alloc()
	rv := objc.Send[STScreenTimeConfigurationObserver](instance.ID, objc.Sel("initWithUpdateQueue:"), updateQueue)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSTScreenTimeConfigurationObserverWithUpdateQueue */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for STScreenTimeConfigurationObserver */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for STScreenTimeConfigurationObserver */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for STScreenTimeConfigurationObserver */

// Starts observing changes to the current configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfigurationObserver/startObserving()
func (s_ STScreenTimeConfigurationObserver) StartObserving() {
	objc.Send[objc.ID](s_.ID, objc.Sel("startObserving"))
}/* debug [instance_methods/method]: StartObserving */


// Stops observing changes to the current configuration.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfigurationObserver/stopObserving()
func (s_ STScreenTimeConfigurationObserver) StopObserving() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stopObserving"))
}/* debug [instance_methods/method]: StopObserving */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for STScreenTimeConfigurationObserver */

// The configuration being observed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfigurationObserver/configuration
func (s_ STScreenTimeConfigurationObserver) Configuration() ISTScreenTimeConfiguration {
	rv := objc.Send[STScreenTimeConfiguration](s_.ID, objc.Sel("configuration"))
	return rv
}/* debug [instance_properties/getter]: configuration */


// A Boolean that indicates whether the device is currently enforcing child
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stscreentimeconfiguration/enforceschildrestrictions
func (s_ STScreenTimeConfigurationObserver) EnforcesChildRestrictions() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("enforcesChildRestrictions"))
	return rv
}/* debug [instance_properties/getter]: enforcesChildRestrictions */


// A Boolean that indicates whether the device is currently enforcing child
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/screentime/stscreentimeconfiguration/enforceschildrestrictions
func (s_ STScreenTimeConfigurationObserver) SetEnforcesChildRestrictions(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setEnforcesChildRestrictions:"), value)
}/* debug [instance_properties/setter]: enforcesChildRestrictions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class STScreenTimeConfigurationObserver */


