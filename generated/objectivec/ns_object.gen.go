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
	AccessibilityFocusedUIElement() IObject
	AccessibilityNotifiesWhenDestroyed() bool
	AttributeKeys() []string
	AutoContentAccessingProxy() IObject
	BrowserAccessibilityContainerType() unsafe.Pointer
	SetBrowserAccessibilityContainerType(value unsafe.Pointer)
	BrowserAccessibilityHasDOMFocus() bool
	SetBrowserAccessibilityHasDOMFocus(value bool)
	BrowserAccessibilityIsRequired() bool
	SetBrowserAccessibilityIsRequired(value bool)
	BrowserAccessibilityPressedState() unsafe.Pointer
	SetBrowserAccessibilityPressedState(value unsafe.Pointer)
	ClassCode() uint32 /* not a class type */
	ClassDescription() IObject
	ClassForArchiver() objc.Class
	ClassForCoder() objc.Class
	ClassForKeyedArchiver() objc.Class
	ClassForPortCoder() objc.Class
	ExposedBindings() []string
	Selectable() bool
	ObjectForWebScript() IObject
	ObjectSpecifier() IObject
	ObservationInfo() unsafe.Pointer
	SetObservationInfo(value unsafe.Pointer)
	ToManyRelationshipKeys() []string
	ToOneRelationshipKeys() []string
	WebPlugInContainerSelectionColor() IObject
	IsSelectable() bool
	SetIsSelectable(value bool)
	// methods:
	AcceptsPreviewPanelControl(panel IObject) bool
	AccessibilityArrayAttributeCount(attribute string) uint
	AccessibilityArrayAttributeValuesIndexMaxCount(attribute string, index uint, maxCount uint) IObject
	AccessibilityHitTest(point IObject) IObject
	AccessibilityIndexOfChild(child IObject) uint
	AccessibilityLineEndPositionFromCurrentSelection() objc.IObject /* cross-framework: Integer */
	AccessibilityLineRangeForPosition(position objc.IObject /* cross-framework: Integer */) IObject
	AccessibilityLineStartPositionFromCurrentSelection() objc.IObject /* cross-framework: Integer */
	ActionProperty() IObject
	AttemptRecoveryFromErrorOptionIndex(error_ IObject, recoveryOptionIndex uint) bool
	AttemptRecoveryFromErrorOptionIndexDelegateDidRecoverSelectorContextInfo(error_ IObject, recoveryOptionIndex uint, delegate IObject, didRecoverSelector objc.SEL, contextInfo unsafe.Pointer)
	AuthorizationViewCreatedAuthorization(view IObject)
	AuthorizationViewDidAuthorize(view IObject)
	AuthorizationViewDidDeauthorize(view IObject)
	AuthorizationViewDidHide(view IObject)
	AuthorizationViewReleasedAuthorization(view IObject)
	AuthorizationViewShouldDeauthorize(view IObject) bool
	BeginPreviewPanelControl(panel IObject)
	BindToObjectWithKeyPathOptions(binding string, observable IObject, keyPath IObject, options IObject)
	BrowserAccessibilityAttributedValueInRange(range_ IObject) IObject
	BrowserAccessibilityDeleteTextAtCursor(numberOfCharacters objc.IObject /* cross-framework: Integer */)
	BrowserAccessibilityInsertTextAtCursor(text IObject)
	BrowserAccessibilitySelectedTextRange() IObject
	BrowserAccessibilitySetSelectedTextRange(range_ IObject)
	BrowserAccessibilityValueInRange(range_ IObject) IObject
	BurnProgressPanelBurnDidFinish(theBurnPanel unsafe.Pointer, burn unsafe.Pointer) bool
	BurnProgressPanelDidFinish(aNotification IObject)
	BurnProgressPanelWillBegin(aNotification IObject)
	Candidates(sender IObject) IObject
	CertificatePanelShowHelp(sender IObject) bool
	ChooseIdentityPanelShowHelp(sender IObject) bool
	CoerceValueForKey(value IObject, key IObject) IObject
	CommitComposition(sender IObject)
	ComposedString(sender IObject) IObject
	CopyScriptingValueForKeyWithProperties(value IObject, key IObject, properties IObject) IObject
	Dealloc()
	DidChangeValueForKeyWithSetMutationUsingObjects(key IObject, mutationKind uint, objects IObject)
	DidCommandBySelectorClient(aSelector objc.SEL, sender IObject) bool
	DoesContain(object IObject) bool
	DoesNotRecognizeSelector(aSelector objc.SEL)
	EndPreviewPanelControl(panel IObject)
	EraseProgressPanelEraseDidFinish(theErasePanel unsafe.Pointer, erase unsafe.Pointer) bool
	EraseProgressPanelDidFinish(aNotification IObject)
	EraseProgressPanelWillBegin(aNotification IObject)
	ExceptionHandlerShouldHandleExceptionMask(sender IObject, exception IObject, aMask uint) bool
	ExceptionHandlerShouldLogExceptionMask(sender IObject, exception IObject, aMask uint) bool
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
	ForwardInvocation(anInvocation IObject)
	ForwardingTargetForSelector(aSelector objc.SEL) IObject
	HandleEventClient(event IObject, sender IObject) bool
	ImageBrowserBackgroundWasRightClickedWithEvent(aBrowser IObject, event IObject)
	ImageBrowserCellWasDoubleClickedAtIndex(aBrowser IObject, index uint)
	ImageBrowserCellWasRightClickedAtIndexWithEvent(aBrowser IObject, index uint, event IObject)
	ImageBrowserGroupAtIndex(aBrowser IObject, index uint) IObject
	ImageBrowserItemAtIndex(aBrowser IObject, index uint) IObject
	ImageBrowserMoveItemsAtIndexesToIndex(aBrowser IObject, indexes IObject, destinationIndex uint) bool
	ImageBrowserRemoveItemsAtIndexes(aBrowser IObject, indexes IObject)
	ImageBrowserWriteItemsAtIndexesToPasteboard(aBrowser IObject, itemIndexes IObject, pasteboard IObject) uint
	ImageBrowserSelectionDidChange(aBrowser IObject)
	ImageRepresentation() IObject
	ImageRepresentationType() IObject
	ImageSubtitle() IObject
	ImageTitle() IObject
	ImageUID() IObject
	ImageVersion() uint
	IndicesOfObjectsByEvaluatingObjectSpecifier(specifier IObject) []objc.ID
	InputTextClient(string_ IObject, sender IObject) bool
	InputTextKeyModifiersClient(string_ IObject, keyCode objc.IObject /* cross-framework: Integer */, flags uint, sender IObject) bool
	InsertValueAtIndexInPropertyWithKey(value IObject, index uint, key IObject)
	InsertValueInPropertyWithKey(value IObject, key IObject)
	InverseForRelationshipKey(relationshipKey IObject) IObject
	InvokeDefaultMethodWithArguments(arguments IObject) IObject
	InvokeUndefinedMethodFromWebScriptWithArguments(name IObject, arguments IObject) IObject
	IsCaseInsensitiveLike(object IObject) bool
	IsGreaterThan(object IObject) bool
	IsGreaterThanOrEqualTo(object IObject) bool
	IsLessThan(object IObject) bool
	IsLessThanOrEqualTo(object IObject) bool
	IsLike(object IObject) bool
	IsNotEqualTo(object IObject) bool
	MethodForSelector(aSelector objc.SEL) IMP
	MethodSignatureForSelector(aSelector objc.SEL) MethodSignature /* not a class type */
	MutableArrayValueForKey(key IObject) IObject
	MutableArrayValueForKeyPath(keyPath IObject) IObject
	MutableOrderedSetValueForKey(key IObject) IObject
	MutableOrderedSetValueForKeyPath(keyPath IObject) IObject
	MutableSetValueForKey(key IObject) IObject
	MutableSetValueForKeyPath(keyPath IObject) IObject
	NewScriptingObjectOfClassForValueForKeyWithContentsValueProperties(objectClass objc.Class, key IObject, contentsValue IObject, properties IObject) IObject
	NumberOfGroupsInImageBrowser(aBrowser IObject) uint
	NumberOfItemsInImageBrowser(aBrowser IObject) uint
	OptionDescriptionsForBinding(binding string) []objc.ID
	OriginalString(sender IObject) IObject
	PerformSelectorOnThreadWithObjectWaitUntilDone(aSelector objc.SEL, thr IObject, arg IObject, wait bool)
	PerformSelectorOnThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, thr IObject, arg IObject, wait bool, array []string)
	PerformSelectorWithObjectAfterDelay(aSelector objc.SEL, anArgument IObject, delay float64)
	PerformSelectorWithObjectAfterDelayInModes(aSelector objc.SEL, anArgument IObject, delay float64, modes []string)
	PerformActionForPersonIdentifier(person IObject, identifier IObject)
	PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, arg IObject, wait bool, array []string)
	ProvideImageToMTLTextureCommandBufferOriginxOriginyWidthHeightUserInfo(texture IObject, commandBuffer IObject, originx uintptr /* not a class type */, originy uintptr /* not a class type */, width uintptr /* not a class type */, height uintptr /* not a class type */, info IObject)
	ProvideImageDataBytesPerRowOriginSizeUserInfo(data unsafe.Pointer, rowbytes uintptr /* not a class type */, originx uintptr /* not a class type */, originy uintptr /* not a class type */, width uintptr /* not a class type */, height uintptr /* not a class type */, info IObject)
	QuartzFilterManagerDidAddFilter(sender IObject, filter IObject)
	QuartzFilterManagerDidModifyFilter(sender IObject, filter IObject)
	QuartzFilterManagerDidRemoveFilter(sender IObject, filter IObject)
	QuartzFilterManagerDidSelectFilter(sender IObject, filter IObject)
	ReadLinkQualityForDeviceCompleteDeviceInfoError(controller IObject, device IObject, info IObject, error_ int)
	ReadRSSIForDeviceCompleteDeviceInfoError(controller IObject, device IObject, info IObject, error_ int)
	RemoveValueAtIndexFromPropertyWithKey(index uint, key IObject)
	ReplaceValueAtIndexInPropertyWithKeyWithValue(index uint, key IObject, value IObject)
	ReplacementObjectForCoder(coder IObject) IObject
	ReplacementObjectForKeyedArchiver(archiver IObject) IObject
	SaveOptionsShouldShowUTType(saveOptions IObject, utType IObject) bool
	ScriptingBeginsWith(object IObject) bool
	ScriptingContains(object IObject) bool
	ScriptingEndsWith(object IObject) bool
	ScriptingIsEqualTo(object IObject) bool
	ScriptingIsGreaterThan(object IObject) bool
	ScriptingIsGreaterThanOrEqualTo(object IObject) bool
	ScriptingIsLessThan(object IObject) bool
	ScriptingIsLessThanOrEqualTo(object IObject) bool
	ScriptingValueForSpecifier(objectSpecifier IObject) IObject
	SetSharedObservers(sharedObservers IObject)
	SetupPanelDetermineBestDeviceOfAOrB(aPanel unsafe.Pointer, deviceA unsafe.Pointer, device unsafe.Pointer) unsafe.Pointer
	SetupPanelDeviceContainsSuitableMediaPromptString(aPanel unsafe.Pointer, device unsafe.Pointer, prompt IObject) bool
	SetupPanelDeviceCouldBeTarget(aPanel unsafe.Pointer, device unsafe.Pointer) bool
	SetupPanelDeviceSelectionChanged(aNotification IObject)
	SetupPanelShouldHandleMediaReservations(aPanel unsafe.Pointer) bool
	ShouldEnableActionForPersonIdentifier(person IObject, identifier IObject) bool
	TitleForPersonIdentifier(person IObject, identifier IObject) IObject
	Unbind(binding string)
	ValidateValueForKeyError(ioValue unsafe.Pointer, inKey IObject, outError unsafe.Pointer) bool
	ValidateValueForKeyPathError(ioValue unsafe.Pointer, inKeyPath IObject, outError unsafe.Pointer) bool
	ValueAtIndexInPropertyWithKey(index uint, key IObject) IObject
	ValueWithNameInPropertyWithKey(name IObject, key IObject) IObject
	ValueWithUniqueIDInPropertyWithKey(uniqueID IObject, key IObject) IObject
	ValueClassForBinding(binding string) objc.Class
	WebPlugInContainerLoadRequestInFrame(request IObject, target IObject)
	WebPlugInContainerShowStatus(message IObject)
	WebPlugInDestroy()
	WebPlugInInitialize()
	WebPlugInMainResourceDidFailWithError(error_ IObject)
	WebPlugInMainResourceDidFinishLoading()
	WebPlugInMainResourceDidReceiveData(data IObject)
	WebPlugInMainResourceDidReceiveResponse(response IObject)
	WebPlugInSetIsSelected(isSelected bool)
	WebPlugInStart()
	WebPlugInStop()
	WillChangeValueForKeyWithSetMutationUsingObjects(key IObject, mutationKind uint, objects IObject)
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
func (oc _ObjectClass) AutomaticallyNotifiesObserversForKey(key IObject) bool {
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


// Returns the class object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/class
func (oc _ObjectClass) Class() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(oc.class), objc.Sel("class"))
	return rv
}


