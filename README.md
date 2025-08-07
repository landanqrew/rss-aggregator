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


