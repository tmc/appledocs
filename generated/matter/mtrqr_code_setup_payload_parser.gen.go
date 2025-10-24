// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTRQRCodeSetupPayloadParser */


/* debug [class_header]: Header for MTRQRCodeSetupPayloadParser */
// The class instance for the [MTRQRCodeSetupPayloadParser] class.
var (
	MTRQRCodeSetupPayloadParserClass     _MTRQRCodeSetupPayloadParserClass
	MTRQRCodeSetupPayloadParserClassOnce sync.Once
)

func getMTRQRCodeSetupPayloadParserClass() _MTRQRCodeSetupPayloadParserClass {
	MTRQRCodeSetupPayloadParserClassOnce.Do(func() {
		MTRQRCodeSetupPayloadParserClass = _MTRQRCodeSetupPayloadParserClass{objc.GetClass("MTRQRCodeSetupPayloadParser")}
	})
	return MTRQRCodeSetupPayloadParserClass
}

type _MTRQRCodeSetupPayloadParserClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTRQRCodeSetupPayloadParser */
// An interface definition for the [MTRQRCodeSetupPayloadParser] class.
type IMTRQRCodeSetupPayloadParser interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTRQRCodeSetupPayloadParser */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTRQRCodeSetupPayloadParser */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTRQRCodeSetupPayloadParser */
// Alloc allocates a new instance without initialization.
func (mc _MTRQRCodeSetupPayloadParserClass) Alloc() MTRQRCodeSetupPayloadParser {
	rv := objc.Send[MTRQRCodeSetupPayloadParser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTRQRCodeSetupPayloadParserClass) New() MTRQRCodeSetupPayloadParser {
	rv := objc.Send[MTRQRCodeSetupPayloadParser](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTRQRCodeSetupPayloadParser) Init() MTRQRCodeSetupPayloadParser {
	rv := objc.Send[MTRQRCodeSetupPayloadParser](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTRQRCodeSetupPayloadParser) Autorelease() MTRQRCodeSetupPayloadParser {
	rv := objc.Send[MTRQRCodeSetupPayloadParser](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTRQRCodeSetupPayloadParser creates a new MTRQRCodeSetupPayloadParser instance.
func NewMTRQRCodeSetupPayloadParser() MTRQRCodeSetupPayloadParser {
	return getMTRQRCodeSetupPayloadParserClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTRQRCodeSetupPayloadParser */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRQRCodeSetupPayloadParser
type MTRQRCodeSetupPayloadParser struct {
	objectivec.Object
}

// MTRQRCodeSetupPayloadParserFrom constructs a [MTRQRCodeSetupPayloadParser] from an unsafe.Pointer.
func MTRQRCodeSetupPayloadParserFrom(ptr unsafe.Pointer) MTRQRCodeSetupPayloadParser {
	return MTRQRCodeSetupPayloadParser{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTRQRCodeSetupPayloadParser */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTRQRCodeSetupPayloadParser/init(base38Representation:)
func NewMTRQRCodeSetupPayloadParserWithBase38Representation(base38Representation objc.IObject /* cross-framework: NSString */) MTRQRCodeSetupPayloadParser {
	instance := getMTRQRCodeSetupPayloadParserClass().Alloc()
	rv := objc.Send[MTRQRCodeSetupPayloadParser](instance.ID, objc.Sel("initWithBase38Representation:"), base38Representation)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMTRQRCodeSetupPayloadParserWithBase38Representation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTRQRCodeSetupPayloadParser */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTRQRCodeSetupPayloadParser */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTRQRCodeSetupPayloadParser */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTRQRCodeSetupPayloadParser */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTRQRCodeSetupPayloadParser */


