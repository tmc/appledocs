// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// shouldConnectExternalDomainWithCompletionHandlerProtocol is the shouldConnectExternalDomainWithCompletionHandler: protocol.
//
// Availability:
//   - macOS 15.0+
//
// Use this protocol when registering custom classes that conform to shouldConnectExternalDomainWithCompletionHandler:.
var shouldConnectExternalDomainWithCompletionHandlerProtocol *objc.Protocol

func init() {
	shouldConnectExternalDomainWithCompletionHandlerProtocol = objc.GetProtocol("shouldConnectExternalDomainWithCompletionHandler:")
}
