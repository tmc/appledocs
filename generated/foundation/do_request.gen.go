// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class doRequest */


/* debug [class_header]: Header for doRequest */
// The class instance for the [doRequest] class.
var (
	DoRequestClass     _doRequestClass
	DoRequestClassOnce sync.Once
)

func getdoRequestClass() _doRequestClass {
	DoRequestClassOnce.Do(func() {
		DoRequestClass = _doRequestClass{objc.GetClass("doRequest")}
	})
	return DoRequestClass
}

type _doRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for doRequest */
// An interface definition for the [doRequest] class.
type IdoRequest interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for doRequest */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for doRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for doRequest */
// Alloc allocates a new instance without initialization.
func (dc _doRequestClass) Alloc() doRequest {
	rv := objc.Send[doRequest](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (dc _doRequestClass) New() doRequest {
	rv := objc.Send[doRequest](objc.ID(dc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (d_ doRequest) Init() doRequest {
	rv := objc.Send[doRequest](d_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d_ doRequest) Autorelease() doRequest {
	rv := objc.Send[doRequest](d_.ID, objc.Sel("autorelease"))
	return rv
}

// NewdoRequest creates a new doRequest instance.
func NewdoRequest() doRequest {
	return getdoRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for doRequest */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSConnection/doRequest
type doRequest struct {
	objectivec.Object
}

// doRequestFrom constructs a [doRequest] from an unsafe.Pointer.
func doRequestFrom(ptr unsafe.Pointer) doRequest {
	return doRequest{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for doRequest *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for doRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for doRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for doRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for doRequest */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class doRequest */



