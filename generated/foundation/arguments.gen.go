// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [arguments] class.
var (
	ArgumentsClass     _argumentsClass
	ArgumentsClassOnce sync.Once
)

func getargumentsClass() _argumentsClass {
	ArgumentsClassOnce.Do(func() {
		ArgumentsClass = _argumentsClass{objc.GetClass("arguments")}
	})
	return ArgumentsClass
}

type _argumentsClass struct {
	class objc.Class
}

// An interface definition for the [arguments] class.
type Iarguments interface {
	objectivec.IObject
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSProcessInfo/arguments-c.ivar
type arguments struct {
	objectivec.Object
}

// argumentsFrom constructs a [arguments] from an unsafe.Pointer.
func argumentsFrom(ptr unsafe.Pointer) arguments {
	return arguments{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _argumentsClass) Alloc() arguments {
	rv := objc.Send[arguments](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _argumentsClass) New() arguments {
	rv := objc.Send[arguments](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ arguments) Init() arguments {
	rv := objc.Send[arguments](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ arguments) Autorelease() arguments {
	rv := objc.Send[arguments](a_.ID, objc.Sel("autorelease"))
	return rv
}

// Newarguments creates a new arguments instance.
func Newarguments() arguments {
	return getargumentsClass().New()
}




