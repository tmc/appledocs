// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

// PDOMEventListener is the DOMEventListener protocol interface.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/DOMEventListener
type PDOMEventListener interface {
	// Required methods
	HandleEvent(event IDOMEvent)
}
