// Code generated from Apple documentation for ClassKit. DO NOT EDIT.

package classkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SContext] class.
var (
	SContextClass     _SContextClass
	SContextClassOnce sync.Once
)

func getSContextClass() _SContextClass {
	SContextClassOnce.Do(func() {
		SContextClass = _SContextClass{objc.GetClass("CLSContext")}
	})
	return SContextClass
}

type _SContextClass struct {
	class objc.Class
}

// An interface definition for the [SContext] class.
type ISContext interface {
	ISObject
	AddChildContext(child unsafe.Pointer)
	AddNavigationChildContext(child unsafe.Pointer)
	AddProgressReportingCapabilities(capabilities unsafe.Pointer)
	BecomeActive()
	CreateNewActivity() unsafe.Pointer
	DescendantMatchingIdentifierPathCompletion(identifierPath unsafe.Pointer, completion unsafe.Pointer)
	RemoveFromParent()
	RemoveNavigationChildContext(child unsafe.Pointer)
	ResetProgressReportingCapabilities()
	ResignActive()
}

// An area of your app that represents an assignable task, like a quiz or a chapter.
//
// Make it easy for teachers to understand the app content a context represents by configuring it with information like a clear, concise title localized for the regions that your app supports. A context can contain groups of other contexts, like a book that contains chapters or a chapter that contains sections. You can assemble contexts into a hierarchy of up to eight levels that acts as a table of contents for teachers who want to assign your app content. See for more details.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext
type SContext struct {
	SObject
}

// SContextFrom constructs a [SContext] from an unsafe.Pointer.
//
// An area of your app that represents an assignable task, like a quiz or a chapter.
func SContextFrom(ptr unsafe.Pointer) SContext {
	return SContext{
		SObject: SObjectFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SContextClass) Alloc() SContext {
	rv := objc.Send[SContext](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SContextClass) New() SContext {
	rv := objc.Send[SContext](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SContext) Init() SContext {
	rv := objc.Send[SContext](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SContext) Autorelease() SContext {
	rv := objc.Send[SContext](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSContext creates a new SContext instance.
func NewSContext() SContext {
	return getSContextClass().New()
}




// Initializes a new context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/init(type:identifier:title:)
func NewSContextWithTypeIdentifierTitle(type_ unsafe.Pointer, identifier string, title string) SContext {
	instance := getSContextClass().Alloc()
	rv := objc.Send[SContext](instance.ID, objc.Sel("initWithType:identifier:title:"), type_, objc.String(identifier), objc.String(title))
	rv.Autorelease()
	return rv
}


// Adds the specifed context as a child of the context receiving the method call.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/addChildContext(_:)
func (s_ SContext) AddChildContext(child unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addChildContext:"), child)
}

// Adds a child context that users can navigate to from this context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/addNavigationChildContext(_:)
func (s_ SContext) AddNavigationChildContext(child unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addNavigationChildContext:"), child)
}

// Adds a progress reporting capability to the set of capabilities for the context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/addProgressReportingCapabilities(_:)
func (s_ SContext) AddProgressReportingCapabilities(capabilities unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("addProgressReportingCapabilities:"), capabilities)
}

// Tells a context to become the active context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/becomeActive()
func (s_ SContext) BecomeActive() {
	objc.Send[objc.ID](s_.ID, objc.Sel("becomeActive"))
}

// Creates and returns a new activity instance for the context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/createNewActivity()
func (s_ SContext) CreateNewActivity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("createNewActivity"))
	return rv
}

// Finds the context with the given identifier path relative to this context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/descendant(matchingIdentifierPath:completion:)
func (s_ SContext) DescendantMatchingIdentifierPathCompletion(identifierPath unsafe.Pointer, completion unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("descendantMatchingIdentifierPath:completion:"), identifierPath, completion)
}

// Removes the context from its parent.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/removeFromParent()
func (s_ SContext) RemoveFromParent() {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeFromParent"))
}

// Removes the specified context as a presentable child of this context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/removeNavigationChildContext(_:)
func (s_ SContext) RemoveNavigationChildContext(child unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("removeNavigationChildContext:"), child)
}

// Resets the set of capabilities for the context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/resetProgressReportingCapabilities()
func (s_ SContext) ResetProgressReportingCapabilities() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resetProgressReportingCapabilities"))
}

// Tells a context to stop being the active context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/resignActive()
func (s_ SContext) ResignActive() {
	objc.Send[objc.ID](s_.ID, objc.Sel("resignActive"))
}

// The activity available for recording progress.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/currentActivity
func (s_ SContext) CurrentActivity() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("currentActivity"))
	return rv
}

// An optional name that the system presents to the user if you choose the custom context type.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/customTypeName
func (s_ SContext) CustomTypeName() string {
	rv := objc.Send[string](s_.ID, objc.Sel("customTypeName"))
	return rv
}


