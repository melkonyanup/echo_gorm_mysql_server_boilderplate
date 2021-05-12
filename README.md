In order to run the project, please do the following:

Create .env file in the root folder of the project (see .env.example file)  
Configure config file (config/config.yml)  
Intall mysql on your os  
Create database with the name that you specified in the config  
Run "make build" (this will get all the dependencies for the your project)  
Run "make docs" (this will generate docs for the project)  
Run db migrations run: "make db.migrate"  
To start the server run: "make start"  

For unit tests run: "make test"  
For integration tests run: "make test.integration"  