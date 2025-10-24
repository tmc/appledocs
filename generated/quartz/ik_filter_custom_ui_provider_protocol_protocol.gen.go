// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PIKFilterCustomUIProvider is the IKFilterCustomUIProvider protocol interface.
//
// A protocol used to provide a custom UI.
//
// Availability:
//   - macOS 10.4+
//
// See: doc://com.apple.quartz/documentation/Quartz/IKFilterCustomUIProvider
type PIKFilterCustomUIProvider interface {
	// Required methods
	ProvideViewForUIConfigurationExcludedKeys(inUIConfiguration objc.IObject /* cross-framework: NSDictionary */, inKeys objc.IObject /* cross-framework: NSArray */) IKFilterUIView/* debug [protocol_interface/required_method]: ProvideViewForUIConfigurationExcludedKeys */
}
