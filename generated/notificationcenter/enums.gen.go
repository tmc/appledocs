// Code generated from Apple documentation for NotificationCenter. DO NOT EDIT.

package notificationcenter

/* debug [enums.gen.go]: Generating 2 enums for NotificationCenter */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum NCUpdateResult (3 cases) */
// NCUpdateResult - The result of updating a widget’s state.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCUpdateResult
type NCUpdateResult uint

const (
	// NCUpdateResultFailed - The update attempt failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCUpdateResult/failed
	NCUpdateResultFailed NCUpdateResult = 0
	// NCUpdateResultNewData - The update resulted in new data to display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCUpdateResult/newData
	NCUpdateResultNewData NCUpdateResult = 0
	// NCUpdateResultNoData - The update did not result in any new data since the last update.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCUpdateResult/noData
	NCUpdateResultNoData NCUpdateResult = 0
)

/* debug [enums.gen.go]: Processing enum NCWidgetDisplayMode (2 cases) */
// NCWidgetDisplayMode - The modes that can be toggled between when the user activates the More button for a widget running in iOS.
//
// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetDisplayMode
type NCWidgetDisplayMode uint

const (
	// NCWidgetDisplayModeCompact - The current height of the widget is compact.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetDisplayMode/compact
	NCWidgetDisplayModeCompact NCWidgetDisplayMode = 0
	// NCWidgetDisplayModeExpanded - The current height of the widget is expanded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/NotificationCenter/NCWidgetDisplayMode/expanded
	NCWidgetDisplayModeExpanded NCWidgetDisplayMode = 0
)


