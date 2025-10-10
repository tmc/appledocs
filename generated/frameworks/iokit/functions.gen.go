// Code generated from Apple documentation for IOKit. DO NOT EDIT.

package iokit

// IOKit Functions
//
// This file contains function declarations discovered from Apple's documentation.
// To use these functions, you need to:
//   1. Map C types to Go types
//   2. Create function variables
//   3. Register them with purego.RegisterLibFunc
//
// Example:
//   var CGContextSetRGBFillColor func(c CGContextRef, red, green, blue, alpha CGFloat)
//   purego.RegisterLibFunc(&CGContextSetRGBFillColor, lib, "CGContextSetRGBFillColor")

// Discovered functions (119 total):

// IOCFSerialize(object _, options :  CFTypeRef!,  _, :  CFOptionFlags) ->  CFData!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IODestroyPlugInInterface(interface _, :  UnsafeMutablePointer< UnsafeMutablePointer< IOCFPlugInInterface>?>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+

// IOCreatePlugInInterfaceForService(service _, pluginType :  io_service_t,  _, interfaceType :  CFUUID!,  _, theInterface :  CFUUID!,  _, theScore :  UnsafeMutablePointer< UnsafeMutablePointer< UnsafeMutablePointer< IOCFPlugInInterface>?>?>!,  _, :  UnsafeMutablePointer< Int32>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.0+


// IOHIDManagerSetDeviceMatching(IOHIDManagerRef manager,  CFDictionaryRef matching);)
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.5+

// IOHIDManagerCopyDevices(IOHIDManagerRef manager);) CFSetRef
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.5+

// IORegistryEntryGetPath(entry _, plane :  io_registry_entry_t,  _, path :  UnsafePointer< CChar>!,  _, :  UnsafeMutablePointer< CChar>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOCatalogueGetData(mainPort _, flag :  mach_port_t,  _, buffer :  UInt32,  _, size :  UnsafeMutablePointer< UnsafeMutablePointer< CChar>?>!,  _, :  UnsafeMutablePointer< UInt32>!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IORegistryCreateIterator(mainPort _, plane :  mach_port_t,  _, options :  UnsafePointer< CChar>!,  _, iterator :  IOOptionBits,  _, :  UnsafeMutablePointer< io_iterator_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOConnectCallMethod(connection _, selector :  mach_port_t,  _, input :  UInt32,  _, inputCnt :  UnsafePointer< UInt64>!,  _, inputStruct :  UInt32,  _, inputStructCnt :  UnsafeRawPointer!,  _, output :  Int,  _, outputCnt :  UnsafeMutablePointer< UInt64>!,  _, outputStruct :  UnsafeMutablePointer< UInt32>!,  _, outputStructCnt :  UnsafeMutableRawPointer!,  _, :  UnsafeMutablePointer< Int>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.5+
//   - visionOS 1.0+


// IORegistryEntryCopyFromPath(mainPort _, path :  mach_port_t,  _, :  CFString!) ->  io_registry_entry_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - visionOS 1.0+

// IORegistryEntryGetProperty(entry _, propertyName :  io_registry_entry_t,  _, buffer :  UnsafePointer< CChar>!,  _, size :  UnsafeMutablePointer< CChar>!,  _, :  UnsafeMutablePointer< UInt32>!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// OSGetNotificationFromMessage(msg _, index :  UnsafeMutablePointer< mach_msg_header_t>!,  _, type :  UInt32,  _, reference :  UnsafeMutablePointer< UInt32>!,  _, content :  UnsafeMutablePointer< UInt>!,  _, size :  UnsafeMutablePointer< UnsafeMutableRawPointer?>!,  _, :  UnsafeMutablePointer< vm_size_t>!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+


// IOCFUnserialize(buffer _, allocator :  UnsafePointer< CChar>!,  _, options :  CFAllocator!,  _, errorString :  CFOptionFlags,  _, :  UnsafeMutablePointer< Unmanaged< CFString>?>!) ->  CFTypeRef!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOURLWriteDataAndPropertiesToResource(url _, dataToWrite :  CFURL!,  _, propertiesToWrite :  CFData!,  _, errorCode :  CFDictionary!,  _, :  UnsafeMutablePointer< Int32>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - Xcode 6.1+
//   - macOS 10.9+

