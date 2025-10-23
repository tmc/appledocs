// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

// Type aliases and typedefs
// AccessibilityActionName - Constants that describe types of actions.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Action
// NSAccessibilityActionName is a string typedef
type AccessibilityActionName = string
// AccessibilityAnnotationAttributeKey - Keys for annotation attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/AnnotationAttributeKey
// NSAccessibilityAnnotationAttributeKey is a string typedef
type AccessibilityAnnotationAttributeKey = string
// AccessibilityAttributeName - Constants that describe attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute
// NSAccessibilityAttributeName is a string typedef
type AccessibilityAttributeName = string
// AccessibilityFontAttributeKey - Keys for font attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/FontAttributeKey
// NSAccessibilityFontAttributeKey is a string typedef
type AccessibilityFontAttributeKey = string
// AccessibilityNotificationName - The name of the notification.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification
// NSAccessibilityNotificationName is a string typedef
type AccessibilityNotificationName = string
// AccessibilityOrientationValue - Values that indicate the orientation of user interface elements, such as scroll bars and split views.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/OrientationValue
// NSAccessibilityOrientationValue is a string typedef
type AccessibilityOrientationValue = string
// AccessibilityParameterizedAttributeName - Values that describe parameterized attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/ParameterizedAttribute
// NSAccessibilityParameterizedAttributeName is a string typedef
type AccessibilityParameterizedAttributeName = string
// AccessibilityRulerMarkerTypeValue - Values that describe ruler marker types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerMarkerTypeValue
// NSAccessibilityRulerMarkerTypeValue is a string typedef
type AccessibilityRulerMarkerTypeValue = string
// AccessibilityRulerUnitValue - Values that indicate the unit values of a ruler or layout area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/RulerUnitValue
// NSAccessibilityRulerUnitValue is a string typedef
type AccessibilityRulerUnitValue = string
// AccessibilitySortDirectionValue - Values that indicate the sort direction of a column.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/SortDirectionValue
// NSAccessibilitySortDirectionValue is a string typedef
type AccessibilitySortDirectionValue = string
// AccessibilityLoadingToken - A token type for loading accessibility elements.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAccessibilityLoadingToken
// NSAccessibilityLoadingToken has base type: id<NSObject,NSSecureCoding>
type AccessibilityLoadingToken uintptr
// AnimatablePropertyKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyKey
// NSAnimatablePropertyKey is a string typedef
type AnimatablePropertyKey = string
// AnimationProgress - The animation progress, as a floating-point number between   and  .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/Progress
// NSAnimationProgress has base type: float
type AnimationProgress uintptr
// AppKitVersion - Constants for determining which version of AppKit is available.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAppKitVersion
// NSAppKitVersion has base type: double
type AppKitVersion uintptr
// ModalResponse - A set of button return values for modal dialogs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse
type ModalResponse int
// ModalSession - Variables of type   point to information used by the system between  ’s   and   messages.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSApplication/ModalSession
// NSModalSession has base type: struct _NSModalSession *
type ModalSession uintptr
// CellStateValue - Constants for specifying a cell’s state and are used mostly for buttons.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCell/StateValue
// NSCellStateValue has base type: NSControlStateValue
type CellStateValue uintptr
// CollectionViewDecorationElementKind type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/DecorationElementKind
// NSCollectionViewDecorationElementKind is a string typedef
type CollectionViewDecorationElementKind = string
// CollectionViewSupplementaryElementKind type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionView/SupplementaryElementKind
// NSCollectionViewSupplementaryElementKind is a string typedef
type CollectionViewSupplementaryElementKind = string
// ColorName - The name of a color.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColor/Name
// NSColorName is a string typedef
type ColorName = string
// ColorListName - The name assigned to a color list.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorList/Name-swift.typealias
// NSColorListName is a string typedef
type ColorListName = string
// ColorSpaceName - Constants that specify color space names.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName
// NSColorSpaceName is a string typedef
type ColorSpaceName = string
// DeviceDescriptionKey - These constants are the keys for device description dictionaries.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSDeviceDescriptionKey
// NSDeviceDescriptionKey is a string typedef
type DeviceDescriptionKey = string
// FontTextStyle - Constants that specify the preferred text styles you use with fonts.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/TextStyle
// NSFontTextStyle is a string typedef
type FontTextStyle = string
// FontTextStyleOptionKey - The options that you apply when requesting the font or font descriptor of a preferred text style.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/TextStyleOptionKey
// NSFontTextStyleOptionKey is a string typedef
type FontTextStyleOptionKey = string
// FontWidth type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFont/Width
// NSFontWidth has base type: CGFloat
type FontWidth uintptr
// FontDescriptorAttributeName - Constants for the names of font attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/AttributeName
// NSFontDescriptorAttributeName is a string typedef
type FontDescriptorAttributeName = string
// FontDescriptorFeatureKey - Constants to use as keys to retrieve information about a font descriptor from its feature dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/FeatureKey
// NSFontDescriptorFeatureKey is a string typedef
type FontDescriptorFeatureKey = string
// FontDescriptorSystemDesign - Constants for font designs, such as monospace, rounded, and serif.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/SystemDesign
// NSFontDescriptorSystemDesign is a string typedef
type FontDescriptorSystemDesign = string
// FontDescriptorTraitKey - Constants that can be used as keys to retrieve information about a font descriptor from its trait dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/TraitKey
// NSFontDescriptorTraitKey is a string typedef
type FontDescriptorTraitKey = string
// FontDescriptorVariationKey - Constants that can be used as keys to retrieve information about a font descriptor from its variation axis dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontDescriptor/VariationKey
// NSFontDescriptorVariationKey is a string typedef
type FontDescriptorVariationKey = string
// FontFamilyClass - Constants that classify certain stylistic qualities of the font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontFamilyClass
// NSFontFamilyClass has base type: uint32_t
type FontFamilyClass uintptr
// FontSymbolicTraits - A symbolic description of stylistic aspects of a font.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSFontSymbolicTraits
// NSFontSymbolicTraits has base type: uint32_t
type FontSymbolicTraits uintptr
// Glyph - The type used to specify glyphs.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGlyph
type Glyph uint32
// GraphicsContextAttributeKey - Constants that specify the dictionary keys for the attributes of the graphics context.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/AttributeKey
// NSGraphicsContextAttributeKey is a string typedef
type GraphicsContextAttributeKey = string
// GraphicsContextRepresentationFormatName - Constants that specify values for the representation format name key in a graphic context’s attributes dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/RepresentationFormatName
// NSGraphicsContextRepresentationFormatName is a string typedef
type GraphicsContextRepresentationFormatName = string
// HelpAnchorName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/AnchorName
// NSHelpAnchorName is a string typedef
type HelpAnchorName = string
// HelpBookName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/BookName
// NSHelpBookName is a string typedef
type HelpBookName = string
// HelpManagerContextHelpKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSHelpManager/ContextHelpKey
// NSHelpManagerContextHelpKey is a string typedef
type HelpManagerContextHelpKey = string
// ImageName - Named images, defined by the system or you, for use in your app.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImage/Name-swift.typealias
// NSImageName is a string typedef
type ImageName = string
// ImageHintKey - Constants for the keys to include in a hints dictionary when drawing the image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSImageRep/HintKey
// NSImageHintKey is a string typedef
type ImageHintKey = string
// LayoutPriority - Layout priority used to indicate the relative importance of constraints, allowing Auto Layout to make appropriate tradeoffs when satisfying the constraints of the system as a whole.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutConstraint/Priority-swift.struct
// NSLayoutPriority has base type: float
type LayoutPriority uintptr
// NibName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib/Name
// NSNibName is a string typedef
type NibName = string
// OpenGLPixelFormatAttribute - Pixel format attributes for OpenGL.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSOpenGLPixelFormatAttribute
// NSOpenGLPixelFormatAttribute has base type: uint32_t
type OpenGLPixelFormatAttribute uintptr
// PasteboardName - Constants that represent the standard pasteboard names.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/Name-swift.struct
// NSPasteboardName is a string typedef
type PasteboardName = string
// PasteboardType - The supported pasteboard types.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType
// NSPasteboardType is a string typedef
type PasteboardType = string
// PasteboardTypeFindPanelSearchOptionKey - Search options for the find panel.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/FindPanelSearchOptionKey
// NSPasteboardTypeFindPanelSearchOptionKey is a string typedef
type PasteboardTypeFindPanelSearchOptionKey = string
// PasteboardTypeTextFinderOptionKey - Search options for text in Finder.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/PasteboardType/TextFinderOptionKey
// NSPasteboardTypeTextFinderOptionKey is a string typedef
type PasteboardTypeTextFinderOptionKey = string
// PasteboardReadingOptionKey - Options for reading pasteboard data.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptionKey
// NSPasteboardReadingOptionKey is a string typedef
type PasteboardReadingOptionKey = string
// PrintInfoAttributeKey - Constants that specify print job attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/AttributeKey
// NSPrintInfoAttributeKey is a string typedef
type PrintInfoAttributeKey = string
// PrintJobDispositionValue - Constants that specify values for the print job disposition.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct
// NSPrintJobDispositionValue is a string typedef
type PrintJobDispositionValue = string
// PrintInfoSettingKey - The type you use to specify a print info setting key.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/SettingKey
// NSPrintInfoSettingKey is a string typedef
type PrintInfoSettingKey = string
// PrintPanelJobStyleHint - Constants that specify job style hints for activating the simplified Print panel interface and setting the options to display.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrintPanel/JobStyleHint-swift.struct
// NSPrintPanelJobStyleHint is a string typedef
type PrintPanelJobStyleHint = string
// PrinterPaperName - The type you use to specify the name of a type of paper.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSPrinter/PaperName
// NSPrinterPaperName is a string typedef
type PrinterPaperName = string
// RulerViewUnitName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSRulerView/UnitName
// NSRulerViewUnitName is a string typedef
type RulerViewUnitName = string
// SearchFieldRecentsAutosaveName - The string that stores the name under which a search field automatically archives a list of recent search strings.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSearchField/RecentsAutosaveName-swift.typealias
// NSSearchFieldRecentsAutosaveName is a string typedef
type SearchFieldRecentsAutosaveName = string
// TextCheckingOptionKey - Constants that define options for text checking.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey
// NSTextCheckingOptionKey is a string typedef
type TextCheckingOptionKey = string
// SplitViewAutosaveName - The type that specifies the split view’s autosave name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSplitView/AutosaveName-swift.typealias
// NSSplitViewAutosaveName is a string typedef
type SplitViewAutosaveName = string
// StackViewVisibilityPriority - The various Auto Layout priorities for a view in the stack view to remain attached.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSStackView/VisibilityPriority
// NSStackViewVisibilityPriority has base type: float
type StackViewVisibilityPriority uintptr
// TextContentType - Constants that identify the semantic meaning for a text-entry area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentType
// NSTextContentType is a string typedef
type TextContentType = string
// ToolbarIdentifier - A string value that you use to differentiate your app’s toolbars.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbar/Identifier-swift.typealias
// NSToolbarIdentifier is a string typedef
type ToolbarIdentifier = string
// ToolbarItemIdentifier - Constants for the standard toolbar items that the system provides.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/Identifier
// NSToolbarItemIdentifier is a string typedef
type ToolbarItemIdentifier = string
// ToolbarItemVisibilityPriority - Constants that indicate which toolbar items to keep in the toolbar when space is limited.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/VisibilityPriority-swift.struct
type ToolbarItemVisibilityPriority int
// ToolbarUserInfoKey - Constants for specifying toolbar-related information in notifications.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSToolbarUserInfoKey
// NSToolbarUserInfoKey is a string typedef
type ToolbarUserInfoKey = string
// TouchBarCustomizationIdentifier - The default type for a Touch Bar customization identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBar/CustomizationIdentifier-swift.typealias
// NSTouchBarCustomizationIdentifier is a string typedef
type TouchBarCustomizationIdentifier = string
// TouchBarItemIdentifier - An identifier for an item in the Touch Bar.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/Identifier-swift.struct
// NSTouchBarItemIdentifier is a string typedef
type TouchBarItemIdentifier = string
// UserInterfaceItemIdentifier type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentifier
// NSUserInterfaceItemIdentifier is a string typedef
type UserInterfaceItemIdentifier = string
// DefinitionOptionKey - Keys to include in your definition.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/DefinitionOptionKey
// NSDefinitionOptionKey is a string typedef
type DefinitionOptionKey = string
// DefinitionPresentationType - Presentation options for the window.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/DefinitionPresentationType
// NSDefinitionPresentationType is a string typedef
type DefinitionPresentationType = string
// ViewFullScreenModeOptionKey - These constants are keys that you can use in the options dictionary in   and  .
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/FullScreenModeOptionKey
// NSViewFullScreenModeOptionKey is a string typedef
type ViewFullScreenModeOptionKey = string
// ToolTipTag - This type describes the rectangle used to identify a tooltip rectangle.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/ToolTipTag
type ToolTipTag int
// TrackingRectTag - This type describes the rectangle used to track the mouse.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSView/TrackingRectTag
type TrackingRectTag int
// WindowFrameAutosaveName - The type of a window’s frame autosave name.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/FrameAutosaveName-swift.typealias
// NSWindowFrameAutosaveName is a string typedef
type WindowFrameAutosaveName = string
// WindowLevel - The standard window levels in macOS.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/Level-swift.struct
type WindowLevel int
// WindowPersistableFrameDescriptor - The type of a window’s frame descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/PersistableFrameDescriptor
// NSWindowPersistableFrameDescriptor is a string typedef
type WindowPersistableFrameDescriptor = string
// WindowTabbingIdentifier - A value that allows a group of related windows.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWindow/TabbingIdentifier-swift.typealias
// NSWindowTabbingIdentifier is a string typedef
type WindowTabbingIdentifier = string
// WorkspaceDesktopImageOptionKey - Keys that indicate how to display a new desktop image.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/DesktopImageOptionKey
// NSWorkspaceDesktopImageOptionKey is a string typedef
type WorkspaceDesktopImageOptionKey = string
// WorkspaceFileOperationName - Constants that define types of file operations.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/FileOperationName
// NSWorkspaceFileOperationName is a string typedef
type WorkspaceFileOperationName = string
// WorkspaceLaunchConfigurationKey - The following keys can be used in the configuration dictionary of the   method.  Each key is optional, and if omitted, default behavior is applied.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchConfigurationKey
// NSWorkspaceLaunchConfigurationKey is a string typedef
type WorkspaceLaunchConfigurationKey = string

