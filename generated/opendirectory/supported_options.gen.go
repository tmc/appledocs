// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class supportedOptions */


/* debug [class_header]: Header for supportedOptions */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for supportedOptions */
// An interface definition for the [supportedOptions] class.
type IsupportedOptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for supportedOptions */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for supportedOptions */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for supportedOptions */
// Alloc allocates a new instance without initialization.
func (sc _supportedOptionsClass) Alloc() supportedOptions {
	rv := objc.Send[supportedOptions](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for supportedOptions */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/supportedOptions-c.ivar
type supportedOptions struct {
	objectivec.Object
}

// supportedOptionsFrom constructs a [supportedOptions] from an unsafe.Pointer.
func supportedOptionsFrom(ptr unsafe.Pointer) supportedOptions {
	return supportedOptions{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for supportedOptions *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for supportedOptions */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for supportedOptions */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for supportedOptions */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for supportedOptions */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class supportedOptions */



