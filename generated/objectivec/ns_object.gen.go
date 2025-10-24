// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Object] class.
var (
	ObjectClass     _ObjectClass
	ObjectClassOnce sync.Once
)

func getObjectClass() _ObjectClass {
	ObjectClassOnce.Do(func() {
		ObjectClass = _ObjectClass{objc.GetClass("NSObject")}
	})
	return ObjectClass
}

type _ObjectClass struct {
	class objc.Class
}

// An interface definition for the [Object] class.
type IObject interface {
	
	// properties:
	AccessibilityActivateBlock() unsafe.Pointer
	SetAccessibilityActivateBlock(value unsafe.Pointer)
	AccessibilityActivationPointBlock() unsafe.Pointer
	SetAccessibilityActivationPointBlock(value unsafe.Pointer)
	AccessibilityAttributedHintBlock() unsafe.Pointer
	SetAccessibilityAttributedHintBlock(value unsafe.Pointer)
	AccessibilityAttributedLabelBlock() unsafe.Pointer
	SetAccessibilityAttributedLabelBlock(value unsafe.Pointer)
	AccessibilityAttributedUserInputLabelsBlock() unsafe.Pointer
	SetAccessibilityAttributedUserInputLabelsBlock(value unsafe.Pointer)
	AccessibilityAttributedValueBlock() unsafe.Pointer
	SetAccessibilityAttributedValueBlock(value unsafe.Pointer)
	AccessibilityContainerType() IObject /* already interface */
	SetAccessibilityContainerType(value IObject /* already interface */)
	AccessibilityContainerTypeBlock() unsafe.Pointer
	SetAccessibilityContainerTypeBlock(value unsafe.Pointer)
	AccessibilityCustomActionsBlock() unsafe.Pointer
	SetAccessibilityCustomActionsBlock(value unsafe.Pointer)
	AccessibilityCustomRotors() IObject /* already interface */
	SetAccessibilityCustomRotors(value IObject /* already interface */)
	AccessibilityCustomRotorsBlock() unsafe.Pointer
	SetAccessibilityCustomRotorsBlock(value unsafe.Pointer)
	AccessibilityDecrementBlock() unsafe.Pointer
	SetAccessibilityDecrementBlock(value unsafe.Pointer)
	AccessibilityDirectTouchOptions() unsafe.Pointer
	SetAccessibilityDirectTouchOptions(value unsafe.Pointer)
	AccessibilityElements() unsafe.Pointer
	SetAccessibilityElements(value unsafe.Pointer)
	AccessibilityElementsBlock() unsafe.Pointer
	SetAccessibilityElementsBlock(value unsafe.Pointer)
	AccessibilityElementsHidden() bool /* primitive/slice/pointer. */
	SetAccessibilityElementsHidden(value bool /* primitive/slice/pointer. */)
	AccessibilityElementsHiddenBlock() unsafe.Pointer
	SetAccessibilityElementsHiddenBlock(value unsafe.Pointer)
	AccessibilityExpandedStatus() unsafe.Pointer
	SetAccessibilityExpandedStatus(value unsafe.Pointer)
	AccessibilityExpandedStatusBlock() unsafe.Pointer
	SetAccessibilityExpandedStatusBlock(value unsafe.Pointer)
	AccessibilityFocusedUIElement() unsafe.Pointer
	SetAccessibilityFocusedUIElement(value unsafe.Pointer)
	AccessibilityFrameBlock() unsafe.Pointer
	SetAccessibilityFrameBlock(value unsafe.Pointer)
	AccessibilityHeaderElements() unsafe.Pointer
	SetAccessibilityHeaderElements(value unsafe.Pointer)
	AccessibilityHeaderElementsBlock() unsafe.Pointer
	SetAccessibilityHeaderElementsBlock(value unsafe.Pointer)
	AccessibilityHintBlock() unsafe.Pointer
	SetAccessibilityHintBlock(value unsafe.Pointer)
	AccessibilityIdentifierBlock() unsafe.Pointer
	SetAccessibilityIdentifierBlock(value unsafe.Pointer)
	AccessibilityIncrementBlock() unsafe.Pointer
	SetAccessibilityIncrementBlock(value unsafe.Pointer)
	AccessibilityLabelBlock() unsafe.Pointer
	SetAccessibilityLabelBlock(value unsafe.Pointer)
	AccessibilityLanguageBlock() unsafe.Pointer
	SetAccessibilityLanguageBlock(value unsafe.Pointer)
	AccessibilityMagicTapBlock() unsafe.Pointer
	SetAccessibilityMagicTapBlock(value unsafe.Pointer)
	AccessibilityNavigationStyle() IObject /* already interface */
	SetAccessibilityNavigationStyle(value IObject /* already interface */)
	AccessibilityNavigationStyleBlock() unsafe.Pointer
	SetAccessibilityNavigationStyleBlock(value unsafe.Pointer)
	AccessibilityNextTextNavigationElement() unsafe.Pointer
	SetAccessibilityNextTextNavigationElement(value unsafe.Pointer)
	AccessibilityNextTextNavigationElementBlock() unsafe.Pointer
	SetAccessibilityNextTextNavigationElementBlock(value unsafe.Pointer)
	AccessibilityNotifiesWhenDestroyed() bool /* primitive/slice/pointer. */
	SetAccessibilityNotifiesWhenDestroyed(value bool /* primitive/slice/pointer. */)
	AccessibilityPathBlock() unsafe.Pointer
	SetAccessibilityPathBlock(value unsafe.Pointer)
	AccessibilityPerformEscapeBlock() unsafe.Pointer
	SetAccessibilityPerformEscapeBlock(value unsafe.Pointer)
	AccessibilityPreviousTextNavigationElement() unsafe.Pointer
	SetAccessibilityPreviousTextNavigationElement(value unsafe.Pointer)
	AccessibilityPreviousTextNavigationElementBlock() unsafe.Pointer
	SetAccessibilityPreviousTextNavigationElementBlock(value unsafe.Pointer)
	AccessibilityRespondsToUserInteraction() bool /* primitive/slice/pointer. */
	SetAccessibilityRespondsToUserInteraction(value bool /* primitive/slice/pointer. */)
	AccessibilityRespondsToUserInteractionBlock() unsafe.Pointer
	SetAccessibilityRespondsToUserInteractionBlock(value unsafe.Pointer)
	AccessibilityShouldGroupAccessibilityChildrenBlock() unsafe.Pointer
	SetAccessibilityShouldGroupAccessibilityChildrenBlock(value unsafe.Pointer)
	AccessibilityTextInputResponder() IObject /* already interface */
	SetAccessibilityTextInputResponder(value IObject /* already interface */)
	AccessibilityTextInputResponderBlock() unsafe.Pointer
	SetAccessibilityTextInputResponderBlock(value unsafe.Pointer)
	AccessibilityTextualContext() IObject /* already interface */
	SetAccessibilityTextualContext(value IObject /* already interface */)
	AccessibilityTextualContextBlock() unsafe.Pointer
	SetAccessibilityTextualContextBlock(value unsafe.Pointer)
	AccessibilityTraits() IObject /* already interface */
	SetAccessibilityTraits(value IObject /* already interface */)
	AccessibilityTraitsBlock() unsafe.Pointer
	SetAccessibilityTraitsBlock(value unsafe.Pointer)
	AccessibilityUserInputLabelsBlock() unsafe.Pointer
	SetAccessibilityUserInputLabelsBlock(value unsafe.Pointer)
	AccessibilityValueBlock() unsafe.Pointer
	SetAccessibilityValueBlock(value unsafe.Pointer)
	AccessibilityViewIsModal() bool /* primitive/slice/pointer. */
	SetAccessibilityViewIsModal(value bool /* primitive/slice/pointer. */)
	AccessibilityViewIsModalBlock() unsafe.Pointer
	SetAccessibilityViewIsModalBlock(value unsafe.Pointer)
	AutoContentAccessingProxy() unsafe.Pointer
	SetAutoContentAccessingProxy(value unsafe.Pointer)
	AutomationElements() unsafe.Pointer
	SetAutomationElements(value unsafe.Pointer)
	AutomationElementsBlock() unsafe.Pointer
	SetAutomationElementsBlock(value unsafe.Pointer)
	BrowserAccessibilityContainerType() unsafe.Pointer
	SetBrowserAccessibilityContainerType(value unsafe.Pointer)
	BrowserAccessibilityHasDOMFocus() bool /* primitive/slice/pointer. */
	SetBrowserAccessibilityHasDOMFocus(value bool /* primitive/slice/pointer. */)
	BrowserAccessibilityIsRequired() bool /* primitive/slice/pointer. */
	SetBrowserAccessibilityIsRequired(value bool /* primitive/slice/pointer. */)
	BrowserAccessibilityPressedState() unsafe.Pointer
	SetBrowserAccessibilityPressedState(value unsafe.Pointer)
	ClassCode() uint32 /* not a class type */
	SetClassCode(value uint32 /* not a class type */)
	ClassForArchiver() objc.Class
	SetClassForArchiver(value objc.Class)
	ClassForCoder() objc.Class
	SetClassForCoder(value objc.Class)
	ClassForKeyedArchiver() objc.Class
	SetClassForKeyedArchiver(value objc.Class)
	IsAccessibilityElement() bool /* primitive/slice/pointer. */
	SetIsAccessibilityElement(value bool /* primitive/slice/pointer. */)
	IsAccessibilityElementBlock() unsafe.Pointer
	SetIsAccessibilityElementBlock(value unsafe.Pointer)
	IsSelectable() bool /* primitive/slice/pointer. */
	SetIsSelectable(value bool /* primitive/slice/pointer. */)
	ShouldGroupAccessibilityChildren() bool /* primitive/slice/pointer. */
	SetShouldGroupAccessibilityChildren(value bool /* primitive/slice/pointer. */)
	// methods:
	AcceptsPreviewPanelControl(panel IObject /* already interface */) bool /* primitive/slice/pointer. */
	AccessibilityActivate() bool /* primitive/slice/pointer. */
	AccessibilityArrayAttributeCount(attribute AccessibilityAttributeName /* not a class type */) uint /* primitive/slice/pointer. */
	AccessibilityAssistiveTechnologyFocusedIdentifiers() unsafe.Pointer
	AccessibilityDecrement()
	AccessibilityElementAtIndex(index int /* primitive/slice/pointer. */) objc.ID
	AccessibilityElementCount() int /* primitive/slice/pointer. */
	AccessibilityElementDidBecomeFocused()
	AccessibilityElementDidLoseFocus()
	AccessibilityElementIsFocused() bool /* primitive/slice/pointer. */
	AccessibilityHitTest(point IObject) objc.ID
	AccessibilityHitTestWithEvent(point IObject, event IObject) objc.ID
	AccessibilityIncrement()
	AccessibilityIndexOfChild(child IObject) uint /* primitive/slice/pointer. */
	AccessibilityLineEndPositionFromCurrentSelection() int /* primitive/slice/pointer. */
	AccessibilityLineStartPositionFromCurrentSelection() int /* primitive/slice/pointer. */
	AccessibilityPerformEscape() bool /* primitive/slice/pointer. */
	AccessibilityPerformMagicTap() bool /* primitive/slice/pointer. */
	AccessibilityScroll(direction IObject /* already interface */) bool /* primitive/slice/pointer. */
	AccessibilityZoomInAtPoint(point IObject) bool /* primitive/slice/pointer. */
	AccessibilityZoomOutAtPoint(point IObject) bool /* primitive/slice/pointer. */
	AttemptRecoveryFromErrorOptionIndex(error_ IObject, recoveryOptionIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	AttemptRecoveryFromErrorOptionIndexDelegateDidRecoverSelectorContextInfo(error_ IObject, recoveryOptionIndex uint /* primitive/slice/pointer. */, delegate IObject, didRecoverSelector objc.SEL, contextInfo unsafe.Pointer)
	AuthorizationViewCreatedAuthorization(view IObject)
	AuthorizationViewDidAuthorize(view IObject)
	AuthorizationViewDidDeauthorize(view IObject)
	AuthorizationViewDidHide(view IObject)
	AuthorizationViewReleasedAuthorization(view IObject)
	AuthorizationViewShouldDeauthorize(view IObject) bool /* primitive/slice/pointer. */
	BeginPreviewPanelControl(panel IObject /* already interface */)
	BindToObjectWithKeyPathOptions(binding BindingName /* not a class type */, observable IObject, keyPath IObject, options IObject)
	BrowserAccessibilityDeleteTextAtCursor(numberOfCharacters int /* primitive/slice/pointer. */)
	BrowserAccessibilityInsertTextAtCursor(text IObject)
	BrowserAccessibilitySetSelectedTextRange(range_ IObject)
	BurnProgressPanelBurnDidFinish(theBurnPanel unsafe.Pointer, burn unsafe.Pointer) bool /* primitive/slice/pointer. */
	BurnProgressPanelDidFinish(aNotification IObject)
	BurnProgressPanelWillBegin(aNotification IObject)
	CertificatePanelShowHelp(sender IObject) bool /* primitive/slice/pointer. */
	ChooseIdentityPanelShowHelp(sender IObject) bool /* primitive/slice/pointer. */
	CoerceValueForKey(value IObject, key IObject) objc.ID
	CommitComposition(sender IObject)
	ComposedString(sender IObject) objc.ID
	CopyScriptingValueForKeyWithProperties(value IObject, key IObject, properties IObject) objc.ID
	DidChangeValueForKeyWithSetMutationUsingObjects(key IObject, mutationKind KeyValueSetMutationKind /* not a class type */, objects IObject)
	DidCommandBySelectorClient(aSelector objc.SEL, sender IObject) bool /* primitive/slice/pointer. */
	DoesContain(object IObject) bool /* primitive/slice/pointer. */
	DoesNotRecognizeSelector(aSelector objc.SEL)
	EndPreviewPanelControl(panel IObject /* already interface */)
	EraseProgressPanelEraseDidFinish(theErasePanel unsafe.Pointer, erase unsafe.Pointer) bool /* primitive/slice/pointer. */
	EraseProgressPanelDidFinish(aNotification IObject)
	EraseProgressPanelWillBegin(aNotification IObject)
	ExceptionHandlerShouldHandleExceptionMask(sender IObject /* already interface */, exception IObject, aMask uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	ExceptionHandlerShouldLogExceptionMask(sender IObject /* already interface */, exception IObject, aMask uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	FileTransferServicesAbortCompleteError(inServices IObject, inError unsafe.Pointer)
	FileTransferServicesConnectionCompleteError(inServices IObject, inError unsafe.Pointer)
	FileTransferServicesCopyRemoteFileCompleteError(inServices IObject, inError unsafe.Pointer)
	FileTransferServicesCopyRemoteFileProgressTransferProgress(inServices IObject, inProgressDescription IObject)
	FileTransferServicesCreateFolderCompleteErrorFolder(inServices IObject, inError unsafe.Pointer, inFolderName IObject)
	FileTransferServicesDisconnectionCompleteError(inServices IObject, inError unsafe.Pointer)
	FileTransferServicesFilePreparationCompleteError(inServices IObject, inError unsafe.Pointer)
	FileTransferServicesPathChangeCompleteErrorFinalPath(inServices IObject, inError unsafe.Pointer, inPath IObject)
	FileTransferServicesRemoveItemCompleteErrorRemovedItem(inServices IObject, inError unsafe.Pointer, inItemName IObject)
	FileTransferServicesRetrieveFolderListingCompleteErrorListing(inServices IObject, inError unsafe.Pointer, inListing IObject)
	FileTransferServicesSendFileCompleteError(inServices IObject, inError unsafe.Pointer)
	FileTransferServicesSendFileProgressTransferProgress(inServices IObject, inProgressDescription IObject)
	Finalize()
	FinalizeForWebScript()
	ForwardingTargetForSelector(aSelector objc.SEL) objc.ID
	HandleEventClient(event IObject, sender IObject) bool /* primitive/slice/pointer. */
	ImageBrowserBackgroundWasRightClickedWithEvent(aBrowser IObject, event IObject)
	ImageBrowserCellWasDoubleClickedAtIndex(aBrowser IObject, index uint /* primitive/slice/pointer. */)
	ImageBrowserCellWasRightClickedAtIndexWithEvent(aBrowser IObject, index uint /* primitive/slice/pointer. */, event IObject)
	ImageBrowserItemAtIndex(aBrowser IObject, index uint /* primitive/slice/pointer. */) objc.ID
	ImageBrowserMoveItemsAtIndexesToIndex(aBrowser IObject, indexes IObject, destinationIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */
	ImageBrowserRemoveItemsAtIndexes(aBrowser IObject, indexes IObject)
	ImageBrowserWriteItemsAtIndexesToPasteboard(aBrowser IObject, itemIndexes IObject, pasteboard IObject) uint /* primitive/slice/pointer. */
	ImageBrowserSelectionDidChange(aBrowser IObject)
	ImageRepresentation() objc.ID
	ImageVersion() uint /* primitive/slice/pointer. */
	IndexOfAccessibilityElement(element IObject) int /* primitive/slice/pointer. */
	IndicesOfObjectsByEvaluatingObjectSpecifier(specifier IObject) IObject
	InputTextClient(string_ IObject, sender IObject) bool /* primitive/slice/pointer. */
	InputTextKeyModifiersClient(string_ IObject, keyCode int /* primitive/slice/pointer. */, flags uint /* primitive/slice/pointer. */, sender IObject) bool /* primitive/slice/pointer. */
	InsertValueAtIndexInPropertyWithKey(value IObject, index uint /* primitive/slice/pointer. */, key IObject)
	InsertValueInPropertyWithKey(value IObject, key IObject)
	InvokeDefaultMethodWithArguments(arguments IObject) objc.ID
	InvokeUndefinedMethodFromWebScriptWithArguments(name IObject, arguments IObject) objc.ID
	IsCaseInsensitiveLike(object IObject) bool /* primitive/slice/pointer. */
	IsGreaterThan(object IObject) bool /* primitive/slice/pointer. */
	IsGreaterThanOrEqualTo(object IObject) bool /* primitive/slice/pointer. */
	IsLessThan(object IObject) bool /* primitive/slice/pointer. */
	IsLessThanOrEqualTo(object IObject) bool /* primitive/slice/pointer. */
	IsLike(object IObject) bool /* primitive/slice/pointer. */
	IsNotEqualTo(object IObject) bool /* primitive/slice/pointer. */
	MethodForSelector(aSelector objc.SEL) unsafe.Pointer
	NewScriptingObjectOfClassForValueForKeyWithContentsValueProperties(objectClass objc.Class, key IObject, contentsValue IObject, properties IObject) objc.ID
	NumberOfGroupsInImageBrowser(aBrowser IObject) uint /* primitive/slice/pointer. */
	NumberOfItemsInImageBrowser(aBrowser IObject) uint /* primitive/slice/pointer. */
	OptionDescriptionsForBinding(binding BindingName /* not a class type */) IObject
	PerformSelectorOnThreadWithObjectWaitUntilDone(aSelector objc.SEL, thr IObject, arg IObject, wait bool /* primitive/slice/pointer. */)
	PerformSelectorOnThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, thr IObject, arg IObject, wait bool /* primitive/slice/pointer. */, array []string /* primitive/slice/pointer. */)
	PerformSelectorWithObjectAfterDelay(aSelector objc.SEL, anArgument IObject, delay TimeInterval /* not a class type */)
	PerformSelectorWithObjectAfterDelayInModes(aSelector objc.SEL, anArgument IObject, delay TimeInterval /* not a class type */, modes []string /* primitive/slice/pointer. */)
	PerformActionForPersonIdentifier(person IObject, identifier IObject)
	PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, arg IObject, wait bool /* primitive/slice/pointer. */, array []string /* primitive/slice/pointer. */)
	ProvideImageToMTLTextureCommandBufferOriginxOriginyWidthHeightUserInfo(texture IObject, commandBuffer IObject, originx uintptr /* not a class type */, originy uintptr /* not a class type */, width uintptr /* not a class type */, height uintptr /* not a class type */, info IObject)
	ProvideImageDataBytesPerRowOriginSizeUserInfo(data unsafe.Pointer, rowbytes uintptr /* not a class type */, originx uintptr /* not a class type */, originy uintptr /* not a class type */, width uintptr /* not a class type */, height uintptr /* not a class type */, info IObject)
	QuartzFilterManagerDidAddFilter(sender IObject, filter IObject)
	QuartzFilterManagerDidModifyFilter(sender IObject, filter IObject)
	QuartzFilterManagerDidRemoveFilter(sender IObject, filter IObject)
	QuartzFilterManagerDidSelectFilter(sender IObject, filter IObject)
	ReadLinkQualityForDeviceCompleteDeviceInfoError(controller IObject, device IObject /* already interface */, info unsafe.Pointer, error_ Return /* not a class type */)
	ReadRSSIForDeviceCompleteDeviceInfoError(controller IObject, device IObject /* already interface */, info unsafe.Pointer, error_ Return /* not a class type */)
	RemoveValueAtIndexFromPropertyWithKey(index uint /* primitive/slice/pointer. */, key IObject)
	ReplaceValueAtIndexInPropertyWithKeyWithValue(index uint /* primitive/slice/pointer. */, key IObject, value IObject)
	ReplacementObjectForCoder(coder IObject) objc.ID
	ReplacementObjectForKeyedArchiver(archiver IObject) objc.ID
	SaveOptionsShouldShowUTType(saveOptions IObject, utType IObject) bool /* primitive/slice/pointer. */
	ScriptingBeginsWith(object IObject) bool /* primitive/slice/pointer. */
	ScriptingContains(object IObject) bool /* primitive/slice/pointer. */
	ScriptingEndsWith(object IObject) bool /* primitive/slice/pointer. */
	ScriptingIsEqualTo(object IObject) bool /* primitive/slice/pointer. */
	ScriptingIsGreaterThan(object IObject) bool /* primitive/slice/pointer. */
	ScriptingIsGreaterThanOrEqualTo(object IObject) bool /* primitive/slice/pointer. */
	ScriptingIsLessThan(object IObject) bool /* primitive/slice/pointer. */
	ScriptingIsLessThanOrEqualTo(object IObject) bool /* primitive/slice/pointer. */
	ScriptingValueForSpecifier(objectSpecifier IObject) objc.ID
	SetSharedObservers(sharedObservers IObject)
	SetupPanelDetermineBestDeviceOfAOrB(aPanel unsafe.Pointer, deviceA unsafe.Pointer, device unsafe.Pointer) unsafe.Pointer
	SetupPanelDeviceContainsSuitableMediaPromptString(aPanel unsafe.Pointer, device unsafe.Pointer, prompt IObject) bool /* primitive/slice/pointer. */
	SetupPanelDeviceCouldBeTarget(aPanel unsafe.Pointer, device unsafe.Pointer) bool /* primitive/slice/pointer. */
	SetupPanelDeviceSelectionChanged(aNotification IObject)
	SetupPanelShouldHandleMediaReservations(aPanel unsafe.Pointer) bool /* primitive/slice/pointer. */
	ShouldEnableActionForPersonIdentifier(person IObject, identifier IObject) bool /* primitive/slice/pointer. */
	Unbind(binding BindingName /* not a class type */)
	ValidateValueForKeyError(ioValue unsafe.Pointer, inKey IObject, outError unsafe.Pointer) bool /* primitive/slice/pointer. */
	ValidateValueForKeyPathError(ioValue unsafe.Pointer, inKeyPath IObject, outError unsafe.Pointer) bool /* primitive/slice/pointer. */
	ValueAtIndexInPropertyWithKey(index uint /* primitive/slice/pointer. */, key IObject) objc.ID
	ValueWithNameInPropertyWithKey(name IObject, key IObject) objc.ID
	ValueWithUniqueIDInPropertyWithKey(uniqueID IObject, key IObject) objc.ID
	ValueClassForBinding(binding BindingName /* not a class type */) objc.Class
	WebPlugInContainerLoadRequestInFrame(request IObject, target IObject)
	WebPlugInContainerShowStatus(message IObject)
	WebPlugInDestroy()
	WebPlugInInitialize()
	WebPlugInMainResourceDidFailWithError(error_ IObject)
	WebPlugInMainResourceDidFinishLoading()
	WebPlugInMainResourceDidReceiveData(data IObject)
	WebPlugInMainResourceDidReceiveResponse(response IObject)
	WebPlugInSetIsSelected(isSelected bool /* primitive/slice/pointer. */)
	WebPlugInStart()
	WebPlugInStop()
	WillChangeValueForKeyWithSetMutationUsingObjects(key IObject, mutationKind KeyValueSetMutationKind /* not a class type */, objects IObject)
}

// The root class of most Objective-C class hierarchies, from which subclasses inherit a basic interface to the runtime system and the ability to behave as Objective-C objects.


// The root class of most Objective-C class hierarchies, from which subclasses inherit a basic interface to the runtime system and the ability to behave as Objective-C objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
type Object struct {
	objc.ID
}

// ObjectFrom constructs a [Object] from an unsafe.Pointer.
//
// The root class of most Objective-C class hierarchies, from which subclasses inherit a basic interface to the runtime system and the ability to behave as Objective-C objects.
func ObjectFrom(ptr unsafe.Pointer) Object {
	return Object{objc.ID(ptr)}
}

// Alloc allocates a new instance without initialization.
func (oc _ObjectClass) Alloc() Object {
	rv := objc.Send[Object](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _ObjectClass) New() Object {
	rv := objc.Send[Object](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ Object) Init() Object {
	rv := objc.Send[Object](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ Object) Autorelease() Object {
	rv := objc.Send[Object](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewObject creates a new Object instance.
func NewObject() Object {
	return getObjectClass().New()
}




// Returns a Boolean value that indicates whether the observed object supports automatic key-value observation for the given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/automaticallyNotifiesObservers(forKey:)
func (oc _ObjectClass) AutomaticallyNotifiesObserversForKey(key IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("automaticallyNotifiesObserversForKey:"), key)
	return rv
}


// Cancels perform requests previously registered with the instance method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/cancelPreviousPerformRequests(withTarget:)
func (oc _ObjectClass) CancelPreviousPerformRequestsWithTarget(aTarget IObject) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("cancelPreviousPerformRequestsWithTarget:"), aTarget)
}


// Cancels perform requests previously registered with .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/cancelPreviousPerformRequests(withTarget:selector:object:)
func (oc _ObjectClass) CancelPreviousPerformRequestsWithTargetSelectorObject(aTarget IObject, aSelector objc.SEL, anArgument IObject) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("cancelPreviousPerformRequestsWithTarget:selector:object:"), aTarget, aSelector, anArgument)
}


// Overridden to return the names of classes that can be used to decode objects if their class is unavailable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classFallbacksForKeyedArchiver()
func (oc _ObjectClass) ClassFallbacksForKeyedArchiver() []string /* primitive/slice/pointer. */ {
	rv := objc.Send[[]string](objc.ID(oc.class), objc.Sel("classFallbacksForKeyedArchiver"))
	return rv
}


// Overridden by subclasses to substitute a new class during keyed unarchiving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classForKeyedUnarchiver()
func (oc _ObjectClass) ClassForKeyedUnarchiver() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(oc.class), objc.Sel("classForKeyedUnarchiver"))
	return rv
}


