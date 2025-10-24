// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROTAHeaderParser */


/* debug [class_header]: Header for MTROTAHeaderParser */
// The class instance for the [MTROTAHeaderParser] class.
var (
	MTROTAHeaderParserClass     _MTROTAHeaderParserClass
	MTROTAHeaderParserClassOnce sync.Once
)

func getMTROTAHeaderParserClass() _MTROTAHeaderParserClass {
	MTROTAHeaderParserClassOnce.Do(func() {
		MTROTAHeaderParserClass = _MTROTAHeaderParserClass{objc.GetClass("MTROTAHeaderParser")}
	})
	return MTROTAHeaderParserClass
}

type _MTROTAHeaderParserClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROTAHeaderParser */
// An interface definition for the [MTROTAHeaderParser] class.
type IMTROTAHeaderParser interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROTAHeaderParser */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROTAHeaderParser */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROTAHeaderParser */
// Alloc allocates a new instance without initialization.
func (mc _MTROTAHeaderParserClass) Alloc() MTROTAHeaderParser {
	rv := objc.Send[MTROTAHeaderParser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROTAHeaderParserClass) New() MTROTAHeaderParser {
	rv := objc.Send[MTROTAHeaderParser](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROTAHeaderParser) Init() MTROTAHeaderParser {
	rv := objc.Send[MTROTAHeaderParser](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROTAHeaderParser) Autorelease() MTROTAHeaderParser {
	rv := objc.Send[MTROTAHeaderParser](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROTAHeaderParser creates a new MTROTAHeaderParser instance.
func NewMTROTAHeaderParser() MTROTAHeaderParser {
	return getMTROTAHeaderParserClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROTAHeaderParser */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeaderParser
type MTROTAHeaderParser struct {
	objectivec.Object
}

// MTROTAHeaderParserFrom constructs a [MTROTAHeaderParser] from an unsafe.Pointer.
func MTROTAHeaderParserFrom(ptr unsafe.Pointer) MTROTAHeaderParser {
	return MTROTAHeaderParser{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROTAHeaderParser *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROTAHeaderParser */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROTAHeaderParser/header(from:)
func (mc _MTROTAHeaderParserClass) HeaderFromDataError(data objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) IMTROTAHeader {
	rv := objc.Send[MTROTAHeader](objc.ID(mc.class), objc.Sel("headerFromData:error:"), data, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=HeaderFromDataError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROTAHeaderParser */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROTAHeaderParser */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROTAHeaderParser */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROTAHeaderParser */



