# Go-HTMX-Hyperscript-Tailwind Base

A base project for making highly interactible websires with Go, HTMX, Hyperscript and Tailwind.

The project is set up to use Go as a server, with HTMX to handle AJAX requests, Hyperscript to handle DOM manipulation, and Tailwind for styling. There are 0 dependencies on either the client or server side, as Go offers an http server out of the box, HTMX & Hyperscript are both standalone libraries that are included in the project directly via script tags in the index.html, and Tailwind is compiled into a single css file that is also included in index.html.

## Prerequisites

Only a couple of tools are needed to install & run the project:
- Go: https://go.dev/doc/install
- Tailwind standalone CLI:https://github.com/tailwindlabs/tailwindcss/releases/tag/v3.4.3 (or latest release)

## Initialising the project
- To initialize the Go project, run `go mod init {name of the project}`
- To initialize Tailwind and create a `tailwind.config.js`, run `{path-to-tailwindcss-cli} init`
- Add `'./public/*.html'` to the 'content' section of `tailwind.config.js`

And you're all set!

## Development
- Run `{path/to/tailwindcss} -i {input/file/path}.css -o public/style.css --watch` to watch for changes & rebuild css file using tailwindcss cli, which will automatically insert styles for tailwind utility classes as they are introduced in templates
- Run `go run main.go` to run the server in development mode - this will look for the 'PORT' environment variable, or bind to port 8000 if this is not available

## Building
- Run the tailwind command above, but without the `--watch` flag to simply build the css file. it can be run with `--minify` to output a minified css file
- Run `go build .` to compile the server code
- It's probably a good idea to create a batch file (.sh on mac or .bat on Windows) to take care of the build process in one step, for example just by running `.\compile.bat` instead of having to do each step individually.
