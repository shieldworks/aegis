# Aegis

![Aegis](assets/aegis-icon.png "Aegis")

keep your secrets… secret

[spire]: https://spiffe.io/ "SPIFFE: Secure Production Identity Framework for Everyone"

## Status of This Software

This project is in **active development**.

The code that have been officially signed and released is stable,
has passed integration tests, and can be used in a production 
environment (*at your own risk—[see LICENSE](LICENSE)*).

However,—semantically-speaking—**Aegis** is still an **alpha** software.
Meaning, until **Aegis** reaches `v1.0.0`, nothing is backwards compatible 
and anything can change with or without notice.

## A Note on Security

We take **Aegis**’ security seriously. If you believe you have found a vulnerability, 
please responsibly disclose by contacting [security@aegis.ist](mailto:security@aegis.ist).

## Links

* **Homepage**: <https://aegis.ist>
* **Documentation**: <https://aegis.ist/docs/>
* **Changelog**: <https://aegis.ist/changelog/>
* **Community**: [Join Aegis’ Slack Workspace][slack-invite]
* **Contact**: <https://aegis.ist/contact/>
* **Media Kit**: <https://aegist.ist/media/>

[slack-invite]: https://join.slack.com/t/aegis-6n41813/shared_invite/zt-1myzqdi6t-jTvuRd1zDLbHX0gN8VkCqg "Join aegis.slack.com"

## A Tour Of Aegis

[Check out this quickstart guide][quickstart] for an overview of **Aegis**.

[quickstart]: https://aegis.ist/docs/

## About Aegis

**Aegis** is a delightfully-secure Kubernetes-native secrets store.

**Aegis** keeps your secrets secret. 

With **Aegis**, you can rest assured that your 
sensitive data is always **secure** and **protected**. 

**Aegis** is perfect for securely storing arbitrary configuration information at
a central location and securely dispatching it to workloads.

[Check out **Aegis**’s website][aegis-web] for more information.

[aegis-web]: https://aegis.ist/

## Projects 

[**Aegis** consists of the several sister projects][aegis-projects] that is
managed from a [central repository][aegis-repo]. This **README** you are
reading right now is from that central repository.

[aegis-projects]: https://aegis.ist/docs/architecture/#projects
[aegis-repo]: https://github.com/zerotohero-dev/aegis

## Community

If you are a security enthusiast, [**join Aegis’ Slack Workspace**][slack-invite]
and let us change the world together 🤘.

## Installation

[Check out this quickstart guide][quickstart] for an overview of **Aegis**,
which also covers **installation** and **uninstallation** instructions.

[quickstart]: https://aegis.ist/docs/

You need a **Kubernetes** cluster and sufficient admin rights on that cluster to
install **Aegis**.

## Usage

[This tutorial about “**Registering Secrets Using Aegis**”][register] covers
several usage scenarios.

[register]: https://aegis.ist/docs/register/

## Architecture Details

[Check out this **Aegis Deep Dive**][deep-dive] article for an overview
of **Aegis** system design and how each component fits together.

[deep-dive]: https://aegis.ist/docs/architecture/

## One More Thing… How Do I Pronounce “Aegis”?

[We have an article for that too 🙂][pronounce].

[pronounce]: https://aegis.ist/docs/pronunciation/

## Changelog

You can find the changelog, and migration/upgrade instructions (*if any*)
on [Aegis’ Changelog Page](https://aegis.ist/changelog/).

## What’s Coming Up Next?

You can see the project’s progress [in these **Aegis** boards][mdp].

The board outlines what are the current outstanding work items, and what is
currently being worked on.

[mdp]: https://github.com/zerotohero-dev/aegis/projects?query=is%3Aopen

[todo-txt]: https://github.com/todotxt "todo.txt"

## Code Of Conduct

[Be a nice citizen](CODE_OF_CONDUCT.md).

## Contributing

It’s a bit chaotic around here, yet if you want to lend a hand,
[here are the contributing guidelines](CONTRIBUTING.md).

## Maintainers

As of now, I, [Volkan Özçelik][me], am the sole maintainer of **Aegis**.

[me]: https://github.com/v0lkan "Volkan Özçelik"

Please send your feedback, suggestions, recommendations, and comments to
[feedback@aegis.ist](mailto:feedback@aegis.ist). 

I’d love to have them.

## License

[MIT License](LICENSE).
