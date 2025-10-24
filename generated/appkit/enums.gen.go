// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

/* debug [enums.gen.go]: Generating 279 enums for AppKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum NSAccessibilityAnnotationPosition (3 cases) */
// AccessibilityAnnotationPosition - Constants that specify the position where the annotation applies.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition
type AccessibilityAnnotationPosition uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition/end
	AccessibilityAnnotationPositionEnd AccessibilityAnnotationPosition = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition/fullRange
	AccessibilityAnnotationPositionFullRange AccessibilityAnnotationPosition = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityAnnotationPosition/start
	AccessibilityAnnotationPositionStart AccessibilityAnnotationPosition = 1
)

/* debug [enums.gen.go]: Processing enum NSAccessibilityCustomRotorType (22 cases) */
// AccessibilityCustomRotorType - Constants that indicate the type of content that the rotor represents.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType
type AccessibilityCustomRotorType uint

const (
	// AccessibilityCustomRotorTypeAnnotation - An annotation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/annotation
	AccessibilityCustomRotorTypeAnnotation AccessibilityCustomRotorType = 2
	// AccessibilityCustomRotorTypeAny - Any type of item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/any
	AccessibilityCustomRotorTypeAny AccessibilityCustomRotorType = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/audiograph
	AccessibilityCustomRotorTypeAudiograph AccessibilityCustomRotorType = 21
	// AccessibilityCustomRotorTypeBoldText - Any bold text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/boldText
	AccessibilityCustomRotorTypeBoldText AccessibilityCustomRotorType = 3
	// AccessibilityCustomRotorTypeCustom - A rotor with a custom label.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/custom
	AccessibilityCustomRotorTypeCustom AccessibilityCustomRotorType = 0
	// AccessibilityCustomRotorTypeHeading - Any heading-level text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/heading
	AccessibilityCustomRotorTypeHeading AccessibilityCustomRotorType = 4
	// AccessibilityCustomRotorTypeHeadingLevel1 - A first-level heading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/headingLevel1
	AccessibilityCustomRotorTypeHeadingLevel1 AccessibilityCustomRotorType = 5
	// AccessibilityCustomRotorTypeHeadingLevel2 - A second-level heading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/headingLevel2
	AccessibilityCustomRotorTypeHeadingLevel2 AccessibilityCustomRotorType = 6
	// AccessibilityCustomRotorTypeHeadingLevel3 - A third-level heading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/headingLevel3
	AccessibilityCustomRotorTypeHeadingLevel3 AccessibilityCustomRotorType = 7
	// AccessibilityCustomRotorTypeHeadingLevel4 - A fourth-level heading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/headingLevel4
	AccessibilityCustomRotorTypeHeadingLevel4 AccessibilityCustomRotorType = 8
	// AccessibilityCustomRotorTypeHeadingLevel5 - A fifth-level heading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/headingLevel5
	AccessibilityCustomRotorTypeHeadingLevel5 AccessibilityCustomRotorType = 9
	// AccessibilityCustomRotorTypeHeadingLevel6 - A sixth-level heading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/headingLevel6
	AccessibilityCustomRotorTypeHeadingLevel6 AccessibilityCustomRotorType = 10
	// AccessibilityCustomRotorTypeImage - An image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/image
	AccessibilityCustomRotorTypeImage AccessibilityCustomRotorType = 11
	// AccessibilityCustomRotorTypeItalicText - Any italicized text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/italicText
	AccessibilityCustomRotorTypeItalicText AccessibilityCustomRotorType = 12
	// AccessibilityCustomRotorTypeLandmark - A landmark.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/landmark
	AccessibilityCustomRotorTypeLandmark AccessibilityCustomRotorType = 13
	// AccessibilityCustomRotorTypeLink - A link.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/link
	AccessibilityCustomRotorTypeLink AccessibilityCustomRotorType = 14
	// AccessibilityCustomRotorTypeList - A list of items.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/list
	AccessibilityCustomRotorTypeList AccessibilityCustomRotorType = 15
	// AccessibilityCustomRotorTypeMisspelledWord - A misspelled word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/misspelledWord
	AccessibilityCustomRotorTypeMisspelledWord AccessibilityCustomRotorType = 16
	// AccessibilityCustomRotorTypeTable - A table of information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/table
	AccessibilityCustomRotorTypeTable AccessibilityCustomRotorType = 17
	// AccessibilityCustomRotorTypeTextField - A text field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/textField
	AccessibilityCustomRotorTypeTextField AccessibilityCustomRotorType = 18
	// AccessibilityCustomRotorTypeUnderlinedText - Any underlined text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/underlinedText
	AccessibilityCustomRotorTypeUnderlinedText AccessibilityCustomRotorType = 19
	// AccessibilityCustomRotorTypeVisitedLink - A visited link.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/RotorType/visitedLink
	AccessibilityCustomRotorTypeVisitedLink AccessibilityCustomRotorType = 20
)

/* debug [enums.gen.go]: Processing enum NSAccessibilityCustomRotorSearchDirection (2 cases) */
// AccessibilityCustomRotorSearchDirection - Constants that describe the direction to search for an item result.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchDirection
type AccessibilityCustomRotorSearchDirection uint

const (
	// AccessibilityCustomRotorSearchDirectionNext - The next search item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchDirection/next
	AccessibilityCustomRotorSearchDirectionNext AccessibilityCustomRotorSearchDirection = 1
	// AccessibilityCustomRotorSearchDirectionPrevious - The previous search item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityCustomRotor/SearchDirection/previous
	AccessibilityCustomRotorSearchDirectionPrevious AccessibilityCustomRotorSearchDirection = 0
)

/* debug [enums.gen.go]: Processing enum NSAccessibilityOrientation (3 cases) */
// AccessibilityOrientation - Values that indicate the orientation of accessibility elements, such as scroll bars and split views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation
type AccessibilityOrientation uint

const (
	// AccessibilityOrientationHorizontal - The element is oriented horizontally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation/horizontal
	AccessibilityOrientationHorizontal AccessibilityOrientation = 2
	// AccessibilityOrientationUnknown - The element has unknown orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation/unknown
	AccessibilityOrientationUnknown AccessibilityOrientation = 0
	// AccessibilityOrientationVertical - The element is oriented vertically.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityOrientation/vertical
	AccessibilityOrientationVertical AccessibilityOrientation = 1
)

/* debug [enums.gen.go]: Processing enum NSAccessibilityPriorityLevel (3 cases) */
// AccessibilityPriorityLevel - A data type for notification priority levels.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityPriorityLevel
type AccessibilityPriorityLevel uint

const (
	// AccessibilityPriorityHigh - The notification is a high priority.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityPriorityLevel/high
	AccessibilityPriorityHigh AccessibilityPriorityLevel = 90
	// AccessibilityPriorityLow - The notification is a low priority.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityPriorityLevel/low
	AccessibilityPriorityLow AccessibilityPriorityLevel = 10
	// AccessibilityPriorityMedium - The notification is a medium priority.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityPriorityLevel/medium
	AccessibilityPriorityMedium AccessibilityPriorityLevel = 50
)

/* debug [enums.gen.go]: Processing enum NSAccessibilityRulerMarkerType (8 cases) */
// AccessibilityRulerMarkerType - Values that indicate the marker type of an accessibility element.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType
type AccessibilityRulerMarkerType uint

const (
	// AccessibilityRulerMarkerTypeIndentFirstLine - First line indent marker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/indentFirstLine
	AccessibilityRulerMarkerTypeIndentFirstLine AccessibilityRulerMarkerType = 7
	// AccessibilityRulerMarkerTypeIndentHead - Head indent marker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/indentHead
	AccessibilityRulerMarkerTypeIndentHead AccessibilityRulerMarkerType = 5
	// AccessibilityRulerMarkerTypeIndentTail - Tail indent marker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/indentTail
	AccessibilityRulerMarkerTypeIndentTail AccessibilityRulerMarkerType = 6
	// AccessibilityRulerMarkerTypeTabStopCenter - Center tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/tabStopCenter
	AccessibilityRulerMarkerTypeTabStopCenter AccessibilityRulerMarkerType = 3
	// AccessibilityRulerMarkerTypeTabStopDecimal - Decimal tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/tabStopDecimal
	AccessibilityRulerMarkerTypeTabStopDecimal AccessibilityRulerMarkerType = 4
	// AccessibilityRulerMarkerTypeTabStopLeft - Left tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/tabStopLeft
	AccessibilityRulerMarkerTypeTabStopLeft AccessibilityRulerMarkerType = 1
	// AccessibilityRulerMarkerTypeTabStopRight - Right tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/tabStopRight
	AccessibilityRulerMarkerTypeTabStopRight AccessibilityRulerMarkerType = 2
	// AccessibilityRulerMarkerTypeUnknown - Unknown marker type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType/unknown
	AccessibilityRulerMarkerTypeUnknown AccessibilityRulerMarkerType = 0
)

/* debug [enums.gen.go]: Processing enum NSAccessibilitySortDirection (3 cases) */
// AccessibilitySortDirection - Values that indicate the sort direction of a column.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortDirection
type AccessibilitySortDirection uint

const (
	// AccessibilitySortDirectionAscending - The column is sorted in ascending values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortDirection/ascending
	AccessibilitySortDirectionAscending AccessibilitySortDirection = 1
	// AccessibilitySortDirectionDescending - The column is sorted in descending values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortDirection/descending
	AccessibilitySortDirectionDescending AccessibilitySortDirection = 2
	// AccessibilitySortDirectionUnknown - The sort direction is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortDirection/unknown
	AccessibilitySortDirectionUnknown AccessibilitySortDirection = 0
)

/* debug [enums.gen.go]: Processing enum NSAccessibilityUnits (5 cases) */
// AccessibilityUnits - Values that indicate the unit values of a ruler or layout area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
type AccessibilityUnits uint

const (
	// AccessibilityUnitsCentimeters - The units are centimeters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits/centimeters
	AccessibilityUnitsCentimeters AccessibilityUnits = 2
	// AccessibilityUnitsInches - The units are inches.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits/inches
	AccessibilityUnitsInches AccessibilityUnits = 1
	// AccessibilityUnitsPicas - The units are picas.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits/picas
	AccessibilityUnitsPicas AccessibilityUnits = 4
	// AccessibilityUnitsPoints - The units are points.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits/points
	AccessibilityUnitsPoints AccessibilityUnits = 3
	// AccessibilityUnitsUnknown - The units are unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits/unknown
	AccessibilityUnitsUnknown AccessibilityUnits = 0
)

/* debug [enums.gen.go]: Processing enum NSAnimationBlockingMode (3 cases) */
// AnimationBlockingMode - These constants indicate the blocking mode of an 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode
type AnimationBlockingMode uint

const (
	// AnimationBlocking - Requests the animation to run in the main thread in a custom run-loop mode that blocks user input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode/blocking
	AnimationBlocking AnimationBlockingMode = 0
	// AnimationNonblocking - Requests the animation to run in a standard or specified run-loop mode that allows user input.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode/nonblocking
	AnimationNonblocking AnimationBlockingMode = 1
	// AnimationNonblockingThreaded - Requests the animation to run in a separate thread that is spawned by the   object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode/nonblockingThreaded
	AnimationNonblockingThreaded AnimationBlockingMode = 2
)

/* debug [enums.gen.go]: Processing enum NSAnimationCurve (4 cases) */
// AnimationCurve - These constants describe the curve of an animation—that is, the relative speed of an animation from start to finish.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/Curve
type AnimationCurve uint

const (
	// AnimationEaseIn - Describes an animation that slows down as it reaches the end.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/Curve/easeIn
	AnimationEaseIn AnimationCurve = 1
	// AnimationEaseInOut - Describes an S-curve in which the animation slowly speeds up and then slows down near the end of the animation. This constant is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/Curve/easeInOut
	AnimationEaseInOut AnimationCurve = 0
	// AnimationEaseOut - Describes an animation that slowly speeds up from the start.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/Curve/easeOut
	AnimationEaseOut AnimationCurve = 2
	// AnimationLinear - Describes an animation in which there is no change in frame rate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/Curve/linear
	AnimationLinear AnimationCurve = 3
)

/* debug [enums.gen.go]: Processing enum NSAnimationEffect (2 cases) */
// AnimationEffect - The type for standard system animation effects, which include both display and sound.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationEffect
type AnimationEffect uint

const (
	// AnimationEffectDisappearingItemDefault - The default effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationEffect/disappearingItemDefault
	AnimationEffectDisappearingItemDefault AnimationEffect = 0
	// AnimationEffectPoof - An effect showing a puff of smoke.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationEffect/poof
	AnimationEffectPoof AnimationEffect = 10
)

/* debug [enums.gen.go]: Processing enum NSApplicationActivationOptions (2 cases) */
// ApplicationActivationOptions - The following flags are for 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationOptions
type ApplicationActivationOptions uint

const (
	// ApplicationActivateAllWindows - By default, activation brings only the main and key windows forward.  If you specify NSApplicationActivateAllWindows, all of the application’s windows are brought forward.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationOptions/activateAllWindows
	ApplicationActivateAllWindows ApplicationActivationOptions = 1
	// ApplicationActivateIgnoringOtherApps - The application is activated regardless of the currently active app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationOptions/activateIgnoringOtherApps
	ApplicationActivateIgnoringOtherApps ApplicationActivationOptions = 2
)

/* debug [enums.gen.go]: Processing enum NSApplicationActivationPolicy (3 cases) */
// ApplicationActivationPolicy - Activation policies (used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum
type ApplicationActivationPolicy uint

const (
	// ApplicationActivationPolicyAccessory - The application doesn’t appear in the Dock and doesn’t have a menu bar, but it may be activated programmatically or by clicking on one of its windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum/accessory
	ApplicationActivationPolicyAccessory ApplicationActivationPolicy = 1
	// ApplicationActivationPolicyProhibited - The application doesn’t appear in the Dock and may not create windows or be activated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum/prohibited
	ApplicationActivationPolicyProhibited ApplicationActivationPolicy = 2
	// ApplicationActivationPolicyRegular - The application is an ordinary app that appears in the Dock and may have a user interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationPolicy-swift.enum/regular
	ApplicationActivationPolicyRegular ApplicationActivationPolicy = 0
)

/* debug [enums.gen.go]: Processing enum NSApplicationDelegateReply (3 cases) */
// ApplicationDelegateReply - Constants that indicate whether a copy or print operation was successful, was canceled, or failed.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/DelegateReply
type ApplicationDelegateReply uint

const (
	// ApplicationDelegateReplyCancel - Indicates the user cancelled the operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/DelegateReply/cancel
	ApplicationDelegateReplyCancel ApplicationDelegateReply = 1
	// ApplicationDelegateReplyFailure - Indicates an error occurred processing the operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/DelegateReply/failure
	ApplicationDelegateReplyFailure ApplicationDelegateReply = 2
	// ApplicationDelegateReplySuccess - Indicates the operation succeeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/DelegateReply/success
	ApplicationDelegateReplySuccess ApplicationDelegateReply = 0
)

/* debug [enums.gen.go]: Processing enum NSApplicationOcclusionState (1 cases) */
// ApplicationOcclusionState - This constant indicates whether at least part of any window owned by this app is visible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/OcclusionState-swift.struct
type ApplicationOcclusionState uint

const (
	// ApplicationOcclusionStateVisible - If set, at least part of any window owned by this app is visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/OcclusionState-swift.struct/visible
	ApplicationOcclusionStateVisible ApplicationOcclusionState = 2
)

/* debug [enums.gen.go]: Processing enum NSApplicationPresentationOptions (14 cases) */
// ApplicationPresentationOptions - Constants that control the presentation of the app, typically for fullscreen apps such as games or kiosks.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct
type ApplicationPresentationOptions uint

