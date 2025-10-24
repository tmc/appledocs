// Code generated from Apple documentation for Quartz. DO NOT EDIT.

package quartz

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"
)

// PQCCompositionRenderer is the QCCompositionRenderer protocol interface.
//
// The   protocol defines the methods used to pass data to the input ports or retrieve data from the output ports of the root patch of a Quartz Composer composition. This protocol is adopted by the  ,  , and   classes.
//
// Availability:
//   - macOS 10.4+ (Deprecated in 10.15)
//
// See: doc://com.apple.quartz/documentation/Quartz/QCCompositionRenderer
type PQCCompositionRenderer interface {
	// Required methods
	Attributes() foundation.Dictionary/* debug [protocol_interface/required_method]: Attributes */
	InputKeys() foundation.Array/* debug [protocol_interface/required_method]: InputKeys */
	OutputKeys() foundation.Array/* debug [protocol_interface/required_method]: OutputKeys */
	PropertyListFromInputValues() objc.ID/* debug [protocol_interface/required_method]: PropertyListFromInputValues */
	SetInputValuesWithPropertyList(plist objc.IObject)/* debug [protocol_interface/required_method]: SetInputValuesWithPropertyList */
	SetValueForInputKey(value objc.IObject, key objc.IObject /* cross-framework: NSString */) bool/* debug [protocol_interface/required_method]: SetValueForInputKey */
	UserInfo() foundation.MutableDictionary/* debug [protocol_interface/required_method]: UserInfo */
	ValueForInputKey(key objc.IObject /* cross-framework: NSString */) objc.ID/* debug [protocol_interface/required_method]: ValueForInputKey */
	ValueForOutputKey(key objc.IObject /* cross-framework: NSString */) objc.ID/* debug [protocol_interface/required_method]: ValueForOutputKey */
	ValueForOutputKeyOfType(key objc.IObject /* cross-framework: NSString */, type_ objc.IObject /* cross-framework: NSString */) objc.ID/* debug [protocol_interface/required_method]: ValueForOutputKeyOfType */
}
