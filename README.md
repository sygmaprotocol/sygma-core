# Sygma Core

<a href="https://discord.gg/ykXsJKfhgq">
  <img alt="discord" src="https://img.shields.io/discord/593655374469660673?label=Discord&style=for-the-badge&logo=discord&logoColor=white" />
</a>
<a href="https://golang.org">
<img alt="go" src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" />
</a>

&nbsp;

Sygma-Core encompasses reusable modules to easily build relayers based on Sygma contracts.

&nbsp;

*Project still in deep beta*

&nbsp;

### Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Local Setup](#local-setup)
4. [Contributing](#contributing)



&nbsp;

## Contributing

Sygma-core is a open project and welcomes contributions of all kinds: code, docs, and more. If you wish to submit more complex changes,
please check up with the core devs first on our [Discord Server](https://discord.gg/ykXsJKfhgq) (look for CHAINBRIDGE category) or submit an issue on the repo to ensure those changes are in line with the general
philoshophy of the project or get some early feedback.

When implementing a change:

1. Adhere to the standard Go formatting guidelines, e.g. [Effective Go](https://golang.org/doc/effective_go.html). Run `go fmt`.
2. Stick to the idioms and patterns used in the codebase. Familiar-looking code has a higher chance of being accepted than eerie code. Pay attention to commonly used variable and parameter names, avoidance of naked returns, error handling patterns, etc.
3. Comments: follow the advice on the [Commentary](https://golang.org/doc/effective_go.html#commentary) section of Effective Go.
4. Minimize code churn. Modify only what is strictly necessary. Well-encapsulated changesets will get a quicker response from maintainers.
5. Lint your code with [`golangci-lint`](https://golangci-lint.run) (CI will reject your PR if unlinted).
6. Add tests.
7. Title the PR in a meaningful way and describe the rationale and the thought process in the PR description.
8. Write clean, thoughtful, and detailed [commit messages](https://chris.beams.io/posts/git-commit/).


### Submiting a PR

Fork the repository, make changes and open a PR to the `main` branch of the repo. Pull requests must be cleanly rebased on top of `main` and changes require at least 2 PR approvals for them to be merged.

### Reporting an issue

A great way to contribute to the project is to send a detailed report when you encounter an issue. We always appreciate a well-written, thorough bug report, and will thank you for it!

When reporting issues, always include:
 - sygma-core version
 - modules used
 - logs (don't forget to remove sensitive data)
 - tx hashes related to issue (if applicable)
 - steps required to reproduce the problem

 Putting large logs into a [gist](https://gist.github.com) will be appreciated.

&nbsp;

# ChainSafe Security Policy

## Reporting a Security Bug

We take all security issues seriously, if you believe you have found a security issue within a ChainSafe
project please notify us immediately. If an issue is confirmed, we will take all necessary precautions
to ensure a statement and patch release is made in a timely manner.

Please email us a description of the flaw and any related information (e.g. reproduction steps, version) to
[security at chainsafe dot io](mailto:security@chainsafe.io).

## License

_GNU Lesser General Public License v3.0_
