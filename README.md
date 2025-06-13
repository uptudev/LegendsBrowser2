# Legends Browser 2 #

Legends Browser 2 is a multi-platform, open source Legends viewer for Dwarf
Fortress v0.47 written in
[Go](https://go.dev/).

It is a complete rewrite of
[the original Legends Browser](https://github.com/robertjanetzko/LegendsBrowser).

## Features ##

* Works in the browser of your choice (just launch Legends Browser 2 and open
`http://localhost:58881`).
* Recreates Legends mode from Dwarf Fortress, with objects being accessible as
pages with links to related objects.
* Adds several statistics and overviews not found in the base game.

### Important Note ###

Some features require additionally exporting the `legends_plus.xml` from
[DFHack](https://docs.dfhack.org).
To export this file, open Legends mode with DFHack, then run
`exportlegends info` or `exportlegends all`.


## Usage ##

1. Download the latest
[release](https://github.com/robertjanetzko/LegendsBrowser2/releases).
2. Run the application.
3. A browser window should open; if not, navigate to `http://localhost:58881` manually.
4. Choose the `legends.xml` file you want to browse by navigating your file
system to your exported Legends data. This is typically in your base directory
for Dwarf Fortress. Loadable exports should show up in green text with an XML
icon beside them, and will additionally have a '+' icon if the corresponding
`legends_plus.xml` file is loadable.

After loading the XML data, you should see an overview of the world containing
a list of civilizations and a map.

### Options ###

Legends Browser 2 has the following options when ran from the command line:

* `-p <arg>` / `--port <arg>`:
    Use the specified port for serving HTTP.
* `-s `/ `--serverMode`:
    Run in server mode (disables XML file browser)
* `-u <arg>` / `--subUri <arg>`:
    Serve HTTP on `http://localhost:58881/<arg>` instead of `http://localhost:58881/`.
* `-w <arg>` / `--world <arg>`:
    Opens a specific Legends XML file. Useful in conjunction with `-s`.

### Compiling from source ###

The only requirement for building from source is the Go compiler.

1. Clone the repository and open a terminal in the `backend` directory.
2. Run `go build -o legendsbrowser`.

The compiled executable is named `legendsbrowser` and is in the same directory.

## Troubleshooting ##

If you find any bugs, feel free to open an issue here on GitHub

If you have any questions, we have
[a thread on Bay12's forums](http://www.bay12forums.com/smf/index.php?topic=179848.0).
