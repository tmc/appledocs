// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [NENetworkRule] class.
var (
	NENetworkRuleClass     _NENetworkRuleClass
	NENetworkRuleClassOnce sync.Once
)

func getNENetworkRuleClass() _NENetworkRuleClass {
	NENetworkRuleClassOnce.Do(func() {
		NENetworkRuleClass = _NENetworkRuleClass{objc.GetClass("NENetworkRule")}
	})
	return NENetworkRuleClass
}

type _NENetworkRuleClass struct {
	class objc.Class
}

// An interface definition for the [NENetworkRule] class.
type INENetworkRule interface {
	objectivec.IObject
	MatchDirection() NETrafficDirection
	MatchLocalNetwork() NWHostEndpoint
	MatchLocalNetworkEndpoint() unsafe.Pointer
	MatchLocalPrefix() uint
	MatchProtocol() NENetworkRuleProtocol
	MatchRemoteEndpoint() NWHostEndpoint
	MatchRemoteHostOrNetworkEndpoint() unsafe.Pointer
	MatchRemotePrefix() uint
}

// A rule to match attributes of network traffic.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule
type NENetworkRule struct {
	objectivec.Object
}

// NENetworkRuleFrom constructs a [NENetworkRule] from an unsafe.Pointer.
//
// A rule to match attributes of network traffic.
func NENetworkRuleFrom(ptr unsafe.Pointer) NENetworkRule {
	return NENetworkRule{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (nc _NENetworkRuleClass) Alloc() NENetworkRule {
	rv := objc.Send[NENetworkRule](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (nc _NENetworkRuleClass) New() NENetworkRule {
	rv := objc.Send[NENetworkRule](objc.ID(nc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (n_ NENetworkRule) Init() NENetworkRule {
	rv := objc.Send[NENetworkRule](n_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n_ NENetworkRule) Autorelease() NENetworkRule {
	rv := objc.Send[NENetworkRule](n_.ID, objc.Sel("autorelease"))
	return rv
}

// NewNENetworkRule creates a new NENetworkRule instance.
func NewNENetworkRule() NENetworkRule {
	return getNENetworkRuleClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/initWithDestinationHostEndpoint:protocol:
func NewNENetworkRuleWithDestinationHostEndpointProtocol(hostEndpoint unsafe.Pointer, protocol_ INENetworkRuleProtocol) NENetworkRule {
	instance := getNENetworkRuleClass().Alloc()
	rv := objc.Send[NENetworkRule](instance.ID, objc.Sel("initWithDestinationHostEndpoint:protocol:"), hostEndpoint, protocol_)
	rv.Autorelease()
	return rv
}



// Creates a rule that matches network traffic destined for a host within a specific DNS domain.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/init(destinationHost:protocol:)
func NewNENetworkRuleWithDestinationHostProtocol(hostEndpoint INWHostEndpoint, protocol_ INENetworkRuleProtocol) NENetworkRule {
	instance := getNENetworkRuleClass().Alloc()
	rv := objc.Send[NENetworkRule](instance.ID, objc.Sel("initWithDestinationHost:protocol:"), hostEndpoint, protocol_)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/initWithDestinationNetworkEndpoint:prefix:protocol:
func NewNENetworkRuleWithDestinationNetworkEndpointPrefixProtocol(networkEndpoint unsafe.Pointer, destinationPrefix uint, protocol_ INENetworkRuleProtocol) NENetworkRule {
	instance := getNENetworkRuleClass().Alloc()
	rv := objc.Send[NENetworkRule](instance.ID, objc.Sel("initWithDestinationNetworkEndpoint:prefix:protocol:"), networkEndpoint, destinationPrefix, protocol_)
	rv.Autorelease()
	return rv
}



// Creates a rule that matches network traffic destined for a host within a specific network.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/init(destinationNetwork:prefix:protocol:)
func NewNENetworkRuleWithDestinationNetworkPrefixProtocol(networkEndpoint INWHostEndpoint, destinationPrefix uint, protocol_ INENetworkRuleProtocol) NENetworkRule {
	instance := getNENetworkRuleClass().Alloc()
	rv := objc.Send[NENetworkRule](instance.ID, objc.Sel("initWithDestinationNetwork:prefix:protocol:"), networkEndpoint, destinationPrefix, protocol_)
	rv.Autorelease()
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/initWithRemoteNetworkEndpoint:remotePrefix:localNetworkEndpoint:localPrefix:protocol:direction:
func NewNENetworkRuleWithRemoteNetworkEndpointRemotePrefixLocalNetworkEndpointLocalPrefixProtocolDirection(remoteNetwork unsafe.Pointer, remotePrefix uint, localNetwork unsafe.Pointer, localPrefix uint, protocol_ INENetworkRuleProtocol, direction NETrafficDirection) NENetworkRule {
	instance := getNENetworkRuleClass().Alloc()
	rv := objc.Send[NENetworkRule](instance.ID, objc.Sel("initWithRemoteNetworkEndpoint:remotePrefix:localNetworkEndpoint:localPrefix:protocol:direction:"), remoteNetwork, remotePrefix, localNetwork, localPrefix, protocol_, direction)
	rv.Autorelease()
	return rv
}



// Creates a rule that matches traffic by remote network, local network, protocol, and direction.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/init(remoteNetwork:remotePrefix:localNetwork:localPrefix:protocol:direction:)
func NewNENetworkRuleWithRemoteNetworkRemotePrefixLocalNetworkLocalPrefixProtocolDirection(remoteNetwork INWHostEndpoint, remotePrefix uint, localNetwork INWHostEndpoint, localPrefix uint, protocol_ INENetworkRuleProtocol, direction NETrafficDirection) NENetworkRule {
	instance := getNENetworkRuleClass().Alloc()
	rv := objc.Send[NENetworkRule](instance.ID, objc.Sel("initWithRemoteNetwork:remotePrefix:localNetwork:localPrefix:protocol:direction:"), remoteNetwork, remotePrefix, localNetwork, localPrefix, protocol_, direction)
	rv.Autorelease()
	return rv
}


// The direction of network traffic that the rule matches.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchDirection
func (n_ NENetworkRule) MatchDirection() NETrafficDirection {
	rv := objc.Send[NETrafficDirection](n_.ID, objc.Sel("matchDirection"))
	return rv
}

// The local network that the rule matches.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchLocalNetwork
func (n_ NENetworkRule) MatchLocalNetwork() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](n_.ID, objc.Sel("matchLocalNetwork"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchLocalNetworkEndpoint-9dyor
func (n_ NENetworkRule) MatchLocalNetworkEndpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("matchLocalNetworkEndpoint"))
	return rv
}

// A number that specifies the local sub-network that the rule matches.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchLocalPrefix
func (n_ NENetworkRule) MatchLocalPrefix() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("matchLocalPrefix"))
	return rv
}

// The protocol that the rule matches.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchProtocol
func (n_ NENetworkRule) MatchProtocol() NENetworkRuleProtocol {
	rv := objc.Send[NENetworkRuleProtocol](n_.ID, objc.Sel("matchProtocol"))
	return rv
}

// The remote endpoint that the rule matches.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchRemoteEndpoint
func (n_ NENetworkRule) MatchRemoteEndpoint() NWHostEndpoint {
	rv := objc.Send[NWHostEndpoint](n_.ID, objc.Sel("matchRemoteEndpoint"))
	return rv
}

//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchRemoteHostOrNetworkEndpoint-80s0l
func (n_ NENetworkRule) MatchRemoteHostOrNetworkEndpoint() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n_.ID, objc.Sel("matchRemoteHostOrNetworkEndpoint"))
	return rv
}

// A number that specifies the remote sub-network that the rule matches.
//
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NENetworkRule/matchRemotePrefix
func (n_ NENetworkRule) MatchRemotePrefix() uint {
	rv := objc.Send[uint](n_.ID, objc.Sel("matchRemotePrefix"))
	return rv
}


