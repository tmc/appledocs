// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension_test

import (
	"github.com/tmc/appledocs/generated/networkextension"
)

// Suppress unused import errors
var _ = networkextension.NewNENetworkRule

// ExampleNewNENetworkRuleWithDestinationHostProtocol demonstrates how to create a NENetworkRule instance using NewNENetworkRuleWithDestinationHostProtocol.
// Creates a rule that matches network traffic destined for a host within a specific DNS domain.
func ExampleNewNENetworkRuleWithDestinationHostProtocol() {
	_ = networkextension.NewNENetworkRuleWithDestinationHostProtocol(
		networkextension.NWHostEndpoint{}, // hostEndpoint NWHostEndpoint
		networkextension.NENetworkRuleProtocol{}, // protocol NENetworkRuleProtocol
	)
	// Output:
}
// ExampleNewNENetworkRuleWithDestinationNetworkPrefixProtocol demonstrates how to create a NENetworkRule instance using NewNENetworkRuleWithDestinationNetworkPrefixProtocol.
// Creates a rule that matches network traffic destined for a host within a specific network.
func ExampleNewNENetworkRuleWithDestinationNetworkPrefixProtocol() {
	_ = networkextension.NewNENetworkRuleWithDestinationNetworkPrefixProtocol(
		networkextension.NWHostEndpoint{}, // networkEndpoint NWHostEndpoint
		0, // destinationPrefix uint
		networkextension.NENetworkRuleProtocol{}, // protocol NENetworkRuleProtocol
	)
	// Output:
}
// ExampleNewNENetworkRuleWithRemoteNetworkRemotePrefixLocalNetworkLocalPrefixProtocolDirection demonstrates how to create a NENetworkRule instance using NewNENetworkRuleWithRemoteNetworkRemotePrefixLocalNetworkLocalPrefixProtocolDirection.
// Creates a rule that matches traffic by remote network, local network, protocol, and direction.
func ExampleNewNENetworkRuleWithRemoteNetworkRemotePrefixLocalNetworkLocalPrefixProtocolDirection() {
	_ = networkextension.NewNENetworkRuleWithRemoteNetworkRemotePrefixLocalNetworkLocalPrefixProtocolDirection(
		networkextension.NWHostEndpoint{}, // remoteNetwork NWHostEndpoint
		0, // remotePrefix uint
		networkextension.NWHostEndpoint{}, // localNetwork NWHostEndpoint
		0, // localPrefix uint
		networkextension.NENetworkRuleProtocol{}, // protocol NENetworkRuleProtocol
		networkextension.NETrafficDirection{}, // direction NETrafficDirection
	)
	// Output:
}
