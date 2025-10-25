// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSSecureUnarchiveFromDataTransformer */


/* debug [class_header]: Header for NSSecureUnarchiveFromDataTransformer */
// The class instance for the [SecureUnarchiveFromDataTransformer] class.
var (
	SecureUnarchiveFromDataTransformerClass     _SecureUnarchiveFromDataTransformerClass
	SecureUnarchiveFromDataTransformerClassOnce sync.Once
)

func getSecureUnarchiveFromDataTransformerClass() _SecureUnarchiveFromDataTransformerClass {
	SecureUnarchiveFromDataTransformerClassOnce.Do(func() {
		SecureUnarchiveFromDataTransformerClass = _SecureUnarchiveFromDataTransformerClass{objc.GetClass("NSSecureUnarchiveFromDataTransformer")}
	})
	return SecureUnarchiveFromDataTransformerClass
}

type _SecureUnarchiveFromDataTransformerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SecureUnarchiveFromDataTransformer */
// An interface definition for the [SecureUnarchiveFromDataTransformer] class.
type ISecureUnarchiveFromDataTransformer interface {
	IValueTransformer
	
/* debug [class_interface_properties]: Properties for SecureUnarchiveFromDataTransformer */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SecureUnarchiveFromDataTransformer */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SecureUnarchiveFromDataTransformer */
// Alloc allocates a new instance without initialization.
func (sc _SecureUnarchiveFromDataTransformerClass) Alloc() SecureUnarchiveFromDataTransformer {
	rv := objc.Send[SecureUnarchiveFromDataTransformer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SecureUnarchiveFromDataTransformerClass) New() SecureUnarchiveFromDataTransformer {
	rv := objc.Send[SecureUnarchiveFromDataTransformer](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SecureUnarchiveFromDataTransformer) Init() SecureUnarchiveFromDataTransformer {
	rv := objc.Send[SecureUnarchiveFromDataTransformer](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SecureUnarchiveFromDataTransformer) Autorelease() SecureUnarchiveFromDataTransformer {
	rv := objc.Send[SecureUnarchiveFromDataTransformer](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSecureUnarchiveFromDataTransformer creates a new SecureUnarchiveFromDataTransformer instance.
func NewSecureUnarchiveFromDataTransformer() SecureUnarchiveFromDataTransformer {
	return getSecureUnarchiveFromDataTransformerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SecureUnarchiveFromDataTransformer */
// A value transformer that converts data to and from classes that support secure coding.
//
// This class provides a default implementation for secure decoding. This class attempts to decode data into the classes listed within , which includes , , , , , , , , , and . To archive or unarchive other classes that support , create a subclass and override to list the classes to transform. To use with , use the name of this class, or the name of a subclass you implement, as the name of the transformer for an entity’s attribute within a Core Data Model. If you use your own transformer subclass, register it with your app before intializing your persistent container with Core Data. For an example of subclassing , see , which has a class that transforms to and the reverse, to support archiving instances of .


// A value transformer that converts data to and from classes that support secure coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSecureUnarchiveFromDataTransformer
type SecureUnarchiveFromDataTransformer struct {
	ValueTransformer
}

// SecureUnarchiveFromDataTransformerFrom constructs a [SecureUnarchiveFromDataTransformer] from an unsafe.Pointer.
//
// A value transformer that converts data to and from classes that support secure coding.
func SecureUnarchiveFromDataTransformerFrom(ptr unsafe.Pointer) SecureUnarchiveFromDataTransformer {
	return SecureUnarchiveFromDataTransformer{
		ValueTransformer: ValueTransformerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SecureUnarchiveFromDataTransformer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SecureUnarchiveFromDataTransformer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SecureUnarchiveFromDataTransformer */

// A list of allowed classes the top-level object in an archive must conform to, for encoding and decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSecureUnarchiveFromDataTransformer/allowedTopLevelClasses
func (sc _SecureUnarchiveFromDataTransformerClass) AllowedTopLevelClasses() []objc.Class {
	rv := objc.Send[[]objc.Class](objc.ID(sc.class), objc.Sel("allowedTopLevelClasses"))
	return rv
}/* debug [class_properties_class/property]: allowedTopLevelClasses */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SecureUnarchiveFromDataTransformer */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SecureUnarchiveFromDataTransformer */

// A list of allowed classes the top-level object in an archive must conform to, for encoding and decoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSecureUnarchiveFromDataTransformer/allowedTopLevelClasses
func (s_ SecureUnarchiveFromDataTransformer) AllowedTopLevelClasses() []objc.Class {
	rv := objc.Send[[]objc.Class](s_.ID, objc.Sel("allowedTopLevelClasses"))
	return rv
}/* debug [instance_properties/getter]: allowedTopLevelClasses */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSSecureUnarchiveFromDataTransformer */



