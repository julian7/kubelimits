# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

No changes so far.

## [v0.3.0] - Jun 1, 2026

Fixed:

* Do not error on missing cpu/memory max cgroup2 file


## [v0.2.0] - Aug 14, 2025

Added:

* GOMAXPROCS support in go v1.25

## [v0.1.0] - Mar 5, 2025

Added:

* GOMAXPROCS setter based on cgroups2's cpu.max
* memory limit setter based on cgroups2's memory.max
* MIT license

[Unreleased]: https://github.com/julian7/kubelimits
[v0.1.0]: https://github.com/julian7/kubelimits/releases/tag/v0.1.0
[v0.2.0]: https://github.com/julian7/kubelimits/releases/tag/v0.2.0
[v0.3.0]: https://github.com/julian7/kubelimits/releases/tag/v0.3.0
