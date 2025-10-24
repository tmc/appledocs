// Code generated from Apple documentation for Contacts. DO NOT EDIT.

package contacts

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CNFetchRequest */


/* debug [class_header]: Header for CNFetchRequest */
// The class instance for the [CNFetchRequest] class.
var (
	CNFetchRequestClass     _CNFetchRequestClass
	CNFetchRequestClassOnce sync.Once
)

func getCNFetchRequestClass() _CNFetchRequestClass {
	CNFetchRequestClassOnce.Do(func() {
		CNFetchRequestClass = _CNFetchRequestClass{objc.GetClass("CNFetchRequest")}
	})
	return CNFetchRequestClass
}

type _CNFetchRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CNFetchRequest */
// An interface definition for the [CNFetchRequest] class.
type ICNFetchRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CNFetchRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CNFetchRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CNFetchRequest */
// Alloc allocates a new instance without initialization.
func (cc _CNFetchRequestClass) Alloc() CNFetchRequest {
	rv := objc.Send[CNFetchRequest](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CNFetchRequestClass) New() CNFetchRequest {
	rv := objc.Send[CNFetchRequest](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CNFetchRequest) Init() CNFetchRequest {
	rv := objc.Send[CNFetchRequest](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CNFetchRequest) Autorelease() CNFetchRequest {
	rv := objc.Send[CNFetchRequest](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCNFetchRequest creates a new CNFetchRequest instance.
func NewCNFetchRequest() CNFetchRequest {
	return getCNFetchRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CNFetchRequest */
// The base class for contact fetch requests.
//
// To fetch contacts, use .


// The base class for contact fetch requests.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Contacts/CNFetchRequest
type CNFetchRequest struct {
	objectivec.Object
}

// CNFetchRequestFrom constructs a [CNFetchRequest] from an unsafe.Pointer.
//
// The base class for contact fetch requests.
func CNFetchRequestFrom(ptr unsafe.Pointer) CNFetchRequest {
	return CNFetchRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CNFetchRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CNFetchRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CNFetchRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CNFetchRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CNFetchRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CNFetchRequest */