const (
	// ApplicationPresentationAutoHideDock - The dock is normally hidden, but automatically appears when moused near.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/autoHideDock
	ApplicationPresentationAutoHideDock ApplicationPresentationOptions = 1
	// ApplicationPresentationAutoHideMenuBar - The menu bar is normally hidden, but automatically appears when moused near.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/autoHideMenuBar
	ApplicationPresentationAutoHideMenuBar ApplicationPresentationOptions = 4
	// ApplicationPresentationAutoHideToolbar - When in fullscreen mode the window toolbar is detached from window and hides and shows with autoHidden menuBar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/autoHideToolbar
	ApplicationPresentationAutoHideToolbar ApplicationPresentationOptions = 514
	// ApplicationPresentationDisableAppleMenu - All Apple Menu items are disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/disableAppleMenu
	ApplicationPresentationDisableAppleMenu ApplicationPresentationOptions = 16
	// ApplicationPresentationDisableCursorLocationAssistance - The behavior that allows the user to shake the mouse to locate the cursor is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/disableCursorLocationAssistance
	ApplicationPresentationDisableCursorLocationAssistance ApplicationPresentationOptions = 515
	// ApplicationPresentationDisableForceQuit - The force quit panel (displayed by pressing Command + Option + Esc) is disabled
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/disableForceQuit
	ApplicationPresentationDisableForceQuit ApplicationPresentationOptions = 64
	// ApplicationPresentationDisableHideApplication - The app’s “Hide” menu item is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/disableHideApplication
	ApplicationPresentationDisableHideApplication ApplicationPresentationOptions = 256
	// ApplicationPresentationDisableMenuBarTransparency - The menu bar transparency appearance is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/disableMenuBarTransparency
	ApplicationPresentationDisableMenuBarTransparency ApplicationPresentationOptions = 512
	// ApplicationPresentationDisableProcessSwitching - The process switching user interface (Command + Tab to cycle through apps) is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/disableProcessSwitching
	ApplicationPresentationDisableProcessSwitching ApplicationPresentationOptions = 32
	// ApplicationPresentationDisableSessionTermination - The panel that shows the Restart, Shut Down, and Log Out options that are displayed as a result of pushing the power key is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/disableSessionTermination
	ApplicationPresentationDisableSessionTermination ApplicationPresentationOptions = 128
	// ApplicationPresentationFullScreen - The app is in fullscreen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/fullScreen
	ApplicationPresentationFullScreen ApplicationPresentationOptions = 513
	// ApplicationPresentationHideDock - The dock is entirely hidden and disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/hideDock
	ApplicationPresentationHideDock ApplicationPresentationOptions = 2
	// ApplicationPresentationHideMenuBar - The menu bar is entirely hidden and disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct/hideMenuBar
	ApplicationPresentationHideMenuBar ApplicationPresentationOptions = 8
	// ApplicationPresentationDefault - This is the default presentation mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplicationPresentationOptions/NSApplicationPresentationDefault
	ApplicationPresentationDefault ApplicationPresentationOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSApplicationPrintReply (4 cases) */
// ApplicationPrintReply - Constants that indicate the outcome of a print request.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PrintReply
type ApplicationPrintReply uint

const (
	// PrintingCancelled - Printing was cancelled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PrintReply/printingCancelled
	PrintingCancelled ApplicationPrintReply = 0
	// PrintingFailure - Printing failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PrintReply/printingFailure
	PrintingFailure ApplicationPrintReply = 3
	// PrintingReplyLater - The result of printing cannot be returned immediately, for example, if printing will cause the presentation of a sheet. If your method returns   it must always invoke   when the entire print operation has been completed, successfully or not.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PrintReply/printingReplyLater
	PrintingReplyLater ApplicationPrintReply = 2
	// PrintingSuccess - Printing was successful.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/PrintReply/printingSuccess
	PrintingSuccess ApplicationPrintReply = 1
)

/* debug [enums.gen.go]: Processing enum NSRemoteNotificationType (4 cases) */
// RemoteNotificationType - These constants determine whether apps launched by remote notifications display a badge.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType
type RemoteNotificationType uint

const (
	// RemoteNotificationTypeAlert - The app should display an alert.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType/alert
	RemoteNotificationTypeAlert RemoteNotificationType = 3
	// RemoteNotificationTypeBadge - The app should display a badge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType/badge
	RemoteNotificationTypeBadge RemoteNotificationType = 1
	// RemoteNotificationTypeSound - The app should play a sound.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType/sound
	RemoteNotificationTypeSound RemoteNotificationType = 2
	// RemoteNotificationTypeNone - The app shouldn’t display a badge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRemoteNotificationType/NSRemoteNotificationTypeNone
	RemoteNotificationTypeNone RemoteNotificationType = 0
)

/* debug [enums.gen.go]: Processing enum NSRequestUserAttentionType (2 cases) */
// RequestUserAttentionType - These constants specify the level of severity of a user attention request and are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RequestUserAttentionType
type RequestUserAttentionType uint

const (
	// CriticalRequest - The user attention request is a critical request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RequestUserAttentionType/criticalRequest
	CriticalRequest RequestUserAttentionType = 0
	// InformationalRequest - The user attention request is an informational request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/RequestUserAttentionType/informationalRequest
	InformationalRequest RequestUserAttentionType = 10
)

/* debug [enums.gen.go]: Processing enum NSApplicationTerminateReply (3 cases) */
// ApplicationTerminateReply - Constants that determine whether an app should terminate.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply
type ApplicationTerminateReply uint

const (
	// TerminateCancel - The app should not be terminated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply/terminateCancel
	TerminateCancel ApplicationTerminateReply = 0
	// TerminateLater - It may be OK to proceed with termination later. Returning this value causes Cocoa to run the run loop in the   until your app subsequently calls   with the value   or  . This return value is for delegates that need to provide document modal alerts (sheets) in order to decide whether to quit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply/terminateLater
	TerminateLater ApplicationTerminateReply = 2
	// TerminateNow - It is OK to proceed with termination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply/terminateNow
	TerminateNow ApplicationTerminateReply = 1
)

/* debug [enums.gen.go]: Processing enum NSWindowListOptions (1 cases) */
// WindowListOptions - This constant indicates a window ordering.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/WindowListOptions
type WindowListOptions uint

const (
	// WindowListOrderedFrontToBack - The app’s onscreen windows in front-to-back order. By default,   is used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/WindowListOptions/orderedFrontToBack
	WindowListOrderedFrontToBack WindowListOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSBitmapImageFileType (6 cases) */
// BitmapImageFileType - Constants that specify bitmap file types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/FileType
type BitmapImageFileType uint

const (
	// BitmapImageFileTypeBMP - Windows bitmap image (BMP) format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/FileType/bmp
	BitmapImageFileTypeBMP BitmapImageFileType = 1
	// BitmapImageFileTypeGIF - Graphics Image Format (GIF), originally created by CompuServe for online downloads.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/FileType/gif
	BitmapImageFileTypeGIF BitmapImageFileType = 2
	// BitmapImageFileTypeJPEG - JPEG format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/FileType/jpeg
	BitmapImageFileTypeJPEG BitmapImageFileType = 3
	// BitmapImageFileTypeJPEG2000 - JPEG 2000 file format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/FileType/jpeg2000
	BitmapImageFileTypeJPEG2000 BitmapImageFileType = 5
	// BitmapImageFileTypePNG - Portable Network Graphics (PNG) format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/FileType/png
	BitmapImageFileTypePNG BitmapImageFileType = 4
	// BitmapImageFileTypeTIFF - Tagged Image File Format (TIFF).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/FileType/tiff
	BitmapImageFileTypeTIFF BitmapImageFileType = 0
)

/* debug [enums.gen.go]: Processing enum NSBitmapFormat (7 cases) */
// BitmapFormat - Constants that represent bitmap component formats.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format
type BitmapFormat uint

const (
	// BitmapFormatAlphaFirst - A format where the alpha value comes first.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format/alphaFirst
	BitmapFormatAlphaFirst BitmapFormat = 1
	// BitmapFormatAlphaNonpremultiplied - A format where alpha values are not premultiplied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format/alphaNonpremultiplied
	BitmapFormatAlphaNonpremultiplied BitmapFormat = 2
	// BitmapFormatFloatingPointSamples - A format where samples are specified using floating-point numbers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format/floatingPointSamples
	BitmapFormatFloatingPointSamples BitmapFormat = 4
	// BitmapFormatSixteenBitBigEndian - A 16-bit, big endian format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format/sixteenBitBigEndian
	BitmapFormatSixteenBitBigEndian BitmapFormat = 7
	// BitmapFormatSixteenBitLittleEndian - A 16-bit, little endian format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format/sixteenBitLittleEndian
	BitmapFormatSixteenBitLittleEndian BitmapFormat = 5
	// BitmapFormatThirtyTwoBitBigEndian - A 32-bit, big endian format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format/thirtyTwoBitBigEndian
	BitmapFormatThirtyTwoBitBigEndian BitmapFormat = 8
	// BitmapFormatThirtyTwoBitLittleEndian - A 32-bit, little endian format.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/Format/thirtyTwoBitLittleEndian
	BitmapFormatThirtyTwoBitLittleEndian BitmapFormat = 6
)

/* debug [enums.gen.go]: Processing enum NSImageRepLoadStatus (6 cases) */
// ImageRepLoadStatus - Constants that identify the loading status of the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus
type ImageRepLoadStatus int

const (
	// ImageRepLoadStatusCompleted - Enough data has been provided to successfully decompress the image (regardless of the complete: flag).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/completed
	ImageRepLoadStatusCompleted ImageRepLoadStatus = -6
	// ImageRepLoadStatusInvalidData - An error occurred during image decompression. The image contains the portions of the data that have already been successfully decompressed, if any
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/invalidData
	ImageRepLoadStatusInvalidData ImageRepLoadStatus = -4
	// ImageRepLoadStatusReadingHeader - The image format is known, but not enough data has been read to determine the size, depth, etc., of the image. You should continue to provide more data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/readingHeader
	ImageRepLoadStatusReadingHeader ImageRepLoadStatus = -2
	// ImageRepLoadStatusUnexpectedEOF - Not enough data was available to fully decompress the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/unexpectedEOF
	ImageRepLoadStatusUnexpectedEOF ImageRepLoadStatus = -5
	// ImageRepLoadStatusUnknownType - Not enough data to determine image format. You should continue to provide more data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/unknownType
	ImageRepLoadStatusUnknownType ImageRepLoadStatus = -1
	// ImageRepLoadStatusWillNeedAllData - Incremental loading cannot be supported.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/LoadStatus/willNeedAllData
	ImageRepLoadStatusWillNeedAllData ImageRepLoadStatus = -3
)

/* debug [enums.gen.go]: Processing enum NSTIFFCompression (8 cases) */
// TIFFCompression - Constants that represent the supported TIFF data-compression schemes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression
type TIFFCompression uint

const (
	// TIFFCompressionCCITTFAX3 - CCITT Fax Group 3 compression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/ccittfax3
	TIFFCompressionCCITTFAX3 TIFFCompression = 3
	// TIFFCompressionCCITTFAX4 - CCITT Fax Group 4 compression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/ccittfax4
	TIFFCompressionCCITTFAX4 TIFFCompression = 4
	// TIFFCompressionJPEG - JPEG compression. No longer supported for input or output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/jpeg
	TIFFCompressionJPEG TIFFCompression = 6
	// TIFFCompressionLZW - LZW compression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/lzw
	TIFFCompressionLZW TIFFCompression = 5
	// TIFFCompressionNEXT - NeXT compressed. Supported for input only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/next
	TIFFCompressionNEXT TIFFCompression = 32766
	// TIFFCompressionNone - No compression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/none
	TIFFCompressionNone TIFFCompression = 1
	// TIFFCompressionOldJPEG - Old JPEG compression. No longer supported for input or output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/oldJPEG
	TIFFCompressionOldJPEG TIFFCompression = 32865
	// TIFFCompressionPackBits - PackBits compression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBitmapImageRep/TIFFCompression/packBits
	TIFFCompressionPackBits TIFFCompression = 32773
)

/* debug [enums.gen.go]: Processing enum NSBorderType (4 cases) */
// BorderType - These constants specify the type of a view’s border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType
type BorderType uint

const (
	// BezelBorder - A concave border that makes the view look sunken.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/bezelBorder
	BezelBorder BorderType = 2
	// GrooveBorder - A thin border that looks etched around the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/grooveBorder
	GrooveBorder BorderType = 3
	// LineBorder - A black line border around the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/lineBorder
	LineBorder BorderType = 1
	// NoBorder - No border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBorderType/noBorder
	NoBorder BorderType = 0
)

/* debug [enums.gen.go]: Processing enum NSBoxType (3 cases) */
// BoxType - These constants and data type identifies box types, which, in conjunction with a box’s border type, define the appearance of the box.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum
type BoxType uint

const (
	// BoxCustom - Specifies that the appearance of the box is determined entirely by the by box-configuration methods, without automatically applying Apple human interface guidelines. See Customizing for details.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum/custom
	BoxCustom BoxType = 3
	// BoxPrimary - Specifies the primary box appearance. This is the default box type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum/primary
	BoxPrimary BoxType = 0
	// BoxSeparator - Specifies that the box is a separator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/BoxType-swift.enum/separator
	BoxSeparator BoxType = 2
)

/* debug [enums.gen.go]: Processing enum NSTitlePosition (7 cases) */
// TitlePosition - Specify the location of a box’s title with respect to its border.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum
type TitlePosition uint

const (
	// AboveBottom - Title positioned above the box’s bottom border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum/aboveBottom
	AboveBottom TitlePosition = 4
	// AboveTop - Title positioned above the box’s top border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum/aboveTop
	AboveTop TitlePosition = 1
	// AtBottom - Title positioned within the box’s bottom border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum/atBottom
	AtBottom TitlePosition = 5
	// AtTop - Title positioned within the box’s top border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum/atTop
	AtTop TitlePosition = 2
	// BelowBottom - Title positioned below the box’s bottom border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum/belowBottom
	BelowBottom TitlePosition = 6
	// BelowTop - Title positioned below the box’s top border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum/belowTop
	BelowTop TitlePosition = 3
	// NoTitle - The box has no title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBox/TitlePosition-swift.enum/noTitle
	NoTitle TitlePosition = 0
)

/* debug [enums.gen.go]: Processing enum NSBrowserColumnResizingType (3 cases) */
// BrowserColumnResizingType - Types of browser column resizing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum
type BrowserColumnResizingType uint

const (
	// BrowserAutoColumnResizing - All columns have the same width, calculated using a combination of the minimum column width and maximum number of visible columns settings. The column width changes as the window size changes. The user cannot resize columns.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/autoColumnResizing
	BrowserAutoColumnResizing BrowserColumnResizingType = 1
	// BrowserNoColumnResizing - Neither   nor the user can change the column width. The developer must explicitly set all column widths.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/noColumnResizing
	BrowserNoColumnResizing BrowserColumnResizingType = 0
	// BrowserUserColumnResizing - The developer chooses the initial column widths, but users can resize all columns simultaneously or each column individually.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/ColumnResizingType-swift.enum/userColumnResizing
	BrowserUserColumnResizing BrowserColumnResizingType = 2
)

/* debug [enums.gen.go]: Processing enum NSBrowserDropOperation (2 cases) */
// BrowserDropOperation - The type used to specify the drop type of a drag-and-drop operation. See 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/DropOperation
type BrowserDropOperation uint

const (
	// BrowserDropAbove - The drop occurs above the row to which the item was dragged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/DropOperation/above
	BrowserDropAbove BrowserDropOperation = 1
	// BrowserDropOn - The drop occurs at the row to which the item was dragged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSBrowser/DropOperation/on
	BrowserDropOn BrowserDropOperation = 0
)

/* debug [enums.gen.go]: Processing enum NSBezelStyle (22 cases) */
// BezelStyle - The set of bezel styles to style buttons in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum
type BezelStyle uint

const (
	// BezelStyleAccessoryBar - A button style that’s typically used in the context of an accessory toolbar for buttons that narrow the focus of a search or other operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/accessoryBar
	BezelStyleAccessoryBar BezelStyle = 13
	// BezelStyleAccessoryBarAction - A button style that you use for extra actions in an accessory toolbar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/accessoryBarAction
	BezelStyleAccessoryBarAction BezelStyle = 12
	// BezelStyleAutomatic - The default button style based on the button’s contents and position within the window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/automatic
	BezelStyleAutomatic BezelStyle = 0
	// BezelStyleBadge - A button style suitable for displaying additional information.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/badge
	BezelStyleBadge BezelStyle = 15
	// BezelStyleCircular - A round button that can contain either a single character or an icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/circular
	BezelStyleCircular BezelStyle = 7
	// BezelStyleDisclosure - A bezel style button for use with a disclosure triangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/disclosure
	BezelStyleDisclosure BezelStyle = 5
	// BezelStyleFlexiblePush - A push button with a flexible height to accommodate longer text labels or an image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/flexiblePush
	BezelStyleFlexiblePush BezelStyle = 2
	// BezelStyleGlass - A bezel style with a glass effect
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/glass
	BezelStyleGlass BezelStyle = 16
	// BezelStyleHelpButton - A round button with a question mark, providing the standard help button look.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/helpButton
	BezelStyleHelpButton BezelStyle = 9
	// BezelStyleInline - A button that has a solid round-rectangle border background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/inline
	BezelStyleInline BezelStyle = 25
	// BezelStylePush - A standard push style button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/push
	BezelStylePush BezelStyle = 1
	// BezelStylePushDisclosure - A bezel style push button with a disclosure triangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/pushDisclosure
	BezelStylePushDisclosure BezelStyle = 14
	// BezelStyleRecessed - A bezel style appropriate for use in scope bars and title bar accessories, similar to the bookmarks bar in Safari.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/recessed
	BezelStyleRecessed BezelStyle = 23
	// BezelStyleRegularSquare - A rectangular button with a two-point border, designed for icons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/regularSquare
	BezelStyleRegularSquare BezelStyle = 20
	// BezelStyleRoundRect - A bezel style appropriate for use as an action or auxiliary button in scope bars and title bar accessories.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/roundRect
	BezelStyleRoundRect BezelStyle = 22
	// BezelStyleRounded - A rounded rectangle button, designed for text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/rounded
	BezelStyleRounded BezelStyle = 19
	// BezelStyleRoundedDisclosure - A bezel style for use with a vertically expanding and collapsing disclosure button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/roundedDisclosure
	BezelStyleRoundedDisclosure BezelStyle = 24
	// BezelStyleShadowlessSquare - A rectangular button with no shadow, so it can abut the cells without overlapping shadows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/shadowlessSquare
	BezelStyleShadowlessSquare BezelStyle = 17
	// BezelStyleSmallSquare - A simple square bezel style that can scale to any size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/smallSquare
	BezelStyleSmallSquare BezelStyle = 10
	// BezelStyleTexturedRounded - A bezel style appropriate for use in the toolbar or title bar regions of a window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/texturedRounded
	BezelStyleTexturedRounded BezelStyle = 21
	// BezelStyleTexturedSquare - A bezel style appropriate for use with textured (metal) windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/texturedSquare
	BezelStyleTexturedSquare BezelStyle = 18
	// BezelStyleToolbar - A button style that’s appropriate for a toolbar item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/BezelStyle-swift.enum/toolbar
	BezelStyleToolbar BezelStyle = 11
)

/* debug [enums.gen.go]: Processing enum NSButtonType (10 cases) */
// ButtonType - Button types that you can specify using 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType
type ButtonType uint

const (
	// ButtonTypeAccelerator - A button that sends repeating actions as pressure changes occur.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/accelerator
	ButtonTypeAccelerator ButtonType = 8
	// ButtonTypeMomentaryChange - A button that displays its alternate content when clicked and returns to its normal content when the user releases it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/momentaryChange
	ButtonTypeMomentaryChange ButtonType = 5
	// ButtonTypeMomentaryLight - A button that displays a highlight when the user clicks it and returns to its normal state when the user releases it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/momentaryLight
	ButtonTypeMomentaryLight ButtonType = 0
	// ButtonTypeMomentaryPushIn - A button that illuminates when the user clicks it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/momentaryPushIn
	ButtonTypeMomentaryPushIn ButtonType = 7
	// ButtonTypeMultiLevelAccelerator - A button that allows for a configurable number of stepped pressure levels and provides tactile feedback as the user reaches each step.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/multiLevelAccelerator
	ButtonTypeMultiLevelAccelerator ButtonType = 9
	// ButtonTypeOnOff - A button that switches between a normal and emphasized bezel on each click.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/onOff
	ButtonTypeOnOff ButtonType = 6
	// ButtonTypePushOnPushOff - A button that switches between on and off states with each click.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/pushOnPushOff
	ButtonTypePushOnPushOff ButtonType = 1
	// ButtonTypeRadio - A button that displays a single selected value from group of possible choices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/radio
	ButtonTypeRadio ButtonType = 4
	// ButtonTypeSwitch - A standard checkbox button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/switch
	ButtonTypeSwitch ButtonType = 3
	// ButtonTypeToggle - A button that switches between its normal and alternate content on each click.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/ButtonType/toggle
	ButtonTypeToggle ButtonType = 2
)

/* debug [enums.gen.go]: Processing enum NSGradientType (5 cases) */
// GradientType - Specify the gradients used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/GradientType
type GradientType uint

const (
	// GradientConcaveStrong - As with  , the top-left corner is light gray, and the bottom-right corner is dark gray, but the difference between the grays is greater, so the appearance of being pushed in is stronger.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/GradientType/concaveStrong
	GradientConcaveStrong GradientType = 2
	// GradientConcaveWeak - The top-left corner is light gray, and the bottom-right corner is dark gray, so the button appears to be pushed in.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/GradientType/concaveWeak
	GradientConcaveWeak GradientType = 1
	// GradientConvexStrong - As with  , the top-left corner is dark gray, and the bottom-right corner is light gray, but the difference between the grays is greater, so the appearance of sticking out is stronger.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/GradientType/convexStrong
	GradientConvexStrong GradientType = 4
	// GradientConvexWeak - The top-left corner is dark gray, and the bottom-right corner is light gray, so the button appears to be sticking out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/GradientType/convexWeak
	GradientConvexWeak GradientType = 3
	// GradientNone - There is no gradient, so the button looks flat.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSButton/GradientType/none
	GradientNone GradientType = 0
)

/* debug [enums.gen.go]: Processing enum NSCellAttribute (17 cases) */
// CellAttribute - Constants for specifying how a button behaves when pressed and how it displays its state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute
type CellAttribute uint

const (
	// CellAllowsMixedState - Lets the cell’s state be  , as well as   and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellAllowsMixedState
	CellAllowsMixedState CellAttribute = 16
	// CellChangesContents - If the cell’s state is   or  , displays the cell’s alternate image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellChangesContents
	CellChangesContents CellAttribute = 14
	// CellDisabled - Does not let the user manipulate the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellDisabled
	CellDisabled CellAttribute = 0
	// CellEditable - Lets the user edit the cell’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellEditable
	CellEditable CellAttribute = 3
	// CellHasImageHorizontal - Controls the position of the cell’s image: places the image on the right of any text in the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellHasImageHorizontal
	CellHasImageHorizontal CellAttribute = 12
	// CellHasImageOnLeftOrBottom - Controls the position of the cell’s image: places the image on the left of or below any text in the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellHasImageOnLeftOrBottom
	CellHasImageOnLeftOrBottom CellAttribute = 13
	// CellHasOverlappingImage - Controls the position of the cell’s image: places the image over any text in the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellHasOverlappingImage
	CellHasOverlappingImage CellAttribute = 11
	// CellHighlighted - Draws the cell with a highlighted appearance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellHighlighted
	CellHighlighted CellAttribute = 5
	// CellIsBordered - Draws a border around the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellIsBordered
	CellIsBordered CellAttribute = 10
	// CellIsInsetButton - Insets the cell’s contents from the border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellIsInsetButton
	CellIsInsetButton CellAttribute = 15
	// CellLightsByBackground - If the cell is pushed in, changes the cell’s background color from gray to white.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellLightsByBackground
	CellLightsByBackground CellAttribute = 9
	// CellLightsByContents - If the cell is pushed in, displays the cell’s alternate image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellLightsByContents
	CellLightsByContents CellAttribute = 6
	// CellLightsByGray - If the cell is pushed in, displays the cell’s image as darkened.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellLightsByGray
	CellLightsByGray CellAttribute = 7
	// CellState - The cell’s state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/cellState
	CellState CellAttribute = 1
	// ChangeBackgroundCell - If the cell’s state is   or  , changes the cell’s background color from gray to white.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/changeBackgroundCell
	ChangeBackgroundCell CellAttribute = 8
	// ChangeGrayCell - If the cell’s state is   or  , displays the cell’s image as darkened.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/changeGrayCell
	ChangeGrayCell CellAttribute = 4
	// PushInCell - Determines whether the cell’s image and text appear to be shifted down and to the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/Attribute/pushInCell
	PushInCell CellAttribute = 2
)

/* debug [enums.gen.go]: Processing enum NSCellType (3 cases) */
// CellType - Constants for specifying how a cell represents its data (as text or as an image).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/CellType
type CellType uint

const (
	// ImageCellType - Cell displays images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/CellType/imageCellType
	ImageCellType CellType = 2
	// NullCellType - Cell displays nothing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/CellType/nullCellType
	NullCellType CellType = 0
	// TextCellType - Cell displays text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/CellType/textCellType
	TextCellType CellType = 1
)

/* debug [enums.gen.go]: Processing enum NSCellHitResult (4 cases) */
// CellHitResult - Constants used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/HitResult
type CellHitResult uint

const (
	// CellHitContentArea - A content area in the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/HitResult/contentArea
	CellHitContentArea CellHitResult = 1
	// CellHitEditableTextArea - An editable text area of the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/HitResult/editableTextArea
	CellHitEditableTextArea CellHitResult = 2
	// CellHitTrackableArea - A trackable area in the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/HitResult/trackableArea
	CellHitTrackableArea CellHitResult = 4
	// CellHitNone - An empty area, or did not hit in the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCellHitResult/NSCellHitNone
	CellHitNone CellHitResult = 0
)

/* debug [enums.gen.go]: Processing enum NSCellStyleMask (5 cases) */
// CellStyleMask - Constants for specifying what happens when a button is pressed or is displaying its alternate state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/StyleMask
type CellStyleMask uint

const (
	// ChangeBackgroundCellMask - Same as  , but only background pixels are changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/StyleMask/changeBackgroundCellMask
	ChangeBackgroundCellMask CellStyleMask = 8
	// ChangeGrayCellMask - The button cell swaps the “control color” (the   method of  ) and white pixels on its background and icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/StyleMask/changeGrayCellMask
	ChangeGrayCellMask CellStyleMask = 4
	// ContentsCellMask - The button cell displays its alternate icon and/or title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/StyleMask/contentsCellMask
	ContentsCellMask CellStyleMask = 1
	// PushInCellMask - The button cell “pushes in” if it has a border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/StyleMask/pushInCellMask
	PushInCellMask CellStyleMask = 2
	// NoCellMask - The button cell doesn’t change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCellStyleMask/NSNoCellMask
	NoCellMask CellStyleMask = 0
)

/* debug [enums.gen.go]: Processing enum NSCharacterCollection (6 cases) */
// CharacterCollection - Values that map character identifiers to glyphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection
type CharacterCollection uint

const (
	// AdobeCNS1CharacterCollection - Indicates the Adobe-CNS1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection/adobeCNS1CharacterCollection
	AdobeCNS1CharacterCollection CharacterCollection = 1
	// AdobeGB1CharacterCollection - Indicates the Adobe-GB1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection/adobeGB1CharacterCollection
	AdobeGB1CharacterCollection CharacterCollection = 2
	// AdobeJapan1CharacterCollection - Indicates the Adobe-Japan1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection/adobeJapan1CharacterCollection
	AdobeJapan1CharacterCollection CharacterCollection = 3
	// AdobeJapan2CharacterCollection - Indicates the Adobe-Japan2 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection/adobeJapan2CharacterCollection
	AdobeJapan2CharacterCollection CharacterCollection = 4
	// AdobeKorea1CharacterCollection - Indicates the Adobe-Korea1 mapping.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection/adobeKorea1CharacterCollection
	AdobeKorea1CharacterCollection CharacterCollection = 5
	// IdentityMappingCharacterCollection - Indicates that the character identifier is equal to the glyph index.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCharacterCollection/identityMappingCharacterCollection
	IdentityMappingCharacterCollection CharacterCollection = 0
)

/* debug [enums.gen.go]: Processing enum NSCollectionElementCategory (4 cases) */
// CollectionElementCategory - Constants specifying the type of element in the collection view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionElementCategory
type CollectionElementCategory uint

const (
	// CollectionElementCategoryDecorationView - The element is a decoration view. Decoration views represent visual adornments that do not contain any data of their own.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionElementCategory/decorationView
	CollectionElementCategoryDecorationView CollectionElementCategory = 2
	// CollectionElementCategoryInterItemGap - The element is an inter-item gap. An inter-item gap element is a custom visual indicator that is displayed between items when dropping items into the collection view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionElementCategory/interItemGap
	CollectionElementCategoryInterItemGap CollectionElementCategory = 3
	// CollectionElementCategoryItem - The element is an item. Items represent the main content of your collection view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionElementCategory/item
	CollectionElementCategoryItem CollectionElementCategory = 0
	// CollectionElementCategorySupplementaryView - The element is a supplementary view. Use supplementary views for single views that contain some data but are associated with an entire section. For example, use them to specify header or footer views for a section.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionElementCategory/supplementaryView
	CollectionElementCategorySupplementaryView CollectionElementCategory = 1
)

/* debug [enums.gen.go]: Processing enum NSCollectionLayoutSectionOrthogonalScrollingBehavior (6 cases) */
// CollectionLayoutSectionOrthogonalScrollingBehavior - The scrolling behavior of the layout’s sections in relation to the main layout axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSectionOrthogonalScrollingBehavior
type CollectionLayoutSectionOrthogonalScrollingBehavior uint

const (
	// CollectionLayoutSectionOrthogonalScrollingBehaviorContinuous - The section allows users to scroll its content orthogonally with continuous scrolling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSectionOrthogonalScrollingBehavior/continuous
	CollectionLayoutSectionOrthogonalScrollingBehaviorContinuous CollectionLayoutSectionOrthogonalScrollingBehavior = 1
	// CollectionLayoutSectionOrthogonalScrollingBehaviorContinuousGroupLeadingBoundary - The section allows users to scroll its content orthogonally, coming to a natural stop at the leading boundary of the visible group.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSectionOrthogonalScrollingBehavior/continuousGroupLeadingBoundary
	CollectionLayoutSectionOrthogonalScrollingBehaviorContinuousGroupLeadingBoundary CollectionLayoutSectionOrthogonalScrollingBehavior = 2
	// CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPaging - The section allows users to page its content orthogonally one group at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSectionOrthogonalScrollingBehavior/groupPaging
	CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPaging CollectionLayoutSectionOrthogonalScrollingBehavior = 4
	// CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPagingCentered - The section allows users to page its content orthogonally one group at a time, centering each group.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSectionOrthogonalScrollingBehavior/groupPagingCentered
	CollectionLayoutSectionOrthogonalScrollingBehaviorGroupPagingCentered CollectionLayoutSectionOrthogonalScrollingBehavior = 5
	// CollectionLayoutSectionOrthogonalScrollingBehaviorNone - The section does not allow users to scroll its content orthogonally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSectionOrthogonalScrollingBehavior/none
	CollectionLayoutSectionOrthogonalScrollingBehaviorNone CollectionLayoutSectionOrthogonalScrollingBehavior = 0
	// CollectionLayoutSectionOrthogonalScrollingBehaviorPaging - The section allows users to page its content orthogonally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSectionOrthogonalScrollingBehavior/paging
	CollectionLayoutSectionOrthogonalScrollingBehaviorPaging CollectionLayoutSectionOrthogonalScrollingBehavior = 3
)

/* debug [enums.gen.go]: Processing enum NSCollectionViewDropOperation (2 cases) */
// CollectionViewDropOperation - These constants specify if acceptance of a drop should be at the item it is dropped on or before the item. These constants are used by the  
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation
type CollectionViewDropOperation uint

const (
	// CollectionViewDropBefore - The drop occurs before the collection view item to which the item was dragged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation/before
	CollectionViewDropBefore CollectionViewDropOperation = 1
	// CollectionViewDropOn - The drop occurs at the collection view item to which the item was dragged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DropOperation/on
	CollectionViewDropOn CollectionViewDropOperation = 0
)

/* debug [enums.gen.go]: Processing enum NSCollectionViewScrollDirection (2 cases) */
// CollectionViewScrollDirection - Constants indicating the scrolling direction for the layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection
type CollectionViewScrollDirection uint

const (
	// CollectionViewScrollDirectionHorizontal - The layout scrolls content horizontally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection/horizontal
	CollectionViewScrollDirectionHorizontal CollectionViewScrollDirection = 1
	// CollectionViewScrollDirectionVertical - The layout scrolls content vertically.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollDirection/vertical
	CollectionViewScrollDirectionVertical CollectionViewScrollDirection = 0
)

/* debug [enums.gen.go]: Processing enum NSCollectionUpdateAction (5 cases) */
// CollectionUpdateAction - Constants indicating the type of action being performed on an item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction
type CollectionUpdateAction uint

const (
	// CollectionUpdateActionDelete - Remove the action from the collection view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction/delete
	CollectionUpdateActionDelete CollectionUpdateAction = 1
	// CollectionUpdateActionInsert - Insert the item into the collection view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction/insert
	CollectionUpdateActionInsert CollectionUpdateAction = 0
	// CollectionUpdateActionMove - Move the item from its current location to a new location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction/move
	CollectionUpdateActionMove CollectionUpdateAction = 3
	// CollectionUpdateActionNone - Take no action on the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction/none
	CollectionUpdateActionNone CollectionUpdateAction = 4
	// CollectionUpdateActionReload - Reload the item, which consists of deleting and then inserting the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/UpdateAction/reload
	CollectionUpdateActionReload CollectionUpdateAction = 2
)

/* debug [enums.gen.go]: Processing enum NSCollectionViewScrollPosition (11 cases) */
// CollectionViewScrollPosition - Constants indicating the options for scrolling the collection view’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition
type CollectionViewScrollPosition uint

const (
	// CollectionViewScrollPositionBottom - Scroll so that the bottom edge of the bounding box is adjacent to the bottom of the collection view’s bounds. This option must not be combined with the  ,  , or   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/bottom
	CollectionViewScrollPositionBottom CollectionViewScrollPosition = 4
	// CollectionViewScrollPositionCenteredHorizontally - Scroll so that the selected items’ bounding box is centered horizontally in the collection view’s bounds. This option must not be combined with the  ,   ,  ,  , or   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/centeredHorizontally
	CollectionViewScrollPositionCenteredHorizontally CollectionViewScrollPosition = 16
	// CollectionViewScrollPositionCenteredVertically - Scroll so that the bounding box of the selected items is centered vertically in the collection view’s bounds. This option must not be combined with the  ,  , or   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/centeredVertically
	CollectionViewScrollPositionCenteredVertically CollectionViewScrollPosition = 2
	// CollectionViewScrollPositionLeadingEdge - Scroll so that the leading edge of the selected items’ bounding box is adjacent to the leading edge of the collection view’s bounds. Use this option to support both left-to-right and right-to-left layouts appropriately. This option must not be combined with the  ,  ,  ,  , or   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/leadingEdge
	CollectionViewScrollPositionLeadingEdge CollectionViewScrollPosition = 64
	// CollectionViewScrollPositionLeft - Scroll so that the left edge of the selected items’ bounding box is adjacent to the left edge of the collection view’s bounds. This option must not be combined with the  ,  ,  ,  , or   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/left
	CollectionViewScrollPositionLeft CollectionViewScrollPosition = 8
	// CollectionViewScrollPositionNearestHorizontalEdge - Scroll so that the bounding box is adjacent to the nearest edge (top or bottom) of the collection view. This option must not be combined with the  ,  , or   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/nearestHorizontalEdge
	CollectionViewScrollPositionNearestHorizontalEdge CollectionViewScrollPosition = 512
	// CollectionViewScrollPositionNearestVerticalEdge - Scroll so that the bounding box is adjacent to the nearest edge (leading or trailing) of the collection view. Use this option to support both left-to-right and right-to-left layouts appropriately. This option must not be combined with the  ,  ,  ,  , or   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/nearestVerticalEdge
	CollectionViewScrollPositionNearestVerticalEdge CollectionViewScrollPosition = 256
	// CollectionViewScrollPositionRight - Scroll so that the right edge of the selected items’ bounding box is adjacent to the right edge of the collection view’s bounds. This option must not be combined with the  ,  ,  ,  , or   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/right
	CollectionViewScrollPositionRight CollectionViewScrollPosition = 32
	// CollectionViewScrollPositionTop - Scroll so that the top edge of the selected items’ bounding box is adjacent to the top edge of the collection view’s bounds. This option must not be combined with the  ,  , and   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/top
	CollectionViewScrollPositionTop CollectionViewScrollPosition = 1
	// CollectionViewScrollPositionTrailingEdge - Scroll so that the trailing edge of the selected items’ bounding box is adjacent to the trailing edge of the collection view’s bounds. Use this option to support both left-to-right and right-to-left layouts appropriately. This option must not be combined with the  ,  ,  ,  , or   options, but may be combined with other options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/ScrollPosition/trailingEdge
	CollectionViewScrollPositionTrailingEdge CollectionViewScrollPosition = 128
	// CollectionViewScrollPositionNone - Do not scroll.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewScrollPosition/NSCollectionViewScrollPositionNone
	CollectionViewScrollPositionNone CollectionViewScrollPosition = 0
)

/* debug [enums.gen.go]: Processing enum NSCollectionViewItemHighlightState (4 cases) */
// CollectionViewItemHighlightState - Constants indicating the type of highlight applied to an item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum
type CollectionViewItemHighlightState uint

const (
	// CollectionViewItemHighlightAsDropTarget - The drop target highlight state. This type of highlight is applied when the item is the target of a drop operation on the collection view. After the drop operation completes, the highlight state returns to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum/asDropTarget
	CollectionViewItemHighlightAsDropTarget CollectionViewItemHighlightState = 3
	// CollectionViewItemHighlightForDeselection - The deselection highlight state. During interactive selection, this state is used to indicate that the item will become deselected when interactions end. After interactions end, the highlight state returns to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum/forDeselection
	CollectionViewItemHighlightForDeselection CollectionViewItemHighlightState = 2
	// CollectionViewItemHighlightForSelection - The selected highlight state. This type of highlight is applied when an item is selected. During interactive highlighting, this state is also applied to indicate that the item will become highlighted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum/forSelection
	CollectionViewItemHighlightForSelection CollectionViewItemHighlightState = 1
	// CollectionViewItemHighlightNone - No highlight state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewItem/HighlightState-swift.enum/none
	CollectionViewItemHighlightNone CollectionViewItemHighlightState = 0
)

/* debug [enums.gen.go]: Processing enum NSColorPanelMode (9 cases) */
// ColorPanelMode - A type defined for the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum
type ColorPanelMode uint

const (
	// ColorPanelModeCMYK - The cyan-magenta-yellow-black color mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum/CMYK
	ColorPanelModeCMYK ColorPanelMode = 2
	// ColorPanelModeHSB - The hue-saturation-brightness color mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum/HSB
	ColorPanelModeHSB ColorPanelMode = 3
	// ColorPanelModeRGB - The red-green-blue color mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum/RGB
	ColorPanelModeRGB ColorPanelMode = 1
	// ColorPanelModeColorList - The custom color list mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum/colorList
	ColorPanelModeColorList ColorPanelMode = 5
	// ColorPanelModeCrayon - The crayon picker mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum/crayon
	ColorPanelModeCrayon ColorPanelMode = 7
	// ColorPanelModeCustomPalette - The custom palette color mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum/customPalette
	ColorPanelModeCustomPalette ColorPanelMode = 4
	// ColorPanelModeGray - The grayscale-alpha color mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum/gray
	ColorPanelModeGray ColorPanelMode = 0
	// ColorPanelModeNone - No color panel mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum/none
	ColorPanelModeNone ColorPanelMode = 0
	// ColorPanelModeWheel - The color wheel mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorPanel/Mode-swift.enum/wheel
	ColorPanelModeWheel ColorPanelMode = 6
)

/* debug [enums.gen.go]: Processing enum NSColorWellStyle (3 cases) */
// ColorWellStyle - Constants that specify the appearance and interaction modes for a color well.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style
type ColorWellStyle uint

const (
	// ColorWellStyleDefault - The default style for color wells.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/default
	ColorWellStyleDefault ColorWellStyle = 0
	// ColorWellStyleExpanded - A style that supports a color picker popover for fast interactions, and adds a dedicated button to display the color panel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/expanded
	ColorWellStyleExpanded ColorWellStyle = 2
	// ColorWellStyleMinimal - A style that adds minimal adornments to the color well.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorWell/Style/minimal
	ColorWellStyleMinimal ColorWellStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSComboButtonStyle (2 cases) */
// ComboButtonStyle - Constants that indicate how a combo button presents its menu.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboButton/Style-swift.enum
type ComboButtonStyle uint

const (
	// ComboButtonStyleSplit - A style that separates the button’s title and image from the menu indicator people use to activate the button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboButton/Style-swift.enum/split
	ComboButtonStyleSplit ComboButtonStyle = 0
	// ComboButtonStyleUnified - A style that unifies the button’s title and image with the menu indicator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSComboButton/Style-swift.enum/unified
	ComboButtonStyleUnified ComboButtonStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSCompositingOperation (29 cases) */
// CompositingOperation - Constants that describe compositing operators in terms of source and destination images, each having an opaque and transparent region.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
type CompositingOperation uint

const (
	// CompositingOperationHighlight - The source image wherever it is opaque, and the destination image elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/NSCompositingOperationHighlight
	CompositingOperationHighlight CompositingOperation = 12
	// CompositingOperationClear - Transparency everywhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/clear
	CompositingOperationClear CompositingOperation = 0
	// CompositingOperationColor - Uses the hue and saturation of the source and the luminosity of the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/color
	CompositingOperationColor CompositingOperation = 27
	// CompositingOperationColorBurn - Darkens the destination color to reflect the source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/colorBurn
	CompositingOperationColorBurn CompositingOperation = 20
	// CompositingOperationColorDodge - Brightens the destination to reflect the source.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/colorDodge
	CompositingOperationColorDodge CompositingOperation = 19
	// CompositingOperationCopy - The source image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/copy
	CompositingOperationCopy CompositingOperation = 1
	// CompositingOperationDarken - Use the darker of the source and destination colors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/darken
	CompositingOperationDarken CompositingOperation = 17
	// CompositingOperationDestinationAtop - The destination image wherever both images are opaque, the source image wherever it is opaque and the destination image is transparent, and transparent elsehwere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/destinationAtop
	CompositingOperationDestinationAtop CompositingOperation = 9
	// CompositingOperationDestinationIn - The destination image wherever both images are opaque, and transparent elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/destinationIn
	CompositingOperationDestinationIn CompositingOperation = 7
	// CompositingOperationDestinationOut - The destination image wherever it is opaque and the source image is transparent, and transparent elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/destinationOut
	CompositingOperationDestinationOut CompositingOperation = 8
	// CompositingOperationDestinationOver - The destination image wherever it is opaque, and the source image elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/destinationOver
	CompositingOperationDestinationOver CompositingOperation = 6
	// CompositingOperationDifference - Subtracts the darker value from the lighter value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/difference
	CompositingOperationDifference CompositingOperation = 23
	// CompositingOperationExclusion - Subtracts the darker value from the lighter value, except lower in contrast.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/exclusion
	CompositingOperationExclusion CompositingOperation = 24
	// CompositingOperationHardLight - Multiplies or screens colors, with the effect of shining a spotlight on the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/hardLight
	CompositingOperationHardLight CompositingOperation = 22
	// CompositingOperationHue - Uses the hue of the source and the saturation and luminosity of the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/hue
	CompositingOperationHue CompositingOperation = 25
	// CompositingOperationLighten - Use the lighter of the source and destination colors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/lighten
	CompositingOperationLighten CompositingOperation = 18
	// CompositingOperationLuminosity - Uses the luminosity of the source and the hue and saturation of the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/luminosity
	CompositingOperationLuminosity CompositingOperation = 28
	// CompositingOperationMultiply - The source color is multiplied by the destination color.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/multiply
	CompositingOperationMultiply CompositingOperation = 14
	// CompositingOperationOverlay - Source colors overlay the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/overlay
	CompositingOperationOverlay CompositingOperation = 16
	// CompositingOperationPlusDarker - The sum of the source and destination images, with color values approach 0 as a limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/plusDarker
	CompositingOperationPlusDarker CompositingOperation = 11
	// CompositingOperationPlusLighter - The sum of the source and destination images, with color values approach 1 as a limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/plusLighter
	CompositingOperationPlusLighter CompositingOperation = 13
	// CompositingOperationSaturation - Uses the saturation value of the source and the hue and luminosity of the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/saturation
	CompositingOperationSaturation CompositingOperation = 26
	// CompositingOperationScreen - Multiplies the complement of the destination and source color values, and then complements the result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/screen
	CompositingOperationScreen CompositingOperation = 15
	// CompositingOperationSoftLight - Darkens or lightens colors, with the effect of shining a diffused spotlight on the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/softLight
	CompositingOperationSoftLight CompositingOperation = 21
	// CompositingOperationSourceAtop - The source image wherever both images are opaque, the destination image wherever it is opaque but the source image is transparent, and transparent elsewhere
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/sourceAtop
	CompositingOperationSourceAtop CompositingOperation = 5
	// CompositingOperationSourceIn - The source image wherever both images are opaque, and transparent elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/sourceIn
	CompositingOperationSourceIn CompositingOperation = 3
	// CompositingOperationSourceOut - The source image wherever it is opaque and the destination image is transparent, and transparent elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/sourceOut
	CompositingOperationSourceOut CompositingOperation = 4
	// CompositingOperationSourceOver - The source image wherever it is opaque, and the destination image elsewhere.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/sourceOver
	CompositingOperationSourceOver CompositingOperation = 2
	// CompositingOperationXOR - Exclusive OR of the source and destination images.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation/xor
	CompositingOperationXOR CompositingOperation = 10
)

/* debug [enums.gen.go]: Processing enum NSControlBorderShape (4 cases) */
// ControlBorderShape enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape
type ControlBorderShape uint

const (
	// ControlBorderShapeAutomatic - The control will resolve this to an appropriate shape for the given control size and context
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape/automatic
	ControlBorderShapeAutomatic ControlBorderShape = 0
	// ControlBorderShapeCapsule - The control will resolve this to an appropriate shape for the given control size and context
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape/capsule
	ControlBorderShapeCapsule ControlBorderShape = 0
	// ControlBorderShapeCircle - The control will resolve this to an appropriate shape for the given control size and context
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape/circle
	ControlBorderShapeCircle ControlBorderShape = 0
	// ControlBorderShapeRoundedRectangle - The control will resolve this to an appropriate shape for the given control size and context
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/BorderShape/roundedRectangle
	ControlBorderShapeRoundedRectangle ControlBorderShape = 0
)

/* debug [enums.gen.go]: Processing enum NSControlSize (5 cases) */
// ControlSize - A constant for specifying a cell’s size.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum
type ControlSize uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/extraLarge
	ControlSizeExtraLarge ControlSize = 4
	// ControlSizeLarge - A size larger than the default control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/large
	ControlSizeLarge ControlSize = 3
	// ControlSizeMini - The smallest control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/mini
	ControlSizeMini ControlSize = 2
	// ControlSizeRegular - The default control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/regular
	ControlSizeRegular ControlSize = 0
	// ControlSizeSmall - A size smaller than the default control size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ControlSize-swift.enum/small
	ControlSizeSmall ControlSize = 1
)

/* debug [enums.gen.go]: Processing enum NSCellImagePosition (9 cases) */
// CellImagePosition - A constant for specifying the position of a button’s image relative to its title.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition
type CellImagePosition uint

const (
	// ImageAbove - The image is above the title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition/imageAbove
	ImageAbove CellImagePosition = 5
	// ImageBelow - The image is below the title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition/imageBelow
	ImageBelow CellImagePosition = 4
	// ImageLeading - The image is on the title’s leading edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition/imageLeading
	ImageLeading CellImagePosition = 7
	// ImageLeft - The image is to the left of the title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition/imageLeft
	ImageLeft CellImagePosition = 2
	// ImageOnly - The cell displays an image but not a title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition/imageOnly
	ImageOnly CellImagePosition = 1
	// ImageOverlaps - The image overlaps the title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition/imageOverlaps
	ImageOverlaps CellImagePosition = 6
	// ImageRight - The image is to the right of the title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition/imageRight
	ImageRight CellImagePosition = 3
	// ImageTrailing - The image is on the title’s trailing edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition/imageTrailing
	ImageTrailing CellImagePosition = 8
	// NoImage - The cell doesn’t display an image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControl/ImagePosition/noImage
	NoImage CellImagePosition = 0
)

/* debug [enums.gen.go]: Processing enum NSControlTint (4 cases) */
// ControlTint - Constants for specifying a cell’s tint color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint
type ControlTint uint

const (
	// BlueControlTint - Aqua control tint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint/blueControlTint
	BlueControlTint ControlTint = 1
	// ClearControlTint - Clear control tint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint/clearControlTint
	ClearControlTint ControlTint = 7
	// DefaultControlTint - The current default tint setting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint/defaultControlTint
	DefaultControlTint ControlTint = 0
	// GraphiteControlTint - Graphite control tint.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSControlTint/graphiteControlTint
	GraphiteControlTint ControlTint = 6
)

/* debug [enums.gen.go]: Processing enum NSCursorFrameResizePosition (8 cases) */
// CursorFrameResizePosition - The position along the perimeter of a rectangular frame (its edges and corners) from which it’s resized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition
type CursorFrameResizePosition uint

const (
	// CursorFrameResizePositionBottom - The bottom edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/bottom
	CursorFrameResizePositionBottom CursorFrameResizePosition = 4
	// CursorFrameResizePositionBottomLeft - The bottom left corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/bottomLeft
	CursorFrameResizePositionBottomLeft CursorFrameResizePosition = 0
	// CursorFrameResizePositionBottomRight - The bottom right corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/bottomRight
	CursorFrameResizePositionBottomRight CursorFrameResizePosition = 0
	// CursorFrameResizePositionLeft - The left edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/left
	CursorFrameResizePositionLeft CursorFrameResizePosition = 2
	// CursorFrameResizePositionRight - The right edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/right
	CursorFrameResizePositionRight CursorFrameResizePosition = 8
	// CursorFrameResizePositionTop - The top edge of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/top
	CursorFrameResizePositionTop CursorFrameResizePosition = 1
	// CursorFrameResizePositionTopLeft - The top left corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/topLeft
	CursorFrameResizePositionTopLeft CursorFrameResizePosition = 0
	// CursorFrameResizePositionTopRight - The top right corner of the frame.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursor/FrameResizePosition/topRight
	CursorFrameResizePositionTopRight CursorFrameResizePosition = 0
)

/* debug [enums.gen.go]: Processing enum NSCursorFrameResizeDirections (3 cases) */
// CursorFrameResizeDirections - The directions in which a rectangular frame can be resized.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections
type CursorFrameResizeDirections uint

const (
	// CursorFrameResizeDirectionsAll - Indicates that the shape can be resized inwards or wards to be either smaller or larger, respectively.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections/NSCursorFrameResizeDirectionsAll
	CursorFrameResizeDirectionsAll CursorFrameResizeDirections = 0
	// CursorFrameResizeDirectionsInward - Indicates that the shape can be resized inwards to be smaller.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections/NSCursorFrameResizeDirectionsInward
	CursorFrameResizeDirectionsInward CursorFrameResizeDirections = 1
	// CursorFrameResizeDirectionsOutward - Indicates that the shape can be resized outwards to be larger.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCursorFrameResizeDirections/NSCursorFrameResizeDirectionsOutward
	CursorFrameResizeDirectionsOutward CursorFrameResizeDirections = 2
)

/* debug [enums.gen.go]: Processing enum NSDatePickerStyle (3 cases) */
// DatePickerStyle - Constants that define the visual appearance of the date picker cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Style
type DatePickerStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Style/clockAndCalendar
	DatePickerStyleClockAndCalendar DatePickerStyle = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Style/textField
	DatePickerStyleTextField DatePickerStyle = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Style/textFieldAndStepper
	DatePickerStyleTextFieldAndStepper DatePickerStyle = 0
)

/* debug [enums.gen.go]: Processing enum NSDatePickerElementFlags (6 cases) */
// DatePickerElementFlags - Constants that specify the date and time elements displayed by the picker.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/ElementFlags
type DatePickerElementFlags uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/ElementFlags/era
	DatePickerElementFlagEra DatePickerElementFlags = 256
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/ElementFlags/hourMinute
	DatePickerElementFlagHourMinute DatePickerElementFlags = 12
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/ElementFlags/hourMinuteSecond
	DatePickerElementFlagHourMinuteSecond DatePickerElementFlags = 14
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/ElementFlags/timeZone
	DatePickerElementFlagTimeZone DatePickerElementFlags = 16
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/ElementFlags/yearMonth
	DatePickerElementFlagYearMonth DatePickerElementFlags = 192
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/ElementFlags/yearMonthDay
	DatePickerElementFlagYearMonthDay DatePickerElementFlags = 224
)

/* debug [enums.gen.go]: Processing enum NSDatePickerMode (2 cases) */
// DatePickerMode - Constants that define whether the picker provides a single date, or a range of dates.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Mode
type DatePickerMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Mode/range
	DatePickerModeRange DatePickerMode = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDatePicker/Mode/single
	DatePickerModeSingle DatePickerMode = 0
)

/* debug [enums.gen.go]: Processing enum NSDirectionalRectEdge (6 cases) */
// DirectionalRectEdge enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge
type DirectionalRectEdge uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge/NSDirectionalRectEdgeNone
	DirectionalRectEdgeNone DirectionalRectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge/all
	DirectionalRectEdgeAll DirectionalRectEdge = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge/bottom
	DirectionalRectEdgeBottom DirectionalRectEdge = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge/leading
	DirectionalRectEdgeLeading DirectionalRectEdge = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge/top
	DirectionalRectEdgeTop DirectionalRectEdge = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDirectionalRectEdge/trailing
	DirectionalRectEdgeTrailing DirectionalRectEdge = 8
)

/* debug [enums.gen.go]: Processing enum NSDisplayGamut (2 cases) */
// DisplayGamut enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDisplayGamut
type DisplayGamut uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDisplayGamut/p3
	DisplayGamutP3 DisplayGamut = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDisplayGamut/sRGB
	DisplayGamutSRGB DisplayGamut = 1
)

/* debug [enums.gen.go]: Processing enum NSDocumentChangeType (7 cases) */
// DocumentChangeType - Values that indicate a document’s edit status.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType
type DocumentChangeType uint

const (
	// ChangeAutosaved - The document’s contents have been autosaved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType/changeAutosaved
	ChangeAutosaved DocumentChangeType = 4
	// ChangeCleared - Set change count to 0.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType/changeCleared
	ChangeCleared DocumentChangeType = 2
	// ChangeDiscardable - A discardable change has been done.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType/changeDiscardable
	ChangeDiscardable DocumentChangeType = 5
	// ChangeDone - Increment change count.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType/changeDone
	ChangeDone DocumentChangeType = 0
	// ChangeReadOtherContents - The document has been initialized with the contents of a file or file package other than the one whose location is in the   property, and therefore can’t possibly be synchronized with its persistent representation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType/changeReadOtherContents
	ChangeReadOtherContents DocumentChangeType = 3
	// ChangeRedone - A single change has been redone.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType/changeRedone
	ChangeRedone DocumentChangeType = 2
	// ChangeUndone - Decrement change count.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType/changeUndone
	ChangeUndone DocumentChangeType = 1
)

/* debug [enums.gen.go]: Processing enum NSSaveOperationType (7 cases) */
// SaveOperationType - Constants for specifying the type of document-save operation to perform.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType
type SaveOperationType uint

const (
	// AutosaveAsOperation - An operation that writes a document’s contents to a new file or file package even though the user has not explicitly requested it, then changes the document’s current location to point to the just-written file or file package.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveAsOperation
	AutosaveAsOperation SaveOperationType = 5
	// AutosaveElsewhereOperation - An operation that writes an autosave version of the file to a different location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveElsewhereOperation
	AutosaveElsewhereOperation SaveOperationType = 4
	// AutosaveInPlaceOperation - An operation that overwrites the document’s current contents with autosave data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveInPlaceOperation
	AutosaveInPlaceOperation SaveOperationType = 3
	// SaveAsOperation - An operation that writes the document’s contents to a new location and updates the document to point to that location
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveAsOperation
	SaveAsOperation SaveOperationType = 1
	// SaveOperation - An operation that overwrites a document’s file or file package with the document’s contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveOperation
	SaveOperation SaveOperationType = 0
	// SaveToOperation - An operation that writes a copy of the document’s contents to the specified location, without changing the original document’s location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveToOperation
	SaveToOperation SaveOperationType = 2
	// AutosaveOperation - Old name for the   operation type.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSaveOperationType/NSAutosaveOperation
	AutosaveOperation SaveOperationType = 6
)

/* debug [enums.gen.go]: Processing enum NSDragOperation (9 cases) */
// DragOperation - A group of constants that represent which operations the dragging source can perform on dragging items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation
type DragOperation uint

