// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [FSTaskOptions] class.
var (
	FSTaskOptionsClass     _FSTaskOptionsClass
	FSTaskOptionsClassOnce sync.Once
)

func getFSTaskOptionsClass() _FSTaskOptionsClass {
	FSTaskOptionsClassOnce.Do(func() {
		FSTaskOptionsClass = _FSTaskOptionsClass{objc.GetClass("FSTaskOptions")}
	})
	return FSTaskOptionsClass
}

type _FSTaskOptionsClass struct {
	class objc.Class
}

// An interface definition for the [FSTaskOptions] class.
type IFSTaskOptions interface {
	objectivec.IObject
	TaskOptions() []string
	UrlForOption(option string) foundation.URL
}

// A class that passes command options to a task, optionally providing security-scoped URLs.


// A class that passes command options to a task, optionally providing security-scoped URLs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTaskOptions
type FSTaskOptions struct {
	objectivec.Object
}

// FSTaskOptionsFrom constructs a [FSTaskOptions] from an unsafe.Pointer.
//
// A class that passes command options to a task, optionally providing security-scoped URLs.
func FSTaskOptionsFrom(ptr unsafe.Pointer) FSTaskOptions {
	return FSTaskOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (fc _FSTaskOptionsClass) Alloc() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (fc _FSTaskOptionsClass) New() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FSTaskOptions) Init() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FSTaskOptions) Autorelease() FSTaskOptions {
	rv := objc.Send[FSTaskOptions](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFSTaskOptions creates a new FSTaskOptions instance.
func NewFSTaskOptions() FSTaskOptions {
	return getFSTaskOptionsClass().New()
}



// Retrieves a URL for a given option.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTaskOptions/url(forOption:)
func (f_ FSTaskOptions) UrlForOption(option string) foundation.URL {
	rv := objc.Send[foundation.URL](f_.ID, objc.Sel("urlForOption:"), objc.String(option))
	return rv
}


// An array of strings that represent command-line options for the task.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/FSKit/FSTaskOptions/taskOptions
func (f_ FSTaskOptions) TaskOptions() []string {
	rv := objc.Send[[]string](f_.ID, objc.Sel("taskOptions"))
	return rv
}



