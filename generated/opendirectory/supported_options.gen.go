// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [supportedOptions] class.
var (
	SupportedOptionsClass     _supportedOptionsClass
	SupportedOptionsClassOnce sync.Once
)

func getsupportedOptionsClass() _supportedOptionsClass {
	SupportedOptionsClassOnce.Do(func() {
		SupportedOptionsClass = _supportedOptionsClass{objc.GetClass("supportedOptions")}
	})
	return SupportedOptionsClass
}

type _supportedOptionsClass struct {
	class objc.Class
}

// An interface definition for the [supportedOptions] class.
type IsupportedOptions interface {
	objectivec.IObject
	// properties:
	// methods:
}



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/supportedOptions-c.ivar
type supportedOptions struct {
	objectivec.Object
}

// supportedOptionsFrom constructs a [supportedOptions] from an unsafe.Pointer.
func supportedOptionsFrom(ptr unsafe.Pointer) supportedOptions {
	return supportedOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _supportedOptionsClass) Alloc() supportedOptions {
	rv := objc.Send[supportedOptions](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _supportedOptionsClass) New() supportedOptions {
	rv := objc.Send[supportedOptions](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ supportedOptions) Init() supportedOptions {
	rv := objc.Send[supportedOptions](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ supportedOptions) Autorelease() supportedOptions {
	rv := objc.Send[supportedOptions](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewsupportedOptions creates a new supportedOptions instance.
func NewsupportedOptions() supportedOptions {
	return getsupportedOptionsClass().New()
}




