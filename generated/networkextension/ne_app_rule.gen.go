// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NEAppRule */


/* debug [class_header]: Header for NEAppRule */
// The class instance for the [NEAppRule] class.
var (
	NEAppRuleClass     _NEAppRuleClass
	NEAppRuleClassOnce sync.Once
)

func getNEAppRuleClass() _NEAppRuleClass {
	NEAppRuleClassOnce.Do(func() {
		NEAppRuleClass = _NEAppRuleClass{objc.GetClass("NEAppRule")}
	})
	return NEAppRuleClass
}

type _NEAppRuleClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for NEAppRule */
// An interface definition for the [NEAppRule] class.
type INEAppRule interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for NEAppRule */
	// properties:
	MatchDesignatedRequirement() objc.IObject /* cross-framework: NSString */
	MatchDomains() objc.IObject /* cross-framework: NSArray */
	SetMatchDomains(value objc.IObject /* cross-framework: NSArray */)
	MatchPath() objc.IObject /* cross-framework: NSString */
	SetMatchPath(value objc.IObject /* cross-framework: NSString */)
	MatchSigningIdentifier() objc.IObject /* cross-framework: NSString */
	MatchTools() []NEAppRule
	SetMatchTools(value []NEAppRule)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for NEAppRule */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for NEAppRule */
// Alloc allocates a new instance without initialization.
func (nc _NEAppRuleClass) Alloc() NEAppRule {
	rv := objc.Send[NEAppRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (nc _NEAppRuleClass) New() NEAppRule {
	rv := objc.Send[NEAppRule](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NEAppRule) Init() NEAppRule {
	rv := objc.Send[NEAppRule](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NEAppRule) Autorelease() NEAppRule {
	rv := objc.Send[NEAppRule](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEAppRule creates a new NEAppRule instance.
func NewNEAppRule() NEAppRule {
	return getNEAppRuleClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for NEAppRule */
// The identity of an app whose traffic is to be routed through the tunnel.


// The identity of an app whose traffic is to be routed through the tunnel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule
type NEAppRule struct {
	objectivec.Object
}

// NEAppRuleFrom constructs a [NEAppRule] from an unsafe.Pointer.
//
// The identity of an app whose traffic is to be routed through the tunnel.
func NEAppRuleFrom(ptr unsafe.Pointer) NEAppRule {
	return NEAppRule{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for NEAppRule */

// Create an app rule that matches an app with a given signing identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:)
func NewNEAppRuleWithSigningIdentifier(signingIdentifier objc.IObject /* cross-framework: NSString */) NEAppRule {
	instance := getNEAppRuleClass().Alloc()
	rv := objc.Send[NEAppRule](instance.ID, objc.Sel("initWithSigningIdentifier:"), signingIdentifier)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEAppRuleWithSigningIdentifier */


// Create an app rule that matches an app with a given signing identifier and a given designated requirement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:designatedRequirement:)
func NewNEAppRuleWithSigningIdentifierDesignatedRequirement(signingIdentifier objc.IObject /* cross-framework: NSString */, designatedRequirement objc.IObject /* cross-framework: NSString */) NEAppRule {
	instance := getNEAppRuleClass().Alloc()
	rv := objc.Send[NEAppRule](instance.ID, objc.Sel("initWithSigningIdentifier:designatedRequirement:"), signingIdentifier, designatedRequirement)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewNEAppRuleWithSigningIdentifierDesignatedRequirement */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for NEAppRule */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for NEAppRule */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for NEAppRule */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for NEAppRule */

// The designated requirement of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDesignatedRequirement
func (n_ NEAppRule) MatchDesignatedRequirement() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchDesignatedRequirement"))
	return rv
}/* debug [instance_properties/getter]: matchDesignatedRequirement */


// The hostname domains that match the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDomains
func (n_ NEAppRule) MatchDomains() objc.IObject /* cross-framework: NSArray */ {
	rv := objc.Send[foundation.NSArray](n_.ID, objc.Sel("matchDomains"))
	return rv
}/* debug [instance_properties/getter]: matchDomains */


// The hostname domains that match the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDomains
func (n_ NEAppRule) SetMatchDomains(value objc.IObject /* cross-framework: NSArray */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}/* debug [instance_properties/setter]: matchDomains */


// The file system path of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchPath
func (n_ NEAppRule) MatchPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchPath"))
	return rv
}/* debug [instance_properties/getter]: matchPath */


// The file system path of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchPath
func (n_ NEAppRule) SetMatchPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchPath:"), value)
}/* debug [instance_properties/setter]: matchPath */


// The signing identifier of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchSigningIdentifier
func (n_ NEAppRule) MatchSigningIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchSigningIdentifier"))
	return rv
}/* debug [instance_properties/getter]: matchSigningIdentifier */


// An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchTools
func (n_ NEAppRule) MatchTools() []NEAppRule {
	rv := objc.Send[[]NEAppRule](n_.ID, objc.Sel("matchTools"))
	return rv
}/* debug [instance_properties/getter]: matchTools */


// An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchTools
func (n_ NEAppRule) SetMatchTools(value []NEAppRule) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchTools:"), nsArray)
}/* debug [instance_properties/setter]: matchTools */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NEAppRule */


