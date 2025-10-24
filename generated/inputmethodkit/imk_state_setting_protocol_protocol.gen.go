// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PIMKStateSetting is the IMKStateSetting protocol interface.
//
// The   protocol defines methods for setting or accessing values that indicate the state of an input method.
//
// Availability:
//   - macOS 10.5+
//
// See: doc://com.apple.inputmethodkit/documentation/InputMethodKit/IMKStateSetting
type PIMKStateSetting interface {
	// Required methods
	ActivateServer(sender objc.IObject)/* debug [protocol_interface/required_method]: ActivateServer */
	DeactivateServer(sender objc.IObject)/* debug [protocol_interface/required_method]: DeactivateServer */
	Modes(sender objc.IObject) foundation.Dictionary/* debug [protocol_interface/required_method]: Modes */
	RecognizedEvents(sender objc.IObject) uint/* debug [protocol_interface/required_method]: RecognizedEvents */
	SetValueForTagClient(value objc.IObject, tag unsafe.Pointer, sender objc.IObject)/* debug [protocol_interface/required_method]: SetValueForTagClient */
	ShowPreferences(sender objc.IObject)/* debug [protocol_interface/required_method]: ShowPreferences */
	ValueForTagClient(tag unsafe.Pointer, sender objc.IObject) objc.ID/* debug [protocol_interface/required_method]: ValueForTagClient */
}
