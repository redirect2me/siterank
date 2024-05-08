# Site Rank API [<img alt="Site Rank API Logo" src="https://siterank.redirect2.me/favicon.svg" height="96" align="right" />](https://siterank.redirect2.me/)

![NodePing status](https://img.shields.io/nodeping/status/mnjrv1uf-ns8t-4b84-8wlm-0rncdu9wb85x)
![NodePing uptime](https://img.shields.io/nodeping/uptime/mnjrv1uf-ns8t-4b84-8wlm-0rncdu9wb85x)
[![deploy](https://github.com/redirect2me/siterank/actions/workflows/gcr-deploy.yaml/badge.svg)](https://github.com/redirect2me/siterank/actions/workflows/gcr-deploy.yaml)

A simple API to get a website's [Tranco](https://tranco-list.eu/) ranking.


## APIs

* `/api/rank.json?domain=redirect2.me` - gets the result for a single domain
* `/api/multiple.json?domains=redirect2.me,resolve.rs` - gets the results for multiple domains

Pass a `callback` parameter for JSONP.

Pass an `apikey` parameter with a contact email if you would like to be notified if the API changes.

## Using

I'm currently running it at [siterank.redirect2.me](https://siterank.redirect2.me/).  You are welcome to use it for light, non-commercial purposes.  For heavy or commercial use, please run your own copy.

I reserve the right to return "unfortunate" results if abused.

## Other lists

The only thing that is tranco-specific is the loader.  It should be easy to adapt to another list.

* [Majestic Million](https://majestic.com/reports/majestic-million)
* [Cisco Umbrella](https://s3-us-west-1.amazonaws.com/umbrella-static/index.html)
* [Cloudflare Radar](https://radar.cloudflare.com/domains)

## Alternatives

* [eest/tranco-list-api](https://github.com/eest/tranco-list-api) - [blog post](https://blog.sigterm.se/posts/building-an-api-for-tranco/)
* [WangYihang/tranco-go-package](https://github.com/WangYihang/tranco-go-package) - CLI and Go package

## License

[AGPLv3](LICENSE.txt)

## Credits

[![Google CloudRun](https://www.vectorlogo.zone/logos/google_cloud_run/google_cloud_run-ar21.svg)](https://cloud.google.com/run/ "Hosting")
[![Docker](https://www.vectorlogo.zone/logos/docker/docker-ar21.svg)](https://www.docker.com/ "Deployment")
[![Git](https://www.vectorlogo.zone/logos/git-scm/git-scm-ar21.svg)](https://git-scm.com/ "Version control")
[![Github](https://www.vectorlogo.zone/logos/github/github-ar21.svg)](https://github.com/ "Code hosting")
[![golang](https://www.vectorlogo.zone/logos/golang/golang-ar21.svg)](https://golang.org/ "Programming language")
[![Google Noto Emoji](https://www.vectorlogo.zone/logos/google/google-ar21.svg)](https://github.com/googlefonts/noto-emoji/ "Logo")
[![NodePing](https://www.vectorlogo.zone/logos/nodeping/nodeping-ar21.svg)](https://nodeping.com?rid=201109281250J5K3P "Uptime monitoring")
[![VectorLogoZone](https://www.vectorlogo.zone/logos/vectorlogozone/vectorlogozone-ar21.svg)](https://www.vectorlogo.zone/ "Logos")
[![water.css](https://www.vectorlogo.zone/logos/netlifyapp_watercss/netlifyapp_watercss-ar21.svg)](https://watercss.netlify.app/ "Classless CSS")