// IOConnectCallStructMethod(connection _, selector :  mach_port_t,  _, inputStruct :  UInt32,  _, inputStructCnt :  UnsafeRawPointer!,  _, outputStruct :  Int,  _, outputStructCnt :  UnsafeMutableRawPointer!,  _, :  UnsafeMutablePointer< Int>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.5+
//   - visionOS 1.0+


// IODataQueueDequeue(dataQueue _, data :  UnsafeMutablePointer< IODataQueueMemory>!,  _, dataSize :  UnsafeMutableRawPointer!,  _, :  UnsafeMutablePointer< UInt32>!) ->  IOReturn) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IORegistryEntryCreateCFProperty(entry _, key :  io_registry_entry_t,  _, allocator :  CFString!,  _, options :  CFAllocator!,  _, :  IOOptionBits) ->  Unmanaged< CFTypeRef>!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IODataQueueSetNotificationPort(dataQueue _, notifyPort :  UnsafeMutablePointer< IODataQueueMemory>!,  _, :  mach_port_t) ->  IOReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.5+
//   - visionOS 1.0+


// IORegistryEntryCreateCFProperties(entry _, properties :  io_registry_entry_t,  _, allocator :  UnsafeMutablePointer< Unmanaged< CFMutableDictionary>?>!,  _, options :  CFAllocator!,  _, :  IOOptionBits) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IORegistryEntryCreateIterator(entry _, plane :  io_registry_entry_t,  _, options :  UnsafePointer< CChar>!,  _, iterator :  IOOptionBits,  _, :  UnsafeMutablePointer< io_iterator_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IORegistryEntryGetName(entry _, name :  io_registry_entry_t,  _, :  UnsafeMutablePointer< CChar>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOObjectGetKernelRetainCount(object _, :  io_object_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - visionOS 1.0+

// IORegistryIteratorExitEntry(iterator _, :  io_iterator_t) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IORegistryEntryGetLocationInPlane(entry _, plane :  io_registry_entry_t,  _, location :  UnsafePointer< CChar>!,  _, :  UnsafeMutablePointer< CChar>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.1+
//   - visionOS 2.4+


// IOConnectTrap5(connect _, index :  io_connect_t,  _, p1 :  UInt32,  _, p2 :  UInt,  _, p3 :  UInt,  _, p4 :  UInt,  _, p5 :  UInt,  _, :  UInt) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IOConnectTrap2(connect _, index :  io_connect_t,  _, p1 :  UInt32,  _, p2 :  UInt,  _, :  UInt) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IOServiceAddMatchingNotification(notifyPort _, notificationType :  IONotificationPortRef!,  _, matching :  UnsafePointer< CChar>!,  _, callback :  CFDictionary!,  _, refCon :  IOServiceMatchingCallback!,  _, notification :  UnsafeMutableRawPointer!,  _, :  UnsafeMutablePointer< io_iterator_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOServiceRequestProbe(service _, options :  io_service_t,  _, :  UInt32) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IORegistryEntryGetParentIterator(entry _, plane :  io_registry_entry_t,  _, iterator :  UnsafePointer< CChar>!,  _, :  UnsafeMutablePointer< io_iterator_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOObjectCopyBundleIdentifierForClass(classname _, :  CFString!) ->  Unmanaged< CFString>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - visionOS 1.0+


// IOConnectMapMemory(connect _, memoryType :  io_connect_t,  _, intoTask :  UInt32,  _, atAddress :  task_port_t,  _, ofSize :  UnsafeMutablePointer< mach_vm_address_t>!,  _, options :  UnsafeMutablePointer< mach_vm_size_t>!,  _, :  IOOptionBits) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOIteratorReset(iterator _, :  io_iterator_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOServiceAddNotification(mainPort _, notificationType :  mach_port_t,  _, matching :  UnsafePointer< CChar>!,  _, wakePort :  CFDictionary!,  _, reference :  mach_port_t,  _, notification :  UInt,  _, :  UnsafeMutablePointer< io_iterator_t>!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.6)
//
// Deprecated: This function is deprecated.


