// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation
import (
	"unsafe"
)


// TimeInterval for non-CoreGraphics frameworks
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
type AEDesc unsafe.Pointer

type AEEventClass unsafe.Pointer

type AEEventID unsafe.Pointer

type AEKeyword unsafe.Pointer

type AEReturnID unsafe.Pointer

type AETransactionID unsafe.Pointer

type APActivationPayload unsafe.Pointer

type BOOL unsafe.Pointer

type CATransform3D unsafe.Pointer

type CFRunLoopRef unsafe.Pointer

type CGAffineTransform unsafe.Pointer

type CGFloat unsafe.Pointer

type CGPoint unsafe.Pointer

type CGRect unsafe.Pointer

type CGSize unsafe.Pointer

type CGVector unsafe.Pointer

type CIBarcodeDescriptor unsafe.Pointer

type CKAllowedSharingOptions unsafe.Pointer

type CKContainer unsafe.Pointer

type CKSharePreparationHandler unsafe.Pointer

type CLLocationCoordinate2D unsafe.Pointer

type CMTime unsafe.Pointer

type CMTimeMapping unsafe.Pointer

type CMTimeRange unsafe.Pointer

type CMVideoDimensions unsafe.Pointer

type CSSearchableItemAttributeSet unsafe.Pointer

type Class unsafe.Pointer

type DescType unsafe.Pointer

type FourCharCode unsafe.Pointer

type GCPoint2 unsafe.Pointer

type GKRandomSource unsafe.Pointer

type INIntent unsafe.Pointer

type INInteraction unsafe.Pointer

type INShortcutAvailabilityOptions unsafe.Pointer

type K unsafe.Pointer

type KeyType unsafe.Pointer

type MKCoordinateSpan unsafe.Pointer

type MKMapItem unsafe.Pointer

type NCWidgetDisplayMode unsafe.Pointer

type NEHotspotHelperCommand unsafe.Pointer

type NFCNDEFMessage unsafe.Pointer

type NSActivityOptions unsafe.Pointer

type NSAlignmentOptions unsafe.Pointer

type NSAttributedStringEnumerationOptions unsafe.Pointer

type NSAttributedStringKey unsafe.Pointer

type NSBinarySearchingOptions unsafe.Pointer

type NSCalculationError unsafe.Pointer

type NSCalendarIdentifier unsafe.Pointer

type NSCalendarOptions unsafe.Pointer

type NSCalendarUnit unsafe.Pointer

type NSComparator unsafe.Pointer

type NSComparisonResult unsafe.Pointer

type NSCompoundPredicateType unsafe.Pointer

type NSDataBase64DecodingOptions unsafe.Pointer

type NSDataBase64EncodingOptions unsafe.Pointer

type NSDataCompressionAlgorithm unsafe.Pointer

type NSDataReadingOptions unsafe.Pointer

type NSDataSearchOptions unsafe.Pointer

type NSDataWritingOptions unsafe.Pointer

type NSDateComponentsFormatterZeroFormattingBehavior unsafe.Pointer

type NSDateFormatterBehavior unsafe.Pointer

type NSDateFormatterStyle unsafe.Pointer

type NSDecimal unsafe.Pointer

type NSDecodingFailurePolicy unsafe.Pointer

type NSDeviceCertification unsafe.Pointer

type NSDirectionalEdgeInsets unsafe.Pointer

type NSDirectoryEnumerationOptions unsafe.Pointer

type NSEdgeInsets unsafe.Pointer

type NSEnumerationOptions unsafe.Pointer

type NSErrorDomain unsafe.Pointer

type NSErrorUserInfoKey unsafe.Pointer

type NSExceptionName unsafe.Pointer

type NSFastEnumerationState unsafe.Pointer

type NSFileCoordinatorReadingOptions unsafe.Pointer

type NSFileManagerItemReplacementOptions unsafe.Pointer

type NSFileManagerResumeSyncBehavior unsafe.Pointer

type NSFileManagerUnmountOptions unsafe.Pointer

type NSFileManagerUploadLocalVersionConflictPolicy unsafe.Pointer

type NSFileProviderItem unsafe.Pointer

type NSFileProviderItemIdentifier unsafe.Pointer

type NSFormattingContext unsafe.Pointer

type NSHTTPCookieAcceptPolicy unsafe.Pointer

type NSHelpManagerContextHelpKey unsafe.Pointer

type NSISO8601DateFormatOptions unsafe.Pointer

type NSImage unsafe.Pointer

type NSImageName unsafe.Pointer

type NSInsertionPosition unsafe.Pointer

type NSInteger unsafe.Pointer

type NSItemProviderFileOptions unsafe.Pointer

type NSItemProviderRepresentationVisibility unsafe.Pointer

type NSKeyValueObservingOptions unsafe.Pointer

type NSLengthFormatterUnit unsafe.Pointer

type NSLineBreakMode unsafe.Pointer

type NSLinguisticTag unsafe.Pointer

type NSLinguisticTagScheme unsafe.Pointer

type NSLinguisticTaggerOptions unsafe.Pointer

type NSLocaleKey unsafe.Pointer

