**Auth**

curl -X POST http://localhost:8080/auth -H "Content-Type: application/json" -d "{\"id\": \"teacher1\", \"password\": \"password1\"}"

**GET**

curl -X GET http://localhost:8080/api/v1/students -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNTE3NzgsImlhdCI6MTcyNDE0ODE3OCwidXNlcl9pZCI6InRlYWNoZXIxIn0.4BfL3syyDAkLTzFNm89A4uwjJLmrESZurc0Ln-HheIg"

**Get by id**

curl -X GET http://localhost:8080/api/v1/student/8 -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNTE3NzgsImlhdCI6MTcyNDE0ODE3OCwidXNlcl9pZCI6InRlYWNoZXIxIn0.4BfL3syyDAkLTzFNm89A4uwjJLmrESZurc0Ln-HheIg"

**Post**

curl -X POST http://localhost:8080/api/v1/student -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNTE3NzgsImlhdCI6MTcyNDE0ODE3OCwidXNlcl9pZCI6InRlYWNoZXIxIn0.4BfL3syyDAkLTzFNm89A4uwjJLmrESZurc0Ln-HheIg" -d "{\"fname\": \"Kamala\",\"lname\": \"Chan\",\"date_of_birth\": \"1994-06-23T23:51:42Z\",\"email\": \"bhavanilakshmegowda@gmail\",\"address\": \"Bangalore\",\"gender\": \"Female\",\"CreatedBy\": \"\",\"CreatedOn\": \"\",\"UpdatedBy\": \"\",\"UpdatedOn\": \"\"}"


**PUT**

curl -X PUT http://localhost:8080/api/v1/student/8 -H "Content-Type: application/json" -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNTE3NzgsImlhdCI6MTcyNDE0ODE3OCwidXNlcl9pZCI6InRlYWNoZXIxIn0.4BfL3syyDAkLTzFNm89A4uwjJLmrESZurc0Ln-HheIg" -d "{\"fname\": \"sarala\",\"lname\": \"Chan\",\"date_of_birth\": \"1994-06-23T23:51:42Z\",\"email\": \"bhavanilakshmegowda@gmail\",\"address\": \"Bangalore\",\"gender\": \"Female\",\"CreatedBy\": \"\",\"CreatedOn\": \"\",\"UpdatedBy\": \"\",\"UpdatedOn\": \"\"}"


**DELETE**

curl -X DELETE http://localhost:8080/api/v1/student/7 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE3MjQxNTE3NzgsImlhdCI6MTcyNDE0ODE3OCwidXNlcl9pZCI6InRlYWNoZXIxIn0.4BfL3syyDAkLTzFNm89A4uwjJLmrESZurc0Ln-HheIg"
