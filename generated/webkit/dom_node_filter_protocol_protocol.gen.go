// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (

	"github.com/tmc/appledocs/generated/objectivec"
)

// PDOMNodeFilter is the DOMNodeFilter protocol interface.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/DOMNodeFilter
type PDOMNodeFilter interface {
	// Required methods
	AcceptNode(n IDOMNode) objectivec.IObject/* debug [protocol_interface/required_method]: AcceptNode */
}
