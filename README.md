This is a simple caddy proxy example, with a static directory, and other
requests sent to a Go (sub)server.

To run, in one terminal window do:

    go run go/hello.go

And in another terminal window do:

    caddy

Now, in your browser go to `localhost:8080` to hit caddy, or `localhost:8081`
to hit the Go server directly.
