# Go Rest Service

#### API Spec
- The OpenAPI spec file which describes the reset services is in resource folder (RecipeService.yaml).
  The file can be used with any spec reader to get better view but for a developer
  it should be easy to understand
- Due to the lack of the time the spec can't be used right now to generate test clients
    or integrate the spec UI into application
- The Spec sheet should serve as extensive documentation to use the REST services provided

#### Application

- The app contains modules which represent the concern associated with the code.
- The code also contains inline documentation to help in understanding the role of the particular section
- Due to the limited time I have tried to test the application with functional tests rather tried to skip 
    the unit test for as the middleware of the application is too thin and after mocking will leave with no code to 
    assert
- For testing purpose the BDD test suites are used as they clearly help in representing the validity of the
    system in terms of HAVING -> GIVEN -> EXPECTED
    
    
#### Docker:
- The application is integrated with docker for easy testing and deployments
- Right now the docker image of the application is not published
- In `docker-compose.yml` file under docker folder has the `server` container and `mongo` container for the application to work
- If the parameters need to be changed they can be changed in docker compose file.
    
    
#### Config
 - Sample Config can be found in resources as config.json
 - Environment specific config can be passed to the app using environment variable
    `$ export CONFIG=prod.json`
    Note: The Environment variable is compulsory for the app to work
    
#### Testing
- All the test are placed under one package for the ease of CI,CD execution and allows us to respect the encapsulation of the service packages
- Goto the `functional-tests` directory under root project directory and execute below command.
- Note: The docker containers / spplication should be started before starting the tests as the test are run against the localhost deployed version of application
- Please change the url for the tests if executing againgst real deployed versions or docker with DNS reachability
 `$ go test`

#### Run
- The service needs an environment variable which specifies the port on which to start the servie
    `$ export PORT=9000`
- From the service root folder `recipe-service`
 `$ sh ../scripts/start-env.sh`
- The command will start the necessary container with the server running on  port specified
- To change the port for the service please provide the `PORT` env variable to the container.
- The startup script will output logs from the container to the current bash session for analysis

#### STOP 
- To stop the containers the startup script has multiple options to get going
- Execute the startup script with `down` option

#### Logs
 - Logs can be found under `application.log` in the go executable path
 - At the same time the logs are output for you when the containers start
 - The startup script also help to accumulate all the logs from mongodb and server container into `docker-compose.log`.
 - All the composite logs from all containers can be found under `docker` directory
 - Degug level for logging debug info can be controlled using environment variable DEBUG
  `$ export DEBUG=True`
  
#### Authentication
- The authentication approach used is JOSE:Javascript Object Signing and Encryption set of standards
- This is the implementation for providing a simple implementation of Auth0 spec also know as JOSE Javascript Object Signing and Encryption set of standards
- The signing is done using static unprotected key using HMAC256
- We are using a single user/pass for authentication for now. 
    Usually, we will want to use a third part IDP(Identity) provider like google 
    to get the IDToken and then based on that give AccessToken to the user based on the 
    IDToken verification using JWK
- The endpoint for authentication is present at `/auth`
- The specs of the auth endpoint can be found in Open API Spec sheet RecipeService.yaml`
- The acquired token should be passed with the subsequent request as header `Authorization`
    `Authorization -> Bearer XXX.YYYY.ZZZ`
