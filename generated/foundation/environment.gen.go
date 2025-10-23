// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [environment] class.
var (
	EnvironmentClass     _environmentClass
	EnvironmentClassOnce sync.Once
)

func getenvironmentClass() _environmentClass {
	EnvironmentClassOnce.Do(func() {
		EnvironmentClass = _environmentClass{objc.GetClass("environment")}
	})
	return EnvironmentClass
}

type _environmentClass struct {
	class objc.Class
}

// An interface definition for the [environment] class.
type Ienvironment interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProcessInfo/environment-c.ivar
type environment struct {
	objectivec.Object
}

// environmentFrom constructs a [environment] from an unsafe.Pointer.
func environmentFrom(ptr unsafe.Pointer) environment {
	return environment{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ec _environmentClass) Alloc() environment {
	rv := objc.Send[environment](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ec _environmentClass) New() environment {
	rv := objc.Send[environment](objc.ID(ec.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (e_ environment) Init() environment {
	rv := objc.Send[environment](e_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e_ environment) Autorelease() environment {
	rv := objc.Send[environment](e_.ID, objc.Sel("autorelease"))
	return rv
}

// Newenvironment creates a new environment instance.
func Newenvironment() environment {
	return getenvironmentClass().New()
}