// Overridden to return the names of classes that can be used to decode objects if their class is unavailable.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classFallbacksForKeyedArchiver()
func (oc _ObjectClass) ClassFallbacksForKeyedArchiver() []string {
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
func (oc _ObjectClass) ConformsToProtocol(protocol_ IProtocol) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("conformsToProtocol:"), protocol_)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/debugDescription()
func (oc _ObjectClass) DebugDescription() IObject {
	rv := objc.Send[Object](objc.ID(oc.class), objc.Sel("debugDescription"))
	return rv
}


// Returns an object that will be used as the placeholder for the , when a key value coding compliant property of an instance of the receiving class returns the value specified by , and no other placeholder has been specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/defaultPlaceholder(for:with:)
func (oc _ObjectClass) DefaultPlaceholderForMarkerWithBinding(marker IObject, binding string) IObject {
	rv := objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("defaultPlaceholderForMarker:withBinding:"), marker, objc.String(binding))
	return Object{ID: rv}
}


// Returns a string that represents the contents of the receiving class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/description()
func (oc _ObjectClass) Description() IObject {
	rv := objc.Send[Object](objc.ID(oc.class), objc.Sel("description"))
	return rv
}


// Exposes the specified , advertising its availability.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/exposeBinding(_:)
func (oc _ObjectClass) ExposeBinding(binding string) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("exposeBinding:"), objc.String(binding))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/hash()
func (oc _ObjectClass) Hash() uint {
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
func (oc _ObjectClass) InstanceMethodForSelector(aSelector objc.SEL) IMP {
	rv := objc.Send[IMP](objc.ID(oc.class), objc.Sel("instanceMethodForSelector:"), aSelector)
	return rv
}


// Returns an object that contains a description of the instance method identified by a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instanceMethodSignatureForSelector:
func (oc _ObjectClass) InstanceMethodSignatureForSelector(aSelector objc.SEL) MethodSignature /* not a class type */ {
	rv := objc.Send[MethodSignature](objc.ID(oc.class), objc.Sel("instanceMethodSignatureForSelector:"), aSelector)
	return rv
}


// Returns a Boolean value that indicates whether instances of the receiver are capable of responding to a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instancesRespond(to:)
func (oc _ObjectClass) InstancesRespondToSelector(aSelector objc.SEL) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("instancesRespondToSelector:"), aSelector)
	return rv
}


