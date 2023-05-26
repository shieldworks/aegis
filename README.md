# Aegis

![Aegis](assets/aegis-icon.png "Aegis")

keep your secrets… secret

[spire]: https://spiffe.io/ "SPIFFE: Secure Production Identity Framework for Everyone"

## The Elevator Pitch

[**Aegis**](https://aegis.ist) is a delightfully-secure Kubernetes-native 
secrets store.

**Aegis** keeps your secrets secret.

With **Aegis**, you can rest assured that your
sensitive data is always **secure** and **protected**.

**Aegis** is perfect for securely storing arbitrary configuration information at
a central location and securely dispatching it to workloads.

## Tell Me More

**Aegis** is a cloud-native secure store for secrets management. 
It provides a minimal and intuitive API, ensuring practical security without 
compromising user experience. 

Aegis is **resilient** and **secure** by default, storing sensitive data in memory 
and encrypting any data saved to disk. 

Endorsed by industry experts, **Aegis** is a ground-up re-imagination of 
secrets management, leveraging SPIFFE for authentication and providing 
a cloud-native way to manage secrets end-to-end.

## Getting Your Hands Dirty

Before trying **Aegis**, you might want to learn about its 
[architecture][architecture] and [design goals][design].

Once you are ready to get started, [see the Quickstart guide][quickstart].

Or, if you one of those who “*learn by doing*”, you might want to dig into the 
implementation details later. If that’s the case, you can directly jump to the 
fun part and [follow the steps here][installation] to install 
**Aegis** to your Kubernetes cluster.

## Dive Into Example Use Cases

There are several examples demonstrating **Aegis** sample use cases 
[inside the `./examples/` folder](./examples).

## Container Images

Pre-built container images of **Aegis** components can be found at:
<https://hub.docker.com/u/aegishub>.

## Build Aegis From the Source

[You can also build **Aegis** from the source][build].

## The Roadmap

[We publicly track all **Aegis** plans on a GitHub project][roadmap]. You can
check it out to get a glimpse of the current planned features, and how the
future of **Aegis** looks like.

[roadmap]:https://github.com/orgs/shieldworks/projects/1  "The Roadmap"
[installation]: https://aegis.ist/docs/#installation "Install Aegis"
[build]: https://aegis.ist/docs/contributing/#building-deploying-and-testing
[architecture]: https://aegis.ist/docs/architecture/ "Aegis Deep Dive"
[design]: https://aegis.ist/docs/philosophy/ "Aegis Philosphy"
[quickstart]: https://aegis.ist/docs/

## Status of This Software

**Aegis**, is under dynamic and progressive development.

The code that we’ve officially signed and released maintains a 
high standard of stability and dependability. However, we do encourage 
it be used in a production environment (*at your own risk—[see LICENSE](LICENSE)*).

It’s important to note that, technically speaking, **Aegis** currently holds the 
status of an alpha software. This simply means that as we journey towards our milestone of 
`v1.0.0`, it's possible for changes to occur—both major and minor. While this might mean 
some aspects are not backward compatible, it's a testament to our unwavering commitment 
to refining and enhancing **Aegis**.

In a nutshell, we are ceaselessly pushing the boundaries of what’s possible, while 
ensuring our software stays as dependable and effective for production use.

## 🦆🦆🦆 (*Docs*)

[Official documentation available on **aegis.ist**](https://aegis.ist).

## A Note on Security

We take **Aegis**’ security seriously. If you believe you have found a vulnerability,
please responsibly disclose by contacting [security@aegis.ist](mailto:security@aegis.ist).

## A Tour Of Aegis

[Check out this quickstart guide][quickstart] for an overview of **Aegis**.

[quickstart]: https://aegis.ist/docs/

## Community

Open Source is better together.

If you are a security enthusiast, [**join Aegis’ Slack Workspace**][slack-invite]
and let us change the world together 🤘.

## Links

### General Links

* **Homepage**: <https://aegis.ist>
* **Documentation**: <https://aegis.ist/docs/>
* **Changelog**: <https://aegis.ist/changelog/>
* **Community**: [Join **Aegis**’ Slack Workspace][slack-invite]
* **Contact**: <https://aegis.ist/contact/>
* **Media Kit**: <https://aegis.ist/media/>
* **Changelog**: <https://aegis.ist/changelog/>

### Guides and Tutorials

* **Installation and Quickstart**: <https://aegis.ist/docs/register/>
* **Local Development Instructions**: <https://aegis.ist/docs/contributing/>
* **Aegis Go SDK**: <https://aegis.ist/docs/sdk/>
* **Aegis CLI**: <https://aegis.ist/docs/sentinel/>
* **Architectural Deep Dive**: <https://aegis.ist/docs/architecture/>
* **Configuration**: <https://aegis.ist/docs/configuration/>
* **Design Philosophy**: <https://aegis.ist/docs/philosophy/>
* **Production Deployment Tips**: <https://aegis.ist/production/>

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

## Folder Structure

Here are the important folders and files in this repository:

* `./app`: Contains core **Aegis** components’ source code.
    * `./app/init-container`: Contains the source code for the **Aegis Init Container**.
    * `./app/safe`: Contains the source code for the **Aegis Safe**.
    * `./app/sentinel`: Contains the source code for the **Aegis Sentinel**.
    * `./app/sidecar`: Contains the source code for the **Aegis Sidecar**.
* `./core`: Contains core modules that are shared across **Aegis** components.
* `./examples`: Contains the source code of example use cases.
* `./hack`: Contains scripts that are used for building, publishing, development
  and testing.
* `./k8s`: Contains Kubernetes manifests that are used to deploy **Aegis** and
  its use cases.
* `./sdk`: Contains the source code of the **Aegis SDK**.
* `./CODE_OF_CONDUCT.md`: Contains **Aegis** Code of Conduct.
* `./SECURITY.md`: Contains **Aegis** Security Policy.
* `./LICENSE`: Contains **Aegis** License.
* `./Makefile`: Contains **Aegis** Makefile that is used for building,
  publishing, deploying, and testing the project.

## One More Thing… How Do I Pronounce “Aegis”?

[We have an article for that too 🙂][pronounce].

[pronounce]: https://aegis.ist/pronunciation/

## Changelog

You can find the changelog, and migration/upgrade instructions (*if any*)
on [**Aegis**’ Changelog Page](https://aegis.ist/changelog/).

## What’s Coming Up Next?

You can see the project’s progress [in these **Aegis** boards][mdp].

The board outlines what are the current outstanding work items, and what is
currently being worked on.

[mdp]: https://github.com/orgs/shieldworks/projects/1/views/2

## Code Of Conduct

[Be a nice citizen](CODE_OF_CONDUCT.md).

## Contributing

To contribute to **Aegis**, [follow the contributing guidelines](CONTRIBUTING.md)
to get started.

Use GitHub issues to request features or file bugs.

## Communications

* [**Slack** is where the community hangs out][slack-invite].
* [Send comments and suggestions to **feedback@aegis.ist**](mailto:feedback@aegis.ist).

## Maintainers

As of now, I, [Volkan Özçelik][me], am the sole maintainer of **Aegis**.

[me]: https://github.com/v0lkan "Volkan Özçelik"

Please send your feedback, suggestions, recommendations, and comments to
[feedback@aegis.ist](mailto:feedback@aegis.ist).

We’d love to have them.

## License

[MIT License](LICENSE).

[slack-invite]: https://join.slack.com/t/aegis-6n41813/shared_invite/zt-1myzqdi6t-jTvuRd1zDLbHX0gN8VkCqg "Join aegis.slack.com"
[aegis-web]: https://aegis.ist/
[aegis-projects]: https://aegis.ist/docs/architecture/#projects
[aegis-repo]: https://github.com/shieldworks/aegis
