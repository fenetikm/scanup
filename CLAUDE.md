# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

ScanUp is a CLI tool for scanning websites for vulnerabilities.

See [the project file](./project.md) file for further information

## Planned Architecture
See [the plan file](./plan.md) file for an outline of how the development will proceed and what to do next.

## Development Commands

No build system is currently set up. Once Go modules are initialized, typical commands will be:
- `go run .` - Run the application
- `go build` - Build the binary
- `go test ./...` - Run tests
- `go mod tidy` - Clean up dependencies

## Security Focus

This tool is designed for **defensive security purposes only**:
- Vulnerability detection and reporting
- Website security assessment
- Compliance monitoring
- Change detection for security monitoring

The tool should never be used for malicious purposes or unauthorized scanning.