const (
	// DragOperationNone - A constant that indicates that the drag cannot perform any operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/NSDragOperationNone
	DragOperationNone DragOperation = 0
	// DragOperationAll - Use   instead.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/all
	DragOperationAll DragOperation = 34
	// DragOperationCopy - A constant that indicates the drag can copy the data that the image represents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/copy
	DragOperationCopy DragOperation = 1
	// DragOperationDelete - A constant that indicates the drag can delete the data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/delete
	DragOperationDelete DragOperation = 32
	// DragOperationEvery - A constant that indicates that drag can perform all of the drag operations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/every
	DragOperationEvery DragOperation = 0
	// DragOperationGeneric - A constant that indicates the destination can define the drag operation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/generic
	DragOperationGeneric DragOperation = 4
	// DragOperationLink - A constant that indicates the drag can share the data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/link
	DragOperationLink DragOperation = 2
	// DragOperationMove - A constant that indicates the drag can move the data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/move
	DragOperationMove DragOperation = 16
	// DragOperationPrivate - A constant that indicates the source and destination negotiate the drag operation privately.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDragOperation/private
	DragOperationPrivate DragOperation = 8
)

/* debug [enums.gen.go]: Processing enum NSDraggingContext (2 cases) */
// DraggingContext - Constants that specify whether a drag terminates within or outside the application.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingContext
type DraggingContext uint

const (
	// DraggingContextOutsideApplication - A constant that indicates dragging terminates outside the application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingContext/outsideApplication
	DraggingContextOutsideApplication DraggingContext = 0
	// DraggingContextWithinApplication - A constant that indicates dragging terminates within the application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingContext/withinApplication
	DraggingContextWithinApplication DraggingContext = 1
)

/* debug [enums.gen.go]: Processing enum NSDraggingFormation (5 cases) */
// DraggingFormation - Constants that control the visual format of multiple dragging items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingFormation
type DraggingFormation uint

const (
	// DraggingFormationDefault - A constant that represents the system determined formation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingFormation/default
	DraggingFormationDefault DraggingFormation = 0
	// DraggingFormationList - A constant that represents a list formation, so drag images display vertically, non-overlapping with the left edges aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingFormation/list
	DraggingFormationList DraggingFormation = 3
	// DraggingFormationNone - A constant that represents no custom formation, so drag images maintain their set positions relative to each other.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingFormation/none
	DraggingFormationNone DraggingFormation = 1
	// DraggingFormationPile - A constant that represents a pile formation, so drag images display on top of each other with random rotations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingFormation/pile
	DraggingFormationPile DraggingFormation = 2
	// DraggingFormationStack - A constant that represents a stack formation, so drag images display overlapping diagonally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingFormation/stack
	DraggingFormationStack DraggingFormation = 4
)

/* debug [enums.gen.go]: Processing enum NSDraggingItemEnumerationOptions (2 cases) */
// DraggingItemEnumerationOptions - A group of constants that specify options to use when enumerating dragging items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItemEnumerationOptions
type DraggingItemEnumerationOptions uint

const (
	// DraggingItemEnumerationClearNonenumeratedImages - A constant that indicates the enumeration clears the image components provider for all dragging items that don’t meet the classes and search options criteria.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItemEnumerationOptions/clearNonenumeratedImages
	DraggingItemEnumerationClearNonenumeratedImages DraggingItemEnumerationOptions = 65536
	// DraggingItemEnumerationConcurrent - A constant that indicates the enumeration processes dragging items concurrently.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDraggingItemEnumerationOptions/concurrent
	DraggingItemEnumerationConcurrent DraggingItemEnumerationOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSDrawerState (4 cases) */
// DrawerState - These constants specify the possible states of a drawer.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/State-swift.enum
type DrawerState uint

const (
	// DrawerClosedState - The drawer is closed (not visible onscreen).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/State-swift.enum/closedState
	DrawerClosedState DrawerState = 0
	// DrawerClosingState - The drawer is in the process of closing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/State-swift.enum/closingState
	DrawerClosingState DrawerState = 3
	// DrawerOpenState - The drawer is open (visible onscreen).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/State-swift.enum/openState
	DrawerOpenState DrawerState = 2
	// DrawerOpeningState - The drawer is in the process of opening.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDrawer/State-swift.enum/openingState
	DrawerOpeningState DrawerState = 1
)

/* debug [enums.gen.go]: Processing enum NSEventButtonMask (1 cases) */
// EventButtonMask - Constants you use to identify the activated tablet buttons in an event.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ButtonMask-swift.struct
type EventButtonMask uint

const (
	// EventButtonMaskPenUpperSide - A mask that matches the button on the upper side of the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ButtonMask-swift.struct/penUpperSide
	EventButtonMaskPenUpperSide EventButtonMask = 4
)

/* debug [enums.gen.go]: Processing enum NSEventSubtype (10 cases) */
// EventSubtype - Subtypes for various types of events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype
type EventSubtype uint

const (
	// EventSubtypeApplicationActivated - An app-activation event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/applicationActivated
	EventSubtypeApplicationActivated EventSubtype = 1
	// EventSubtypeApplicationDeactivated - An app-deactivation event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/applicationDeactivated
	EventSubtypeApplicationDeactivated EventSubtype = 2
	// EventSubtypeMouseEvent - A mouse event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/mouseEvent
	EventSubtypeMouseEvent EventSubtype = 0
	// EventSubtypePowerOff - An event that indicates a system shutdown or restart operation is in progress.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/powerOff
	EventSubtypePowerOff EventSubtype = 1
	// EventSubtypeScreenChanged - An event that indicates a window changed screens.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/screenChanged
	EventSubtypeScreenChanged EventSubtype = 8
	// EventSubtypeTabletPoint - A tablet-pointer event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/tabletPoint
	EventSubtypeTabletPoint EventSubtype = 1
	// EventSubtypeTabletProximity - A tablet-proximity event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/tabletProximity
	EventSubtypeTabletProximity EventSubtype = 2
	// EventSubtypeTouch - A touch event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/touch
	EventSubtypeTouch EventSubtype = 3
	// EventSubtypeWindowExposed - An event that indicates a window’s contents are visible again.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/windowExposed
	EventSubtypeWindowExposed EventSubtype = 0
	// EventSubtypeWindowMoved - An event that indicates a window moved.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventSubtype/windowMoved
	EventSubtypeWindowMoved EventSubtype = 4
)

/* debug [enums.gen.go]: Processing enum NSEventType (35 cases) */
// EventType - Constants for the types of events that responder objects can handle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType
type EventType uint

const (
	// EventTypeAppKitDefined - An AppKit-related event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/appKitDefined
	EventTypeAppKitDefined EventType = 13
	// EventTypeApplicationDefined - An app-defined event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/applicationDefined
	EventTypeApplicationDefined EventType = 15
	// EventTypeBeginGesture - An event marking the beginning of a gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/beginGesture
	EventTypeBeginGesture EventType = 32
	// EventTypeChangeMode - The user changed the mode of a connected device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/changeMode
	EventTypeChangeMode EventType = 38
	// EventTypeCursorUpdate - An event that updates the cursor.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/cursorUpdate
	EventTypeCursorUpdate EventType = 17
	// EventTypeDirectTouch - The user touched a portion of the touch bar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/directTouch
	EventTypeDirectTouch EventType = 37
	// EventTypeEndGesture - An event that marks the end of a gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/endGesture
	EventTypeEndGesture EventType = 33
	// EventTypeFlagsChanged - The event flags changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/flagsChanged
	EventTypeFlagsChanged EventType = 12
	// EventTypeGesture - The user performed a nonspecific type of gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/gesture
	EventTypeGesture EventType = 28
	// EventTypeKeyDown - The user pressed a key on the keyboard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/keyDown
	EventTypeKeyDown EventType = 10
	// EventTypeKeyUp - The user released a key on the keyboard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/keyUp
	EventTypeKeyUp EventType = 11
	// EventTypeLeftMouseDown - The user pressed the left mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/leftMouseDown
	EventTypeLeftMouseDown EventType = 1
	// EventTypeLeftMouseDragged - The user moved the mouse while holding down the left mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/leftMouseDragged
	EventTypeLeftMouseDragged EventType = 6
	// EventTypeLeftMouseUp - The user released the left mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/leftMouseUp
	EventTypeLeftMouseUp EventType = 2
	// EventTypeMagnify - The user performed a pinch-open or pinch-close gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/magnify
	EventTypeMagnify EventType = 29
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/mouseCancelled
	EventTypeMouseCancelled EventType = 39
	// EventTypeMouseEntered - The cursor entered a well-defined area, such as a view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/mouseEntered
	EventTypeMouseEntered EventType = 8
	// EventTypeMouseExited - The cursor exited a well-defined area, such as a view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/mouseExited
	EventTypeMouseExited EventType = 9
	// EventTypeMouseMoved - The user moved the mouse in a way that caused the cursor to move onscreen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/mouseMoved
	EventTypeMouseMoved EventType = 5
	// EventTypeOtherMouseDown - The user pressed a tertiary mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/otherMouseDown
	EventTypeOtherMouseDown EventType = 25
	// EventTypeOtherMouseDragged - The user moved the mouse while holding down a tertiary mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/otherMouseDragged
	EventTypeOtherMouseDragged EventType = 27
	// EventTypeOtherMouseUp - The user released a tertiary mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/otherMouseUp
	EventTypeOtherMouseUp EventType = 26
	// EventTypePeriodic - An event that provides execution time to periodic tasks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/periodic
	EventTypePeriodic EventType = 16
	// EventTypePressure - An event that reports a change in pressure on a pressure-sensitive device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/pressure
	EventTypePressure EventType = 36
	// EventTypeQuickLook - An event that initiates a Quick Look request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/quickLook
	EventTypeQuickLook EventType = 35
	// EventTypeRightMouseDown - The user pressed the right mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/rightMouseDown
	EventTypeRightMouseDown EventType = 3
	// EventTypeRightMouseDragged - The user moved the mouse while holding down the right mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/rightMouseDragged
	EventTypeRightMouseDragged EventType = 7
	// EventTypeRightMouseUp - The user released the right mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/rightMouseUp
	EventTypeRightMouseUp EventType = 4
	// EventTypeRotate - The user performed a rotate gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/rotate
	EventTypeRotate EventType = 31
	// EventTypeScrollWheel - The scroll wheel position changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/scrollWheel
	EventTypeScrollWheel EventType = 22
	// EventTypeSmartMagnify - The user performed a smart-zoom gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/smartMagnify
	EventTypeSmartMagnify EventType = 34
	// EventTypeSwipe - The user performed a swipe gesture.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/swipe
	EventTypeSwipe EventType = 30
	// EventTypeSystemDefined - A system-related event occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/systemDefined
	EventTypeSystemDefined EventType = 14
	// EventTypeTabletPoint - The user touched a point on a tablet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/tabletPoint
	EventTypeTabletPoint EventType = 23
	// EventTypeTabletProximity - A pointing device is near, but not touching, the associated tablet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventType/tabletProximity
	EventTypeTabletProximity EventType = 24
)

/* debug [enums.gen.go]: Processing enum NSEventMask (12 cases) */
// EventMask - Constants that you use to filter out specific event types from the stream of incoming events.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask
type EventMask uint

const (
	// EventMaskBeginGesture - A mask for begin-gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/beginGesture
	EventMaskBeginGesture EventMask = 4
	// EventMaskDirectTouch - A mask for touch events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/directTouch
	EventMaskDirectTouch EventMask = 8
	// EventMaskEndGesture - A mask for end-gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/endGesture
	EventMaskEndGesture EventMask = 5
	// EventMaskGesture - A mask for generic gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/gesture
	EventMaskGesture EventMask = 0
	// EventMaskLeftMouseUp - A mask for left mouse-up events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/leftMouseUp
	EventMaskLeftMouseUp EventMask = 0
	// EventMaskMagnify - A mask for magnify-gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/magnify
	EventMaskMagnify EventMask = 1
	// EventMaskPressure - A mask for pressure-change events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/pressure
	EventMaskPressure EventMask = 7
	// EventMaskRotate - A mask for rotate-gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/rotate
	EventMaskRotate EventMask = 3
	// EventMaskSmartMagnify - A mask for smart-zoom gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/smartMagnify
	EventMaskSmartMagnify EventMask = 6
	// EventMaskSwipe - A mask for swipe-gesture events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/swipe
	EventMaskSwipe EventMask = 2
	// EventMaskTabletPoint - A mask for tablet-point events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/tabletPoint
	EventMaskTabletPoint EventMask = 0
	// EventMaskTabletProximity - A mask for tablet-proximity events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/EventTypeMask/tabletProximity
	EventMaskTabletProximity EventMask = 0
)

/* debug [enums.gen.go]: Processing enum NSEventGestureAxis (3 cases) */
// EventGestureAxis - Constants that specify the direction of travel for a gesture.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis
type EventGestureAxis uint

const (
	// EventGestureAxisHorizontal - The horizontal axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis/horizontal
	EventGestureAxisHorizontal EventGestureAxis = 1
	// EventGestureAxisNone - No specific axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis/none
	EventGestureAxisNone EventGestureAxis = 0
	// EventGestureAxisVertical - The vertical axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/GestureAxis/vertical
	EventGestureAxisVertical EventGestureAxis = 2
)

/* debug [enums.gen.go]: Processing enum NSEventModifierFlags (9 cases) */
// EventModifierFlags - Flags that represent key states in an event object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/ModifierFlags-swift.struct
type EventModifierFlags uint

const (
	EventModifierFlagCapsLock EventModifierFlags = 65536
	EventModifierFlagShift EventModifierFlags = 131072
	EventModifierFlagControl EventModifierFlags = 262144
	EventModifierFlagOption EventModifierFlags = 524288
	EventModifierFlagCommand EventModifierFlags = 1048576
	EventModifierFlagNumericPad EventModifierFlags = 2097152
	EventModifierFlagHelp EventModifierFlags = 4194304
	EventModifierFlagFunction EventModifierFlags = 8388608
	EventModifierFlagDeviceIndependentFlagsMask EventModifierFlags = 4294901760
)

/* debug [enums.gen.go]: Processing enum NSEventPhase (7 cases) */
// EventPhase - Constants that represent the possible phases during an event phase.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct
type EventPhase uint

const (
	// EventPhaseBegan - An event phase has begun.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/began
	EventPhaseBegan EventPhase = 1
	// EventPhaseCancelled - The system canceled the event phase.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/cancelled
	EventPhaseCancelled EventPhase = 1
	// EventPhaseChanged - An event phase has changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/changed
	EventPhaseChanged EventPhase = 1
	// EventPhaseEnded - The event phase ended.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/ended
	EventPhaseEnded EventPhase = 1
	// EventPhaseMayBegin - The system event phase may begin.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/mayBegin
	EventPhaseMayBegin EventPhase = 1
	// EventPhaseStationary - An event phase is in progress but hasn’t moved since the previous event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/Phase-swift.struct/stationary
	EventPhaseStationary EventPhase = 1
	// EventPhaseNone - The event is not associated with a phase.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEventPhase/NSEventPhaseNone
	EventPhaseNone EventPhase = 0
)

/* debug [enums.gen.go]: Processing enum NSPressureBehavior (7 cases) */
// PressureBehavior - These constants describe the behavior and progression of a pressure gesture.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum
type PressureBehavior int

const (
	// PressureBehaviorPrimaryAccelerator - A pressure gesture’s behavior begins on left mouse-down events. A maximum of one stage is supported, and a stage transition animation occurs when moving from stage 1 to stage 0. Actuations occur during the mouse-down and mouse-up events when this behavior is configured. This configuration uses specific pressure mappings that are ideal for controlling speed as variable pressure occurs between the mouse-down and mouse-up events. The   class uses this behavior. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryAccelerator
	PressureBehaviorPrimaryAccelerator PressureBehavior = 3
	// PressureBehaviorPrimaryClick - A pressure gesture’s behavior begins on left mouse-down events. A maximum of one stage is supported, and a stage transition animation occurs when moving from stage 1 to stage 0. Actuations (haptic feedback the user feels) occur during mouse-down and mouse-up events when this behavior is configured. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryClick
	PressureBehaviorPrimaryClick PressureBehavior = 1
	// PressureBehaviorPrimaryDeepClick - A pressure gesture’s behavior begins on left mouse-down events. Two stages are supported, and a stage transition animation may occur when moving between stages—from stage 1 to stage 0, stage 1 to stage 2, stage 2 to stage 1, and stage 2 to stage 0. With this behavior type, stage 2 becomes disabled once dragging occurs. When this behavior is configured, actuations occur during the mouse-down and mouse-up events, as well as when force click is activated and released when entering or exiting stage 2. This configuration is ideal for responding to force clicks. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryDeepClick
	PressureBehaviorPrimaryDeepClick PressureBehavior = 5
	// PressureBehaviorPrimaryDeepDrag - A pressure gesture’s behavior begins on left mouse-down events. Two stages are supported, and a stage transition animation may occur when moving between stages—from stage 1 to stage 0, stage 1 to stage 2, stage 2 to stage 1, or stage 2 to stage 0. Actuations occur during the mouse-down and mouse-up events, as well as during the transitions up and down between stage 1 and stage 2, when this behavior is configured. This configuration is ideal for responding to force clicks during drag operations. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryDeepDrag
	PressureBehaviorPrimaryDeepDrag PressureBehavior = 6
	// PressureBehaviorPrimaryDefault - This is the default behavior when a pressure gesture’s behavior has not been explicitly configured. In OS X 10.10.3, this behavior defaults to the behavior of  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryDefault
	PressureBehaviorPrimaryDefault PressureBehavior = 0
	// PressureBehaviorPrimaryGeneric - A pressure gesture’s behavior begins on left mouse-down events. A maximum of one stage is supported, and a stage transition animation occurs when moving from stage 1 to stage 0. Actuations occur during the mouse-down and mouse-up events when this behavior is configured. This configuration is ideal for drawing, painting, and general use. Variable pressure occurs throughout the course of the gesture. Note that the pressure gesture operates on a separate event stream from the mouse events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/primaryGeneric
	PressureBehaviorPrimaryGeneric PressureBehavior = 2
	// PressureBehaviorUnknown - A pressure gesture’s behavior is not known, perhaps because the input device does not support pressure gestures.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/PressureBehavior-swift.enum/unknown
	PressureBehaviorUnknown PressureBehavior = -1
)

/* debug [enums.gen.go]: Processing enum NSEventSwipeTrackingOptions (2 cases) */
// EventSwipeTrackingOptions - Constants that specify swipe-tracking options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions
type EventSwipeTrackingOptions uint

const (
	// EventSwipeTrackingClampGestureAmount - Don’t allow gestureAmount to go beyond +/-1.0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions/clampGestureAmount
	EventSwipeTrackingClampGestureAmount EventSwipeTrackingOptions = 1
	// EventSwipeTrackingLockDirection - Clamp gestureAmount to 0 if the user starts to swipe in the opposite direction than they started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSEvent/SwipeTrackingOptions/lockDirection
	EventSwipeTrackingLockDirection EventSwipeTrackingOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSFindPanelAction (10 cases) */
// FindPanelAction - These constants define the tags for 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction
type FindPanelAction uint

const (
	// FindPanelActionNext - Finds the next instance of the queried text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/next
	FindPanelActionNext FindPanelAction = 2
	// FindPanelActionPrevious - Finds the previous instance of the queried text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/previous
	FindPanelActionPrevious FindPanelAction = 3
	// FindPanelActionReplace - Replaces a single query instance within the text view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/replace
	FindPanelActionReplace FindPanelAction = 5
	// FindPanelActionReplaceAll - Replaces all query instances within the text view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/replaceAll
	FindPanelActionReplaceAll FindPanelAction = 4
	// FindPanelActionReplaceAllInSelection - Replaces all query instances within the selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/replaceAllInSelection
	FindPanelActionReplaceAllInSelection FindPanelAction = 8
	// FindPanelActionReplaceAndFind - Replaces a single query instance and finds the next.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/replaceAndFind
	FindPanelActionReplaceAndFind FindPanelAction = 6
	// FindPanelActionSelectAll - Selects all query instances in the text view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/selectAll
	FindPanelActionSelectAll FindPanelAction = 9
	// FindPanelActionSelectAllInSelection - Selects all query instances within the selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/selectAllInSelection
	FindPanelActionSelectAllInSelection FindPanelAction = 10
	// FindPanelActionSetFindString - Sets the query string to the current selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/setFindString
	FindPanelActionSetFindString FindPanelAction = 7
	// FindPanelActionShowFindPanel - Displays the find panel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelAction/showFindPanel
	FindPanelActionShowFindPanel FindPanelAction = 1
)

/* debug [enums.gen.go]: Processing enum NSFindPanelSubstringMatchType (4 cases) */
// FindPanelSubstringMatchType - The type of substring matching used by the Find panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelSubstringMatchType
type FindPanelSubstringMatchType uint

const (
	// FindPanelSubstringMatchTypeContains - Finds a word containing the search string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelSubstringMatchType/contains
	FindPanelSubstringMatchTypeContains FindPanelSubstringMatchType = 0
	// FindPanelSubstringMatchTypeEndsWith - Finds a word ending with the search string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelSubstringMatchType/endsWith
	FindPanelSubstringMatchTypeEndsWith FindPanelSubstringMatchType = 3
	// FindPanelSubstringMatchTypeFullWord - Finds a word exactly matching the search string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelSubstringMatchType/fullWord
	FindPanelSubstringMatchTypeFullWord FindPanelSubstringMatchType = 2
	// FindPanelSubstringMatchTypeStartsWith - Finds a word starting with the search string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFindPanelSubstringMatchType/startsWith
	FindPanelSubstringMatchTypeStartsWith FindPanelSubstringMatchType = 1
)

/* debug [enums.gen.go]: Processing enum NSFocusRingPlacement (3 cases) */
// FocusRingPlacement - Constants that indicate how the system draws the focus ring.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement
type FocusRingPlacement uint

const (
	// FocusRingAbove - Draw the focus ring over an image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/above
	FocusRingAbove FocusRingPlacement = 2
	// FocusRingBelow - Draw the focus ring under text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/below
	FocusRingBelow FocusRingPlacement = 1
	// FocusRingOnly - Draw the focus ring if you don’t have an image or text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingPlacement/only
	FocusRingOnly FocusRingPlacement = 0
)

/* debug [enums.gen.go]: Processing enum NSFocusRingType (3 cases) */
// FocusRingType - Constants that describe the style of the focus ring.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType
type FocusRingType uint

const (
	// FocusRingTypeDefault - The default focus ring type for a view or cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType/default
	FocusRingTypeDefault FocusRingType = 0
	// FocusRingTypeExterior - The standard Aqua focus ring.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType/exterior
	FocusRingTypeExterior FocusRingType = 2
	// FocusRingTypeNone - No focus ring.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFocusRingType/none
	FocusRingTypeNone FocusRingType = 1
)

/* debug [enums.gen.go]: Processing enum NSFontAction (8 cases) */
// FontAction - Actions that modify a font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAction
type FontAction uint

const (
	// AddTraitFontAction - Converts the font to have an additional trait using  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAction/addTraitFontAction
	AddTraitFontAction FontAction = 2
	// HeavierFontAction - Converts the font to a heavier weight using  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAction/heavierFontAction
	HeavierFontAction FontAction = 5
	// LighterFontAction - Converts the font to a lighter weight using  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAction/lighterFontAction
	LighterFontAction FontAction = 6
	// NoFontChangeAction - No action; the font is returned unchanged.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAction/noFontChangeAction
	NoFontChangeAction FontAction = 0
	// RemoveTraitFontAction - Converts the font to remove a trait using  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAction/removeTraitFontAction
	RemoveTraitFontAction FontAction = 7
	// SizeDownFontAction - Converts the font to a smaller size using  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAction/sizeDownFontAction
	SizeDownFontAction FontAction = 4
	// SizeUpFontAction - Converts the font to a larger size using  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAction/sizeUpFontAction
	SizeUpFontAction FontAction = 3
	// ViaPanelFontAction - Converts the font according to the   method  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAction/viaPanelFontAction
	ViaPanelFontAction FontAction = 1
)

/* debug [enums.gen.go]: Processing enum NSFontAssetRequestOptions (1 cases) */
// FontAssetRequestOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/Options
type FontAssetRequestOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontAssetRequest/Options/usesStandardUI
	FontAssetRequestOptionUsesStandardUI FontAssetRequestOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSFontCollectionVisibility (3 cases) */
// FontCollectionVisibility - Constants that specify the visibility of font collections.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/Visibility
type FontCollectionVisibility uint

const (
	// FontCollectionVisibilityComputer - The font collection is visible to all users and is stored persistently.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/Visibility/computer
	FontCollectionVisibilityComputer FontCollectionVisibility = 4
	// FontCollectionVisibilityProcess - The font collection is visible within this process and is not persistent.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/Visibility/process
	FontCollectionVisibilityProcess FontCollectionVisibility = 1
	// FontCollectionVisibilityUser - The font collection is visible to all processes and is stored persistently.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollection/Visibility/user
	FontCollectionVisibilityUser FontCollectionVisibility = 2
)

/* debug [enums.gen.go]: Processing enum NSFontCollectionOptions (1 cases) */
// FontCollectionOptions - Constants that support font collection management.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollectionOptions
type FontCollectionOptions uint

const (
	// FontCollectionApplicationOnlyMask - Makes the collection available only to the application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontCollectionOptions/applicationOnlyMask
	FontCollectionApplicationOnlyMask FontCollectionOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSFontDescriptorSymbolicTraits (22 cases) */
// FontDescriptorSymbolicTraits - A symbolic description of the stylistic aspects of a font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct
type FontDescriptorSymbolicTraits uint

const (
	// FontDescriptorTraitUIOptimized - The font synthesizes appropriate attributes for user interface rendering, such as in control titles, if necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/UIOptimized
	FontDescriptorTraitUIOptimized FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitBold - The font’s style is boldface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/bold
	FontDescriptorTraitBold FontDescriptorSymbolicTraits = 1
	// FontDescriptorClassClarendonSerifs - The font’s characters include variations of old style and transitional serifs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classClarendonSerifs
	FontDescriptorClassClarendonSerifs FontDescriptorSymbolicTraits = 4
	// FontDescriptorClassFreeformSerifs - The font’s characters include serifs, and don’t generally fit within other serif design classifications.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classFreeformSerifs
	FontDescriptorClassFreeformSerifs FontDescriptorSymbolicTraits = 7
	// FontDescriptorClassMask - The font family class mask that you use to access font descriptor values.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classMask
	FontDescriptorClassMask FontDescriptorSymbolicTraits = 4026531840
	// FontDescriptorClassModernSerifs - The font’s characters include serifs, and reflect the Latin printing style of the 20th century.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classModernSerifs
	FontDescriptorClassModernSerifs FontDescriptorSymbolicTraits = 3
	// FontDescriptorClassOldStyleSerifs - The font’s characters include serifs, and reflect the Latin printing style of the 15th to 17th centuries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classOldStyleSerifs
	FontDescriptorClassOldStyleSerifs FontDescriptorSymbolicTraits = 1
	// FontDescriptorClassOrnamentals - The font’s characters use highly decorated or stylized character shapes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classOrnamentals
	FontDescriptorClassOrnamentals FontDescriptorSymbolicTraits = 9
	// FontDescriptorClassSansSerif - The font’s characters don’t have serifs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classSansSerif
	FontDescriptorClassSansSerif FontDescriptorSymbolicTraits = 8
	// FontDescriptorClassScripts - The font’s characters simulate handwriting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classScripts
	FontDescriptorClassScripts FontDescriptorSymbolicTraits = 10
	// FontDescriptorClassSlabSerifs - The font’s characters use square transitions, without brackets, between strokes and serifs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classSlabSerifs
	FontDescriptorClassSlabSerifs FontDescriptorSymbolicTraits = 5
	// FontDescriptorClassSymbolic - The font’s characters consist mainly of symbols rather than letters and numbers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classSymbolic
	FontDescriptorClassSymbolic FontDescriptorSymbolicTraits = 12
	// FontDescriptorClassTransitionalSerifs - The font’s characters include serifs, and reflect the Latin printing style of the 18th to 19th centuries.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/classTransitionalSerifs
	FontDescriptorClassTransitionalSerifs FontDescriptorSymbolicTraits = 2
	// FontDescriptorTraitCondensed - The font’s characters have a condensed width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/condensed
	FontDescriptorTraitCondensed FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitExpanded - The font’s characters have an expanded width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/expanded
	FontDescriptorTraitExpanded FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitItalic - The font’s style is italic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/italic
	FontDescriptorTraitItalic FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitLooseLeading - The font uses a leading value that’s greater than the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/looseLeading
	FontDescriptorTraitLooseLeading FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitMonoSpace - The font’s characters all have the same width.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/monoSpace
	FontDescriptorTraitMonoSpace FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitTightLeading - The font uses a leading value that’s less than the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/tightLeading
	FontDescriptorTraitTightLeading FontDescriptorSymbolicTraits = 1
	// FontDescriptorTraitVertical - The font uses vertical glyph variants and metrics.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SymbolicTraits-swift.struct/vertical
	FontDescriptorTraitVertical FontDescriptorSymbolicTraits = 1
	// FontDescriptorClassUnknown - The font has no design classification.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptorSymbolicTraits/NSFontDescriptorClassUnknown
	FontDescriptorClassUnknown FontDescriptorSymbolicTraits = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptorSymbolicTraits/NSFontDescriptorTraitEmphasized
	FontDescriptorTraitEmphasized FontDescriptorSymbolicTraits = 2
)

/* debug [enums.gen.go]: Processing enum NSFontPanelModeMask (1 cases) */
// FontPanelModeMask enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/ModeMask
type FontPanelModeMask uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontPanel/ModeMask/face
	FontPanelModeMaskFace FontPanelModeMask = 1
)

/* debug [enums.gen.go]: Processing enum NSFontRenderingMode (4 cases) */
// FontRenderingMode - The font rendering mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode
type FontRenderingMode uint

const (
	// FontAntialiasedIntegerAdvancementsRenderingMode - Specifies antialiased, integer advancements rendering mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode/antialiasedIntegerAdvancementsRenderingMode
	FontAntialiasedIntegerAdvancementsRenderingMode FontRenderingMode = 3
	// FontAntialiasedRenderingMode - Specifies antialiased, floating-point advancements rendering mode (synonymous with printerFont).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode/antialiasedRenderingMode
	FontAntialiasedRenderingMode FontRenderingMode = 1
	// FontDefaultRenderingMode - Determines the actual mode based on the user preference settings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode/defaultRenderingMode
	FontDefaultRenderingMode FontRenderingMode = 0
	// FontIntegerAdvancementsRenderingMode - Specifies integer advancements rendering mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontRenderingMode/integerAdvancementsRenderingMode
	FontIntegerAdvancementsRenderingMode FontRenderingMode = 2
)

/* debug [enums.gen.go]: Processing enum NSFontTraitMask (12 cases) */
// FontTraitMask - Constants for isolating specific traits of a font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask
type FontTraitMask uint

const (
	// BoldFontMask - A mask that specifies a bold font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/boldFontMask
	BoldFontMask FontTraitMask = 2
	// CompressedFontMask - A mask that specifies a compressed font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/compressedFontMask
	CompressedFontMask FontTraitMask = 512
	// CondensedFontMask - A mask that specifies a condensed font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/condensedFontMask
	CondensedFontMask FontTraitMask = 64
	// ExpandedFontMask - A mask that specifies an expanded font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/expandedFontMask
	ExpandedFontMask FontTraitMask = 32
	// FixedPitchFontMask - A mask that specifies a fixed pitch font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/fixedPitchFontMask
	FixedPitchFontMask FontTraitMask = 1024
	// ItalicFontMask - A mask that specifies an italic font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/italicFontMask
	ItalicFontMask FontTraitMask = 1
	// NarrowFontMask - A mask that specifies a narrow font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/narrowFontMask
	NarrowFontMask FontTraitMask = 16
	// NonStandardCharacterSetFontMask - A mask that specifies a font containing a non-standard character set.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/nonStandardCharacterSetFontMask
	NonStandardCharacterSetFontMask FontTraitMask = 8
	// PosterFontMask - A mask that specifies a poster-style font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/posterFontMask
	PosterFontMask FontTraitMask = 256
	// SmallCapsFontMask - A mask that specifies a small-caps font.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/smallCapsFontMask
	SmallCapsFontMask FontTraitMask = 128
	// UnboldFontMask - A mask that specifies a font that is not bold.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/unboldFontMask
	UnboldFontMask FontTraitMask = 4
	// UnitalicFontMask - A mask that specifies a font that is not italic.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontTraitMask/unitalicFontMask
	UnitalicFontMask FontTraitMask = 16777216
)

/* debug [enums.gen.go]: Processing enum NSGestureRecognizerState (7 cases) */
// GestureRecognizerState - The current state of the gesture recognizer.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum
type GestureRecognizerState uint

const (
	// GestureRecognizerStateBegan - The gesture recognizer has recognized a sequence of events as a continuous gesture. It calls its action method at the next cycle of the run loop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/began
	GestureRecognizerStateBegan GestureRecognizerState = 1
	// GestureRecognizerStateCancelled - The gesture recognizer received events that resulted in the cancellation of a continuous gesture. It calls its action method at the next cycle of the run loop and resets its state to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/cancelled
	GestureRecognizerStateCancelled GestureRecognizerState = 4
	// GestureRecognizerStateChanged - The gesture recognizer has detected a change to a continuous gesture. It calls its action method at the next cycle of the run loop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/changed
	GestureRecognizerStateChanged GestureRecognizerState = 2
	// GestureRecognizerStateEnded - The gesture recognizer has detected the end of a continuous gesture. It calls its action method at the next cycle of the run loop and resets its state to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/ended
	GestureRecognizerStateEnded GestureRecognizerState = 3
	// GestureRecognizerStateFailed - The gesture recognizer failed to recognize its gesture and will not call its action method. The gesture recognizer resets itself to the   state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/failed
	GestureRecognizerStateFailed GestureRecognizerState = 5
	// GestureRecognizerStatePossible - The gesture recognizer has not yet recognized its gesture but may be evaluating events. This is the default state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/possible
	GestureRecognizerStatePossible GestureRecognizerState = 0
	// GestureRecognizerStateRecognized - The gesture recognizer successfully recognized its gesture. It calls its action method at the next cycle of the run loop and resets its state to  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/State-swift.enum/recognized
	GestureRecognizerStateRecognized GestureRecognizerState = 0
)

/* debug [enums.gen.go]: Processing enum NSGlassEffectViewStyle (2 cases) */
// GlassEffectViewStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/Style-swift.enum
type GlassEffectViewStyle uint

const (
	// GlassEffectViewStyleClear - Clear glass effect style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/Style-swift.enum/clear
	GlassEffectViewStyleClear GlassEffectViewStyle = 0
	// GlassEffectViewStyleRegular - Standard glass effect style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlassEffectView/Style-swift.enum/regular
	GlassEffectViewStyleRegular GlassEffectViewStyle = 0
)

/* debug [enums.gen.go]: Processing enum NSGlyphInscription (5 cases) */
// GlyphInscription - Constants that specify how a glyph is laid out relative to the previous glyph.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription
type GlyphInscription uint

const (
	// GlyphInscribeAbove - A glyph is rendered above the previous glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeAbove
	GlyphInscribeAbove GlyphInscription = 2
	// GlyphInscribeBase - A base glyph; a character that the font can represent with a single glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeBase
	GlyphInscribeBase GlyphInscription = 0
	// GlyphInscribeBelow - A glyph is rendered below the previous glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeBelow
	GlyphInscribeBelow GlyphInscription = 1
	// GlyphInscribeOverBelow - A glyph is rendered on top and below the previous glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeOverBelow
	GlyphInscribeOverBelow GlyphInscription = 4
	// GlyphInscribeOverstrike - A glyph is rendered on top of the previous glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyphInscription/NSGlyphInscribeOverstrike
	GlyphInscribeOverstrike GlyphInscription = 3
)

/* debug [enums.gen.go]: Processing enum NSGridCellPlacement (8 cases) */
// GridCellPlacement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement
type GridCellPlacement uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement/bottom
	GridCellPlacementBottom GridCellPlacement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement/center
	GridCellPlacementCenter GridCellPlacement = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement/fill
	GridCellPlacementFill GridCellPlacement = 5
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement/inherited
	GridCellPlacementInherited GridCellPlacement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement/leading
	GridCellPlacementLeading GridCellPlacement = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement/none
	GridCellPlacementNone GridCellPlacement = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement/top
	GridCellPlacementTop GridCellPlacement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridCell/Placement/trailing
	GridCellPlacementTrailing GridCellPlacement = 3
)

/* debug [enums.gen.go]: Processing enum NSGridRowAlignment (4 cases) */
// GridRowAlignment enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/Alignment
type GridRowAlignment uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/Alignment/firstBaseline
	GridRowAlignmentFirstBaseline GridRowAlignment = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/Alignment/inherited
	GridRowAlignmentInherited GridRowAlignment = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/Alignment/lastBaseline
	GridRowAlignmentLastBaseline GridRowAlignment = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGridRow/Alignment/none
	GridRowAlignmentNone GridRowAlignment = 1
)

