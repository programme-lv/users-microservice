```
aws-lambda-users-microservice/
├── cmd/                         # Main applications for this project
│   └── lambda/
│       └── main.go              # Entry point for the Lambda function
├── config/                      # Configuration files
│   └── config.toml              # Configuration file for the service
├── internal/                    # Private application and library code
│   ├── handlers/                # Lambda handlers
│   │   ├── create_user.go
│   │   ├── get_user.go
│   │   ├── update_user.go
│   │   └── delete_user.go
│   ├── entities/                # Domain models with business logic
│   │   └── user.go
│   ├── repository/              # Persistence adapters and data access layer interfaces
│   │   └── dynamodb.go
│   ├── services/                # Use cases
│   │   └── user_service.go
│   ├── interfaces.go            # Interfaces for communication between subpackages
├── scripts/                     # Scripts for various tasks
│   └── deploy.sh                # Deployment script
├── .gitignore                   # Git ignore file
├── go.mod                       # Go module file
├── go.sum                       # Go dependencies file
└── README.md                    # Project documentation
```


## services/

The software in this layer contains application specific business rules.
It encapsulates and implements all of the use cases of the system.
These use cases orchestrate the flow of data to and from the entities,
and direct those entities to use their enterprise wide business rules to achieve the goals of the use case.

## handlers/

The software in this layer is a set of adapters that convert data from the format most convenient for the use cases and entities,
to the format most convenient for some external agency such as the Database or the Web.
It is this layer, for example, that will wholly contain the MVC architecture of a GUI.
The Presenters, Views, and Controllers all belong in here.
The models are likely just data structures that are passed from the controllers to the use cases,
and then back from the use cases to the presenters and views.

## entities/

Entities encapsulate Enterprise wide business rules.
An entity can be an object with methods,
or it can be a set of data structures and functions.
It doesn’t matter so long as the entities could be used by many different applications in the enterprise.