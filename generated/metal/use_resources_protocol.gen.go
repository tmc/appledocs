// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// useResourcesProtocol is the useResources: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.15+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to useResources:.
var useResourcesProtocol *objc.Protocol

func init() {
	useResourcesProtocol = objc.GetProtocol("useResources:")
}
