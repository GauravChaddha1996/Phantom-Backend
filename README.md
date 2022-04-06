# Phantom-Backend
A fashion e-commerce backend powered by a Golang using MySQL as database and Redis cache for querying and filtering.

The android repository is hosted ![here](https://github.com/GauravChaddha1996/Phantom-Android).

### Code-structure
The code is divided into following packages and files:
1. **apis**
    * Contains packages for home, filter & product endpoints. Each package holds a handler.go that is invoked for the respctive endpoint by the router
    * Contains apiCommons which holds easy logging for API error logs from all the layers from dbRead, cache misses to curation problems. 
2. **assets**
	  * Holds the images this backend supplies for all products
3. **config**
    * Holds the configuration for database and server
4. **dataLayer**
    * cacheDaos - holds the redis cache daos
    * databaseDaos - holds the mysql database daos
    * dbModels - holds the mysql database models
    * uiModels - holds the atomic models and snippets model this backend can send
5. **main.go**
  	*  The entry point to this monolith. It sets up the database, re-populates the redis cache with new values and starts the router.****
6. **ginRouter**
    * Actual router with it's middlewares lives here
7. **validator**
    * Holds the logic for custom validations on struct tags like productId, categoryId etc.

### API standards
Each API is divided into following
1. handler.go
	1. Handles the incoming request and transforms it into the request model (if needed)
	2. Creates the daos (cache and database as needed)
	3. Reads the data from DB using dbRead.go
	4. Invokes the logic in sections go files
	5. Finally forms the API response
2. models folder
	1. Holds different modles like API request, API response and it's sub-models.
  2. **Input validation** using struct tags is driven by a custom written validator
3. Sections folder
		1. Each API response is curated by converting different sections into snippets
		2. This folder holds each section logic separately, combining db read models with dao reads or previous section results.
4. dbRead.go
	1. Handles the initial read of all the data required from the database
	2. Uses one or more DAO’s for the db tasks