// Returns whether a key should be hidden from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isKeyExcluded(fromWebScript:)
func (oc _ObjectClass) IsKeyExcludedFromWebScript(name unsafe.Pointer) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("isKeyExcludedFromWebScript:"), name)
	return rv
}


// Returns whether a selector should be hidden from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isSelectorExcluded(fromWebScript:)
func (oc _ObjectClass) IsSelectorExcludedFromWebScript(selector objc.SEL) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("isSelectorExcludedFromWebScript:"), selector)
	return rv
}


// Returns a Boolean value that indicates whether the receiving class is a subclass of, or identical to, a given class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isSubclass(of:)
func (oc _ObjectClass) IsSubclassOfClass(aClass objc.Class) bool {
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
func (oc _ObjectClass) ResolveClassMethod(sel objc.SEL) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("resolveClassMethod:"), sel)
	return rv
}


// Dynamically provides an implementation for a given selector for an instance method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/resolveInstanceMethod(_:)
func (oc _ObjectClass) ResolveInstanceMethod(sel objc.SEL) bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("resolveInstanceMethod:"), sel)
	return rv
}


// Sets as the default placeholder for the , when a key value coding compliant property of an instance of the receiving class returns the value specified by , and no other placeholder has been specified.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setDefaultPlaceholder(_:for:with:)
func (oc _ObjectClass) SetDefaultPlaceholderForMarkerWithBinding(placeholder IObject, marker IObject, binding string) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("setDefaultPlaceholder:forMarker:withBinding:"), placeholder, marker, objc.String(binding))
}


// Configures the observed object to post change notifications for a given property if any of the properties specified in a given array changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setKeys:triggerChangeNotificationsForDependentKey:
func (oc _ObjectClass) SetKeysTriggerChangeNotificationsForDependentKey(keys IObject, dependentKey IObject) {
	objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("setKeys:triggerChangeNotificationsForDependentKey:"), keys, dependentKey)
}


// Sets the receiver’s version number.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setVersion(_:)
func (oc _ObjectClass) SetVersion(aVersion objc.IObject /* cross-framework: Integer */) {
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
func (oc _ObjectClass) UseStoredAccessor() bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("useStoredAccessor"))
	return rv
}


// Returns the version number assigned to the class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/version()
func (oc _ObjectClass) Version() objc.IObject /* cross-framework: Integer */ {
	rv := objc.Send[objc.ID](objc.ID(oc.class), objc.Sel("version"))
	return rv
}


// Returns the scripting environment name for a selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webScriptName(for:)
func (oc _ObjectClass) WebScriptNameForSelector(selector objc.SEL) IObject {
	rv := objc.Send[Object](objc.ID(oc.class), objc.Sel("webScriptNameForSelector:"), selector)
	return rv
}


// Returns the scripting environment name for an attribute specified by a key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webScriptName(forKey:)
func (oc _ObjectClass) WebScriptNameForKey(name unsafe.Pointer) IObject {
	rv := objc.Send[Object](objc.ID(oc.class), objc.Sel("webScriptNameForKey:"), name)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/withL2CAPChannelRef:
func (oc _ObjectClass) WithL2CAPChannelRef(l2capChannelRef uintptr /* not a class type */) IObject {
	rv := objc.Send[Object](objc.ID(oc.class), objc.Sel("withL2CAPChannelRef:"), l2capChannelRef)
	return rv
}


// Returns a Boolean value that indicates whether the key-value coding methods should access the corresponding instance variable directly on finding no accessor method for a property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessInstanceVariablesDirectly
func (oc _ObjectClass) AccessInstanceVariablesDirectly() bool {
	rv := objc.Send[bool](objc.ID(oc.class), objc.Sel("accessInstanceVariablesDirectly"))
	return rv
}

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/acceptsPreviewPanelControl(_:)
func (o_ Object) AcceptsPreviewPanelControl(panel IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("acceptsPreviewPanelControl:"), panel)
	return rv
}


// Returns the count of the specified accessibility array attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityArrayAttributeCount(_:)
func (o_ Object) AccessibilityArrayAttributeCount(attribute string) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("accessibilityArrayAttributeCount:"), objc.String(attribute))
	return rv
}


