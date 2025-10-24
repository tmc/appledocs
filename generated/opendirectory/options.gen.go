// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class options */


/* debug [class_header]: Header for options */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for options */
// An interface definition for the [options] class.
type Ioptions interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for options */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for options */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for options */
// Alloc allocates a new instance without initialization.
func (oc _optionsClass) Alloc() options {
	rv := objc.Send[options](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for options */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODModuleEntry/options
type options struct {
	objectivec.Object
}

// optionsFrom constructs a [options] from an unsafe.Pointer.
func optionsFrom(ptr unsafe.Pointer) options {
	return options{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for options *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for options */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for options */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for options */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for options */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class options */



