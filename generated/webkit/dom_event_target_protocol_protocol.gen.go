// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// PDOMEventTarget is the DOMEventTarget protocol interface.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.14)
//
// See: doc://com.apple.webkit/documentation/WebKit/DOMEventTarget
type PDOMEventTarget interface {
	// Required methods
	AddEventListenerListenerUseCapture(type_ objc.IObject /* cross-framework: NSString */, listener unsafe.Pointer, useCapture bool)
	DispatchEvent(event IDOMEvent) bool
	RemoveEventListenerListenerUseCapture(type_ objc.IObject /* cross-framework: NSString */, listener unsafe.Pointer, useCapture bool)
}
