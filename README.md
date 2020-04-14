# Find Aspect Ratio

A small program to learn some Go by finding the aspect ratio of a rectangle of a given width and height.

## Compiling

The executable is output in the `bin` directory, and will be named `gofar` or `gofar.exe`.

There are also `tasks.json` and `launch.json` files for VS Code users.

### Unix-like

    sh build.sh

### Windows

    build.cmd

## Usage

    gofar -width=1920 -height=1080 [-nonewline]

This should output:

    16:9

Specifying the `-nonewline` flag will omit the newline from the output, making it simpler to pipe the result into something else.