// Returns a subarray of values of an accessibility array attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityArrayAttributeValues(_:index:maxCount:)
func (o_ Object) AccessibilityArrayAttributeValuesIndexMaxCount(attribute string, index uint, maxCount uint) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityArrayAttributeValues:index:maxCount:"), objc.String(attribute), index, maxCount)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityHitTest(_:)
func (o_ Object) AccessibilityHitTest(point IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityHitTest:"), point)
	return Object{ID: rv}
}


// Returns the index of the specified accessibility child in the parent.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityIndex(ofChild:)
func (o_ Object) AccessibilityIndexOfChild(child IObject) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("accessibilityIndexOfChild:"), child)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineEndPositionFromCurrentSelection()
func (o_ Object) AccessibilityLineEndPositionFromCurrentSelection() objc.IObject /* cross-framework: Integer */ {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityLineEndPositionFromCurrentSelection"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineRange(forPosition:)
func (o_ Object) AccessibilityLineRangeForPosition(position objc.IObject /* cross-framework: Integer */) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("accessibilityLineRangeForPosition:"), position)
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityLineStartPositionFromCurrentSelection()
func (o_ Object) AccessibilityLineStartPositionFromCurrentSelection() objc.IObject /* cross-framework: Integer */ {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityLineStartPositionFromCurrentSelection"))
	return rv
}


// Sent to the delegate to request the property the action applies to.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/actionProperty()
func (o_ Object) ActionProperty() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("actionProperty"))
	return rv
}


// Registers the observer object to receive KVO notifications for the key path relative to the object receiving this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/addObserver(_:forKeyPath:options:context:)
func (o_ Object) AddObserverForKeyPathOptionsContext(observer IObject, keyPath IObject, options uint, context unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addObserver:forKeyPath:options:context:"), observer, keyPath, options, context)
}


// Implemented to attempt a recovery from an error noted in an application-modal dialog.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attemptRecovery(fromError:optionIndex:)
func (o_ Object) AttemptRecoveryFromErrorOptionIndex(error_ IObject, recoveryOptionIndex uint) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("attemptRecoveryFromError:optionIndex:"), error_, recoveryOptionIndex)
	return rv
}


// Implemented to attempt a recovery from an error noted in a document-modal sheet.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attemptRecovery(fromError:optionIndex:delegate:didRecoverSelector:contextInfo:)
func (o_ Object) AttemptRecoveryFromErrorOptionIndexDelegateDidRecoverSelectorContextInfo(error_ IObject, recoveryOptionIndex uint, delegate IObject, didRecoverSelector objc.SEL, contextInfo unsafe.Pointer) {
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
func (o_ Object) AuthorizationViewShouldDeauthorize(view IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("authorizationViewShouldDeauthorize:"), view)
	return rv
}


// Overridden by subclasses to substitute another object in place of the object that was decoded and subsequently received this message.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/awakeAfter(using:)
func (o_ Object) AwakeAfterUsingCoder(coder IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("awakeAfterUsingCoder:"), coder)
	return Object{ID: rv}
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/beginPreviewPanelControl(_:)
func (o_ Object) BeginPreviewPanelControl(panel IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("beginPreviewPanelControl:"), panel)
}


// Establishes a binding between a given property of the receiver and the property of a given object specified by a given key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/bind(_:to:withKeyPath:options:)
func (o_ Object) BindToObjectWithKeyPathOptions(binding string, observable IObject, keyPath IObject, options IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("bind:toObject:withKeyPath:options:"), objc.String(binding), observable, keyPath, options)
}


// Returns the value for this element within the given range, as an attributed string.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityAttributedValue(in:)
func (o_ Object) BrowserAccessibilityAttributedValueInRange(range_ IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("browserAccessibilityAttributedValueInRange:"), range_)
	return rv
}


// Deletes text from the element at the current cursor position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityDeleteTextAtCursor(numberOfCharacters:)
func (o_ Object) BrowserAccessibilityDeleteTextAtCursor(numberOfCharacters objc.IObject /* cross-framework: Integer */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilityDeleteTextAtCursor:"), numberOfCharacters)
}


// Inserts text into the element at the current cursor position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityInsertTextAtCursor(text:)
func (o_ Object) BrowserAccessibilityInsertTextAtCursor(text IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilityInsertTextAtCursor:"), text)
}


// Returns the range of selected text in the element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilitySelectedTextRange()
func (o_ Object) BrowserAccessibilitySelectedTextRange() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("browserAccessibilitySelectedTextRange"))
	return rv
}


// Updates the element’s selected text.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilitySetSelectedTextRange(_:)
func (o_ Object) BrowserAccessibilitySetSelectedTextRange(range_ IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("browserAccessibilitySetSelectedTextRange:"), range_)
}


// Returns this element’s value in the given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityValue(in:)
func (o_ Object) BrowserAccessibilityValueInRange(range_ IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("browserAccessibilityValueInRange:"), range_)
	return rv
}


// Allows the delegate to handle the end-of-burn feedback.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/burnProgressPanel(_:burnDidFinish:)
func (o_ Object) BurnProgressPanelBurnDidFinish(theBurnPanel unsafe.Pointer, burn unsafe.Pointer) bool {
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


// Returns an array of candidates.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/candidates(_:)
func (o_ Object) Candidates(sender IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("candidates:"), sender)
	return rv
}


// Implements custom help behavior for the modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/certificatePanelShowHelp(_:)
func (o_ Object) CertificatePanelShowHelp(sender IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("certificatePanelShowHelp:"), sender)
	return rv
}


// Implements custom help behavior for the modal panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/chooseIdentityPanelShowHelp(_:)
func (o_ Object) ChooseIdentityPanelShowHelp(sender IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("chooseIdentityPanelShowHelp:"), sender)
	return rv
}


// Uses type info from the class description and to attempt to convert for to the proper type, if necessary.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/coerceValue(_:forKey:)
func (o_ Object) CoerceValueForKey(value IObject, key IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("coerceValue:forKey:"), value, key)
	return Object{ID: rv}
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
func (o_ Object) ComposedString(sender IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("composedString:"), sender)
	return Object{ID: rv}
}


// Creates and returns one or more scripting objects to be inserted into the specified relationship by copying the passed-in value and setting the properties in the copied object or objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copyScriptingValue(_:forKey:withProperties:)
func (o_ Object) CopyScriptingValueForKeyWithProperties(value IObject, key IObject, properties IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("copyScriptingValue:forKey:withProperties:"), value, key, properties)
	return Object{ID: rv}
}


// Deallocates the memory occupied by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/dealloc
func (o_ Object) Dealloc() {
	objc.Send[objc.ID](o_.ID, objc.Sel("dealloc"))
}


