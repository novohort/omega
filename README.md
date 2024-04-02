# The Omega Programming Language

This is the main source code repository for Omega. It contains the interpreter, Alpha.

## About Omega

Omega is (obviously) in its extreme infancy stage; it hasn't even attached to the uterus lining yet. Be advised that everything has the potential to change as time goes on, but we're planning to prevent any breaking changes as much as possible so we can maintain a singular backwards-compatible spec as opposed to Rust's tendency to break previous versions.

Omega is designed to be a general-purpose programming language. Sure, it'd be great if we can set it up so it can accept multiple programming paradigms (functional, OOP, etc) in order to ease adoption from developers of all flavors, but that's wildly ambitious especially at this stage.

## About Alpha, the Omega interpreter

Alpha originally started as a compiler written in Javascript and leveraging rustc, but has since changed direction.

Starting in v0.3.0, Alpha is a live interpreter (REPL) for Omega, and is written in Go.

## Getting Started

You'll need to install Go in order to make use of Alpha (and thereby, Omega).

Once that's done, you may launch the REPL by running `go run alpha.go` in your terminal just like you would with any other REPL. Be sure that you're in the correct directory (`interpreter`)


## Contributing

We're thrilled you're interested in contributing to Omega!

### Branch Guidelines

#### `main`

This is the primary branch. It contains the most stable and tested version of the language. All development eventually merges back to `main`, but only after thorough testing and review. The `main` branch is what the average user should be using in production.

#### `dev`

This branch serves as the primary integration branch for new features. It's the "working" version of the language and contains new features that are under development but might notbe fully tested. Features should be developed in separate branches and merged into `dev` when they reach a stable state.

#### `feature/*`

For new features or significant changes, create individual branches off `dev`. Name these branches clearly based on the feature or change being implemented, e.g., `feature/new-syntax`, `feature/variable-bindings`. Once the feature is completed, tested, and approved, it gets merged back into the `dev` branch.

#### `release/*`

When `dev` reaches a state where it's ready to be released (after all features planned for the release are merged and tests are passed), create a release branch, e.g., `release/v0.3.0`. This branch is used for final testing and documentation updates. After it's fully tested and ready, it merges into `main` and back into `dev` to ensure any versionupdates or last-minute fixes are not lost.

#### `hotfix/*`

For urgent fixes that need to be applied directly to the stable version, create hotfix branches from `main`. For example, `hotfix/critical-bug-fix`. After the fix is implemented and tested, it merges back into both `main` and `dev` to ensure the fix is applied across all relevant versions.

### Development Guidelines

#### Develop in Isolation

Work on new features and non-urgent fixes in their respective `feature/*` branches to keep changes isolated until they're ready.

#### Regular Merges

Regularly merge changse from `main` into `develop` and then into feature branches to keep them up to date and minimize merge conflicts.

#### Code Reviews

Use pull requests for merging from `feature/*`, `release/*`, and `hotfix/*` branches to facilitate code reviews and ensure quality.

#### Continuous Integration

Automate testing as much as possible. Use CI tools to run tests on all branches but especially on `dev`, `release/*`, and `hotfix/*` branches to catch issues early.
