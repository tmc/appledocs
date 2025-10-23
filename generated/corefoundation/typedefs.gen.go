// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

// Type aliases and typedefs
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
// Bit - A binary value of either   or  .
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFBit
// CFBit has base type: UInt32
type Bit uintptr
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
// ByteOrder - Flags that identify byte order.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFByteOrder
// CFByteOrder has base type: CFIndex
type ByteOrder uintptr
// CalendarRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendar
// CFCalendarRef has base type: struct __CFCalendar *
type CalendarRef uintptr
// CalendarIdentifier type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFCalendarIdentifier
// CFCalendarIdentifier has base type: CFStringRef
type CalendarIdentifier uintptr
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
// ErrorDomain type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFErrorDomain
// CFErrorDomain has base type: CFStringRef
type ErrorDomain uintptr
// FileDescriptorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptor
// CFFileDescriptorRef has base type: struct __CFFileDescriptor *
type FileDescriptorRef uintptr
// FileDescriptorCallBack - Defines a structure for a callback for a CFFileDescriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorCallBack
// CFFileDescriptorCallBack has base type: void (*)(struct __CFFileDescriptor *, unsigned long, void *)
type FileDescriptorCallBack uintptr
// FileDescriptorNativeDescriptor - Defines a type for the native file descriptor.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileDescriptorNativeDescriptor
type FileDescriptorNativeDescriptor int32
// FileSecurityRef - Encapsulates a file system object’s security information in a Core Foundation object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurity
// CFFileSecurityRef has base type: struct __CFFileSecurity *
type FileSecurityRef uintptr
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
// LocaleKey type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFLocaleKey
// CFLocaleKey has base type: CFStringRef
type LocaleKey uintptr
// MachPortRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPort
// CFMachPortRef has base type: struct __CFMachPort *
type MachPortRef uintptr
// MachPortCallBack - Callback invoked to process a message received on a CFMachPort object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortCallBack
// CFMachPortCallBack has base type: void (*)(struct __CFMachPort *, void *, long, void *)
type MachPortCallBack uintptr
// MachPortInvalidationCallBack - Callback invoked when a CFMachPort object is invalidated.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMachPortInvalidationCallBack
// CFMachPortInvalidationCallBack has base type: void (*)(struct __CFMachPort *, void *)
type MachPortInvalidationCallBack uintptr
// MessagePortRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePort
// CFMessagePortRef has base type: struct __CFMessagePort *
type MessagePortRef uintptr
// MessagePortCallBack - Callback invoked to process a message received on a CFMessagePort object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortCallBack
// CFMessagePortCallBack has base type: const struct __CFData *(*)(struct __CFMessagePort *, int, const struct __CFData *, void *)
type MessagePortCallBack uintptr
// MessagePortInvalidationCallBack - Callback invoked when a CFMessagePort object is invalidated.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMessagePortInvalidationCallBack
// CFMessagePortInvalidationCallBack has base type: void (*)(struct __CFMessagePort *, void *)
type MessagePortInvalidationCallBack uintptr
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
// MutableBitVectorRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableBitVector
// CFMutableBitVectorRef has base type: struct __CFBitVector *
type MutableBitVectorRef uintptr
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
// MutableSetRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableSet
// CFMutableSetRef has base type: struct __CFSet *
type MutableSetRef uintptr
// MutableStringRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFMutableString
// CFMutableStringRef has base type: struct __CFString *
type MutableStringRef uintptr
// NotificationCallback - Callback function invoked for each observer of a notification when the notification is posted.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCallback
// CFNotificationCallback has base type: void (*)(struct __CFNotificationCenter *, void *, const struct __CFString *, const void *, const struct __CFDictionary *)
type NotificationCallback uintptr
// NotificationCenterRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationCenter
// CFNotificationCenterRef has base type: struct __CFNotificationCenter *
type NotificationCenterRef uintptr
// NotificationName type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNotificationName
// CFNotificationName has base type: CFStringRef
type NotificationName uintptr
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
// NumberFormatterRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatter
// CFNumberFormatterRef has base type: struct __CFNumberFormatter *
type NumberFormatterRef uintptr
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
// PlugInDynamicRegisterFunction - A callback which provides a plug-in the opportunity to dynamically register its types with a host.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInDynamicRegisterFunction
// CFPlugInDynamicRegisterFunction has base type: void (*)(struct __CFBundle *)
type PlugInDynamicRegisterFunction uintptr
// PlugInFactoryFunction - Callback function that a plug-in author must implement to create a plug-in instance.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInFactoryFunction
// CFPlugInFactoryFunction has base type: void *(*)(const struct __CFAllocator *, const struct __CFUUID *)
type PlugInFactoryFunction uintptr
// PlugInInstanceRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstance
// CFPlugInInstanceRef has base type: struct __CFPlugInInstance *
type PlugInInstanceRef uintptr
// PlugInInstanceDeallocateInstanceDataFunction - Not recommended.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceDeallocateInstanceDataFunction
// CFPlugInInstanceDeallocateInstanceDataFunction has base type: void (*)(void *)
type PlugInInstanceDeallocateInstanceDataFunction uintptr
// PlugInInstanceGetInterfaceFunction - Not recommended.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInInstanceGetInterfaceFunction
// CFPlugInInstanceGetInterfaceFunction has base type: unsigned char (*)(struct __CFPlugInInstance *, const struct __CFString *, void **)
type PlugInInstanceGetInterfaceFunction uintptr
// PlugInUnloadFunction - Callback function that is called, if present, just before a plug-in’s code is unloaded.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPlugInUnloadFunction
// CFPlugInUnloadFunction has base type: void (*)(struct __CFBundle *)
type PlugInUnloadFunction uintptr
// PropertyListRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFPropertyList
// CFPropertyListRef has base type: CFTypeRef
type PropertyListRef uintptr
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
// RunLoopObserverRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserver
// CFRunLoopObserverRef has base type: struct __CFRunLoopObserver *
type RunLoopObserverRef uintptr
// RunLoopObserverCallBack - Callback invoked when a CFRunLoopObserver object is fired.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopObserverCallBack
// CFRunLoopObserverCallBack has base type: void (*)(struct __CFRunLoopObserver *, enum CFRunLoopActivity, void *)
type RunLoopObserverCallBack uintptr
// RunLoopSourceRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopSource
// CFRunLoopSourceRef has base type: struct __CFRunLoopSource *
type RunLoopSourceRef uintptr
// RunLoopTimerRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimer
// CFRunLoopTimerRef has base type: struct __CFRunLoopTimer *
type RunLoopTimerRef uintptr
// RunLoopTimerCallBack - Callback invoked when a CFRunLoopTimer object fires.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimerCallBack
// CFRunLoopTimerCallBack has base type: void (*)(struct __CFRunLoopTimer *, void *)
type RunLoopTimerCallBack uintptr
// SetRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSet
// CFSetRef has base type: const struct __CFSet *
type SetRef uintptr
// SetApplierFunction - Prototype of a callback function that may be applied to every value in a set.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetApplierFunction
// CFSetApplierFunction has base type: void (*)(const void *, void *)
type SetApplierFunction uintptr
// SetCopyDescriptionCallBack - Prototype of a callback function used to get a description of a value in a set.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetCopyDescriptionCallBack
// CFSetCopyDescriptionCallBack has base type: const struct __CFString *(*)(const void *)
type SetCopyDescriptionCallBack uintptr
// SetEqualCallBack - Prototype of a callback function used to determine if two values in a set are equal.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetEqualCallBack
// CFSetEqualCallBack has base type: unsigned char (*)(const void *, const void *)
type SetEqualCallBack uintptr
// SetHashCallBack - Prototype of a callback function called to compute a hash code for a value. Hash codes are used when values are accessed, added, or removed from a collection.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetHashCallBack
// CFSetHashCallBack has base type: unsigned long (*)(const void *)
type SetHashCallBack uintptr
// SetReleaseCallBack - Prototype of a callback function used to release a value before it’s removed from a set.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetReleaseCallBack
// CFSetReleaseCallBack has base type: void (*)(const struct __CFAllocator *, const void *)
type SetReleaseCallBack uintptr
// SetRetainCallBack - Prototype of a callback function used to retain a value being added to a set.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSetRetainCallBack
// CFSetRetainCallBack has base type: const void *(*)(const struct __CFAllocator *, const void *)
type SetRetainCallBack uintptr
// SocketRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocket
// CFSocketRef has base type: struct __CFSocket *
type SocketRef uintptr
// SocketCallBack - Callback invoked when certain types of activity takes place on a CFSocket object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBack
// CFSocketCallBack has base type: void (*)(struct __CFSocket *, enum CFSocketCallBackType, const struct __CFData *, const void *, void *)
type SocketCallBack uintptr
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
// TimeInterval - Type used to represent elapsed time in seconds.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeInterval
// CFTimeInterval has base type: double
type TimeInterval uintptr
// TimeZoneRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZone
// CFTimeZoneRef has base type: const struct __CFTimeZone *
type TimeZoneRef uintptr
// TreeRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTree
// CFTreeRef has base type: struct __CFTree *
type TreeRef uintptr
// TreeApplierFunction - Type of the callback function used by the CFTree apply function.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeApplierFunction
// CFTreeApplierFunction has base type: void (*)(const void *, void *)
type TreeApplierFunction uintptr
// TreeCopyDescriptionCallBack - Callback function used to provide a description of the program-defined information pointer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeCopyDescriptionCallBack
// CFTreeCopyDescriptionCallBack has base type: const struct __CFString *(*)(const void *)
type TreeCopyDescriptionCallBack uintptr
// TreeReleaseCallBack - Callback function used to release a previously retained program-defined information pointer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeReleaseCallBack
// CFTreeReleaseCallBack has base type: void (*)(const void *)
type TreeReleaseCallBack uintptr
// TreeRetainCallBack - Callback function used to retain a program-defined information pointer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFTreeRetainCallBack
// CFTreeRetainCallBack has base type: const void *(*)(const void *)
type TreeRetainCallBack uintptr
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
// URLBookmarkFileCreationOptions - Type for bookmark file creation options.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkFileCreationOptions
// CFURLBookmarkFileCreationOptions has base type: CFOptionFlags
type URLBookmarkFileCreationOptions uintptr
// URLEnumeratorRef - A reference to a   object.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumerator
// CFURLEnumeratorRef has base type: const struct __CFURLEnumerator *
type URLEnumeratorRef uintptr
// UUIDRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUUID
// CFUUIDRef has base type: const struct __CFUUID *
type UUIDRef uintptr
// UserNotificationRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotification
// CFUserNotificationRef has base type: struct __CFUserNotification *
type UserNotificationRef uintptr
// UserNotificationCallBack - Callback invoked when an asynchronous user notification dialog is dismissed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFUserNotificationCallBack
// CFUserNotificationCallBack has base type: void (*)(struct __CFUserNotification *, unsigned long)
type UserNotificationCallBack uintptr
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
// XMLNodeRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLNode
// CFXMLNodeRef has base type: const struct __CFXMLNode *
type XMLNodeRef uintptr
// XMLParserRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParser
// CFXMLParserRef has base type: struct __CFXMLParser *
type XMLParserRef uintptr
// XMLParserAddChildCallBack - Callback function invoked by the parser to notify your application of parent/child relationships between XML structures.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserAddChildCallBack
// CFXMLParserAddChildCallBack has base type: void (*)(struct __CFXMLParser *, void *, void *, void *)
type XMLParserAddChildCallBack uintptr
// XMLParserCopyDescriptionCallBack - Callback function invoked by the parser when handling the information pointer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCopyDescriptionCallBack
// CFXMLParserCopyDescriptionCallBack has base type: const struct __CFString *(*)(const void *)
type XMLParserCopyDescriptionCallBack uintptr
// XMLParserCreateXMLStructureCallBack - Callback function invoked when the parser encounters an XML open tag.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserCreateXMLStructureCallBack
// CFXMLParserCreateXMLStructureCallBack has base type: void *(*)(struct __CFXMLParser *, const struct __CFXMLNode *, void *)
type XMLParserCreateXMLStructureCallBack uintptr
// XMLParserEndXMLStructureCallBack - Callback function invoked by the parser to notify your application that an XML structure (and all its children) have been completely parsed.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserEndXMLStructureCallBack
// CFXMLParserEndXMLStructureCallBack has base type: void (*)(struct __CFXMLParser *, void *, void *)
type XMLParserEndXMLStructureCallBack uintptr
// XMLParserHandleErrorCallBack - Callback function invoked by the parser to notify your application that an error has occurred.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserHandleErrorCallBack
// CFXMLParserHandleErrorCallBack has base type: unsigned char (*)(struct __CFXMLParser *, enum CFXMLParserStatusCode, void *)
type XMLParserHandleErrorCallBack uintptr
// XMLParserReleaseCallBack - Callback function invoked by the parser when it wants to release a reference to the information pointer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserReleaseCallBack
// CFXMLParserReleaseCallBack has base type: void (*)(const void *)
type XMLParserReleaseCallBack uintptr
// XMLParserResolveExternalEntityCallBack - Callback function invoked by the parser to notify your application that an external entity has been referenced.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserResolveExternalEntityCallBack
// CFXMLParserResolveExternalEntityCallBack has base type: const struct __CFData *(*)(struct __CFXMLParser *, CFXMLExternalID *, void *)
type XMLParserResolveExternalEntityCallBack uintptr
// XMLParserRetainCallBack - Callback function invoked by the parser when it needs another reference to the information pointer.
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserRetainCallBack
// CFXMLParserRetainCallBack has base type: const void *(*)(const void *)
type XMLParserRetainCallBack uintptr
// XMLTreeRef type alias
//
// [Full Topic]: https://developer.apple.com/documentation/CoreFoundation/CFXMLTree
// CFXMLTreeRef has base type: CFTreeRef
type XMLTreeRef uintptr

