```text
 .-'_.---._'-.
 ||####|(__)||   Protect your secrets, protect your business.
   \\()|##//       Secure your sensitive data with Aegis.
    \\ |#//                  <aegis.z2h.dev>
     .\_/.
```

## Aegis Sentinel

**Sentinel** is the only pod that can directly talk to **Sidecar**.

[Watch this demo video][aegis-demo-video] to learn more about how **Sentinel**
is used to register secrets to workloads.

## Netshoot

The development version of `aegis-sentinel` comes bundled with a modified version
of [netshoot][netshoot] the help debugging the network. This is not recommended
for a production deployment as it increases the attack surface.

`z2hdev/aegis-sentinel` v0.5.43 and above do not bundle netshoot by default. To
use `netshoot` inside sentinel, you’d need `z2hdev/aegis-sentinel-dev` as your
container image instead.

[aegis-demo-video]: https://vimeo.com/v0lkan/secrets "Aegis: Keep your secrets… Secret."
[netshoot]: https://github.com/nicolaka/netshoot "Netshoot:  Docker + Kubernetes network trouble-shooting swiss-army container"