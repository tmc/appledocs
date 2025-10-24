// Code generated from Apple documentation for QuickLookThumbnailing. DO NOT EDIT.

package quicklookthumbnailing

/* debug [enums.gen.go]: Generating 3 enums for QuickLookThumbnailing */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum QLThumbnailError (6 cases) */
// QLThumbnailError - Error codes that may be returned when generating a thumbnail.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailError-swift.struct/Code
type QLThumbnailError uint

const (
	// QLThumbnailErrorGenerationFailed - The thumbnail couldn’t be created for the given file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailError-swift.struct/Code/generationFailed
	QLThumbnailErrorGenerationFailed QLThumbnailError = 0
	// QLThumbnailErrorNoCachedThumbnail - A low-quality thumbnail couldn’t be created.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailError-swift.struct/Code/noCachedThumbnail
	QLThumbnailErrorNoCachedThumbnail QLThumbnailError = 0
	// QLThumbnailErrorNoCloudThumbnail - The thumbnail for a remote file couldn’t be created.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailError-swift.struct/Code/noCloudThumbnail
	QLThumbnailErrorNoCloudThumbnail QLThumbnailError = 0
	// QLThumbnailErrorRequestCancelled - The request to create a thumbnail was canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailError-swift.struct/Code/requestCancelled
	QLThumbnailErrorRequestCancelled QLThumbnailError = 0
	// QLThumbnailErrorRequestInvalid - The request to create a thumbnail was invalid, for example, there’s no file at a provided URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailError-swift.struct/Code/requestInvalid
	QLThumbnailErrorRequestInvalid QLThumbnailError = 0
	// QLThumbnailErrorSavingToURLFailed - The thumbnail couldn’t be saved at the given URL.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailError-swift.struct/Code/savingToURLFailed
	QLThumbnailErrorSavingToURLFailed QLThumbnailError = 0
)

/* debug [enums.gen.go]: Processing enum QLThumbnailGenerationRequestRepresentationTypes (4 cases) */
// QLThumbnailGenerationRequestRepresentationTypes - The various types of thumbnails that you can request.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/RepresentationTypes-swift.struct
type QLThumbnailGenerationRequestRepresentationTypes uint

const (
	// QLThumbnailGenerationRequestRepresentationTypeAll - The thumbnail type to generate all possible thumbnail representations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/RepresentationTypes-swift.struct/all
	QLThumbnailGenerationRequestRepresentationTypeAll QLThumbnailGenerationRequestRepresentationTypes = 0
	// QLThumbnailGenerationRequestRepresentationTypeIcon - A file icon representation of a file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/RepresentationTypes-swift.struct/icon
	QLThumbnailGenerationRequestRepresentationTypeIcon QLThumbnailGenerationRequestRepresentationTypes = 0
	// QLThumbnailGenerationRequestRepresentationTypeLowQualityThumbnail - A faster to generate version of the thumbnail that may sacrifice quality for speed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/RepresentationTypes-swift.struct/lowQualityThumbnail
	QLThumbnailGenerationRequestRepresentationTypeLowQualityThumbnail QLThumbnailGenerationRequestRepresentationTypes = 0
	// QLThumbnailGenerationRequestRepresentationTypeThumbnail - A thumbnail representation of a file.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailGenerator/Request/RepresentationTypes-swift.struct/thumbnail
	QLThumbnailGenerationRequestRepresentationTypeThumbnail QLThumbnailGenerationRequestRepresentationTypes = 0
)

/* debug [enums.gen.go]: Processing enum QLThumbnailRepresentationType (3 cases) */
// QLThumbnailRepresentationType - The different types of thumbnails that you can create.
//
// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/RepresentationType
type QLThumbnailRepresentationType uint

const (
	// QLThumbnailRepresentationTypeIcon - A file icon representation of an image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/RepresentationType/icon
	QLThumbnailRepresentationTypeIcon QLThumbnailRepresentationType = 0
	// QLThumbnailRepresentationTypeLowQualityThumbnail - A cached thumbnail representation of an image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/RepresentationType/lowQualityThumbnail
	QLThumbnailRepresentationTypeLowQualityThumbnail QLThumbnailRepresentationType = 0
	// QLThumbnailRepresentationTypeThumbnail - A thumbnail representation of an image.
	//
	// [Full Topic]: https://developer.apple.com/documentation/QuickLookThumbnailing/QLThumbnailRepresentation/RepresentationType/thumbnail
	QLThumbnailRepresentationTypeThumbnail QLThumbnailRepresentationType = 0
)


