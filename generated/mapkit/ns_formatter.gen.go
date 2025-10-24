// Code generated from Apple documentation for MapKit. DO NOT EDIT.

package mapkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSFormatter */


/* debug [class_header]: Header for NSFormatter */
// The class instance for the [Formatter] class.
var (
	FormatterClass     _FormatterClass
	FormatterClassOnce sync.Once
)

func getFormatterClass() _FormatterClass {
	FormatterClassOnce.Do(func() {
		FormatterClass = _FormatterClass{objc.GetClass("NSFormatter")}
	})
	return FormatterClass
}

type _FormatterClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Formatter */
// An interface definition for the [Formatter] class.
type IFormatter interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Formatter */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Formatter */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Formatter */
// Alloc allocates a new instance without initialization.
func (fc _FormatterClass) Alloc() Formatter {
	rv := objc.Send[Formatter](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FormatterClass) New() Formatter {
	rv := objc.Send[Formatter](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ Formatter) Init() Formatter {
	rv := objc.Send[Formatter](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ Formatter) Autorelease() Formatter {
	rv := objc.Send[Formatter](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFormatter creates a new Formatter instance.
func NewFormatter() Formatter {
	return getFormatterClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Formatter */
// A parent class referenced by other MapKit classes.


// A parent class referenced by other MapKit classes. [Full Topic]
type Formatter struct {
	objectivec.Object
}

// FormatterFrom constructs a [Formatter] from an unsafe.Pointer.
//
// A parent class referenced by other MapKit classes.
func FormatterFrom(ptr unsafe.Pointer) Formatter {
	return Formatter{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Formatter *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Formatter */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Formatter */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Formatter */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Formatter */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSFormatter */



