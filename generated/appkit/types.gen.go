// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit
import (
	"unsafe"
)


// TimeInterval for non-CoreGraphics frameworks
type TimeInterval = float64  // NSTimeInterval
// AppKit-specific type aliases
// WindowLevel represents a window's position in the z-axis
type WindowLevel = int  // NSWindowLevel


// Fallback type aliases for undefined types
// These types are referenced in method signatures but not fully documented.
// Using unsafe.Pointer as fallback to allow code generation.
type BOOL unsafe.Pointer

type CADisplayLink unsafe.Pointer

type CALayer unsafe.Pointer

type CGColorSpaceRef unsafe.Pointer

type CGContextRef unsafe.Pointer

type CGDirectDisplayID unsafe.Pointer

type CGEventRef unsafe.Pointer

type CGFloat unsafe.Pointer

type CGImageRef unsafe.Pointer

type CGLContextObj unsafe.Pointer

type CGLPixelFormatObj unsafe.Pointer

type CGPoint unsafe.Pointer

type CGRect unsafe.Pointer

type CIContext unsafe.Pointer

type CIFilter unsafe.Pointer

type CIImage unsafe.Pointer

type Class unsafe.Pointer

type GLbitfield unsafe.Pointer

type GLint unsafe.Pointer

type NSAccessibilityActionName unsafe.Pointer

type NSAccessibilityAttributeName unsafe.Pointer

type NSAccessibilityNotificationName unsafe.Pointer

type NSAccessibilityRole unsafe.Pointer

type NSAffineTransform unsafe.Pointer

type NSAlignmentOptions unsafe.Pointer

type NSAnimationEffect unsafe.Pointer

type NSAnimationProgress unsafe.Pointer

type NSAppearanceName unsafe.Pointer

type NSApplicationActivationPolicy unsafe.Pointer

type NSApplicationDelegateReply unsafe.Pointer

type NSArray unsafe.Pointer

type NSAttributedString unsafe.Pointer

type NSAttributedStringKey unsafe.Pointer

type NSAutoresizingMaskOptions unsafe.Pointer

type NSBackgroundStyle unsafe.Pointer

type NSBackingStoreType unsafe.Pointer

type NSBorderType unsafe.Pointer

type NSBoxType unsafe.Pointer

type NSBundle unsafe.Pointer

type NSCellAttribute unsafe.Pointer

type NSCellHitResult unsafe.Pointer

type NSCellType unsafe.Pointer

type NSCharacterSet unsafe.Pointer

type NSCloseCommand unsafe.Pointer

type NSCoder unsafe.Pointer

type NSCollectionViewLayoutAttributes unsafe.Pointer

type NSCollectionViewScrollPosition unsafe.Pointer

type NSCollectionViewSupplementaryElementKind unsafe.Pointer

type NSColorListName unsafe.Pointer

type NSColorName unsafe.Pointer

type NSColorPanelMode unsafe.Pointer

type NSColorPanelOptions unsafe.Pointer

type NSColorRenderingIntent unsafe.Pointer

type NSColorSpaceModel unsafe.Pointer

type NSColorSpaceName unsafe.Pointer

type NSColorSystemEffect unsafe.Pointer

type NSColorType unsafe.Pointer

type NSColorWellStyle unsafe.Pointer

type NSComparisonResult unsafe.Pointer

type NSCompositingOperation unsafe.Pointer

type NSControlBorderShape unsafe.Pointer

type NSControlSize unsafe.Pointer

type NSControlStateValue unsafe.Pointer

type NSControlTint unsafe.Pointer

type NSData unsafe.Pointer

type NSDate unsafe.Pointer

type NSDictionary unsafe.Pointer

type NSDisplayGamut unsafe.Pointer

type NSDocumentChangeType unsafe.Pointer

type NSDragOperation unsafe.Pointer

type NSEdgeInsets unsafe.Pointer

type NSError unsafe.Pointer

type NSEventGestureAxis unsafe.Pointer

type NSEventMask unsafe.Pointer

type NSEventModifierFlags unsafe.Pointer

type NSException unsafe.Pointer

type NSFileVersion unsafe.Pointer

type NSFileWrapper unsafe.Pointer

type NSFocusRingType unsafe.Pointer

type NSFontDescriptorAttributeName unsafe.Pointer

type NSFontDescriptorSymbolicTraits unsafe.Pointer

type NSFontTextStyle unsafe.Pointer

type NSFontWeight unsafe.Pointer

type NSFormatter unsafe.Pointer

type NSGlassEffectViewStyle unsafe.Pointer

type NSGlyph unsafe.Pointer

type NSGradientDrawingOptions unsafe.Pointer

