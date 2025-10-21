// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [NEAppRule] class.
type INEAppRule interface {
	objectivec.IObject
}

// The identity of an app whose traffic is to be routed through the tunnel.
//
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

// Alloc allocates a new instance without initialization.
func (nc _NEAppRuleClass) Alloc() NEAppRule {
	rv := objc.Send[NEAppRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Create an app rule that matches an app with a given signing identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:)
func NewNEAppRuleWithSigningIdentifier(signingIdentifier string) NEAppRule {
	instance := getNEAppRuleClass().Alloc()
	rv := objc.Send[NEAppRule](instance.ID, objc.Sel("initWithSigningIdentifier:"), objc.String(signingIdentifier))
	rv.Autorelease()
	return rv
}



// Create an app rule that matches an app with a given signing identifier and a given designated requirement.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:designatedRequirement:)
func NewNEAppRuleWithSigningIdentifierDesignatedRequirement(signingIdentifier string, designatedRequirement string) NEAppRule {
	instance := getNEAppRuleClass().Alloc()
	rv := objc.Send[NEAppRule](instance.ID, objc.Sel("initWithSigningIdentifier:designatedRequirement:"), objc.String(signingIdentifier), objc.String(designatedRequirement))
	rv.Autorelease()
	return rv
}


// The designated requirement of the app that matches the rule.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDesignatedRequirement
func (n_ NEAppRule) MatchDesignatedRequirement() string {
	rv := objc.Send[string](n_.ID, objc.Sel("matchDesignatedRequirement"))
	return rv
}

// The hostname domains that match the rule.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDomains
func (n_ NEAppRule) MatchDomains() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// SetMatchDomains sets the value of the matchDomains property.
// The hostname domains that match the rule.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDomains
func (n_ NEAppRule) SetMatchDomains(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}

// The file system path of the app that matches the rule.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchPath
func (n_ NEAppRule) MatchPath() string {
	rv := objc.Send[string](n_.ID, objc.Sel("matchPath"))
	return rv
}


// SetMatchPath sets the value of the matchPath property.
// The file system path of the app that matches the rule.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchPath
func (n_ NEAppRule) SetMatchPath(value string) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchPath:"), objc.String(value))
}

// The signing identifier of the app that matches the rule.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchSigningIdentifier
func (n_ NEAppRule) MatchSigningIdentifier() string {
	rv := objc.Send[string](n_.ID, objc.Sel("matchSigningIdentifier"))
	return rv
}

// An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchTools
func (n_ NEAppRule) MatchTools() []NEAppRule {
	rv := objc.Send[[]NEAppRule](n_.ID, objc.Sel("matchTools"))
	return rv
}


// SetMatchTools sets the value of the matchTools property.
// An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchTools
func (n_ NEAppRule) SetMatchTools(value []NEAppRule) {
	// Convert Go slice to NSArray
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
}


