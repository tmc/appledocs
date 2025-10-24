// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	MatchDesignatedRequirement() objc.IObject /* cross-framework: NSString */
	SetMatchDesignatedRequirement(value objc.IObject /* cross-framework: NSString */)
	MatchDomains() unsafe.Pointer
	SetMatchDomains(value unsafe.Pointer)
	MatchPath() objc.IObject /* cross-framework: NSString */
	SetMatchPath(value objc.IObject /* cross-framework: NSString */)
	MatchSigningIdentifier() objc.IObject /* cross-framework: NSString */
	SetMatchSigningIdentifier(value objc.IObject /* cross-framework: NSString */)
	MatchTools() INEAppRule
	SetMatchTools(value INEAppRule)
	// methods:
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



// The designated requirement of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchdesignatedrequirement
func (n_ NEAppRule) MatchDesignatedRequirement() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchDesignatedRequirement"))
	return rv
}


// The designated requirement of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchdesignatedrequirement
func (n_ NEAppRule) SetMatchDesignatedRequirement(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDesignatedRequirement:"), value)
}


// The hostname domains that match the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchdomains
func (n_ NEAppRule) MatchDomains() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("matchDomains"))
	return rv
}


// The hostname domains that match the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchdomains
func (n_ NEAppRule) SetMatchDomains(value unsafe.Pointer) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchDomains:"), value)
}


// The file system path of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchpath
func (n_ NEAppRule) MatchPath() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchPath"))
	return rv
}


// The file system path of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchpath
func (n_ NEAppRule) SetMatchPath(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchPath:"), value)
}


// The signing identifier of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchsigningidentifier
func (n_ NEAppRule) MatchSigningIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](n_.ID, objc.Sel("matchSigningIdentifier"))
	return rv
}


// The signing identifier of the app that matches the rule.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchsigningidentifier
func (n_ NEAppRule) SetMatchSigningIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchSigningIdentifier:"), value)
}


// An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchtools
func (n_ NEAppRule) MatchTools() INEAppRule {
	rv := objc.Send[NEAppRule](n_.ID, objc.Sel("matchTools"))
	return rv
}


// An array of app rule objects that restrict the rule so it only matches network traffic generated from helper processes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/networkextension/neapprule/matchtools
func (n_ NEAppRule) SetMatchTools(value INEAppRule) {
	objc.Send[objc.ID](n_.ID, objc.Sel("setMatchTools:"), value)
}