// IODataQueueDataAvailable(dataQueue _, :  UnsafeMutablePointer< IODataQueueMemory>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOConnectCallAsyncStructMethod(connection _, selector :  mach_port_t,  _, wake_port :  UInt32,  _, reference :  mach_port_t,  _, referenceCnt :  UnsafeMutablePointer< UInt64>!,  _, inputStruct :  UInt32,  _, inputStructCnt :  UnsafeRawPointer!,  _, outputStruct :  Int,  _, outputStructCnt :  UnsafeMutableRawPointer!,  _, :  UnsafeMutablePointer< Int>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.5+
//   - visionOS 1.0+

// IOCatalogueSendData(mainPort _, flag :  mach_port_t,  _, buffer :  UInt32,  _, size :  UnsafePointer< CChar>!,  _, :  UInt32) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+


// IOConnectTrap4(connect _, index :  io_connect_t,  _, p1 :  UInt32,  _, p2 :  UInt,  _, p3 :  UInt,  _, p4 :  UInt,  _, :  UInt) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IORegistryEntrySetCFProperties(entry _, properties :  io_registry_entry_t,  _, :  CFTypeRef!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IOServiceNameMatching(name _, :  UnsafePointer< CChar>!) ->  CFMutableDictionary!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOConnectCallAsyncMethod(connection _, selector :  mach_port_t,  _, wake_port :  UInt32,  _, reference :  mach_port_t,  _, referenceCnt :  UnsafeMutablePointer< UInt64>!,  _, input :  UInt32,  _, inputCnt :  UnsafePointer< UInt64>!,  _, inputStruct :  UInt32,  _, inputStructCnt :  UnsafeRawPointer!,  _, output :  Int,  _, outputCnt :  UnsafeMutablePointer< UInt64>!,  _, outputStruct :  UnsafeMutablePointer< UInt32>!,  _, outputStructCnt :  UnsafeMutableRawPointer!,  _, :  UnsafeMutablePointer< Int>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.5+
//   - visionOS 1.0+

// IOConnectGetService(connect _, service :  io_connect_t,  _, :  UnsafeMutablePointer< io_service_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOKitWaitQuiet(mainPort _, waitTime :  mach_port_t,  _, :  UnsafeMutablePointer< mach_timespec_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IORegistryEntryGetParentEntry(entry _, plane :  io_registry_entry_t,  _, parent :  UnsafePointer< CChar>!,  _, :  UnsafeMutablePointer< io_registry_entry_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOKitGetBusyState(mainPort _, busyState :  mach_port_t,  _, :  UnsafeMutablePointer< UInt32>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOObjectGetUserRetainCount(object _, :  io_object_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.6+
//   - visionOS 1.0+


// IORegistryEntryGetNameInPlane(entry _, plane :  io_registry_entry_t,  _, name :  UnsafePointer< CChar>!,  _, :  UnsafeMutablePointer< CChar>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IONotificationPortCreate(mainPort _, :  mach_port_t) ->  IONotificationPortRef!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IODataQueueEnqueue(dataQueue _, data :  UnsafeMutablePointer< IODataQueueMemory>!,  _, dataSize :  UnsafeMutableRawPointer!,  _, :  UInt32) ->  IOReturn) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.5+
//   - visionOS 1.0+


// IOBSDNameMatching(mainPort _, options :  mach_port_t,  _, bsdName :  UInt32,  _, :  UnsafePointer< CChar>!) ->  CFMutableDictionary!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOConnectTrap6(connect _, index :  io_connect_t,  _, p1 :  UInt32,  _, p2 :  UInt,  _, p3 :  UInt,  _, p4 :  UInt,  _, p5 :  UInt,  _, p6 :  UInt,  _, :  UInt) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IOServiceGetMatchingServices(mainPort _, matching :  mach_port_t,  _, existing :  CFDictionary!,  _, :  UnsafeMutablePointer< io_iterator_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IODataQueueAllocateNotificationPort() func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IORegistryEntryGetChildEntry(entry _, plane :  io_registry_entry_t,  _, child :  UnsafePointer< CChar>!,  _, :  UnsafeMutablePointer< io_registry_entry_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOURLCreatePropertyFromResource(alloc _, url :  CFAllocator!,  _, property :  CFURL!,  _, errorCode :  CFString!,  _, :  UnsafeMutablePointer< Int32>!) ->  Unmanaged< CFTypeRef>!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - Xcode 6.1+
//   - macOS 10.9+


