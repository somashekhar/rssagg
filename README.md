# rssagg
RSS Aggregator


# Using goose
1. Install goose
2. Create schema in sql/schema directory and cd to it.
3. Run goose up or down based on usage scenario
```
    cd sql/schema
    goose postgres postgres://postgres:password@localhost:5432/rssagg up
    goose postgres postgres://postgres:password@localhost:5432/rssagg down
```

# Using sqlc
1. Install sqlc.
2. Write the sqlc.yaml
3. Write the queries in sql/queries
4. Run it from root of project
```
    sqlc generate
```