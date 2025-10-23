// Code generated from Apple documentation for CoreSpotlight. DO NOT EDIT.

package corespotlight

// Enum types and constants
// CSIndexErrorCode - Error codes that describe indexing-specific errors.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSIndexError/Code
type CSIndexErrorCode uint

// CSSearchQuerySourceOptions - The query source options to allow or deny Mail messages in the search.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryContext/SourceOptions-swift.struct
type CSSearchQuerySourceOptions uint

// CSSearchQueryErrorCode - Error codes that describe reasons a query might fail.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSSearchQueryError/Code
type CSSearchQueryErrorCode uint

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

// CSUserInteraction - Constants that indicate how someone engaged with search-related content.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreSpotlight/CSUserQuery/UserInteractionKind
type CSUserInteraction uint