// Returns a Boolean value that indicates whether the target conforms to a given protocol.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/conforms(to:)
func (oc _ObjectClass) ConformsToProtocol(protocol_ objc.IObject /* cross-framework: Protocol */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("conformsToProtocol:"), protocol_)
	return rv
}


// Returns an object that will be used as the placeholder for the , when a key value coding compliant property of an instance of the receiving class returns the value specified by , and no other placeholder has been specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/defaultPlaceholder(for:with:)
func (oc _ObjectClass) DefaultPlaceholderForMarkerWithBinding(marker IObject, binding BindingName /* not a class type */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("defaultPlaceholderForMarker:withBinding:"), marker, binding)
	return rv
}


// Exposes the specified , advertising its availability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/exposeBinding(_:)
func (oc _ObjectClass) ExposeBinding(binding BindingName /* not a class type */) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("exposeBinding:"), binding)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/hash()
func (oc _ObjectClass) Hash() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](objc.ID(oc.class), objc.Sel("hash"))
	return rv
}


// Initializes the class before it receives its first message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/initialize()
func (oc _ObjectClass) Initialize() {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("initialize"))
}


// Locates and returns the address of the implementation of the instance method identified by a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instanceMethod(for:)
func (oc _ObjectClass) InstanceMethodForSelector(aSelector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("instanceMethodForSelector:"), aSelector)
	return rv
}


