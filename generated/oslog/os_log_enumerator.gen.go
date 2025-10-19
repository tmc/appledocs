// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [OSLogEnumerator] class.
var oSLogEnumeratorClass = _OSLogEnumeratorClass{objc.GetClass("OSLogEnumerator")}

type _OSLogEnumeratorClass struct {
	class objc.Class
}

// An interface definition for the [OSLogEnumerator] class.
type IOSLogEnumerator interface {
	foundation.IEnumerator
}

// An enumerator that can access and list log entries. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (oc _OSLogEnumeratorClass) Alloc() OSLogEnumerator {
	rv := objc.Send[OSLogEnumerator](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return oSLogEnumeratorClass.New()
}




