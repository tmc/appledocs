// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OSLogMessageComponent] class.
var oSLogMessageComponentClass = _OSLogMessageComponentClass{objc.GetClass("OSLogMessageComponent")}

type _OSLogMessageComponentClass struct {
	class objc.Class
}

// An interface definition for the [OSLogMessageComponent] class.
type IOSLogMessageComponent interface {
	objectivec.IObject
}

// The message arguments for a particular entry. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent

type OSLogMessageComponent struct {
	objectivec.Object
}

// OSLogMessageComponentFrom constructs a [OSLogMessageComponent] from an unsafe.Pointer.
//
// The message arguments for a particular entry.
func OSLogMessageComponentFrom(ptr unsafe.Pointer) OSLogMessageComponent {
	return OSLogMessageComponent{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (oc _OSLogMessageComponentClass) Alloc() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (oc _OSLogMessageComponentClass) New() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OSLogMessageComponent) Init() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OSLogMessageComponent) Autorelease() OSLogMessageComponent {
	rv := objc.Send[OSLogMessageComponent](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOSLogMessageComponent creates a new OSLogMessageComponent instance.
func NewOSLogMessageComponent() OSLogMessageComponent {
	return oSLogMessageComponentClass.New()
}