// SetCustomTypeName sets the value of the customTypeName property.
// An optional name that the system presents to the user if you choose the custom context type.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/customTypeName
func (s_ SContext) SetCustomTypeName(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCustomTypeName:"), objc.String(value))
}

// The position of a context relative to its siblings.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/displayOrder
func (s_ SContext) DisplayOrder() int {
	rv := objc.Send[int](s_.ID, objc.Sel("displayOrder"))
	return rv
}


// SetDisplayOrder sets the value of the displayOrder property.
// The position of a context relative to its siblings.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/displayOrder
func (s_ SContext) SetDisplayOrder(value int) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDisplayOrder:"), value)
}

// A string that uniquely identifies a context among its siblings.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/identifier
func (s_ SContext) Identifier() string {
	rv := objc.Send[string](s_.ID, objc.Sel("identifier"))
	return rv
}

// The identifier path that locates the context within the data store’s context hierarchy.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/identifierPath
func (s_ SContext) IdentifierPath() []string {
	rv := objc.Send[[]string](s_.ID, objc.Sel("identifierPath"))
	return rv
}

// A Boolean indicating whether the context is active.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/isActive
func (s_ SContext) Active() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("active"))
	return rv
}

// A Boolean that indicates whether teachers can assign the context as a task.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/isAssignable
func (s_ SContext) Assignable() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("assignable"))
	return rv
}


// SetAssignable sets the value of the assignable property.
// A Boolean that indicates whether teachers can assign the context as a task.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/isAssignable
func (s_ SContext) SetAssignable(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAssignable:"), value)
}

// The child contexts that a user can navigate to from this context in the Schoolwork app.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/navigationChildContexts
func (s_ SContext) NavigationChildContexts() []SContext {
	rv := objc.Send[[]SContext](s_.ID, objc.Sel("navigationChildContexts"))
	return rv
}

// The direct ancestor of this context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/parent
func (s_ SContext) Parent() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("parent"))
	return rv
}

// The kinds of progress reporting that the context can perform.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/progressReportingCapabilities
func (s_ SContext) ProgressReportingCapabilities() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("progressReportingCapabilities"))
	return rv
}

// The range of ages, measured in years, for which you deem a context’s content suitable.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedAge
func (s_ SContext) SuggestedAge() Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("suggestedAge"))
	return rv
}


// SetSuggestedAge sets the value of the suggestedAge property.
// The range of ages, measured in years, for which you deem a context’s content suitable.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedAge
func (s_ SContext) SetSuggestedAge(value Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuggestedAge:"), value)
}

// A suggested time range to complete a task, measured in minutes.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedCompletionTime
func (s_ SContext) SuggestedCompletionTime() Range {
	rv := objc.Send[Range](s_.ID, objc.Sel("suggestedCompletionTime"))
	return rv
}


// SetSuggestedCompletionTime sets the value of the suggestedCompletionTime property.
// A suggested time range to complete a task, measured in minutes.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/suggestedCompletionTime
func (s_ SContext) SetSuggestedCompletionTime(value Range) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSuggestedCompletionTime:"), value)
}

// An optional, user-visible description of the context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/summary
func (s_ SContext) Summary() string {
	rv := objc.Send[string](s_.ID, objc.Sel("summary"))
	return rv
}


// SetSummary sets the value of the summary property.
// An optional, user-visible description of the context.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/summary
func (s_ SContext) SetSummary(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSummary:"), objc.String(value))
}

// An optional thumbnail image associated with the context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/thumbnail
func (s_ SContext) Thumbnail() CGImageRef {
	rv := objc.Send[CGImageRef](s_.ID, objc.Sel("thumbnail"))
	return rv
}


// SetThumbnail sets the value of the thumbnail property.
// An optional thumbnail image associated with the context.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/thumbnail
func (s_ SContext) SetThumbnail(value CGImageRef) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setThumbnail:"), value)
}

// The name of the context as it appears to users.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/title
func (s_ SContext) Title() string {
	rv := objc.Send[string](s_.ID, objc.Sel("title"))
	return rv
}


// SetTitle sets the value of the title property.
// The name of the context as it appears to users.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/title
func (s_ SContext) SetTitle(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The area of study to which a context relates.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/topic
func (s_ SContext) Topic() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("topic"))
	return rv
}


// SetTopic sets the value of the topic property.
// The area of study to which a context relates.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/topic
func (s_ SContext) SetTopic(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setTopic:"), value)
}

// The kind of content a context represents.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/type
func (s_ SContext) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("type"))
	return rv
}

// A URL that leads to the content in your app associated with the current context.
//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/universalLinkURL
func (s_ SContext) UniversalLinkURL() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("universalLinkURL"))
	return rv
}


// SetUniversalLinkURL sets the value of the universalLinkURL property.
// A URL that leads to the content in your app associated with the current context.

//
// [Full Topic]: https://developer.apple.com/documentation/ClassKit/CLSContext/universalLinkURL
func (s_ SContext) SetUniversalLinkURL(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setUniversalLinkURL:"), value)
}


