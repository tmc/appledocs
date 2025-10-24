// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [automaticTerminationOptOutCounter] class.
var (
	AutomaticTerminationOptOutCounterClass     _automaticTerminationOptOutCounterClass
	AutomaticTerminationOptOutCounterClassOnce sync.Once
)

func getautomaticTerminationOptOutCounterClass() _automaticTerminationOptOutCounterClass {
	AutomaticTerminationOptOutCounterClassOnce.Do(func() {
		AutomaticTerminationOptOutCounterClass = _automaticTerminationOptOutCounterClass{objc.GetClass("automaticTerminationOptOutCounter")}
	})
	return AutomaticTerminationOptOutCounterClass
}

type _automaticTerminationOptOutCounterClass struct {
	class objc.Class
}

// An interface definition for the [automaticTerminationOptOutCounter] class.
type IautomaticTerminationOptOutCounter interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProcessInfo/automaticTerminationOptOutCounter
type automaticTerminationOptOutCounter struct {
	objectivec.Object
}

// automaticTerminationOptOutCounterFrom constructs a [automaticTerminationOptOutCounter] from an unsafe.Pointer.
func automaticTerminationOptOutCounterFrom(ptr unsafe.Pointer) automaticTerminationOptOutCounter {
	return automaticTerminationOptOutCounter{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _automaticTerminationOptOutCounterClass) Alloc() automaticTerminationOptOutCounter {
	rv := objc.Send[automaticTerminationOptOutCounter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _automaticTerminationOptOutCounterClass) New() automaticTerminationOptOutCounter {
	rv := objc.Send[automaticTerminationOptOutCounter](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ automaticTerminationOptOutCounter) Init() automaticTerminationOptOutCounter {
	rv := objc.Send[automaticTerminationOptOutCounter](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ automaticTerminationOptOutCounter) Autorelease() automaticTerminationOptOutCounter {
	rv := objc.Send[automaticTerminationOptOutCounter](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewautomaticTerminationOptOutCounter creates a new automaticTerminationOptOutCounter instance.
func NewautomaticTerminationOptOutCounter() automaticTerminationOptOutCounter {
	return getautomaticTerminationOptOutCounterClass().New()
}