// IOObjectConformsTo(object _, className :  io_object_t,  _, :  UnsafePointer< CChar>!) ->  boolean_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOConnectRelease(connect _, :  io_connect_t) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOServiceOpen(service _, owningTask :  io_service_t,  _, type :  task_port_t,  _, connect :  UInt32,  _, :  UnsafeMutablePointer< io_connect_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOConnectUnmapMemory(connect _, memoryType :  io_connect_t,  _, fromTask :  UInt32,  _, atAddress :  task_port_t,  _, :  mach_vm_address_t) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOServiceAuthorize(service _, options :  io_service_t,  _, :  UInt32) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IOServiceGetMatchingService(mainPort _, matching :  mach_port_t,  _, :  CFDictionary!) ->  io_service_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.2+
//   - visionOS 2.4+


// IORegistryEntrySearchCFProperty(entry _, plane :  io_registry_entry_t,  _, key :  UnsafePointer< CChar>!,  _, allocator :  CFString!,  _, options :  CFAllocator!,  _, :  IOOptionBits) ->  CFTypeRef!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.1+
//   - visionOS 2.4+

// IOConnectSetNotificationPort(connect _, type :  io_connect_t,  _, port :  UInt32,  _, reference :  mach_port_t,  _, :  UInt) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOIteratorIsValid(iterator _, :  io_iterator_t) ->  boolean_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOObjectIsEqualTo(object _, anObject :  io_object_t,  _, :  io_object_t) ->  boolean_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOServiceWaitQuiet(service _, waitTime :  io_service_t,  _, :  UnsafeMutablePointer< mach_timespec_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IONotificationPortSetDispatchQueue(notify _, queue :  IONotificationPortRef!,  _, :  dispatch_queue_t!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.3+
//   - iPadOS 4.3+
//   - macOS 10.6+
//   - visionOS 1.0+


// IONotificationPortGetRunLoopSource(notify _, :  IONotificationPortRef!) ->  Unmanaged< CFRunLoopSource>!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOServiceGetBusyState(service _, busyState :  io_service_t,  _, :  UnsafeMutablePointer< UInt32>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOConnectAddClient(connect _, client :  io_connect_t,  _, :  io_connect_t) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+


// IOObjectRelease(object _, :  io_object_t) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOObjectCopySuperclassForClass(classname _, :  CFString!) ->  Unmanaged< CFString>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - visionOS 1.0+

// IOServiceClose(connect _, :  io_connect_t) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IODataQueuePeek(dataQueue _, :  UnsafeMutablePointer< IODataQueueMemory>!) ->  UnsafeMutablePointer< IODataQueueEntry>!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOMasterPort(bootstrapPort _, mainPort :  mach_port_t,  _, :  UnsafeMutablePointer< mach_port_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 10.14+
//   - macOS 10.0+ (Deprecated in 12.0)
//
// Deprecated: This function is deprecated.

// IOServiceOFPathToBSDName(mainPort _, openFirmwarePath :  mach_port_t,  _, bsdName :  UnsafePointer< CChar>!,  _, :  UnsafeMutablePointer< CChar>!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.


// IOCatalogueTerminate(mainPort _, flag :  mach_port_t,  _, description :  UInt32,  _, :  UnsafeMutablePointer< CChar>!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IORegistryEntryInPlane(entry _, plane :  io_registry_entry_t,  _, :  UnsafePointer< CChar>!) ->  boolean_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOConnectTrap0(connect _, index :  io_connect_t,  _, :  UInt32) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+


