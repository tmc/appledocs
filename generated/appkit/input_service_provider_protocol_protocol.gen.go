// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/foundation"

	"github.com/tmc/appledocs/generated/objectivec"
)

// PInputServiceProvider is the NSInputServiceProvider protocol interface.
//
// Availability:
//   - macOS +
//
// See: doc://com.apple.appkit/documentation/AppKit/NSInputServiceProvider
type PInputServiceProvider interface {
	// Required methods
	ActiveConversationChangedToNewConversation(sender objectivec.IObject, newConversation int)
	ActiveConversationWillChangeFromOldConversation(sender objectivec.IObject, oldConversation int)
	CanBeDisabled() bool
	DoCommandBySelectorClient(selector objc.SEL, sender objectivec.IObject)
	InputClientBecomeActive(sender objectivec.IObject)
	InputClientDisabled(sender objectivec.IObject)
	InputClientEnabled(sender objectivec.IObject)
	InputClientResignActive(sender objectivec.IObject)
	InsertTextClient(string_ objectivec.IObject, sender objectivec.IObject)
	MarkedTextAbandoned(sender objectivec.IObject)
	MarkedTextSelectionChangedClient(newSel foundation.Range, sender objectivec.IObject)
	Terminate(sender objectivec.IObject)
	WantsToDelayTextChangeNotifications() bool
	WantsToHandleMouseEvents() bool
	WantsToInterpretAllKeystrokes() bool
}