// Returns a Boolean value that indicates whether instances of the receiver are capable of responding to a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instancesRespond(to:)
func (oc _ObjectClass) InstancesRespondToSelector(aSelector objc.SEL) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("instancesRespondToSelector:"), aSelector)
	return rv
}


// Returns whether a key should be hidden from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isKeyExcluded(fromWebScript:)
func (oc _ObjectClass) IsKeyExcludedFromWebScript(name unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("isKeyExcludedFromWebScript:"), name)
	return rv
}


// Returns whether a selector should be hidden from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isSelectorExcluded(fromWebScript:)
func (oc _ObjectClass) IsSelectorExcludedFromWebScript(selector objc.SEL) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("isSelectorExcludedFromWebScript:"), selector)
	return rv
}


// Returns a Boolean value that indicates whether the receiving class is a subclass of, or identical to, a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isSubclass(of:)
func (oc _ObjectClass) IsSubclassOfClass(aClass objc.Class) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("isSubclassOfClass:"), aClass)
	return rv
}


// Returns a set of key paths for properties whose values affect the value of the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/keyPathsForValuesAffectingValue(forKey:)
func (oc _ObjectClass) KeyPathsForValuesAffectingValueForKey(key IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(oc.class), objc.Sel("keyPathsForValuesAffectingValueForKey:"), key)
	return rv
}


// Invoked whenever a class or category is added to the Objective-C runtime; implement this method to perform class-specific behavior upon loading.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/load()
func (oc _ObjectClass) Load() {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("load"))
}


// Dynamically provides an implementation for a given selector for a class method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/resolveClassMethod(_:)
func (oc _ObjectClass) ResolveClassMethod(sel objc.SEL) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("resolveClassMethod:"), sel)
	return rv
}


// Dynamically provides an implementation for a given selector for an instance method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/resolveInstanceMethod(_:)
func (oc _ObjectClass) ResolveInstanceMethod(sel objc.SEL) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("resolveInstanceMethod:"), sel)
	return rv
}


// Sets as the default placeholder for the , when a key value coding compliant property of an instance of the receiving class returns the value specified by , and no other placeholder has been specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setDefaultPlaceholder(_:for:with:)
func (oc _ObjectClass) SetDefaultPlaceholderForMarkerWithBinding(placeholder IObject, marker IObject, binding BindingName /* not a class type */) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("setDefaultPlaceholder:forMarker:withBinding:"), placeholder, marker, binding)
}


// Sets the receiver’s version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setVersion(_:)
func (oc _ObjectClass) SetVersion(aVersion int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("setVersion:"), aVersion)
}


// Returns the class object for the receiver’s superclass.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/superclass()
func (oc _ObjectClass) Superclass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(oc.class), objc.Sel("superclass"))
	return rv
}


// Returns if the stored value methods and should use private accessor methods in preference to public accessors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/useStoredAccessor()
func (oc _ObjectClass) UseStoredAccessor() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("useStoredAccessor"))
	return rv
}


// Returns the version number assigned to the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/version()
func (oc _ObjectClass) Version() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](objc.ID(oc.class), objc.Sel("version"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/acceptsPreviewPanelControl(_:)
func (o_ Object) AcceptsPreviewPanelControl(panel IObject /* already interface */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("acceptsPreviewPanelControl:"), panel)
	return rv
}


// Tells the element to activate itself and report the success or failure of the operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityActivate()
func (o_ Object) AccessibilityActivate() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityActivate"))
	return rv
}


// Returns the count of the specified accessibility array attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityArrayAttributeCount(_:)
func (o_ Object) AccessibilityArrayAttributeCount(attribute AccessibilityAttributeName /* not a class type */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](o_.ID, objc.Sel("accessibilityArrayAttributeCount:"), attribute)
	return rv
}


// Returns a set of identifier keys indicating which assistive app has focus on the accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityAssistiveTechnologyFocusedIdentifiers()
func (o_ Object) AccessibilityAssistiveTechnologyFocusedIdentifiers() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAssistiveTechnologyFocusedIdentifiers"))
	return rv
}


// Tells the accessibility element to decrement the value of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityDecrement()
func (o_ Object) AccessibilityDecrement() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityDecrement"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElement(at:)
func (o_ Object) AccessibilityElementAtIndex(index int /* primitive/slice/pointer. */) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityElementAtIndex:"), index)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementCount()
func (o_ Object) AccessibilityElementCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](o_.ID, objc.Sel("accessibilityElementCount"))
	return rv
}


// Sent after an assistive technology has set its virtual focus on the accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementDidBecomeFocused()
func (o_ Object) AccessibilityElementDidBecomeFocused() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityElementDidBecomeFocused"))
}


// Sent after an assistive technology has removed its virtual focus from an accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementDidLoseFocus()
func (o_ Object) AccessibilityElementDidLoseFocus() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityElementDidLoseFocus"))
}