/* debug [enums.gen.go]: Processing enum NSHapticFeedbackPattern (3 cases) */
// HapticFeedbackPattern - A pattern of haptic feedback to be provided to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern
type HapticFeedbackPattern uint

const (
	// HapticFeedbackPatternAlignment - A haptic feedback pattern to be used in response to the alignment of an object the user is dragging around. For example, this pattern of feedback could be used in a drawing app when the user drags a shape into alignment with another shape. Other scenarios where this type of feedback could be used might include scaling an object to fit within specific dimensions, positioning an object at a preferred location, or reaching the beginning/minimum or end/maximum of something, such as a track view in an audio/video app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern/alignment
	HapticFeedbackPatternAlignment HapticFeedbackPattern = 1
	// HapticFeedbackPatternGeneric - A general haptic feedback pattern. Use this when no other feedback patterns apply.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern/generic
	HapticFeedbackPatternGeneric HapticFeedbackPattern = 0
	// HapticFeedbackPatternLevelChange - A haptic feedback pattern to be used as the user moves between discrete levels of pressure. This pattern of feedback is used by multilevel accelerator buttons (class  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/FeedbackPattern/levelChange
	HapticFeedbackPatternLevelChange HapticFeedbackPattern = 2
)

/* debug [enums.gen.go]: Processing enum NSHapticFeedbackPerformanceTime (3 cases) */
// HapticFeedbackPerformanceTime - A time at which to provide haptic feedback to the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/PerformanceTime
type HapticFeedbackPerformanceTime uint

const (
	// HapticFeedbackPerformanceTimeDefault - Allows the system to choose the most appropriate time for feedback to be provided. Currently, this is the next time the screen is updated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/PerformanceTime/default
	HapticFeedbackPerformanceTimeDefault HapticFeedbackPerformanceTime = 0
	// HapticFeedbackPerformanceTimeDrawCompleted - Instructs the system to provide haptic feedback to the user the next time the screen is updated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/PerformanceTime/drawCompleted
	HapticFeedbackPerformanceTimeDrawCompleted HapticFeedbackPerformanceTime = 2
	// HapticFeedbackPerformanceTimeNow - Instructs the system to provide immediate haptic feedback to the user, rather than waiting for synchronization to occur with something visual occurring on screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHapticFeedbackManager/PerformanceTime/now
	HapticFeedbackPerformanceTimeNow HapticFeedbackPerformanceTime = 1
)

/* debug [enums.gen.go]: Processing enum NSHorizontalDirections (3 cases) */
// HorizontalDirections - The absolute directions on the horizontal axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections
type HorizontalDirections uint

const (
	// HorizontalDirectionsAll - All horizontal directions (left and right).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections/NSHorizontalDirectionsAll
	HorizontalDirectionsAll HorizontalDirections = 0
	// HorizontalDirectionsLeft - The left direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections/NSHorizontalDirectionsLeft
	HorizontalDirectionsLeft HorizontalDirections = 1
	// HorizontalDirectionsRight - The right direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHorizontalDirections/NSHorizontalDirectionsRight
	HorizontalDirectionsRight HorizontalDirections = 2
)

/* debug [enums.gen.go]: Processing enum NSImageCacheMode (4 cases) */
// ImageCacheMode - Constants that specify the caching policy on a per-image basis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum
type ImageCacheMode uint

const (
	// ImageCacheAlways - Always generate a cache when drawing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum/always
	ImageCacheAlways ImageCacheMode = 1
	// ImageCacheBySize - Cache if the cache size is smaller than the original data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum/bySize
	ImageCacheBySize ImageCacheMode = 2
	// ImageCacheDefault - Caching is unspecified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum/default
	ImageCacheDefault ImageCacheMode = 0
	// ImageCacheNever - Never cache; always draw direct.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum/never
	ImageCacheNever ImageCacheMode = 3
)

/* debug [enums.gen.go]: Processing enum NSImageDynamicRange (4 cases) */
// ImageDynamicRange - Describes how High Dynamic Range (HDR) image content displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange
type ImageDynamicRange int

const (
	// ImageDynamicRangeConstrainedHigh - Allows for constrained High Dynamic Range (HDR) image content which is useful for mixing HDR and Standard Dynamic Range (SDR) content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/constrainedHigh
	ImageDynamicRangeConstrainedHigh ImageDynamicRange = 1
	// ImageDynamicRangeHigh - Allows image content to use extended dynamic range if it has dynamic range content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/high
	ImageDynamicRangeHigh ImageDynamicRange = 2
	// ImageDynamicRangeStandard - Restricts the image content dynamic range to the standard range regardless of the actual range of the image content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/standard
	ImageDynamicRangeStandard ImageDynamicRange = 0
	// ImageDynamicRangeUnspecified - Indicates that the dynamic range treatment of the image is unknown or otherwise unspecified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/DynamicRange/unspecified
	ImageDynamicRangeUnspecified ImageDynamicRange = -1
)

/* debug [enums.gen.go]: Processing enum NSImageLayoutDirection (3 cases) */
// ImageLayoutDirection - Constants that describe the layout direction for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection
type ImageLayoutDirection int

const (
	// ImageLayoutDirectionLeftToRight - A left-to-right layout direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection/leftToRight
	ImageLayoutDirectionLeftToRight ImageLayoutDirection = 2
	// ImageLayoutDirectionRightToLeft - A right-to-left layout direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection/rightToLeft
	ImageLayoutDirectionRightToLeft ImageLayoutDirection = 3
	// ImageLayoutDirectionUnspecified - An unspecified layout direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection/unspecified
	ImageLayoutDirectionUnspecified ImageLayoutDirection = -1
)

/* debug [enums.gen.go]: Processing enum NSImageLoadStatus (5 cases) */
// ImageLoadStatus - Status values for incremental image loading.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus
type ImageLoadStatus uint

const (
	// ImageLoadStatusCancelled - Image loading was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus/cancelled
	ImageLoadStatusCancelled ImageLoadStatus = 1
	// ImageLoadStatusCompleted - Enough data is available to completely decompress the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus/completed
	ImageLoadStatusCompleted ImageLoadStatus = 0
	// ImageLoadStatusInvalidData - An error occurred during image decompression.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus/invalidData
	ImageLoadStatusInvalidData ImageLoadStatus = 2
	// ImageLoadStatusReadError - Not enough data was available for full decompression of the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus/readError
	ImageLoadStatusReadError ImageLoadStatus = 4
	// ImageLoadStatusUnexpectedEOF - Not enough data was available to fully decompress the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/LoadStatus/unexpectedEOF
	ImageLoadStatusUnexpectedEOF ImageLoadStatus = 3
)

/* debug [enums.gen.go]: Processing enum NSImageResizingMode (2 cases) */
// ImageResizingMode - Constants that describe the resizing mode for the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/ResizingMode-swift.enum
type ImageResizingMode uint

const (
	// ImageResizingModeStretch - The image stretches when it resizes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/ResizingMode-swift.enum/stretch
	ImageResizingModeStretch ImageResizingMode = 1
	// ImageResizingModeTile - The image tiles when it resizes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/ResizingMode-swift.enum/tile
	ImageResizingModeTile ImageResizingMode = 0
)

/* debug [enums.gen.go]: Processing enum NSImageSymbolColorRenderingMode (3 cases) */
// ImageSymbolColorRenderingMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolColorRenderingMode
type ImageSymbolColorRenderingMode uint

const (
	// ImageSymbolColorRenderingModeAutomatic - Automatically uses an appropriate color rendering mode for the symbol’s color layers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolColorRenderingMode/automatic
	ImageSymbolColorRenderingModeAutomatic ImageSymbolColorRenderingMode = 0
	// ImageSymbolColorRenderingModeFlat - Renders the symbol’s color layers using flat colors.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolColorRenderingMode/flat
	ImageSymbolColorRenderingModeFlat ImageSymbolColorRenderingMode = 1
	// ImageSymbolColorRenderingModeGradient - Renders the symbol’s color layers using gradients.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolColorRenderingMode/gradient
	ImageSymbolColorRenderingModeGradient ImageSymbolColorRenderingMode = 2
)

/* debug [enums.gen.go]: Processing enum NSImageSymbolScale (3 cases) */
// ImageSymbolScale - Constants that specify which scale variant of a symbol image to use.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolScale
type ImageSymbolScale uint

const (
	// ImageSymbolScaleLarge - The symbol uses the large scale variant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolScale/large
	ImageSymbolScaleLarge ImageSymbolScale = 3
	// ImageSymbolScaleMedium - The symbol uses the default medium scale variant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolScale/medium
	ImageSymbolScaleMedium ImageSymbolScale = 2
	// ImageSymbolScaleSmall - The symbol uses the small scale variant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolScale/small
	ImageSymbolScaleSmall ImageSymbolScale = 1
)

/* debug [enums.gen.go]: Processing enum NSImageSymbolVariableValueMode (3 cases) */
// ImageSymbolVariableValueMode enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolVariableValueMode
type ImageSymbolVariableValueMode uint

const (
	// ImageSymbolVariableValueModeAutomatic - Automatically selects an appropriate variable value mode for the symbol.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolVariableValueMode/automatic
	ImageSymbolVariableValueModeAutomatic ImageSymbolVariableValueMode = 0
	// ImageSymbolVariableValueModeColor - The “color” variable value mode. Sets the opacity of each variable layer to   either on or off depending on how its threshold compared to the current value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolVariableValueMode/color
	ImageSymbolVariableValueModeColor ImageSymbolVariableValueMode = 1
	// ImageSymbolVariableValueModeDraw - The “draw” variable value mode. Changes the drawn length of each variable layer   to either based on how its range relates to the current value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/SymbolVariableValueMode/draw
	ImageSymbolVariableValueModeDraw ImageSymbolVariableValueMode = 2
)

/* debug [enums.gen.go]: Processing enum NSImageAlignment (9 cases) */
// ImageAlignment - Constants used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment
type ImageAlignment uint

const (
	// ImageAlignBottom - Align the image with the bottom edge of the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignBottom
	ImageAlignBottom ImageAlignment = 5
	// ImageAlignBottomLeft - Align the image with the bottom and left edges of the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignBottomLeft
	ImageAlignBottomLeft ImageAlignment = 6
	// ImageAlignBottomRight - Align the image with the bottom and right edges of the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignBottomRight
	ImageAlignBottomRight ImageAlignment = 7
	// ImageAlignCenter - Center the image in the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignCenter
	ImageAlignCenter ImageAlignment = 0
	// ImageAlignLeft - Align the image with the left edge of the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignLeft
	ImageAlignLeft ImageAlignment = 4
	// ImageAlignRight - Position the image along the right edge of the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignRight
	ImageAlignRight ImageAlignment = 8
	// ImageAlignTop - Position the image along the top edge of the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignTop
	ImageAlignTop ImageAlignment = 1
	// ImageAlignTopLeft - Align the image with the top and left edges of the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignTopLeft
	ImageAlignTopLeft ImageAlignment = 2
	// ImageAlignTopRight - Align the image with the top and right edges of the cell.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageAlignment/alignTopRight
	ImageAlignTopRight ImageAlignment = 3
)

/* debug [enums.gen.go]: Processing enum NSImageInterpolation (5 cases) */
// ImageInterpolation - Constants that specify the interpolation, or image smoothing, behavior used by the image interpolation property.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageInterpolation
type ImageInterpolation uint

const (
	// ImageInterpolationDefault - Use the context’s default interpolation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageInterpolation/default
	ImageInterpolationDefault ImageInterpolation = 0
	// ImageInterpolationHigh - Highest quality, slower than the medium interpolation option.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageInterpolation/high
	ImageInterpolationHigh ImageInterpolation = 3
	// ImageInterpolationLow - Fast, low-quality interpolation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageInterpolation/low
	ImageInterpolationLow ImageInterpolation = 2
	// ImageInterpolationMedium - Medium quality, slower than the low interpolation option.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageInterpolation/medium
	ImageInterpolationMedium ImageInterpolation = 3
	// ImageInterpolationNone - No interpolation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageInterpolation/none
	ImageInterpolationNone ImageInterpolation = 1
)

/* debug [enums.gen.go]: Processing enum NSImageScaling (7 cases) */
// ImageScaling - Constants that specify a cell’s image scaling behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling
type ImageScaling uint

const (
	// ScaleNone - Use  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/NSScaleNone
	ScaleNone ImageScaling = 6
	// ScaleProportionally - Use  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/NSScaleProportionally
	ScaleProportionally ImageScaling = 4
	// ScaleToFit - Use  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/NSScaleToFit
	ScaleToFit ImageScaling = 5
	// ImageScaleAxesIndependently - Scale each dimension to exactly fit destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleAxesIndependently
	ImageScaleAxesIndependently ImageScaling = 1
	// ImageScaleNone - Do not scale the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleNone
	ImageScaleNone ImageScaling = 2
	// ImageScaleProportionallyDown - If it is too large for the destination, scale the image down while preserving the aspect ratio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleProportionallyDown
	ImageScaleProportionallyDown ImageScaling = 0
	// ImageScaleProportionallyUpOrDown - Scale the image to its maximum possible dimensions while both staying within the destination area and preserving its aspect ratio.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleProportionallyUpOrDown
	ImageScaleProportionallyUpOrDown ImageScaling = 3
)

/* debug [enums.gen.go]: Processing enum NSImageFrameStyle (5 cases) */
// ImageFrameStyle - Constants that allow you to specify the kind of frame bordering the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/FrameStyle
type ImageFrameStyle uint

const (
	// ImageFrameButton - A convex bezel that makes the image stand out in relief, like a button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/FrameStyle/button
	ImageFrameButton ImageFrameStyle = 4
	// ImageFrameGrayBezel - A gray, concave bezel that makes the image look sunken.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/FrameStyle/grayBezel
	ImageFrameGrayBezel ImageFrameStyle = 2
	// ImageFrameGroove - A thin groove that looks etched around the image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/FrameStyle/groove
	ImageFrameGroove ImageFrameStyle = 3
	// ImageFrameNone - An invisible frame
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/FrameStyle/none
	ImageFrameNone ImageFrameStyle = 0
	// ImageFramePhoto - A thin black outline and a dropped shadow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageView/FrameStyle/photo
	ImageFramePhoto ImageFrameStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSLayoutAttribute (14 cases) */
// LayoutAttribute - The part of the object’s visual representation that should be used to get the value for the constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute
type LayoutAttribute uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAttribute/NSLayoutAttributeBaseline
	LayoutAttributeBaseline LayoutAttribute = 0
	// LayoutAttributeBottom - The bottom of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/bottom
	LayoutAttributeBottom LayoutAttribute = 4
	// LayoutAttributeCenterX - The center along the x-axis of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/centerX
	LayoutAttributeCenterX LayoutAttribute = 9
	// LayoutAttributeCenterY - The center along the y-axis of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/centerY
	LayoutAttributeCenterY LayoutAttribute = 10
	// LayoutAttributeFirstBaseline - The object’s baseline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/firstBaseline
	LayoutAttributeFirstBaseline LayoutAttribute = 12
	// LayoutAttributeHeight - The height of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/height
	LayoutAttributeHeight LayoutAttribute = 8
	// LayoutAttributeLastBaseline - The object’s baseline.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/lastBaseline
	LayoutAttributeLastBaseline LayoutAttribute = 11
	// LayoutAttributeLeading - The leading edge of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/leading
	LayoutAttributeLeading LayoutAttribute = 5
	// LayoutAttributeLeft - The left side of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/left
	LayoutAttributeLeft LayoutAttribute = 1
	// LayoutAttributeNotAnAttribute - A placeholder value for indicating that the constraint’s second item and second attribute aren’t used in any calculations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/notAnAttribute
	LayoutAttributeNotAnAttribute LayoutAttribute = 0
	// LayoutAttributeRight - The right side of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/right
	LayoutAttributeRight LayoutAttribute = 2
	// LayoutAttributeTop - The top of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/top
	LayoutAttributeTop LayoutAttribute = 3
	// LayoutAttributeTrailing - The trailing edge of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/trailing
	LayoutAttributeTrailing LayoutAttribute = 6
	// LayoutAttributeWidth - The width of the object’s alignment rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Attribute/width
	LayoutAttributeWidth LayoutAttribute = 7
)

/* debug [enums.gen.go]: Processing enum NSLayoutFormatOptions (16 cases) */
// LayoutFormatOptions - A bit mask that specifies both a part of an interface element to align and a direction for the alignment between two interface elements.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions
type LayoutFormatOptions uint

const (
	// LayoutFormatAlignAllBottom - Align all specified interface elements using   on each.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllBottom
	LayoutFormatAlignAllBottom LayoutFormatOptions = 0
	// LayoutFormatAlignAllCenterX - Align all specified interface elements using   on each.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllCenterX
	LayoutFormatAlignAllCenterX LayoutFormatOptions = 0
	// LayoutFormatAlignAllCenterY - Align all specified interface elements using   on each.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllCenterY
	LayoutFormatAlignAllCenterY LayoutFormatOptions = 0
	// LayoutFormatAlignAllFirstBaseline - Align all specified interface elements using the first baseline of each one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllFirstBaseline
	LayoutFormatAlignAllFirstBaseline LayoutFormatOptions = 0
	// LayoutFormatAlignAllLastBaseline - Align all specified interface elements using the last baseline of each one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllLastBaseline
	LayoutFormatAlignAllLastBaseline LayoutFormatOptions = 0
	// LayoutFormatAlignAllLeading - Align all specified interface elements using   on each.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllLeading
	LayoutFormatAlignAllLeading LayoutFormatOptions = 0
	// LayoutFormatAlignAllLeft - Align all specified interface elements using   on each.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllLeft
	LayoutFormatAlignAllLeft LayoutFormatOptions = 0
	// LayoutFormatAlignAllRight - Align all specified interface elements using   on each.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllRight
	LayoutFormatAlignAllRight LayoutFormatOptions = 0
	// LayoutFormatAlignAllTop - Align all specified interface elements using   on each.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllTop
	LayoutFormatAlignAllTop LayoutFormatOptions = 0
	// LayoutFormatAlignAllTrailing - Align all specified interface elements using   on each.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignAllTrailing
	LayoutFormatAlignAllTrailing LayoutFormatOptions = 0
	// LayoutFormatAlignmentMask - Bit mask that can be combined with an   variable to yield only the alignment portion of the format options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/alignmentMask
	LayoutFormatAlignmentMask LayoutFormatOptions = 65535
	// LayoutFormatDirectionLeadingToTrailing - Arrange objects in order based on the normal text flow for the current user interface language. In left-to-right languages (like English), this arrangement results in the first object being placed farthest to the left, the next one to its right, and so on. In right-to-left languages (like Arabic or Hebrew), the ordering is reversed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/directionLeadingToTrailing
	LayoutFormatDirectionLeadingToTrailing LayoutFormatOptions = 0
	// LayoutFormatDirectionLeftToRight - Arrange objects in order from left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/directionLeftToRight
	LayoutFormatDirectionLeftToRight LayoutFormatOptions = 65536
	// LayoutFormatDirectionMask - A bit mask that can be combined with an   variable to yield only the direction portion of the format options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/directionMask
	LayoutFormatDirectionMask LayoutFormatOptions = 3
	// LayoutFormatDirectionRightToLeft - Arrange objects in order from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/FormatOptions/directionRightToLeft
	LayoutFormatDirectionRightToLeft LayoutFormatOptions = 131072
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutFormatOptions/NSLayoutFormatAlignAllBaseline
	LayoutFormatAlignAllBaseline LayoutFormatOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSLayoutConstraintOrientation (2 cases) */
// LayoutConstraintOrientation - The layout constraint orientation, either horizontal or vertical, that the constraint uses to enforce layout between objects.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation
type LayoutConstraintOrientation uint

const (
	// LayoutConstraintOrientationHorizontal - The constraint orientation applied to laying out the horizontal relationship between objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation/horizontal
	LayoutConstraintOrientationHorizontal LayoutConstraintOrientation = 0
	// LayoutConstraintOrientationVertical - The constraint orientation applied to laying out the vertical relationship between objects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Orientation/vertical
	LayoutConstraintOrientationVertical LayoutConstraintOrientation = 1
)

/* debug [enums.gen.go]: Processing enum NSLayoutRelation (3 cases) */
// LayoutRelation - The relation between the first attribute and the modified second attribute in a constraint.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Relation-swift.enum
type LayoutRelation int

const (
	// LayoutRelationEqual - The constraint requires the first attribute to be exactly equal to the modified second attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Relation-swift.enum/equal
	LayoutRelationEqual LayoutRelation = 0
	// LayoutRelationGreaterThanOrEqual - The constraint requires the first attribute to be greater than or equal to the modified second attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Relation-swift.enum/greaterThanOrEqual
	LayoutRelationGreaterThanOrEqual LayoutRelation = 1
	// LayoutRelationLessThanOrEqual - The constraint requires the first attribute to be less than or equal to the modified second attribute.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Relation-swift.enum/lessThanOrEqual
	LayoutRelationLessThanOrEqual LayoutRelation = -1
)

/* debug [enums.gen.go]: Processing enum NSControlCharacterAction (6 cases) */
// ControlCharacterAction - Constants that describe actions for control characters.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction
type ControlCharacterAction uint

const (
	// ControlCharacterActionContainerBreak - An action that triggers a break in layout for the current container.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/containerBreak
	ControlCharacterActionContainerBreak ControlCharacterAction = 32
	// ControlCharacterActionHorizontalTab - An action that inserts a horizontal tab.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/horizontalTab
	ControlCharacterActionHorizontalTab ControlCharacterAction = 4
	// ControlCharacterActionLineBreak - An action that causes a line break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/lineBreak
	ControlCharacterActionLineBreak ControlCharacterAction = 8
	// ControlCharacterActionParagraphBreak - An action that causes a paragraph break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/paragraphBreak
	ControlCharacterActionParagraphBreak ControlCharacterAction = 16
	// ControlCharacterActionWhitespace - An action that adds whitespace.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/whitespace
	ControlCharacterActionWhitespace ControlCharacterAction = 2
	// ControlCharacterActionZeroAdvancement - An action that removes the glyph from layout.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/ControlCharacterAction/zeroAdvancement
	ControlCharacterActionZeroAdvancement ControlCharacterAction = 1
)

/* debug [enums.gen.go]: Processing enum NSGlyphProperty (4 cases) */
// GlyphProperty - Glyph properties.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty
type GlyphProperty uint

const (
	// GlyphPropertyControlCharacter - A glyph representing a control character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty/controlCharacter
	GlyphPropertyControlCharacter GlyphProperty = 2
	// GlyphPropertyElastic - A glyph with a changeable width, such as a white space character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty/elastic
	GlyphPropertyElastic GlyphProperty = 4
	// GlyphPropertyNonBaseCharacter - A glyph that combines several properties.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty/nonBaseCharacter
	GlyphPropertyNonBaseCharacter GlyphProperty = 8
	// GlyphPropertyNull - The null glyph, which the layout manager ignores.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/GlyphProperty/null
	GlyphPropertyNull GlyphProperty = 1
)

/* debug [enums.gen.go]: Processing enum NSTextLayoutOrientation (2 cases) */
// TextLayoutOrientation - Constants that describe the text layout orientation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation
type TextLayoutOrientation uint

const (
	// TextLayoutOrientationHorizontal - Lines render horizontally, each line following the previous from top to bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation/horizontal
	TextLayoutOrientationHorizontal TextLayoutOrientation = 0
	// TextLayoutOrientationVertical - Lines render vertically, each line following the previous from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TextLayoutOrientation/vertical
	TextLayoutOrientationVertical TextLayoutOrientation = 1
)

/* debug [enums.gen.go]: Processing enum NSTypesetterBehavior (6 cases) */
// TypesetterBehavior - Constants that determine the layout manager’s behavior during layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum
type TypesetterBehavior int

const (
	// TypesetterBehavior_10_2 - The typesetter behavior introduced in macOS 10.2.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum/behavior_10_2
	TypesetterBehavior_10_2 TypesetterBehavior = 2
	// TypesetterBehavior_10_2_WithCompatibility - The macOS 10.2 typesetting behavior that is still compatible with the original typesetter behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum/behavior_10_2_WithCompatibility
	TypesetterBehavior_10_2_WithCompatibility TypesetterBehavior = 1
	// TypesetterBehavior_10_3 - The typesetter behavior introduced in macOS 10.3.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum/behavior_10_3
	TypesetterBehavior_10_3 TypesetterBehavior = 3
	// TypesetterBehavior_10_4 - The typesetter behavior introduced in macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum/behavior_10_4
	TypesetterBehavior_10_4 TypesetterBehavior = 4
	// TypesetterLatestBehavior - The current typesetter behavior in the current operating system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum/latestBehavior
	TypesetterLatestBehavior TypesetterBehavior = -1
	// TypesetterOriginalBehavior - The original typesetter behavior, as shipped with macOS 10.1 and earlier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutManager/TypesetterBehavior-swift.enum/originalBehavior
	TypesetterOriginalBehavior TypesetterBehavior = 0
)

/* debug [enums.gen.go]: Processing enum NSLevelIndicatorPlaceholderVisibility (3 cases) */
// LevelIndicatorPlaceholderVisibility enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/PlaceholderVisibility-swift.enum
type LevelIndicatorPlaceholderVisibility uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/PlaceholderVisibility-swift.enum/always
	LevelIndicatorPlaceholderVisibilityAlways LevelIndicatorPlaceholderVisibility = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/PlaceholderVisibility-swift.enum/automatic
	LevelIndicatorPlaceholderVisibilityAutomatic LevelIndicatorPlaceholderVisibility = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/PlaceholderVisibility-swift.enum/whileEditing
	LevelIndicatorPlaceholderVisibilityWhileEditing LevelIndicatorPlaceholderVisibility = 2
)

/* debug [enums.gen.go]: Processing enum NSLevelIndicatorStyle (4 cases) */
// LevelIndicatorStyle - Constants that specify a level indicator’s appearance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/Style
type LevelIndicatorStyle uint

const (
	// LevelIndicatorStyleContinuousCapacity - A style that indicates the capacity of something, such as how much data is on a hard disk.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/Style/continuousCapacity
	LevelIndicatorStyleContinuousCapacity LevelIndicatorStyle = 1
	// LevelIndicatorStyleDiscreteCapacity - A style that displays discrete segments that indicate the capacity of something, such as an audio level.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/Style/discreteCapacity
	LevelIndicatorStyleDiscreteCapacity LevelIndicatorStyle = 2
	// LevelIndicatorStyleRating - A style that indicates a rank, such as a star ranking display.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/Style/rating
	LevelIndicatorStyleRating LevelIndicatorStyle = 3
	// LevelIndicatorStyleRelevancy - A style that indicates the relevancy of an item, such as a search result.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLevelIndicator/Style/relevancy
	LevelIndicatorStyleRelevancy LevelIndicatorStyle = 0
)

/* debug [enums.gen.go]: Processing enum NSLineBreakMode (6 cases) */
// LineBreakMode - Constants that specify what happens when a line is too long for a container.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode
type LineBreakMode uint

const (
	// LineBreakByCharWrapping - The value that indicates wrapping occurs before the first character that doesn’t fit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byCharWrapping
	LineBreakByCharWrapping LineBreakMode = 1
	// LineBreakByClipping - The value that indicates lines don’t extend past the edge of the text container.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byClipping
	LineBreakByClipping LineBreakMode = 2
	// LineBreakByTruncatingHead - The value that indicates that a line displays so that the end fits in the container and an ellipsis glyph indicates the missing text at the beginning of the line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byTruncatingHead
	LineBreakByTruncatingHead LineBreakMode = 3
	// LineBreakByTruncatingMiddle - The value that indicates that a line displays so that the beginning and end fit in the container and an ellipsis glyph indicates the missing text in the middle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byTruncatingMiddle
	LineBreakByTruncatingMiddle LineBreakMode = 5
	// LineBreakByTruncatingTail - The value that indicates a line displays so that the beginning fits in the container and an ellipsis glyph indicates the missing text at the end of the line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byTruncatingTail
	LineBreakByTruncatingTail LineBreakMode = 4
	// LineBreakByWordWrapping - The value that indicates wrapping occurs at word boundaries, unless the word doesn’t fit on a single line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakMode/byWordWrapping
	LineBreakByWordWrapping LineBreakMode = 0
)

/* debug [enums.gen.go]: Processing enum NSLineMovementDirection (5 cases) */
// LineMovementDirection - The direction in which a line moves.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection
type LineMovementDirection uint

const (
	// LineDoesntMove - Line has no movement.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineDoesntMove
	LineDoesntMove LineMovementDirection = 0
	// LineMovesDown - Lines move from top to bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineMovesDown
	LineMovesDown LineMovementDirection = 3
	// LineMovesLeft - Lines move from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineMovesLeft
	LineMovesLeft LineMovementDirection = 1
	// LineMovesRight - Lines move from left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineMovesRight
	LineMovesRight LineMovementDirection = 2
	// LineMovesUp - Lines move from bottom to top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineMovementDirection/NSLineMovesUp
	LineMovesUp LineMovementDirection = 4
)

/* debug [enums.gen.go]: Processing enum NSLineSweepDirection (4 cases) */
// LineSweepDirection - Values that describe the progression of text on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection
type LineSweepDirection uint

const (
	// LineSweepDown - Characters move from top to bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection/NSLineSweepDown
	LineSweepDown LineSweepDirection = 2
	// LineSweepLeft - Characters move from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection/NSLineSweepLeft
	LineSweepLeft LineSweepDirection = 0
	// LineSweepRight - Characters move from left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection/NSLineSweepRight
	LineSweepRight LineSweepDirection = 1
	// LineSweepUp - Characters move from bottom to top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineSweepDirection/NSLineSweepUp
	LineSweepUp LineSweepDirection = 3
)

/* debug [enums.gen.go]: Processing enum NSMatrixMode (4 cases) */
// MatrixMode - These constants determine how 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum
type MatrixMode uint

const (
	// HighlightModeMatrix - An   is highlighted before it’s asked to track the mouse, then unhighlighted when it’s done tracking.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum/highlightModeMatrix
	HighlightModeMatrix MatrixMode = 1
	// ListModeMatrix -  objects are highlighted, but don’t track the mouse.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum/listModeMatrix
	ListModeMatrix MatrixMode = 2
	// RadioModeMatrix - Selects no more than one   at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum/radioModeMatrix
	RadioModeMatrix MatrixMode = 0
	// TrackModeMatrix - The   objects are asked to track the mouse with   whenever the cursor is inside their bounds. No highlighting is performed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMatrix/Mode-swift.enum/trackModeMatrix
	TrackModeMatrix MatrixMode = 3
)

/* debug [enums.gen.go]: Processing enum NSMediaLibrary (3 cases) */
// MediaLibrary - These constants are masks used to configure a Media Library Browser to display specific types of media. Combined masks are not yet supported.  In other words, only one nonzero mask value is supported at a time.  If masks are combined, the lowest mask value is used.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/Library
type MediaLibrary uint

const (
	// MediaLibraryAudio - Display audio media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/Library/audio
	MediaLibraryAudio MediaLibrary = 1
	// MediaLibraryImage - Display image media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/Library/image
	MediaLibraryImage MediaLibrary = 2
	// MediaLibraryMovie - Display movie media.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMediaLibraryBrowserController/Library/movie
	MediaLibraryMovie MediaLibrary = 4
)

/* debug [enums.gen.go]: Processing enum NSMultibyteGlyphPacking (1 cases) */
// MultibyteGlyphPacking - A constant for glyph packing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMultibyteGlyphPacking
type MultibyteGlyphPacking uint

const (
	// NativeShortGlyphPacking - The native format for macOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSMultibyteGlyphPacking/nativeShortGlyphPacking
	NativeShortGlyphPacking MultibyteGlyphPacking = 0
)

/* debug [enums.gen.go]: Processing enum NSOpenGLContextParameter (15 cases) */
// OpenGLContextParameter - Constants that specify context parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter
type OpenGLContextParameter uint

const (
	// OpenGLContextParameterCurrentRendererID - Get the current renderer ID.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/currentRendererID
	OpenGLContextParameterCurrentRendererID OpenGLContextParameter = 5
	// OpenGLContextParameterGPUFragmentProcessing - Get whether the CPU is currently processing fragments with the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/gpuFragmentProcessing
	OpenGLContextParameterGPUFragmentProcessing OpenGLContextParameter = 7
	// OpenGLContextParameterGPUVertexProcessing - Get whether the CPU is currently processing vertices with the GPU.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/gpuVertexProcessing
	OpenGLContextParameterGPUVertexProcessing OpenGLContextParameter = 6
	// OpenGLContextParameterHasDrawable - Returns a Boolean that indicates whether a drawable is attached to the context.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/hasDrawable
	OpenGLContextParameterHasDrawable OpenGLContextParameter = 8
	// OpenGLContextParameterMPSwapsInFlight - The number of frames that the multithreaded OpenGL engine can process before stalling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/mpSwapsInFlight
	OpenGLContextParameterMPSwapsInFlight OpenGLContextParameter = 9
	// OpenGLContextParameterRasterizationEnable - If disabled, all rasterization of 2D and 3D primitives is disabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/rasterizationEnable
	OpenGLContextParameterRasterizationEnable OpenGLContextParameter = 12
	// OpenGLContextParameterReclaimResources - Enable or disable reclaiming resources.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/reclaimResources
	OpenGLContextParameterReclaimResources OpenGLContextParameter = 4
	// OpenGLContextParameterStateValidation - If enabled, OpenGL inspects the context state each time the update method is called to ensure that it is in an appropriate state for switching between renderers.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/stateValidation
	OpenGLContextParameterStateValidation OpenGLContextParameter = 13
	// OpenGLContextParameterSurfaceBackingSize - Set or get the height and width of the back buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/surfaceBackingSize
	OpenGLContextParameterSurfaceBackingSize OpenGLContextParameter = 3
	// OpenGLContextParameterSurfaceOpacity - Set or get the surface opacity.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/surfaceOpacity
	OpenGLContextParameterSurfaceOpacity OpenGLContextParameter = 2
	// OpenGLContextParameterSurfaceOrder - Set or get the surface order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/surfaceOrder
	OpenGLContextParameterSurfaceOrder OpenGLContextParameter = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/surfaceSurfaceVolatile
	OpenGLContextParameterSurfaceSurfaceVolatile OpenGLContextParameter = 14
	// OpenGLContextParameterSwapInterval - Set or get the swap interval.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/swapInterval
	OpenGLContextParameterSwapInterval OpenGLContextParameter = 0
	// OpenGLContextParameterSwapRectangle - Sets or gets the swap rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/swapRectangle
	OpenGLContextParameterSwapRectangle OpenGLContextParameter = 10
	// OpenGLContextParameterSwapRectangleEnable - Enables or disables the swap rectangle in the context’s drawable object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLContext/Parameter/swapRectangleEnable
	OpenGLContextParameterSwapRectangleEnable OpenGLContextParameter = 11
)

/* debug [enums.gen.go]: Processing enum NSOpenGLGlobalOption (5 cases) */
// OpenGLGlobalOption - Constants that specify OpenGL options.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption
type OpenGLGlobalOption uint

const (
	// OpenGLGOResetLibrary - Does a soft reset of the CGL library if true.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption/NSOpenGLGOResetLibrary
	OpenGLGOResetLibrary OpenGLGlobalOption = 4
	// OpenGLGOClearFormatCache - Resets the pixel format cache if true.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption/clearFormatCache
	OpenGLGOClearFormatCache OpenGLGlobalOption = 1
	// OpenGLGOFormatCacheSize - Sets the size of the pixel format cache.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption/formatCacheSize
	OpenGLGOFormatCacheSize OpenGLGlobalOption = 0
	// OpenGLGORetainRenderers - Whether to retain loaded renderers in memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption/retainRenderers
	OpenGLGORetainRenderers OpenGLGlobalOption = 2
	// OpenGLGOUseBuildCache - Whether to enable the function compilation block cache. This is off by default. It must be enabled at startup.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLGlobalOption/useBuildCache
	OpenGLGOUseBuildCache OpenGLGlobalOption = 3
)

/* debug [enums.gen.go]: Processing enum NSPageControllerTransitionStyle (3 cases) */
// PageControllerTransitionStyle - These constants control the transition style of the page controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/TransitionStyle-swift.enum
type PageControllerTransitionStyle uint

