// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [SharedEventListener] class.
var (
	SharedEventListenerClass     _SharedEventListenerClass
	SharedEventListenerClassOnce sync.Once
)

func getSharedEventListenerClass() _SharedEventListenerClass {
	SharedEventListenerClassOnce.Do(func() {
		SharedEventListenerClass = _SharedEventListenerClass{objc.GetClass("MTLSharedEventListener")}
	})
	return SharedEventListenerClass
}

type _SharedEventListenerClass struct {
	class objc.Class
}





// An interface definition for the [SharedEventListener] class.
type ISharedEventListener interface {
	objectivec.IObject
	

	// properties:
	DispatchQueue() objectivec.IObject


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (sc _SharedEventListenerClass) Alloc() SharedEventListener {
	rv := objc.Send[SharedEventListener](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SharedEventListenerClass) New() SharedEventListener {
	rv := objc.Send[SharedEventListener](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SharedEventListener) Init() SharedEventListener {
	rv := objc.Send[SharedEventListener](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SharedEventListener) Autorelease() SharedEventListener {
	rv := objc.Send[SharedEventListener](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSharedEventListener creates a new SharedEventListener instance.
func NewSharedEventListener() SharedEventListener {
	return getSharedEventListenerClass().New()
}





// A listener for shareable event notifications.


// A listener for shareable event notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSharedEventListener
type SharedEventListener struct {
	objectivec.Object
}

// SharedEventListenerFrom constructs a [SharedEventListener] from an unsafe.Pointer.
//
// A listener for shareable event notifications.
func SharedEventListenerFrom(ptr unsafe.Pointer) SharedEventListener {
	return SharedEventListener{objectivec.Object{objc.ID(ptr)}}
}






// Creates a new shareable event listener with a specific dispatch queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSharedEventListener/init(dispatchQueue:)
func NewSharedEventListenerWithDispatchQueue(dispatchQueue objectivec.IObject) SharedEventListener {
	instance := getSharedEventListenerClass().Alloc()
	rv := objc.Send[SharedEventListener](instance.ID, objc.Sel("initWithDispatchQueue:"), dispatchQueue)
	rv.Autorelease()
	return rv
}







// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSharedEventListener/shared()
func (sc _SharedEventListenerClass) SharedListener() ISharedEventListener {
	rv := objc.Send[SharedEventListener](objc.ID(sc.class), objc.Sel("sharedListener"))
	return rv
}

















// The dispatch queue used to dispatch any notifications.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSharedEventListener/dispatchQueue
func (s_ SharedEventListener) DispatchQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("dispatchQueue"))
	return rv
}







