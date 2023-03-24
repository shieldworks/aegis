---
#
# .-'_.---._'-.
# ||####|(__)||   Protect your secrets, protect your business.
#   \\()|##//       Secure your sensitive data with Aegis.
#    \\ |#//                    <aegis.ist>
#     .\_/.
#

layout: default
keywords: Aegis, Changelog
title: Aegis Changelog
description: what happened so far
micro_nav: true
page_nav:
  prev:
    content: Miscellaneous
    url: '/misc'
---

# Aegis Changelog

## Currently Working On

## Changed

* Upgraded Aegis to use `SPIRE v1.6.1`.

## [v0.13.0] - 2023-03-03

### Added

* Aegis has a new website: [aegis.ist](https://aegis.ist/).
* Added a documentation page for Aegis Sentinel CLI usage.
* Template transformations now apply to the secret values that are
  reflected to the workloads as well.

### Fixed

* Fixed minor errors in documentation.
* Minor bug fixes.

### Changed

* **BREAKING**: Trust root of entities changed from `z2h.dev` to `aegist.ist`.
* Updated to Aegis Sentinel commands.
* Documentation updates to talk about how to use Aegis Sidecar to propagate 
  Aegis-Safe-managed secrets into Kubernetes.
* Documentation update: Changed contact, support, and security feedback emails
  to official `@aegis.ist` emails.

## [v0.12.70] - 2023-02-17

### Added

* Ability to wait for the secret to initialize first, using Aegis Init Container.
* Ability to use go templates to transform secrets. In this version, the 
  transformation only applies to the generated Kubernetes `Secret`s. In the
  upcoming versions, Aegis Safe API will also honor the transformations when
  returning the secret values.

### Changed

* Upgraded all builder images to Go `v1.20.1`.

### Fixed

* Fixed a channel overflow bug that was blocking secret operations when an
  error occurs.

## [v0.12.55] - 2023-02-15

### Added

* Added ability to use actual Kubernetes `Secret` object to save the values
  of the registered secrets—this is (*mostly*) for legacy support.
* Added an option to use the cluster as a backing store (*work in progress*)
* Implemented additional configuration options through environment variables
* Created a script to list the status of all Aegis projects (*especially
  useful when doing deployments and release cuts*)

### Changed

* There are breaking changes in certain environment variable names. The 
  documentation ha been updated to reflect these changes.

## [v0.12.30] - 2023-02-05

### Added

* Added contributing guidelines.
* Added local development instructions.
* Added default values to sample `yaml` manifests.
* Added the ability to list secrets to Aegis Sentinel.
* Other documentation updates.

### Changed

* Improvements to local development workflow.
* **BREAKING**: Changes in Aegis Safe REST API that required changes in
  demo workload implementations, and Aegis Sentinel.

## [v0.11.20] - 2023-02-03

### Added

* Added a [media kit ](/media).
* Added more configuration options through environment variables.

### Changed

* Aegis Sidecar now exponentially backs off if it cannot fetch secrets
  in a timely manner.
* **BREAKING**: All time units in environment variables 
  are now milliseconds, instead of seconds.

## [v0.11.5] - 2023-02-01

### Added

* Improved the website’s information architecture.
* Added audit logs to Safe API methods.
* When a secret persist error occurs, it is logged.
* Improvements in development workflow (*enabling local registries*)

### Changed

* Retry persisting secrets to disk one more time before erring out.
* Added a channel mechanism to funnel disk errors instead of suppressing them.

## [v1.11.0] - 2023-01-28

### Current State

As per this release, Aegis is able to securely dispatch secrets to workloads
within a single cluster; it encrypts and backs up secrets to a volume; and
if it crashes, it recovers its state from the backups. The code is stable
enough and the solution can be used at a production capacity.

That’s also why we started a changelog: Before this point in time, it
was a figurative cambrian soup, and it was too chaotic to even keep a changelog.

From this point on, we will record the changes, deprecations,
removals, fixes, and security patches more carefully.

### Added

* Added a script to do functional test before every release cut.
* Whenever a secret is saved, the old secret is backed up.
* Ability to delete secrets.
* In-memory mode: Secrets can optionally NOT be persisted on disk and kept
  only in memory. This is not the default behavior and needs to be set
  using an environment variable.

### Changed

* Usability improvements to the website.
* Added more tutorials and articles to the website to explain how Aegis works.
* Aegis Safe responds with RFC3339 instead of RubyDate. RFC3339 is also what Go
  uses when serializing dates into JSON.
* Added a buffered channel to save secrets on disk to improve performance
  especially for slower disks.
* Better logs.

<!--
Added
Changed
Deprecated
Removed
Security
-->