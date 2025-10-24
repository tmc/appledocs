// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"
)

// PDOMNodeFilter is the DOMNodeFilter protocol interface.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/DOMNodeFilter
type PDOMNodeFilter interface {
	// Required methods
	AcceptNode(n IDOMNode) unsafe.Pointer
}
