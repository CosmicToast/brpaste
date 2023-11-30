Burning Rubber Paste
====================

[![Build Status](https://cloud.drone.io/api/badges/5paceToast/brpaste/status.svg)](https://cloud.drone.io/5paceToast/brpaste)
[![Go Report Card](https://goreportcard.com/badge/toast.cafe/x/brpaste)](https://goreportcard.com/report/toast.cafe/x/brpaste)

`brpaste` is a small and fast pastebin service.
It provides a lightweight REST-like interface and client-side syntax highlighting (if desired).
It's small and fast because it relies on redis to perform the actual storage.

### Project Status
Brpaste has been in pure maintenance mode for a while.
Do not mistake the recent (at the time of writing) commits as it being revived: I added boltdb mode to make it easier to migrate (again).
I'm planning to write a new (smaller) thing eventually, but for now I'm still hosting this, and am making it simpler to host.

### Quickstart
#### Go edition
`go get -u toast.cafe/x/brpaste`
#### Github edition
Download the correct binary from the releases page.
#### CI (nightly master) edition
Download your build from https://minio.toast.cafe/cicd/brpaste/brpaste-$OS-$ARCH where `$OS` is something like "openbsd" and `$ARCH` is something like "amd64".
Note that the github edition binaries are just these from immediately after a release.

### Platform Support
Linux AMD64 is the primary platform.
Everything else is "best effort".
For a full list of supported platforms, see the releases page (all the binaries on there).

### Speed
It's just fast.
I could put a bunch of benchmarks here but people didn't really seem to care in the previous version anyway.
Just trust me.
It's fast.

### Configuration
There is no configuration within `brpaste`, besides the basics (what address/port to listen on, how to connect to redis).
This is because the actual job being done is relatively minimal.
However, because redis is used, your instance can be greatly configured.
For instance, you could make all your pastes expire in 10 minutes.
Or you could make them never expire.
The hosted instance over at https://brpaste.xyz limits memory usage to 250mb, and expires the least frequently used keys first.
This is the recommended configuration.

### Deployment Difficulty
`brpaste` is distributed as a single binary file.
All other files (such as html) are baked into the binary.

### Stable IDs
`brpaste` IDs are not the shortest.
What they are, however, is stable.
What does that mean?
When you upload something to `brpaste`, the ID is generated through Murmurhash3 32 bit, and converted into a string of letters and symbols using base64.
Murmurhash3 is suitable for lookups, so collision are sufficiently unlikely within the lifetime of your paste.
However, if you upload the same paste twice, the ID will stay the same.
The memory usage will not increase.
This unlocks a few interesting use-cases (e.g not needing to keep around an open tab of a paste - just paste it again, you'll get the original link back).
The bitsize of the ID is 32 bits, which translates to roughly 6 base64 "digits".
This may not be the shortest, but it is short enough to memorize in one go (see: magic 7 of human working memory; approximated to more likely 6 if letters are involved).
As such, the disadvantage is rather minimal, while the advantage is a nice-to-have, consistent, and cheap (murmurhash3 is fast, and there's no need to do things like keep a counter around).

### Other Utilities
For server-side helpers and utilities (such as openrc scripts, systemd unit files, and anything else of the sort), see the `server/` directory.
For client-side helpers and utilities (such as a `sprunge(1)`-like script, an `anypaste` plugin, and anything else of the sort), see the `client/` directory.
For documentation, see the `doc/` directory, but do note that you will need a compliant asciidoc parser to compile it.
