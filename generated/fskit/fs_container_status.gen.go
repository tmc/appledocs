// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [FSContainerStatus] class.
var (
	FSContainerStatusClass     _FSContainerStatusClass
	FSContainerStatusClassOnce sync.Once
)

func getFSContainerStatusClass() _FSContainerStatusClass {
	FSContainerStatusClassOnce.Do(func() {
		FSContainerStatusClass = _FSContainerStatusClass{objc.GetClass("FSContainerStatus")}
	})
	return FSContainerStatusClass
}

type _FSContainerStatusClass struct {
	class objc.Class
}





// An interface definition for the [FSContainerStatus] class.
type IFSContainerStatus interface {
	objectivec.IObject
	

	// properties:
	State() FSContainerState
	Status() foundation.foundation.INSError


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (fc _FSContainerStatusClass) Alloc() FSContainerStatus {
	rv := objc.Send[FSContainerStatus](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FSContainerStatusClass) New() FSContainerStatus {
	rv := objc.Send[FSContainerStatus](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSContainerStatus) Init() FSContainerStatus {
	rv := objc.Send[FSContainerStatus](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSContainerStatus) Autorelease() FSContainerStatus {
	rv := objc.Send[FSContainerStatus](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSContainerStatus creates a new FSContainerStatus instance.
func NewFSContainerStatus() FSContainerStatus {
	return getFSContainerStatusClass().New()
}





// A type that represents a container’s status.
//
// This type contains two properties: The value that indicates the state of the container, such as or . The is an error (optional in Swift, nullable in Objective-C) that provides further information about the state, such as why the container is blocked. Examples of statuses that require intervention include errors that indicate the container isn’t ready (POSIX or ), the container needs authentication ( ), or that authentication failed ( ). The status can also be an informative error, such as the FSKit error .


// A type that represents a container’s status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus
type FSContainerStatus struct {
	objectivec.Object
}

// FSContainerStatusFrom constructs a [FSContainerStatus] from an unsafe.Pointer.
//
// A type that represents a container’s status.
func FSContainerStatusFrom(ptr unsafe.Pointer) FSContainerStatus {
	return FSContainerStatus{objectivec.Object{objc.ID(ptr)}}
}










// Returns a active container status instance with the provided error status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/active(status:)
func (fc _FSContainerStatusClass) ActiveWithStatus(errorStatus foundation.foundation.INSError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("activeWithStatus:"), errorStatus)
	return rv
}


// Returns a blocked container status instance with the provided error status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/blocked(status:)
func (fc _FSContainerStatusClass) BlockedWithStatus(errorStatus foundation.foundation.INSError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("blockedWithStatus:"), errorStatus)
	return rv
}


// Returns a not-ready container status instance with the provided error status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/notReady(status:)
func (fc _FSContainerStatusClass) NotReadyWithStatus(errorStatus foundation.foundation.INSError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("notReadyWithStatus:"), errorStatus)
	return rv
}


// Returns a ready container status instance with the provided error status.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/ready(status:)
func (fc _FSContainerStatusClass) ReadyWithStatus(errorStatus foundation.foundation.INSError) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(fc.class), objc.Sel("readyWithStatus:"), errorStatus)
	return rv
}







// A status that represents an active container with no error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/active
func (fc _FSContainerStatusClass) Active() FSContainerStatus {
	rv := objc.Send[FSContainerStatus](objc.ID(fc.class), objc.Sel("active"))
	return rv
}

// A status that represents a ready container with no error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/ready
func (fc _FSContainerStatusClass) Ready() FSContainerStatus {
	rv := objc.Send[FSContainerStatus](objc.ID(fc.class), objc.Sel("ready"))
	return rv
}











// A status that represents an active container with no error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/active
func (f_ FSContainerStatus) Active() IFSContainerStatus {
	rv := objc.Send[FSContainerStatus](f_.ID, objc.Sel("active"))
	return rv
}


// A status that represents a ready container with no error.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/ready
func (f_ FSContainerStatus) Ready() IFSContainerStatus {
	rv := objc.Send[FSContainerStatus](f_.ID, objc.Sel("ready"))
	return rv
}


// A value that represents the container state, such as ready, active, or blocked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/state
func (f_ FSContainerStatus) State() FSContainerState {
	rv := objc.Send[FSContainerState](f_.ID, objc.Sel("state"))
	return rv
}


// An optional error that provides further information about the state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSContainerStatus/status
func (f_ FSContainerStatus) Status() foundation.foundation.INSError {
	rv := objc.Send[foundation.NSError](f_.ID, objc.Sel("status"))
	return rv
}








