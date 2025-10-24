// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

package dispatch

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [OS_object] class.
var (
	OS_objectClass     _OS_objectClass
	OS_objectClassOnce sync.Once
)

func getOS_objectClass() _OS_objectClass {
	OS_objectClassOnce.Do(func() {
		OS_objectClass = _OS_objectClass{objc.GetClass("OS_object")}
	})
	return OS_objectClass
}

type _OS_objectClass struct {
	class objc.Class
}

// An interface definition for the [OS_object] class.
type IOS_object interface {
	objectivec.IObject
	// properties:
	// methods:
}

// A parent class referenced by other Dispatch classes.


// A parent class referenced by other Dispatch classes. [Full Topic]
type OS_object struct {
	objectivec.Object
}

// OS_objectFrom constructs a [OS_object] from an unsafe.Pointer.
//
// A parent class referenced by other Dispatch classes.
func OS_objectFrom(ptr unsafe.Pointer) OS_object {
	return OS_object{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OS_objectClass) Alloc() OS_object {
	rv := objc.Send[OS_object](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OS_objectClass) New() OS_object {
	rv := objc.Send[OS_object](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OS_object) Init() OS_object {
	rv := objc.Send[OS_object](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OS_object) Autorelease() OS_object {
	rv := objc.Send[OS_object](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOS_object creates a new OS_object instance.
func NewOS_object() OS_object {
	return getOS_objectClass().New()
}