// Informs the observed object that the specified change has occurred on the indexes for a specified ordered to-many relationship.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/didChange(_:valuesAt:forKey:)
func (o_ Object) DidChangeValuesAtIndexesForKey(changeKind uint, indexes IObject, key IObject) {
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
func (o_ Object) DidChangeValueForKeyWithSetMutationUsingObjects(key IObject, mutationKind uint, objects IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("didChangeValueForKey:withSetMutation:usingObjects:"), key, mutationKind, objects)
}


// Processes a command generated by user action such as typing certain keys or pressing the mouse button.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/didCommand(by:client:)
func (o_ Object) DidCommandBySelectorClient(aSelector objc.SEL, sender IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("didCommandBySelector:client:"), aSelector, sender)
	return rv
}


// Returns a Boolean value that indicates whether the receiver contains a given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/doesContain(_:)
func (o_ Object) DoesContain(object IObject) bool {
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
func (o_ Object) EndPreviewPanelControl(panel IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("endPreviewPanelControl:"), panel)
}


// Notification sent by the panel before display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/eraseProgressPanel(_:eraseDidFinish:)
func (o_ Object) EraseProgressPanelEraseDidFinish(theErasePanel unsafe.Pointer, erase unsafe.Pointer) bool {
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
func (o_ Object) ExceptionHandlerShouldHandleExceptionMask(sender IObject, exception IObject, aMask uint) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("exceptionHandler:shouldHandleException:mask:"), sender, exception, aMask)
	return rv
}


// Implemented by the delegate to evaluate whether the delegating exception hangler should log a given exception.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/exceptionHandler(_:shouldLogException:mask:)
func (o_ Object) ExceptionHandlerShouldLogExceptionMask(sender IObject, exception IObject, aMask uint) bool {
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


// Overridden by subclasses to forward messages to other objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/forwardInvocation:
func (o_ Object) ForwardInvocation(anInvocation IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("forwardInvocation:"), anInvocation)
}


// Returns the object to which unrecognized messages should first be directed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/forwardingTarget(for:)
func (o_ Object) ForwardingTargetForSelector(aSelector objc.SEL) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("forwardingTargetForSelector:"), aSelector)
	return Object{ID: rv}
}


// Handles key down and mouse events.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/handle(_:client:)
func (o_ Object) HandleEventClient(event IObject, sender IObject) bool {
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
func (o_ Object) ImageBrowserCellWasDoubleClickedAtIndex(aBrowser IObject, index uint) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:cellWasDoubleClickedAtIndex:"), aBrowser, index)
}


// Performs custom tasks when the user right-clicks an item in the image browser view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:cellWasRightClickedAt:with:)
func (o_ Object) ImageBrowserCellWasRightClickedAtIndexWithEvent(aBrowser IObject, index uint, event IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:cellWasRightClickedAtIndex:withEvent:"), aBrowser, index, event)
}


// Returns the group at the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:groupAt:)
func (o_ Object) ImageBrowserGroupAtIndex(aBrowser IObject, index uint) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("imageBrowser:groupAtIndex:"), aBrowser, index)
	return rv
}


// Returns an object for the item in an image browser view that corresponds to the specified index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:itemAt:)
func (o_ Object) ImageBrowserItemAtIndex(aBrowser IObject, index uint) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("imageBrowser:itemAtIndex:"), aBrowser, index)
	return Object{ID: rv}
}


// Signals that the specified items should be moved to the specified destination.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageBrowser(_:moveItemsAt:to:)
func (o_ Object) ImageBrowserMoveItemsAtIndexesToIndex(aBrowser IObject, indexes IObject, destinationIndex uint) bool {
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
func (o_ Object) ImageBrowserWriteItemsAtIndexesToPasteboard(aBrowser IObject, itemIndexes IObject, pasteboard IObject) uint {
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
func (o_ Object) ImageRepresentation() IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("imageRepresentation"))
	return Object{ID: rv}
}


// Returns the representation type of the image to display.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageRepresentationType()
func (o_ Object) ImageRepresentationType() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("imageRepresentationType"))
	return rv
}


// Returns the display subtitle of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageSubtitle()
func (o_ Object) ImageSubtitle() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("imageSubtitle"))
	return rv
}


// Returns the display title of the image.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageTitle()
func (o_ Object) ImageTitle() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("imageTitle"))
	return rv
}


// Returns a unique string that identifies the data source item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageUID()
func (o_ Object) ImageUID() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("imageUID"))
	return rv
}


// Returns the version of the item.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/imageVersion()
func (o_ Object) ImageVersion() uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("imageVersion"))
	return rv
}


// Returns the indices of the specified container objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/indicesOfObjects(byEvaluatingObjectSpecifier:)
func (o_ Object) IndicesOfObjectsByEvaluatingObjectSpecifier(specifier IObject) []objc.ID {
	rv := objc.Send[[]foundation.Number](o_.ID, objc.Sel("indicesOfObjectsByEvaluatingObjectSpecifier:"), specifier)
	return rv
}


// Handles key down events that do not map to an action method.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/inputText(_:client:)
func (o_ Object) InputTextClient(string_ IObject, sender IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("inputText:client:"), string_, sender)
	return rv
}


// Receives Unicode, the key code that generated it, and any modifier flags.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/inputText(_:key:modifiers:client:)
func (o_ Object) InputTextKeyModifiersClient(string_ IObject, keyCode objc.IObject /* cross-framework: Integer */, flags uint, sender IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("inputText:key:modifiers:client:"), string_, keyCode, flags, sender)
	return rv
}


// Inserts an object at the specified index in the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/insertValue(_:at:inPropertyWithKey:)
func (o_ Object) InsertValueAtIndexInPropertyWithKey(value IObject, index uint, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("insertValue:atIndex:inPropertyWithKey:"), value, index, key)
}


// Inserts an object in the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/insertValue(_:inPropertyWithKey:)
func (o_ Object) InsertValueInPropertyWithKey(value IObject, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("insertValue:inPropertyWithKey:"), value, key)
}


// For a given key that defines the name of the relationship from the receiver’s class to another class, returns the name of the relationship from the other class to the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/inverse(forRelationshipKey:)
func (o_ Object) InverseForRelationshipKey(relationshipKey IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("inverseForRelationshipKey:"), relationshipKey)
	return rv
}


// Executes when a script attempts to invoke a method on an exposed object directly.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/invokeDefaultMethod(withArguments:)
func (o_ Object) InvokeDefaultMethodWithArguments(arguments IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("invokeDefaultMethodWithArguments:"), arguments)
	return Object{ID: rv}
}


// Handles undefined method invocation from the scripting environment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/invokeUndefinedMethod(fromWebScript:withArguments:)
func (o_ Object) InvokeUndefinedMethodFromWebScriptWithArguments(name IObject, arguments IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("invokeUndefinedMethodFromWebScript:withArguments:"), name, arguments)
	return Object{ID: rv}
}