// Returns a Boolean value indicating whether an assistive technology is focused on the accessibility element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityElementIsFocused()
func (o_ Object) AccessibilityElementIsFocused() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityElementIsFocused"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHitTest(_:)
func (o_ Object) AccessibilityHitTest(point IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityHitTest:"), point)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHitTest(_:event:)
func (o_ Object) AccessibilityHitTestWithEvent(point IObject, event IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityHitTest:withEvent:"), point, event)
	return rv
}


// Tells the accessibility element to increment the value of its content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityIncrement()
func (o_ Object) AccessibilityIncrement() {
	objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityIncrement"))
}


// Returns the index of the specified accessibility child in the parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityIndex(ofChild:)
func (o_ Object) AccessibilityIndexOfChild(child IObject) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](o_.ID, objc.Sel("accessibilityIndexOfChild:"), child)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineEndPositionFromCurrentSelection()
func (o_ Object) AccessibilityLineEndPositionFromCurrentSelection() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](o_.ID, objc.Sel("accessibilityLineEndPositionFromCurrentSelection"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineStartPositionFromCurrentSelection()
func (o_ Object) AccessibilityLineStartPositionFromCurrentSelection() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](o_.ID, objc.Sel("accessibilityLineStartPositionFromCurrentSelection"))
	return rv
}


// Dismisses a modal view and returns the success or failure of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPerformEscape()
func (o_ Object) AccessibilityPerformEscape() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityPerformEscape"))
	return rv
}


// Performs a salient action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityPerformMagicTap()
func (o_ Object) AccessibilityPerformMagicTap() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityPerformMagicTap"))
	return rv
}


// Scrolls screen content in an application-specific way and returns the success or failure of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityScroll(_:)
func (o_ Object) AccessibilityScroll(direction IObject /* already interface */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityScroll:"), direction)
	return rv
}


// Zooms in on the content at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityZoomIn(at:)
func (o_ Object) AccessibilityZoomInAtPoint(point IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityZoomInAtPoint:"), point)
	return rv
}


// Zooms out from the content at the specified point.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityZoomOut(at:)
func (o_ Object) AccessibilityZoomOutAtPoint(point IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityZoomOutAtPoint:"), point)
	return rv
}


// Registers the observer object to receive KVO notifications for the key path relative to the object receiving this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/addObserver(_:forKeyPath:options:context:)
func (o_ Object) AddObserverForKeyPathOptionsContext(observer IObject, keyPath IObject, options KeyValueObservingOptions /* not a class type */, context unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addObserver:forKeyPath:options:context:"), observer, keyPath, options, context)
}


// Implemented to attempt a recovery from an error noted in an application-modal dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attemptRecovery(fromError:optionIndex:)
func (o_ Object) AttemptRecoveryFromErrorOptionIndex(error_ IObject, recoveryOptionIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("attemptRecoveryFromError:optionIndex:"), error_, recoveryOptionIndex)
	return rv
}


// Implemented to attempt a recovery from an error noted in a document-modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attemptRecovery(fromError:optionIndex:delegate:didRecoverSelector:contextInfo:)
func (o_ Object) AttemptRecoveryFromErrorOptionIndexDelegateDidRecoverSelectorContextInfo(error_ IObject, recoveryOptionIndex uint /* primitive/slice/pointer. */, delegate IObject, didRecoverSelector objc.SEL, contextInfo unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("attemptRecoveryFromError:optionIndex:delegate:didRecoverSelector:contextInfo:"), error_, recoveryOptionIndex, delegate, didRecoverSelector, contextInfo)
}


// Sent to the delegate to indicate the authorization object has been created or changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewCreatedAuthorization(_:)
func (o_ Object) AuthorizationViewCreatedAuthorization(view IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewCreatedAuthorization:"), view)
}


// Sent to the delegate to indicate the user was authorized and the authorization view was changed to unlocked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewDidAuthorize(_:)
func (o_ Object) AuthorizationViewDidAuthorize(view IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewDidAuthorize:"), view)
}


// Sent to the delegate to indicate the user was deauthorized and the authorization view was changed to locked.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewDidDeauthorize(_:)
func (o_ Object) AuthorizationViewDidDeauthorize(view IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewDidDeauthorize:"), view)
}


// Sent to the delegate to indicate that the view’s visibility has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewDidHide(_:)
func (o_ Object) AuthorizationViewDidHide(view IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewDidHide:"), view)
}


// Sent to the delegate to indicate that deauthorization is about to occur.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewReleasedAuthorization(_:)
func (o_ Object) AuthorizationViewReleasedAuthorization(view IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("authorizationViewReleasedAuthorization:"), view)
}


// Sent to the delegate when a user clicks the open lock icon.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/authorizationViewShouldDeauthorize(_:)
func (o_ Object) AuthorizationViewShouldDeauthorize(view IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("authorizationViewShouldDeauthorize:"), view)
	return rv
}


// Overridden by subclasses to substitute another object in place of the object that was decoded and subsequently received this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/awakeAfter(using:)
func (o_ Object) AwakeAfterUsingCoder(coder IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("awakeAfterUsingCoder:"), coder)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/beginPreviewPanelControl(_:)
func (o_ Object) BeginPreviewPanelControl(panel IObject /* already interface */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("beginPreviewPanelControl:"), panel)
}


// Establishes a binding between a given property of the receiver and the property of a given object specified by a given key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/bind(_:to:withKeyPath:options:)
func (o_ Object) BindToObjectWithKeyPathOptions(binding BindingName /* not a class type */, observable IObject, keyPath IObject, options IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("bind:toObject:withKeyPath:options:"), binding, observable, keyPath, options)
}


// Deletes text from the element at the current cursor position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityDeleteTextAtCursor(numberOfCharacters:)
func (o_ Object) BrowserAccessibilityDeleteTextAtCursor(numberOfCharacters int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilityDeleteTextAtCursor:"), numberOfCharacters)
}


// Inserts text into the element at the current cursor position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityInsertTextAtCursor(text:)
func (o_ Object) BrowserAccessibilityInsertTextAtCursor(text IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilityInsertTextAtCursor:"), text)
}


// Updates the element’s selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilitySetSelectedTextRange(_:)
func (o_ Object) BrowserAccessibilitySetSelectedTextRange(range_ IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilitySetSelectedTextRange:"), range_)
}


// Allows the delegate to handle the end-of-burn feedback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/burnProgressPanel(_:burnDidFinish:)
func (o_ Object) BurnProgressPanelBurnDidFinish(theBurnPanel unsafe.Pointer, burn unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("burnProgressPanel:burnDidFinish:"), theBurnPanel, burn)
	return rv
}


// Notification sent by the panel after ordering out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/burnProgressPanelDidFinish(_:)
func (o_ Object) BurnProgressPanelDidFinish(aNotification IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("burnProgressPanelDidFinish:"), aNotification)
}


// Notification sent by the panel before display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/burnProgressPanelWillBegin(_:)
func (o_ Object) BurnProgressPanelWillBegin(aNotification IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("burnProgressPanelWillBegin:"), aNotification)
}


// Implements custom help behavior for the modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/certificatePanelShowHelp(_:)
func (o_ Object) CertificatePanelShowHelp(sender IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("certificatePanelShowHelp:"), sender)
	return rv
}


// Implements custom help behavior for the modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/chooseIdentityPanelShowHelp(_:)
func (o_ Object) ChooseIdentityPanelShowHelp(sender IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("chooseIdentityPanelShowHelp:"), sender)
	return rv
}


// Uses type info from the class description and to attempt to convert for to the proper type, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/coerceValue(_:forKey:)
func (o_ Object) CoerceValueForKey(value IObject, key IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("coerceValue:forKey:"), value, key)
	return rv
}


// Informs the controller that the composition should be committed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/commitComposition(_:)
func (o_ Object) CommitComposition(sender IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("commitComposition:"), sender)
}


// Return the current composed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/composedString(_:)
func (o_ Object) ComposedString(sender IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("composedString:"), sender)
	return rv
}


// Creates and returns one or more scripting objects to be inserted into the specified relationship by copying the passed-in value and setting the properties in the copied object or objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyScriptingValue(_:forKey:withProperties:)
func (o_ Object) CopyScriptingValueForKeyWithProperties(value IObject, key IObject, properties IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("copyScriptingValue:forKey:withProperties:"), value, key, properties)
	return rv
}


// Informs the observed object that the specified change has occurred on the indexes for a specified ordered to-many relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/didChange(_:valuesAt:forKey:)
func (o_ Object) DidChangeValuesAtIndexesForKey(changeKind KeyValueChange /* not a class type */, indexes IObject, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("didChange:valuesAtIndexes:forKey:"), changeKind, indexes, key)
}


// Informs the observed object that the value of a given property has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/didChangeValue(forKey:)
func (o_ Object) DidChangeValueForKey(key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("didChangeValueForKey:"), key)
}


// Informs the observed object that the specified change was made to a specified unordered to-many relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/didChangeValue(forKey:withSetMutation:using:)
func (o_ Object) DidChangeValueForKeyWithSetMutationUsingObjects(key IObject, mutationKind KeyValueSetMutationKind /* not a class type */, objects IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("didChangeValueForKey:withSetMutation:usingObjects:"), key, mutationKind, objects)
}


// Processes a command generated by user action such as typing certain keys or pressing the mouse button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/didCommand(by:client:)
func (o_ Object) DidCommandBySelectorClient(aSelector objc.SEL, sender IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("didCommandBySelector:client:"), aSelector, sender)
	return rv
}


// Returns a Boolean value that indicates whether the receiver contains a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/doesContain(_:)
func (o_ Object) DoesContain(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("doesContain:"), object)
	return rv
}


// Handles messages the receiver doesn’t recognize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/doesNotRecognizeSelector(_:)
func (o_ Object) DoesNotRecognizeSelector(aSelector objc.SEL) {
	objc.Send[objc.ID](o_.ID, objc.Sel("doesNotRecognizeSelector:"), aSelector)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/endPreviewPanelControl(_:)
func (o_ Object) EndPreviewPanelControl(panel IObject /* already interface */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("endPreviewPanelControl:"), panel)
}


// Notification sent by the panel before display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/eraseProgressPanel(_:eraseDidFinish:)
func (o_ Object) EraseProgressPanelEraseDidFinish(theErasePanel unsafe.Pointer, erase unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("eraseProgressPanel:eraseDidFinish:"), theErasePanel, erase)
	return rv
}


// Notification sent by the panel after ordering out.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/eraseProgressPanelDidFinish(_:)
func (o_ Object) EraseProgressPanelDidFinish(aNotification IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("eraseProgressPanelDidFinish:"), aNotification)
}


// Notification sent by the panel before display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/eraseProgressPanelWillBegin(_:)
func (o_ Object) EraseProgressPanelWillBegin(aNotification IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("eraseProgressPanelWillBegin:"), aNotification)
}


// Implemented by the delegate to evaluate whether the delegating exception handler should handle a given exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/exceptionHandler(_:shouldHandle:mask:)
func (o_ Object) ExceptionHandlerShouldHandleExceptionMask(sender IObject /* already interface */, exception IObject, aMask uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("exceptionHandler:shouldHandleException:mask:"), sender, exception, aMask)
	return rv
}