// IOServiceMatchPropertyTable(service _, matching :  io_service_t,  _, matches :  CFDictionary!,  _, :  UnsafeMutablePointer< boolean_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOServiceMatching(name _, :  UnsafePointer< CChar>!) ->  CFMutableDictionary!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IODataQueueWaitForAvailableData(dataQueue _, notificationPort :  UnsafeMutablePointer< IODataQueueMemory>!,  _, :  mach_port_t) ->  IOReturn) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOCreateReceivePort(msgType _, recvPort :  UInt32,  _, :  UnsafeMutablePointer< mach_port_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOCatalogueReset(mainPort _, flag :  mach_port_t,  _, :  UInt32) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IORegistryEntryGetChildIterator(entry _, plane :  io_registry_entry_t,  _, iterator :  UnsafePointer< CChar>!,  _, :  UnsafeMutablePointer< io_iterator_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOConnectSetCFProperties(connect _, properties :  io_connect_t,  _, :  CFTypeRef!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IOOpenFirmwarePathMatching(mainPort _, options :  mach_port_t,  _, path :  UInt32,  _, :  UnsafePointer< CChar>!) ->  Unmanaged< CFMutableDictionary>!) func
//
// Availability:
//   - macOS 10.0+ (Deprecated in 10.8)
//
// Deprecated: This function is deprecated.

// IORegistryEntryGetRegistryEntryID(entry _, entryID :  io_registry_entry_t,  _, :  UnsafeMutablePointer< UInt64>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.6+
//   - visionOS 2.4+


// IOConnectAddRef(connect _, :  io_connect_t) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOIteratorNext(iterator _, :  io_iterator_t) ->  io_object_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOCFUnserializeWithSize(buffer _, bufferSize :  UnsafePointer< CChar>!,  _, allocator :  Int,  _, options :  CFAllocator!,  _, errorString :  CFOptionFlags,  _, :  UnsafeMutablePointer< Unmanaged< CFString>?>!) ->  CFTypeRef!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.10+
//   - visionOS 2.4+


// IONotificationPortDestroy(notify _, :  IONotificationPortRef!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOObjectGetClass(object _, className :  io_object_t,  _, :  UnsafeMutablePointer< CChar>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOConnectUnmapMemory64(connect _, memoryType :  io_connect_t,  _, fromTask :  UInt32,  _, atAddress :  task_port_t,  _, :  mach_vm_address_t) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.5+
//   - visionOS 2.4+


// IOObjectRetain(object _, :  io_object_t) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.1+
//   - visionOS 2.4+

// IODispatchCalloutFromMessage(unused _, msg :  UnsafeMutableRawPointer!,  _, reference :  UnsafeMutablePointer< mach_msg_header_t>!,  _, :  UnsafeMutableRawPointer!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOObjectCopyClass(object _, :  io_object_t) ->  Unmanaged< CFString>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.4+
//   - visionOS 1.0+


// IOConnectCallScalarMethod(connection _, selector :  mach_port_t,  _, input :  UInt32,  _, inputCnt :  UnsafePointer< UInt64>!,  _, output :  UInt32,  _, outputCnt :  UnsafeMutablePointer< UInt64>!,  _, :  UnsafeMutablePointer< UInt32>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.5+
//   - visionOS 1.0+

// IOConnectSetCFProperty(connect _, propertyName :  io_connect_t,  _, property :  CFString!,  _, :  CFTypeRef!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IORegistryEntryFromPath(mainPort _, path :  mach_port_t,  _, :  UnsafePointer< CChar>!) ->  io_registry_entry_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOConnectTrap1(connect _, index :  io_connect_t,  _, p1 :  UInt32,  _, :  UInt) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IORegistryIteratorEnterEntry(iterator _, :  io_iterator_t) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOObjectGetRetainCount(object _, :  io_object_t) ->  UInt32) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOConnectTrap3(connect _, index :  io_connect_t,  _, p1 :  UInt32,  _, p2 :  UInt,  _, p3 :  UInt,  _, :  UInt) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IOURLCreateDataAndPropertiesFromResource(alloc _, url :  CFAllocator!,  _, resourceData :  CFURL!,  _, properties :  UnsafeMutablePointer< Unmanaged< CFData>?>!,  _, desiredProperties :  UnsafeMutablePointer< Unmanaged< CFDictionary>?>!,  _, errorCode :  CFArray!,  _, :  UnsafeMutablePointer< Int32>!) ->  Bool) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - Xcode 6.1+
//   - macOS 10.9+

