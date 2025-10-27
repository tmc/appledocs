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
	

	// properties:
	MatchDesignatedRequirement() foundation.foundation.INSString
	MatchDomains() foundation.foundation.INSArray
	SetMatchDomains(value foundation.foundation.INSArray)
	MatchPath() foundation.foundation.INSString
	SetMatchPath(value foundation.foundation.INSString)
	MatchSigningIdentifier() foundation.foundation.INSString
	MatchTools() []NEAppRule
	SetMatchTools(value []NEAppRule)


	

	// methods:


}





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






// Create an app rule that matches an app with a given signing identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:)
func NewNEAppRuleWithSigningIdentifier(signingIdentifier foundation.foundation.INSString) NEAppRule {
	instance := getNEAppRuleClass().Alloc()
	rv := objc.Send[NEAppRule](instance.ID, objc.Sel("initWithSigningIdentifier:"), signingIdentifier)
	rv.Autorelease()
	return rv
}


// Create an app rule that matches an app with a given signing identifier and a given designated requirement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/init(signingIdentifier:designatedRequirement:)
func NewNEAppRuleWithSigningIdentifierDesignatedRequirement(signingIdentifier foundation.foundation.INSString, designatedRequirement foundation.foundation.INSString) NEAppRule {
	instance := getNEAppRuleClass().Alloc()
	rv := objc.Send[NEAppRule](instance.ID, objc.Sel("initWithSigningIdentifier:designatedRequirement:"), signingIdentifier, designatedRequirement)
	rv.Autorelease()
	return rv
}






















// The designated requirement of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDesignatedRequirement
func (n_ NEAppRule) MatchDesignatedRequirement() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchDesignatedRequirement"))
	return rv
}


// The hostname domains that match the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDomains
func (n_ NEAppRule) MatchDomains() foundation.foundation.INSArray {
	rv := objc.Send[foundation.NSArray](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// The hostname domains that match the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchDomains
func (n_ NEAppRule) SetMatchDomains(value foundation.foundation.INSArray) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}


// The file system path of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchPath
func (n_ NEAppRule) MatchPath() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchPath"))
	return rv
}


// The file system path of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchPath
func (n_ NEAppRule) SetMatchPath(value foundation.foundation.INSString) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchPath:"), value)
}


// The signing identifier of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchSigningIdentifier
func (n_ NEAppRule) MatchSigningIdentifier() foundation.foundation.INSString {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchSigningIdentifier"))
	return rv
}


// An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEAppRule/matchTools
func (n_ NEAppRule) MatchTools() []NEAppRule {
	rv := objc.Send[[]NEAppRule](n_.ID, objc.Sel("matchTools"))
	return rv
}


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
}







