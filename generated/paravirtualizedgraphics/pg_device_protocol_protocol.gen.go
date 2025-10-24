// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PPGDevice is the PGDevice protocol interface.
//
// A paravirtualized GPU device object.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// See: doc://com.apple.paravirtualizedgraphics/documentation/ParavirtualizedGraphics/PGDevice
type PPGDevice interface {
	// Required methods
	DidResume()/* debug [protocol_interface/required_method]: DidResume */
	FinishSuspend() foundation.Data/* debug [protocol_interface/required_method]: FinishSuspend */
	MmioReadAtOffset(offset uintptr /* not a class type */) uint32/* debug [protocol_interface/required_method]: MmioReadAtOffset */
	MmioWriteAtOffsetValue(offset uintptr /* not a class type */, value uint32 /* not a class type */)/* debug [protocol_interface/required_method]: MmioWriteAtOffsetValue */
	NewDisplayWithDescriptorPortSerialNum(descriptor IPGDisplayDescriptor, port uint, serialNum uint32 /* not a class type */) unsafe.Pointer/* debug [protocol_interface/required_method]: NewDisplayWithDescriptorPortSerialNum */
	Pause()/* debug [protocol_interface/required_method]: Pause */
	Reset()/* debug [protocol_interface/required_method]: Reset */
	Stop()/* debug [protocol_interface/required_method]: Stop */
	Unpause()/* debug [protocol_interface/required_method]: Unpause */
	WillResumeWithSuspendStateError(suspendState objc.IObject /* cross-framework: NSData */, error_ unsafe.Pointer) bool/* debug [protocol_interface/required_method]: WillResumeWithSuspendStateError */
	WillSuspend()/* debug [protocol_interface/required_method]: WillSuspend */
}
