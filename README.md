BaseApp
=======

#### A [Bootstrap](http://getbootstrap.com)-based web application built in [Go](http://golang.org) on top of the [Revel Web Framework](https://revel.github.io) ####

BaseApp is a web application that provides web application scaffolding with the following features out of the box:

* Basic pages (Home, About Us, Contact Us, etc)
* Account registration (including password strength meter and safe storage of passwords with [bcrypt](https://en.wikipedia.org/wiki/Bcrypt))
* Account confirmation via email (via configurable e-mail settings)
* Account recovery via email (via configurable e-mail settings)
* Account log in (by username or email address)
* Public user profiles (with [Gravatar](http://gravatar.com/) integration for profile photos and profile edit functionality)
* Public user posts (with full [Markdown](http://daringfireball.net/projects/markdown/) support and edit + delete functionality)
* Social Features (Profile Follower and Following Counts and Post Likes)
* People finder search (by username or full name)
* Auto-provisioned backend data store (SQLite3, MySQL, PostgreSQL)
* Full live client-side and static server-side form validation
* Full non-interactive and interactive testing framework (Test-driven development process)
* Responsive front-end design with [Bootstrap](http://getbootstrap.com) (looks good on mobile, tablet and desktop screen sizes)

### Contents ###

* [Screenshots](#screenshots)
* [Quick Start](#quick-start)
* [BaseApp Run Modes](#baseapp-run-modes)
* [Docker Usage](#docker-usage)
* [Feedback](#feedback)
* [License](#license)

### Screenshots ###

Example Home Page:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/01.HomePage.png" style="max-width: 100%"/>

User Registration Form:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/02.Register.png" style="max-width: 100%"/>

User Registration Form - Form Validation Failed:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/09.Register_Fail.png" style="max-width: 100%"/>

User Login Form:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/03.Login.png" style="max-width: 100%"/>

Display User Profile:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/04.Profile.png" style="max-width: 100%"/>

Followers Display:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/11.Followers.png" style="max-width: 100%"/>

Following Display:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/12.Following.png" style="max-width: 100%"/>

Edit User Profile Form:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/05.EditProfile.png" style="max-width: 100%"/>

New Post Form:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/06.NewPost.png" style="max-width: 100%"/>

Post Display:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/10.Post.png" style="max-width: 100%"/>

User Search:

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/13.Search.png" style="max-width: 100%"/>

Reset Password Form (if email settings are provided in `conf/app.conf`):

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/07.ResetPwd.png" style="max-width: 100%"/>

Interactive Test Suite (at `http://localhost:9000/@tests`)

<img src="https://github.com/mr-klaatu/baseapp/raw/master/screenshots/08.TestRunner.png" style="max-width: 100%"/>

### Quick Start ####

#### Prerequisites ####

You will need a [functioning Revel installation](https://revel.github.io/tutorial/gettingstarted.html) for BaseApp to work.

#### Installing BaseApp ####

To get [BaseApp](https://github.com/mr-klaatu/baseapp), run

    go get github.com/mr-klaatu/baseapp/...

This command does a couple things:

* Go uses git to clone the repository into $GOPATH/src/github.com/mr-klaatu/baseapp/
* Go transitively finds all of the dependencies and runs go get on them as well.

#### Configuring BaseApp ####

Before you can start using BaseApp you will need to add your own `app.conf` file in the `conf/` directory. You can copy and use the default configuration file provided at [`/conf/app.conf.default`](https://github.com/mr-klaatu/baseapp/blob/master/conf/app.conf.default) with

    cp $GOPATH/src/github.com/mr-klaatu/baseapp/conf/app.conf.default $GOPATH/src/github.com/mr-klaatu/baseapp/conf/app.conf

Note: It is highly recommended that you review the configuration options available in [`app.conf.default`](https://github.com/mr-klaatu/baseapp/blob/master/conf/app.conf.default) before you run your project for the first time!

You can customize the configuration of your BaseApp application either directly in your `app.conf` or via environment variables (see [Running BaseApp](#running-baseapp) below for instructions on running BaseApp with environment variable overrides).

| Configuration Name | Environment Variable Name | Description |
| --- | --- | --- |
| - | `BASEAPP_RUN_LEVEL` | The [run level](#baseapp-run-modes) of your application. |
| `app.name` | `BASEAPP_APPLICATION_NAME` | The name of your application.<br><br>_default: `"BaseApp"`_ |
| `http.host` | `BASEAPP_SERVER_HOST` | The fully qualified address for your application.<br><br>_default: `"http://localhost:9000"`_ |
| `http.port` | `BASEAPP_SERVER_PORT` | The port on which to expose your application.<br><br>_default: `9000`_ |
| `http.ssl` | `BASEAPP_SERVER_SSL` | Flag to indicate that BaseApp should be served over HTTPS<br><br>_default: `false`_ |
| `http.sslcert` | `BASEAPP_SERVER_SSLCERT` | Your SSL certificate as a string<br><br>_default: ``_ |
| `http.sslkey` | `BASEAPP_SERVER_SSLKEY` | Your SSL private key as a string<br><br>_default: ``_ |
| `db.driver` | `BASEAPP_DATABASE_DRIVER` | The Go driver name for your database. |
| `db.import` | `BASEAPP_DATABASE_IMPORT` | The Go driver library for your database |
| `db.spec` | `BASEAPP_DATABASE_SPEC` | The Go driver connection string for your database |
| `mailer.server` | `BASEAPP_MAILER_SERVER` | SMTP server address |
| `mailer.port` | `BASEAPP_MAILER_PORT` | SMTP server port |
| `mailer.username` | `BASEAPP_MAILER_USERNAME` | SMTP server user name |
| `mailer.password` | `BASEAPP_MAILER_PASSWORD` | SMTP server password |
| `mailer.fromaddress` | `BASEAPP_MAILER_FROMADDRESS` | The email address to include as the sender address |

#### Running BaseApp ####

Once you have set up an `app.conf` file you can start BaseApp [_in test mode_](#baseapp-run-modes), with

    revel run github.com/mr-klaatu/baseapp test

If you have set up [BaseApp environment variable overrides](#configuring-baseapp) you should start BaseApp with the included [`start.sh`](https://github.com/mr-klaatu/baseapp/blob/master/start.sh) script instead:

    ./start.sh

Point your browser to your BaseApp installation at `http://localhost:9000` (or the host/port you specified in your configuration) and BaseApp should compile and run.

### BaseApp Run Modes ###

BaseApp can be run in three different modes that are each useful for different stages of application development:

1. `test` mode: Uses an *in-memory* sqlite3 datastore that is created and populated with basic data which is always wiped when your application ends.

    You can run BaseApp in `test` mode as follows:

        $> revel run github.com/mr-klaatu/baseapp test 9000

    Once BaseApp is up and running you can point your browser to `http://localhost:9000` to use the application or `http://localhost:9000/@tests` to run the interactive test suite.

    To run the BaseApp test suite in non-interactive mode you can use:

        $> revel test github.com/mr-klaatu/baseapp test 9000

    If testing is successful then a `test-results/result.passed` is written. Otherwise a `test-results/result.failed` is written. You can use check for these files when testing before deployment within your own continuous integration system.

2. `dev` mode [default mode]: Uses a blank sqlite3/mysql/postgres datastore that is created if it does not already exist at your chosen location and persists whenever your application is restarted. Outputs detailed error messages if something goes wrong in your application.

    You can run BaseApp in `dev` mode as follows:

        $> revel run github.com/mr-klaatu/baseapp dev 9000

    Or, simply:

        $> revel run github.com/mr-klaatu/baseapp

3. `prod` mode: Uses a blank sqlite3/mysql/postgres datastore that is created if it does not already exist at your chosen location and persists whenever your application is restarted. Outputs user-friendly error messages if something goes wrong in your application.

    You can run BaseApp in `prod` mode as follows:

        $> revel run github.com/mr-klaatu/baseapp prod 9000

Note: Both `dev` and `prod` modes require a configured backend DB. See [app.conf.default](https://github.com/mr-klaatu/baseapp/blob/master/conf/app.conf.default). The `test` mode creates an in-memory database representation that dies when the app dies.

### Docker Usage ###

This project comes with its own `Dockerfile` for building your own LXC container via [Docker](http://www.docker.io) to host BaseApp.

You will need a [functioning Docker installation](http://docs.docker.io/installation/) to build and use BaseApp as an LXC container.

To build your own BaseApp LXC container via Docker is easy:

```
$ git clone https://github.com/mr-klaatu/baseapp.git
$ cd baseapp
$ docker build -t="mr-klaatu/baseapp" .
```

To spawn a new instance of BaseApp:

```bash
$ docker run -p 9000:9000 -d mr-klaatu/baseapp
```

Then you can the visit the following URL in a browser on your host machine to get started:

```
http://127.0.0.1:9000
```

#### Docker Compose ####

An example `docker-compose` configuration is provided with this repo. To build and run the infrastructure defined in the `docker-compose` file, run the following:

```
$ docker-compose up
```

It can take a while to bring this architecture up but once BaseApp has been built in its docker container, you can navigate to `http://dockerhost:80` (or `http://localhost:80`) to see everything running.

See the [Configuring BaseApp](#configuring-baseapp) section above for available docker compose environment variable overrides.

### Feedback ###

If you find any bugs or issues please report them on the [BaseApp Issue Tracker](https://github.com/mr-klaatu/baseapp/issues).

If you would like to contribute to this project please consider [forking this repo](https://github.com/mr-klaatu/baseapp/fork), making your changes and then creating a new [Pull Request](https://github.com/mr-klaatu/baseapp/pulls) back to the main code repository.

### License ###

MIT. Copyright (c) [Rich Tibbett](https://twitter.com/richtibbett).

See the [LICENSE](https://github.com/mr-klaatu/baseapp/blob/master/LICENSE) file.
