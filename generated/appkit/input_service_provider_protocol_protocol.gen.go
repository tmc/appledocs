// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/corefoundation"
)

// PInputServiceProvider is the NSInputServiceProvider protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSInputServiceProvider
type PInputServiceProvider interface {
	// Required methods
	ActiveConversationChangedToNewConversation(sender objc.IObject, newConversation int)/* debug [protocol_interface/required_method]: ActiveConversationChangedToNewConversation */
	ActiveConversationWillChangeFromOldConversation(sender objc.IObject, oldConversation int)/* debug [protocol_interface/required_method]: ActiveConversationWillChangeFromOldConversation */
	CanBeDisabled() bool/* debug [protocol_interface/required_method]: CanBeDisabled */
	DoCommandBySelectorClient(selector objc.SEL, sender objc.IObject)/* debug [protocol_interface/required_method]: DoCommandBySelectorClient */
	InputClientBecomeActive(sender objc.IObject)/* debug [protocol_interface/required_method]: InputClientBecomeActive */
	InputClientDisabled(sender objc.IObject)/* debug [protocol_interface/required_method]: InputClientDisabled */
	InputClientEnabled(sender objc.IObject)/* debug [protocol_interface/required_method]: InputClientEnabled */
	InputClientResignActive(sender objc.IObject)/* debug [protocol_interface/required_method]: InputClientResignActive */
	InsertTextClient(string_ objc.IObject, sender objc.IObject)/* debug [protocol_interface/required_method]: InsertTextClient */
	MarkedTextAbandoned(sender objc.IObject)/* debug [protocol_interface/required_method]: MarkedTextAbandoned */
	MarkedTextSelectionChangedClient(newSel corefoundation.Range, sender objc.IObject)/* debug [protocol_interface/required_method]: MarkedTextSelectionChangedClient */
	Terminate(sender objc.IObject)/* debug [protocol_interface/required_method]: Terminate */
	WantsToDelayTextChangeNotifications() bool/* debug [protocol_interface/required_method]: WantsToDelayTextChangeNotifications */
	WantsToHandleMouseEvents() bool/* debug [protocol_interface/required_method]: WantsToHandleMouseEvents */
	WantsToInterpretAllKeystrokes() bool/* debug [protocol_interface/required_method]: WantsToInterpretAllKeystrokes */
}
