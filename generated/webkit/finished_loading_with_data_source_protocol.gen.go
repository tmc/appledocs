// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// finishedLoadingWithDataSourceProtocol is the finishedLoadingWithDataSource: protocol.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// Use this protocol when registering custom classes that conform to finishedLoadingWithDataSource:.
var finishedLoadingWithDataSourceProtocol *objc.Protocol

func init() {
	finishedLoadingWithDataSourceProtocol = objc.GetProtocol("finishedLoadingWithDataSource:")
}
