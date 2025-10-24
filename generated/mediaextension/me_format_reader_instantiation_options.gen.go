// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MEFormatReaderInstantiationOptions] class.
var (
	MEFormatReaderInstantiationOptionsClass     _MEFormatReaderInstantiationOptionsClass
	MEFormatReaderInstantiationOptionsClassOnce sync.Once
)

func getMEFormatReaderInstantiationOptionsClass() _MEFormatReaderInstantiationOptionsClass {
	MEFormatReaderInstantiationOptionsClassOnce.Do(func() {
		MEFormatReaderInstantiationOptionsClass = _MEFormatReaderInstantiationOptionsClass{objc.GetClass("MEFormatReaderInstantiationOptions")}
	})
	return MEFormatReaderInstantiationOptionsClass
}

type _MEFormatReaderInstantiationOptionsClass struct {
	class objc.Class
}

// An interface definition for the [MEFormatReaderInstantiationOptions] class.
type IMEFormatReaderInstantiationOptions interface {
	objectivec.IObject
	// properties:
	AllowIncrementalFragmentParsing() bool
	// methods:
}

// An object that contains options to pass to a format reader extension.
//
// This object is mutable with options set through instance properties.


// An object that contains options to pass to a format reader extension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderInstantiationOptions
type MEFormatReaderInstantiationOptions struct {
	objectivec.Object
}

// MEFormatReaderInstantiationOptionsFrom constructs a [MEFormatReaderInstantiationOptions] from an unsafe.Pointer.
//
// An object that contains options to pass to a format reader extension.
func MEFormatReaderInstantiationOptionsFrom(ptr unsafe.Pointer) MEFormatReaderInstantiationOptions {
	return MEFormatReaderInstantiationOptions{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MEFormatReaderInstantiationOptionsClass) Alloc() MEFormatReaderInstantiationOptions {
	rv := objc.Send[MEFormatReaderInstantiationOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MEFormatReaderInstantiationOptionsClass) New() MEFormatReaderInstantiationOptions {
	rv := objc.Send[MEFormatReaderInstantiationOptions](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MEFormatReaderInstantiationOptions) Init() MEFormatReaderInstantiationOptions {
	rv := objc.Send[MEFormatReaderInstantiationOptions](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MEFormatReaderInstantiationOptions) Autorelease() MEFormatReaderInstantiationOptions {
	rv := objc.Send[MEFormatReaderInstantiationOptions](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMEFormatReaderInstantiationOptions creates a new MEFormatReaderInstantiationOptions instance.
func NewMEFormatReaderInstantiationOptions() MEFormatReaderInstantiationOptions {
	return getMEFormatReaderInstantiationOptionsClass().New()
}



// Enables support for parsing additional fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderInstantiationOptions/allowIncrementalFragmentParsing
func (m_ MEFormatReaderInstantiationOptions) AllowIncrementalFragmentParsing() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowIncrementalFragmentParsing"))
	return rv
}



