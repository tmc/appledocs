// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CLSObject */


/* debug [class_header]: Header for CLSObject */
// The class instance for the [SObject] class.
var (
	SObjectClass     _SObjectClass
	SObjectClassOnce sync.Once
)

func getSObjectClass() _SObjectClass {
	SObjectClassOnce.Do(func() {
		SObjectClass = _SObjectClass{objc.GetClass("CLSObject")}
	})
	return SObjectClass
}

type _SObjectClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SObject */
// An interface definition for the [SObject] class.
type ISObject interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for SObject */
	// properties:
	DateCreated() objc.IObject /* cross-framework: NSDate */
	DateLastModified() objc.IObject /* cross-framework: NSDate */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SObject */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SObject */
// Alloc allocates a new instance without initialization.
func (sc _SObjectClass) Alloc() SObject {
	rv := objc.Send[SObject](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (sc _SObjectClass) New() SObject {
	rv := objc.Send[SObject](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SObject) Init() SObject {
	rv := objc.Send[SObject](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SObject) Autorelease() SObject {
	rv := objc.Send[SObject](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSObject creates a new SObject instance.
func NewSObject() SObject {
	return getSObjectClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SObject */
// The abstract base class for objects managed by ClassKit.


// The abstract base class for objects managed by ClassKit.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSObject
type SObject struct {
	objectivec.Object
}

// SObjectFrom constructs a [SObject] from an unsafe.Pointer.
//
// The abstract base class for objects managed by ClassKit.
func SObjectFrom(ptr unsafe.Pointer) SObject {
	return SObject{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SObject *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SObject */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SObject */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SObject */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SObject */

// The date on which the object was created.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSObject/dateCreated
func (s_ SObject) DateCreated() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](s_.ID, objc.Sel("dateCreated"))
	return rv
}/* debug [instance_properties/getter]: dateCreated */


// The date on which the object was last modified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSObject/dateLastModified
func (s_ SObject) DateLastModified() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](s_.ID, objc.Sel("dateLastModified"))
	return rv
}/* debug [instance_properties/getter]: dateLastModified */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CLSObject */



