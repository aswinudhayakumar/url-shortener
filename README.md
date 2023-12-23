\*\*

# URL_SHORTENER

This service is writen in **go fiber** framework, as per its name, this returns a short URL for a URL that you are sharing.

This service has 3 Endpoints,

- **shortURL** - http://localhost:3000/?url=https://www.youtube.com/, here you can share the URL that need to be shortened in the URL query param called `url`, in response this returns the shortened URL.
- **getMetrics** - http://0.0.0.0:3000/metrics. this returns the top 3 URLs shortened in this service.

## How to run ?

With these below commands, you can run this service,

- `make dev` to run this in developer mode.
- `make build` to build the service
- `make run-build` to run the service, which already has a build
- `make run` to build and run the build

This service, runs on the port `3000`

You can also pull this service as a image from here - https://hub.docker.com/r/aswin8799/url-shortener
