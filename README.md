# bouli.com.br "Hub" Websites

This repository is a simple generic webserver in GoLang for my domains.
Some websites are onepage websites, with a very simple html structure,
and other ones are just redirects to other external websites.

Dependencies
----

- Golang@1.24+

Install
----

```shell
brew install go
go run .
```

Deploy
----

The project is ready to be deployed in Heroku. So to deploy, you can follow the
script bellow based on [this tutorial](https://devcenter.heroku.com/articles/getting-started-with-go#prepare-the-app).

```shell
brew install heroku/brew/heroku
heroku login
heroku create
heroku push heroku main
heroku open
```

To install the domains are in [scripts/heroku-add-domains.sh](scripts/heroku-add-domains.sh).
Nowadays we have the domains bellow:
- bouli.com.br
- cesarbouli.com
- byeceebee.com
- salviasupernova.com.br
- paralelossa.com.br
- littletalks.org