const (
	// PageControllerTransitionStyleHorizontalStrip - Each page is laid out next to each other in one long horizontal strip
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/TransitionStyle-swift.enum/horizontalStrip
	PageControllerTransitionStyleHorizontalStrip PageControllerTransitionStyle = 2
	// PageControllerTransitionStyleStackBook - Pages are stacked on top of each other. Pages animate out to the left to reveal the next page. Previous pages animate in from the left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/TransitionStyle-swift.enum/stackBook
	PageControllerTransitionStyleStackBook PageControllerTransitionStyle = 1
	// PageControllerTransitionStyleStackHistory - Pages are stacked on top of each other. Pages animate out to the right to reveal the previous page. Next pages animate in from the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageController/TransitionStyle-swift.enum/stackHistory
	PageControllerTransitionStyleStackHistory PageControllerTransitionStyle = 0
)

/* debug [enums.gen.go]: Processing enum NSPageLayoutResult (2 cases) */
// PageLayoutResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/Result
type PageLayoutResult uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/Result/cancelled
	PageLayoutResultCancelled PageLayoutResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPageLayout/Result/changed
	PageLayoutResultChanged PageLayoutResult = 1
)

/* debug [enums.gen.go]: Processing enum NSLineBreakStrategy (4 cases) */
// LineBreakStrategy - Constants that specify how the text system breaks lines while laying out paragraphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct
type LineBreakStrategy uint

const (
	// LineBreakStrategyNone - The text system doesn’t use any line-break strategies.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLineBreakStrategy/NSLineBreakStrategyNone
	LineBreakStrategyNone LineBreakStrategy = 0
	// LineBreakStrategyHangulWordPriority - The text system prohibits breaking between Hangul characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct/hangulWordPriority
	LineBreakStrategyHangulWordPriority LineBreakStrategy = 2
	// LineBreakStrategyPushOut - The text system pushes out individual lines to avoid an orphan word on the last line of the paragraph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct/pushOut
	LineBreakStrategyPushOut LineBreakStrategy = 1
	// LineBreakStrategyStandard - The text system uses the same configuration of line-break strategies that it uses for standard UI labels.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/LineBreakStrategy-swift.struct/standard
	LineBreakStrategyStandard LineBreakStrategy = 3
)

/* debug [enums.gen.go]: Processing enum NSTextTabType (4 cases) */
// TextTabType - Constants that specify the type of tab stop.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType
type TextTabType uint

const (
	// CenterTabStopType - A center-aligned tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType/centerTabStopType
	CenterTabStopType TextTabType = 2
	// DecimalTabStopType - A tab stop that aligns columns of numbers to each number’s decimal point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType/decimalTabStopType
	DecimalTabStopType TextTabType = 3
	// LeftTabStopType - A left-aligned tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType/leftTabStopType
	LeftTabStopType TextTabType = 0
	// RightTabStopType - A right-aligned tab stop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSParagraphStyle/TextTabType/rightTabStopType
	RightTabStopType TextTabType = 1
)

/* debug [enums.gen.go]: Processing enum NSPasteboardAccessBehavior (4 cases) */
// PasteboardAccessBehavior - A value indicating pasteboard access behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum
type PasteboardAccessBehavior uint

const (
	// PasteboardAccessBehaviorAlwaysAllow - The system will automatically allow all pasteboard access, without notifying the user.  The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/alwaysAllow
	PasteboardAccessBehaviorAlwaysAllow PasteboardAccessBehavior = 2
	// PasteboardAccessBehaviorAlwaysDeny - The system will automatically deny all pasteboard access, without notifying the user. However, access that is both user originated and paste related will always be allowed, and will not result in a notification. The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/alwaysDeny
	PasteboardAccessBehaviorAlwaysDeny PasteboardAccessBehavior = 3
	// PasteboardAccessBehaviorAsk - The system will notify the user and ask for permission before granting pasteboard access. However, access that is both user originated and paste related will always be allowed, and will not result in a notification. The app is listed in the corresponding System Settings pane.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/ask
	PasteboardAccessBehaviorAsk PasteboardAccessBehavior = 1
	// PasteboardAccessBehaviorDefault - The default behavior for the General pasteboard is to ask upon programmatic access. All other pasteboards default to always allow access.   If an app has never triggered a pasteboard access alert, its General pasteboard will report   behavior. Such an app is not shown in the corresponding System Settings pane.   Once programmatic pasteboard access triggers the first pasteboard access alert, the state automatically changes to  . At this point the app starts being shown in System Settings, where the user can toggle the behavior between  ,  , and  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/AccessBehavior-swift.enum/default
	PasteboardAccessBehaviorDefault PasteboardAccessBehavior = 0
)

/* debug [enums.gen.go]: Processing enum NSPasteboardContentsOptions (1 cases) */
// PasteboardContentsOptions - Options for preparing the pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ContentsOptions
type PasteboardContentsOptions uint

const (
	// PasteboardContentsCurrentHostOnly - The pasteboard contents are available only on the current device, and not on any other devices.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ContentsOptions/currentHostOnly
	PasteboardContentsCurrentHostOnly PasteboardContentsOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSPasteboardReadingOptions (4 cases) */
// PasteboardReadingOptions - Options that specify how to interpret data on the pasteboard when initializing pasteboard data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
type PasteboardReadingOptions uint

const (
	// PasteboardReadingAsData - An option to read data from the pasteboard as-is and return it as a data object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asData
	PasteboardReadingAsData PasteboardReadingOptions = 0
	// PasteboardReadingAsKeyedArchive - An option to read data from the pasteboard and use it to initialize the object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asKeyedArchive
	PasteboardReadingAsKeyedArchive PasteboardReadingOptions = 4
	// PasteboardReadingAsPropertyList - An option to read data from the pasteboard and unserialize it as a property list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asPropertyList
	PasteboardReadingAsPropertyList PasteboardReadingOptions = 2
	// PasteboardReadingAsString - An option to read data from the pasteboard and convert it to a string object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions/asString
	PasteboardReadingAsString PasteboardReadingOptions = 1
)

/* debug [enums.gen.go]: Processing enum NSPasteboardWritingOptions (1 cases) */
// PasteboardWritingOptions - Type to specify options for writing to a pasteboard.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions
type PasteboardWritingOptions uint

const (
	// PasteboardWritingPromised - Data for a type with this option is promised, not immediately written.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/WritingOptions/promised
	PasteboardWritingPromised PasteboardWritingOptions = 512
)

/* debug [enums.gen.go]: Processing enum NSPathStyle (3 cases) */
// PathStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/Style
type PathStyle uint

const (
	// PathStylePopUp - The pop-up display style and behavior. Only the last path component is displayed with an icon image and component name. The full path is shown when the user clicks on the cell. If the cell is editable, a Choose item is included to enable selecting a different path.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/Style/popUp
	PathStylePopUp PathStyle = 2
	// PathStyleStandard - The standard display style and behavior. All path component cells are displayed with an icon image and component name. If the path can not fully be displayed, the middle parts are truncated as required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathControl/Style/standard
	PathStyleStandard PathStyle = 0
	// PathStyleNavigationBar - The navigation bar display style and behavior. Similar to the   with the navigation bar drawing style. Also known as the breadcrumb style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPathStyle/NSPathStyleNavigationBar
	PathStyleNavigationBar PathStyle = 3
)

/* debug [enums.gen.go]: Processing enum NSPickerTouchBarItemControlRepresentation (3 cases) */
// PickerTouchBarItemControlRepresentation - Constants that specify display styles for picker bar items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum
type PickerTouchBarItemControlRepresentation uint

const (
	// PickerTouchBarItemControlRepresentationAutomatic - The system dynamically changes the display mode based on the available space.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum/automatic
	PickerTouchBarItemControlRepresentationAutomatic PickerTouchBarItemControlRepresentation = 0
	// PickerTouchBarItemControlRepresentationCollapsed - The system displays the control’s options through a popover.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum/collapsed
	PickerTouchBarItemControlRepresentationCollapsed PickerTouchBarItemControlRepresentation = 2
	// PickerTouchBarItemControlRepresentationExpanded - The system displays the control and all of its options directly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/ControlRepresentation-swift.enum/expanded
	PickerTouchBarItemControlRepresentationExpanded PickerTouchBarItemControlRepresentation = 1
)

/* debug [enums.gen.go]: Processing enum NSPickerTouchBarItemSelectionMode (3 cases) */
// PickerTouchBarItemSelectionMode - Constants that specify selection modes for picker bar items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum
type PickerTouchBarItemSelectionMode uint

const (
	// PickerTouchBarItemSelectionModeMomentary - A mode in which an option is only selected while a person is interacting within the bounds of that option.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum/momentary
	PickerTouchBarItemSelectionModeMomentary PickerTouchBarItemSelectionMode = 2
	// PickerTouchBarItemSelectionModeSelectAny - A mode in which a person can select one or more options in the control at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum/selectAny
	PickerTouchBarItemSelectionModeSelectAny PickerTouchBarItemSelectionMode = 1
	// PickerTouchBarItemSelectionModeSelectOne - A mode in which a person can only select one option in the control at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPickerTouchBarItem/SelectionMode-swift.enum/selectOne
	PickerTouchBarItemSelectionModeSelectOne PickerTouchBarItemSelectionMode = 0
)

/* debug [enums.gen.go]: Processing enum NSPopUpArrowPosition (3 cases) */
// PopUpArrowPosition - These constants are defined for use with the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/ArrowPosition
type PopUpArrowPosition uint

const (
	// PopUpArrowAtBottom - Arrow is drawn at the edge of the button, pointing toward the  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/ArrowPosition/arrowAtBottom
	PopUpArrowAtBottom PopUpArrowPosition = 2
	// PopUpArrowAtCenter - Arrow is centered vertically, pointing toward the  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/ArrowPosition/arrowAtCenter
	PopUpArrowAtCenter PopUpArrowPosition = 1
	// PopUpNoArrow - Does not display any arrow in the control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/ArrowPosition/noArrow
	PopUpNoArrow PopUpArrowPosition = 0
)

/* debug [enums.gen.go]: Processing enum NSPrintingOrientation (2 cases) */
// PrintingOrientation - Constants that specify page orientations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/Orientation-swift.enum
type PrintingOrientation uint

const (
	// LandscapeOrientation - Orientation is landscape (page is wider than it is tall).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/Orientation-swift.enum/landscapeOrientation
	LandscapeOrientation PrintingOrientation = 1
	// PortraitOrientation - Orientation is portrait (page is taller than it is wide).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/Orientation-swift.enum/portraitOrientation
	PortraitOrientation PrintingOrientation = 0
)

/* debug [enums.gen.go]: Processing enum NSPrintingPaginationMode (3 cases) */
// PrintingPaginationMode - Constants that specify the different ways in which an image is divided into pages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaginationMode
type PrintingPaginationMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaginationMode/automatic
	PrintingPaginationModeAutomatic PrintingPaginationMode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaginationMode/clip
	PrintingPaginationModeClip PrintingPaginationMode = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaginationMode/fit
	PrintingPaginationModeFit PrintingPaginationMode = 1
)

/* debug [enums.gen.go]: Processing enum NSPaperOrientation (2 cases) */
// PaperOrientation - Constants that describe the orientation of printing on a page.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaperOrientation
type PaperOrientation uint

const (
	// PaperOrientationLandscape - Pages are printed in landscape orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaperOrientation/landscape
	PaperOrientationLandscape PaperOrientation = 1
	// PaperOrientationPortrait - Pages are printed in portrait orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/PaperOrientation/portrait
	PaperOrientationPortrait PaperOrientation = 0
)

/* debug [enums.gen.go]: Processing enum NSPrintingPageOrder (4 cases) */
// PrintingPageOrder - Constants that specify the page order.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/PageOrder-swift.enum
type PrintingPageOrder int

const (
	// AscendingPageOrder - Ascending (back to front) page order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/PageOrder-swift.enum/ascendingPageOrder
	AscendingPageOrder PrintingPageOrder = 1
	// DescendingPageOrder - Descending (front to back) page order.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/PageOrder-swift.enum/descendingPageOrder
	DescendingPageOrder PrintingPageOrder = -1
	// SpecialPageOrder - The spooler does not rearrange pages—they are printed in the order received by the spooler.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/PageOrder-swift.enum/specialPageOrder
	SpecialPageOrder PrintingPageOrder = 0
	// UnknownPageOrder - No page order specified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/PageOrder-swift.enum/unknownPageOrder
	UnknownPageOrder PrintingPageOrder = 2
)

/* debug [enums.gen.go]: Processing enum NSPrintRenderingQuality (2 cases) */
// PrintRenderingQuality - Constants that specify the print quality in use.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/RenderingQuality
type PrintRenderingQuality uint

const (
	// PrintRenderingQualityBest - Renders the printing at the best possible quality, regardless of speed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/RenderingQuality/best
	PrintRenderingQualityBest PrintRenderingQuality = 0
	// PrintRenderingQualityResponsive - Sacrifices the least possible amount of rendering quality for speed to maintain a responsive user interface. This option should be used only after establishing that best quality rendering does indeed make the user interface unresponsive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintOperation/RenderingQuality/responsive
	PrintRenderingQualityResponsive PrintRenderingQuality = 1
)

/* debug [enums.gen.go]: Processing enum NSPrintPanelOptions (8 cases) */
// PrintPanelOptions - Constants that specify options for configuring the contents of the main Print panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct
type PrintPanelOptions uint

const (
	// PrintPanelShowsCopies - The Print panel includes a field for manipulating the number of copies being printed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct/showsCopies
	PrintPanelShowsCopies PrintPanelOptions = 1
	// PrintPanelShowsOrientation - The Print panel includes a control for manipulating the page orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct/showsOrientation
	PrintPanelShowsOrientation PrintPanelOptions = 8
	// PrintPanelShowsPageRange - The Print panel includes a set of fields for manipulating the range of pages being printed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct/showsPageRange
	PrintPanelShowsPageRange PrintPanelOptions = 2
	// PrintPanelShowsPageSetupAccessory - The Print panel includes a separate accessory view for manipulating the paper size, orientation, and scaling attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct/showsPageSetupAccessory
	PrintPanelShowsPageSetupAccessory PrintPanelOptions = 256
	// PrintPanelShowsPaperSize - The Print panel includes a control for manipulating the paper size of the printer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct/showsPaperSize
	PrintPanelShowsPaperSize PrintPanelOptions = 4
	// PrintPanelShowsPreview - The Print panel displays a built-in preview of the document contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct/showsPreview
	PrintPanelShowsPreview PrintPanelOptions = 131072
	// PrintPanelShowsPrintSelection - The Print panel includes an additional selection option for paper range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct/showsPrintSelection
	PrintPanelShowsPrintSelection PrintPanelOptions = 17
	// PrintPanelShowsScaling - The Print panel includes a control for scaling the printed output.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Options-swift.struct/showsScaling
	PrintPanelShowsScaling PrintPanelOptions = 16
)

/* debug [enums.gen.go]: Processing enum NSPrintPanelResult (2 cases) */
// PrintPanelResult enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Result
type PrintPanelResult uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Result/cancelled
	PrintPanelResultCancelled PrintPanelResult = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/Result/printed
	PrintPanelResultPrinted PrintPanelResult = 1
)

/* debug [enums.gen.go]: Processing enum NSPrinterTableStatus (3 cases) */
// PrinterTableStatus - Constants that describe the state of a printer information table stored by a printer object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/TableStatus
type PrinterTableStatus uint

const (
	// PrinterTableError - Printer table is not valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/TableStatus/error
	PrinterTableError PrinterTableStatus = 2
	// PrinterTableNotFound - Printer table was not found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/TableStatus/notFound
	PrinterTableNotFound PrinterTableStatus = 1
	// PrinterTableOK - Printer table was found and is valid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/TableStatus/ok
	PrinterTableOK PrinterTableStatus = 0
)

/* debug [enums.gen.go]: Processing enum NSProgressIndicatorStyle (2 cases) */
// ProgressIndicatorStyle - Constants that specify the progress indicator’s style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/Style-swift.enum
type ProgressIndicatorStyle uint

const (
	// ProgressIndicatorStyleBar - A rectangular indicator that can be determinate or indeterminate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/Style-swift.enum/bar
	ProgressIndicatorStyleBar ProgressIndicatorStyle = 0
	// ProgressIndicatorStyleSpinning - A small circular indicator that can be determinate or indeterminate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/Style-swift.enum/spinning
	ProgressIndicatorStyleSpinning ProgressIndicatorStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSProgressIndicatorThickness (4 cases) */
// ProgressIndicatorThickness - Specify the height of a progress indicator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicatorThickness
type ProgressIndicatorThickness uint

const (
	// ProgressIndicatorPreferredAquaThickness - 12
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicatorThickness/NSProgressIndicatorPreferredAquaThickness
	ProgressIndicatorPreferredAquaThickness ProgressIndicatorThickness = 3
	// ProgressIndicatorPreferredLargeThickness - 18
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicatorThickness/NSProgressIndicatorPreferredLargeThickness
	ProgressIndicatorPreferredLargeThickness ProgressIndicatorThickness = 2
	// ProgressIndicatorPreferredSmallThickness - 10
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicatorThickness/NSProgressIndicatorPreferredSmallThickness
	ProgressIndicatorPreferredSmallThickness ProgressIndicatorThickness = 1
	// ProgressIndicatorPreferredThickness - 14
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSProgressIndicatorThickness/NSProgressIndicatorPreferredThickness
	ProgressIndicatorPreferredThickness ProgressIndicatorThickness = 0
)

/* debug [enums.gen.go]: Processing enum NSRectAlignment (9 cases) */
// RectAlignment - Constants that specify alignment to an edge or a set of edges depending on the user interface layout direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment
type RectAlignment uint

const (
	// RectAlignmentBottom - Aligns to the bottom edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment/bottom
	RectAlignmentBottom RectAlignment = 5
	// RectAlignmentBottomLeading - Aligns to the bottom and leading edges.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment/bottomLeading
	RectAlignmentBottomLeading RectAlignment = 4
	// RectAlignmentBottomTrailing - Aligns to the bottom and trailing edges.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment/bottomTrailing
	RectAlignmentBottomTrailing RectAlignment = 6
	// RectAlignmentLeading - Aligns to the leading edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment/leading
	RectAlignmentLeading RectAlignment = 3
	// RectAlignmentNone - Has no specified alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment/none
	RectAlignmentNone RectAlignment = 0
	// RectAlignmentTop - Aligns to the top edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment/top
	RectAlignmentTop RectAlignment = 1
	// RectAlignmentTopLeading - Aligns to the top and leading edges.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment/topLeading
	RectAlignmentTopLeading RectAlignment = 2
	// RectAlignmentTopTrailing - Aligns to the top and trailing edges.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment/topTrailing
	RectAlignmentTopTrailing RectAlignment = 8
	// RectAlignmentTrailing - Aligns to the trailing edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRectAlignment/trailing
	RectAlignmentTrailing RectAlignment = 7
)

/* debug [enums.gen.go]: Processing enum NSRuleEditorNestingMode (4 cases) */
// RuleEditorNestingMode - Specifies a type for nesting modes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/NestingMode-swift.enum
type RuleEditorNestingMode uint

const (
	// RuleEditorNestingModeCompound - Unlimited nesting and compound rows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/NestingMode-swift.enum/compound
	RuleEditorNestingModeCompound RuleEditorNestingMode = 2
	// RuleEditorNestingModeList - Allows a single list, with no nesting and no compound rows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/NestingMode-swift.enum/list
	RuleEditorNestingModeList RuleEditorNestingMode = 1
	// RuleEditorNestingModeSimple - One compound row at the top with subrows beneath it, and no further nesting allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/NestingMode-swift.enum/simple
	RuleEditorNestingModeSimple RuleEditorNestingMode = 3
	// RuleEditorNestingModeSingle - Only a single row is allowed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/NestingMode-swift.enum/single
	RuleEditorNestingModeSingle RuleEditorNestingMode = 0
)

/* debug [enums.gen.go]: Processing enum NSRuleEditorRowType (2 cases) */
// RuleEditorRowType - Specifies a type for row types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/RowType
type RuleEditorRowType uint

const (
	// RuleEditorRowTypeCompound - Specifies a compound row.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/RowType/compound
	RuleEditorRowTypeCompound RuleEditorRowType = 1
	// RuleEditorRowTypeSimple - Specifies a simple row.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRuleEditor/RowType/simple
	RuleEditorRowTypeSimple RuleEditorRowType = 0
)

/* debug [enums.gen.go]: Processing enum NSRulerOrientation (2 cases) */
// RulerOrientation - These constants are defined to specify a ruler’s orientation and are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/Orientation-swift.enum
type RulerOrientation uint

const (
	// HorizontalRuler - Ruler is oriented horizontally.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/Orientation-swift.enum/horizontalRuler
	HorizontalRuler RulerOrientation = 0
	// VerticalRuler - Ruler is oriented vertically.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/Orientation-swift.enum/verticalRuler
	VerticalRuler RulerOrientation = 1
)

/* debug [enums.gen.go]: Processing enum NSScrollElasticity (3 cases) */
// ScrollElasticity - These constants determine the elasticity behavior for an axis of the scrollview.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity
type ScrollElasticity uint

const (
	// ScrollElasticityAllowed - Allow content to be scrolled past its bounds on this axis in an elastic fashion.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/allowed
	ScrollElasticityAllowed ScrollElasticity = 2
	// ScrollElasticityAutomatic - Automatically determine whether to allow elasticity on this axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/automatic
	ScrollElasticityAutomatic ScrollElasticity = 0
	// ScrollElasticityNone - Disallow scrolling beyond document bounds on this axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/Elasticity/none
	ScrollElasticityNone ScrollElasticity = 1
)

/* debug [enums.gen.go]: Processing enum NSScrollViewFindBarPosition (3 cases) */
// ScrollViewFindBarPosition - These constants define the position of the find bar in relation to the scroll view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum
type ScrollViewFindBarPosition uint

const (
	// ScrollViewFindBarPositionAboveContent - The find bar is displayed above the scroll view content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum/aboveContent
	ScrollViewFindBarPositionAboveContent ScrollViewFindBarPosition = 1
	// ScrollViewFindBarPositionAboveHorizontalRuler - The find bar is displayed above the horizontal ruler, if visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum/aboveHorizontalRuler
	ScrollViewFindBarPositionAboveHorizontalRuler ScrollViewFindBarPosition = 0
	// ScrollViewFindBarPositionBelowContent - The find bar is displayed below the scroll view content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollView/FindBarPosition-swift.enum/belowContent
	ScrollViewFindBarPositionBelowContent ScrollViewFindBarPosition = 2
)

/* debug [enums.gen.go]: Processing enum NSScrollerArrow (2 cases) */
// ScrollerArrow - These constants describe the two scroller buttons and are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Arrow
type ScrollerArrow uint

const (
	// ScrollerDecrementArrow - The up or left scroll button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Arrow/decrementArrow
	ScrollerDecrementArrow ScrollerArrow = 1
	// ScrollerIncrementArrow - The down or right scroll button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Arrow/incrementArrow
	ScrollerIncrementArrow ScrollerArrow = 0
)

/* debug [enums.gen.go]: Processing enum NSScrollArrowPosition (4 cases) */
// ScrollArrowPosition - These constants specify where the scroller’s buttons appear and are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/ArrowPosition
type ScrollArrowPosition uint

const (
	// ScrollerArrowsDefaultSetting - Buttons are displayed according to the system-wide appearance preferences.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/ArrowPosition/scrollerArrowsDefaultSetting
	ScrollerArrowsDefaultSetting ScrollArrowPosition = 0
	// ScrollerArrowsMaxEnd - Buttons at bottom or right. This constant has been deprecated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/ArrowPosition/scrollerArrowsMaxEnd
	ScrollerArrowsMaxEnd ScrollArrowPosition = 0
	// ScrollerArrowsMinEnd - Buttons at top or left. This has been deprecated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/ArrowPosition/scrollerArrowsMinEnd
	ScrollerArrowsMinEnd ScrollArrowPosition = 1
	// ScrollerArrowsNone - No buttons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/ArrowPosition/scrollerArrowsNone
	ScrollerArrowsNone ScrollArrowPosition = 2
)

/* debug [enums.gen.go]: Processing enum NSScrollerKnobStyle (3 cases) */
// ScrollerKnobStyle - Specify different knob styles.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/KnobStyle-swift.enum
type ScrollerKnobStyle uint

const (
	// ScrollerKnobStyleDark - Specifies a dark knob.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/KnobStyle-swift.enum/dark
	ScrollerKnobStyleDark ScrollerKnobStyle = 1
	// ScrollerKnobStyleDefault - Specifies a dark knob with a light border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/KnobStyle-swift.enum/default
	ScrollerKnobStyleDefault ScrollerKnobStyle = 0
	// ScrollerKnobStyleLight - Specifies a light knob.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/KnobStyle-swift.enum/light
	ScrollerKnobStyleLight ScrollerKnobStyle = 2
)

/* debug [enums.gen.go]: Processing enum NSScrollerPart (7 cases) */
// ScrollerPart - These constants specify the different parts of the scroller:
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Part
type ScrollerPart uint

const (
	// ScrollerDecrementLine - Up or left by a small amount.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/decrementLine
	ScrollerDecrementLine ScrollerPart = 4
	// ScrollerDecrementPage - Up or left by a large amount.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/decrementPage
	ScrollerDecrementPage ScrollerPart = 1
	// ScrollerIncrementLine - Down or right by a small amount.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/incrementLine
	ScrollerIncrementLine ScrollerPart = 5
	// ScrollerIncrementPage - Down or right by a large amount.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/incrementPage
	ScrollerIncrementPage ScrollerPart = 3
	// ScrollerKnob - Directly to the scroller’s value, as given by  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/knob
	ScrollerKnob ScrollerPart = 2
	// ScrollerKnobSlot - Directly to the scroller’s value, as given by  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/knobSlot
	ScrollerKnobSlot ScrollerPart = 6
	// ScrollerNoPart - Don’t scroll at all.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Part/noPart
	ScrollerNoPart ScrollerPart = 0
)

/* debug [enums.gen.go]: Processing enum NSScrollerStyle (2 cases) */
// ScrollerStyle - Constants to specify the scroller style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Style
type ScrollerStyle uint

const (
	// ScrollerStyleLegacy - Specifies legacy-style scrollers as prior to macOS 10.7.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Style/legacy
	ScrollerStyleLegacy ScrollerStyle = 0
	// ScrollerStyleOverlay - Specifies overlay-style scrollers in macOS 10.7 and later.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/Style/overlay
	ScrollerStyleOverlay ScrollerStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSUsableScrollerParts (3 cases) */
// UsableScrollerParts - These constants specify which parts of the scroller are visible.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/UsableParts-swift.enum
type UsableScrollerParts uint

const (
	// AllScrollerParts - Specifies that the scroller has at least a knob, possibly also scroll buttons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/UsableParts-swift.enum/allScrollerParts
	AllScrollerParts UsableScrollerParts = 2
	// NoScrollerParts - Specifies that the scroller has neither a knob nor scroll buttons, only the knob slot.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/UsableParts-swift.enum/noScrollerParts
	NoScrollerParts UsableScrollerParts = 0
	// OnlyScrollerArrows - Specifies that the scroller has only scroll buttons, no knob.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScroller/UsableParts-swift.enum/onlyScrollerArrows
	OnlyScrollerArrows UsableScrollerParts = 1
)

/* debug [enums.gen.go]: Processing enum NSScrubberAlignment (4 cases) */
// ScrubberAlignment - The specified preferred alignment of items within the scrubber, when they come to rest following a user’s scrolling or paging interaction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment
type ScrubberAlignment uint

const (
	// ScrubberAlignmentCenter - Center alignment of items within the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/center
	ScrubberAlignmentCenter ScrubberAlignment = 3
	// ScrubberAlignmentLeading - Leading alignment of items within the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/leading
	ScrubberAlignmentLeading ScrubberAlignment = 1
	// ScrubberAlignmentNone - No preference for item alignment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/none
	ScrubberAlignmentNone ScrubberAlignment = 0
	// ScrubberAlignmentTrailing - Trailing alignment of items within the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Alignment/trailing
	ScrubberAlignmentTrailing ScrubberAlignment = 2
)

/* debug [enums.gen.go]: Processing enum NSScrubberMode (2 cases) */
// ScrubberMode - The scrolling behavior for a scrubber.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum
type ScrubberMode uint

const (
	// ScrubberModeFixed - A scrolling mode in which scrubber items remain fixed in place, and the item under the user’s finger is highlighted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum/fixed
	ScrubberModeFixed ScrubberMode = 0
	// ScrubberModeFree - A scrolling mode in which the scrubber scrolls as the user swipes horizontally across the scrubber.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrubber/Mode-swift.enum/free
	ScrubberModeFree ScrubberMode = 1
)

/* debug [enums.gen.go]: Processing enum NSSegmentDistribution (4 cases) */
// SegmentDistribution enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Distribution
type SegmentDistribution uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Distribution/fill
	SegmentDistributionFill SegmentDistribution = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Distribution/fillEqually
	SegmentDistributionFillEqually SegmentDistribution = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Distribution/fillProportionally
	SegmentDistributionFillProportionally SegmentDistribution = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Distribution/fit
	SegmentDistributionFit SegmentDistribution = 0
)

/* debug [enums.gen.go]: Processing enum NSSegmentStyle (8 cases) */
// SegmentStyle - The following constants specify the visual style used to display the segmented control. They are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style
type SegmentStyle uint

const (
	// SegmentStyleAutomatic - The appearance of the segmented control is automatically determined based on the type of window in which the control is displayed and the position within the window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style/automatic
	SegmentStyleAutomatic SegmentStyle = 0
	// SegmentStyleCapsule - The control is displayed using the capsule style. In macOS 10.7 and later, this style uses the artwork defined for  , so you should specify   instead.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style/capsule
	SegmentStyleCapsule SegmentStyle = 9
	// SegmentStyleRoundRect - The control is displayed using the round rect style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style/roundRect
	SegmentStyleRoundRect SegmentStyle = 3
	// SegmentStyleRounded - The control is displayed using the rounded style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style/rounded
	SegmentStyleRounded SegmentStyle = 1
	// SegmentStyleSeparated - The segments in the control are displayed very close to each other but not touching. For example, Safari in macOS 10.10 and later uses this style for the previous and next page segmented control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style/separated
	SegmentStyleSeparated SegmentStyle = 7
	// SegmentStyleSmallSquare - The control is displayed using the small square style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style/smallSquare
	SegmentStyleSmallSquare SegmentStyle = 6
	// SegmentStyleTexturedRounded - The control is displayed using the textured rounded style. In macOS 10.7 and later, this style uses the artwork defined for  , so you should specify   instead.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style/texturedRounded
	SegmentStyleTexturedRounded SegmentStyle = 8
	// SegmentStyleTexturedSquare - The control is displayed using the textured square style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/Style/texturedSquare
	SegmentStyleTexturedSquare SegmentStyle = 4
)

/* debug [enums.gen.go]: Processing enum NSSegmentSwitchTracking (4 cases) */
// SegmentSwitchTracking - The following constants specify the type of tracking behavior a segmented control exhibits. They are used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking
type SegmentSwitchTracking uint

const (
	// SegmentSwitchTrackingMomentary - A segment is selected only when the user is pressing the mouse down within the bounds of the segment. When the mouse is no longer down within the segment, the segment is automatically deselected. A momentary segmented control sends an action when the user clicks a segment, and another action when the user releases the segment. If configured as continuous (see  ), the control also sends actions at repeating intervals until the user releases the segment, at which point the control sends its final action.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking/momentary
	SegmentSwitchTrackingMomentary SegmentSwitchTracking = 2
	// SegmentSwitchTrackingMomentaryAccelerator - On pressure-sensitive systems, when the user force clicks a segment, a momentary accelerator segmented control sends repeating actions as pressure changes occur. The control stops sending actions when the user releases pressure. A document-based app, for example, might implement a momentary accelerator segmented control in order to allow a user to adjust the speed of paging by using variable pressure. In this example, actions are sent to the app to indicate when pressure on the control has changed. The app then determines the amount of pressure currently applied, and adjusts navigation speed accordingly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking/momentaryAccelerator
	SegmentSwitchTrackingMomentaryAccelerator SegmentSwitchTracking = 3
	// SegmentSwitchTrackingSelectAny - One or more segment cells in the control can be selected at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking/selectAny
	SegmentSwitchTrackingSelectAny SegmentSwitchTracking = 1
	// SegmentSwitchTrackingSelectOne - Only one segment in the control can be selected at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSegmentedControl/SwitchTracking/selectOne
	SegmentSwitchTrackingSelectOne SegmentSwitchTracking = 0
)

/* debug [enums.gen.go]: Processing enum NSSelectionAffinity (2 cases) */
// SelectionAffinity - These constants specify the preferred direction of selection. They’re used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSelectionAffinity
type SelectionAffinity uint

const (
	// SelectionAffinityDownstream - The selection is moving toward the bottom of the document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSelectionAffinity/downstream
	SelectionAffinityDownstream SelectionAffinity = 1
	// SelectionAffinityUpstream - The selection is moving toward the top of the document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSelectionAffinity/upstream
	SelectionAffinityUpstream SelectionAffinity = 0
)

/* debug [enums.gen.go]: Processing enum NSSelectionGranularity (3 cases) */
// SelectionGranularity - These constants specify how much the text view extends the selection when the user drags the mouse. They’re used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSelectionGranularity
type SelectionGranularity uint

const (
	// SelectByCharacter - Extends the selection character by character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSelectionGranularity/selectByCharacter
	SelectByCharacter SelectionGranularity = 0
	// SelectByParagraph - Extends the selection paragraph by paragraph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSelectionGranularity/selectByParagraph
	SelectByParagraph SelectionGranularity = 2
	// SelectByWord - Extends the selection word by word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSelectionGranularity/selectByWord
	SelectByWord SelectionGranularity = 1
)

/* debug [enums.gen.go]: Processing enum NSSharingCollaborationMode (2 cases) */
// SharingCollaborationMode - Represents the types of sharing (collaborating on an item vs. sending a copy of the item)
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode
type SharingCollaborationMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode/collaborate
	SharingCollaborationModeCollaborate SharingCollaborationMode = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingCollaborationMode/sendCopy
	SharingCollaborationModeSendCopy SharingCollaborationMode = 0
)

/* debug [enums.gen.go]: Processing enum NSCloudKitSharingServiceOptions (5 cases) */
// CloudKitSharingServiceOptions - Constants that describe how a participant can configure a CloudKit share.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/CloudKitOptions
type CloudKitSharingServiceOptions uint