// Returns a Boolean value that indicates whether receiver is considered to be “like” a given string when the case of characters in the receiver is ignored.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isCaseInsensitiveLike(_:)
func (o_ Object) IsCaseInsensitiveLike(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isCaseInsensitiveLike:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is equal to another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isEqual(to:)
func (o_ Object) IsEqualTo(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isEqualTo:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is greater than another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isGreaterThan(_:)
func (o_ Object) IsGreaterThan(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isGreaterThan:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is greater than or equal to another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isGreaterThanOrEqual(to:)
func (o_ Object) IsGreaterThanOrEqualTo(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isGreaterThanOrEqualTo:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is less than another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isLessThan(_:)
func (o_ Object) IsLessThan(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isLessThan:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is less than or equal to another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isLessThanOrEqual(to:)
func (o_ Object) IsLessThanOrEqualTo(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isLessThanOrEqualTo:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is “like” another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isLike(_:)
func (o_ Object) IsLike(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isLike:"), object)
	return rv
}


// Returns a Boolean value that indicates whether the receiver is not equal to another given object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isNotEqual(to:)
func (o_ Object) IsNotEqualTo(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isNotEqualTo:"), object)
	return rv
}


// Locates and returns the address of the receiver’s implementation of a method so it can be called as a function.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/method(for:)
func (o_ Object) MethodForSelector(aSelector objc.SEL) IMP {
	rv := objc.Send[IMP](o_.ID, objc.Sel("methodForSelector:"), aSelector)
	return rv
}


// Returns an object that contains a description of the method identified by a given selector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/methodSignatureForSelector:
func (o_ Object) MethodSignatureForSelector(aSelector objc.SEL) MethodSignature /* not a class type */ {
	rv := objc.Send[MethodSignature](o_.ID, objc.Sel("methodSignatureForSelector:"), aSelector)
	return rv
}


// Returns a mutable array proxy that provides read-write access to an ordered to-many relationship specified by a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/mutableArrayValue(forKey:)
func (o_ Object) MutableArrayValueForKey(key IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("mutableArrayValueForKey:"), key)
	return rv
}


// Returns a mutable array that provides read-write access to the ordered to-many relationship specified by a given key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/mutableArrayValue(forKeyPath:)
func (o_ Object) MutableArrayValueForKeyPath(keyPath IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("mutableArrayValueForKeyPath:"), keyPath)
	return rv
}


// Returns a mutable ordered set that provides read-write access to the uniquing ordered to-many relationship specified by a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/mutableOrderedSetValue(forKey:)
func (o_ Object) MutableOrderedSetValueForKey(key IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("mutableOrderedSetValueForKey:"), key)
	return rv
}


// Returns a mutable ordered set that provides read-write access to the uniquing ordered to-many relationship specified by a given key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/mutableOrderedSetValue(forKeyPath:)
func (o_ Object) MutableOrderedSetValueForKeyPath(keyPath IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("mutableOrderedSetValueForKeyPath:"), keyPath)
	return rv
}


// Returns a mutable set proxy that provides read-write access to the unordered to-many relationship specified by a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/mutableSetValue(forKey:)
func (o_ Object) MutableSetValueForKey(key IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("mutableSetValueForKey:"), key)
	return rv
}


// Returns a mutable set that provides read-write access to the unordered to-many relationship specified by a given key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/mutableSetValue(forKeyPath:)
func (o_ Object) MutableSetValueForKeyPath(keyPath IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("mutableSetValueForKeyPath:"), keyPath)
	return rv
}


// Creates and returns an instance of a scriptable class, setting its contents and properties, for insertion into the relationship identified by the key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/newScriptingObject(of:forValueForKey:withContentsValue:properties:)
func (o_ Object) NewScriptingObjectOfClassForValueForKeyWithContentsValueProperties(objectClass objc.Class, key IObject, contentsValue IObject, properties IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("newScriptingObjectOfClass:forValueForKey:withContentsValue:properties:"), objectClass, key, contentsValue, properties)
	return Object{ID: rv}
}


// Returns the number of groups in an image browser view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/numberOfGroups(inImageBrowser:)
func (o_ Object) NumberOfGroupsInImageBrowser(aBrowser IObject) uint {
	rv := objc.Send[uint](o_.ID, objc.Sel("numberOfGroupsInImageBrowser:"), aBrowser)
	return rv
}


// Returns the number of records managed by the data source object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/numberOfItems(inImageBrowser:)
func (o_ Object) NumberOfItemsInImageBrowser(aBrowser IObject) uint {
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
func (o_ Object) OptionDescriptionsForBinding(binding string) []objc.ID {
	rv := objc.Send[[]coredata.AttributeDescription](o_.ID, objc.Sel("optionDescriptionsForBinding:"), objc.String(binding))
	return rv
}


// Return the string that consists of the precomposed Unicode characters.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/originalString(_:)
func (o_ Object) OriginalString(sender IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("originalString:"), sender)
	return rv
}


// Invokes a method of the receiver on the specified thread using the default mode.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:on:with:waitUntilDone:)
func (o_ Object) PerformSelectorOnThreadWithObjectWaitUntilDone(aSelector objc.SEL, thr IObject, arg IObject, wait bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:onThread:withObject:waitUntilDone:"), aSelector, thr, arg, wait)
}


// Invokes a method of the receiver on the specified thread using the specified modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:on:with:waitUntilDone:modes:)
func (o_ Object) PerformSelectorOnThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, thr IObject, arg IObject, wait bool, array []string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:onThread:withObject:waitUntilDone:modes:"), aSelector, thr, arg, wait, array)
}


// Invokes a method of the receiver on the current thread using the default mode after a delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:with:afterDelay:)
func (o_ Object) PerformSelectorWithObjectAfterDelay(aSelector objc.SEL, anArgument IObject, delay float64) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelector:withObject:afterDelay:"), aSelector, anArgument, delay)
}


// Invokes a method of the receiver on the current thread using the specified modes after a delay.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/perform(_:with:afterDelay:inModes:)
func (o_ Object) PerformSelectorWithObjectAfterDelayInModes(aSelector objc.SEL, anArgument IObject, delay float64, modes []string) {
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
func (o_ Object) PerformSelectorOnMainThreadWithObjectWaitUntilDone(aSelector objc.SEL, arg IObject, wait bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("performSelectorOnMainThread:withObject:waitUntilDone:"), aSelector, arg, wait)
}


