// Code generated from Apple documentation for AudioToolbox. DO NOT EDIT.

package audiotoolbox

import (

	"github.com/tmc/appledocs/generated/appkit"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PCocoaUIBase is the AUCocoaUIBase protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.audiotoolbox/documentation/AudioToolbox/AUCocoaUIBase
type PCocoaUIBase interface {
	// Required methods
	InterfaceVersion() objectivec.IObject/* debug [protocol_interface/required_method]: InterfaceVersion */
	UiViewForAudioUnitWithSize(inAudioUnit AudioUnit /* typedef */, inPreferredSize Size /* not a class type */) appkit.View/* debug [protocol_interface/required_method]: UiViewForAudioUnitWithSize */
}
