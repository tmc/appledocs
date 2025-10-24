// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSString */


/* debug [class_header]: Header for NSString */
// The class instance for the [String] class.
var (
	StringClass     _StringClass
	StringClassOnce sync.Once
)

func getStringClass() _StringClass {
	StringClassOnce.Do(func() {
		StringClass = _StringClass{objc.GetClass("NSString")}
	})
	return StringClass
}

type _StringClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for String */
// An interface definition for the [String] class.
type IString interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for String */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for String */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for String */
// Alloc allocates a new instance without initialization.
func (sc _StringClass) Alloc() String {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _StringClass) New() String {
	rv := objc.Send[String](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ String) Init() String {
	rv := objc.Send[String](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ String) Autorelease() String {
	rv := objc.Send[String](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewString creates a new String instance.
func NewString() String {
	return getStringClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for String */
// A parent class referenced by other CoreSpotlight classes.


// A parent class referenced by other CoreSpotlight classes. [Full Topic]
type String struct {
	objectivec.Object
}

// StringFrom constructs a [String] from an unsafe.Pointer.
//
// A parent class referenced by other CoreSpotlight classes.
func StringFrom(ptr unsafe.Pointer) String {
	return String{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for String *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for String */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for String */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for String */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for String */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSString */



