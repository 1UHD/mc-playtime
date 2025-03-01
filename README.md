<div align="center">
<h1>MC-Playtime</h1>
<h3>a Minecraft playtime calculator written in Go</h2>
</div>

## Functionality

This program analyzes the log files of your installed Minecraft launchers and instances to determine the time spent in each session. It does this by identifying the first and last log messages of each session and calculating the total duration.

Currently, the following Minecraft launchers are supported:

<ul>
<li>Vanilla
<li>Lunar Client
<li>Badlion Client
<li>Curseforge Launcher
<li>MultiMC Launcher
<li>PrismLauncher
</ul>

## Installation

### Precompiled Binaries

You can download the latest release for your platform from the Releases tab.

> [!Warning]
> The Linux executable is currently unavailable due to issues with the Linux for Darwin C cross-compiler.

### Building from source

**Prerequisites**

1. Install the [Go programming language](https://go.dev/doc/install).
2. Clone the repository:

```sh
git clone https://github.com/1UHD/mc-playtime.git
```

3. Install the required dependencies:

```sh
go mod download
```

**Compiling the project**
This project uses [Fyne](https://fyne.io/) for its GUI, which relies on CGO. Therefore, a C compiler is required for compilation. I recommend the GCC compiler.

#### MacOS

MacOS comes with the clang compiler, so we can skip the step of installing one.

**Intel Macs:**

```sh
CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o mc-playtime-mac-intel
```

**Apple Silicon Macs:**

```sh
CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o mc-playtime-mac-arm
```

#### Linux

For installing GCC:
**Debian based:**

```sh
sudo apt update && sudo apt install -y build-essential gcc
```

**Fedora based:**

```sh
sudo dnf install gcc glibc-devel
```

**Arch based:**

```sh
sudo pacman -S base-devel
```

Now you can compile mc-playtime
**x86_64 (basically every desktop):**

```sh
CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o mc-playtime-linux-x86_64
```

**ARM64 (f.e. Raspberry Pi):**

```sh
CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o mc-playtime-linux-arm64
```

#### Windows

Install [GCC](https://gcc.gnu.org/install/download.html). If you use the scoop package manager:

```sh
scoop install gcc
```

Now you can compile mc-playtime
**x86-64:**

```sh
SET CGO_ENABLED=1 SET CC=gcc go build -o mc-playtime-windows-x86_64.exe
```

**32-bit:**

```sh
SET CGO_ENABLED=1 SET CC=gcc go build GOARCH=386 -o mc-playtime-windows-32-bit.exe
```
