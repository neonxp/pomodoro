# pomodoro
Pomodoro timer for nerds

## Installation

### Binaries

Grab binaries for different OS from https://github.com/neonxp/pomodoro/releases

### Homebrew

```
brew install neonxp/tap/pomodoro
```

### Docker

Build image:

```shell
docker build -t pomodoro:master .
```

or pull image from GitHub:

```shell
docker pull ghcr.io/neonxp/pomodoro:master
```

Run image
```shell
docker run -it --rm pomodoro:master
```

### With golang
```
go install github.com/neonxp/pomodoro@latest
```

## Usage

Just run

```shell
pomodoro
```

to stop - press Ctrl+C or kill process
