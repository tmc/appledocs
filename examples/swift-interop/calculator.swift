// calculator.swift - Example of Swift class bridged via Objective-C for Go interop

import Foundation

// Swift class that can be exposed to Objective-C runtime
@objc public class Calculator: NSObject {
    private var memory: Double = 0

    @objc public func add(_ a: Double, _ b: Double) -> Double {
        return a + b
    }

    @objc public func multiply(_ a: Double, _ b: Double) -> Double {
        return a * b
    }

    @objc public func store(_ value: Double) {
        memory = value
    }

    @objc public func recall() -> Double {
        return memory
    }
}

// C-compatible wrapper functions for purego
@_cdecl("calculator_create")
public func calculatorCreate() -> UnsafeMutableRawPointer {
    let calc = Calculator()
    return Unmanaged.passRetained(calc).toOpaque()
}

@_cdecl("calculator_add")
public func calculatorAdd(_ ptr: UnsafeMutableRawPointer, _ a: Double, _ b: Double) -> Double {
    let calc = Unmanaged<Calculator>.fromOpaque(ptr).takeUnretainedValue()
    return calc.add(a, b)
}

@_cdecl("calculator_multiply")
public func calculatorMultiply(_ ptr: UnsafeMutableRawPointer, _ a: Double, _ b: Double) -> Double {
    let calc = Unmanaged<Calculator>.fromOpaque(ptr).takeUnretainedValue()
    return calc.multiply(a, b)
}

@_cdecl("calculator_store")
public func calculatorStore(_ ptr: UnsafeMutableRawPointer, _ value: Double) {
    let calc = Unmanaged<Calculator>.fromOpaque(ptr).takeUnretainedValue()
    calc.store(value)
}

@_cdecl("calculator_recall")
public func calculatorRecall(_ ptr: UnsafeMutableRawPointer) -> Double {
    let calc = Unmanaged<Calculator>.fromOpaque(ptr).takeUnretainedValue()
    return calc.recall()
}

@_cdecl("calculator_release")
public func calculatorRelease(_ ptr: UnsafeMutableRawPointer) {
    Unmanaged<Calculator>.fromOpaque(ptr).release()
}