type NSGradientType unsafe.Pointer

type NSGridCellPlacement unsafe.Pointer

type NSGridRow unsafe.Pointer

type NSGridRowAlignment unsafe.Pointer

type NSHelpAnchorName unsafe.Pointer

type NSHelpBookName unsafe.Pointer

type NSImageAlignment unsafe.Pointer

type NSImageName unsafe.Pointer

type NSImageSymbolColorRenderingMode unsafe.Pointer

type NSImageSymbolScale unsafe.Pointer

type NSImageSymbolVariableValueMode unsafe.Pointer

type NSIndexPath unsafe.Pointer

type NSIndexSet unsafe.Pointer

type NSInteger unsafe.Pointer

type NSInterfaceStyle unsafe.Pointer

type NSItemProvider unsafe.Pointer

type NSLayoutAttribute unsafe.Pointer

type NSLayoutConstraintOrientation unsafe.Pointer

type NSLayoutPriority unsafe.Pointer

type NSLineBreakMode unsafe.Pointer

type NSLineBreakStrategy unsafe.Pointer

type NSMenuItemBadgeType unsafe.Pointer

type NSModalResponse unsafe.Pointer

type NSModalSession unsafe.Pointer

type NSMultibyteGlyphPacking unsafe.Pointer

type NSMutableData unsafe.Pointer

type NSMutableDictionary unsafe.Pointer

type NSNibName unsafe.Pointer

type NSNotification unsafe.Pointer

type NSNotificationCenter unsafe.Pointer

type NSNumber unsafe.Pointer

type NSObject unsafe.Pointer

type NSOpenGLGlobalOption unsafe.Pointer

type NSOperationQueue unsafe.Pointer

type NSOrthography unsafe.Pointer

type NSPaperOrientation unsafe.Pointer

type NSPasteboardAccessBehavior unsafe.Pointer

type NSPasteboardType unsafe.Pointer

type NSPickerTouchBarItemSelectionMode unsafe.Pointer

type NSPoint unsafe.Pointer

type NSPointPointer unsafe.Pointer

type NSPopoverBehavior unsafe.Pointer

type NSPredicateEditorRowTemplate unsafe.Pointer

type NSPrintJobDispositionValue unsafe.Pointer

type NSPrintPanelJobStyleHint unsafe.Pointer

type NSPrinterPaperName unsafe.Pointer

type NSPrinterTableStatus unsafe.Pointer

type NSPrinterTypeName unsafe.Pointer

type NSProgress unsafe.Pointer

type NSProgressIndicatorStyle unsafe.Pointer

type NSRange unsafe.Pointer

type NSRangePointer unsafe.Pointer

type NSRect unsafe.Pointer

type NSRectEdge unsafe.Pointer

type NSRemoteNotificationType unsafe.Pointer

type NSRulerOrientation unsafe.Pointer

type NSRulerViewUnitName unsafe.Pointer

type NSRunLoopMode unsafe.Pointer

type NSSaveOperationType unsafe.Pointer

type NSScriptCommand unsafe.Pointer

type NSScriptObjectSpecifier unsafe.Pointer

type NSScrollElasticity unsafe.Pointer

type NSScrollViewFindBarPosition unsafe.Pointer

type NSScrollerKnobStyle unsafe.Pointer

type NSScrollerStyle unsafe.Pointer

type NSScrubberAlignment unsafe.Pointer

type NSScrubberMode unsafe.Pointer

type NSSearchFieldRecentsAutosaveName unsafe.Pointer

type NSSelectionDirection unsafe.Pointer

type NSServiceProviderName unsafe.Pointer

type NSSet unsafe.Pointer

type NSSharingServiceName unsafe.Pointer

type NSSize unsafe.Pointer

type NSSliderType unsafe.Pointer

type NSSplitViewAutosaveName unsafe.Pointer

type NSSplitViewDividerStyle unsafe.Pointer

type NSSplitViewItemBehavior unsafe.Pointer

type NSStackViewDistribution unsafe.Pointer

type NSStackViewGravity unsafe.Pointer

type NSStackViewVisibilityPriority unsafe.Pointer

type NSStoryboardControllerCreator unsafe.Pointer

type NSStoryboardSceneIdentifier unsafe.Pointer

type NSStoryboardSegueIdentifier unsafe.Pointer

type NSString unsafe.Pointer

type NSSymbolEffect unsafe.Pointer

type NSSymbolEffectOptions unsafe.Pointer

type NSTabPosition unsafe.Pointer

type NSTabState unsafe.Pointer

type NSTabViewBorderType unsafe.Pointer

type NSTabViewType unsafe.Pointer

