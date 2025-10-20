// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [options] class.
var (
	OptionsClass     _optionsClass
	OptionsClassOnce sync.Once
)

func getoptionsClass() _optionsClass {
	OptionsClassOnce.Do(func() {
		OptionsClass = _optionsClass{objc.GetClass("options")}
	})
	return OptionsClass
}

type _optionsClass struct {
	class objc.Class
}

// An interface definition for the [options] class.
type Ioptions interface {
	objectivec.IObject
}

//
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/options
type options struct {
	objectivec.Object
}

// optionsFrom constructs a [options] from an unsafe.Pointer.
func optionsFrom(ptr unsafe.Pointer) options {
	return options{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _optionsClass) Alloc() options {
	rv := objc.Send[options](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _optionsClass) New() options {
	rv := objc.Send[options](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ options) Init() options {
	rv := objc.Send[options](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ options) Autorelease() options {
	rv := objc.Send[options](o_.ID, objc.Sel("autorelease"))
	return rv
}

// Newoptions creates a new options instance.
func Newoptions() options {
	return getoptionsClass().New()
}