// Implemented by the delegate to evaluate whether the delegating exception hangler should log a given exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/exceptionHandler(_:shouldLogException:mask:)
func (o_ Object) ExceptionHandlerShouldLogExceptionMask(sender IObject /* already interface */, exception IObject, aMask uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("exceptionHandler:shouldLogException:mask:"), sender, exception, aMask)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesAbortComplete(_:error:)
func (o_ Object) FileTransferServicesAbortCompleteError(inServices IObject, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesAbortComplete:error:"), inServices, inError)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesConnectionComplete(_:error:)
func (o_ Object) FileTransferServicesConnectionCompleteError(inServices IObject, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesConnectionComplete:error:"), inServices, inError)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesCopyRemoteFileComplete(_:error:)
func (o_ Object) FileTransferServicesCopyRemoteFileCompleteError(inServices IObject, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesCopyRemoteFileComplete:error:"), inServices, inError)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesCopyRemoteFileProgress(_:transferProgress:)
func (o_ Object) FileTransferServicesCopyRemoteFileProgressTransferProgress(inServices IObject, inProgressDescription IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesCopyRemoteFileProgress:transferProgress:"), inServices, inProgressDescription)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesCreateFolderComplete(_:error:folder:)
func (o_ Object) FileTransferServicesCreateFolderCompleteErrorFolder(inServices IObject, inError unsafe.Pointer, inFolderName IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesCreateFolderComplete:error:folder:"), inServices, inError, inFolderName)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesDisconnectionComplete(_:error:)
func (o_ Object) FileTransferServicesDisconnectionCompleteError(inServices IObject, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesDisconnectionComplete:error:"), inServices, inError)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesFilePreparationComplete(_:error:)
func (o_ Object) FileTransferServicesFilePreparationCompleteError(inServices IObject, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesFilePreparationComplete:error:"), inServices, inError)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesPathChangeComplete(_:error:finalPath:)
func (o_ Object) FileTransferServicesPathChangeCompleteErrorFinalPath(inServices IObject, inError unsafe.Pointer, inPath IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesPathChangeComplete:error:finalPath:"), inServices, inError, inPath)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesRemoveItemComplete(_:error:removedItem:)
func (o_ Object) FileTransferServicesRemoveItemCompleteErrorRemovedItem(inServices IObject, inError unsafe.Pointer, inItemName IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesRemoveItemComplete:error:removedItem:"), inServices, inError, inItemName)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesRetrieveFolderListingComplete(_:error:listing:)
func (o_ Object) FileTransferServicesRetrieveFolderListingCompleteErrorListing(inServices IObject, inError unsafe.Pointer, inListing IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesRetrieveFolderListingComplete:error:listing:"), inServices, inError, inListing)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesSendFileComplete(_:error:)
func (o_ Object) FileTransferServicesSendFileCompleteError(inServices IObject, inError unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesSendFileComplete:error:"), inServices, inError)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/fileTransferServicesSendFileProgress(_:transferProgress:)
func (o_ Object) FileTransferServicesSendFileProgressTransferProgress(inServices IObject, inProgressDescription IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("fileTransferServicesSendFileProgress:transferProgress:"), inServices, inProgressDescription)
}


// The garbage collector invokes this method on the receiver before disposing of the memory it uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/finalize()
func (o_ Object) Finalize() {
	objc.Send[objc.ID](o_.ID, objc.Sel("finalize"))
}


// Performs cleanup when the scripting environment is reset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/finalizeForWebScript()
func (o_ Object) FinalizeForWebScript() {
	objc.Send[objc.ID](o_.ID, objc.Sel("finalizeForWebScript"))
}


// Returns the object to which unrecognized messages should first be directed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/forwardingTarget(for:)
func (o_ Object) ForwardingTargetForSelector(aSelector objc.SEL) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("forwardingTargetForSelector:"), aSelector)
	return rv
}


// Handles key down and mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/handle(_:client:)
func (o_ Object) HandleEventClient(event IObject, sender IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("handleEvent:client:"), event, sender)
	return rv
}


// Performs custom tasks when the user right-clicks the image browser view background.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:backgroundWasRightClickedWith:)
func (o_ Object) ImageBrowserBackgroundWasRightClickedWithEvent(aBrowser IObject, event IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:backgroundWasRightClickedWithEvent:"), aBrowser, event)
}


// Performs custom tasks when the user double-clicks an item in the image browser view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:cellWasDoubleClickedAt:)
func (o_ Object) ImageBrowserCellWasDoubleClickedAtIndex(aBrowser IObject, index uint /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:cellWasDoubleClickedAtIndex:"), aBrowser, index)
}


// Performs custom tasks when the user right-clicks an item in the image browser view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:cellWasRightClickedAt:with:)
func (o_ Object) ImageBrowserCellWasRightClickedAtIndexWithEvent(aBrowser IObject, index uint /* primitive/slice/pointer. */, event IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:cellWasRightClickedAtIndex:withEvent:"), aBrowser, index, event)
}


// Returns an object for the item in an image browser view that corresponds to the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:itemAt:)
func (o_ Object) ImageBrowserItemAtIndex(aBrowser IObject, index uint /* primitive/slice/pointer. */) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:itemAtIndex:"), aBrowser, index)
	return rv
}


// Signals that the specified items should be moved to the specified destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:moveItemsAt:to:)
func (o_ Object) ImageBrowserMoveItemsAtIndexesToIndex(aBrowser IObject, indexes IObject, destinationIndex uint /* primitive/slice/pointer. */) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("imageBrowser:moveItemsAtIndexes:toIndex:"), aBrowser, indexes, destinationIndex)
	return rv
}


// Signals that a remove operation should be applied to the specified items.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:removeItemsAt:)
func (o_ Object) ImageBrowserRemoveItemsAtIndexes(aBrowser IObject, indexes IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:removeItemsAtIndexes:"), aBrowser, indexes)
}


// Signals that a drag should begin.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:writeItemsAt:to:)
func (o_ Object) ImageBrowserWriteItemsAtIndexesToPasteboard(aBrowser IObject, itemIndexes IObject, pasteboard IObject) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](o_.ID, objc.Sel("imageBrowser:writeItemsAtIndexes:toPasteboard:"), aBrowser, itemIndexes, pasteboard)
	return rv
}


// Performs custom tasks when the selection changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowserSelectionDidChange(_:)
func (o_ Object) ImageBrowserSelectionDidChange(aBrowser IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowserSelectionDidChange:"), aBrowser)
}


// Returns the image to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageRepresentation()
func (o_ Object) ImageRepresentation() objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("imageRepresentation"))
	return rv
}


// Returns the version of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageVersion()
func (o_ Object) ImageVersion() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](o_.ID, objc.Sel("imageVersion"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/index(ofAccessibilityElement:)
func (o_ Object) IndexOfAccessibilityElement(element IObject) int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](o_.ID, objc.Sel("indexOfAccessibilityElement:"), element)
	return rv
}


// Returns the indices of the specified container objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/indicesOfObjects(byEvaluatingObjectSpecifier:)
func (o_ Object) IndicesOfObjectsByEvaluatingObjectSpecifier(specifier IObject) IObject {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("indicesOfObjectsByEvaluatingObjectSpecifier:"), specifier)
	return rv
}


// Handles key down events that do not map to an action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/inputText(_:client:)
func (o_ Object) InputTextClient(string_ IObject, sender IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("inputText:client:"), string_, sender)
	return rv
}


// Receives Unicode, the key code that generated it, and any modifier flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/inputText(_:key:modifiers:client:)
func (o_ Object) InputTextKeyModifiersClient(string_ IObject, keyCode int /* primitive/slice/pointer. */, flags uint /* primitive/slice/pointer. */, sender IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("inputText:key:modifiers:client:"), string_, keyCode, flags, sender)
	return rv
}


// Inserts an object at the specified index in the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/insertValue(_:at:inPropertyWithKey:)
func (o_ Object) InsertValueAtIndexInPropertyWithKey(value IObject, index uint /* primitive/slice/pointer. */, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("insertValue:atIndex:inPropertyWithKey:"), value, index, key)
}


// Inserts an object in the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/insertValue(_:inPropertyWithKey:)
func (o_ Object) InsertValueInPropertyWithKey(value IObject, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("insertValue:inPropertyWithKey:"), value, key)
}


// Executes when a script attempts to invoke a method on an exposed object directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/invokeDefaultMethod(withArguments:)
func (o_ Object) InvokeDefaultMethodWithArguments(arguments IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("invokeDefaultMethodWithArguments:"), arguments)
	return rv
}


// Handles undefined method invocation from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/invokeUndefinedMethod(fromWebScript:withArguments:)
func (o_ Object) InvokeUndefinedMethodFromWebScriptWithArguments(name IObject, arguments IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("invokeUndefinedMethodFromWebScript:withArguments:"), name, arguments)
	return rv
}


// Returns a Boolean value that indicates whether receiver is considered to be “like” a given string when the case of characters in the receiver is ignored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isCaseInsensitiveLike(_:)
func (o_ Object) IsCaseInsensitiveLike(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isCaseInsensitiveLike:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is equal to another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isEqual(to:)
func (o_ Object) IsEqualTo(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isEqualTo:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is greater than another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isGreaterThan(_:)
func (o_ Object) IsGreaterThan(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isGreaterThan:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is greater than or equal to another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isGreaterThanOrEqual(to:)
func (o_ Object) IsGreaterThanOrEqualTo(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isGreaterThanOrEqualTo:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is less than another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isLessThan(_:)
func (o_ Object) IsLessThan(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isLessThan:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is less than or equal to another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isLessThanOrEqual(to:)
func (o_ Object) IsLessThanOrEqualTo(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isLessThanOrEqualTo:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is “like” another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isLike(_:)
func (o_ Object) IsLike(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isLike:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is not equal to another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isNotEqual(to:)
func (o_ Object) IsNotEqualTo(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isNotEqualTo:"), object)
	return rv
}


// Locates and returns the address of the receiver’s implementation of a method so it can be called as a function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/method(for:)
func (o_ Object) MethodForSelector(aSelector objc.SEL) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("methodForSelector:"), aSelector)
	return rv
}


// Creates and returns an instance of a scriptable class, setting its contents and properties, for insertion into the relationship identified by the key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/newScriptingObject(of:forValueForKey:withContentsValue:properties:)
func (o_ Object) NewScriptingObjectOfClassForValueForKeyWithContentsValueProperties(objectClass objc.Class, key IObject, contentsValue IObject, properties IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("newScriptingObjectOfClass:forValueForKey:withContentsValue:properties:"), objectClass, key, contentsValue, properties)
	return rv
}


// Returns the number of groups in an image browser view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/numberOfGroups(inImageBrowser:)
func (o_ Object) NumberOfGroupsInImageBrowser(aBrowser IObject) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](o_.ID, objc.Sel("numberOfGroupsInImageBrowser:"), aBrowser)
	return rv
}


// Returns the number of records managed by the data source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/numberOfItems(inImageBrowser:)
func (o_ Object) NumberOfItemsInImageBrowser(aBrowser IObject) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](o_.ID, objc.Sel("numberOfItemsInImageBrowser:"), aBrowser)
	return rv
}


