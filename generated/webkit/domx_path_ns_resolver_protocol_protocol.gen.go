// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

// PDOMXPathNSResolver is the DOMXPathNSResolver protocol interface.
//
// Availability:
//   - macOS 10.5+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/DOMXPathNSResolver
type PDOMXPathNSResolver interface {
	// Required methods
	LookupNamespaceURI(prefix objc.IObject /* cross-framework: NSString */) foundation.String
}
