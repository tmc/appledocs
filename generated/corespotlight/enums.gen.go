// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

// Enum types and constants
// CSIndexErrorCode - Error codes that describe indexing-specific errors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code
type CSIndexErrorCode uint

const (
// CSIndexErrorCodeIndexUnavailableError - The indexer is unavailable.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code/indexUnavailableError
CSIndexErrorCodeIndexUnavailableError CSIndexErrorCode = 0
// CSIndexErrorCodeIndexingUnsupported - Indexing isn’t supported on the device.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code/indexingUnsupported
CSIndexErrorCodeIndexingUnsupported CSIndexErrorCode = 0
// CSIndexErrorCodeInvalidClientStateError - The provided client state data is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code/invalidClientStateError
CSIndexErrorCodeInvalidClientStateError CSIndexErrorCode = 0
// CSIndexErrorCodeInvalidItemError - The searchable item object is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code/invalidItemError
CSIndexErrorCodeInvalidItemError CSIndexErrorCode = 0
// CSIndexErrorCodeMismatchedClientState - The provided client state did not match the information in the index.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code/mismatchedClientState
CSIndexErrorCodeMismatchedClientState CSIndexErrorCode = 0
// CSIndexErrorCodeQuotaExceeded - The quota for the bundle has been exceeded.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code/quotaExceeded
CSIndexErrorCodeQuotaExceeded CSIndexErrorCode = 0
// CSIndexErrorCodeRemoteConnectionError - An error occurred while communicating with the remote process.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code/remoteConnectionError
CSIndexErrorCodeRemoteConnectionError CSIndexErrorCode = 0
// CSIndexErrorCodeUnknownError - An unknown error occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code/unknownError
CSIndexErrorCodeUnknownError CSIndexErrorCode = 0
)

// CSSearchQuerySourceOptions - The query source options to allow or deny Mail messages in the search.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/SourceOptions-swift.struct
type CSSearchQuerySourceOptions uint

const (
// CSSearchQuerySourceOptionAllowMail - The query allows Mail messages in the search.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/SourceOptions-swift.struct/allowMail
CSSearchQuerySourceOptionAllowMail CSSearchQuerySourceOptions = 0
// CSSearchQuerySourceOptionDefault - The query uses the default search option that excludes Mail messages.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQuerySourceOptions/CSSearchQuerySourceOptionDefault
CSSearchQuerySourceOptionDefault CSSearchQuerySourceOptions = 0
)

// CSSearchQueryErrorCode - Error codes that describe reasons a query might fail.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryError/Code
type CSSearchQueryErrorCode uint

const (
// CSSearchQueryErrorCodeCancelled - The query stopped because someone canceled it.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryError/Code/cancelled
CSSearchQueryErrorCodeCancelled CSSearchQueryErrorCode = 0
// CSSearchQueryErrorCodeIndexUnreachable - The index is unreachable.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryError/Code/indexUnreachable
CSSearchQueryErrorCodeIndexUnreachable CSSearchQueryErrorCode = 0
// CSSearchQueryErrorCodeInvalidQuery - The query is syntactically invalid or specifies items that your app   doesn’t have access to.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryError/Code/invalidQuery
CSSearchQueryErrorCodeInvalidQuery CSSearchQueryErrorCode = 0
// CSSearchQueryErrorCodeUnknown - An unknown error occurred.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryError/Code/unknown
CSSearchQueryErrorCodeUnknown CSSearchQueryErrorCode = 0
)

// CSSearchableItemUpdateListenerOptions - The set of options that contain metadata-associated summarization and
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/UpdateListenerOptions-swift.struct
type CSSearchableItemUpdateListenerOptions uint

const (
// CSSearchableItemUpdateListenerOptionPriority - A value that describes the listener priority options.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/UpdateListenerOptions-swift.struct/priority
CSSearchableItemUpdateListenerOptionPriority CSSearchableItemUpdateListenerOptions = 0
// CSSearchableItemUpdateListenerOptionSummarization - A value that describes the listener summarization options.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItem/UpdateListenerOptions-swift.struct/summarization
CSSearchableItemUpdateListenerOptionSummarization CSSearchableItemUpdateListenerOptions = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchableItemUpdateListenerOptions/CSSearchableItemUpdateListenerOptionDefault
CSSearchableItemUpdateListenerOptionDefault CSSearchableItemUpdateListenerOptions = 0
)

// CSSuggestionKind - The suggestion type that determines how the system handles a suggestion.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/SuggestionKind-swift.enum
type CSSuggestionKind uint

const (
// CSSuggestionKindCustom - Sorts the custom suggestions together.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/SuggestionKind-swift.enum/custom
CSSuggestionKindCustom CSSuggestionKind = 0
// CSSuggestionKindDefault - Displays the suggestion normally.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/SuggestionKind-swift.enum/default
CSSuggestionKindDefault CSSuggestionKind = 0
// CSSuggestionKindNone - Blocks the system from displaying the suggestion.
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSuggestion/SuggestionKind-swift.enum/none
CSSuggestionKindNone CSSuggestionKind = 0
)

// CSUserInteraction - Constants that indicate how someone engaged with search-related content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/UserInteractionKind
type CSUserInteraction uint

const (
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/UserInteractionKind/default
CSUserInteractionDefault CSUserInteraction = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/UserInteractionKind/focus
CSUserInteractionFocus CSUserInteraction = 0
//
	// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/UserInteractionKind/select
CSUserInteractionSelect CSUserInteraction = 0
)


