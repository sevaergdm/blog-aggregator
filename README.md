# Blog Aggregator

This project provides basic CLI called "gator", and gives a user the ability to follow RSS feeds, fetch their posts to store in a local postgres db and then display those posts in the terminal.

## Requirements
* Go
* Postgres

## Installation

Simply run `go install -i github.com/sevaergdm/blog-aggregator`

## Setup

You'll need to first create a config file and then set up your local postgres db.

### config

Create a configuration file `~/.gatorconfig.json` and set the following content: 

```
{
    "db_url": "postgres://<user name>@localhost:5432/gator?sslmode=disable",
}
```

### db setup

* Start your postgres server in the background (how to do this will vary depending on OS)
* Connect to your server by entering the `psql` shell (e.g. `psql postgres` if on Mac OS)
* Create the "gator" database `CREATE DATABSE gator;`

## Usage

Once installed, you should be able to execute the cli commands via `gator <command>`

Available commands are:
* login: logs in as a specifed user like `login <username>`
* register: adds a new user like `register <username>`
* reset: resets the database
* addfeed: adds the specified feed to the database and follows for the current user: `addfeed <url>`
* feeds: lists the available feeds in the database
* follow: adds the current user as a follower for the feed specified: `follow <url>`
* following: lists the feeds followed by the current user
* unfollow: removes the current user as a follower of the specified feed `unfollow <url>`
* browse: prints the posts from the feeds followed by the current user. Optionally include limit to the number of posts returned `browse <limit>`
* agg: fetches posts from the available feeds at a given duration value `agg <duration>`
