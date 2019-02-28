Burning Rubber Paste
====================

`brpaste` is a small and fast pastebin service.
It provides a lightweight REST-like interface and client-side syntax highlighting (if desired).
It's small and fast because it relies on redis to perform the actual storage.

### Speed
Redis is [fast](https://redis.io/topics/benchmarks).
D is [pretty fast](https://github.com/kostya/benchmarks).
D is in the general range of C++ and Rust, significantly beating out common other choices (such as python, nodejs and lua).
However, D is (subjectively) a much more friendly syntax than either C++ or Rust (note: this is written as someone that's spent a good portion of their life writing C++ professionally).
Further, D is safe (it includes a garbage collector, and more robust RAII than C++ does, as well as more compile-time tests than C++).
Rust would be another good choice for writing this.

### Configuration
There is no configuration within `brpaste`, besides the basics (what addresses to listen on, what port to listen on, how to connect to redis).
This is because the actual job being done is relatively minimal.
However, because redis is used, your instance can be greatly configured.
For instance, you could make all your pastes expire in 10 minutes.
Or you could make them never expire.
The hosted instance over at https://p.toastin.space limits memory usage to 250mb, and expires the least frequently used keys first.
This is the recommended configuration.

### Deployment Difficulty
`brpaste` is distributed as a single binary file.
All other files (such as html) are baked into the binary.
It is planned to offer statically linked to musl versions in the future.
This is possible thanks to the `diet` templates (inspired by pugjs) provided by vibe-d, which are computed at compile-time.

### Other Utilities
For server-side helpers and utilities (such as openrc scripts, systemd unit files, and anything else of the sort), see the `server/` directory.
For client-side helpers and utilities (such as a `sprunge(1)`-like script, an `anypaste` plugin, and anything else of the sort), see the `client/` directory.
For documentation, see the `doc/` directory, but do note that you will need a compliant asciidoc parser to compile it.
