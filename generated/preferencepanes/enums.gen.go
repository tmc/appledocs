// Code generated from Apple documentation for PreferencePanes. DO NOT EDIT.

package preferencepanes

/* debug [enums.gen.go]: Generating 1 enums for PreferencePanes */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum NSPreferencePaneUnselectReply (3 cases) */
// NSPreferencePaneUnselectReply - Constants that indicate the preference pane’s availability to be deselected.
//
// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePaneUnselectReply
type NSPreferencePaneUnselectReply uint

const (
	// NSUnselectCancel - Cancel the deselection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePaneUnselectReply/unselectCancel
	NSUnselectCancel NSPreferencePaneUnselectReply = 0
	// NSUnselectLater - Delay the deselection until the preference pane invokes  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePaneUnselectReply/unselectLater
	NSUnselectLater NSPreferencePaneUnselectReply = 2
	// NSUnselectNow - Continue the deselection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/PreferencePanes/NSPreferencePaneUnselectReply/unselectNow
	NSUnselectNow NSPreferencePaneUnselectReply = 1
)


