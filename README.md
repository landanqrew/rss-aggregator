# rss-aggregator
An RSS Aggregator build with Go and Postgres

## Lessons learned
#### Local Data Stores for Desktop based tools / CLIs
* ~/.toolName/your-sub-directories-or-files

#### Integrating with Postgres locally
* postgresql
 * connection string: postgres://username:@localhost:port/db_name
 * `psql "<connection-string>"` to start your local instance
* sqlc
 * `sqlc generate` to generate code logic that coincides with specific sql commands
 * sqlc needs a slightly different version of the <connection-string>:
  * protocol://username:password@host:port/database?sslmode=disable
 * requires a sqlc.yaml file that stands as a config file
* goose
 * `goose postgres <connection_string> up` to perform an up migration (your terminal needs to be in the context of your /schema/ directory. see the ./gooseUp.sh file)

 ## Program Dependencies
 * Ensure you have Go installed locally on your device
 * Ensure you have postgresql installed locally on your device
 * start your local instance using the connection string directions above (default port is 5432)
 * run `go install .` once you have cloned the repo


## CLI Commands Overview

Here's a breakdown of the available commands in the RSS Aggregator CLI:

*   **`register <username>`**:
    *   **Purpose**: Registers a new user with the given username. If the user already exists, it will indicate that.
    *   **Arguments**: `username` (string) - The desired username.

*   **`login <username>`**:
    *   **Purpose**: Sets the current active user for the CLI. This user will be used for subsequent commands that require a logged-in user.
    *   **Arguments**: `username` (string) - The username to log in as.

*   **`users`**:
    *   **Purpose**: Lists all registered users in the system, indicating the currently logged-in user.
    *   **Arguments**: None.

*   **`addFeed <name> <url>`**:
    *   **Purpose**: Adds a new RSS feed to the system and automatically creates a feed follow for the currently logged-in user.
    *   **Arguments**: 
        *   `name` (string) - A display name for the feed.
        *   `url` (string) - The URL of the RSS feed.

*   **`feeds`**:
    *   **Purpose**: Lists all RSS feeds that the current user is following.
    *   **Arguments**: None.

*   **`follow <feed_url>`**:
    *   **Purpose**: Allows the currently logged-in user to follow an existing RSS feed by its URL.
    *   **Arguments**: `feed_url` (string) - The URL of the feed to follow.

*   **`following`**:
    *   **Purpose**: Displays all the feeds that the currently logged-in user is following.
    *   **Arguments**: None.

*   **`unfollow <feed_url>`**:
    *   **Purpose**: Allows the currently logged-in user to unfollow a specific RSS feed by its URL.
    *   **Arguments**: `feed_url` (string) - The URL of the feed to unfollow.

*   **`browse [max_results]`**:
    *   **Purpose**: Displays a list of recent posts from the feeds the current user is following. Optionally limits the number of results.
    *   **Arguments**: `max_results` (integer, optional) - The maximum number of posts to display. Defaults to 2.

*   **`agg <time_between_requests>`**:
    *   **Purpose**: Starts an aggregation process that periodically fetches new posts from the followed RSS feeds.
    *   **Arguments**: `time_between_requests` (duration string) - The time interval between fetch requests (e.g., "5s", "1m", "1h").

*   **`reset`**:
    *   **Purpose**: *Danger Zone!* Removes all user data from the database. Use with caution.
    *   **Arguments**: None.


