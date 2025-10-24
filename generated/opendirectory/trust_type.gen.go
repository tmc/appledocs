// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class trustType */


/* debug [class_header]: Header for trustType */
// The class instance for the [trustType] class.
var (
	TrustTypeClass     _trustTypeClass
	TrustTypeClassOnce sync.Once
)

func gettrustTypeClass() _trustTypeClass {
	TrustTypeClassOnce.Do(func() {
		TrustTypeClass = _trustTypeClass{objc.GetClass("trustType")}
	})
	return TrustTypeClass
}

type _trustTypeClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for trustType */
// An interface definition for the [trustType] class.
type ItrustType interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for trustType */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for trustType */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for trustType */
// Alloc allocates a new instance without initialization.
func (tc _trustTypeClass) Alloc() trustType {
	rv := objc.Send[trustType](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _trustTypeClass) New() trustType {
	rv := objc.Send[trustType](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ trustType) Init() trustType {
	rv := objc.Send[trustType](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ trustType) Autorelease() trustType {
	rv := objc.Send[trustType](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewtrustType creates a new trustType instance.
func NewtrustType() trustType {
	return gettrustTypeClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for trustType */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/trustType-c.ivar
type trustType struct {
	objectivec.Object
}

// trustTypeFrom constructs a [trustType] from an unsafe.Pointer.
func trustTypeFrom(ptr unsafe.Pointer) trustType {
	return trustType{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for trustType *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for trustType */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for trustType */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for trustType */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for trustType */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class trustType */



