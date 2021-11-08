# test-harness

This test-harness can be used to investigate the behavior of HTTP servers and proxies. You can for example use it to look for HTTP Request Smuggling issues. You set the harness up with some services it should run, you supply it with some requests to send and then it sends all the requests to all the systems and creates a report with the resulting responses.

When testing proxies the harness starts a dummy backend server that receives the proxied request and responds with a static 200 OK response every time.

If you want to look for HRS issues in servers, make sure that they are set up to echo the body of the received request back in the response. Otherwise there is no way of knowing how the server interprets the request. For proxies, you can just look at how the proxy forwarded the request.

This repo includes some exaple services in the `/services` directory. The harness uses Docker.

This harness was inspired by https://github.com/regilero/HTTPWookiee

## Using it

Make sure that you are in the `test-harness` directory when running the harness!

1. Install go
2. Run `go build .` to build the harness
3. Run the harness using `sudo ./test-harness <.req-files>`. The program needs `sudo` to be able to run `docker` commands.

When you run the command you need to specify a path to some `.req`-files. Example `.req`-files can be found in the top-level `/requests` directory in this repo.

After the harness has finished it will produce a `report.html` and a `report.mini.html` which contain the results. All reports over all runs will be saved in the `/reports` directory.

## Disabling a system
Remove the respective folder in `/services/proxies` or in `/services/servers`.

## Adding a new system to test
You must add a new folder in `/services/servers` or in `/services/proxies`. In that folder you must have a Dockerfile which starts the system and listens on port 80 inside the container. The harness will build the Dockerfile in that directory, so place any files needed there.

If you are adding a new proxy you must also add a BACKEND_PORT file containing the port that the dummy backend should use. You must also configure the proxy to forward requests to `host.docker.internal:<port>` or alternatively `172.17.0.1:<port>`, where `<port>` is the one specified in BACKEND_PORT. Make sure that this number doesn't collide with any of the other proxies BACKEND_PORTs!

## Disclaimer

This software is kinda crappy, but it works good enough for its intended purpose. Feel free to make a better version of it. Here are some known ways that it is crappy (there are probably more ways):

- It starts one goroutine for each server and two for each proxy (one for running the proxy and one for running the dummy backend). So if you want to test 8 proxies at the same time the harness will start 16 goroutines at the same time. This is fine if you have many cores, but if you have few cores the harness might hang. So try running it with fewer systems at a time. The fix to this would be to make the program only run a few of the systems in parallell, not more than the given number of cores are able to handle.
- The program is largely timing-based. There are sleeps here and there to wait for systems starting up and messages being sent and delivered. If you add a new system and it seems to not work, maybe you should investigate this.
- I'm certain there is a bug in how the dummy backend waits for a request. I know that I even found the bug some time ago and tried to fix it. But now I've forgotten what exactly the problem was.

(But despite all these issues, it worked well enough to find bugs.)
