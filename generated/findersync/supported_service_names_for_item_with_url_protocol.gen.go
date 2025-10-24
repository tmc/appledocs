// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

package findersync

import "github.com/ebitengine/purego/objc"

// supportedServiceNamesForItemWithURLProtocol is the supportedServiceNamesForItemWithURL: protocol.
//
// Availability:
//   - macOS 10.10+
//
// Use this protocol when registering custom classes that conform to supportedServiceNamesForItemWithURL:.
var supportedServiceNamesForItemWithURLProtocol *objc.Protocol

func init() {
	supportedServiceNamesForItemWithURLProtocol = objc.GetProtocol("supportedServiceNamesForItemWithURL:")
}

