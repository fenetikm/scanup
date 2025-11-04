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
Language: Go - queue processor running concurrently.
Storage: file, sqlite, redis backends.
Testing: TBD.

Idea:
- Have different aspects to scan for such as:
    * domain (check ownership, expiry, a records)
    * csp, headers
    * libraries / frameworks (up to date?)
    * uptime
    * existence of bad files
- Pull everything into key/value pairs
- For each of those, have analysers, simplest being regex but could have AI
- Reports are created, like NeoVim healthcheck
- Compile all into a full report
- Report to DB, file, stdout
- Queues for all the things (nice to have)
- Different timing schedules, essentially cron
- Retry policies

### Other
Bubbletea? if we make a TUI

## Config
How to store sites?
