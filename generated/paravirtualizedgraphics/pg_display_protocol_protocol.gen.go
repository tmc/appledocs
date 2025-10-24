// Code generated from Apple documentation for ParavirtualizedGraphics. DO NOT EDIT.

package paravirtualizedgraphics

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/corelocation"
)

// PPGDisplay is the PGDisplay protocol interface.
//
// An object that provides display functionality to the guest operating system in a way that the host-side virtual machine app can intercept.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - macOS 11.0+
//
// See: doc://com.apple.paravirtualizedgraphics/documentation/ParavirtualizedGraphics/PGDisplay
type PPGDisplay interface {
	// Required methods
	EncodeCurrentFrameToCommandBufferTextureRegion(commandBuffer unsafe.Pointer, texture unsafe.Pointer, region corelocation.Region) bool/* debug [protocol_interface/required_method]: EncodeCurrentFrameToCommandBufferTextureRegion */
}
