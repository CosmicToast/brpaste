Burning Rubber Paste
====================

`brpaste` is a small and fast pastebin service.
It provides a lightweight REST-like interface and client-side syntax highlighting (if desired).
It's small and fast because it relies on redis to perform the actual storage.

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