// Invokes a method of the receiver on the main thread using the specified modes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performSelector(onMainThread:with:waitUntilDone:modes:)
func (o_ Object) PerformSelectorOnMainThreadWithObjectWaitUntilDoneModes(aSelector objc.SEL, arg IObject, wait bool, array []string) {
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
func (o_ Object) ReadLinkQualityForDeviceCompleteDeviceInfoError(controller IObject, device IObject, info IObject, error_ int) {
	objc.Send[objc.ID](o_.ID, objc.Sel("readLinkQualityForDeviceComplete:device:info:error:"), controller, device, info, error_)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/readRSSI(forDeviceComplete:device:info:error:)
func (o_ Object) ReadRSSIForDeviceCompleteDeviceInfoError(controller IObject, device IObject, info IObject, error_ int) {
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
func (o_ Object) RemoveValueAtIndexFromPropertyWithKey(index uint, key IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("removeValueAtIndex:fromPropertyWithKey:"), index, key)
}


// Replaces the object at the specified index in the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replaceValue(at:inPropertyWithKey:withValue:)
func (o_ Object) ReplaceValueAtIndexInPropertyWithKeyWithValue(index uint, key IObject, value IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("replaceValueAtIndex:inPropertyWithKey:withValue:"), index, key, value)
}


// Overridden by subclasses to substitute another object for itself during encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replacementObject(for:)-2l8ox
func (o_ Object) ReplacementObjectForCoder(coder IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("replacementObjectForCoder:"), coder)
	return Object{ID: rv}
}


// Overridden by subclasses to substitute another object for itself during keyed archiving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/replacementObject(for:)-60vwc
func (o_ Object) ReplacementObjectForKeyedArchiver(archiver IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("replacementObjectForKeyedArchiver:"), archiver)
	return Object{ID: rv}
}


// Called to determine if the specified uniform type identifier should be shown in the save panel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/saveOptions(_:shouldShowUTType:)
func (o_ Object) SaveOptionsShouldShowUTType(saveOptions IObject, utType IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("saveOptions:shouldShowUTType:"), saveOptions, utType)
	return rv
}


// Returns if, in a scripting comparison, the compared object matches the beginning of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingBegins(with:)
func (o_ Object) ScriptingBeginsWith(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingBeginsWith:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object contains .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingContains(_:)
func (o_ Object) ScriptingContains(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingContains:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object matches the end of .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingEnds(with:)
func (o_ Object) ScriptingEndsWith(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingEndsWith:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is equal to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsEqual(to:)
func (o_ Object) ScriptingIsEqualTo(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsEqualTo:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is greater than .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsGreaterThan(_:)
func (o_ Object) ScriptingIsGreaterThan(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsGreaterThan:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is greater than or equal to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsGreaterThanOrEqual(to:)
func (o_ Object) ScriptingIsGreaterThanOrEqualTo(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsGreaterThanOrEqualTo:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is less than .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsLessThan(_:)
func (o_ Object) ScriptingIsLessThan(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsLessThan:"), object)
	return rv
}


// Returns if, in a scripting comparison, the compared object is less than or equal to .
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingIsLessThanOrEqual(to:)
func (o_ Object) ScriptingIsLessThanOrEqualTo(object IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("scriptingIsLessThanOrEqualTo:"), object)
	return rv
}


// Given an object specifier, returns the specified object or objects in the receiving container.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/scriptingValue(for:)
func (o_ Object) ScriptingValueForSpecifier(objectSpecifier IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("scriptingValueForSpecifier:"), objectSpecifier)
	return Object{ID: rv}
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
func (o_ Object) SetupPanelDeviceContainsSuitableMediaPromptString(aPanel unsafe.Pointer, device unsafe.Pointer, prompt IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setupPanel:deviceContainsSuitableMedia:promptString:"), aPanel, device, prompt)
	return rv
}


// Allows the delegate to determine if device can be used as a target.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/setupPanel(_:deviceCouldBeTarget:)
func (o_ Object) SetupPanelDeviceCouldBeTarget(aPanel unsafe.Pointer, device unsafe.Pointer) bool {
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
func (o_ Object) SetupPanelShouldHandleMediaReservations(aPanel unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("setupPanelShouldHandleMediaReservations:"), aPanel)
	return rv
}


// Sent to the delegate to determine whether the action should be enabled.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/shouldEnableAction(for:identifier:)
func (o_ Object) ShouldEnableActionForPersonIdentifier(person IObject, identifier IObject) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("shouldEnableActionForPerson:identifier:"), person, identifier)
	return rv
}


// Sent to the delegate to request the title of the menu item for the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/title(for:identifier:)
func (o_ Object) TitleForPersonIdentifier(person IObject, identifier IObject) IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("titleForPerson:identifier:"), person, identifier)
	return rv
}


// Removes a given binding between the receiver and a controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/unbind(_:)
func (o_ Object) Unbind(binding string) {
	objc.Send[objc.ID](o_.ID, objc.Sel("unbind:"), objc.String(binding))
}


// Indicates whether the value specified by a given pointer is valid, or can be made valid, for the property identified by a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/validateValue(_:forKey:)
func (o_ Object) ValidateValueForKeyError(ioValue unsafe.Pointer, inKey IObject, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("validateValue:forKey:error:"), ioValue, inKey, outError)
	return rv
}


// Indicates whether the value specified by a given pointer is not valid for a given key path relative to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/validateValue(_:forKeyPath:)
func (o_ Object) ValidateValueForKeyPathError(ioValue unsafe.Pointer, inKeyPath IObject, outError unsafe.Pointer) bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("validateValue:forKeyPath:error:"), ioValue, inKeyPath, outError)
	return rv
}


// Retrieves an indexed object from the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(at:inPropertyWithKey:)
func (o_ Object) ValueAtIndexInPropertyWithKey(index uint, key IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueAtIndex:inPropertyWithKey:"), index, key)
	return Object{ID: rv}
}


// Returns the value for the property identified by a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forKey:)
func (o_ Object) ValueForKey(key IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueForKey:"), key)
	return Object{ID: rv}
}


// Returns the value for the derived property identified by a given key path.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forKeyPath:)
func (o_ Object) ValueForKeyPath(keyPath IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueForKeyPath:"), keyPath)
	return Object{ID: rv}
}


// Invoked by when it finds no property corresponding to a given key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(forUndefinedKey:)
func (o_ Object) ValueForUndefinedKey(key IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueForUndefinedKey:"), key)
	return Object{ID: rv}
}


// Retrieves a named object from the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(withName:inPropertyWithKey:)
func (o_ Object) ValueWithNameInPropertyWithKey(name IObject, key IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueWithName:inPropertyWithKey:"), name, key)
	return Object{ID: rv}
}


// Retrieves an object by ID from the collection specified by the passed key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/value(withUniqueID:inPropertyWithKey:)
func (o_ Object) ValueWithUniqueIDInPropertyWithKey(uniqueID IObject, key IObject) IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("valueWithUniqueID:inPropertyWithKey:"), uniqueID, key)
	return Object{ID: rv}
}


