// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)





// The class instance for the [OSLogEnumerator] class.
var (
	OSLogEnumeratorClass     _OSLogEnumeratorClass
	OSLogEnumeratorClassOnce sync.Once
)

func getOSLogEnumeratorClass() _OSLogEnumeratorClass {
	OSLogEnumeratorClassOnce.Do(func() {
		OSLogEnumeratorClass = _OSLogEnumeratorClass{objc.GetClass("OSLogEnumerator")}
	})
	return OSLogEnumeratorClass
}

type _OSLogEnumeratorClass struct {
	class objc.Class
}





// An interface definition for the [OSLogEnumerator] class.
type IOSLogEnumerator interface {
	foundation.IEnumerator
	

	// properties:


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (oc _OSLogEnumeratorClass) Alloc() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (oc _OSLogEnumeratorClass) New() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogEnumerator) Init() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogEnumerator) Autorelease() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogEnumerator creates a new OSLogEnumerator instance.
func NewOSLogEnumerator() OSLogEnumerator {
	return getOSLogEnumeratorClass().New()
}





// An enumerator that can access and list log entries.


// An enumerator that can access and list log entries.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogEnumerator
type OSLogEnumerator struct {
	foundation.Enumerator
}

// OSLogEnumeratorFrom constructs a [OSLogEnumerator] from an unsafe.Pointer.
//
// An enumerator that can access and list log entries.
func OSLogEnumeratorFrom(ptr unsafe.Pointer) OSLogEnumerator {
	return OSLogEnumerator{
		Enumerator: foundation.EnumeratorFrom(ptr),
	}
}































