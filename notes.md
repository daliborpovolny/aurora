# To do

- restructure to include folders like:
	- /config -> air.toml, ...

- utils.Decode crashes if given not a pointer - TODO WHOLE REWRITE MUST UNDERSTAND THE CODE THIS TIME !!!!
- add openapi
- change handlers to use auth service and include authinfo in publicDeps, require authInfo in privateDeps
- create custom something (handlers?) that require specific user type
- add tests - a lot of them - crud operations on user, teacher, student, admin, parent
- add AuthInfo to every template -> conditional navbar based on userType, my account page


# SQL FORMAT
- separate queries into individual files
- all tables are in the singular case


# DOCKER
- dockirize the app

# OPENAPI
- include openapi documentation