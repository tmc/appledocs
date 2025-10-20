// Code generated from Apple documentation for ScreenTime. DO NOT EDIT.

package screentime

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [STScreenTimeConfigurationObserver] class.
type ISTScreenTimeConfigurationObserver interface {
	objectivec.IObject
	StartObserving()
	StopObserving()
}

// The object you use to observe changes to the current configuration.
//
// Use this class to start and stop observing the current configuration. For example, you can opt to disable private browsing in your web browser’s view controller when is .
//
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

// Alloc allocates a new instance without initialization.
func (sc _STScreenTimeConfigurationObserverClass) Alloc() STScreenTimeConfigurationObserver {
	rv := objc.Send[STScreenTimeConfigurationObserver](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Creates a configuration observer that reports updates on the queue you specify.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfigurationObserver/init(updateQueue:)
func NewSTScreenTimeConfigurationObserverWithUpdateQueue(updateQueue unsafe.Pointer) STScreenTimeConfigurationObserver {
	instance := getSTScreenTimeConfigurationObserverClass().Alloc()
	rv := objc.Send[STScreenTimeConfigurationObserver](instance.ID, objc.Sel("initWithUpdateQueue:"), updateQueue)
	rv.Autorelease()
	return rv
}


// Starts observing changes to the current configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfigurationObserver/startObserving()
func (s_ STScreenTimeConfigurationObserver) StartObserving() {
	objc.Send[objc.ID](s_.ID, objc.Sel("startObserving"))
}

// Stops observing changes to the current configuration.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfigurationObserver/stopObserving()
func (s_ STScreenTimeConfigurationObserver) StopObserving() {
	objc.Send[objc.ID](s_.ID, objc.Sel("stopObserving"))
}

// The configuration being observed.
//
// [Full Topic]: https://developer.apple.com/documentation/ScreenTime/STScreenTimeConfigurationObserver/configuration
func (s_ STScreenTimeConfigurationObserver) Configuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("configuration"))
	return rv
}