type NSTableViewAnimationOptions unsafe.Pointer

type NSTableViewColumnAutoresizingStyle unsafe.Pointer

type NSTableViewDraggingDestinationFeedbackStyle unsafe.Pointer

type NSTableViewGridLineStyle unsafe.Pointer

type NSTableViewRowSizeStyle unsafe.Pointer

type NSTableViewSelectionHighlightStyle unsafe.Pointer

type NSTextAlignment unsafe.Pointer

type NSTextCheckingResult unsafe.Pointer

type NSTextCheckingTypes unsafe.Pointer

type NSTextFieldBezelStyle unsafe.Pointer

type NSTextInsertionIndicatorAutomaticModeOptions unsafe.Pointer

type NSTextInsertionIndicatorDisplayMode unsafe.Pointer

type NSTextLayoutManagerSegmentOptions unsafe.Pointer

type NSTextLayoutManagerSegmentType unsafe.Pointer

type NSTimeInterval unsafe.Pointer

type NSTitlePosition unsafe.Pointer

type NSTitlebarSeparatorStyle unsafe.Pointer

type NSTokenStyle unsafe.Pointer

type NSToolTipTag unsafe.Pointer

type NSToolbarDisplayMode unsafe.Pointer

type NSToolbarIdentifier unsafe.Pointer

type NSToolbarItemIdentifier unsafe.Pointer

type NSToolbarItemStyle unsafe.Pointer

type NSToolbarItemVisibilityPriority unsafe.Pointer

type NSToolbarSizeMode unsafe.Pointer

type NSTouchBarCustomizationIdentifier unsafe.Pointer

type NSTouchBarItemIdentifier unsafe.Pointer

type NSTouchBarItemPriority unsafe.Pointer

type NSTouchPhase unsafe.Pointer

type NSTouchTypeMask unsafe.Pointer

type NSTrackingRectTag unsafe.Pointer

type NSTypesetterBehavior unsafe.Pointer

type NSUInteger unsafe.Pointer

type NSURL unsafe.Pointer

type NSUUID unsafe.Pointer

type NSUndoManager unsafe.Pointer

type NSUserActivity unsafe.Pointer

type NSUserInterfaceItemIdentifier unsafe.Pointer

type NSUserInterfaceLayoutDirection unsafe.Pointer

type NSUserInterfaceLayoutOrientation unsafe.Pointer

type NSViewLayerContentsPlacement unsafe.Pointer

type NSViewLayerContentsRedrawPolicy unsafe.Pointer

type NSViewLayoutRegionAdaptivityAxis unsafe.Pointer

type NSVisualEffectBlendingMode unsafe.Pointer

type NSVisualEffectMaterial unsafe.Pointer

type NSVisualEffectState unsafe.Pointer

type NSWindowAnimationBehavior unsafe.Pointer

type NSWindowBackingLocation unsafe.Pointer

type NSWindowButton unsafe.Pointer

type NSWindowCollectionBehavior unsafe.Pointer

type NSWindowDepth unsafe.Pointer

type NSWindowFrameAutosaveName unsafe.Pointer

type NSWindowLevel unsafe.Pointer

type NSWindowNumberListOptions unsafe.Pointer

type NSWindowOcclusionState unsafe.Pointer

type NSWindowOrderingMode unsafe.Pointer

type NSWindowPersistableFrameDescriptor unsafe.Pointer

type NSWindowSharingType unsafe.Pointer

type NSWindowStyleMask unsafe.Pointer

type NSWindowTabbingIdentifier unsafe.Pointer

type NSWindowTabbingMode unsafe.Pointer

type NSWindowTitleVisibility unsafe.Pointer

type NSWindowToolbarStyle unsafe.Pointer

type NSWindowUserTabbingPreference unsafe.Pointer

type NSWorkspaceAuthorization unsafe.Pointer

type NSWorkspaceAuthorizationType unsafe.Pointer

type NSWorkspaceLaunchOptions unsafe.Pointer

type NSWritingDirection unsafe.Pointer

type NSWritingToolsBehavior unsafe.Pointer

type NSWritingToolsCoordinatorState unsafe.Pointer

type NSWritingToolsCoordinatorTextUpdateReason unsafe.Pointer

type NSWritingToolsResultOptions unsafe.Pointer

type QTMovie unsafe.Pointer

type SEL unsafe.Pointer

type SKNode unsafe.Pointer

type SWCollaborationMetadata unsafe.Pointer

type UIActivityItemsConfigurationReading unsafe.Pointer

type UIBarButtonItem unsafe.Pointer

type UIImage unsafe.Pointer

type UIMenuElement unsafe.Pointer

type UTType unsafe.Pointer



