# zet

A Zettelkasten management CLI, because I procrastinate

## Building

There is a bundled Makefile you can use:

```shell
make
```

And to install the utilities, use the `install` target.

```shell
# local install
make install PREFIX=$HOME/.local/bin

# global install
sudo make install
```

## Usage

First, initialize a Zettelkasten (configurable)

```shell
zet init
```

Then, add a Zet

```shell
zet new
<opens in vi or $EDITOR>
```

It will generate a git repo and track/manage changes for you

- **(Not yet implemented)** To sync those changes remotely, run `zet sync`
- **(Not yet implemented)** To sync those changes remotely, run `zet sync`
