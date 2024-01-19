# Github Issues Notifier 📢
A daemon process that listens for new issues in Github repositories and notifies about them via email.

I run this in a Raspberry Pi 4 🫐 and use it to monitor Open Source repos that I like to contribute to.

## Prerequisites
- `docker` (but not really, you could also compile the binary)
- `make` (but not really, you can just check the contents of the `Makefile` and use `docker` directly)
- `go 1.21+` (but not really, only if you decide not to use `docker`)

## Installation
This is intended to be used as a Docker container, although it could also be compiled and run as a standalone Go binary using something like `systemd`.

1. Clone this repository
``` bash
$ git clone https://github.com/yugarinn/github-issues-notifier
```

2. Change to the notifier folder and execute the run command
``` bash
$ cd github-issues-notifier
$ make run
```

It is also possible to:

3. Check the logs:

``` bash
$ make logs
```

4. Stop the service

``` bash
$ make stop
```

## Configuration
The notifier relies on a `yml` file to define which repositories it should listen to. This is an example configuration file:
``` yaml
listeners:
  - name: 'lila'
    repository: 'lichess-org/lila'
    email_to: 'address@mail.com'
    is_active: true
    filters:
      labels: 'good first issue'
      assignee: 'none'

  - name: 'doom'
    repository: 'doomemacs/doomemacs'
    email_to: 'address@mail.com'
    is_active: true
    filters:
      labels: 'good first issue,help wanted'

  - name: 'rails'
    repository: 'rails/rails'
    email_to: 'address@mail.com'
    is_active: true
    filters:
      labels: 'good first issue,help wanted'
```

Each listener is composed of:
- Its name.
- The repository it should listen to in the `owner/repository`  format.
- The email to send the notification to.
- An active flag.
- A filters list.

Only the issues that match the defined filters in each listener are taken into account. Also, keep in mind that just label filters are implemented.

### Environment variables
Since you need to **bring your own SMTP server**, the following variables MUST be defined in a `.env` file in the root folder of the project:
- `SMTP_HOST`
- `SMTP_PORT`
- `SMTP_EMAIL_FROM`
- `SMTP_PASSWORD`

The following variables are optional:
- `WORKER_CRON_FREQUENCY` defaults to `"*/30 * * * *"` (at every 30th minute)
- `LISTENERS_FILE_PATH` Where the configuration file defining the listeners is located. Defaults to `"./listeners.yml"`.
- `LISTENERS_DATABASE_PATH` Where the database file (this uses `bbolt` under the hood) is located. This database is needed jsut to keep track of the already notified issues, nothing fancy really. Defaults to `"./listener.db"`

## Contributing
Just open a PR!
