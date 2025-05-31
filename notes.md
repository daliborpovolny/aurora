# To do

restructure to include folders like:
/internal -> /handlers, /servicer, /utils, /auth, /tests
/templates -> /partials, /layouts
/database -> /gen, queries.sql, schema.sql
/static
/config -> air.toml, ...

add openapi


# SQL FORMAT
- separate into individual files
- all tables are in the singular case


# Handlers
- public, private
- maybe json and html variants?
- modify handlers to return error
	- log the error
	- error should be a struct with a message and status code
	- return the error to the client in appropriate format - html / json 
