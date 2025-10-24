//go:build darwin && ios

// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// iOS-only methods for NEHotspotHelperCommand


// Create a response to the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommand/createResponse(_:)
func (n_ NEHotspotHelperCommand) CreateResponse(result NEHotspotHelperResult) INEHotspotHelperResponse {
	rv := objc.Send[NEHotspotHelperResponse](n_.ID, objc.Sel("createResponse:"), result)
	return rv
}

// iOS-only properties

// The type of the command
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommand/commandType
func (n_ NEHotspotHelperCommand) CommandType() NEHotspotHelperCommandType {
	rv := objc.Send[NEHotspotHelperCommandType](n_.ID, objc.Sel("commandType"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommand/interface-7rt15
func (n_ NEHotspotHelperCommand) Interface() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](n_.ID, objc.Sel("interface"))
	return rv
}

// The network associated with the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommand/network
func (n_ NEHotspotHelperCommand) Network() INEHotspotNetwork {
	rv := objc.Send[NEHotspotNetwork](n_.ID, objc.Sel("network"))
	return rv
}

// The list of networks associated with the command.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/NetworkExtension/NEHotspotHelperCommand/networkList
func (n_ NEHotspotHelperCommand) NetworkList() []NEHotspotNetwork {
	rv := objc.Send[[]NEHotspotNetwork](n_.ID, objc.Sel("networkList"))
	return rv
}