const (
	// CloudKitSharingServiceAllowPrivate - An option that allows the participant to privately distribute the share to other iCloud users.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/CloudKitOptions/allowPrivate
	CloudKitSharingServiceAllowPrivate CloudKitSharingServiceOptions = 2
	// CloudKitSharingServiceAllowPublic - An option that allows the participant to publicly distribute the share to other iCloud users.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/CloudKitOptions/allowPublic
	CloudKitSharingServiceAllowPublic CloudKitSharingServiceOptions = 1
	// CloudKitSharingServiceAllowReadOnly - An option that allows the participant to grant other participants read-only permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/CloudKitOptions/allowReadOnly
	CloudKitSharingServiceAllowReadOnly CloudKitSharingServiceOptions = 16
	// CloudKitSharingServiceAllowReadWrite - An option that allows the participant to grant other participants read-write permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/CloudKitOptions/allowReadWrite
	CloudKitSharingServiceAllowReadWrite CloudKitSharingServiceOptions = 32
	// CloudKitSharingServiceStandard - An option that allows the participant to configure the share with a standard set of options.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/CloudKitOptions/standard
	CloudKitSharingServiceStandard CloudKitSharingServiceOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSSharingContentScope (3 cases) */
// SharingContentScope - The sharing scope constants specify the nature of the things you are sharing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/SharingContentScope
type SharingContentScope uint

const (
	// SharingContentScopeFull - Used when sharing the whole content of the current document, for example, the URL of the webpage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/SharingContentScope/full
	SharingContentScopeFull SharingContentScope = 2
	// SharingContentScopeItem - Used when sharing a clearly identified item, for example, a file represented by its icon.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/SharingContentScope/item
	SharingContentScopeItem SharingContentScope = 0
	// SharingContentScopePartial - Used when sharing a portion of a more global content, for example, part of a webpage.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSharingService/SharingContentScope/partial
	SharingContentScopePartial SharingContentScope = 1
)

/* debug [enums.gen.go]: Processing enum NSSliderType (2 cases) */
// SliderType - The types of sliders, used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/SliderType-swift.enum
type SliderType uint

const (
	// SliderTypeCircular - A dial representing an angular range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/SliderType-swift.enum/circular
	SliderTypeCircular SliderType = 1
	// SliderTypeLinear - A bar representing a range, and a knob indicating the currently selected value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/SliderType-swift.enum/linear
	SliderTypeLinear SliderType = 0
)

/* debug [enums.gen.go]: Processing enum NSTickMarkPosition (4 cases) */
// TickMarkPosition - The position where a linear slider’s tick marks appear (above, below, leading, or trailing).
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum
type TickMarkPosition uint

const (
	// TickMarkPositionAbove - A constant indicating that tick marks are displayed above the slider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum/above
	TickMarkPositionAbove TickMarkPosition = 1
	// TickMarkPositionBelow - A constant indicating that tick marks are displayed below the slider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum/below
	TickMarkPositionBelow TickMarkPosition = 0
	// TickMarkPositionLeading - A constant indicating that tick marks are displayed on the leading side of the slider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum/leading
	TickMarkPositionLeading TickMarkPosition = 0
	// TickMarkPositionTrailing - A constant indicating that tick marks are displayed on the trailing side of the slider.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSlider/TickMarkPosition-swift.enum/trailing
	TickMarkPositionTrailing TickMarkPosition = 0
)

/* debug [enums.gen.go]: Processing enum NSSpeechBoundary (3 cases) */
// SpeechBoundary - These constants are used to indicate where speech should be stopped and paused. See 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/Boundary
type SpeechBoundary uint

const (
	// SpeechImmediateBoundary - Speech should be paused or stopped immediately.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/Boundary/immediateBoundary
	SpeechImmediateBoundary SpeechBoundary = 0
	// SpeechSentenceBoundary - Speech should be paused or stopped at the end of the sentence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/Boundary/sentenceBoundary
	SpeechSentenceBoundary SpeechBoundary = 2
	// SpeechWordBoundary - Speech should be paused or stopped at the end of the word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpeechSynthesizer/Boundary/wordBoundary
	SpeechWordBoundary SpeechBoundary = 1
)

/* debug [enums.gen.go]: Processing enum NSSpellingState (2 cases) */
// SpellingState - Constants for the spelling state attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellingState
type SpellingState int

const (
	// SpellingStateGrammarFlag - Flag for grammar issues.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellingState/NSSpellingStateGrammarFlag
	SpellingStateGrammarFlag SpellingState = 1
	// SpellingStateSpellingFlag - Flag for spelling issues.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellingState/NSSpellingStateSpellingFlag
	SpellingStateSpellingFlag SpellingState = 0
)

/* debug [enums.gen.go]: Processing enum NSSplitViewDividerStyle (3 cases) */
// SplitViewDividerStyle - Constants that specify the style of the split view’s dividers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum
type SplitViewDividerStyle uint

const (
	// SplitViewDividerStylePaneSplitter - A thick style divider with a 3D appearance displays between subviews.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum/paneSplitter
	SplitViewDividerStylePaneSplitter SplitViewDividerStyle = 3
	// SplitViewDividerStyleThick - A thick style divider displays between subviews.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum/thick
	SplitViewDividerStyleThick SplitViewDividerStyle = 1
	// SplitViewDividerStyleThin - A thin style divider displays between subviews.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/DividerStyle-swift.enum/thin
	SplitViewDividerStyleThin SplitViewDividerStyle = 2
)

/* debug [enums.gen.go]: Processing enum NSSplitViewItemBehavior (4 cases) */
// SplitViewItemBehavior - Constants that describe the behavior of the split view item.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/Behavior-swift.enum
type SplitViewItemBehavior uint

const (
	// SplitViewItemBehaviorContentList - The content list behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/Behavior-swift.enum/contentList
	SplitViewItemBehaviorContentList SplitViewItemBehavior = 2
	// SplitViewItemBehaviorDefault - The default split view item behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/Behavior-swift.enum/default
	SplitViewItemBehaviorDefault SplitViewItemBehavior = 0
	// SplitViewItemBehaviorInspector - The inspector behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/Behavior-swift.enum/inspector
	SplitViewItemBehaviorInspector SplitViewItemBehavior = 3
	// SplitViewItemBehaviorSidebar - The sidebar behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/Behavior-swift.enum/sidebar
	SplitViewItemBehaviorSidebar SplitViewItemBehavior = 1
)

/* debug [enums.gen.go]: Processing enum NSSplitViewItemCollapseBehavior (4 cases) */
// SplitViewItemCollapseBehavior - Constants that describe the split view item’s collapsing behavior.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/CollapseBehavior-swift.enum
type SplitViewItemCollapseBehavior uint

const (
	// SplitViewItemCollapseBehaviorDefault - The item uses the default collapsing behavior.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/CollapseBehavior-swift.enum/default
	SplitViewItemCollapseBehaviorDefault SplitViewItemCollapseBehavior = 0
	// SplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView - The item’s preference is to resize the other split panes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/CollapseBehavior-swift.enum/preferResizingSiblingsWithFixedSplitView
	SplitViewItemCollapseBehaviorPreferResizingSiblingsWithFixedSplitView SplitViewItemCollapseBehavior = 2
	// SplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings - The item’s preference is to keep the other panes at their current size and position onscreen, potentially growing or shrinking the window in the direction to best preserve that.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/CollapseBehavior-swift.enum/preferResizingSplitViewWithFixedSiblings
	SplitViewItemCollapseBehaviorPreferResizingSplitViewWithFixedSiblings SplitViewItemCollapseBehavior = 1
	// SplitViewItemCollapseBehaviorUseConstraints - The item collapses and expands using a constraint animation, with a constraint priority of the item’s holding priority.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitViewItem/CollapseBehavior-swift.enum/useConstraints
	SplitViewItemCollapseBehaviorUseConstraints SplitViewItemCollapseBehavior = 3
)

/* debug [enums.gen.go]: Processing enum NSSpringLoadingHighlight (3 cases) */
// SpringLoadingHighlight - A group of constants that indicate a highlighting style for your app’s user interface to display during a spring-loading operation.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingHighlight
type SpringLoadingHighlight uint

const (
	// SpringLoadingHighlightEmphasized - A constant that indicates emphasized highlighting to show active spring-loading on the destination.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingHighlight/emphasized
	SpringLoadingHighlightEmphasized SpringLoadingHighlight = 2
	// SpringLoadingHighlightNone - A constant that indicates no highlighting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingHighlight/none
	SpringLoadingHighlightNone SpringLoadingHighlight = 0
	// SpringLoadingHighlightStandard - A constant that indicates standard highlighting to show the destination supports spring-loading.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingHighlight/standard
	SpringLoadingHighlightStandard SpringLoadingHighlight = 1
)

/* debug [enums.gen.go]: Processing enum NSSpringLoadingOptions (4 cases) */
// SpringLoadingOptions - These constants denote the type of spring-loading behavior configured for the destination object.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingOptions
type SpringLoadingOptions uint

const (
	// SpringLoadingContinuousActivation - Spring-loading on the destination object is enabled. The user can drag an object over a destination object and hover or force click to initiate spring-loading and activate the destination object. When initiated by a force click, spring-loading is invoked once the force click begins and deactivated when the force click is released. When initiated by hovering, spring-loading is invoked at the hover timeout and deactivated when the drag exits the destination object. Use this constant sparingly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingOptions/continuousActivation
	SpringLoadingContinuousActivation SpringLoadingOptions = 2
	// SpringLoadingDisabled - Spring-loading on the destination object is disabled. No spring-loading operations can occur.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingOptions/disabled
	SpringLoadingDisabled SpringLoadingOptions = 0
	// SpringLoadingEnabled - Spring-loading on the destination object is enabled. The user can drag an object over a destination object and hover or force click to initiate spring-loading and activate the destination object. When initiated by a force click, spring-loading is invoked once the force click is released.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingOptions/enabled
	SpringLoadingEnabled SpringLoadingOptions = 1
	// SpringLoadingNoHover - Spring-loading on the destination object is enabled, but cannot be invoked by hovering. The user can drag an object over a destination object and force click to initiate spring-loading and activate the destination object. This option may be useful in situations where a long hover, such as dragging across a large destination object, initiates undesired spring-loading. Use this constant sparingly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpringLoadingOptions/noHover
	SpringLoadingNoHover SpringLoadingOptions = 8
)

/* debug [enums.gen.go]: Processing enum NSStackViewDistribution (6 cases) */
// StackViewDistribution enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum
type StackViewDistribution int

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum/equalCentering
	StackViewDistributionEqualCentering StackViewDistribution = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum/equalSpacing
	StackViewDistributionEqualSpacing StackViewDistribution = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum/fill
	StackViewDistributionFill StackViewDistribution = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum/fillEqually
	StackViewDistributionFillEqually StackViewDistribution = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum/fillProportionally
	StackViewDistributionFillProportionally StackViewDistribution = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Distribution-swift.enum/gravityAreas
	StackViewDistributionGravityAreas StackViewDistribution = -1
)

/* debug [enums.gen.go]: Processing enum NSStackViewGravity (5 cases) */
// StackViewGravity - The gravity areas available in a stack view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity
type StackViewGravity uint

const (
	// StackViewGravityBottom - The bottommost gravity area in a vertically oriented stack view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity/bottom
	StackViewGravityBottom StackViewGravity = 3
	// StackViewGravityCenter - The center gravity area, regardless of stack view layout direction or user interface language.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity/center
	StackViewGravityCenter StackViewGravity = 2
	// StackViewGravityLeading - The leftmost or rightmost gravity area in a horizontally oriented stack view, based on the user interface language or the explicitly set user interface layout direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity/leading
	StackViewGravityLeading StackViewGravity = 1
	// StackViewGravityTop - The topmost gravity area in a vertically oriented stack view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity/top
	StackViewGravityTop StackViewGravity = 1
	// StackViewGravityTrailing - The leftmost or rightmost gravity area in a horizontally oriented stack view, based on the user interface language or the explicitly set user interface layout direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/Gravity/trailing
	StackViewGravityTrailing StackViewGravity = 3
)

/* debug [enums.gen.go]: Processing enum NSStringDrawingOptions (1 cases) */
// StringDrawingOptions - Constants that specify the rendering options for drawing a string.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions
type StringDrawingOptions int

const (
	// StringDrawingUsesDeviceMetrics - Uses image glyph bounds instead of typographic bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStringDrawingOptions/NSStringDrawingUsesDeviceMetrics
	StringDrawingUsesDeviceMetrics StringDrawingOptions = 8
)

/* debug [enums.gen.go]: Processing enum NSTabPosition (5 cases) */
// TabPosition enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum
type TabPosition uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum/bottom
	TabPositionBottom TabPosition = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum/left
	TabPositionLeft TabPosition = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum/none
	TabPositionNone TabPosition = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum/right
	TabPositionRight TabPosition = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabPosition-swift.enum/top
	TabPositionTop TabPosition = 1
)

/* debug [enums.gen.go]: Processing enum NSTabViewType (7 cases) */
// TabViewType - These constants specify the tab view’s type as used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType
type TabViewType uint

const (
	// BottomTabsBezelBorder - Tabs are on the bottom of the view with a bezeled border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType/bottomTabsBezelBorder
	BottomTabsBezelBorder TabViewType = 2
	// LeftTabsBezelBorder - Tabs are on the left of the view with a bezeled border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType/leftTabsBezelBorder
	LeftTabsBezelBorder TabViewType = 1
	// NoTabsBezelBorder - The view does not include tabs and has a bezeled border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType/noTabsBezelBorder
	NoTabsBezelBorder TabViewType = 4
	// NoTabsLineBorder - The view does not include tabs and has a lined border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType/noTabsLineBorder
	NoTabsLineBorder TabViewType = 5
	// NoTabsNoBorder - The view does not include tabs and has no border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType/noTabsNoBorder
	NoTabsNoBorder TabViewType = 6
	// RightTabsBezelBorder - Tabs are on the right of the view with a bezeled border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType/rightTabsBezelBorder
	RightTabsBezelBorder TabViewType = 3
	// TopTabsBezelBorder - The view includes tabs on the top of the view and has a bezeled border (the default).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabType/topTabsBezelBorder
	TopTabsBezelBorder TabViewType = 0
)

/* debug [enums.gen.go]: Processing enum NSTabViewBorderType (3 cases) */
// TabViewBorderType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabViewBorderType-swift.enum
type TabViewBorderType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabViewBorderType-swift.enum/bezel
	TabViewBorderTypeBezel TabViewBorderType = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabViewBorderType-swift.enum/line
	TabViewBorderTypeLine TabViewBorderType = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabView/TabViewBorderType-swift.enum/none
	TabViewBorderTypeNone TabViewBorderType = 0
)

/* debug [enums.gen.go]: Processing enum NSTabViewControllerTabStyle (4 cases) */
// TabViewControllerTabStyle - Tab control style options for a tab view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum
type TabViewControllerTabStyle int

const (
	// TabViewControllerTabStyleSegmentedControlOnBottom - A style that displays a segmented control along the bottom edge of the tab view interface. Access the configuration of the tab items through the tab view, which you can get from the   property.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum/segmentedControlOnBottom
	TabViewControllerTabStyleSegmentedControlOnBottom TabViewControllerTabStyle = 1
	// TabViewControllerTabStyleSegmentedControlOnTop - A style that displays a segmented control along the top edge of the tab view interface. Access the configuration of the tab items through the tab view, which you can get from the   property.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum/segmentedControlOnTop
	TabViewControllerTabStyleSegmentedControlOnTop TabViewControllerTabStyle = 0
	// TabViewControllerTabStyleToolbar - A style that automatically adds any tabs to the window’s toolbar. The tab view controller takes control of the window’s toolbar and sets itself as the toolbar’s delegate. Customization of the toolbar is handled using the methods in Responding to Toolbar Events.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum/toolbar
	TabViewControllerTabStyleToolbar TabViewControllerTabStyle = 2
	// TabViewControllerTabStyleUnspecified - A style that indicates the tab view controller does not provide the tab selection UI. Your app provides the control (such as an   or  ) for navigating between tabs. You can bind an existing control to the tab view controller object so that interactions with the control automatically change tabs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewController/TabStyle-swift.enum/unspecified
	TabViewControllerTabStyleUnspecified TabViewControllerTabStyle = -1
)

/* debug [enums.gen.go]: Processing enum NSTabState (3 cases) */
// TabState - These constants describe the current display state of a tab:
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/State
type TabState uint

const (
	// BackgroundTab - A tab that’s not being displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/State/backgroundTab
	BackgroundTab TabState = 1
	// PressedTab - A tab that the user is in the process of clicking. That is, the user has pressed the mouse button while the cursor is over the tab but has not released the mouse button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/State/pressedTab
	PressedTab TabState = 2
	// SelectedTab - The tab that’s being displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTabViewItem/State/selectedTab
	SelectedTab TabState = 0
)

/* debug [enums.gen.go]: Processing enum NSTableColumnResizingOptions (3 cases) */
// TableColumnResizingOptions enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/ResizingOptions
type TableColumnResizingOptions uint

