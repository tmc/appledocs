# DarwinKit-Style Bindings Documentation

Comprehensive documentation for the DarwinKit-style Go bindings to Apple frameworks.

## Documentation Index

### Getting Started

1. **[MIGRATION_GUIDE.md](./MIGRATION_GUIDE.md)** - Start here if you're migrating from function-style bindings
   - Key differences between styles
   - Step-by-step migration examples
   - Memory management changes
   - Common patterns and troubleshooting

### Understanding the Design

2. **[API_DESIGN_RATIONALE.md](./API_DESIGN_RATIONALE.md)** - Understand why the API is designed this way
   - Design goals and principles
   - Architectural decisions
   - Type system choices
   - Performance considerations
   - Trade-offs and alternatives

### Writing Good Code

3. **[BEST_PRACTICES.md](./BEST_PRACTICES.md)** - Learn how to use the bindings effectively
   - Memory management patterns
   - Type safety guidelines
   - Concurrency best practices
   - Performance optimization
   - Testing strategies
   - Common pitfalls and solutions

## Quick Links

### For New Users
→ Start with [MIGRATION_GUIDE.md](./MIGRATION_GUIDE.md) to understand the basics

### For Existing Users
→ Check [BEST_PRACTICES.md](./BEST_PRACTICES.md) for patterns and anti-patterns

### For Contributors
→ Read [API_DESIGN_RATIONALE.md](./API_DESIGN_RATIONALE.md) to understand design decisions

## Example Code

See the `/examples/` directory for working examples:
- `purego-objc-demo/` - Direct purego/objc usage
- `darwinkit-bindings-test/` - DarwinKit-style bindings example
- More examples coming soon...

## Additional Resources

### External Documentation
- [Apple Objective-C Documentation](https://developer.apple.com/documentation/objectivec)
- [progrium/darwinkit](https://github.com/progrium/darwinkit) - Reference implementation
- [ebitengine/purego](https://github.com/ebitengine/purego) - Underlying FFI layer
- [Effective Go](https://go.dev/doc/effective_go) - Go programming guide

### Technical Documentation
- [DARWINKIT_ANALYSIS.md](../cmd/generate-framework-bindings/DARWINKIT_ANALYSIS.md) - Detailed analysis of DarwinKit patterns
- [TEMPLATE_DESIGN.md](../cmd/generate-framework-bindings/TEMPLATE_DESIGN.md) - Template architecture
- [TESTING.md](../cmd/generate-framework-bindings/TESTING.md) - Testing strategy

## Documentation Philosophy

Our documentation follows these principles:

1. **Examples First** - Show working code before explaining theory
2. **Practical Focus** - Real-world patterns over academic completeness
3. **Progressive Disclosure** - Start simple, add complexity as needed
4. **Visual Clarity** - Use tables, code blocks, and formatting
5. **Linked Context** - Cross-reference related topics

## Contributing to Documentation

Improvements welcome! When contributing:

1. **Keep it Practical** - Include runnable code examples
2. **Explain Why** - Don't just show how, explain the rationale
3. **Update Examples** - Keep code examples tested and current
4. **Cross-Link** - Reference related sections and external docs
5. **Test Locally** - Verify all code examples compile and run

## Getting Help

### If you have a question:
1. Check these docs first
2. Look at `/examples/` for working code
3. Search [GitHub Issues](https://github.com/tmc/appledocs/issues)
4. Ask in the community

### If you found a bug:
1. Check if it's already reported
2. Create a minimal reproduction
3. Open a new issue with details

### If you want to contribute:
1. Read the relevant design docs
2. Discuss in an issue first (for large changes)
3. Submit a PR with tests and documentation

## Documentation Status

| Document | Status | Last Updated |
|----------|--------|--------------|
| MIGRATION_GUIDE.md | ✅ Complete | 2025-10-16 |
| API_DESIGN_RATIONALE.md | ✅ Complete | 2025-10-16 |
| BEST_PRACTICES.md | ✅ Complete | 2025-10-16 |
| README.md (this file) | ✅ Complete | 2025-10-16 |

## Roadmap

Planned documentation additions:

- [ ] Tutorial: Building Your First macOS App
- [ ] Guide: Working with Delegates and Protocols
- [ ] Guide: Memory Debugging Techniques
- [ ] Reference: Complete Type Mapping Table
- [ ] Cookbook: Common UI Patterns
- [ ] Troubleshooting: Known Issues and Workarounds

## License

Documentation is part of the appledocs project. See parent repository for license details.

---

**Navigate:** [Migration Guide](./MIGRATION_GUIDE.md) | [API Design](./API_DESIGN_RATIONALE.md) | [Best Practices](./BEST_PRACTICES.md)
