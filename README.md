# yarn

[![Build Status](https://ci.mills.io/api/badges/yarnsocial/yarn/status.svg)](https://ci.mills.io/yarnsocial/yarn)
[![GoDoc](https://pkg.go.dev/git.mills.io/yarnsocial/yarn?status.svg)](https://pkg.go.dev/git.mills.io/yarnsocial/yarn)

📕 yarn is a Self-Hosted, Twitter™-like Decentralised micro-Blogging platform.

- No ads
- No tracking
- Your content
- Your data!

![](https://twtxt.net/media/XsLsDHuisnXcL6NuUkYguK.png)

> Technically `yarnd` is a multi-user
> [twtxt](https://twtxt.readthedocs.io/en/latest/) client in the form  of a
> web app and api. It supports multiple users and also hosts user feeds
> directly and provides a familiar "social" experience with minimal user
> profiles.
>
> It also supports "rich" text by utilising Markdown as well as multimedia
> such as photos, videos and even audio.

There is also a publicly (_free_) available Pod available at:

- https://twtxt.net/

> __NOTE:__ I, [James Mills](https://github.com/prologic), run this first
> (_of which I hope to be many_) `twtxt` instance on pretty cheap hardware on
> a limited budget. Please use it fairly so everyone can enjoy using it
> equally! Please be sure to read the [/privacy](https://twtxt.net/privacy)
> policy before signing up (_pretty straight forward_) and happy Twt'ing! 🤗

> **[Sponsor](#Sponsor)** this project to support the development of new
> features, improving existings ones and fix bugs!
> Or contact [Support](https://twtxt.net) for help with running your own Pod!
> Or host your own Twtxt feed and support our [Extensions](https://dev.twtxt.net).

![Demo_1](https://user-images.githubusercontent.com/15314237/90351548-cac74b80-dffd-11ea-8288-b347af548465.gif)

## Installation

### Using Homebrew

We provide [Homebrew](https://brew.sh) formulae for macOS users for both the
command-line client (`yarn`) as well as the server (`yarnd`).

```console
brew tap yarnsocial/yarn https://git.mills.io/yarnsocial/homebrew-yarn.git
brew install yarn
```

Run the server:

```console
yarnd
```

Run the command-line client:

```console
yarn
```

### Building from source

This is an option if you are familiar with [Go](https://golang.org) development.

1. Clone this repository (_this is important_)

```console
git clone https://git.mills.io/yarnsocial/yarn.git
```

2. Install required dependencies (_this is important_)

Linux, macOS:

```console
make deps
```
Note that in order to get the media upload functions to work, you need to
install ffmpeg and its associated `-dev` packages. Consult your distribution's
package repository for availability and names.

FreeBSD:

- Install `gmake`
- Install `pkgconf` that brings `pkg-config`
`gmake deps`

3. Build the binaries

Linux, macOS:

```console
make
```

FreeBSD:

```console
gmake
```


## Usage

### Command-line Client

1. Login to  your pod:

```#!console
$ ./yarn login
INFO[0000] Using config file: /Users/prologic/.yarn.yaml
Username:
```

2. Viewing your timeline

```#!console
$ ./yarn timeline
INFO[0000] Using config file: /Users/prologic/.yarn.yaml
> prologic (50 minutes ago)
Hey @rosaelefanten 👋 Nice to see you have a Twtxt feed! Saw your [Tweet](https://twitter.com/koehr_in/status/1326914925348982784?s=20) (_or at least I assume it was yours?_). Never heard of `aria2c` till now! 🤣 TIL

> dilbert (2 hours ago)
Angry Techn Writers ‣ https://dilbert.com/strip/2020-11-14
```

3. Posting a Yarn (_post_):

```#!console
$ ./yarn post
INFO[0000] Using config file: /Users/prologic/.yarn.yaml
Testing `yarn` the command-line client
INFO[0015] posting twt...
INFO[0016] post successful
```

### Deploy with Docker Compose

Run the compose configuration:

```console
docker-compose up -d
```

Then visit: http://localhost:8000/

### Web App

Run yarnd:

```console
yarnd -R
```

__NOTE:__ Registrations are disabled by default so hence the `-R` flag above.

Then visit: http://localhost:8000/

You can configure other options by specifying them on the command-line:

```console
$ ./yarnd --help
Usage of ./yarnd:
  -E, --admin-email string          default admin user email (default "support@yarn.social")
  -N, --admin-name string           default admin user name (default "Administrator")
  -A, --admin-user string           default admin user to use (default "admin")
      --api-session-time duration   timeout for api tokens to expire (default 240h0m0s)
      --api-signing-key string      secret to use for signing api tokens (default "PLEASE_CHANGE_ME!!!")
  -u, --base-url string             base url to use (default "http://0.0.0.0:8000")
  -b, --bind string                 [int]:<port> to bind to (default "0.0.0.0:8000")
      --cookie-secret string        cookie secret to use secure sessions (default "PLEASE_CHANGE_ME!!!")
  -d, --data string                 data directory (default "./data")
  -D, --debug                       enable debug logging
      --feed-sources strings        external feed sources for discovery of other feeds (default [https://feeds.twtxt.net/we-are-feeds.txt,https://raw.githubusercontent.com/jointwt/we-are-twtxt/master/we-are-bots.txt,https://raw.githubusercontent.com/jointwt/we-are-twtxt/master/we-are-twtxt.txt])
      --magiclink-secret string     magiclink secret to use for password reset tokens (default "PLEASE_CHANGE_ME!!!")
  -F, --max-fetch-limit int         maximum feed fetch limit in bytes (default 2097152)
  -L, --max-twt-length int          maximum length of posts (default 288)
  -U, --max-upload-size int         maximum upload size of media (default 16777216)
  -n, --name string                 set the pod's name (default "twtxt.net")
  -O, --open-profiles               whether or not to have open user profiles
  -R, --open-registrations          whether or not to have open user registgration
      --session-expiry duration     timeout for sessions to expire (default 240h0m0s)
      --smtp-from string            SMTP From to use for email sending (default "PLEASE_CHANGE_ME!!!")
      --smtp-host string            SMTP Host to use for email sending (default "smtp.gmail.com")
      --smtp-pass string            SMTP Pass to use for email sending (default "PLEASE_CHANGE_ME!!!")
      --smtp-port int               SMTP Port to use for email sending (default 587)
      --smtp-user string            SMTP User to use for email sending (default "PLEASE_CHANGE_ME!!!")
  -s, --store string                store to use (default "bitcask://twtxt.db")
  -t, --theme string                set the default theme (default "dark")
  -T, --twts-per-page int           maximum twts per page to display (default 50)
  -v, --version                     display version information
      --whitelist-domain strings    whitelist of external domains to permit for display of inline images (default [imgur\.com,giphy\.com,reactiongifs\.com,githubusercontent\.com])
pflag: help requested
```

## Configuring your Pod

At a bare minimum you should set the following options:

- `-d /path/to/data`
- `-s bitcask:///path/to/data/twtxt.db` (_we will likely simplify/default this_)
- `-R` to enable open registrations.
- `-O` to enable open profiles.

Most other configuration values _should_ be done via environment variables.

It is _recommended_ you pick an account you want to use to "administer" the
pod with and set the following environment values:

- `ADMIN_USER=username`
- `ADMIN_EMAIL=email`

In order to configure email settings for password recovery and the `/support`
and `/abuse` endpoints, you should set appropriate `SMTP_` values.

It is **highly** recommended you also set the following values to secure your Pod:

- `API_SIGNING_KEY`
- `COOKIE_SECRET`
- `MAGICLINK_SECRET`

These values _should_ be generated with a secure random number generator and
be of length `64` characters long. You can use the following shell snippet
to generate secrets for your pod for the above values:

```console
$ cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 64 | head -n 1
```

**DO NOT** publish or share these values. **BE SURE** to only set them as env vars.

## Production Deployments

### Docker Swarm

You can deploy `twtxt` to a [Docker Swarm](https://docs.docker.com/engine/swarm/)
cluster by utilising the provided `twtxt.yaml` Docker Stack. This also depends on
and uses the [Traefik](https://docs.traefik.io/) ingress load balancer so you must
also have that configured and running in your cluster appropriately.

```console
docker stack deploy -c yarn.yml
```

## Sponsor

Support the ongoing development of twtxt!

**Sponsor**

- Become a [Sponsor](https://www.patreon.com/prologic)
- Contribute! See [Issues](https://git.mills.io/yarnsocial/yarn/issues)

## Contributing

Interested in contributing to this project? You are welcome! Here are some ways
you can contribute:

- [File an Issue](https://git.mills.io/yarnsocial/yarn/issues/new) -- For a bug,
  or interesting idea you have for a new feature or just general questions.
- Submit a Pull-Request or two! We welcome all PR(s) that improve the project!

Please see the [Contributing Guidelines](/CONTRIBUTING.md) and checkout the
[Developer Documentation](https://dev.twtxt.net) or over at [/docs](/docs).

## Contributors

Thank you to all those that have contributed to this project, battle-tested it,
used it in their own projects or products, fixed bugs, improved performance
and even fix tiny typos in documentation! Thank you and keep contributing!

You can find an [AUTHORS](/AUTHORS) file where we keep a list of contributors
to the project. If you contribute a PR please consider adding your name there.

## License

`yarn` is licensed under the terms of the [MIT License](./LICENSE)
