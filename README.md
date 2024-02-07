[![Get-zap](https://github.com/SiliconLabs/get-zap/actions/workflows/go.yml/badge.svg)](https://github.com/SiliconLabs/get-zap/actions/workflows/go.yml)

# What is this?

This is a simple program that anyone can use to retrieve `zap` from the release page on Github.

# Who should use this?

Anyone who wants to get the "correct" version of zap, should be able to use this. Whether it's Matter SDK, or something else.

# Instructions

This is a [Go](https://go.dev/) project. You need go toolchain installed to build it from source code. Many platforms (Linuxes, brew) come with Go toolchains easily installable through your package manager of choice, or you can follow [instructions here](https://go.dev/doc/install).

Once you have go installed:
  1. You can simply run `go run .` to execute the program from source.
  2. Or build it using `go build` and run `get-zap` executable that gets created.
  3. You can run `go install` to build and deploy the executable into your Go bin directory.
  4. If you want to build for a different platform than local, then set the GOOS and GOARCH environment variables as described [here](https://go.dev/doc/install/source#environment) before you run `go build`.

When executing `get-zap` without any arguments, it will by default download the latest stable release of Zap for the local platform.

For all other options, type `get-zap --help` or `go run . --help` if you run from source.

# Examples


1. Default execution without any arguments:
```
[~/git/get-zap (main)]$ ./get-zap
```

2. Download from a different repo than zap (example, latest release of PTI library):
```
[~/git/get-zap (main)]$ ./get-zap --owner SiliconLabs --repo java_packet_trace_library
```

3. List latest zap release:
```
[~/git/get-zap (main)]$ ./get-zap gh list
```

4. List all zap releases:
```
[~/git/get-zap (main)]$ ./get-zap gh list --release all
```

5. List all releases of a different repo:
```
[~/git/get-zap (main)]$ ./get-zap gh list --release all --owner SiliconLabs --repo java_packet_trace_library
```

6. Download a specific zap release:
```
[~/git/get-zap (main)]$ ./get-zap --release v2024.01.05-nightly
```

7. Print help:
```
[~/git/get-zap (main)]$ ./get-zap --help
```