// IORegistryEntryCopyPath(entry _, plane :  io_registry_entry_t,  _, :  UnsafePointer< CChar>!) ->  Unmanaged< CFString>!) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - visionOS 1.0+


// IOConnectMapMemory64(connect _, memoryType :  io_connect_t,  _, intoTask :  UInt32,  _, atAddress :  task_port_t,  _, ofSize :  UnsafeMutablePointer< mach_vm_address_t>!,  _, options :  UnsafeMutablePointer< mach_vm_size_t>!,  _, :  IOOptionBits) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.5+
//   - visionOS 2.4+

// IOServiceAddInterestNotification(notifyPort _, service :  IONotificationPortRef!,  _, interestType :  io_service_t,  _, callback :  UnsafePointer< CChar>!,  _, refCon :  IOServiceInterestCallback!,  _, notification :  UnsafeMutableRawPointer!,  _, :  UnsafeMutablePointer< io_object_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IONotificationPortGetMachPort(notify _, :  IONotificationPortRef!) ->  mach_port_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+


// IOCFUnserializeBinary(buffer _, bufferSize :  UnsafePointer< CChar>!,  _, allocator :  Int,  _, options :  CFAllocator!,  _, errorString :  CFOptionFlags,  _, :  UnsafeMutablePointer< Unmanaged< CFString>?>!) ->  CFTypeRef!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.10+
//   - visionOS 2.4+

// IORegistryGetRootEntry(mainPort _, :  mach_port_t) ->  io_registry_entry_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.0+
//   - visionOS 2.4+

// IOServiceOpenAsFileDescriptor(service _, oflag :  io_service_t,  _, :  Int32) ->  Int32) func
//
// Availability:
//   - macOS 10.0+


// IORegistryEntryIDMatching(entryID _, :  UInt64) ->  CFMutableDictionary!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.6+
//   - visionOS 2.4+

// IORegistryEntrySetCFProperty(entry _, propertyName :  io_registry_entry_t,  _, property :  CFString!,  _, :  CFTypeRef!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IOConnectCallAsyncScalarMethod(connection _, selector :  mach_port_t,  _, wake_port :  UInt32,  _, reference :  mach_port_t,  _, referenceCnt :  UnsafeMutablePointer< UInt64>!,  _, input :  UInt32,  _, inputCnt :  UnsafePointer< UInt64>!,  _, output :  UInt32,  _, outputCnt :  UnsafeMutablePointer< UInt64>!,  _, :  UnsafeMutablePointer< UInt32>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - macOS 10.5+
//   - visionOS 1.0+


// IOCatalogueModuleLoaded(mainPort _, name :  mach_port_t,  _, :  UnsafeMutablePointer< CChar>!) ->  kern_return_t) func
//
// Availability:
//   - macOS 10.0+

// IOHIDDeviceOpen(IOHIDDeviceRef device,  IOOptionBits options);) IOReturn
//
// Availability:
//   - Mac Catalyst 13.1+
//   - macOS 10.5+

// IONotificationPortSetImportanceReceiver(notify _, :  IONotificationPortRef!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 18.4+
//   - iPadOS 18.4+
//   - macOS 10.13+
//   - visionOS 2.4+


// IORPCMessageFromMach(msg _, reply :  UnsafeMutablePointer< IORPCMessageMach>!,  _, :  Bool) ->  UnsafeMutablePointer< IORPCMessage>!) func
//
// Availability:
//   - Mac Catalyst 13.0+
//   - macOS 10.15+

// IOMainPort(bootstrapPort _, mainPort :  mach_port_t,  _, :  UnsafeMutablePointer< mach_port_t>!) ->  kern_return_t) func
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - macOS 12.0+
//   - visionOS 1.0+

