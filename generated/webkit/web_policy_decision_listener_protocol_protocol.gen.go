// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

// PWebPolicyDecisionListener is the WebPolicyDecisionListener protocol interface.
//
// This protocol enables   policy delegates to communicate with listener objects. A listener object conforming to this protocol is passed as one of the arguments to web view policy delegate methods.
//
// Availability:
//   - macOS 10.3+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/WebPolicyDecisionListener
type PWebPolicyDecisionListener interface {
	// Required methods
	Download()/* debug [protocol_interface/required_method]: Download */
	Ignore()/* debug [protocol_interface/required_method]: Ignore */
	Use()/* debug [protocol_interface/required_method]: Use */
}
