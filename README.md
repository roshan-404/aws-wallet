# AWS-Wallet
A service where users can register themselves and can create a personal space for uploading their files. 

###  How To get this working:
- You must have GoLang installed. 
- You will need to run go mod download to install dependencies for this project. Do that in the root directory of the project. 

## Starting - Local 
- Copy .env.example to .env in the same directory 
- Remember to create your own .env based on the template provided. 

## AWS configuration 
- You will need to have a AWS account and then add the follwing in the .evn configuration file. 
1. Access key ID 
2. Secret access key 

## Run Command
```sh 
go run main.go 
```

## Documentation of API's
> Run the app, and browse to http://localhost:3000/swagger/index.html. Find the Swagger 2.0 Api documents. 


## swagger documentation
documentation of API's are available here https://go-aws.herokuapp.com/swagger/index.html
