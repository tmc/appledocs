// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

// Type aliases and typedefs
// FileAttributeKey - Keys in dictionaries used to get and set file attributes.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileAttributeKey
// NSFileAttributeKey is a string typedef
type FileAttributeKey = string
// FileAttributeType - Values representing a file’s type attribute.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileAttributeType
// NSFileAttributeType is a string typedef
type FileAttributeType = string
// FileProtectionType - Protection level values that can be associated with a file attribute key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/FileProtectionType
// NSFileProtectionType is a string typedef
type FileProtectionType = string
// AttributedStringKey - The attributes you apply to ranges of characters in an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key
// NSAttributedStringKey is a string typedef
type AttributedStringKey = string
// CalendarIdentifier - The supported calendar types.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier
// NSCalendarIdentifier is a string typedef
type CalendarIdentifier = string
// ErrorDomain type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSErrorDomain
// NSErrorDomain is a string typedef
type ErrorDomain = string
// UncaughtExceptionHandler type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSUncaughtExceptionHandler
// NSUncaughtExceptionHandler has base type: void (NSException *)
type UncaughtExceptionHandler uintptr
// StringEncodingDetectionOptionsKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/StringEncodingDetectionOptionsKey
// NSStringEncodingDetectionOptionsKey is a string typedef
type StringEncodingDetectionOptionsKey = string
// URLResourceKey - Keys that apply to file system URLs.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLResourceKey
// NSURLResourceKey is a string typedef
type URLResourceKey = string
// ExceptionName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSExceptionName
// NSExceptionName is a string typedef
type ExceptionName = string
// HashTableOptions - Components in a bit-field to specify the behavior of elements in an   object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSHashTableOptions
type HashTableOptions uint
// LinguisticTag - A token, lexical class, name, lemma, language, or script returned by a linguistic tagger for natural language text.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTag
// NSLinguisticTag is a string typedef
type LinguisticTag = string
// LinguisticTagScheme - Constants for the tag schemes specified when initializing a linguistic tagger.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLinguisticTagScheme
// NSLinguisticTagScheme is a string typedef
type LinguisticTagScheme = string
// LocaleKey - The keys used to access components of a locale.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSLocale/Key
// NSLocaleKey is a string typedef
type LocaleKey = string
// MapTableOptions - Constants used as components in a bitfield to specify the behavior of elements (keys and values) in an   object.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMapTableOptions
type MapTableOptions uint
// NotificationName - A structure that defines the name of a notification.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct
// NSNotificationName is a string typedef
type NotificationName = string
// SKIPPED: NSPoint - conflicts with existing enum "Point"
// PointArray - Type indicating a parameter is array of   structures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointArray
// NSPointArray has base type: NSPoint *
type PointArray uintptr
// PointPointer - Type indicating a parameter is a pointer to an   structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointPointer
// NSPointPointer has base type: NSPoint *
type PointPointer uintptr
// RangePointer - Type indicating a parameter is a pointer to an   structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRangePointer
// NSRangePointer has base type: NSRange *
type RangePointer uintptr
// SKIPPED: NSRect - conflicts with existing enum "Rect"
// RectArray - Type indicating a parameter is array of   structures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectArray
// NSRectArray has base type: NSRect *
type RectArray uintptr
// RectPointer - Type indicating a parameter is a pointer to an   structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSRectPointer
// NSRectPointer has base type: NSRect *
type RectPointer uintptr
// SKIPPED: NSSize - conflicts with existing enum "Size"
// SizeArray - Type indicating a parameter is an array of   structures.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSizeArray
// NSSizeArray has base type: NSSize *
type SizeArray uintptr
// SizePointer - Type indicating parameter is a pointer to an   structure.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSSizePointer
// NSSizePointer has base type: NSSize *
type SizePointer uintptr
// TextCheckingKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingKey
// NSTextCheckingKey is a string typedef
type TextCheckingKey = string
// TextCheckingTypes - Defines the types of checking that are available. These values can be combined using the C-bitwise OR operator. The system supports its own internal types, and the user can extend those types by subclassing   and adding their own custom types.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSTextCheckingTypes
// NSTextCheckingTypes has base type: uint64_t
type TextCheckingTypes uintptr
// URLBookmarkFileCreationOptions - Options used when creating file bookmark data
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkFileCreationOptions
type URLBookmarkFileCreationOptions uint
// Audio3DVector - A structure that represents a vector in 3D space, in degrees.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFAudio/AVAudio3DVector
// AVAudio3DVector has base type: struct AVAudio3DPoint
type Audio3DVector uintptr
// CaptureDeviceType - A structure that defines the device types the framework supports.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct
// AVCaptureDeviceType is a string typedef
type CaptureDeviceType = string
// AXTechnology type alias
//
// [Full Topic]: https://developer.apple.com/documentation/Accessibility/AccessibilityTechnology
// AXTechnology has base type: NSString * const
type AXTechnology uintptr
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
// NibName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSNib/Name
// NSNibName is a string typedef
type NibName = string
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
// TextCheckingOptionKey - Constants that define options for text checking.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSSpellChecker/OptionKey
// NSTextCheckingOptionKey is a string typedef
type TextCheckingOptionKey = string
// TextContentType - Constants that identify the semantic meaning for a text-entry area.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTextContentType
// NSTextContentType is a string typedef
type TextContentType = string
// UserInterfaceItemIdentifier type alias
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentifier
// NSUserInterfaceItemIdentifier is a string typedef
type UserInterfaceItemIdentifier = string
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
// WorkspaceLaunchConfigurationKey - The following keys can be used in the configuration dictionary of the   method.  Each key is optional, and if omitted, default behavior is applied.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSWorkspace/LaunchConfigurationKey
// NSWorkspaceLaunchConfigurationKey is a string typedef
type WorkspaceLaunchConfigurationKey = string
// BADownloaderPriority - A type that determines the execution priority of a scheduled asset download.
//
// [Full Topic]: https://developer.apple.com/documentation/BackgroundAssets/BADownload/Priority-swift.struct
type BADownloaderPriority int
// FontRef - A font object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTFont
// CTFontRef has base type: const struct __CTFont *
type FontRef uintptr
// GlyphInfoRef - Override a font’s specified mapping from Unicode to the glyph ID.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTGlyphInfo
// CTGlyphInfoRef has base type: const struct __CTGlyphInfo *
type GlyphInfoRef uintptr
// ParagraphStyleRef - Paragraph or ruler attributes in an attributed string.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTParagraphStyle
// CTParagraphStyleRef has base type: const struct __CTParagraphStyle *
type ParagraphStyleRef uintptr
// RunDelegateRef - A run delegate.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreText/CTRunDelegate
// CTRunDelegateRef has base type: const struct __CTRunDelegate *
type RunDelegateRef uintptr
// HKClinicalTypeIdentifier - A type identifier for the different categories of clinical records.
//
// [Full Topic]: https://developer.apple.com/documentation/HealthKit/HKClinicalTypeIdentifier
// HKClinicalTypeIdentifier is a string typedef
type HKClinicalTypeIdentifier = string
// Integer - Describes an integer.
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSInteger
// NSInteger has base type: long
type Integer uintptr
// objc_exception_handler type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_handler
// objc_exception_handler has base type: void (*)(id, void *)
type objc_exception_handler uintptr
// objc_exception_matcher type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_matcher
// objc_exception_matcher has base type: int (*)(Class, id)
type objc_exception_matcher uintptr
// objc_exception_preprocessor type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_exception_preprocessor
// objc_exception_preprocessor has base type: id (*)(id)
type objc_exception_preprocessor uintptr
// objc_func_loadImage type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_func_loadImage
// objc_func_loadImage has base type: void (*)(const struct mach_header *)
type objc_func_loadImage uintptr
// objc_hook_getClass type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_hook_getClass
// objc_hook_getClass has base type: _Bool (*)(const char *, Class *)
type objc_hook_getClass uintptr
// objc_hook_getImageName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_hook_getImageName
// objc_hook_getImageName has base type: _Bool (*)(Class, const char **)
type objc_hook_getImageName uintptr
// objc_hook_lazyClassNamer type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_hook_lazyClassNamer
// objc_hook_lazyClassNamer has base type: const char *(*)(Class)
type objc_hook_lazyClassNamer uintptr
// objc_objectptr_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_objectptr_t
// objc_objectptr_t has base type: const void *
type objc_objectptr_t uintptr
// objc_uncaught_exception_handler type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_uncaught_exception_handler
// objc_uncaught_exception_handler has base type: void (*)(id)
type objc_uncaught_exception_handler uintptr
// objc_zone_t type alias
//
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/objc_zone_t
// objc_zone_t has base type: struct _malloc_zone_t *
type objc_zone_t uintptr
// PassLibraryNotificationName - The types of notifications that the pass library posts.
//
// [Full Topic]: https://developer.apple.com/documentation/PassKit/PKPassLibraryNotificationName
// PKPassLibraryNotificationName is a string typedef
type PassLibraryNotificationName = string
// AbsoluteTime - Type used to represent a specific point in time relative to the absolute reference date of 1 Jan 2001 00:00:00 GMT.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAbsoluteTime
// CFAbsoluteTime has base type: CFTimeInterval
type AbsoluteTime uintptr
// AllocatorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocator
// CFAllocatorRef has base type: const struct __CFAllocator *
type AllocatorRef uintptr
// AllocatorAllocateCallBack - A prototype for a function callback that allocates memory of a requested size.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorAllocateCallBack
// CFAllocatorAllocateCallBack has base type: void *(*)(long, unsigned long, void *)
type AllocatorAllocateCallBack uintptr
// AllocatorCopyDescriptionCallBack - A prototype for a function callback that provides a description of the specified data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorCopyDescriptionCallBack
// CFAllocatorCopyDescriptionCallBack has base type: const struct __CFString *(*)(const void *)
type AllocatorCopyDescriptionCallBack uintptr
// AllocatorDeallocateCallBack - A prototype for a function callback that deallocates a block of memory.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorDeallocateCallBack
// CFAllocatorDeallocateCallBack has base type: void (*)(void *, void *)
type AllocatorDeallocateCallBack uintptr
// AllocatorPreferredSizeCallBack - A prototype for a function callback that gives the size of memory likely to be allocated, given a certain request.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorPreferredSizeCallBack
// CFAllocatorPreferredSizeCallBack has base type: long (*)(long, unsigned long, void *)
type AllocatorPreferredSizeCallBack uintptr
// AllocatorReallocateCallBack - A prototype for a function callback that reallocates memory of a requested size for an existing block of memory.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReallocateCallBack
// CFAllocatorReallocateCallBack has base type: void *(*)(void *, long, unsigned long, void *)
type AllocatorReallocateCallBack uintptr
// AllocatorReleaseCallBack - A prototype for a function callback that releases the given data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorReleaseCallBack
// CFAllocatorReleaseCallBack has base type: void (*)(const void *)
type AllocatorReleaseCallBack uintptr
// AllocatorRetainCallBack - A prototype for a function callback that retains the given data.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorRetainCallBack
// CFAllocatorRetainCallBack has base type: const void *(*)(const void *)
type AllocatorRetainCallBack uintptr
// AllocatorTypeID type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAllocatorTypeID
// CFAllocatorTypeID has base type: unsigned long long
type AllocatorTypeID uintptr
// ArrayRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArray
// CFArrayRef has base type: const struct __CFArray *
type ArrayRef uintptr
// ArrayApplierFunction - Prototype of a callback function that may be applied to every value in an array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayApplierFunction
// CFArrayApplierFunction has base type: void (*)(const void *, void *)
type ArrayApplierFunction uintptr
// ArrayCopyDescriptionCallBack - Prototype of a callback function used to get a description of a value in an array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayCopyDescriptionCallBack
// CFArrayCopyDescriptionCallBack has base type: const struct __CFString *(*)(const void *)
type ArrayCopyDescriptionCallBack uintptr
// ArrayEqualCallBack - Prototype of a callback function used to determine if two values in an array are equal.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayEqualCallBack
// CFArrayEqualCallBack has base type: unsigned char (*)(const void *, const void *)
type ArrayEqualCallBack uintptr
// ArrayReleaseCallBack - Prototype of a callback function used to release a value before it’s removed from an array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayReleaseCallBack
// CFArrayReleaseCallBack has base type: void (*)(const struct __CFAllocator *, const void *)
type ArrayReleaseCallBack uintptr
// ArrayRetainCallBack - Prototype of a callback function used to retain a value being added to an array.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFArrayRetainCallBack
// CFArrayRetainCallBack has base type: const void *(*)(const struct __CFAllocator *, const void *)
type ArrayRetainCallBack uintptr
// AttributedStringRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFAttributedString
// CFAttributedStringRef has base type: const struct __CFAttributedString *
type AttributedStringRef uintptr
// BagRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBag
// CFBagRef has base type: const struct __CFBag *
type BagRef uintptr
// BagApplierFunction - Prototype of a callback function that may be applied to every value in a bag.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagApplierFunction
// CFBagApplierFunction has base type: void (*)(const void *, void *)
type BagApplierFunction uintptr
// BagCopyDescriptionCallBack - Prototype of a callback function used to get a description of a value in a bag.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagCopyDescriptionCallBack
// CFBagCopyDescriptionCallBack has base type: const struct __CFString *(*)(const void *)
type BagCopyDescriptionCallBack uintptr
// BagEqualCallBack - Prototype of a callback function used to determine if two values in a bag are equal.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagEqualCallBack
// CFBagEqualCallBack has base type: unsigned char (*)(const void *, const void *)
type BagEqualCallBack uintptr
// BagHashCallBack - Prototype of a callback function invoked to compute a hash code for a value. Hash codes are used when values are accessed, added, or removed from a collection.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagHashCallBack
// CFBagHashCallBack has base type: unsigned long (*)(const void *)
type BagHashCallBack uintptr
// BagReleaseCallBack - Prototype of a callback function used to release a value before it’s removed from a bag.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagReleaseCallBack
// CFBagReleaseCallBack has base type: void (*)(const struct __CFAllocator *, const void *)
type BagReleaseCallBack uintptr
// BagRetainCallBack - Prototype of a callback function used to retain a value being added to a bag.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBagRetainCallBack
// CFBagRetainCallBack has base type: const void *(*)(const struct __CFAllocator *, const void *)
type BagRetainCallBack uintptr
// BinaryHeapRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeap
// CFBinaryHeapRef has base type: struct __CFBinaryHeap *
type BinaryHeapRef uintptr
// BinaryHeapApplierFunction - Callback function used to apply a function to all members of a binary heap.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBinaryHeapApplierFunction
// CFBinaryHeapApplierFunction has base type: void (*)(const void *, void *)
type BinaryHeapApplierFunction uintptr
// BitVectorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBitVector
// CFBitVectorRef has base type: const struct __CFBitVector *
type BitVectorRef uintptr
// BooleanRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBoolean
// CFBooleanRef has base type: const struct __CFBoolean *
type BooleanRef uintptr
// BundleRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundle
// CFBundleRef has base type: struct __CFBundle *
type BundleRef uintptr
// BundleRefNum - Type that identifies a distinct reference number for a resource map.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBundleRefNum
type BundleRefNum int32
// CalendarRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendar
// CFCalendarRef has base type: struct __CFCalendar *
type CalendarRef uintptr
// CharacterSetRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSet
// CFCharacterSetRef has base type: const struct __CFCharacterSet *
type CharacterSetRef uintptr
// ComparatorFunction - Callback function that compares two values. You provide a pointer to this callback in certain Core Foundation sorting functions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFComparatorFunction
// CFComparatorFunction has base type: enum CFComparisonResult (*)(const void *, const void *, void *)
type ComparatorFunction uintptr
// DataRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFData
// CFDataRef has base type: const struct __CFData *
type DataRef uintptr
// DateRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDate
// CFDateRef has base type: const struct __CFDate *
type DateRef uintptr
// DateFormatterRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatter
// CFDateFormatterRef has base type: struct __CFDateFormatter *
type DateFormatterRef uintptr
// DateFormatterKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterKey
// CFDateFormatterKey has base type: CFStringRef
type DateFormatterKey uintptr
// DictionaryRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionary
// CFDictionaryRef has base type: const struct __CFDictionary *
type DictionaryRef uintptr
// DictionaryApplierFunction - Prototype of a callback function that may be applied to every key-value pair in a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryApplierFunction
// CFDictionaryApplierFunction has base type: void (*)(const void *, const void *, void *)
type DictionaryApplierFunction uintptr
// DictionaryCopyDescriptionCallBack - Prototype of a callback function used to get a description of a value or key in a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryCopyDescriptionCallBack
// CFDictionaryCopyDescriptionCallBack has base type: const struct __CFString *(*)(const void *)
type DictionaryCopyDescriptionCallBack uintptr
// DictionaryEqualCallBack - Prototype of a callback function used to determine if two values or keys in a dictionary are equal.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryEqualCallBack
// CFDictionaryEqualCallBack has base type: unsigned char (*)(const void *, const void *)
type DictionaryEqualCallBack uintptr
// DictionaryHashCallBack - Prototype of a callback function invoked to compute a hash code for a key. Hash codes are used when key-value pairs are accessed, added, or removed from a collection.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryHashCallBack
// CFDictionaryHashCallBack has base type: unsigned long (*)(const void *)
type DictionaryHashCallBack uintptr
// DictionaryReleaseCallBack - Prototype of a callback function used to release a key-value pair before it’s removed from a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryReleaseCallBack
// CFDictionaryReleaseCallBack has base type: void (*)(const struct __CFAllocator *, const void *)
type DictionaryReleaseCallBack uintptr
// DictionaryRetainCallBack - Prototype of a callback function used to retain a value or key being added to a dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFDictionaryRetainCallBack
// CFDictionaryRetainCallBack has base type: const void *(*)(const struct __CFAllocator *, const void *)
type DictionaryRetainCallBack uintptr
// ErrorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFError
// CFErrorRef has base type: struct __CFError *
type ErrorRef uintptr
// FileDescriptorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptor
// CFFileDescriptorRef has base type: struct __CFFileDescriptor *
type FileDescriptorRef uintptr
// HashCode - A type for hash codes returned by the   function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFHashCode
// CFHashCode has base type: unsigned long
type HashCode uintptr
// Index - Priority values used for kAXPriorityKey
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFIndex
// CFIndex has base type: long
type Index uintptr
// LocaleRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocale
// CFLocaleRef has base type: const struct __CFLocale *
type LocaleRef uintptr
// LocaleIdentifier type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleIdentifier
// CFLocaleIdentifier has base type: CFStringRef
type LocaleIdentifier uintptr
// MutableArrayRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableArray
// CFMutableArrayRef has base type: struct __CFArray *
type MutableArrayRef uintptr
// MutableAttributedStringRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableAttributedString
// CFMutableAttributedStringRef has base type: struct __CFAttributedString *
type MutableAttributedStringRef uintptr
// MutableBagRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableBag
// CFMutableBagRef has base type: struct __CFBag *
type MutableBagRef uintptr
// MutableCharacterSetRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableCharacterSet
// CFMutableCharacterSetRef has base type: struct __CFCharacterSet *
type MutableCharacterSetRef uintptr
// MutableDataRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableData
// CFMutableDataRef has base type: struct __CFData *
type MutableDataRef uintptr
// MutableDictionaryRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableDictionary
// CFMutableDictionaryRef has base type: struct __CFDictionary *
type MutableDictionaryRef uintptr
// MutableStringRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableString
// CFMutableStringRef has base type: struct __CFString *
type MutableStringRef uintptr
// NullRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNull
// CFNullRef has base type: const struct __CFNull *
type NullRef uintptr
// NumberRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumber
// CFNumberRef has base type: const struct __CFNumber *
type NumberRef uintptr
// NumberFormatterKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterKey
// CFNumberFormatterKey has base type: CFStringRef
type NumberFormatterKey uintptr
// OptionFlags - A bitfield used for passing special allocation and other requests into Core Foundation functions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFOptionFlags
// CFOptionFlags has base type: unsigned long
type OptionFlags uintptr
// PlugInRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugIn
// CFPlugInRef has base type: struct __CFBundle *
type PlugInRef uintptr
// ReadStreamRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStream
// CFReadStreamRef has base type: struct __CFReadStream *
type ReadStreamRef uintptr
// ReadStreamClientCallBack - Callback invoked when certain types of activity takes place on a readable stream.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFReadStreamClientCallBack
// CFReadStreamClientCallBack has base type: void (*)(struct __CFReadStream *, enum CFStreamEventType, void *)
type ReadStreamClientCallBack uintptr
// RunLoopRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoop
// CFRunLoopRef has base type: struct __CFRunLoop *
type RunLoopRef uintptr
// RunLoopMode type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopMode
// CFRunLoopMode has base type: CFStringRef
type RunLoopMode uintptr
// SocketRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocket
// CFSocketRef has base type: struct __CFSocket *
type SocketRef uintptr
// SocketNativeHandle - Type for the platform-specific native socket handle.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketNativeHandle
type SocketNativeHandle int32
// StreamPropertyKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStreamPropertyKey
// CFStreamPropertyKey has base type: CFStringRef
type StreamPropertyKey uintptr
// StringRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFString
// CFStringRef has base type: const struct __CFString *
type StringRef uintptr
// StringEncoding - An integer type for constants used to specify supported string encodings in various CFString functions.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringEncoding
// CFStringEncoding has base type: UInt32
type StringEncoding uintptr
// StringTokenizerRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizer
// CFStringTokenizerRef has base type: struct __CFStringTokenizer *
type StringTokenizerRef uintptr
// TimeZoneRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZone
// CFTimeZoneRef has base type: const struct __CFTimeZone *
type TimeZoneRef uintptr
// TypeID - A type for unique, constant integer values that identify particular Core Foundation opaque types.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTypeID
// CFTypeID has base type: unsigned long
type TypeID uintptr
// TypeRef - An untyped “generic” reference to any Core Foundation object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTypeRef
// CFTypeRef has base type: const void *
type TypeRef uintptr
// URLRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURL
// CFURLRef has base type: const struct __CFURL *
type URLRef uintptr
// UUIDRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUID
// CFUUIDRef has base type: const struct __CFUUID *
type UUIDRef uintptr
// WriteStreamRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStream
// CFWriteStreamRef has base type: struct __CFWriteStream *
type WriteStreamRef uintptr
// WriteStreamClientCallBack - Callback invoked when certain types of activity takes place on a writable stream.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStreamClientCallBack
// CFWriteStreamClientCallBack has base type: void (*)(struct __CFWriteStream *, enum CFStreamEventType, void *)
type WriteStreamClientCallBack uintptr
// ColorRef - A set of components that define a color, with a color space specifying how to interpret them.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColor
// CGColorRef has base type: struct CGColor *
type ColorRef uintptr
// ColorSpaceRef - A profile that specifies how to interpret a color value for display.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// CGColorSpaceRef has base type: struct CGColorSpace *
type ColorSpaceRef uintptr
// ImageRef - A bitmap image or image mask.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// CGImageRef has base type: struct CGImage *
type ImageRef uintptr
// URLFileProtectionType - Protection-level values for a URL resource key.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLFileProtection
// NSURLFileProtectionType is a string typedef
type URLFileProtectionType = string
// URLFileResourceType - Possible values for the type of file resource.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLFileResourceType
// NSURLFileResourceType is a string typedef
type URLFileResourceType = string
// URLThumbnailDictionaryItem - Possible keys for the   dictionary.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLThumbnailDictionaryItem
// NSURLThumbnailDictionaryItem is a string typedef
type URLThumbnailDictionaryItem = string
// URLUbiquitousItemDownloadingStatus - Values that describe the iCloud storage state of a file.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLUbiquitousItemDownloadingStatus
// NSURLUbiquitousItemDownloadingStatus is a string typedef
type URLUbiquitousItemDownloadingStatus = string
// URLUbiquitousSharedItemPermissions - The key for the permissions of a shared item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLUbiquitousSharedItemPermissions
// NSURLUbiquitousSharedItemPermissions is a string typedef
type URLUbiquitousSharedItemPermissions = string
// URLUbiquitousSharedItemRole - The key for the role of a shared item.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/URLUbiquitousSharedItemRole
// NSURLUbiquitousSharedItemRole is a string typedef
type URLUbiquitousSharedItemRole = string
// TimeInterval - A number of seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/TimeInterval
// NSTimeInterval has base type: double
type TimeInterval uintptr
// unichar - Type for UTF-16 code units.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/unichar
// unichar has base type: unsigned short
type unichar uintptr
// NetworkReachabilityRef - The handle to a network address or name.
//
// [Full Topic]: https://developer.apple.com/documentation/SystemConfiguration/SCNetworkReachability
// SCNetworkReachabilityRef has base type: const struct __SCNetworkReachability *
type NetworkReachabilityRef uintptr
// AccessibilityAssistiveTechnologyIdentifier - Identifiers for assistive apps.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibility/AssistiveTechnologyIdentifier
// UIAccessibilityAssistiveTechnologyIdentifier is a string typedef
type AccessibilityAssistiveTechnologyIdentifier = string
// AccessibilityPriority - Constants that specify priorities for accessibility announcements.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibilityPriority
// UIAccessibilityPriority is a string typedef
type AccessibilityPriority = string
// AccessibilityTextualContext - Constants that describe a named context that helps identify and classify the type of text inside an element.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibilityTextualContext
// UIAccessibilityTextualContext is a string typedef
type AccessibilityTextualContext = string
// AccessibilityTraits - Constants that describe how an accessibility element behaves.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIAccessibilityTraits
// UIAccessibilityTraits has base type: uint64_t
type AccessibilityTraits uintptr
// ApplicationOpenExternalURLOptionsKey - Options for opening a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplication/OpenExternalURLOptionsKey
// UIApplicationOpenExternalURLOptionsKey is a string typedef
type ApplicationOpenExternalURLOptionsKey = string
// ApplicationOpenURLOptionsKey - Keys you use to access values in the options dictionary when opening a URL.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIApplication/OpenURLOptionsKey
// UIApplicationOpenURLOptionsKey is a string typedef
type ApplicationOpenURLOptionsKey = string
// BackgroundTaskIdentifier - A unique token that identifies a request to run in the background.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIBackgroundTaskIdentifier
type BackgroundTaskIdentifier uint
// ContentSizeCategory - Constants that indicate the preferred size of your content.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIContentSizeCategory
// UIContentSizeCategory is a string typedef
type ContentSizeCategory = string
// DocumentCreationIntent - An app intent that creates new documents for your app.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIDocument/CreationIntent
// UIDocumentCreationIntent is a string typedef
type DocumentCreationIntent = string
// FontWeight - Constants that represent standard typeface styles.
//
// [Full Topic]: https://developer.apple.com/documentation/UIKit/UIFont/Weight
// UIFontWeight has base type: CGFloat
type FontWeight uintptr

