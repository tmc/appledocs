// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PInputServiceProvider is the NSInputServiceProvider protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSInputServiceProvider
type PInputServiceProvider interface {
	// Required methods
	ActiveConversationChangedToNewConversation(sender objc.IObject, newConversation int)
	ActiveConversationWillChangeFromOldConversation(sender objc.IObject, oldConversation int)
	CanBeDisabled() bool
	DoCommandBySelectorClient(selector objc.SEL, sender objc.IObject)
	InputClientBecomeActive(sender objc.IObject)
	InputClientDisabled(sender objc.IObject)
	InputClientEnabled(sender objc.IObject)
	InputClientResignActive(sender objc.IObject)
	InsertTextClient(string_ objc.IObject, sender objc.IObject)
	MarkedTextAbandoned(sender objc.IObject)
	MarkedTextSelectionChangedClient(newSel corefoundation.Range, sender objc.IObject)
	Terminate(sender objc.IObject)
	WantsToDelayTextChangeNotifications() bool
	WantsToHandleMouseEvents() bool
	WantsToInterpretAllKeystrokes() bool
}
