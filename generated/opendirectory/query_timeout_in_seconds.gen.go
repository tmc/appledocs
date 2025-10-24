// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class queryTimeoutInSeconds */


/* debug [class_header]: Header for queryTimeoutInSeconds */
// The class instance for the [queryTimeoutInSeconds] class.
var (
	QueryTimeoutInSecondsClass     _queryTimeoutInSecondsClass
	QueryTimeoutInSecondsClassOnce sync.Once
)

func getqueryTimeoutInSecondsClass() _queryTimeoutInSecondsClass {
	QueryTimeoutInSecondsClassOnce.Do(func() {
		QueryTimeoutInSecondsClass = _queryTimeoutInSecondsClass{objc.GetClass("queryTimeoutInSeconds")}
	})
	return QueryTimeoutInSecondsClass
}

type _queryTimeoutInSecondsClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for queryTimeoutInSeconds */
// An interface definition for the [queryTimeoutInSeconds] class.
type IqueryTimeoutInSeconds interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for queryTimeoutInSeconds */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for queryTimeoutInSeconds */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for queryTimeoutInSeconds */
// Alloc allocates a new instance without initialization.
func (qc _queryTimeoutInSecondsClass) Alloc() queryTimeoutInSeconds {
	rv := objc.Send[queryTimeoutInSeconds](objc.ID(qc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (qc _queryTimeoutInSecondsClass) New() queryTimeoutInSeconds {
	rv := objc.Send[queryTimeoutInSeconds](objc.ID(qc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (q_ queryTimeoutInSeconds) Init() queryTimeoutInSeconds {
	rv := objc.Send[queryTimeoutInSeconds](q_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (q_ queryTimeoutInSeconds) Autorelease() queryTimeoutInSeconds {
	rv := objc.Send[queryTimeoutInSeconds](q_.ID, objc.Sel("autorelease"))
	return rv
}

// NewqueryTimeoutInSeconds creates a new queryTimeoutInSeconds instance.
func NewqueryTimeoutInSeconds() queryTimeoutInSeconds {
	return getqueryTimeoutInSecondsClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for queryTimeoutInSeconds */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/OpenDirectory/ODConfiguration/queryTimeoutInSeconds-c.ivar
type queryTimeoutInSeconds struct {
	objectivec.Object
}

// queryTimeoutInSecondsFrom constructs a [queryTimeoutInSeconds] from an unsafe.Pointer.
func queryTimeoutInSecondsFrom(ptr unsafe.Pointer) queryTimeoutInSeconds {
	return queryTimeoutInSeconds{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for queryTimeoutInSeconds *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for queryTimeoutInSeconds */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for queryTimeoutInSeconds */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for queryTimeoutInSeconds */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for queryTimeoutInSeconds */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class queryTimeoutInSeconds */