type NSLocaleLanguageDirection unsafe.Pointer

type NSMachPortOptions unsafe.Pointer

type NSNetServiceOptions unsafe.Pointer

type NSNibName unsafe.Pointer

type NSNotificationCoalescing unsafe.Pointer

type NSNotificationName unsafe.Pointer

type NSNumberFormatterBehavior unsafe.Pointer

type NSNumberFormatterPadPosition unsafe.Pointer

type NSNumberFormatterRoundingMode unsafe.Pointer

type NSNumberFormatterStyle unsafe.Pointer

type NSObject unsafe.Pointer

type NSOperatingSystemVersion unsafe.Pointer

type NSOperationQueuePriority unsafe.Pointer

type NSOrderedCollectionChange unsafe.Pointer

type NSOrderedCollectionDifferenceCalculationOptions unsafe.Pointer

type NSPasteboard unsafe.Pointer

type NSPoint unsafe.Pointer

type NSPointerFunctionsOptions unsafe.Pointer

type NSPostingStyle unsafe.Pointer

type NSProgressFileOperationKind unsafe.Pointer

type NSProgressKind unsafe.Pointer

type NSQualityOfService unsafe.Pointer

type NSRange unsafe.Pointer

type NSRangePointer unsafe.Pointer

type NSRect unsafe.Pointer

type NSRoundingMode unsafe.Pointer

type NSRunLoopMode unsafe.Pointer

type NSSaveOptions unsafe.Pointer

type NSSearchPathDirectory unsafe.Pointer

type NSSearchPathDomainMask unsafe.Pointer

type NSSize unsafe.Pointer

type NSSocketNativeHandle unsafe.Pointer

type NSSortOptions unsafe.Pointer

type NSSoundName unsafe.Pointer

type NSStreamPropertyKey unsafe.Pointer

type NSStreamStatus unsafe.Pointer

type NSStringCompareOptions unsafe.Pointer

type NSStringDrawingContext unsafe.Pointer

type NSStringDrawingOptions unsafe.Pointer

type NSStringEncoding unsafe.Pointer

type NSStringEncodingConversionOptions unsafe.Pointer

type NSStringEnumerationOptions unsafe.Pointer

type NSStringTransform unsafe.Pointer

type NSTestComparisonOperation unsafe.Pointer

type NSTextAlignment unsafe.Pointer

type NSTextBlock unsafe.Pointer

type NSTextCheckingType unsafe.Pointer

type NSTextList unsafe.Pointer

type NSTextTable unsafe.Pointer

type NSTimeInterval unsafe.Pointer

type NSUInteger unsafe.Pointer

type NSURLBookmarkCreationOptions unsafe.Pointer

type NSURLBookmarkFileCreationOptions unsafe.Pointer

type NSURLBookmarkResolutionOptions unsafe.Pointer

type NSURLCredentialPersistence unsafe.Pointer

type NSURLHandle unsafe.Pointer

type NSURLRelationship unsafe.Pointer

type NSURLRequestCachePolicy unsafe.Pointer

type NSURLRequestNetworkServiceType unsafe.Pointer

type NSURLResourceKey unsafe.Pointer

type NSURLSessionMultipathServiceType unsafe.Pointer

type NSURLSessionTaskState unsafe.Pointer

type NSURLSessionWebSocketCloseCode unsafe.Pointer

type NSUndoManagerUserInfoKey unsafe.Pointer

type NSUserActivityPersistentIdentifier unsafe.Pointer

type NSUserAppleScriptTaskCompletionHandler unsafe.Pointer

type NSUserAutomatorTaskCompletionHandler unsafe.Pointer

type NSUserNotificationActivationType unsafe.Pointer

type NSUserScriptTaskCompletionHandler unsafe.Pointer

type NSUserUnixTaskCompletionHandler unsafe.Pointer

type NSValueTransformerName unsafe.Pointer

type NSVolumeEnumerationOptions unsafe.Pointer

type NSWorkspaceAuthorization unsafe.Pointer

type NSXMLNodeKind unsafe.Pointer

type NSXMLNodeOptions unsafe.Pointer

type NSXPCConnectionOptions unsafe.Pointer

type NSZone unsafe.Pointer

type OBEXError unsafe.Pointer

type OSType unsafe.Pointer

type ObjectType unsafe.Pointer

type Protocol unsafe.Pointer

type RPBroadcastConfiguration unsafe.Pointer

type SCNMatrix4 unsafe.Pointer

type SCNVector3 unsafe.Pointer

type SCNVector4 unsafe.Pointer

type SEL unsafe.Pointer

type SRSensor unsafe.Pointer

type SSLProtocol unsafe.Pointer

type SecIdentityRef unsafe.Pointer

type SecTrustRef unsafe.Pointer

type UIBaselineAdjustment unsafe.Pointer

type UIEdgeInsets unsafe.Pointer

type UIFont unsafe.Pointer

type UIImage unsafe.Pointer

type UIOffset unsafe.Pointer

type UIPreferredPresentationStyle unsafe.Pointer

type UNNotificationAction unsafe.Pointer

type UTType unsafe.Pointer

type UnitType unsafe.Pointer



