
# User Management gRPC Service 
This project implements a simple gRPC service in Go for managing user details, including fetching user details by ID, retrieving a list of users by IDs, and searching for users based on specific criteria.  
## Prerequisites  
- Go 1.16+  
- Protobuf Compiler (`protoc`)  

- `protoc-gen-go` and `protoc-gen-go-grpc` plugins

## Project Structure  

- `pb/`: Contains the protobuf definition file. 
- `pkg/service/`: Contains the gRPC server implementation.  
- `cmd/main.go`: Entry point for running the gRPC server.  

- `client/client.go`: Example client to test the gRPC service. 

## Getting Started  
### Clone the Repository  
 ```sh
   git clone https://github.com/abdullahnettoor/totality-grpc-user-service  
   cd totality-grpc-user-service 
 ```

### Install Dependencies
Run  `go mod tidy` to install dependencies used in the project.

### Specify Port
Specify the port if necessary (default `:9000`) in `dev.env` in root folder.

### Build the Application

Build the gRPC server for user service and client.

```sh 
go build -o user_server cmd/main.go   
go build -o client_instance cmd/main.go   
```

### Run the Application

#### Start the gRPC Server

```sh
./user_server   
```
The server will start on localhost:9000.

#### Run the gRPC Client
In a separate terminal, run the gRPC client.
```sh
./client_instance   
```

Created a Data Simulation for User List
---------------------------------

*   **Initialization**: `f := gofakeit.New(0)` initializes the **GoFakeIt package** with a seed value of 0 to ensure consistent random data generation across runs.
    
*   **User List Generation**: `userList := make([]*pb.User, count)` creates a slice of count length to store pb.User structs.
    
*   **User Data Generation**: Inside the loop for `i := range userList`, each user is generated with:
    
    *   **Id:** Sequentially assigned starting from 1.
        
    *   **Fname:** Random first name generated by `f.FirstName()`.
        
    *   **City:** Random city name generated by `f.City()`.
        
    *   **Phone:** Random phone number generated by `f.Phone()`.
        
    *   **Height:** Random height between 4.5 and 6.5 rounded to one decimal place.
        
    *   **IsMarried:** Random boolean value generated by `f.Bool()`.
        
*   **Output**: Each user's details are printed in a formatted table-like structure using `fmt.Printf()`.
    
*   **Return**: The function returns `userList`, which contains all generated `pb.User` structs.


Access the gRPC Service Endpoints
---------------------------------

The gRPC service provides the following endpoints:

1.  **GetUserByID**: Fetches a user by ID.
    
    *   Request: `UserIDRequest{ id: 1 }`
        
    *   Response: `UserResponse { user: User { id: 1, fname: "Steve", ... } }`
        
2.  **GetUserListByIDs**: Retrieves a list of users by IDs.
    
    *   Request: `UserIDsRequest { ids: [1, 2] }`
        
    *   Response: `UsersResponse { userList: [User { id: 1, ... }, User { id: 2, ... }] }`
        
3.  **SearchUsers**: Searches for users based on criteria *(e.g., fname, city, phone, marital status)*.
    
    *   Request: `SearchUserReq { city: "LA", isMarried: true }`
        
    *   Response: `SearchUserRes { userList: [User { id: 1, city: "LA", ... }] }`
        

Containeraisation using Docker
----------------
To run the application in a container, follow these steps: 
- Generate docker image from the `Dockerfile` provided in the repository.
- Run the `docker build -t grpc-user-service .` command to create docker image. 
- Run the container using the command `docker run -d -p  9000:9000 grpc-user-service`
- The application will be available at `localhost:9000`

Configuration Details
---------------------

*   **gRPC Server Address**: The server listens on `localhost:9000` by default. You can change the port in the `dev.env` file.
    
*   **Protobuf Definition**: The protobuf definition file `pb/userservice.proto` defines the messages and services for the gRPC service. If you  update proto file, Generate the Go code from the protobuf definition. Using command `protoc --go_out=. --go-grpc_out=. pb/*.proto`.

*   **Generated Code**: The generated Go code from the protobuf definition is located in the `pb/` directory.

---

***Acknowledgment:** This assignment is a Coding Assessment given by Totality Corp for the purpose of a job opening. This project is completely created from scratch.*