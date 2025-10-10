// hello.swift - Simple Swift library with C-compatible exports

import Foundation

// Export a simple function with C calling convention
@_cdecl("swift_hello")
public func swiftHello() {
    print("Hello from Swift!")
}

// Export a function that returns a value
@_cdecl("swift_add")
public func swiftAdd(_ a: Int32, _ b: Int32) -> Int32 {
    return a + b
}

// Export a function that works with C strings
@_cdecl("swift_greet")
public func swiftGreet(_ name: UnsafePointer<CChar>) {
    let swiftName = String(cString: name)
    print("Hello, \(swiftName), from Swift!")
}

// Export a function that returns a C string (caller must free)
@_cdecl("swift_get_message")
public func swiftGetMessage() -> UnsafeMutablePointer<CChar> {
    let message = "Message from Swift"
    let cString = strdup(message)
    return cString!
}
