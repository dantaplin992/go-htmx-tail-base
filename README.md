# Go-HTMX-Hyperscript-Tailwind Base

A base project for making highly interactible websites with Go, HTMX, Hyperscript and Tailwind.

The project is set up to use Go as a server, with HTMX & Hyperscript to handle DOM manipulation, and Tailwind for styling. There are 0 dependencies on either the client or server side, as Go offers an http server out of the box, HTMX & Hyperscript are both standalone libraries that are included in the project directly via script tags in the index.html, and Tailwin compiles the CSS into a single file that is also included in index.html directly.

A Dockerfile is also included for easy(ish) deployment.

### How It Works

HTMX extends HTML elements to allow for AJAX requests to be made to the server, and for the server to respond with HTML that can be inserted into the DOM. The `hx-trigger` attribute is used to specify what event should trigger the request, and the `hx-get` and `hx-post` attributes are used to specify the route & method to request. The server responds with an HTML string, which is then inserted into the DOM at the location specified by the `hx-target` attribute, with the optional `hx-swap` attribute specifying how the new content should be inserted.

Hyperscript is a complementary library that allows for smaller, client-side DOM manipulation such as the addition/removal of classes on elements, or changes to an element's text content. It is used in conjunction with HTMX where an AJAX request is not necessary, but the DOM still needs to be manipulated. The user can define trigger events, target elements, scoped variables, loops etc with plain-english syntax.

Tailwind is used because styling everything with vanilla CSS is a hellish experience.

Using this framework it's possible to create a fast & highly interactible website with minimal effort & without writing a single line of JavaScript (The dream is alive).

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
