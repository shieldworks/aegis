
## Aegis Design Decisions

We follow the **guidelines** outlined in the next few sections to achieve these.

Since **Aegis** is still in development, some of these goals discussed in the
following sections may still need to be fully implemented. Regardless,
they are the **guiding principles** we steer towards while shaping the future
of **Aegis**.

### Be Cloud Native

**Aegis** is designed to run on Kubernetes and **only** on Kubernetes.
That helps us leverage Kubernetes concepts like *Operators*, *Custom Resources*,
and *Controllers* to our advantage to simplify workflow and state management.

If you are looking for a solution that runs outside Kubernetes or as a
standalone binary, then Aegis is not the Secrets Store you’re looking for.

### Do One Thing Well

At a 5000-feet level, **Aegis** is a secure Key-Value store. It can securely
store arbitrary values that you, as an administrator, associate with keys. It
does that, and it does that well.

If you are searching for a solution to create and store X.509 certificates,
create dynamic secrets, automate your PKI infrastructure, federate your
identities, use as an OTP generator, policy manager, in short, anything other
than a secure key-value store, then Aegis is likely not the solution you are
looking for.

### Have a Minimal and Intuitive API

As an administrator, there is a limited set of API endpoints that you can
interact with **Aegis**. This makes **Aegis** easy to manage. In addition,
a minimal set of APIs means a smaller attack surface, a smaller footprint, and
a codebase that is easy to understand, test, audit, and develop; all good things.

### Be Practically Secure

Corollary: Do not be annoyingly secure. Provide a delightful user experience
while taking security seriously.

**Aegis** is a secure solution, yet still delightful to operate.
You won’t have to jump over the hoops or wake up in the middle of the night
to keep it up and running. Instead, **Aegis** will work seamlessly, as if it
doesn’t exist at all.

## Secure By Default

**Aegis** stores your sensitive data in memory by default. Yes, that brings
up resource limitations since you cannot store a gorilla holding a banana and
the entire jungle in your store; however, a couple of gigabytes of RAM can store
a lot of plain text secrets in it, so it’s good enough for most practical
purposes.

More importantly, almost all modern instruction set architectures and
operating systems implement [*memory protection*][memory-protection]. The primary
purpose of memory protection is to prevent a process from accessing memory that
has not been allocated to it. This prevents a bug or malware within a process
from affecting other processes or the operating system itself.

[memory-protection]: https://en.wikipedia.org/wiki/Memory_protection "Memory Protection (Wikipedia)"

Therefore, reading a variable’s value from a process’s memory is practically
impossible unless you attach a debugger to it.

So, until we implement ways to securely store a backup of the state data
encrypted on disk, all the secrets **Aegis** has will be held in memory
only: Never persisted to disk and never written or streamed to log files.

Other related works are in progress to make **Aegis**’s system
architecture secure by default. We are slowly and steadily getting there.
You can check out [aegis.txt](aegis.txt) for the overall progress of
what has been done and what is in progress.

## Disaster Recovery and Fault Tolerance

This is an feature that we are actively working on.

As of the current version, recovering the secrets when a workload restarts is
automatic. However, if **Safe** or **Notary** pods are evicted for any reason,
then the in-memory secrets are lost, and an administrator will have to
delete and redeploy everything under the `aegis-system` namespace and
create secrets (*ideally, through a CI pipeline*).

Yes, this can get annoying. And the future versions of **Aegis** will have
measures to prevent this from happening so you, as the ops person,
can `#sleepmore`.