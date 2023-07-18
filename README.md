# instruction to run the application

This repository is for creating an API using the Go programming language with MySQL database, running on a local server using XAMPP. 
In this implementation, I used GORM and Gin for developing the API.

# To run this application, please follow these steps:
1. Start your local server by running XAMPP and enabling Apache and MySQL services.
2. Open this repository in your preferred text editor.
3. Open the terminal in your text editor and run the command "go run main.go".
4. The application will start running, and you can test it using API testing tools like Postman, Netlify, or others.
5. Here are the endpoints of the API:
   1. Farm
   - (CREATE) http://localhost:8080/farm/create
   - (PUT) http://localhost:8080/farm/:id
   - (DELETE) http://localhost:8080/farm/:id
   - (GET) http://localhost:8080/farm/
   - (GET) http://localhost:8080/farm/:id
  
   2. Pounds
   - (CREATE) http://localhost:8080/pounds/create
   - (PUT) http://localhost:8080/pounds/:id
   - (DELETE) http://localhost:8080/pounds/:id
   - (GET) http://localhost:8080/pounds/
   - (GET) http://localhost:8080/pounds/:id
  
   3. Statistics
   - (GET) http://localhost:8080/statistics/

Thank you for the opportunity. If there are any mistakes, please accept my apologies.
