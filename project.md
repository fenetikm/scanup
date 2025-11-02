# Project - ScanUp!
## Overview
CLI tool. Scans websites for vulns.

Ideas:
- Runs as a service (send it sites to scan) or one shot
- Spits out a report or stores it
- Different scanners e.g. regex pattern, AI, 3rd party service etc.
- Different sources e.g. file, URL etc.
- See changes to a website over time, run on a schedule

## Architecture
Go - queue processor running concurrently.
Bubbletea? for TUI, to come later.
Storage, file or sqlite should be fine.

## Config
Ingestion: file, web.
Cli, try out bubbletea as a tui.
Service that ingests sites, scans them, put results in db.
view results.
Scanners: pattern, ai.
Scan upguard website?
Find an out of date package or vuln? Obfuscated using unicode whitespace?
Or .env file in the open.
Service
- Add site
- List sites

Config file:
- retry
- how often
- User agent

## Steps
### To-do

### Done

