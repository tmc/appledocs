// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation
import (
	"unsafe"
)


// Type alias for NSTimeInterval
type TimeInterval = float64  // NSTimeInterval

// Foundation-specific types

// Foundation geometry types - compatible with NSPoint, NSSize, NSRect, NSRange
type Point struct {
	X float64
	Y float64
}

type Size struct {
	Width  float64
	Height float64
}

type Rect struct {
	Origin Point
	Size   Size
}

type Range struct {
	Location int
	Length   int
}

// RectEdge defines which edge of a rectangle.
type RectEdge int

const (
	RectEdgeMinX RectEdge = 0
	RectEdgeMinY RectEdge = 1
	RectEdgeMaxX RectEdge = 2
	RectEdgeMaxY RectEdge = 3
)


// Fallback type aliases for undefined types
// These types are referenced in method signatures but not fully documented.
// Using unsafe.Pointer as fallback to allow code generation.
type Class unsafe.Pointer

type FileAttributeKey unsafe.Pointer

type NSAffineTransformStruct unsafe.Pointer

type NSCacheDelegate unsafe.Pointer

type NSCalculationError unsafe.Pointer

type NSCalendarOptions unsafe.Pointer

type NSComparator unsafe.Pointer

type NSDateComponentsFormatterZeroFormattingBehavior unsafe.Pointer

type NSDateFormatterBehavior unsafe.Pointer

type NSDecimal unsafe.Pointer

type NSDecodingFailurePolicy unsafe.Pointer

type NSDeviceCertification unsafe.Pointer

type NSDirectionalEdgeInsets unsafe.Pointer

type NSEdgeInsets unsafe.Pointer

type NSFastEnumerationState unsafe.Pointer

type NSFileCoordinatorReadingOptions unsafe.Pointer

type NSFileProviderItem unsafe.Pointer

type NSFileProviderItemIdentifier unsafe.Pointer

type NSFormattingContext unsafe.Pointer

type NSHTTPCookieAcceptPolicy unsafe.Pointer

type NSHelpManagerContextHelpKey unsafe.Pointer

type NSImageName unsafe.Pointer

type NSInsertionPosition unsafe.Pointer

type NSLengthFormatterUnit unsafe.Pointer

type NSLineBreakMode unsafe.Pointer

type NSMetadataQueryAttributeValueTuple unsafe.Pointer

type NSMetadataQueryResultGroup unsafe.Pointer

type NSNibName unsafe.Pointer

type NSObject unsafe.Pointer

type NSOperatingSystemVersion unsafe.Pointer

type NSOrderedCollectionChange unsafe.Pointer

type NSPostingStyle unsafe.Pointer

type NSProgressKind unsafe.Pointer

type NSRange unsafe.Pointer

type NSRoundingMode unsafe.Pointer

type NSSoundName unsafe.Pointer

type NSTestComparisonOperation unsafe.Pointer

type NSTextAlignment unsafe.Pointer

type NSTextCheckingTypes unsafe.Pointer

type NSURLHandle unsafe.Pointer

type NSUndoManagerUserInfoKey unsafe.Pointer

type NSUserAppleScriptTaskCompletionHandler unsafe.Pointer

type NSUserAutomatorTaskCompletionHandler unsafe.Pointer

type NSUserScriptTaskCompletionHandler unsafe.Pointer

type NSUserUnixTaskCompletionHandler unsafe.Pointer

type NSWorkspaceAuthorization unsafe.Pointer

type NSZone unsafe.Pointer



