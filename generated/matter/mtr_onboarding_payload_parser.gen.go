// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MTROnboardingPayloadParser */


/* debug [class_header]: Header for MTROnboardingPayloadParser */
// The class instance for the [MTROnboardingPayloadParser] class.
var (
	MTROnboardingPayloadParserClass     _MTROnboardingPayloadParserClass
	MTROnboardingPayloadParserClassOnce sync.Once
)

func getMTROnboardingPayloadParserClass() _MTROnboardingPayloadParserClass {
	MTROnboardingPayloadParserClassOnce.Do(func() {
		MTROnboardingPayloadParserClass = _MTROnboardingPayloadParserClass{objc.GetClass("MTROnboardingPayloadParser")}
	})
	return MTROnboardingPayloadParserClass
}

type _MTROnboardingPayloadParserClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MTROnboardingPayloadParser */
// An interface definition for the [MTROnboardingPayloadParser] class.
type IMTROnboardingPayloadParser interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MTROnboardingPayloadParser */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MTROnboardingPayloadParser */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MTROnboardingPayloadParser */
// Alloc allocates a new instance without initialization.
func (mc _MTROnboardingPayloadParserClass) Alloc() MTROnboardingPayloadParser {
	rv := objc.Send[MTROnboardingPayloadParser](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MTROnboardingPayloadParserClass) New() MTROnboardingPayloadParser {
	rv := objc.Send[MTROnboardingPayloadParser](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MTROnboardingPayloadParser) Init() MTROnboardingPayloadParser {
	rv := objc.Send[MTROnboardingPayloadParser](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MTROnboardingPayloadParser) Autorelease() MTROnboardingPayloadParser {
	rv := objc.Send[MTROnboardingPayloadParser](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTROnboardingPayloadParser creates a new MTROnboardingPayloadParser instance.
func NewMTROnboardingPayloadParser() MTROnboardingPayloadParser {
	return getMTROnboardingPayloadParserClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MTROnboardingPayloadParser */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnboardingPayloadParser
type MTROnboardingPayloadParser struct {
	objectivec.Object
}

// MTROnboardingPayloadParserFrom constructs a [MTROnboardingPayloadParser] from an unsafe.Pointer.
func MTROnboardingPayloadParserFrom(ptr unsafe.Pointer) MTROnboardingPayloadParser {
	return MTROnboardingPayloadParser{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MTROnboardingPayloadParser *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MTROnboardingPayloadParser */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Matter/MTROnboardingPayloadParser/setupPayload(forOnboardingPayload:)
func (mc _MTROnboardingPayloadParserClass) SetupPayloadForOnboardingPayloadError(onboardingPayload objc.IObject /* cross-framework: NSString */, error_ unsafe.Pointer) objc.IObject /* cross-framework: MTRSetupPayload */ {
	rv := objc.Send[objc.ID](objc.ID(mc.class), objc.Sel("setupPayloadForOnboardingPayload:error:"), onboardingPayload, error_)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SetupPayloadForOnboardingPayloadError) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MTROnboardingPayloadParser */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MTROnboardingPayloadParser */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MTROnboardingPayloadParser */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MTROnboardingPayloadParser */



