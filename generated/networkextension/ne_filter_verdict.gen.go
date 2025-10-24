// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEFilterVerdict */


/* debug [class_header]: Header for NEFilterVerdict */
// The class instance for the [NEFilterVerdict] class.
var (
	NEFilterVerdictClass     _NEFilterVerdictClass
	NEFilterVerdictClassOnce sync.Once
)

func getNEFilterVerdictClass() _NEFilterVerdictClass {
	NEFilterVerdictClassOnce.Do(func() {
		NEFilterVerdictClass = _NEFilterVerdictClass{objc.GetClass("NEFilterVerdict")}
	})
	return NEFilterVerdictClass
}

type _NEFilterVerdictClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEFilterVerdict */
// An interface definition for the [NEFilterVerdict] class.
type INEFilterVerdict interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEFilterVerdict */
	// properties:
	ShouldReport() bool
	SetShouldReport(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEFilterVerdict */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEFilterVerdict */
// Alloc allocates a new instance without initialization.
func (nc _NEFilterVerdictClass) Alloc() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEFilterVerdictClass) New() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEFilterVerdict) Init() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEFilterVerdict) Autorelease() NEFilterVerdict {
	rv := objc.Send[NEFilterVerdict](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEFilterVerdict creates a new NEFilterVerdict instance.
func NewNEFilterVerdict() NEFilterVerdict {
	return getNEFilterVerdictClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEFilterVerdict */
// The abstract base class for filter verdict classes.
//
// Filter providers use instances this class to inform the system about how to handle flows of network data.


// The abstract base class for filter verdict classes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterVerdict
type NEFilterVerdict struct {
	objectivec.Object
}

// NEFilterVerdictFrom constructs a [NEFilterVerdict] from an unsafe.Pointer.
//
// The abstract base class for filter verdict classes.
func NEFilterVerdictFrom(ptr unsafe.Pointer) NEFilterVerdict {
	return NEFilterVerdict{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEFilterVerdict *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEFilterVerdict */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEFilterVerdict */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEFilterVerdict */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEFilterVerdict */

// A Boolean value that indicates whether to send a report to the control provider when processing this verdict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterVerdict/shouldReport
func (n_ NEFilterVerdict) ShouldReport() bool {
	rv := objc.Send[bool](n_.ID, objc.Sel("shouldReport"))
	return rv
}/* debug [instance_properties/getter]: shouldReport */


// A Boolean value that indicates whether to send a report to the control provider when processing this verdict.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEFilterVerdict/shouldReport
func (n_ NEFilterVerdict) SetShouldReport(value bool) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setShouldReport:"), value)
}/* debug [instance_properties/setter]: shouldReport */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEFilterVerdict */



