**create table query**

CREATE TABLE `students` (
   `id` bigint NOT NULL AUTO_INCREMENT,
   `fname` varchar(50) NOT NULL,
   `lname` varchar(50) NOT NULL,
   `date_of_birth` datetime NOT NULL,
   `email` varchar(50) NOT NULL,
   `address` varchar(50) NOT NULL,
   `gender` varchar(50) NOT NULL,
   `createdBy` varchar(50) DEFAULT NULL,
   `createdOn` datetime DEFAULT NULL,
   `updatedBy` varchar(50) DEFAULT NULL,
   `updatedOn` datetime DEFAULT NULL,
   PRIMARY KEY (`id`)
 ) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci


**Config file path**

cmd\server\.env

**Log file path**

cmd\server\app.log

**Auth**

curl -X POST http://localhost:8080/auth -H "Content-Type: application/json" -d "{\"id\": \"teacher1\", \"password\": \"password1\"}"

**GET**

curl -v -X GET http://localhost:8080/api/v1/students -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNTE3NzgsImlhdCI6MTcyNDE0ODE3OCwidXNlcl9pZCI6InRlYWNoZXIxIn0.4BfL3syyDAkLTzFNm89A4uwjJLmrESZurc0Ln-HheIg"

**Get by id**

curl -v -X GET http://localhost:8080/api/v1/student/8 -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNTE3NzgsImlhdCI6MTcyNDE0ODE3OCwidXNlcl9pZCI6InRlYWNoZXIxIn0.4BfL3syyDAkLTzFNm89A4uwjJLmrESZurc0Ln-HheIg"

**Post**

curl -v -X POST http://localhost:8080/api/v1/student -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNTE3NzgsImlhdCI6MTcyNDE0ODE3OCwidXNlcl9pZCI6InRlYWNoZXIxIn0.4BfL3syyDAkLTzFNm89A4uwjJLmrESZurc0Ln-HheIg" -d "{\"fname\": \"Kamala\",\"lname\": \"Chan\",\"date_of_birth\": \"1994-06-23T23:51:42Z\",\"email\": \"bhavanilakshmegowda@gmail\",\"address\": \"Bangalore\",\"gender\": \"Female\"}"


**PUT**

curl -v -X PUT http://localhost:8080/api/v1/student/14 -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNzM0MjMsImlhdCI6MTcyNDE2OTgyMywidXNlcl9pZCI6InRlYWNoZXIxIn0.4VHUh3WwEpPJvfwZ14yCBsjJgiGy-ctQPkk_G1qfM9M" -d "{\"fname\": \"Benaka\",\"lname\": \"Bhavimane\",\"date_of_birth\": \"1994-06-23T23:51:42Z\",\"email\": \"bhavanilakshmegowda@gmail\",\"address\": \"Bangalore\",\"gender\": \"Female\"}"


**DELETE**

curl -v -X DELETE http://localhost:8080/api/v1/student/7 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNTE3NzgsImlhdCI6MTcyNDE0ODE3OCwidXNlcl9pZCI6InRlYWNoZXIxIn0.4BfL3syyDAkLTzFNm89A4uwjJLmrESZurc0Ln-HheIg"
