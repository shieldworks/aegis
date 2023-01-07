![Aegis](../assets/aegis-banner.png "Aegis")

## Aegis Safe

**Safe** (`aegis-safe`) is the part that does most of the dirty work:

* It acts as the central in-memory secrets store.
* **Sentinel** talks to **Safe** to register secrets.
* **Sidecar** talks to **Safe** to get the secrets that the workload needs.