// Informs the observing object when the value at the specified key path relative to the observed object has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/observeValue(forKeyPath:of:change:context:)
func (o_ Object) ObserveValueForKeyPathOfObjectChangeContext(keyPath IObject, object IObject, change IObject, context unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("observeValueForKeyPath:ofObject:change:context:"), keyPath, object, change, context)
}


// Returns an array describing the options for the specified binding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/optionDescriptionsForBinding(_:)
func (o_ Object) OptionDescriptionsForBinding(binding BindingName /* not a class type */) IObject {
	rv := objc.Send[[]objc.ID](o_.ID, objc.Sel("optionDescriptionsForBinding:"), binding)
	return rv
}


// Invokes a method of the receiver on the specified thread using the default mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:on:with:waitUntilDone:)
func (o_ Object) PerformSelectorOnThreadWithObjectWaitUntilDone(aSelector objc.SEL, thr IObject, arg IObject, wait bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:onThread:withObject:waitUntilDone:"), aSelector, thr, arg, wait)
}


// Invokes a method of the receiver on the specified thread using the specified modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:on:with:waitUntilDone:modes:)
func (o_ Object) PerformSelectorOnThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, thr IObject, arg IObject, wait bool /* primitive/slice/pointer. */, array []string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:onThread:withObject:waitUntilDone:modes:"), aSelector, thr, arg, wait, array)
}


// Invokes a method of the receiver on the current thread using the default mode after a delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:with:afterDelay:)
func (o_ Object) PerformSelectorWithObjectAfterDelay(aSelector objc.SEL, anArgument IObject, delay TimeInterval /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:withObject:afterDelay:"), aSelector, anArgument, delay)
}


// Invokes a method of the receiver on the current thread using the specified modes after a delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:with:afterDelay:inModes:)
func (o_ Object) PerformSelectorWithObjectAfterDelayInModes(aSelector objc.SEL, anArgument IObject, delay TimeInterval /* not a class type */, modes []string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:withObject:afterDelay:inModes:"), aSelector, anArgument, delay, modes)
}


// Sent to the delegate to perform the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performAction(for:identifier:)
func (o_ Object) PerformActionForPersonIdentifier(person IObject, identifier IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performActionForPerson:identifier:"), person, identifier)
}


// Invokes a method of the receiver on a new background thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performSelector(inBackground:with:)
func (o_ Object) PerformSelectorInBackgroundWithObject(aSelector objc.SEL, arg IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelectorInBackground:withObject:"), aSelector, arg)
}


// Invokes a method of the receiver on the main thread using the default mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performSelector(onMainThread:with:waitUntilDone:)
func (o_ Object) PerformSelectorOnMainThreadWithObjectWaitUntilDone(aSelector objc.SEL, arg IObject, wait bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelectorOnMainThread:withObject:waitUntilDone:"), aSelector, arg, wait)
}


// Invokes a method of the receiver on the main thread using the specified modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performSelector(onMainThread:with:waitUntilDone:modes:)
func (o_ Object) PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, arg IObject, wait bool /* primitive/slice/pointer. */, array []string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelectorOnMainThread:withObject:waitUntilDone:modes:"), aSelector, arg, wait, array)
}


// An optional method that an image provider object way implement. With this method, the provider object can use the Metal API to provide pixel data into a MTLTexture when the image object is rendered.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/provideImage(to:commandBuffer:originx:originy:width:height:userInfo:)
func (o_ Object) ProvideImageToMTLTextureCommandBufferOriginxOriginyWidthHeightUserInfo(texture IObject, commandBuffer IObject, originx uintptr /* not a class type */, originy uintptr /* not a class type */, width uintptr /* not a class type */, height uintptr /* not a class type */, info IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("provideImageToMTLTexture:commandBuffer:originx:originy:width:height:userInfo:"), texture, commandBuffer, originx, originy, width, height, info)
}


// Supplies data to a object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/provideImageData(_:bytesPerRow:origin:_:size:_:userInfo:)
func (o_ Object) ProvideImageDataBytesPerRowOriginSizeUserInfo(data unsafe.Pointer, rowbytes uintptr /* not a class type */, originx uintptr /* not a class type */, originy uintptr /* not a class type */, width uintptr /* not a class type */, height uintptr /* not a class type */, info IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("provideImageData:bytesPerRow:origin::size::userInfo:"), data, rowbytes, originx, originy, width, height, info)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/quartzFilterManager(_:didAdd:)
func (o_ Object) QuartzFilterManagerDidAddFilter(sender IObject, filter IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("quartzFilterManager:didAddFilter:"), sender, filter)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/quartzFilterManager(_:didModifyFilter:)
func (o_ Object) QuartzFilterManagerDidModifyFilter(sender IObject, filter IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("quartzFilterManager:didModifyFilter:"), sender, filter)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/quartzFilterManager(_:didRemove:)
func (o_ Object) QuartzFilterManagerDidRemoveFilter(sender IObject, filter IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("quartzFilterManager:didRemoveFilter:"), sender, filter)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/quartzFilterManager(_:didSelect:)
func (o_ Object) QuartzFilterManagerDidSelectFilter(sender IObject, filter IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("quartzFilterManager:didSelectFilter:"), sender, filter)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/readLinkQuality(forDeviceComplete:device:info:error:)
func (o_ Object) ReadLinkQualityForDeviceCompleteDeviceInfoError(controller IObject, device IObject /* already interface */, info unsafe.Pointer, error_ Return /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("readLinkQualityForDeviceComplete:device:info:error:"), controller, device, info, error_)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/readRSSI(forDeviceComplete:device:info:error:)
func (o_ Object) ReadRSSIForDeviceCompleteDeviceInfoError(controller IObject, device IObject /* already interface */, info unsafe.Pointer, error_ Return /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("readRSSIForDeviceComplete:device:info:error:"), controller, device, info, error_)
}


// Stops the observer object from receiving change notifications for the property specified by the key path relative to the object receiving this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/removeObserver(_:forKeyPath:)
func (o_ Object) RemoveObserverForKeyPath(observer IObject, keyPath IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeObserver:forKeyPath:"), observer, keyPath)
}


// Stops the observer object from receiving change notifications for the property specified by the key path relative to the object receiving this message, given the context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/removeObserver(_:forKeyPath:context:)
func (o_ Object) RemoveObserverForKeyPathContext(observer IObject, keyPath IObject, context unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeObserver:forKeyPath:context:"), observer, keyPath, context)
}


// Removes the object at the specified index from the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/removeValue(at:fromPropertyWithKey:)
func (o_ Object) RemoveValueAtIndexFromPropertyWithKey(index uint /* primitive/slice/pointer. */, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeValueAtIndex:fromPropertyWithKey:"), index, key)
}


// Replaces the object at the specified index in the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replaceValue(at:inPropertyWithKey:withValue:)
func (o_ Object) ReplaceValueAtIndexInPropertyWithKeyWithValue(index uint /* primitive/slice/pointer. */, key IObject, value IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("replaceValueAtIndex:inPropertyWithKey:withValue:"), index, key, value)
}


// Overridden by subclasses to substitute another object for itself during encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replacementObject(for:)-2l8ox
func (o_ Object) ReplacementObjectForCoder(coder IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("replacementObjectForCoder:"), coder)
	return rv
}


// Overridden by subclasses to substitute another object for itself during keyed archiving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replacementObject(for:)-60vwc
func (o_ Object) ReplacementObjectForKeyedArchiver(archiver IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("replacementObjectForKeyedArchiver:"), archiver)
	return rv
}


// Called to determine if the specified uniform type identifier should be shown in the save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/saveOptions(_:shouldShowUTType:)
func (o_ Object) SaveOptionsShouldShowUTType(saveOptions IObject, utType IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("saveOptions:shouldShowUTType:"), saveOptions, utType)
	return rv
}


// Returns if, in a scripting comparison, the compared object matches the beginning of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingBegins(with:)
func (o_ Object) ScriptingBeginsWith(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingBeginsWith:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object contains .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingContains(_:)
func (o_ Object) ScriptingContains(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingContains:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object matches the end of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingEnds(with:)
func (o_ Object) ScriptingEndsWith(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingEndsWith:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is equal to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsEqual(to:)
func (o_ Object) ScriptingIsEqualTo(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsEqualTo:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is greater than .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsGreaterThan(_:)
func (o_ Object) ScriptingIsGreaterThan(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsGreaterThan:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is greater than or equal to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsGreaterThanOrEqual(to:)
func (o_ Object) ScriptingIsGreaterThanOrEqualTo(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsGreaterThanOrEqualTo:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is less than .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsLessThan(_:)
func (o_ Object) ScriptingIsLessThan(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsLessThan:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is less than or equal to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsLessThanOrEqual(to:)
func (o_ Object) ScriptingIsLessThanOrEqualTo(object IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsLessThanOrEqualTo:"), object)
	return rv
}


// Given an object specifier, returns the specified object or objects in the receiving container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingValue(for:)
func (o_ Object) ScriptingValueForSpecifier(objectSpecifier IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("scriptingValueForSpecifier:"), objectSpecifier)
	return rv
}


// Invoked by when it’s given a value for a scalar value (such as an or ).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setNilValueForKey(_:)
func (o_ Object) SetNilValueForKey(key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setNilValueForKey:"), key)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setSharedObservers(_:)
func (o_ Object) SetSharedObservers(sharedObservers IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setSharedObservers:"), sharedObservers)
}


// Sets the property of the receiver specified by a given key to a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forKey:)
func (o_ Object) SetValueForKey(value IObject, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValue:forKey:"), value, key)
}


// Sets the value for the property identified by a given key path to a given value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forKeyPath:)
func (o_ Object) SetValueForKeyPath(value IObject, keyPath IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValue:forKeyPath:"), value, keyPath)
}


// Invoked by when it finds no property for a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValue(_:forUndefinedKey:)
func (o_ Object) SetValueForUndefinedKey(value IObject, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValue:forUndefinedKey:"), value, key)
}


// Sets properties of the receiver with values from a given dictionary, using its keys to identify the properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setValuesForKeys(_:)
func (o_ Object) SetValuesForKeysWithDictionary(keyedValues IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setValuesForKeysWithDictionary:"), keyedValues)
}


// Allows the delegate to specify which device is its preferred.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanel(_:determineBestDeviceOfA:orB:)
func (o_ Object) SetupPanelDetermineBestDeviceOfAOrB(aPanel unsafe.Pointer, deviceA unsafe.Pointer, device unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("setupPanel:determineBestDeviceOfA:orB:"), aPanel, deviceA, device)
	return rv
}


// This delegate method allows the delegate to determine if the media inserted in the device is suitable for whatever operation is to be performed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanel(_:deviceContainsSuitableMedia:promptString:)
func (o_ Object) SetupPanelDeviceContainsSuitableMediaPromptString(aPanel unsafe.Pointer, device unsafe.Pointer, prompt IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("setupPanel:deviceContainsSuitableMedia:promptString:"), aPanel, device, prompt)
	return rv
}


// Allows the delegate to determine if device can be used as a target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanel(_:deviceCouldBeTarget:)
func (o_ Object) SetupPanelDeviceCouldBeTarget(aPanel unsafe.Pointer, device unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("setupPanel:deviceCouldBeTarget:"), aPanel, device)
	return rv
}


// Sent by the default notification center when the device selection in the panel has changed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanelDeviceSelectionChanged(_:)
func (o_ Object) SetupPanelDeviceSelectionChanged(aNotification IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setupPanelDeviceSelectionChanged:"), aNotification)
}


// This delegate method allows the delegate to control how media reservations are handled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanelShouldHandleMediaReservations(_:)
func (o_ Object) SetupPanelShouldHandleMediaReservations(aPanel unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("setupPanelShouldHandleMediaReservations:"), aPanel)
	return rv
}


// Sent to the delegate to determine whether the action should be enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/shouldEnableAction(for:identifier:)
func (o_ Object) ShouldEnableActionForPersonIdentifier(person IObject, identifier IObject) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("shouldEnableActionForPerson:identifier:"), person, identifier)
	return rv
}


// Removes a given binding between the receiver and a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/unbind(_:)
func (o_ Object) Unbind(binding BindingName /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("unbind:"), binding)
}