// Returns the class of the value that will be returned for the specified binding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/valueClassForBinding(_:)
func (o_ Object) ValueClassForBinding(binding string) objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("valueClassForBinding:"), objc.String(binding))
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
func (o_ Object) WebPlugInSetIsSelected(isSelected bool) {
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
func (o_ Object) WillChangeValuesAtIndexesForKey(changeKind uint, indexes IObject, key IObject) {
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
func (o_ Object) WillChangeValueForKeyWithSetMutationUsingObjects(key IObject, mutationKind uint, objects IObject) {
	objc.Send[objc.ID](o_.ID, objc.Sel("willChangeValueForKey:withSetMutation:usingObjects:"), key, mutationKind, objects)
}


// Returns a Boolean value that indicates whether the key-value coding methods should access the corresponding instance variable directly on finding no accessor method for a property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessInstanceVariablesDirectly
func (o_ Object) AccessInstanceVariablesDirectly() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessInstanceVariablesDirectly"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityFocusedUIElement
func (o_ Object) AccessibilityFocusedUIElement() IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("accessibilityFocusedUIElement"))
	return Object{ID: rv}
}


// A Boolean value that indicates whether a custom accessibility object sends a notification when its corresponding UI element is destroyed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/accessibilityNotifiesWhenDestroyed
func (o_ Object) AccessibilityNotifiesWhenDestroyed() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("accessibilityNotifiesWhenDestroyed"))
	return rv
}


// An array of objects containing the names of immutable values that instances of the receiver’s class contain.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attributeKeys
func (o_ Object) AttributeKeys() []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("attributeKeys"))
	return rv
}


// A proxy for the receiving object
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/autoContentAccessingProxy
func (o_ Object) AutoContentAccessingProxy() IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("autoContentAccessingProxy"))
	return Object{ID: rv}
}


// The kind of container that contains this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityContainerType
func (o_ Object) BrowserAccessibilityContainerType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("browserAccessibilityContainerType"))
	return rv
}


// The kind of container that contains this element.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityContainerType
func (o_ Object) SetBrowserAccessibilityContainerType(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBrowserAccessibilityContainerType:"), value)
}


// A Boolean value that indicates whether the element has native focus in the browser Document Object Model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityHasDOMFocus
func (o_ Object) BrowserAccessibilityHasDOMFocus() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("browserAccessibilityHasDOMFocus"))
	return rv
}


// A Boolean value that indicates whether the element has native focus in the browser Document Object Model.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityHasDOMFocus
func (o_ Object) SetBrowserAccessibilityHasDOMFocus(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBrowserAccessibilityHasDOMFocus:"), value)
}


// A Boolean value that’s the element’s value for aria-required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityIsRequired
func (o_ Object) BrowserAccessibilityIsRequired() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("browserAccessibilityIsRequired"))
	return rv
}


// A Boolean value that’s the element’s value for aria-required.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityIsRequired
func (o_ Object) SetBrowserAccessibilityIsRequired(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBrowserAccessibilityIsRequired:"), value)
}


// The element’s value for aria-pressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityPressedState
func (o_ Object) BrowserAccessibilityPressedState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("browserAccessibilityPressedState"))
	return rv
}


// The element’s value for aria-pressed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/browserAccessibilityPressedState
func (o_ Object) SetBrowserAccessibilityPressedState(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setBrowserAccessibilityPressedState:"), value)
}


// The receiver’s Apple event type code, as stored in the object for the object’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classCode
func (o_ Object) ClassCode() uint32 /* not a class type */ {
	rv := objc.Send[uint32](o_.ID, objc.Sel("classCode"))
	return rv
}


// An object containing information about the attributes and relationships of the receiver’s class.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classDescription
func (o_ Object) ClassDescription() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("classDescription"))
	return rv
}


// The class to substitute for the receiver’s own class during archiving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classForArchiver
func (o_ Object) ClassForArchiver() objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("classForArchiver"))
	return rv
}


// Overridden by subclasses to substitute a class other than its own during coding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classForCoder
func (o_ Object) ClassForCoder() objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("classForCoder"))
	return rv
}


// Subclasses to substitute a new class for instances during keyed archiving.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classForKeyedArchiver
func (o_ Object) ClassForKeyedArchiver() objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("classForKeyedArchiver"))
	return rv
}


// Class to substitute for the receiver in distribution encoding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/classForPortCoder
func (o_ Object) ClassForPortCoder() objc.Class {
	rv := objc.Send[objc.Class](o_.ID, objc.Sel("classForPortCoder"))
	return rv
}


// Returns an array containing the bindings exposed by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/exposedBindings
func (o_ Object) ExposedBindings() []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("exposedBindings"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/isSelectable
func (o_ Object) Selectable() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("selectable"))
	return rv
}


// Returns an object that exposes the plug-in’s scripting interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/objectForWebScript
func (o_ Object) ObjectForWebScript() IObject {
	rv := objc.Send[objc.ID](o_.ID, objc.Sel("objectForWebScript"))
	return Object{ID: rv}
}


// Returns an object specifier for the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/objectSpecifier
func (o_ Object) ObjectSpecifier() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("objectSpecifier"))
	return rv
}


// Returns a pointer that identifies information about all of the observers that are registered with the observed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/observationInfo
func (o_ Object) ObservationInfo() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o_.ID, objc.Sel("observationInfo"))
	return rv
}


// Returns a pointer that identifies information about all of the observers that are registered with the observed object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/observationInfo
func (o_ Object) SetObservationInfo(value unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setObservationInfo:"), value)
}


// An array containing the keys for the to-many relationship properties of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/toManyRelationshipKeys
func (o_ Object) ToManyRelationshipKeys() []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("toManyRelationshipKeys"))
	return rv
}


// The keys for the to-one relationship properties of the receiver, if any.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/toOneRelationshipKeys
func (o_ Object) ToOneRelationshipKeys() []string {
	rv := objc.Send[[]string](o_.ID, objc.Sel("toOneRelationshipKeys"))
	return rv
}


// Returns the plug-in selection color.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/webPlugInContainerSelectionColor
func (o_ Object) WebPlugInContainerSelectionColor() IObject {
	rv := objc.Send[Object](o_.ID, objc.Sel("webPlugInContainerSelectionColor"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/isselectable
func (o_ Object) IsSelectable() bool {
	rv := objc.Send[bool](o_.ID, objc.Sel("isSelectable"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject-swift.class/isselectable
func (o_ Object) SetIsSelectable(value bool) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setIsSelectable:"), value)
}


