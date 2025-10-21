// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [INObject] class.
var (
	INObjectClass     _INObjectClass
	INObjectClassOnce sync.Once
)

func getINObjectClass() _INObjectClass {
	INObjectClassOnce.Do(func() {
		INObjectClass = _INObjectClass{objc.GetClass("INObject")}
	})
	return INObjectClass
}

type _INObjectClass struct {
	class objc.Class
}

// An interface definition for the [INObject] class.
type IINObject interface {
	objectivec.IObject
}

// A representation of a custom intent parameter or response property.
//
// Use to create custom parameters and response properties for intent data that doesn’t fit into one of the System Types, such as Boolean, Duration, or Location. Define custom types and associate them with your custom intents and responses in the Intent Definition file. Xcode uses the type defined in the Intent Definition file to generate a subclass of . Create an instance of this subclass to structure data in intents and intent responses.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INObject
type INObject struct {
	objectivec.Object
}

// INObjectFrom constructs a [INObject] from an unsafe.Pointer.
//
// A representation of a custom intent parameter or response property.
func INObjectFrom(ptr unsafe.Pointer) INObject {
	return INObject{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _INObjectClass) Alloc() INObject {
	rv := objc.Send[INObject](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INObjectClass) New() INObject {
	rv := objc.Send[INObject](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INObject) Init() INObject {
	rv := objc.Send[INObject](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INObject) Autorelease() INObject {
	rv := objc.Send[INObject](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINObject creates a new INObject instance.
func NewINObject() INObject {
	return getINObjectClass().New()
}