// Indicates whether the value specified by a given pointer is valid, or can be made valid, for the property identified by a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/validateValue(_:forKey:)
func (o_ Object) ValidateValueForKeyError(ioValue unsafe.Pointer, inKey IObject, outError unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("validateValue:forKey:error:"), ioValue, inKey, outError)
	return rv
}


// Indicates whether the value specified by a given pointer is not valid for a given key path relative to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/validateValue(_:forKeyPath:)
func (o_ Object) ValidateValueForKeyPathError(ioValue unsafe.Pointer, inKeyPath IObject, outError unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("validateValue:forKeyPath:error:"), ioValue, inKeyPath, outError)
	return rv
}


// Retrieves an indexed object from the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(at:inPropertyWithKey:)
func (o_ Object) ValueAtIndexInPropertyWithKey(index uint /* primitive/slice/pointer. */, key IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueAtIndex:inPropertyWithKey:"), index, key)
	return rv
}


// Returns the value for the property identified by a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forKey:)
func (o_ Object) ValueForKey(key IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueForKey:"), key)
	return rv
}


// Returns the value for the derived property identified by a given key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forKeyPath:)
func (o_ Object) ValueForKeyPath(keyPath IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueForKeyPath:"), keyPath)
	return rv
}


// Invoked by when it finds no property corresponding to a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forUndefinedKey:)
func (o_ Object) ValueForUndefinedKey(key IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueForUndefinedKey:"), key)
	return rv
}


// Retrieves a named object from the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(withName:inPropertyWithKey:)
func (o_ Object) ValueWithNameInPropertyWithKey(name IObject, key IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueWithName:inPropertyWithKey:"), name, key)
	return rv
}


// Retrieves an object by ID from the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(withUniqueID:inPropertyWithKey:)
func (o_ Object) ValueWithUniqueIDInPropertyWithKey(uniqueID IObject, key IObject) objc.ID {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueWithUniqueID:inPropertyWithKey:"), uniqueID, key)
	return rv
}


// Returns the class of the value that will be returned for the specified binding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/valueClassForBinding(_:)
func (o_ Object) ValueClassForBinding(binding BindingName /* not a class type */) objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("valueClassForBinding:"), binding)
	return rv
}


// Loads a URL into a web frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInContainerLoad(_:inFrame:)
func (o_ Object) WebPlugInContainerLoadRequestInFrame(request IObject, target IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInContainerLoadRequest:inFrame:"), request, target)
}


// Tells the container to show a status message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInContainerShowStatus(_:)
func (o_ Object) WebPlugInContainerShowStatus(message IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInContainerShowStatus:"), message)
}


// Prepares the plug-in for deallocation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInDestroy()
func (o_ Object) WebPlugInDestroy() {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInDestroy"))
}


// Initializes the plug-in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInInitialize()
func (o_ Object) WebPlugInInitialize() {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInInitialize"))
}


// Invoked when an error occurs loading the main resource.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInMainResourceDidFailWithError(_:)
func (o_ Object) WebPlugInMainResourceDidFailWithError(error_ IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInMainResourceDidFailWithError:"), error_)
}


// Invoked when the connection successfully finishes loading data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInMainResourceDidFinishLoading()
func (o_ Object) WebPlugInMainResourceDidFinishLoading() {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInMainResourceDidFinishLoading"))
}


// Invoked when the connection loads data incrementally.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInMainResourceDidReceive(_:)-5b6f6
func (o_ Object) WebPlugInMainResourceDidReceiveData(data IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInMainResourceDidReceiveData:"), data)
}


// Invoked when the connection receives sufficient data to construct the URL response for its request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInMainResourceDidReceive(_:)-6x7b9
func (o_ Object) WebPlugInMainResourceDidReceiveResponse(response IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInMainResourceDidReceiveResponse:"), response)
}


// Controls plug-in behavior based on its selection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInSetIsSelected(_:)
func (o_ Object) WebPlugInSetIsSelected(isSelected bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInSetIsSelected:"), isSelected)
}


// Tells the plug-in to start normal operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInStart()
func (o_ Object) WebPlugInStart() {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInStart"))
}


// Tells the plug-in to stop normal operation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInStop()
func (o_ Object) WebPlugInStop() {
	objc.Send[objc.ID](o_.ID, objc.Sel("webPlugInStop"))
}


// Informs the observed object that the specified change is about to be executed at given indexes for a specified ordered to-many relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/willChange(_:valuesAt:forKey:)
func (o_ Object) WillChangeValuesAtIndexesForKey(changeKind KeyValueChange /* not a class type */, indexes IObject, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("willChange:valuesAtIndexes:forKey:"), changeKind, indexes, key)
}


// Informs the observed object that the value of a given property is about to change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/willChangeValue(forKey:)
func (o_ Object) WillChangeValueForKey(key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("willChangeValueForKey:"), key)
}


