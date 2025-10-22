// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [String] class.
var (
	StringClass     _StringClass
	StringClassOnce sync.Once
)

func getStringClass() _StringClass {
	StringClassOnce.Do(func() {
		StringClass = _StringClass{objc.GetClass("NSString")}
	})
	return StringClass
}

type _StringClass struct {
	class objc.Class
}

// An interface definition for the [String] class.
type IString interface {
	objectivec.IObject
}

// A parent class referenced by other CoreSpotlight classes.


// A parent class referenced by other CoreSpotlight classes. [Full Topic]

type String struct {
	objectivec.Object
}

// StringFrom constructs a [String] from an unsafe.Pointer.
//
// A parent class referenced by other CoreSpotlight classes.
func StringFrom(ptr unsafe.Pointer) String {
	return String{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StringClass) Alloc() String {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StringClass) New() String {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ String) Init() String {
	rv := objc.Send[String](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ String) Autorelease() String {
	rv := objc.Send[String](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewString creates a new String instance.
func NewString() String {
	return getStringClass().New()
}




