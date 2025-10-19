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
	nEAppRuleClass     _NEAppRuleClass
	nEAppRuleClassOnce sync.Once
)

func getNEAppRuleClass() _NEAppRuleClass {
	nEAppRuleClassOnce.Do(func() {
		nEAppRuleClass = _NEAppRuleClass{objc.GetClass("NEAppRule")}
	})
	return nEAppRuleClass
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