// Informs the observed object that the specified change is about to be made to a specified unordered to-many relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/willChangeValue(forKey:withSetMutation:using:)
func (o_ Object) WillChangeValueForKeyWithSetMutationUsingObjects(key IObject, mutationKind KeyValueSetMutationKind /* not a class type */, objects IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("willChangeValueForKey:withSetMutation:usingObjects:"), key, mutationKind, objects)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityactivateblock
func (o_ Object) AccessibilityActivateBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityActivateBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityactivateblock
func (o_ Object) SetAccessibilityActivateBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityActivateBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityactivationpointblock
func (o_ Object) AccessibilityActivationPointBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityActivationPointBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityactivationpointblock
func (o_ Object) SetAccessibilityActivationPointBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityActivationPointBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityattributedhintblock
func (o_ Object) AccessibilityAttributedHintBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAttributedHintBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityattributedhintblock
func (o_ Object) SetAccessibilityAttributedHintBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityAttributedHintBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityattributedlabelblock
func (o_ Object) AccessibilityAttributedLabelBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAttributedLabelBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityattributedlabelblock
func (o_ Object) SetAccessibilityAttributedLabelBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityAttributedLabelBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityattributeduserinputlabelsblock
func (o_ Object) AccessibilityAttributedUserInputLabelsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAttributedUserInputLabelsBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityattributeduserinputlabelsblock
func (o_ Object) SetAccessibilityAttributedUserInputLabelsBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityAttributedUserInputLabelsBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityattributedvalueblock
func (o_ Object) AccessibilityAttributedValueBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityAttributedValueBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityattributedvalueblock
func (o_ Object) SetAccessibilityAttributedValueBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityAttributedValueBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycontainertype
func (o_ Object) AccessibilityContainerType() IObject /* already interface */ {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityContainerType"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycontainertype
func (o_ Object) SetAccessibilityContainerType(value IObject /* already interface */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityContainerType:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycontainertypeblock
func (o_ Object) AccessibilityContainerTypeBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityContainerTypeBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycontainertypeblock
func (o_ Object) SetAccessibilityContainerTypeBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityContainerTypeBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycustomactionsblock
func (o_ Object) AccessibilityCustomActionsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityCustomActionsBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycustomactionsblock
func (o_ Object) SetAccessibilityCustomActionsBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityCustomActionsBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycustomrotors
func (o_ Object) AccessibilityCustomRotors() IObject /* already interface */ {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityCustomRotors"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycustomrotors
func (o_ Object) SetAccessibilityCustomRotors(value IObject /* already interface */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityCustomRotors:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycustomrotorsblock
func (o_ Object) AccessibilityCustomRotorsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityCustomRotorsBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitycustomrotorsblock
func (o_ Object) SetAccessibilityCustomRotorsBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityCustomRotorsBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitydecrementblock
func (o_ Object) AccessibilityDecrementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityDecrementBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitydecrementblock
func (o_ Object) SetAccessibilityDecrementBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityDecrementBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitydirecttouchoptions
func (o_ Object) AccessibilityDirectTouchOptions() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityDirectTouchOptions"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitydirecttouchoptions
func (o_ Object) SetAccessibilityDirectTouchOptions(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityDirectTouchOptions:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityelements
func (o_ Object) AccessibilityElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityElements"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityelements
func (o_ Object) SetAccessibilityElements(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityElements:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityelementsblock
func (o_ Object) AccessibilityElementsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityElementsBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityelementsblock
func (o_ Object) SetAccessibilityElementsBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityElementsBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityelementshidden
func (o_ Object) AccessibilityElementsHidden() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityElementsHidden"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityelementshidden
func (o_ Object) SetAccessibilityElementsHidden(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityElementsHidden:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityelementshiddenblock
func (o_ Object) AccessibilityElementsHiddenBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityElementsHiddenBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityelementshiddenblock
func (o_ Object) SetAccessibilityElementsHiddenBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityElementsHiddenBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityexpandedstatus
func (o_ Object) AccessibilityExpandedStatus() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityExpandedStatus"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityexpandedstatus
func (o_ Object) SetAccessibilityExpandedStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityExpandedStatus:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityexpandedstatusblock
func (o_ Object) AccessibilityExpandedStatusBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityExpandedStatusBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityexpandedstatusblock
func (o_ Object) SetAccessibilityExpandedStatusBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityExpandedStatusBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityfocuseduielement
func (o_ Object) AccessibilityFocusedUIElement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityFocusedUIElement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityfocuseduielement
func (o_ Object) SetAccessibilityFocusedUIElement(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityFocusedUIElement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityframeblock
func (o_ Object) AccessibilityFrameBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityFrameBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityframeblock
func (o_ Object) SetAccessibilityFrameBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityFrameBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityheaderelements
func (o_ Object) AccessibilityHeaderElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityHeaderElements"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityheaderelements
func (o_ Object) SetAccessibilityHeaderElements(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityHeaderElements:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityheaderelementsblock
func (o_ Object) AccessibilityHeaderElementsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityHeaderElementsBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityheaderelementsblock
func (o_ Object) SetAccessibilityHeaderElementsBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityHeaderElementsBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityhintblock
func (o_ Object) AccessibilityHintBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityHintBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityhintblock
func (o_ Object) SetAccessibilityHintBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityHintBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityidentifierblock
func (o_ Object) AccessibilityIdentifierBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityIdentifierBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityidentifierblock
func (o_ Object) SetAccessibilityIdentifierBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityIdentifierBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityincrementblock
func (o_ Object) AccessibilityIncrementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityIncrementBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityincrementblock
func (o_ Object) SetAccessibilityIncrementBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityIncrementBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitylabelblock
func (o_ Object) AccessibilityLabelBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityLabelBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitylabelblock
func (o_ Object) SetAccessibilityLabelBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityLabelBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitylanguageblock
func (o_ Object) AccessibilityLanguageBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityLanguageBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitylanguageblock
func (o_ Object) SetAccessibilityLanguageBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityLanguageBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitymagictapblock
func (o_ Object) AccessibilityMagicTapBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityMagicTapBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitymagictapblock
func (o_ Object) SetAccessibilityMagicTapBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityMagicTapBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynavigationstyle
func (o_ Object) AccessibilityNavigationStyle() IObject /* already interface */ {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityNavigationStyle"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynavigationstyle
func (o_ Object) SetAccessibilityNavigationStyle(value IObject /* already interface */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityNavigationStyle:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynavigationstyleblock
func (o_ Object) AccessibilityNavigationStyleBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityNavigationStyleBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynavigationstyleblock
func (o_ Object) SetAccessibilityNavigationStyleBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityNavigationStyleBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynexttextnavigationelement
func (o_ Object) AccessibilityNextTextNavigationElement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityNextTextNavigationElement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynexttextnavigationelement
func (o_ Object) SetAccessibilityNextTextNavigationElement(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityNextTextNavigationElement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynexttextnavigationelementblock
func (o_ Object) AccessibilityNextTextNavigationElementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityNextTextNavigationElementBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynexttextnavigationelementblock
func (o_ Object) SetAccessibilityNextTextNavigationElementBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityNextTextNavigationElementBlock:"), value)
}


// A Boolean value that indicates whether a custom accessibility object sends a notification when its corresponding UI element is destroyed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynotifieswhendestroyed
func (o_ Object) AccessibilityNotifiesWhenDestroyed() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityNotifiesWhenDestroyed"))
	return rv
}


// A Boolean value that indicates whether a custom accessibility object sends a notification when its corresponding UI element is destroyed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitynotifieswhendestroyed
func (o_ Object) SetAccessibilityNotifiesWhenDestroyed(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityNotifiesWhenDestroyed:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitypathblock
func (o_ Object) AccessibilityPathBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityPathBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitypathblock
func (o_ Object) SetAccessibilityPathBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityPathBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityperformescapeblock
func (o_ Object) AccessibilityPerformEscapeBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityPerformEscapeBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityperformescapeblock
func (o_ Object) SetAccessibilityPerformEscapeBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityPerformEscapeBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityprevioustextnavigationelement
func (o_ Object) AccessibilityPreviousTextNavigationElement() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityPreviousTextNavigationElement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityprevioustextnavigationelement
func (o_ Object) SetAccessibilityPreviousTextNavigationElement(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityPreviousTextNavigationElement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityprevioustextnavigationelementblock
func (o_ Object) AccessibilityPreviousTextNavigationElementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityPreviousTextNavigationElementBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityprevioustextnavigationelementblock
func (o_ Object) SetAccessibilityPreviousTextNavigationElementBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityPreviousTextNavigationElementBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityrespondstouserinteraction
func (o_ Object) AccessibilityRespondsToUserInteraction() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityRespondsToUserInteraction"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityrespondstouserinteraction
func (o_ Object) SetAccessibilityRespondsToUserInteraction(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityRespondsToUserInteraction:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityrespondstouserinteractionblock
func (o_ Object) AccessibilityRespondsToUserInteractionBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityRespondsToUserInteractionBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityrespondstouserinteractionblock
func (o_ Object) SetAccessibilityRespondsToUserInteractionBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityRespondsToUserInteractionBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityshouldgroupaccessibilitychildrenblock
func (o_ Object) AccessibilityShouldGroupAccessibilityChildrenBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityShouldGroupAccessibilityChildrenBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityshouldgroupaccessibilitychildrenblock
func (o_ Object) SetAccessibilityShouldGroupAccessibilityChildrenBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityShouldGroupAccessibilityChildrenBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytextinputresponder
func (o_ Object) AccessibilityTextInputResponder() IObject /* already interface */ {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityTextInputResponder"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytextinputresponder
func (o_ Object) SetAccessibilityTextInputResponder(value IObject /* already interface */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityTextInputResponder:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytextinputresponderblock
func (o_ Object) AccessibilityTextInputResponderBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityTextInputResponderBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytextinputresponderblock
func (o_ Object) SetAccessibilityTextInputResponderBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityTextInputResponderBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytextualcontext
func (o_ Object) AccessibilityTextualContext() IObject /* already interface */ {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityTextualContext"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytextualcontext
func (o_ Object) SetAccessibilityTextualContext(value IObject /* already interface */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityTextualContext:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytextualcontextblock
func (o_ Object) AccessibilityTextualContextBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityTextualContextBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytextualcontextblock
func (o_ Object) SetAccessibilityTextualContextBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityTextualContextBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytraits
func (o_ Object) AccessibilityTraits() IObject /* already interface */ {
	rv := objc.Send[IObject](o_.ID, objc.Sel("accessibilityTraits"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytraits
func (o_ Object) SetAccessibilityTraits(value IObject /* already interface */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityTraits:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytraitsblock
func (o_ Object) AccessibilityTraitsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityTraitsBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilitytraitsblock
func (o_ Object) SetAccessibilityTraitsBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityTraitsBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityuserinputlabelsblock
func (o_ Object) AccessibilityUserInputLabelsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityUserInputLabelsBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityuserinputlabelsblock
func (o_ Object) SetAccessibilityUserInputLabelsBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityUserInputLabelsBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityvalueblock
func (o_ Object) AccessibilityValueBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityValueBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityvalueblock
func (o_ Object) SetAccessibilityValueBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityValueBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityviewismodal
func (o_ Object) AccessibilityViewIsModal() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityViewIsModal"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityviewismodal
func (o_ Object) SetAccessibilityViewIsModal(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityViewIsModal:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityviewismodalblock
func (o_ Object) AccessibilityViewIsModalBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("accessibilityViewIsModalBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/accessibilityviewismodalblock
func (o_ Object) SetAccessibilityViewIsModalBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAccessibilityViewIsModalBlock:"), value)
}


// A proxy for the receiving object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/autocontentaccessingproxy
func (o_ Object) AutoContentAccessingProxy() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("autoContentAccessingProxy"))
	return rv
}


// A proxy for the receiving object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/autocontentaccessingproxy
func (o_ Object) SetAutoContentAccessingProxy(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutoContentAccessingProxy:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/automationelements
func (o_ Object) AutomationElements() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("automationElements"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/automationelements
func (o_ Object) SetAutomationElements(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutomationElements:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/automationelementsblock
func (o_ Object) AutomationElementsBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("automationElementsBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/automationelementsblock
func (o_ Object) SetAutomationElementsBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setAutomationElementsBlock:"), value)
}


// The kind of container that contains this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/browseraccessibilitycontainertype
func (o_ Object) BrowserAccessibilityContainerType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("browserAccessibilityContainerType"))
	return rv
}


// The kind of container that contains this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/browseraccessibilitycontainertype
func (o_ Object) SetBrowserAccessibilityContainerType(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBrowserAccessibilityContainerType:"), value)
}


// A Boolean value that indicates whether the element has native focus in the browser Document Object Model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/browseraccessibilityhasdomfocus
func (o_ Object) BrowserAccessibilityHasDOMFocus() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("browserAccessibilityHasDOMFocus"))
	return rv
}


// A Boolean value that indicates whether the element has native focus in the browser Document Object Model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/browseraccessibilityhasdomfocus
func (o_ Object) SetBrowserAccessibilityHasDOMFocus(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBrowserAccessibilityHasDOMFocus:"), value)
}


// A Boolean value that’s the element’s value for aria-required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/browseraccessibilityisrequired
func (o_ Object) BrowserAccessibilityIsRequired() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("browserAccessibilityIsRequired"))
	return rv
}


// A Boolean value that’s the element’s value for aria-required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/browseraccessibilityisrequired
func (o_ Object) SetBrowserAccessibilityIsRequired(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBrowserAccessibilityIsRequired:"), value)
}


// The element’s value for aria-pressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/browseraccessibilitypressedstate
func (o_ Object) BrowserAccessibilityPressedState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("browserAccessibilityPressedState"))
	return rv
}


// The element’s value for aria-pressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/browseraccessibilitypressedstate
func (o_ Object) SetBrowserAccessibilityPressedState(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBrowserAccessibilityPressedState:"), value)
}


// The receiver’s Apple event type code, as stored in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/classcode
func (o_ Object) ClassCode() uint32 /* not a class type */ {
	rv := objc.Send[uint32](o_.ID, objc.Sel("classCode"))
	return rv
}


// The receiver’s Apple event type code, as stored in the
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/classcode
func (o_ Object) SetClassCode(value uint32 /* not a class type */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setClassCode:"), value)
}


// The class to substitute for the receiver’s own class during archiving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/classforarchiver
func (o_ Object) ClassForArchiver() objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("classForArchiver"))
	return rv
}


// The class to substitute for the receiver’s own class during archiving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/classforarchiver
func (o_ Object) SetClassForArchiver(value objc.Class) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setClassForArchiver:"), value)
}


// Overridden by subclasses to substitute a class other than its own during coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/classforcoder
func (o_ Object) ClassForCoder() objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("classForCoder"))
	return rv
}


// Overridden by subclasses to substitute a class other than its own during coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/classforcoder
func (o_ Object) SetClassForCoder(value objc.Class) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setClassForCoder:"), value)
}


// Subclasses to substitute a new class for instances during keyed archiving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/classforkeyedarchiver
func (o_ Object) ClassForKeyedArchiver() objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("classForKeyedArchiver"))
	return rv
}


// Subclasses to substitute a new class for instances during keyed archiving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/classforkeyedarchiver
func (o_ Object) SetClassForKeyedArchiver(value objc.Class) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setClassForKeyedArchiver:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/isaccessibilityelement
func (o_ Object) IsAccessibilityElement() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isAccessibilityElement"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/isaccessibilityelement
func (o_ Object) SetIsAccessibilityElement(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsAccessibilityElement:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/isaccessibilityelementblock
func (o_ Object) IsAccessibilityElementBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("isAccessibilityElementBlock"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/isaccessibilityelementblock
func (o_ Object) SetIsAccessibilityElementBlock(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsAccessibilityElementBlock:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/isselectable
func (o_ Object) IsSelectable() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("isSelectable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/isselectable
func (o_ Object) SetIsSelectable(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsSelectable:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/shouldgroupaccessibilitychildren
func (o_ Object) ShouldGroupAccessibilityChildren() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](o_.ID, objc.Sel("shouldGroupAccessibilityChildren"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/shouldgroupaccessibilitychildren
func (o_ Object) SetShouldGroupAccessibilityChildren(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setShouldGroupAccessibilityChildren:"), value)
}


