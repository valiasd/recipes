# recipes
Website that holds recipes for each user. It supports adding comments and ratings for each recipe. Due to lack of time I did not implement the frontend for these two features though. Registering, log in, create/list/read/delete of recipes works.

### how to run
`docker-compose up` and navigate to localhost

### Generic information
Swagger documentation: http://localhost:8080/swagger/index.html or [here](backend/docs/swagger.yaml)

Github link: https://github.com/valiasd/recipes


Improvements needed:
* Implement cookies so that the solution is more robust
* Implement adding comments and ratings on the frontend
* Implement users seeing more recipes than just their own