const (
	// TableColumnAutoresizingMask - Allows the table column to resize automatically in response to resizing the table view. The resizing behavior for the table view is set using the   method  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/ResizingOptions/autoresizingMask
	TableColumnAutoresizingMask TableColumnResizingOptions = 1
	// TableColumnUserResizingMask - Allows the table column to be resized by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumn/ResizingOptions/userResizingMask
	TableColumnUserResizingMask TableColumnResizingOptions = 2
	// TableColumnNoResizing - Prevents the table column from resizing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableColumnResizingOptions/NSTableColumnNoResizing
	TableColumnNoResizing TableColumnResizingOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSTableViewAnimationOptions (7 cases) */
// TableViewAnimationOptions - Specifies the animation effects to apply when inserting or removing rows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions
type TableViewAnimationOptions uint

const (
	// TableViewAnimationEffectFade - Use a fade for row or column removal. The effect can be combined with any of the slide constants.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions/effectFade
	TableViewAnimationEffectFade TableViewAnimationOptions = 1
	// TableViewAnimationEffectGap - Creates a gap for newly inserted rows. This is useful for drag and drop animations that animate to a newly opened gap and should be used in the delegate method  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions/effectGap
	TableViewAnimationEffectGap TableViewAnimationOptions = 2
	// TableViewAnimationSlideDown - Animates a row insertion or removal by sliding downward.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions/slideDown
	TableViewAnimationSlideDown TableViewAnimationOptions = 32
	// TableViewAnimationSlideLeft - Animates a row insertion by sliding from the left. Animates a row removal by sliding towards the left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions/slideLeft
	TableViewAnimationSlideLeft TableViewAnimationOptions = 48
	// TableViewAnimationSlideRight - Animates a row insertion by sliding from the right. Animates a row removal by sliding towards the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions/slideRight
	TableViewAnimationSlideRight TableViewAnimationOptions = 64
	// TableViewAnimationSlideUp - Animates a row insertion or removal by sliding upward.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/AnimationOptions/slideUp
	TableViewAnimationSlideUp TableViewAnimationOptions = 16
	// TableViewAnimationEffectNone - Use no animation effects.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewAnimationOptions/NSTableViewAnimationEffectNone
	TableViewAnimationEffectNone TableViewAnimationOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSTableRowActionEdge (2 cases) */
// TableRowActionEdge - These constants define table row edges on which row actions are attached. They are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowActionEdge
type TableRowActionEdge uint

const (
	// TableRowActionEdgeLeading - Denotes the leading, or left, edge of a table row view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowActionEdge/leading
	TableRowActionEdgeLeading TableRowActionEdge = 0
	// TableRowActionEdgeTrailing - Denotes the trailing, or right, edge of a table row view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowActionEdge/trailing
	TableRowActionEdgeTrailing TableRowActionEdge = 1
)

/* debug [enums.gen.go]: Processing enum NSTableViewSelectionHighlightStyle (3 cases) */
// TableViewSelectionHighlightStyle - The following constants specify the selection highlight styles. These constants are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/SelectionHighlightStyle-swift.enum
type TableViewSelectionHighlightStyle uint

const (
	// TableViewSelectionHighlightStyleNone - Displays no highlight style at all.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/SelectionHighlightStyle-swift.enum/none
	TableViewSelectionHighlightStyleNone TableViewSelectionHighlightStyle = 0
	// TableViewSelectionHighlightStyleRegular - The regular highlight style of NSTableView. In OS X v10.7 a light blue (returned by sending   a   message) or light gray color (returned by sending NSColor a   message).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/SelectionHighlightStyle-swift.enum/regular
	TableViewSelectionHighlightStyleRegular TableViewSelectionHighlightStyle = 0
	// TableViewSelectionHighlightStyleSourceList - The source list style of NSTableView. On 10.5, a light blue gradient is used to highlight selected rows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/SelectionHighlightStyle-swift.enum/sourceList
	TableViewSelectionHighlightStyleSourceList TableViewSelectionHighlightStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSTableViewColumnAutoresizingStyle (6 cases) */
// TableViewColumnAutoresizingStyle - The following constants specify the autoresizing styles. These constants are used by the  
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum
type TableViewColumnAutoresizingStyle uint

const (
	// TableViewFirstColumnOnlyAutoresizingStyle - Autoresize only the first table column.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum/firstColumnOnlyAutoresizingStyle
	TableViewFirstColumnOnlyAutoresizingStyle TableViewColumnAutoresizingStyle = 5
	// TableViewLastColumnOnlyAutoresizingStyle - Autoresize only the last table column.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum/lastColumnOnlyAutoresizingStyle
	TableViewLastColumnOnlyAutoresizingStyle TableViewColumnAutoresizingStyle = 4
	// TableViewNoColumnAutoresizing - Disable table column autoresizing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum/noColumnAutoresizing
	TableViewNoColumnAutoresizing TableViewColumnAutoresizingStyle = 0
	// TableViewReverseSequentialColumnAutoresizingStyle - Autoresize each table column sequentially, from the first auto-resizable column to the last auto-resizable column; proceed to the next column when the current column has reached its minimum or maximum size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum/reverseSequentialColumnAutoresizingStyle
	TableViewReverseSequentialColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 3
	// TableViewSequentialColumnAutoresizingStyle - Autoresize each table column sequentially, from the last auto-resizable column to the first auto-resizable column; proceed to the next column when the current column has reached its minimum or maximum size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum/sequentialColumnAutoresizingStyle
	TableViewSequentialColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 2
	// TableViewUniformColumnAutoresizingStyle - Autoresize all columns by distributing space equally, simultaneously.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/ColumnAutoresizingStyle-swift.enum/uniformColumnAutoresizingStyle
	TableViewUniformColumnAutoresizingStyle TableViewColumnAutoresizingStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSTableViewDraggingDestinationFeedbackStyle (4 cases) */
// TableViewDraggingDestinationFeedbackStyle - These constants specify the drag styles displayed by the table view. They’re used by 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum
type TableViewDraggingDestinationFeedbackStyle int

const (
	// TableViewDraggingDestinationFeedbackStyleGap - Provides a gap insertion when dragging over the table. Note that this style is only officially supported for  -based table views, but may partially work in Cell Based TableViews. The decision to use the gap style (compared to another style) can be made in  , or it can dynamically be changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum/gap
	TableViewDraggingDestinationFeedbackStyleGap TableViewDraggingDestinationFeedbackStyle = 2
	// TableViewDraggingDestinationFeedbackStyleNone - Provides no feedback when the user drags over the table view. This option exists to allow subclasses to implement their dragging destination highlighting, or to make it not show anything all.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum/none
	TableViewDraggingDestinationFeedbackStyleNone TableViewDraggingDestinationFeedbackStyle = -1
	// TableViewDraggingDestinationFeedbackStyleRegular - Draws a solid round-rect background on drop target rows, and an insertion marker between rows. This style should be used in most cases.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum/regular
	TableViewDraggingDestinationFeedbackStyleRegular TableViewDraggingDestinationFeedbackStyle = 0
	// TableViewDraggingDestinationFeedbackStyleSourceList - Draws an outline on drop target rows, and an insertion marker between rows. This style will automatically be set for source lists when the table’s   is set to  . This is the standard look for Source Lists, but may be used in other areas as needed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DraggingDestinationFeedbackStyle-swift.enum/sourceList
	TableViewDraggingDestinationFeedbackStyleSourceList TableViewDraggingDestinationFeedbackStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSTableViewDropOperation (2 cases) */
// TableViewDropOperation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DropOperation
type TableViewDropOperation uint

const (
	// TableViewDropAbove - Specifies that the drop should occur above the specified row.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DropOperation/above
	TableViewDropAbove TableViewDropOperation = 1
	// TableViewDropOn - Specifies that the drop should occur on the specified row.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/DropOperation/on
	TableViewDropOn TableViewDropOperation = 0
)

/* debug [enums.gen.go]: Processing enum NSTableViewGridLineStyle (4 cases) */
// TableViewGridLineStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/GridLineStyle
type TableViewGridLineStyle uint

const (
	// TableViewDashedHorizontalGridLineMask - Specifies that the horizontal grid lines should be drawn dashed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/GridLineStyle/dashedHorizontalGridLineMask
	TableViewDashedHorizontalGridLineMask TableViewGridLineStyle = 3
	// TableViewSolidHorizontalGridLineMask - Specifies that horizontal grid lines should be displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/GridLineStyle/solidHorizontalGridLineMask
	TableViewSolidHorizontalGridLineMask TableViewGridLineStyle = 2
	// TableViewSolidVerticalGridLineMask - Specifies that vertical grid lines should be displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/GridLineStyle/solidVerticalGridLineMask
	TableViewSolidVerticalGridLineMask TableViewGridLineStyle = 1
	// TableViewGridNone - Specifies that no grid lines should be displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewGridLineStyle/NSTableViewGridNone
	TableViewGridNone TableViewGridLineStyle = 0
)

/* debug [enums.gen.go]: Processing enum NSTableViewRowSizeStyle (5 cases) */
// TableViewRowSizeStyle - The row size style constants define the size of the rows in the table view. They are used by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum
type TableViewRowSizeStyle int

const (
	// TableViewRowSizeStyleCustom - The table will use the   or invoke the delegate method  , if implemented. The cell layout is not changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/custom
	TableViewRowSizeStyleCustom TableViewRowSizeStyle = 0
	// TableViewRowSizeStyleDefault - The table will use the system default layout size: small, medium or large.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/default
	TableViewRowSizeStyleDefault TableViewRowSizeStyle = -1
	// TableViewRowSizeStyleLarge - The table will use a row height specified for a large table. It is required that the size be fully tested and supported if   is not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/large
	TableViewRowSizeStyleLarge TableViewRowSizeStyle = 3
	// TableViewRowSizeStyleMedium - The table will use a row height specified for a medium table. It is required that the size be fully tested and supported if   is not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/medium
	TableViewRowSizeStyleMedium TableViewRowSizeStyle = 2
	// TableViewRowSizeStyleSmall - The table will use a row height specified for a small table. It is required that the size be fully tested and supported if   is not used.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/RowSizeStyle-swift.enum/small
	TableViewRowSizeStyleSmall TableViewRowSizeStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSTableViewStyle (5 cases) */
// TableViewStyle - Contains the possible style values for a table view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum
type TableViewStyle uint

const (
	// TableViewStyleAutomatic - The system resolves the table view style based on the table view hierarchy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/automatic
	TableViewStyleAutomatic TableViewStyle = 0
	// TableViewStyleFullWidth - The table view style resolves to a full-width style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/fullWidth
	TableViewStyleFullWidth TableViewStyle = 1
	// TableViewStyleInset - The table view style resolves to an inset style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/inset
	TableViewStyleInset TableViewStyle = 2
	// TableViewStylePlain - The table view style resolves to a plain style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/plain
	TableViewStylePlain TableViewStyle = 4
	// TableViewStyleSourceList - The table view style resolves to a source-list style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableView/Style-swift.enum/sourceList
	TableViewStyleSourceList TableViewStyle = 3
)

/* debug [enums.gen.go]: Processing enum NSTableViewRowActionStyle (2 cases) */
// TableViewRowActionStyle - Constants that help define the appearance and behavior of action buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/Style-swift.enum
type TableViewRowActionStyle uint

const (
	// TableViewRowActionStyleDestructive - Apply a style that indicates that the action might change or delete data. This style changes the value of the   property to an appropriate value to reflect the destructive action. After creating the action object, you can change the background color as needed. Destructive actions require a longer swipe to activate, and trigger an animation when a table row is deleted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/Style-swift.enum/destructive
	TableViewRowActionStyleDestructive TableViewRowActionStyle = 1
	// TableViewRowActionStyleRegular - Apply the default style to the button. This style does not apply any special coloring to the button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTableViewRowAction/Style-swift.enum/regular
	TableViewRowActionStyleRegular TableViewRowActionStyle = 0
)

/* debug [enums.gen.go]: Processing enum NSTextAlignment (5 cases) */
// TextAlignment - Constants that specify text alignment.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment
type TextAlignment uint

const (
	// TextAlignmentCenter - Text is center-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment/center
	TextAlignmentCenter TextAlignment = 1
	// TextAlignmentJustified - Text is justified.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment/justified
	TextAlignmentJustified TextAlignment = 3
	// TextAlignmentLeft - Text is left-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment/left
	TextAlignmentLeft TextAlignment = 0
	// TextAlignmentNatural - Text uses the default alignment for the current localization of the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment/natural
	TextAlignmentNatural TextAlignment = 4
	// TextAlignmentRight - Text is right-aligned.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextAlignment/right
	TextAlignmentRight TextAlignment = 2
)

/* debug [enums.gen.go]: Processing enum NSTextBlockDimension (6 cases) */
// TextBlockDimension - The following constants specify values used by the methods 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension
type TextBlockDimension uint

const (
	// TextBlockHeight - Height of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/height
	TextBlockHeight TextBlockDimension = 4
	// TextBlockMaximumHeight - Maximum height of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/maximumHeight
	TextBlockMaximumHeight TextBlockDimension = 6
	// TextBlockMaximumWidth - Maximum width of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/maximumWidth
	TextBlockMaximumWidth TextBlockDimension = 2
	// TextBlockMinimumHeight - Minimum height of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/minimumHeight
	TextBlockMinimumHeight TextBlockDimension = 5
	// TextBlockMinimumWidth - Minimum width of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/minimumWidth
	TextBlockMinimumWidth TextBlockDimension = 1
	// TextBlockWidth - Width of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Dimension/width
	TextBlockWidth TextBlockDimension = 0
)

/* debug [enums.gen.go]: Processing enum NSTextBlockLayer (3 cases) */
// TextBlockLayer - The following constants specify values used by the properties and methods 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Layer
type TextBlockLayer int

const (
	// TextBlockBorder - The border of the text block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Layer/border
	TextBlockBorder TextBlockLayer = 0
	// TextBlockMargin - Margin of the text block: space surrounding the border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Layer/margin
	TextBlockMargin TextBlockLayer = 1
	// TextBlockPadding - Padding of the text block: space surrounding the content area extending to the border.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/Layer/padding
	TextBlockPadding TextBlockLayer = -1
)

/* debug [enums.gen.go]: Processing enum NSTextBlockValueType (2 cases) */
// TextBlockValueType - The following constants specify values used by the methods 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/ValueType
type TextBlockValueType uint

const (
	// TextBlockAbsoluteValueType - Absolute value in points.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/ValueType/absoluteValueType
	TextBlockAbsoluteValueType TextBlockValueType = 0
	// TextBlockPercentageValueType - Percentage value (out of 100).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/ValueType/percentageValueType
	TextBlockPercentageValueType TextBlockValueType = 1
)

/* debug [enums.gen.go]: Processing enum NSTextBlockVerticalAlignment (4 cases) */
// TextBlockVerticalAlignment - The following constants specify values used by the property 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/VerticalAlignment-swift.enum
type TextBlockVerticalAlignment uint

const (
	// TextBlockBaselineAlignment - Aligns adjacent blocks at the baseline of the first line of text in the block.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/VerticalAlignment-swift.enum/baselineAlignment
	TextBlockBaselineAlignment TextBlockVerticalAlignment = 3
	// TextBlockBottomAlignment - Aligns adjacent blocks at their bottom.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/VerticalAlignment-swift.enum/bottomAlignment
	TextBlockBottomAlignment TextBlockVerticalAlignment = 2
	// TextBlockMiddleAlignment - Aligns adjacent blocks at their middle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/VerticalAlignment-swift.enum/middleAlignment
	TextBlockMiddleAlignment TextBlockVerticalAlignment = 1
	// TextBlockTopAlignment - Aligns adjacent blocks at their top.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextBlock/VerticalAlignment-swift.enum/topAlignment
	TextBlockTopAlignment TextBlockVerticalAlignment = 0
)

/* debug [enums.gen.go]: Processing enum NSTextContentManagerEnumerationOptions (2 cases) */
// TextContentManagerEnumerationOptions - Values that control the order in which the framework enumerates text elements.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/EnumerationOptions
type TextContentManagerEnumerationOptions uint

const (
	// TextContentManagerEnumerationOptionsReverse - Returns whether enumerations start from the end of the text element.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManager/EnumerationOptions/reverse
	TextContentManagerEnumerationOptionsReverse TextContentManagerEnumerationOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentManagerEnumerationOptions/NSTextContentManagerEnumerationOptionsNone
	TextContentManagerEnumerationOptionsNone TextContentManagerEnumerationOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSTextCursorAccessoryPlacement (9 cases) */
// TextCursorAccessoryPlacement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement
type TextCursorAccessoryPlacement uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/backward
	TextCursorAccessoryPlacementBackward TextCursorAccessoryPlacement = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/center
	TextCursorAccessoryPlacementCenter TextCursorAccessoryPlacement = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/forward
	TextCursorAccessoryPlacementForward TextCursorAccessoryPlacement = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/invisible
	TextCursorAccessoryPlacementInvisible TextCursorAccessoryPlacement = 3
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/offscreenBottom
	TextCursorAccessoryPlacementOffscreenBottom TextCursorAccessoryPlacement = 8
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/offscreenLeft
	TextCursorAccessoryPlacementOffscreenLeft TextCursorAccessoryPlacement = 5
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/offscreenRight
	TextCursorAccessoryPlacementOffscreenRight TextCursorAccessoryPlacement = 7
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/offscreenTop
	TextCursorAccessoryPlacementOffscreenTop TextCursorAccessoryPlacement = 6
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextCursorAccessoryPlacement/unspecified
	TextCursorAccessoryPlacementUnspecified TextCursorAccessoryPlacement = 0
)

/* debug [enums.gen.go]: Processing enum NSTextFieldBezelStyle (2 cases) */
// TextFieldBezelStyle - The style of bezel the text field displays.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/BezelStyle-swift.enum
type TextFieldBezelStyle uint

const (
	// TextFieldRoundedBezel - A style that draws a bezel with rounded corners around a single-line text field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/BezelStyle-swift.enum/roundedBezel
	TextFieldRoundedBezel TextFieldBezelStyle = 1
	// TextFieldSquareBezel - A style that draws a bezel with square corners around a text field.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextField/BezelStyle-swift.enum/squareBezel
	TextFieldSquareBezel TextFieldBezelStyle = 0
)

/* debug [enums.gen.go]: Processing enum NSTextFinderAction (13 cases) */
// TextFinderAction - These constants specify the user interface item tags that correspond find action. These constants are passed to the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action
type TextFinderAction uint

const (
	// TextFinderActionHideFindInterface - Hides the find bar interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/hideFindInterface
	TextFinderActionHideFindInterface TextFinderAction = 11
	// TextFinderActionHideReplaceInterface - Displays the find bar interface including the replace functionality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/hideReplaceInterface
	TextFinderActionHideReplaceInterface TextFinderAction = 13
	// TextFinderActionNextMatch - The next match, if any, is displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/nextMatch
	TextFinderActionNextMatch TextFinderAction = 2
	// TextFinderActionPreviousMatch - The previous match, if any, is displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/previousMatch
	TextFinderActionPreviousMatch TextFinderAction = 3
	// TextFinderActionReplace - Replaces a single instance of the string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/replace
	TextFinderActionReplace TextFinderAction = 5
	// TextFinderActionReplaceAll - All occurrences of the string are replaced.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/replaceAll
	TextFinderActionReplaceAll TextFinderAction = 4
	// TextFinderActionReplaceAllInSelection - Replaces all occurrences of the string within the current selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/replaceAllInSelection
	TextFinderActionReplaceAllInSelection TextFinderAction = 8
	// TextFinderActionReplaceAndFind - Replaces a single instance of the string and searches for the next match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/replaceAndFind
	TextFinderActionReplaceAndFind TextFinderAction = 6
	// TextFinderActionSelectAll - Selects all matching search strings.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/selectAll
	TextFinderActionSelectAll TextFinderAction = 9
	// TextFinderActionSelectAllInSelection - Selects all matching search strings within the current selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/selectAllInSelection
	TextFinderActionSelectAllInSelection TextFinderAction = 10
	// TextFinderActionSetSearchString - Sets the search string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/setSearchString
	TextFinderActionSetSearchString TextFinderAction = 7
	// TextFinderActionShowFindInterface - The find bar interface is displayed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/showFindInterface
	TextFinderActionShowFindInterface TextFinderAction = 1
	// TextFinderActionShowReplaceInterface - Displays the find bar interface including the replace functionality.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/Action/showReplaceInterface
	TextFinderActionShowReplaceInterface TextFinderAction = 12
)

/* debug [enums.gen.go]: Processing enum NSTextFinderMatchingType (4 cases) */
// TextFinderMatchingType - The following constants indicate the type of search anchor an action should perform.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/MatchingType
type TextFinderMatchingType uint

const (
	// TextFinderMatchingTypeContains - The match contains the string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/MatchingType/contains
	TextFinderMatchingTypeContains TextFinderMatchingType = 0
	// TextFinderMatchingTypeEndsWith - The match ends with the string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/MatchingType/endsWith
	TextFinderMatchingTypeEndsWith TextFinderMatchingType = 3
	// TextFinderMatchingTypeFullWord - The match exactly matches the string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/MatchingType/fullWord
	TextFinderMatchingTypeFullWord TextFinderMatchingType = 2
	// TextFinderMatchingTypeStartsWith - The match begins with the string.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextFinder/MatchingType/startsWith
	TextFinderMatchingTypeStartsWith TextFinderMatchingType = 1
)

/* debug [enums.gen.go]: Processing enum NSTextInputTraitType (1 cases) */
// TextInputTraitType enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputTraitType
type TextInputTraitType uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInputTraitType/yes
	TextInputTraitTypeYes TextInputTraitType = 2
)

/* debug [enums.gen.go]: Processing enum NSTextInsertionIndicatorAutomaticModeOptions (2 cases) */
// TextInsertionIndicatorAutomaticModeOptions - Options that affect the automatic display mode.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/AutomaticModeOptions-swift.struct
type TextInsertionIndicatorAutomaticModeOptions uint

const (
	TextInsertionIndicatorAutomaticModeOptionsShowEffectsView TextInsertionIndicatorAutomaticModeOptions = 1
	TextInsertionIndicatorAutomaticModeOptionsShowWhileTracking TextInsertionIndicatorAutomaticModeOptions = 2
)

/* debug [enums.gen.go]: Processing enum NSTextInsertionIndicatorDisplayMode (1 cases) */
// TextInsertionIndicatorDisplayMode - Constants that determine how to display the system text cursor in a custom text UI.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum
type TextInsertionIndicatorDisplayMode uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextInsertionIndicator/DisplayMode-swift.enum/hidden
	TextInsertionIndicatorDisplayModeHidden TextInsertionIndicatorDisplayMode = 1
)

/* debug [enums.gen.go]: Processing enum NSTextLayoutFragmentEnumerationOptions (5 cases) */
// TextLayoutFragmentEnumerationOptions - Values that describe options for enumerating text layout fragments.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/EnumerationOptions
type TextLayoutFragmentEnumerationOptions uint

const (
	// TextLayoutFragmentEnumerationOptionsEnsuresExtraLineFragment - Synthesize the extra line fragment when necessary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/EnumerationOptions/ensuresExtraLineFragment
	TextLayoutFragmentEnumerationOptionsEnsuresExtraLineFragment TextLayoutFragmentEnumerationOptions = 8
	// TextLayoutFragmentEnumerationOptionsEnsuresLayout - When enumerating, tell the layout fragments to layout their contents.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/EnumerationOptions/ensuresLayout
	TextLayoutFragmentEnumerationOptionsEnsuresLayout TextLayoutFragmentEnumerationOptions = 4
	// TextLayoutFragmentEnumerationOptionsEstimatesSize - When enumerating, tell the layout fragments to estimate their size.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/EnumerationOptions/estimatesSize
	TextLayoutFragmentEnumerationOptionsEstimatesSize TextLayoutFragmentEnumerationOptions = 2
	// TextLayoutFragmentEnumerationOptionsReverse - Causes the enumeration to start from the last element.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/EnumerationOptions/reverse
	TextLayoutFragmentEnumerationOptionsReverse TextLayoutFragmentEnumerationOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragmentEnumerationOptions/NSTextLayoutFragmentEnumerationOptionsNone
	TextLayoutFragmentEnumerationOptionsNone TextLayoutFragmentEnumerationOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSTextLayoutFragmentState (4 cases) */
// TextLayoutFragmentState - Values that describe the possible layout states.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum
type TextLayoutFragmentState uint

const (
	// TextLayoutFragmentStateCalculatedUsageBounds - The layout fragment measurements are available without text line fragments.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum/calculatedUsageBounds
	TextLayoutFragmentStateCalculatedUsageBounds TextLayoutFragmentState = 2
	// TextLayoutFragmentStateEstimatedUsageBounds - The text layout manager hasn’t performed a full layout yet for the region covered by this layout fragment and is returning an estimated bounds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum/estimatedUsageBounds
	TextLayoutFragmentStateEstimatedUsageBounds TextLayoutFragmentState = 1
	// TextLayoutFragmentStateLayoutAvailable - Measurements for the text line fragments and layout fragment are available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum/layoutAvailable
	TextLayoutFragmentStateLayoutAvailable TextLayoutFragmentState = 3
	// TextLayoutFragmentStateNone - No layout information is available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutFragment/State-swift.enum/none
	TextLayoutFragmentStateNone TextLayoutFragmentState = 0
)

/* debug [enums.gen.go]: Processing enum NSTextLayoutManagerSegmentOptions (6 cases) */
// TextLayoutManagerSegmentOptions - Values that describe where and how the framework extends segments of a selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions
type TextLayoutManagerSegmentOptions uint

const (
	// TextLayoutManagerSegmentOptionsHeadSegmentExtended - Returns the value that causes the framework to extend the segment to the tail edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions/headSegmentExtended
	TextLayoutManagerSegmentOptionsHeadSegmentExtended TextLayoutManagerSegmentOptions = 4
	// TextLayoutManagerSegmentOptionsMiddleFragmentsExcluded - Returns the value that causes the framework to enumerate segments in only the first and last line fragments.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions/middleFragmentsExcluded
	TextLayoutManagerSegmentOptionsMiddleFragmentsExcluded TextLayoutManagerSegmentOptions = 2
	// TextLayoutManagerSegmentOptionsRangeNotRequired - Returns the value that causes the framework enumerate text segment rectangles, but avoids preparing a range object.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions/rangeNotRequired
	TextLayoutManagerSegmentOptionsRangeNotRequired TextLayoutManagerSegmentOptions = 1
	// TextLayoutManagerSegmentOptionsTailSegmentExtended - Returns the value that causes the framework to extend the segment to the tail edge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions/tailSegmentExtended
	TextLayoutManagerSegmentOptionsTailSegmentExtended TextLayoutManagerSegmentOptions = 8
	// TextLayoutManagerSegmentOptionsUpstreamAffinity - Returns the value that causes the framework to the place the segment based on the upstream affinity for an empty range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentOptions/upstreamAffinity
	TextLayoutManagerSegmentOptionsUpstreamAffinity TextLayoutManagerSegmentOptions = 16
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManagerSegmentOptions/NSTextLayoutManagerSegmentOptionsNone
	TextLayoutManagerSegmentOptionsNone TextLayoutManagerSegmentOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSTextLayoutManagerSegmentType (3 cases) */
// TextLayoutManagerSegmentType - Values that describe the rendering of selection boundaries.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentType
type TextLayoutManagerSegmentType uint

const (
	// TextLayoutManagerSegmentTypeHighlight - The segment behavior suitable for highlighting.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentType/highlight
	TextLayoutManagerSegmentTypeHighlight TextLayoutManagerSegmentType = 2
	// TextLayoutManagerSegmentTypeSelection - The segment behavior suitable for selection rendering.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentType/selection
	TextLayoutManagerSegmentTypeSelection TextLayoutManagerSegmentType = 1
	// TextLayoutManagerSegmentTypeStandard - The standard segment, matching the typographic bounds of the range.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextLayoutManager/SegmentType/standard
	TextLayoutManagerSegmentTypeStandard TextLayoutManagerSegmentType = 0
)

/* debug [enums.gen.go]: Processing enum NSTextListOptions (1 cases) */
// TextListOptions - Values that available options for text list items.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextList/Options
type TextListOptions uint

const (
	// TextListPrependEnclosingMarker - Specifies that a nested list should include the marker for its enclosing superlist before its own marker.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextList/Options/prependEnclosingMarker
	TextListPrependEnclosingMarker TextListOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSTextMovement (9 cases) */
// TextMovement enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement
type TextMovement uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement/backtab
	TextMovementBacktab TextMovement = 18
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement/cancel
	TextMovementCancel TextMovement = 23
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement/down
	TextMovementDown TextMovement = 22
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement/left
	TextMovementLeft TextMovement = 19
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement/other
	TextMovementOther TextMovement = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement/return
	TextMovementReturn TextMovement = 16
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement/right
	TextMovementRight TextMovement = 20
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement/tab
	TextMovementTab TextMovement = 17
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextMovement/up
	TextMovementUp TextMovement = 21
)

/* debug [enums.gen.go]: Processing enum NSTextScalingType (2 cases) */
// TextScalingType - Constants that specify the text scaling.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextScalingType
type TextScalingType uint

const (
	// TextScalingiOS - Font sizes throughout the document appear visually similar to how they would render in iOS.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextScalingType/iOS
	TextScalingiOS TextScalingType = 1
	// TextScalingStandard - Font sizes throughout the document appear visually similar to how they would render in macOS and non-Apple platforms.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextScalingType/standard
	TextScalingStandard TextScalingType = 0
)

/* debug [enums.gen.go]: Processing enum NSTextSelectionAffinity (2 cases) */
// TextSelectionAffinity - Values that describe the visual location of the text cursor, or the direction of the non-anchored edge of the selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum
type TextSelectionAffinity uint

const (
	// TextSelectionAffinityDownstream - The value that defines the visual location of the text cursor between the head of line that contains the selection location.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum/downstream
	TextSelectionAffinityDownstream TextSelectionAffinity = 1
	// TextSelectionAffinityUpstream - The value that defines the visual location of the text cursor to the tail of the previous line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Affinity-swift.enum/upstream
	TextSelectionAffinityUpstream TextSelectionAffinity = 0
)

/* debug [enums.gen.go]: Processing enum NSTextSelectionGranularity (5 cases) */
// TextSelectionGranularity - Values that describe the different granularities available to make a selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum
type TextSelectionGranularity uint

const (
	// TextSelectionGranularityCharacter - A value that represents selection by character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum/character
	TextSelectionGranularityCharacter TextSelectionGranularity = 0
	// TextSelectionGranularityLine - A value that represents selection by line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum/line
	TextSelectionGranularityLine TextSelectionGranularity = 3
	// TextSelectionGranularityParagraph - A value that represents selection by paragraph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum/paragraph
	TextSelectionGranularityParagraph TextSelectionGranularity = 2
	// TextSelectionGranularitySentence - A value that represents selection by sentence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum/sentence
	TextSelectionGranularitySentence TextSelectionGranularity = 4
	// TextSelectionGranularityWord - A value that represents selection by word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelection/Granularity-swift.enum/word
	TextSelectionGranularityWord TextSelectionGranularity = 1
)

/* debug [enums.gen.go]: Processing enum NSTextSelectionNavigationDestination (7 cases) */
// TextSelectionNavigationDestination - Values that affect how the framework handles navigation across different textual boundaries during a selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination
type TextSelectionNavigationDestination uint

const (
	// TextSelectionNavigationDestinationCharacter - The selection moves to the next extended grapheme cluster boundary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination/character
	TextSelectionNavigationDestinationCharacter TextSelectionNavigationDestination = 0
	// TextSelectionNavigationDestinationContainer - The selection moves to the next container or page boundary after boundary of the current container, ignoring the end of line elastic characters.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination/container
	TextSelectionNavigationDestinationContainer TextSelectionNavigationDestination = 5
	// TextSelectionNavigationDestinationDocument - The selection moves to the document boundary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination/document
	TextSelectionNavigationDestinationDocument TextSelectionNavigationDestination = 6
	// TextSelectionNavigationDestinationLine - The selection moves to the next line boundary.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination/line
	TextSelectionNavigationDestinationLine TextSelectionNavigationDestination = 2
	// TextSelectionNavigationDestinationParagraph - The selection moves to the next paragraph boundary, ignoring the end of line elastic characters and paragraph separators.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination/paragraph
	TextSelectionNavigationDestinationParagraph TextSelectionNavigationDestination = 4
	// TextSelectionNavigationDestinationSentence - The selection moves to the next sentence boundary, ignoring punctuation, whitespace, and format characters preceding the next sentence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination/sentence
	TextSelectionNavigationDestinationSentence TextSelectionNavigationDestination = 3
	// TextSelectionNavigationDestinationWord - The selection moves to the next word boundary ignoring punctuation, whitespace, and format characters preceding the next word.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Destination/word
	TextSelectionNavigationDestinationWord TextSelectionNavigationDestination = 1
)

/* debug [enums.gen.go]: Processing enum NSTextSelectionNavigationDirection (6 cases) */
// TextSelectionNavigationDirection - Values that describe the direction of a selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction
type TextSelectionNavigationDirection uint

const (
	// TextSelectionNavigationDirectionBackward - The value that represents a backward selection based on the flow of text stored in the document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction/backward
	TextSelectionNavigationDirectionBackward TextSelectionNavigationDirection = 1
	// TextSelectionNavigationDirectionDown - The value that represents a selection in the down direction, below the current line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction/down
	TextSelectionNavigationDirectionDown TextSelectionNavigationDirection = 5
	// TextSelectionNavigationDirectionForward - The value that represents a logical forward selection based on the flow of text stored in the document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction/forward
	TextSelectionNavigationDirectionForward TextSelectionNavigationDirection = 0
	// TextSelectionNavigationDirectionLeft - The value that represents a selection in the left direction along the current line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction/left
	TextSelectionNavigationDirectionLeft TextSelectionNavigationDirection = 3
	// TextSelectionNavigationDirectionRight - The value that represents a selection in the right direction along the current line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction/right
	TextSelectionNavigationDirectionRight TextSelectionNavigationDirection = 2
	// TextSelectionNavigationDirectionUp - The value that represents a selection in the up direction, above the current line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Direction/up
	TextSelectionNavigationDirectionUp TextSelectionNavigationDirection = 4
)

/* debug [enums.gen.go]: Processing enum NSTextSelectionNavigationLayoutOrientation (2 cases) */
// TextSelectionNavigationLayoutOrientation - Values that describe the possible layout orientations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/LayoutOrientation
type TextSelectionNavigationLayoutOrientation uint

const (
	// TextSelectionNavigationLayoutOrientationHorizontal - The value that defines horizontal layout orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/LayoutOrientation/horizontal
	TextSelectionNavigationLayoutOrientationHorizontal TextSelectionNavigationLayoutOrientation = 0
	// TextSelectionNavigationLayoutOrientationVertical - The value that defines vertical layout orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/LayoutOrientation/vertical
	TextSelectionNavigationLayoutOrientationVertical TextSelectionNavigationLayoutOrientation = 1
)

/* debug [enums.gen.go]: Processing enum NSTextSelectionNavigationModifier (3 cases) */
// TextSelectionNavigationModifier - Values that describe how the framework handles different kinds of selection modifiers.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Modifier
type TextSelectionNavigationModifier uint

const (
	// TextSelectionNavigationModifierExtend - The value that indicates the framework extends the selection by not moving the initial location while in a drag selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Modifier/extend
	TextSelectionNavigationModifierExtend TextSelectionNavigationModifier = 1
	// TextSelectionNavigationModifierMultiple - The value that indicates the framework extends the selection visually inside the rectangular area defined by the anchor and dragged positions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Modifier/multiple
	TextSelectionNavigationModifierMultiple TextSelectionNavigationModifier = 4
	// TextSelectionNavigationModifierVisual - The value that indicates the framework extends the selection visually inside the rectangular area defined by the anchor and drag positions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/Modifier/visual
	TextSelectionNavigationModifierVisual TextSelectionNavigationModifier = 2
)

/* debug [enums.gen.go]: Processing enum NSTextSelectionNavigationWritingDirection (2 cases) */
// TextSelectionNavigationWritingDirection - Values that describe the writing direction inside a text selection.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/WritingDirection
type TextSelectionNavigationWritingDirection uint

const (
	// TextSelectionNavigationWritingDirectionLeftToRight - The value that defines the left to right writing direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/WritingDirection/leftToRight
	TextSelectionNavigationWritingDirectionLeftToRight TextSelectionNavigationWritingDirection = 0
	// TextSelectionNavigationWritingDirectionRightToLeft - The value that defines the right to left writing direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextSelectionNavigation/WritingDirection/rightToLeft
	TextSelectionNavigationWritingDirectionRightToLeft TextSelectionNavigationWritingDirection = 1
)

/* debug [enums.gen.go]: Processing enum NSTextStorageEditActions (2 cases) */
// TextStorageEditActions - Constants that indicate the types of changes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorageEditActions
type TextStorageEditActions uint

const (
	// TextStorageEditedAttributes - Attributes were added, removed, or changed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorageEditActions/editedAttributes
	TextStorageEditedAttributes TextStorageEditActions = 0
	// TextStorageEditedCharacters - Characters were added, removed, or replaced.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextStorageEditActions/editedCharacters
	TextStorageEditedCharacters TextStorageEditActions = 1
)

/* debug [enums.gen.go]: Processing enum NSTextTableLayoutAlgorithm (2 cases) */
// TextTableLayoutAlgorithm - These constants, specifying the type of text table layout algorithm, are used with 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/LayoutAlgorithm-swift.enum
type TextTableLayoutAlgorithm uint

const (
	// TextTableAutomaticLayoutAlgorithm - Specifies automatic layout algorithm
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/LayoutAlgorithm-swift.enum/automaticLayoutAlgorithm
	TextTableAutomaticLayoutAlgorithm TextTableLayoutAlgorithm = 0
	// TextTableFixedLayoutAlgorithm - Specifies fixed layout algorithm
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextTable/LayoutAlgorithm-swift.enum/fixedLayoutAlgorithm
	TextTableFixedLayoutAlgorithm TextTableLayoutAlgorithm = 1
)

/* debug [enums.gen.go]: Processing enum NSTintProminence (4 cases) */
// TintProminence - Controls how strongly the tint color applies in a view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence
type TintProminence uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence/automatic
	TintProminenceAutomatic TintProminence = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence/none
	TintProminenceNone TintProminence = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence/primary
	TintProminencePrimary TintProminence = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTintProminence/secondary
	TintProminenceSecondary TintProminence = 3
)

/* debug [enums.gen.go]: Processing enum NSTitlebarSeparatorStyle (4 cases) */
// TitlebarSeparatorStyle - Styles that determine the type of separator displayed between the title bar and content of a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle
type TitlebarSeparatorStyle uint

const (
	// TitlebarSeparatorStyleAutomatic - A style indicating that the system determines the type of separator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/automatic
	TitlebarSeparatorStyleAutomatic TitlebarSeparatorStyle = 0
	// TitlebarSeparatorStyleLine - A style indicating that the title bar separator is a line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/line
	TitlebarSeparatorStyleLine TitlebarSeparatorStyle = 2
	// TitlebarSeparatorStyleNone - A style indicating that there’s no title bar separator.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/none
	TitlebarSeparatorStyleNone TitlebarSeparatorStyle = 1
	// TitlebarSeparatorStyleShadow - A style indicating that the title bar separator is a shadow.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTitlebarSeparatorStyle/shadow
	TitlebarSeparatorStyleShadow TitlebarSeparatorStyle = 3
)

/* debug [enums.gen.go]: Processing enum NSTokenStyle (5 cases) */
// TokenStyle - The NSTokenStyle constants define how tokens are displayed and editable in the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum
type TokenStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum/default
	TokenStyleDefault TokenStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum/none
	TokenStyleNone TokenStyle = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum/plainSquared
	TokenStylePlainSquared TokenStyle = 4
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum/rounded
	TokenStyleRounded TokenStyle = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum/squared
	TokenStyleSquared TokenStyle = 3
)

/* debug [enums.gen.go]: Processing enum NSToolbarDisplayMode (4 cases) */
// ToolbarDisplayMode - Constants that indicate whether the toolbar displays items using a name, icon, or combination of elements.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum
type ToolbarDisplayMode uint

const (
	// ToolbarDisplayModeDefault - The default display mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/default
	ToolbarDisplayModeDefault ToolbarDisplayMode = 0
	// ToolbarDisplayModeIconAndLabel - The toolbar displays an icon and label for each item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/iconAndLabel
	ToolbarDisplayModeIconAndLabel ToolbarDisplayMode = 1
	// ToolbarDisplayModeIconOnly - The toolbar displays only an icon for each item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/iconOnly
	ToolbarDisplayModeIconOnly ToolbarDisplayMode = 2
	// ToolbarDisplayModeLabelOnly - The toolbar displays only a label for each item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum/labelOnly
	ToolbarDisplayModeLabelOnly ToolbarDisplayMode = 3
)

/* debug [enums.gen.go]: Processing enum NSToolbarSizeMode (3 cases) */
// ToolbarSizeMode - Constants that specify toolbar display modes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum
type ToolbarSizeMode uint

const (
	// ToolbarSizeModeDefault - The toolbar uses the system-defined default size, which is  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum/default
	ToolbarSizeModeDefault ToolbarSizeMode = 0
	// ToolbarSizeModeRegular - The toolbar uses regular-sized controls and 32 by 32 pixel icons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum/regular
	ToolbarSizeModeRegular ToolbarSizeMode = 1
	// ToolbarSizeModeSmall - The toolbar uses small-sized controls and 24 by 24 pixel icons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/SizeMode-swift.enum/small
	ToolbarSizeModeSmall ToolbarSizeMode = 2
)

/* debug [enums.gen.go]: Processing enum NSToolbarItemStyle (2 cases) */
// ToolbarItemStyle enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Style-swift.enum
type ToolbarItemStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Style-swift.enum/plain
	ToolbarItemStylePlain ToolbarItemStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Style-swift.enum/prominent
	ToolbarItemStyleProminent ToolbarItemStyle = 1
)

/* debug [enums.gen.go]: Processing enum NSToolbarItemGroupControlRepresentation (3 cases) */
// ToolbarItemGroupControlRepresentation enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum
type ToolbarItemGroupControlRepresentation uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum/automatic
	ToolbarItemGroupControlRepresentationAutomatic ToolbarItemGroupControlRepresentation = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum/collapsed
	ToolbarItemGroupControlRepresentationCollapsed ToolbarItemGroupControlRepresentation = 2
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/ControlRepresentation-swift.enum/expanded
	ToolbarItemGroupControlRepresentationExpanded ToolbarItemGroupControlRepresentation = 1
)

/* debug [enums.gen.go]: Processing enum NSToolbarItemGroupSelectionMode (3 cases) */
// ToolbarItemGroupSelectionMode - A value that indicates how a grouped toolbar item selects its subitems.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum
type ToolbarItemGroupSelectionMode uint

const (
	// ToolbarItemGroupSelectionModeMomentary - The system temporarily highlights the select group item when the user selects the item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/momentary
	ToolbarItemGroupSelectionModeMomentary ToolbarItemGroupSelectionMode = 2
	// ToolbarItemGroupSelectionModeSelectAny - The system toggles a highlight on any item selected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/selectAny
	ToolbarItemGroupSelectionModeSelectAny ToolbarItemGroupSelectionMode = 1
	// ToolbarItemGroupSelectionModeSelectOne - The system displays a highlighted mode on the most recent item selected.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItemGroup/SelectionMode-swift.enum/selectOne
	ToolbarItemGroupSelectionModeSelectOne ToolbarItemGroupSelectionMode = 0
)

/* debug [enums.gen.go]: Processing enum NSTouchPhase (7 cases) */
// TouchPhase - The possible phases of a touch.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct
type TouchPhase uint

const (
	// TouchPhaseAny - Matches any phase of a touch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct/any
	TouchPhaseAny TouchPhase = 0
	// TouchPhaseBegan - A finger touched the device. Or, a resting touch transitioned to an active touch and resting touches are not wanted by the view hierarchy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct/began
	TouchPhaseBegan TouchPhase = 1
	// TouchPhaseCancelled - The system cancelled tracking for the touch, as when (for example) the window associated with the touch resigns key or is deactivated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct/cancelled
	TouchPhaseCancelled TouchPhase = 16
	// TouchPhaseEnded - A finger was lifted from the screen. Or, an active touch transitioned to a resting touch and resting touches are not wanted by the view hierarchy.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct/ended
	TouchPhaseEnded TouchPhase = 8
	// TouchPhaseMoved - A finger moved on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct/moved
	TouchPhaseMoved TouchPhase = 2
	// TouchPhaseStationary - A finger is touching the device, but hasn’t moved since the previous event.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct/stationary
	TouchPhaseStationary TouchPhase = 4
	// TouchPhaseTouching - Matches the  ,  , or   phases of a touch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/Phase-swift.struct/touching
	TouchPhaseTouching TouchPhase = 0
)

/* debug [enums.gen.go]: Processing enum NSTouchType (2 cases) */
// TouchType - A bit mask identifying a direct or indirect touch type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType
type TouchType uint

const (
	// TouchTypeDirect - A direct touch from a user’s finger on a screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType/direct
	TouchTypeDirect TouchType = 0
	// TouchTypeIndirect - An indirect touch that is not on a screen, like a digitizer touch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchType/indirect
	TouchTypeIndirect TouchType = 1
)

/* debug [enums.gen.go]: Processing enum NSTouchTypeMask (2 cases) */
// TouchTypeMask - A bit mask identifying a direct or indirect touch type.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchTypeMask
type TouchTypeMask uint

const (
	// TouchTypeMaskDirect - A direct touch from a user’s finger on a screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchTypeMask/direct
	TouchTypeMaskDirect TouchTypeMask = 0
	// TouchTypeMaskIndirect - An indirect touch that is not on a screen, like a digitizer touch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouch/TouchTypeMask/indirect
	TouchTypeMaskIndirect TouchTypeMask = 0
)

/* debug [enums.gen.go]: Processing enum NSTypesetterControlCharacterAction (6 cases) */
// TypesetterControlCharacterAction - The following constants are possible values returned by the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetterControlCharacterAction
type TypesetterControlCharacterAction uint

const (
	// TypesetterContainerBreakAction - Causes container break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetterControlCharacterAction/containerBreakAction
	TypesetterContainerBreakAction TypesetterControlCharacterAction = 32
	// TypesetterHorizontalTabAction - Treated as tab character.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetterControlCharacterAction/horizontalTabAction
	TypesetterHorizontalTabAction TypesetterControlCharacterAction = 4
	// TypesetterLineBreakAction - Causes line break.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetterControlCharacterAction/lineBreakAction
	TypesetterLineBreakAction TypesetterControlCharacterAction = 8
	// TypesetterParagraphBreakAction - Causes paragraph break; the value returned by   is the advancement used for the following glyph.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetterControlCharacterAction/paragraphBreakAction
	TypesetterParagraphBreakAction TypesetterControlCharacterAction = 16
	// TypesetterWhitespaceAction - The width for glyphs with this action are determined by  , if the method is implemented; otherwise, same as  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetterControlCharacterAction/whitespaceAction
	TypesetterWhitespaceAction TypesetterControlCharacterAction = 2
	// TypesetterZeroAdvancementAction - Glyphs with this action are filtered out from layout  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTypesetterControlCharacterAction/zeroAdvancementAction
	TypesetterZeroAdvancementAction TypesetterControlCharacterAction = 1
)

/* debug [enums.gen.go]: Processing enum NSUnderlineStyle (10 cases) */
// UnderlineStyle - Constants for the underline style and strikethrough style attribute keys.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle
type UnderlineStyle uint

const (
	// UnderlineStyleNone - Don’t draw a line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/NSUnderlineStyleNone
	UnderlineStyleNone UnderlineStyle = 0
	// UnderlineStylePatternSolid - Draw a solid line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/NSUnderlineStylePatternSolid
	UnderlineStylePatternSolid UnderlineStyle = 4
	// UnderlineStyleByWord - Draw the line only beneath or through words, not whitespace.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/byWord
	UnderlineStyleByWord UnderlineStyle = 9
	// UnderlineStyleDouble - Draw a double line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/double
	UnderlineStyleDouble UnderlineStyle = 3
	// UnderlineStylePatternDash - Draw a line of dashes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDash
	UnderlineStylePatternDash UnderlineStyle = 6
	// UnderlineStylePatternDashDot - Draw a line of alternating dashes and dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDashDot
	UnderlineStylePatternDashDot UnderlineStyle = 7
	// UnderlineStylePatternDashDotDot - Draw a line of alternating dashes and two dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDashDotDot
	UnderlineStylePatternDashDotDot UnderlineStyle = 8
	// UnderlineStylePatternDot - Draw a line of dots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/patternDot
	UnderlineStylePatternDot UnderlineStyle = 5
	// UnderlineStyleSingle - Draw a single line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/single
	UnderlineStyleSingle UnderlineStyle = 1
	// UnderlineStyleThick - Draw a thick line.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUnderlineStyle/thick
	UnderlineStyleThick UnderlineStyle = 2
)

/* debug [enums.gen.go]: Processing enum NSUserInterfaceLayoutDirection (2 cases) */
// UserInterfaceLayoutDirection - Specifies the directional flow of the user interface.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection
type UserInterfaceLayoutDirection uint

const (
	// UserInterfaceLayoutDirectionLeftToRight - Layout direction is left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection/leftToRight
	UserInterfaceLayoutDirectionLeftToRight UserInterfaceLayoutDirection = 0
	// UserInterfaceLayoutDirectionRightToLeft - Layout direction is right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection/rightToLeft
	UserInterfaceLayoutDirectionRightToLeft UserInterfaceLayoutDirection = 1
)

/* debug [enums.gen.go]: Processing enum NSUserInterfaceLayoutOrientation (2 cases) */
// UserInterfaceLayoutOrientation - The stack view layout directions, and user interface axes for hugging priority and clipping resistance.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation
type UserInterfaceLayoutOrientation uint

const (
	// UserInterfaceLayoutOrientationHorizontal - The horizontal orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation/horizontal
	UserInterfaceLayoutOrientationHorizontal UserInterfaceLayoutOrientation = 0
	// UserInterfaceLayoutOrientationVertical - The vertical orientation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutOrientation/vertical
	UserInterfaceLayoutOrientationVertical UserInterfaceLayoutOrientation = 1
)

/* debug [enums.gen.go]: Processing enum NSVerticalDirections (3 cases) */
// VerticalDirections - The directions on the vertical axis.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections
type VerticalDirections uint

const (
	// VerticalDirectionsAll - All vertical directions (up and down).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections/NSVerticalDirectionsAll
	VerticalDirectionsAll VerticalDirections = 0
	// VerticalDirectionsDown - The downward direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections/NSVerticalDirectionsDown
	VerticalDirectionsDown VerticalDirections = 2
	// VerticalDirectionsUp - The upwards direction.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVerticalDirections/NSVerticalDirectionsUp
	VerticalDirectionsUp VerticalDirections = 1
)

/* debug [enums.gen.go]: Processing enum NSAutoresizingMaskOptions (7 cases) */
// AutoresizingMaskOptions - Constants that specify the autoresizing behaviors for views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct
type AutoresizingMaskOptions uint

const (
	// ViewHeightSizable - The view’s height is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct/height
	ViewHeightSizable AutoresizingMaskOptions = 16
	// ViewMaxXMargin - The right margin between the view and its superview is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct/maxXMargin
	ViewMaxXMargin AutoresizingMaskOptions = 4
	// ViewMaxYMargin - The top margin between the view and its superview is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct/maxYMargin
	ViewMaxYMargin AutoresizingMaskOptions = 32
	// ViewMinXMargin - The left margin between the view and its superview is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct/minXMargin
	ViewMinXMargin AutoresizingMaskOptions = 1
	// ViewMinYMargin - The bottom margin between the view and its superview is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct/minYMargin
	ViewMinYMargin AutoresizingMaskOptions = 8
	// ViewNotSizable - The view cannot be resized.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct/none
	ViewNotSizable AutoresizingMaskOptions = 0
	// ViewWidthSizable - The view’s width is flexible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/AutoresizingMask-swift.struct/width
	ViewWidthSizable AutoresizingMaskOptions = 2
)

/* debug [enums.gen.go]: Processing enum NSBackgroundStyle (4 cases) */
// BackgroundStyle - Background styles to apply to a view’s cell.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle
type BackgroundStyle uint

const (
	// BackgroundStyleEmphasized - A style that adds emphasis to the background using an alternate color or visual effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle/emphasized
	BackgroundStyleEmphasized BackgroundStyle = 1
	// BackgroundStyleLowered - A style that makes the background appear lower than the content drawn on it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle/lowered
	BackgroundStyleLowered BackgroundStyle = 3
	// BackgroundStyleNormal - A style that reflects the predominant color scheme of the view’s appearance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle/normal
	BackgroundStyleNormal BackgroundStyle = 0
	// BackgroundStyleRaised - A style that makes the background appear higher than the content drawn on it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/BackgroundStyle/raised
	BackgroundStyleRaised BackgroundStyle = 2
)

/* debug [enums.gen.go]: Processing enum NSViewLayerContentsPlacement (12 cases) */
// ViewLayerContentsPlacement - These constants specify the location of the layer content when the content is not rerendered in response to view resizing. For more information, see the 
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum
type ViewLayerContentsPlacement uint

const (
	// ViewLayerContentsPlacementBottom - The content is horizontally centered at the bottom-edge of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/bottom
	ViewLayerContentsPlacementBottom ViewLayerContentsPlacement = 8
	// ViewLayerContentsPlacementBottomLeft - The content is positioned in the bottom-left corner of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/bottomLeft
	ViewLayerContentsPlacementBottomLeft ViewLayerContentsPlacement = 9
	// ViewLayerContentsPlacementBottomRight - The content is positioned in the bottom-right corner of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/bottomRight
	ViewLayerContentsPlacementBottomRight ViewLayerContentsPlacement = 7
	// ViewLayerContentsPlacementCenter - The content is horizontally and vertically centered in the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/center
	ViewLayerContentsPlacementCenter ViewLayerContentsPlacement = 3
	// ViewLayerContentsPlacementLeft - The content is vertically centered at the left-edge of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/left
	ViewLayerContentsPlacementLeft ViewLayerContentsPlacement = 10
	// ViewLayerContentsPlacementRight - The content is vertically centered at the right-edge of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/right
	ViewLayerContentsPlacementRight ViewLayerContentsPlacement = 6
	// ViewLayerContentsPlacementScaleAxesIndependently - The content is resized to fit the entire bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/scaleAxesIndependently
	ViewLayerContentsPlacementScaleAxesIndependently ViewLayerContentsPlacement = 0
	// ViewLayerContentsPlacementScaleProportionallyToFill - The content is resized to completely fill the bounds rectangle, while still preserving the aspect of the content. The content is centered in the axis it exceeds.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/scaleProportionallyToFill
	ViewLayerContentsPlacementScaleProportionallyToFill ViewLayerContentsPlacement = 2
	// ViewLayerContentsPlacementScaleProportionallyToFit - The content is resized to fit the bounds rectangle, preserving the aspect of the content. If the content does not completely fill the bounds rectangle, the content is centered in the partial axis.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/scaleProportionallyToFit
	ViewLayerContentsPlacementScaleProportionallyToFit ViewLayerContentsPlacement = 1
	// ViewLayerContentsPlacementTop - The content is horizontally centered at the top-edge of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/top
	ViewLayerContentsPlacementTop ViewLayerContentsPlacement = 4
	// ViewLayerContentsPlacementTopLeft - The content is positioned in the top-left corner of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/topLeft
	ViewLayerContentsPlacementTopLeft ViewLayerContentsPlacement = 11
	// ViewLayerContentsPlacementTopRight - The content is positioned in the top-right corner of the bounds rectangle.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsPlacement-swift.enum/topRight
	ViewLayerContentsPlacementTopRight ViewLayerContentsPlacement = 5
)

/* debug [enums.gen.go]: Processing enum NSViewLayerContentsRedrawPolicy (5 cases) */
// ViewLayerContentsRedrawPolicy - Constants that specify how layer resizing is handled when a view is layer-backed or layer-hosting. For more information, see the  
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum
type ViewLayerContentsRedrawPolicy uint

const (
	// ViewLayerContentsRedrawBeforeViewResize - Resize the layer and redraw the view to the layer when the view’s size changes. This will be done just once at the beginning of a resize animation, not at each frame of the animation. Affected parts of the layer will also be redrawn when the view is marked as needing display. This mode is a superset of  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum/beforeViewResize
	ViewLayerContentsRedrawBeforeViewResize ViewLayerContentsRedrawPolicy = 3
	// ViewLayerContentsRedrawCrossfade - Redraw the layer contents at the new size and crossfade from the old contents to the new contents. Use this in conjunction with the   constants to get a nice crossfade animation for complex layer-backed views that cannot update correctly at each step of the animation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum/crossfade
	ViewLayerContentsRedrawCrossfade ViewLayerContentsRedrawPolicy = 4
	// ViewLayerContentsRedrawDuringViewResize - Resize the view’s backing-layer and redraw the view to the layer when the view’s size changes. If the resize is animated, AppKit will drive the resize animation itself and will do this resize and redraw at each step of the animation. Affected parts of the layer will also be redrawn when the view is marked as needing display. This mode is a superset of  . This is the way that layer-backed views are currently treated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum/duringViewResize
	ViewLayerContentsRedrawDuringViewResize ViewLayerContentsRedrawPolicy = 2
	// ViewLayerContentsRedrawNever - Leave the layer’s contents alone. Never mark the layer as needing display, or draw the view’s contents to the layer. This is how developer created layers (layer-hosting views) are treated.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum/never
	ViewLayerContentsRedrawNever ViewLayerContentsRedrawPolicy = 0
	// ViewLayerContentsRedrawOnSetNeedsDisplay - Any of the   methods sent to the view will cause the view redraw the affected layer parts by invoking the view’s  , but neither the layer or the view are marked as needing display when the view’s size changes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/LayerContentsRedrawPolicy-swift.enum/onSetNeedsDisplay
	ViewLayerContentsRedrawOnSetNeedsDisplay ViewLayerContentsRedrawPolicy = 1
)

/* debug [enums.gen.go]: Processing enum NSViewControllerTransitionOptions (9 cases) */
// ViewControllerTransitionOptions - Animation options for view transitions in a view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions
type ViewControllerTransitionOptions uint

const (
	// ViewControllerTransitionAllowUserInteraction - A transition animation that allows user interaction during the transition.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/allowUserInteraction
	ViewControllerTransitionAllowUserInteraction ViewControllerTransitionOptions = 4096
	// ViewControllerTransitionCrossfade - A transition animation that fades the new view in and simultaneously fades the old view out. You can combine this animation option with any of the “slide” options in this enumeration.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/crossfade
	ViewControllerTransitionCrossfade ViewControllerTransitionOptions = 1
	// ViewControllerTransitionSlideBackward - A transition animation that reflects the user interface layout direction ( ) in a “backward” manner, as follows
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideBackward
	ViewControllerTransitionSlideBackward ViewControllerTransitionOptions = 384
	// ViewControllerTransitionSlideDown - A transition animation that slides the old view down while the new view slides into view from the top. In other words, both views slide down.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideDown
	ViewControllerTransitionSlideDown ViewControllerTransitionOptions = 32
	// ViewControllerTransitionSlideForward - A transition animation that reflects the user interface layout direction ( ) in a “forward” manner, as follows:
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideForward
	ViewControllerTransitionSlideForward ViewControllerTransitionOptions = 320
	// ViewControllerTransitionSlideLeft - A transition animation that slides the old view to the left while the new view slides into view from the right. In other words, both views slide to the left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideLeft
	ViewControllerTransitionSlideLeft ViewControllerTransitionOptions = 64
	// ViewControllerTransitionSlideRight - A transition animation that slides the old view to the right while the new view slides into view from the left.  In other words, both views slide to the right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideRight
	ViewControllerTransitionSlideRight ViewControllerTransitionOptions = 128
	// ViewControllerTransitionSlideUp - A transition animation that slides the old view up while the new view comes into view from the bottom.  In other words, both views slide up.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewController/TransitionOptions/slideUp
	ViewControllerTransitionSlideUp ViewControllerTransitionOptions = 16
	// ViewControllerTransitionNone - A transition with no animation (the default). Specifying another animation option from this enumeration overrides this option.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewControllerTransitionOptions/NSViewControllerTransitionNone
	ViewControllerTransitionNone ViewControllerTransitionOptions = 0
)

/* debug [enums.gen.go]: Processing enum NSViewLayoutRegionAdaptivityAxis (3 cases) */
// ViewLayoutRegionAdaptivityAxis enum type
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis
type ViewLayoutRegionAdaptivityAxis int

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis/NSViewLayoutRegionAdaptivityAxisHorizontal
	ViewLayoutRegionAdaptivityAxisHorizontal ViewLayoutRegionAdaptivityAxis = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis/NSViewLayoutRegionAdaptivityAxisNone
	ViewLayoutRegionAdaptivityAxisNone ViewLayoutRegionAdaptivityAxis = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewLayoutRegionAdaptivityAxis/NSViewLayoutRegionAdaptivityAxisVertical
	ViewLayoutRegionAdaptivityAxisVertical ViewLayoutRegionAdaptivityAxis = 0
)

/* debug [enums.gen.go]: Processing enum NSVisualEffectBlendingMode (2 cases) */
// VisualEffectBlendingMode - Constants that specify whether the visual effect view blends with what’s either behind or within the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/BlendingMode-swift.enum
type VisualEffectBlendingMode uint

const (
	// VisualEffectBlendingModeBehindWindow - A mode that blends and blurs the visual effect view with the contents behind the window, such as the desktop or other windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/BlendingMode-swift.enum/behindWindow
	VisualEffectBlendingModeBehindWindow VisualEffectBlendingMode = 0
	// VisualEffectBlendingModeWithinWindow - A mode that blends and blurs the visual effect view with contents behind the view in the current window only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/BlendingMode-swift.enum/withinWindow
	VisualEffectBlendingModeWithinWindow VisualEffectBlendingMode = 1
)

/* debug [enums.gen.go]: Processing enum NSVisualEffectMaterial (19 cases) */
// VisualEffectMaterial - Constants to specify the material shown by the visual effect view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum
type VisualEffectMaterial uint

const (
	// VisualEffectMaterialAppearanceBased - A default material for the view’s effective appearance.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/appearanceBased
	VisualEffectMaterialAppearanceBased VisualEffectMaterial = 17
	// VisualEffectMaterialContentBackground - The material for the background of opaque content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/contentBackground
	VisualEffectMaterialContentBackground VisualEffectMaterial = 14
	// VisualEffectMaterialDark - A material with a dark effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/dark
	VisualEffectMaterialDark VisualEffectMaterial = 19
	// VisualEffectMaterialFullScreenUI - The material for the background of a full-screen modal interface.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/fullScreenUI
	VisualEffectMaterialFullScreenUI VisualEffectMaterial = 12
	// VisualEffectMaterialHeaderView - The material for in-line header or footer views.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/headerView
	VisualEffectMaterialHeaderView VisualEffectMaterial = 8
	// VisualEffectMaterialHUDWindow - The material for the background of heads-up display (HUD) windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/hudWindow
	VisualEffectMaterialHUDWindow VisualEffectMaterial = 11
	// VisualEffectMaterialLight - A material with a light effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/light
	VisualEffectMaterialLight VisualEffectMaterial = 18
	// VisualEffectMaterialMediumLight - A material with a medium-light effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/mediumLight
	VisualEffectMaterialMediumLight VisualEffectMaterial = 20
	// VisualEffectMaterialMenu - The material for menus.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/menu
	VisualEffectMaterialMenu VisualEffectMaterial = 5
	// VisualEffectMaterialPopover - The material for the background of popover windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/popover
	VisualEffectMaterialPopover VisualEffectMaterial = 6
	// VisualEffectMaterialSelection - The material used to indicate a selection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/selection
	VisualEffectMaterialSelection VisualEffectMaterial = 4
	// VisualEffectMaterialSheet - The material for the background of sheet windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/sheet
	VisualEffectMaterialSheet VisualEffectMaterial = 9
	// VisualEffectMaterialSidebar - The material for the background of window sidebars.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/sidebar
	VisualEffectMaterialSidebar VisualEffectMaterial = 7
	// VisualEffectMaterialTitlebar - The material for a window’s titlebar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/titlebar
	VisualEffectMaterialTitlebar VisualEffectMaterial = 3
	// VisualEffectMaterialToolTip - The material for the background of a tool tip.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/toolTip
	VisualEffectMaterialToolTip VisualEffectMaterial = 13
	// VisualEffectMaterialUltraDark - A material with an ultra-dark effect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/ultraDark
	VisualEffectMaterialUltraDark VisualEffectMaterial = 21
	// VisualEffectMaterialUnderPageBackground - The material for the area behind the pages of a document.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/underPageBackground
	VisualEffectMaterialUnderPageBackground VisualEffectMaterial = 16
	// VisualEffectMaterialUnderWindowBackground - The material to show under a window’s background.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/underWindowBackground
	VisualEffectMaterialUnderWindowBackground VisualEffectMaterial = 15
	// VisualEffectMaterialWindowBackground - The material for the background of opaque windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/Material-swift.enum/windowBackground
	VisualEffectMaterialWindowBackground VisualEffectMaterial = 10
)

/* debug [enums.gen.go]: Processing enum NSVisualEffectState (3 cases) */
// VisualEffectState - Constants to specify how the material appearance should reflect window activity state.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum
type VisualEffectState uint

const (
	// VisualEffectStateActive - The backdrop should always appear active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum/active
	VisualEffectStateActive VisualEffectState = 1
	// VisualEffectStateFollowsWindowActiveState - The backdrop should automatically appear active when the window is active, and inactive when it is not.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum/followsWindowActiveState
	VisualEffectStateFollowsWindowActiveState VisualEffectState = 0
	// VisualEffectStateInactive - The backdrop should always appear inactive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSVisualEffectView/State-swift.enum/inactive
	VisualEffectStateInactive VisualEffectState = 2
)

/* debug [enums.gen.go]: Processing enum NSWindowAnimationBehavior (5 cases) */
// WindowAnimationBehavior - Constants that control the automatic window animation behavior windows use when ordering to the front or out of view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum
type WindowAnimationBehavior uint

const (
	// WindowAnimationBehaviorAlertPanel - The animation behavior that’s appropriate to the alert window style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/alertPanel
	WindowAnimationBehaviorAlertPanel WindowAnimationBehavior = 5
	// WindowAnimationBehaviorDefault - The automatic animation that’s appropriate to the window type. This is the default.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/default
	WindowAnimationBehaviorDefault WindowAnimationBehavior = 0
	// WindowAnimationBehaviorDocumentWindow - The animation behavior that’s appropriate to the document window style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/documentWindow
	WindowAnimationBehaviorDocumentWindow WindowAnimationBehavior = 3
	// WindowAnimationBehaviorNone - No automatic animation used. This may be useful when you perform your own window animation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/none
	WindowAnimationBehaviorNone WindowAnimationBehavior = 2
	// WindowAnimationBehaviorUtilityWindow - The animation behavior that’s appropriate to the utility window style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/AnimationBehavior-swift.enum/utilityWindow
	WindowAnimationBehaviorUtilityWindow WindowAnimationBehavior = 4
)

/* debug [enums.gen.go]: Processing enum NSWindowBackingLocation (3 cases) */
// WindowBackingLocation - The following constants and the related data type represent a window’s possible backing locations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingLocation-swift.enum
type WindowBackingLocation uint

const (
	// WindowBackingLocationDefault - Determined by the operating system.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingLocation-swift.enum/default
	WindowBackingLocationDefault WindowBackingLocation = 0
	// WindowBackingLocationMainMemory - Physical memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingLocation-swift.enum/mainMemory
	WindowBackingLocationMainMemory WindowBackingLocation = 2
	// WindowBackingLocationVideoMemory - Video memory.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingLocation-swift.enum/videoMemory
	WindowBackingLocationVideoMemory WindowBackingLocation = 1
)

/* debug [enums.gen.go]: Processing enum NSBackingStoreType (3 cases) */
// BackingStoreType - Constants that specify how the window device buffers the drawing done in a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
type BackingStoreType uint

const (
	// BackingStoreBuffered - The window renders all drawing into a display buffer and then flushes it to the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType/buffered
	BackingStoreBuffered BackingStoreType = 2
	// BackingStoreNonretained - The window draws directly to the screen without using any buffer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType/nonretained
	BackingStoreNonretained BackingStoreType = 1
	// BackingStoreRetained - The window uses a buffer, but draws directly to the screen where possible and to the buffer for obscured portions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType/retained
	BackingStoreRetained BackingStoreType = 0
)

/* debug [enums.gen.go]: Processing enum NSWindowButton (6 cases) */
// WindowButton - Constants that provide a way to access standard title bar buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType
type WindowButton uint

const (
	// WindowCloseButton - The close button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/closeButton
	WindowCloseButton WindowButton = 0
	// WindowDocumentIconButton - The document icon button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/documentIconButton
	WindowDocumentIconButton WindowButton = 4
	// WindowDocumentVersionsButton - The document versions button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/documentVersionsButton
	WindowDocumentVersionsButton WindowButton = 5
	// WindowMiniaturizeButton - The minimize button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/miniaturizeButton
	WindowMiniaturizeButton WindowButton = 1
	// WindowToolbarButton - The toolbar button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/toolbarButton
	WindowToolbarButton WindowButton = 3
	// WindowZoomButton - The zoom button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/ButtonType/zoomButton
	WindowZoomButton WindowButton = 2
)

/* debug [enums.gen.go]: Processing enum NSWindowCollectionBehavior (16 cases) */
// WindowCollectionBehavior - Window collection behaviors related to Mission Control, Spaces, and Stage Manager.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct
type WindowCollectionBehavior uint

const (
	// WindowCollectionBehaviorAuxiliary - The behavior marking this window as auxiliary for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/auxiliary
	WindowCollectionBehaviorAuxiliary WindowCollectionBehavior = 14
	// WindowCollectionBehaviorCanJoinAllApplications - The behavior marking this window as one that can join all apps for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/canJoinAllApplications
	WindowCollectionBehaviorCanJoinAllApplications WindowCollectionBehavior = 15
	// WindowCollectionBehaviorCanJoinAllSpaces - The window can appear in all spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/canJoinAllSpaces
	WindowCollectionBehaviorCanJoinAllSpaces WindowCollectionBehavior = 1
	// WindowCollectionBehaviorFullScreenAllowsTiling - The window can be a secondary full screen tile even if it can’t be a full screen window itself.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenAllowsTiling
	WindowCollectionBehaviorFullScreenAllowsTiling WindowCollectionBehavior = 11
	// WindowCollectionBehaviorFullScreenAuxiliary - The window displays on the same space as the full screen window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenAuxiliary
	WindowCollectionBehaviorFullScreenAuxiliary WindowCollectionBehavior = 9
	// WindowCollectionBehaviorFullScreenDisallowsTiling - The window doesn’t support being a full-screen tile window, but may support being a full-screen window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenDisallowsTiling
	WindowCollectionBehaviorFullScreenDisallowsTiling WindowCollectionBehavior = 12
	// WindowCollectionBehaviorFullScreenNone - The window doesn’t support full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenNone
	WindowCollectionBehaviorFullScreenNone WindowCollectionBehavior = 10
	// WindowCollectionBehaviorFullScreenPrimary - The window can enter full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/fullScreenPrimary
	WindowCollectionBehaviorFullScreenPrimary WindowCollectionBehavior = 8
	// WindowCollectionBehaviorIgnoresCycle - The window isn’t part of the window cycle for use with the Cycle Through Windows menu item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/ignoresCycle
	WindowCollectionBehaviorIgnoresCycle WindowCollectionBehavior = 7
	// WindowCollectionBehaviorManaged - The window participates in Mission Control and Spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/managed
	WindowCollectionBehaviorManaged WindowCollectionBehavior = 3
	// WindowCollectionBehaviorMoveToActiveSpace - When the window becomes active, move it to the active space instead of switching spaces.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/moveToActiveSpace
	WindowCollectionBehaviorMoveToActiveSpace WindowCollectionBehavior = 2
	// WindowCollectionBehaviorParticipatesInCycle - The window participates in the window cycle for use with the Cycle Through Windows menu item.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/participatesInCycle
	WindowCollectionBehaviorParticipatesInCycle WindowCollectionBehavior = 6
	// WindowCollectionBehaviorPrimary - The behavior marking this window as primary for both Stage Manager and full screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/primary
	WindowCollectionBehaviorPrimary WindowCollectionBehavior = 13
	// WindowCollectionBehaviorStationary - Mission Control doesn’t affect the window, so it stays visible and stationary, like the desktop window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/stationary
	WindowCollectionBehaviorStationary WindowCollectionBehavior = 5
	// WindowCollectionBehaviorTransient - The window floats in Spaces and hides in Mission Control.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/CollectionBehavior-swift.struct/transient
	WindowCollectionBehaviorTransient WindowCollectionBehavior = 4
	// WindowCollectionBehaviorDefault - The window appears in only one space at a time.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindowCollectionBehavior/NSWindowCollectionBehaviorDefault
	WindowCollectionBehaviorDefault WindowCollectionBehavior = 0
)

/* debug [enums.gen.go]: Processing enum NSWindowDepth (3 cases) */
// WindowDepth - A type that represents the depth, or amount of memory, for a single pixel in a window or screen.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth
type WindowDepth uint

const (
	// WindowDepthOnehundredtwentyeightBitRGB - One hundred and twenty eight bit RGB depth limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/onehundredtwentyeightBitRGB
	WindowDepthOnehundredtwentyeightBitRGB WindowDepth = 544
	// WindowDepthSixtyfourBitRGB - Sixty four bit RGB depth limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/sixtyfourBitRGB
	WindowDepthSixtyfourBitRGB WindowDepth = 528
	// WindowDepthTwentyfourBitRGB - Twenty four bit RGB depth limit.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Depth/twentyfourBitRGB
	WindowDepthTwentyfourBitRGB WindowDepth = 520
)

/* debug [enums.gen.go]: Processing enum NSWindowNumberListOptions (2 cases) */
// WindowNumberListOptions - Options to use when retrieving window numbers from the system.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions
type WindowNumberListOptions uint

const (
	// WindowNumberListAllApplications - The window numbers of windows visible on any space and belonging to any application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions/allApplications
	WindowNumberListAllApplications WindowNumberListOptions = 1
	// WindowNumberListAllSpaces - The window numbers of windows visible on any space and belonging to the calling application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/NumberListOptions/allSpaces
	WindowNumberListAllSpaces WindowNumberListOptions = 16
)

/* debug [enums.gen.go]: Processing enum NSWindowOcclusionState (1 cases) */
// WindowOcclusionState - Specifies whether the window is occluded.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct
type WindowOcclusionState uint

const (
	// WindowOcclusionStateVisible - If set, at least part of the window is visible; if not set, the entire window is occluded. A window that has a nonrectangular shape can be entirely occluded onscreen, but if its bounding box falls into a visible region, the window is considered to be visible. Note that a completely transparent window may also be considered visible.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OcclusionState-swift.struct/visible
	WindowOcclusionStateVisible WindowOcclusionState = 2
)

/* debug [enums.gen.go]: Processing enum NSWindowOrderingMode (3 cases) */
// WindowOrderingMode - Constants that let you specify how a window is ordered relative to another window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode
type WindowOrderingMode int

const (
	// WindowAbove - Moves the window above the indicated window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/above
	WindowAbove WindowOrderingMode = 1
	// WindowBelow - Moves the window below the indicated window.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/below
	WindowBelow WindowOrderingMode = -1
	// WindowOut - Moves the window off the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/OrderingMode/out
	WindowOut WindowOrderingMode = 0
)

/* debug [enums.gen.go]: Processing enum NSSelectionDirection (3 cases) */
// SelectionDirection - Constants that specify the direction a window is currently using to change the key view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection
type SelectionDirection uint

const (
	// DirectSelection - The window isn’t traversing the key view loop.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection/directSelection
	DirectSelection SelectionDirection = 0
	// SelectingNext - The window is proceeding to the next valid key view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection/selectingNext
	SelectingNext SelectionDirection = 1
	// SelectingPrevious - The window is proceeding to the previous valid key view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SelectionDirection/selectingPrevious
	SelectingPrevious SelectionDirection = 2
)

/* debug [enums.gen.go]: Processing enum NSWindowSharingType (2 cases) */
// WindowSharingType - Constants that represent the access levels other processes can have to a window’s content.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum
type WindowSharingType uint

const (
	// WindowSharingNone - A legacy constant that macOS no longer uses.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum/none
	WindowSharingNone WindowSharingType = 0
	// WindowSharingReadOnly - The window’s contents can be read but not modified by another process.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/SharingType-swift.enum/readOnly
	WindowSharingReadOnly WindowSharingType = 1
)

/* debug [enums.gen.go]: Processing enum NSWindowStyleMask (13 cases) */
// WindowStyleMask - Constants that specify the style of a window, and that you can combine with the C bitwise OR operator.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
type WindowStyleMask uint

const (
	// WindowStyleMaskBorderless - The window displays none of the usual peripheral elements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/borderless
	WindowStyleMaskBorderless WindowStyleMask = 0
	// WindowStyleMaskClosable - The window displays a close button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/closable
	WindowStyleMaskClosable WindowStyleMask = 2
	// WindowStyleMaskDocModalWindow - The window is a document-modal panel (or  a subclass of  ).
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/docModalWindow
	WindowStyleMaskDocModalWindow WindowStyleMask = 64
	// WindowStyleMaskFullScreen - The window can appear full screen. A fullscreen window does not draw its title bar, and may have special handling for its toolbar. (This mask is automatically toggled when   is called.)
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/fullScreen
	WindowStyleMaskFullScreen WindowStyleMask = 4097
	// WindowStyleMaskFullSizeContentView - When set, the window’s   consumes the full size of the window. Although you can combine this constant with other window style masks, it is respected only for windows with a title bar. Note that using this mask opts in to layer-backing. Use the   or the   to lay out views underneath the title bar–toolbar area.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/fullSizeContentView
	WindowStyleMaskFullSizeContentView WindowStyleMask = 4098
	// WindowStyleMaskHUDWindow - The window is a HUD panel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/hudWindow
	WindowStyleMaskHUDWindow WindowStyleMask = 129
	// WindowStyleMaskMiniaturizable - The window displays a minimize button.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/miniaturizable
	WindowStyleMaskMiniaturizable WindowStyleMask = 4
	// WindowStyleMaskNonactivatingPanel - The window is a panel or a subclass of   that does not activate the owning app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/nonactivatingPanel
	WindowStyleMaskNonactivatingPanel WindowStyleMask = 128
	// WindowStyleMaskResizable - The window can be resized by the user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/resizable
	WindowStyleMaskResizable WindowStyleMask = 8
	// WindowStyleMaskTexturedBackground - The window uses a textured background that darkens when the window is key or main and lightens when it is inactive, and may have a second gradient in the section below the window content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/texturedBackground
	WindowStyleMaskTexturedBackground WindowStyleMask = 9
	// WindowStyleMaskTitled - The window displays a title bar.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/titled
	WindowStyleMaskTitled WindowStyleMask = 1
	// WindowStyleMaskUnifiedTitleAndToolbar - This constant has no effect, because all windows that include a toolbar use the unified style.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/unifiedTitleAndToolbar
	WindowStyleMaskUnifiedTitleAndToolbar WindowStyleMask = 4096
	// WindowStyleMaskUtilityWindow - The window is a panel or a subclass of  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct/utilityWindow
	WindowStyleMaskUtilityWindow WindowStyleMask = 16
)

/* debug [enums.gen.go]: Processing enum NSWindowTabbingMode (3 cases) */
// WindowTabbingMode - The preferred tabbing behavior of a window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum
type WindowTabbingMode uint

const (
	// WindowTabbingModeAutomatic - A window that automatically tabs together based on the user’s tabbing preference.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum/automatic
	WindowTabbingModeAutomatic WindowTabbingMode = 0
	// WindowTabbingModeDisallowed - A window that explicitly does not prefer to tab together with other windows.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum/disallowed
	WindowTabbingModeDisallowed WindowTabbingMode = 2
	// WindowTabbingModePreferred - A window that explicitly prefers to tab together with other windows with the same tabbing identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingMode-swift.enum/preferred
	WindowTabbingModePreferred WindowTabbingMode = 1
)

/* debug [enums.gen.go]: Processing enum NSWindowTitleVisibility (2 cases) */
// WindowTitleVisibility - Specifies the appearance of the window’s title bar area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum
type WindowTitleVisibility uint

const (
	// WindowTitleHidden - The window hides the title and moves the toolbar up into the area previously occupied by the title.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum/hidden
	WindowTitleHidden WindowTitleVisibility = 1
	// WindowTitleVisible - The window has the regular window title and title bar buttons.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TitleVisibility-swift.enum/visible
	WindowTitleVisible WindowTitleVisibility = 0
)

/* debug [enums.gen.go]: Processing enum NSWindowUserTabbingPreference (3 cases) */
// WindowUserTabbingPreference - A value that indicates the user’s preference for window tabbing.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum
type WindowUserTabbingPreference uint

const (
	// WindowUserTabbingPreferenceAlways - A value that indicates a window should always display as tabs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum/always
	WindowUserTabbingPreferenceAlways WindowUserTabbingPreference = 1
	// WindowUserTabbingPreferenceInFullScreen - A value that indicates a window should only display as tabs when in full-screen mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum/inFullScreen
	WindowUserTabbingPreferenceInFullScreen WindowUserTabbingPreference = 2
	// WindowUserTabbingPreferenceManual - A value that indicates a window should display as tabs according to the window’s tabbing mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/UserTabbingPreference-swift.enum/manual
	WindowUserTabbingPreferenceManual WindowUserTabbingPreference = 0
)

/* debug [enums.gen.go]: Processing enum NSWorkspaceAuthorizationType (3 cases) */
// WorkspaceAuthorizationType - The types of privileged file operations that can be authorized by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType
type WorkspaceAuthorizationType uint

const (
	// WorkspaceAuthorizationTypeCreateSymbolicLink - Authorization for the app to create a symbolic link.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType/createSymbolicLink
	WorkspaceAuthorizationTypeCreateSymbolicLink WorkspaceAuthorizationType = 0
	// WorkspaceAuthorizationTypeReplaceFile - Authorization for the app to perform an atomic file write without changing the target file’s permissions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType/replaceFile
	WorkspaceAuthorizationTypeReplaceFile WorkspaceAuthorizationType = 2
	// WorkspaceAuthorizationTypeSetAttributes - Authorization for the app to change specific file attributes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/AuthorizationType/setAttributes
	WorkspaceAuthorizationTypeSetAttributes WorkspaceAuthorizationType = 1
)

/* debug [enums.gen.go]: Processing enum NSWorkspaceIconCreationOptions (2 cases) */
// WorkspaceIconCreationOptions - Constants that describe options for creating icons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions
type WorkspaceIconCreationOptions uint

const (
	// Exclude10_4ElementsIconCreationOption - An option to suppress generation of the new higher resolution icon representations that are supported in macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions/exclude10_4ElementsIconCreationOption
	Exclude10_4ElementsIconCreationOption WorkspaceIconCreationOptions = 4
	// ExcludeQuickDrawElementsIconCreationOption - An option to suppress generation of the QuickDraw format icon representations that are used in macOS 10.0 through macOS 10.4.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/IconCreationOptions/excludeQuickDrawElementsIconCreationOption
	ExcludeQuickDrawElementsIconCreationOption WorkspaceIconCreationOptions = 2
)

/* debug [enums.gen.go]: Processing enum NSWorkspaceLaunchOptions (12 cases) */
// WorkspaceLaunchOptions - Constants specifying how you want to launch an app
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions
type WorkspaceLaunchOptions uint

const (
	// WorkspaceLaunchAllowingClassicStartup - Start up the Classic compatibility environment, if it is required by the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/allowingClassicStartup
	WorkspaceLaunchAllowingClassicStartup WorkspaceLaunchOptions = 10
	// WorkspaceLaunchAndHide - Tell the app to hide itself as soon as it finishes launching.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/andHide
	WorkspaceLaunchAndHide WorkspaceLaunchOptions = 7
	// WorkspaceLaunchAndHideOthers - Hide all apps except the newly launched one.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/andHideOthers
	WorkspaceLaunchAndHideOthers WorkspaceLaunchOptions = 8
	// WorkspaceLaunchAndPrint - Print items instead of opening them.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/andPrint
	WorkspaceLaunchAndPrint WorkspaceLaunchOptions = 0
	// WorkspaceLaunchAsync - Launch the app and return the results asynchronously.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/async
	WorkspaceLaunchAsync WorkspaceLaunchOptions = 5
	// WorkspaceLaunchDefault - Launch the app asynchronously and launch it in the Classic environment, if required.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/default
	WorkspaceLaunchDefault WorkspaceLaunchOptions = 9
	// WorkspaceLaunchInhibitingBackgroundOnly - Causes launch to fail if the target is background-only.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/inhibitingBackgroundOnly
	WorkspaceLaunchInhibitingBackgroundOnly WorkspaceLaunchOptions = 2
	// WorkspaceLaunchNewInstance - Create a new instance of the app, even if one is already running.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/newInstance
	WorkspaceLaunchNewInstance WorkspaceLaunchOptions = 6
	// WorkspaceLaunchPreferringClassic - Force the app to launch in the Classic compatibility environment.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/preferringClassic
	WorkspaceLaunchPreferringClassic WorkspaceLaunchOptions = 11
	// WorkspaceLaunchWithErrorPresentation - Display an error panel to the user if a failure occurs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/withErrorPresentation
	WorkspaceLaunchWithErrorPresentation WorkspaceLaunchOptions = 1
	// WorkspaceLaunchWithoutActivation - Launch the app but do not bring it into the foreground.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/withoutActivation
	WorkspaceLaunchWithoutActivation WorkspaceLaunchOptions = 4
	// WorkspaceLaunchWithoutAddingToRecents - Do not add the app or documents to the Recents menu.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchOptions/withoutAddingToRecents
	WorkspaceLaunchWithoutAddingToRecents WorkspaceLaunchOptions = 3
)

/* debug [enums.gen.go]: Processing enum NSWritingDirection (3 cases) */
// WritingDirection - Constants that specify the writing direction.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection
type WritingDirection int

const (
	// WritingDirectionLeftToRight - The writing direction is left to right.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/leftToRight
	WritingDirectionLeftToRight WritingDirection = 0
	// WritingDirectionNatural - The writing direction of the current script that the system determines using the Unicode Bidi Algorithm rules P2 and P3.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/natural
	WritingDirectionNatural WritingDirection = -1
	// WritingDirectionRightToLeft - The writing direction is right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirection/rightToLeft
	WritingDirectionRightToLeft WritingDirection = 1
)

/* debug [enums.gen.go]: Processing enum NSWritingDirectionFormatType (2 cases) */
// WritingDirectionFormatType - Constants for the writing direction attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirectionFormatType
type WritingDirectionFormatType uint

const (
	// WritingDirectionEmbedding - Text is embedded in text with another writing direction. For example, an English quotation in the middle of an Arabic sentence could be marked as being embedded left-to-right text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirectionFormatType/embedding
	WritingDirectionEmbedding WritingDirectionFormatType = 0
	// WritingDirectionOverride - Enables character types with inherent directionality to be overridden when required for special cases, such as for part numbers made of mixed English, digits, and Hebrew letters to be written from right to left.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingDirectionFormatType/override
	WritingDirectionOverride WritingDirectionFormatType = 2
)

/* debug [enums.gen.go]: Processing enum NSWritingToolsBehavior (4 cases) */
// WritingToolsBehavior - Constants that specify the Writing Tools experience for the underlying view.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior
type WritingToolsBehavior int

const (
	// WritingToolsBehaviorComplete - An option to provide the complete Writing Tools experience for the text view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/complete
	WritingToolsBehaviorComplete WritingToolsBehavior = 1
	// WritingToolsBehaviorDefault - An option to let the system determine the best way to enable Writing Tools for the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/default
	WritingToolsBehaviorDefault WritingToolsBehavior = 0
	// WritingToolsBehaviorLimited - An option to provide a limited, overlay-panel experience for the text view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/limited
	WritingToolsBehaviorLimited WritingToolsBehavior = 2
	// WritingToolsBehaviorNone - An option to prevent Writing Tools from modifying the text in the view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsBehavior/none
	WritingToolsBehaviorNone WritingToolsBehavior = -1
)

/* debug [enums.gen.go]: Processing enum NSWritingToolsCoordinatorContextScope (3 cases) */
// WritingToolsCoordinatorContextScope - Options that indicate how much of your content Writing Tools requested.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/ContextScope
type WritingToolsCoordinatorContextScope uint

const (
	// WritingToolsCoordinatorContextScopeFullDocument - An option to provide all of your view’s text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/ContextScope/fullDocument
	WritingToolsCoordinatorContextScopeFullDocument WritingToolsCoordinatorContextScope = 1
	// WritingToolsCoordinatorContextScopeUserSelection - An option to provide only the view’s currently selected text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/ContextScope/userSelection
	WritingToolsCoordinatorContextScopeUserSelection WritingToolsCoordinatorContextScope = 0
	// WritingToolsCoordinatorContextScopeVisibleArea - An option to provide only the text in the currently visible portion   of your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/ContextScope/visibleArea
	WritingToolsCoordinatorContextScopeVisibleArea WritingToolsCoordinatorContextScope = 2
)

/* debug [enums.gen.go]: Processing enum NSWritingToolsCoordinatorState (4 cases) */
// WritingToolsCoordinatorState - The states that indicate the current activity, if any, Writing Tools
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum
type WritingToolsCoordinatorState uint

const (
	// WritingToolsCoordinatorStateInactive - A state that indicates Writing Tools isn’t currently performing   any work on your view’s content.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/inactive
	WritingToolsCoordinatorStateInactive WritingToolsCoordinatorState = 0
	// WritingToolsCoordinatorStateInteractiveResting - A state that indicates Writing Tools is in the resting state   for an inline editing experience.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/interactiveResting
	WritingToolsCoordinatorStateInteractiveResting WritingToolsCoordinatorState = 2
	// WritingToolsCoordinatorStateInteractiveStreaming - A state that indicates Writing Tools is processing a request and   incorporating changes interactively into your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/interactiveStreaming
	WritingToolsCoordinatorStateInteractiveStreaming WritingToolsCoordinatorState = 3
	// WritingToolsCoordinatorStateNoninteractive - A state that indicates Writing Tools is handling interactions in   the system UI, instead of in your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/State-swift.enum/noninteractive
	WritingToolsCoordinatorStateNoninteractive WritingToolsCoordinatorState = 1
)

/* debug [enums.gen.go]: Processing enum NSWritingToolsCoordinatorTextAnimation (5 cases) */
// WritingToolsCoordinatorTextAnimation - The types of animations that Writing Tools performs during an
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextAnimation
type WritingToolsCoordinatorTextAnimation uint

const (
	// WritingToolsCoordinatorTextAnimationAnticipate - The animation that Writing Tools performs when waiting to receive   results from the large language model.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextAnimation/anticipate
	WritingToolsCoordinatorTextAnimationAnticipate WritingToolsCoordinatorTextAnimation = 0
	// WritingToolsCoordinatorTextAnimationAnticipateInactive - The animation effect that Writing Tools performs when the view is waiting   for results, but the system isn’t actively evaluating the text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextAnimation/anticipateInactive
	WritingToolsCoordinatorTextAnimationAnticipateInactive WritingToolsCoordinatorTextAnimation = 8
	// WritingToolsCoordinatorTextAnimationInsert - The animation that Writing Tools performs when inserting text into your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextAnimation/insert
	WritingToolsCoordinatorTextAnimationInsert WritingToolsCoordinatorTextAnimation = 2
	// WritingToolsCoordinatorTextAnimationRemove - The animation that Writing Tools performs when removing text from your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextAnimation/remove
	WritingToolsCoordinatorTextAnimationRemove WritingToolsCoordinatorTextAnimation = 1
	// WritingToolsCoordinatorTextAnimationTranslate - The animation effect that Writing Tools performs on text situated after   the insertion point.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextAnimation/translate
	WritingToolsCoordinatorTextAnimationTranslate WritingToolsCoordinatorTextAnimation = 9
)

/* debug [enums.gen.go]: Processing enum NSWritingToolsCoordinatorTextReplacementReason (2 cases) */
// WritingToolsCoordinatorTextReplacementReason - Options that indicate whether Writing Tools is animating changes to
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextReplacementReason
type WritingToolsCoordinatorTextReplacementReason uint

const (
	// WritingToolsCoordinatorTextReplacementReasonInteractive - An option to animate the replacement of text in your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextReplacementReason/interactive
	WritingToolsCoordinatorTextReplacementReasonInteractive WritingToolsCoordinatorTextReplacementReason = 0
	// WritingToolsCoordinatorTextReplacementReasonNoninteractive - An option to replace the text in your view without animating the change.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextReplacementReason/noninteractive
	WritingToolsCoordinatorTextReplacementReasonNoninteractive WritingToolsCoordinatorTextReplacementReason = 1
)

/* debug [enums.gen.go]: Processing enum NSWritingToolsCoordinatorTextUpdateReason (2 cases) */
// WritingToolsCoordinatorTextUpdateReason - Constants that specify the reason you updated your view’s content
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextUpdateReason
type WritingToolsCoordinatorTextUpdateReason uint

const (
	// WritingToolsCoordinatorTextUpdateReasonTyping - An operation that involved a person editing the text in your view.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextUpdateReason/typing
	WritingToolsCoordinatorTextUpdateReasonTyping WritingToolsCoordinatorTextUpdateReason = 0
	// WritingToolsCoordinatorTextUpdateReasonUndoRedo - An operation that changed the view’s text as part of an undo or   redo command.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsCoordinator/TextUpdateReason/undoRedo
	WritingToolsCoordinatorTextUpdateReasonUndoRedo WritingToolsCoordinatorTextUpdateReason = 1
)

/* debug [enums.gen.go]: Processing enum NSWritingToolsResultOptions (6 cases) */
// WritingToolsResultOptions - Constants to specify what type of content to allow in Writing Tools
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions
type WritingToolsResultOptions uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions/NSWritingToolsResultDefault
	WritingToolsResultDefault WritingToolsResultOptions = 0
	// WritingToolsResultList - An option to allow list-style formatting in the returned text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions/list
	WritingToolsResultList WritingToolsResultOptions = 4
	// WritingToolsResultPlainText - An option to allow only plain text without any attributes in   the returned text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions/plainText
	WritingToolsResultPlainText WritingToolsResultOptions = 1
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions/presentationIntent
	WritingToolsResultPresentationIntent WritingToolsResultOptions = 9
	// WritingToolsResultRichText - An option to include style attributes consistent with the RTF   format in the returned text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions/richText
	WritingToolsResultRichText WritingToolsResultOptions = 2
	// WritingToolsResultTable - An option to allow tabular layout attributes in the returned text.
	//
	// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWritingToolsResultOptions/table
	WritingToolsResultTable WritingToolsResultOptions = 8
)


