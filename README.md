# GATOR - Go RSS Aggregator

GATOR is a command-line RSS feed aggregator built with Go. It allows you to manage RSS feeds, follow your favorite content creators, and browse their latest posts all from your terminal.

## Prerequisites

Before you can use GATOR, you'll need to have the following installed on your system:

- [Go](https://golang.org/doc/install) (version 1.23 or higher)
- [PostgreSQL](https://www.postgresql.org/download/) database server

## Installation

You can install the GATOR CLI directly using Go:

```bash
go install github.com/NeroQue/GATOR@latest
```

This will install the `GATOR` executable in your Go binary directory.

## Configuration

1. Create a `.env` file in the directory where you run GATOR with the following content:

```
DB_CONNECTION_STRING=postgres://username:password@localhost:5432/gator?sslmode=disable
```

Replace `username`, `password`, and potentially the database name `gator` with your PostgreSQL credentials.

2. The application will automatically create a `.gatorconfig.json` file in your home directory to store your current user session and configuration.

## Usage

GATOR provides several commands to manage your RSS feeds:

### User Management

```bash
# Register a new user
gator register <username>

# Log in as a user
gator login <username>

# List all users
gator users
```

### Feed Management

```bash
# Add a new RSS feed
gator addfeed <feed_name> <feed_url>

# List all available feeds
gator feeds

# Follow a feed
gator follow <feed_url>

# List feeds you're following
gator following

# Unfollow a feed
gator unfollow <feed_url>
```

### Content Browsing

```bash
# Browse the latest posts from feeds you follow
gator browse [limit]

# Aggregate and fetch the latest content from your feeds
gator agg <duration>
```

Note: The `duration` parameter for `agg` command specifies the time between fetches (e.g., `1m`, `5m`, `1h`).

### Maintenance

```bash
# Reset your configuration
gator reset
```

## Example Workflow

1. Register yourself as a user: `gator register johndoe`
2. Add an RSS feed: `gator addfeed "Example Blog" https://example.com/feed.xml`
3. Follow another feed (if needed): `gator follow https://another-example.com/feed.xml` (feeds you add are automatically followed)
4. Fetch the latest content: `gator agg 5m` (checks for new content every 5 minutes)
5. Browse your latest posts: `gator browse 10` (shows 10 most recent posts)


