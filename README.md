# smex

A blazing fast CLI application that processes sitemaps in Go Lang. 

smex is short for **S**ite**M**ap **EX**trator. It can handle various operations related to sitemaps with more being build over time.

smex supports: 
- extraction of urls
- conversion of sitemap into other formats such as json and csv
- url status checking

## Usage

### Extract

To extract information from sitemap

The following command extracts yoast's post sitemaps and prints to stdout
`smex extract https://yoast.com/post-sitemap.xml --remote --loc`

You can also point the extract at a sitemap locally
`smex extract ~/Download/sitemap.xml`

### Convert

### Check

### Validate 

TODO

### Help

To get help simply run `smex` without any commands and flags.

## Installation

### Using Go

`go get -u github.com/hbish/smex`

### Binary

TODO

## Documentation

For package documentation please check on [pkg.go.dev](https://pkg.go.dev/github.com/hbish/smex).

## Contribute

### Getting the code

Clone the repo

`git clone git@github.com:hbish/smex.git`

Initialise local environment

`make init`

Running tests

`make test`

Running the source

`go run github.com/hbish/smex [command] [flags...]`

