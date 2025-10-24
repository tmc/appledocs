// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MEFormatReaderInstantiationOptions */


/* debug [class_header]: Header for MEFormatReaderInstantiationOptions */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MEFormatReaderInstantiationOptions */
// An interface definition for the [MEFormatReaderInstantiationOptions] class.
type IMEFormatReaderInstantiationOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MEFormatReaderInstantiationOptions */
	// properties:
	AllowIncrementalFragmentParsing() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MEFormatReaderInstantiationOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MEFormatReaderInstantiationOptions */
// Alloc allocates a new instance without initialization.
func (mc _MEFormatReaderInstantiationOptionsClass) Alloc() MEFormatReaderInstantiationOptions {
	rv := objc.Send[MEFormatReaderInstantiationOptions](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MEFormatReaderInstantiationOptions */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MEFormatReaderInstantiationOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MEFormatReaderInstantiationOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MEFormatReaderInstantiationOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MEFormatReaderInstantiationOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MEFormatReaderInstantiationOptions */

// Enables support for parsing additional fragments.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderInstantiationOptions/allowIncrementalFragmentParsing
func (m_ MEFormatReaderInstantiationOptions) AllowIncrementalFragmentParsing() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("allowIncrementalFragmentParsing"))
	return rv
}/* debug [instance_properties/getter]: allowIncrementalFragmentParsing */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MEFormatReaderInstantiationOptions */



