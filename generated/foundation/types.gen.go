// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation
import (
	"unsafe"
)


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
// Using inferred base types as fallback to allow code generation.
type ActivityItemsConfigurationReading = int

type AffineTransformStruct = int

type BaselineAdjustment = int

type CalculationError = int

type Comparator = int

type Decimal = int

type DecodingFailurePolicy = int

type DeviceCertification = int

type DirectionalEdgeInsets = int

type EdgeInsets = int

type ExceptionName = int

type FastEnumerationState = int

type FileProviderItem = int

type FileProviderItemIdentifier = int

type HTTPCookieAcceptPolicy = int

type HelpManagerContextHelpKey = int

type ImageName = int

type Integer = int

type ItemProviderCompletionHandler = int

type ItemProviderLoadHandler = int

type KeyedArchiverDelegate = int

type KeyedUnarchiverDelegate = int

type LineBreakMode = int

type LinguisticTagScheme = int

type LocationCoordinate2D = int

type NMatrix4 = int

type NVector3 = int

type NVector4 = int

type NibName = int

type Offset = int

type OperatingSystemVersion = int

type OrderedCollectionChange = int

type ProcessPerformanceProfile = int

type ProgressPublishingHandler = int

type PropertyListFormat = int

type RangePointer = int

type RoundingMode = int

type RunLoopRef = int

type SoundName = int

type SpellServerDelegate = int

type StreamStatus = int

type StringEncoding = int

type StringTransform = int

type TextAlignment = int

type TextCheckingTypes = int

type Time = int

type TimeMapping = int

type TimeRange = int

type Transform3D = int

type UInteger = int

type UserAppleScriptTaskCompletionHandler = int

type UserAutomatorTaskCompletionHandler = int

type UserScriptTaskCompletionHandler = int

type UserUnixTaskCompletionHandler = int

type ValueTransformerName = int

type VideoDimensions = int

type XPCListenerDelegate = int

type Zone = int



