import Foundation
import SwiftSyntax
import SwiftParser

// MARK: - API Models

struct APIDefinition: Codable {
    let framework: String
    let version: String
    let extensions: [ExtensionDef]
}

struct ExtensionDef: Codable {
    let extendedType: String
    let availability: [String]
    let methods: [MethodDef]
    let properties: [PropertyDef]
}

struct MethodDef: Codable {
    let name: String
    let availability: [String]
    let parameters: [ParameterDef]
    let returnType: String
    let isAsync: Bool
    let canThrow: Bool
    let isGeneric: Bool
    let genericConstraints: [String]
}

struct ParameterDef: Codable {
    let label: String?
    let name: String
    let type: String
}

struct PropertyDef: Codable {
    let name: String
    let type: String
    let availability: [String]
    let getter: Bool
    let setter: Bool
}

// MARK: - Visitor

class InterfaceVisitor: SyntaxVisitor {
    var extensions: [ExtensionDef] = []

    override func visit(_ node: ExtensionDeclSyntax) -> SyntaxVisitorContinueKind {
        // Extract extended type name
        let extendedType = node.extendedType.trimmedDescription

        // Extract availability attributes
        let availability = extractAvailability(from: node.attributes)

        // Extract methods and properties
        var methods: [MethodDef] = []
        var properties: [PropertyDef] = []

        for member in node.memberBlock.members {
            if let function = member.decl.as(FunctionDeclSyntax.self) {
                methods.append(extractMethod(function))
            } else if let property = member.decl.as(VariableDeclSyntax.self) {
                properties.append(contentsOf: extractProperties(property))
            }
        }

        let ext = ExtensionDef(
            extendedType: extendedType,
            availability: availability,
            methods: methods,
            properties: properties
        )

        extensions.append(ext)

        return .skipChildren
    }

    private func extractMethod(_ function: FunctionDeclSyntax) -> MethodDef {
        let name = function.name.text
        let availability = extractAvailability(from: function.attributes)

        // Extract parameters
        let parameters = function.signature.parameterClause.parameters.map { param in
            ParameterDef(
                label: param.firstName.text == "_" ? nil : param.firstName.text,
                name: param.secondName?.text ?? param.firstName.text,
                type: param.type.trimmedDescription
            )
        }

        // Extract return type
        let returnType = function.signature.returnClause?.type.trimmedDescription ?? "Void"

        // Check for async/throws
        let isAsync = function.signature.effectSpecifiers?.asyncSpecifier != nil
        let throwsEffect = function.signature.effectSpecifiers?.throwsSpecifier != nil

        // Check for generics
        let isGeneric = function.genericParameterClause != nil
        let genericConstraints = extractGenericConstraints(function.genericWhereClause)

        return MethodDef(
            name: name,
            availability: availability,
            parameters: parameters,
            returnType: returnType,
            isAsync: isAsync,
            canThrow: throwsEffect,
            isGeneric: isGeneric,
            genericConstraints: genericConstraints
        )
    }

    private func extractProperties(_ variable: VariableDeclSyntax) -> [PropertyDef] {
        let availability = extractAvailability(from: variable.attributes)

        return variable.bindings.compactMap { binding in
            guard let identifier = binding.pattern.as(IdentifierPatternSyntax.self) else {
                return nil
            }

            let name = identifier.identifier.text
            let type = binding.typeAnnotation?.type.trimmedDescription ?? "Unknown"

            // Check if has getter/setter
            var hasGetter = false
            var hasSetter = false

            if let accessor = binding.accessorBlock {
                switch accessor.accessors {
                case .accessors(let list):
                    for acc in list {
                        if acc.accessorSpecifier.text == "get" {
                            hasGetter = true
                        } else if acc.accessorSpecifier.text == "set" {
                            hasSetter = true
                        }
                    }
                case .getter:
                    hasGetter = true
                }
            }

            return PropertyDef(
                name: name,
                type: type,
                availability: availability,
                getter: hasGetter,
                setter: hasSetter
            )
        }
    }

    private func extractAvailability(from attributes: AttributeListSyntax?) -> [String] {
        guard let attributes = attributes else { return [] }

        var availability: [String] = []

        for attribute in attributes {
            if let attr = attribute.as(AttributeSyntax.self),
               attr.attributeName.trimmedDescription == "available" {
                availability.append(attr.trimmedDescription)
            }
        }

        return availability
    }

    private func extractGenericConstraints(_ whereClause: GenericWhereClauseSyntax?) -> [String] {
        guard let whereClause = whereClause else { return [] }

        return whereClause.requirements.map { requirement in
            requirement.trimmedDescription
        }
    }
}

// MARK: - Main

func main() {
    let args = CommandLine.arguments

    guard args.count >= 2 else {
        print("Usage: swift-interface-parser <path-to-swiftinterface>", to: &standardError)
        exit(1)
    }

    let inputPath = args[1]

    do {
        // Read the .swiftinterface file
        let source = try String(contentsOfFile: inputPath, encoding: .utf8)

        // Parse with SwiftSyntax
        let tree = Parser.parse(source: source)

        // Visit and extract APIs
        let visitor = InterfaceVisitor(viewMode: .sourceAccurate)
        visitor.walk(tree)

        // Extract framework name and version from header comments
        let lines = source.split(separator: "\n", maxSplits: 10)
        var frameworkName = "Unknown"
        var version = "Unknown"

        for line in lines {
            if line.contains("module-name") {
                if let name = line.split(separator: " ").last {
                    frameworkName = String(name)
                }
            }
            if line.contains("user-module-version") {
                if let ver = line.split(separator: " ").last {
                    version = String(ver)
                }
            }
        }

        // Create API definition
        let apiDef = APIDefinition(
            framework: frameworkName,
            version: version,
            extensions: visitor.extensions
        )

        // Output as JSON
        let encoder = JSONEncoder()
        encoder.outputFormatting = [.prettyPrinted, .sortedKeys]
        let jsonData = try encoder.encode(apiDef)

        if let jsonString = String(data: jsonData, encoding: .utf8) {
            print(jsonString)
        }

    } catch {
        print("Error: \(error)", to: &standardError)
        exit(1)
    }
}

// Helper to print to stderr
struct StdErr: TextOutputStream {
    mutating func write(_ string: String) {
        guard let data = string.data(using: .utf8) else { return }
        FileHandle.standardError.write(data)
    }
}

var standardError = StdErr()

main()
