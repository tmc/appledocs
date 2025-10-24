// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PURLSchemeTask is the WKURLSchemeTask protocol interface.
//
// An interface that WebKit uses to request custom resources from your app.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - visionOS 1.0+
//
// See: doc://com.apple.webkit/documentation/WebKit/WKURLSchemeTask
type PURLSchemeTask interface {
	// Required methods
	DidFailWithError(error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: DidFailWithError */
	DidFinish()/* debug [protocol_interface/required_method]: DidFinish */
	DidReceiveResponse(response foundation.URLResponse)/* debug [protocol_interface/required_method]: DidReceiveResponse */
	DidReceiveData(data objc.IObject /* cross-framework: NSData */)/* debug [protocol_interface/required_method]: DidReceiveData */
}
