# 404-placeholder

Simple placeholder for cloud environments until real app is ready.
This is a very simple container that returns 404s for most requests.
One endpoint gets a 200 with an empty JSON object for health checks.
By default, the code listens on port 8080 and serves the health checks
on /health_check, but you can change those with parameters.  This works
as a placeholder for cloud environments like Google Cloud Run for you to
configure your cloud until you are ready to deploy a real server.

## Getting Started

The container is available on Docker Hub and you can run it straight
from there.  The container is hosted on Docker Hub at
jimp/404-placerholder

## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/your/project/tags).

## License

This project is licensed under the MIT License - see the
[LICENSE](LICENSE) file for details
