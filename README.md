# Lab04 Assignment
## How to Run GO Assignment

---
## Install Go and Git
### Dowload if you don't have GO : https://go.dev/dl/

check that you have installed GO by command : 
```bash
go version
 ```
>`Output should be : go version go...`

### Dowload if you don't have Git : https://git-scm.com/install/

check that you have installed Git by command : 
```bash
git version
 ```

>`Output should be : git version...`

Go to link Github Repository 
---
> https://github.com/puttipong-6609650541/go-api-gin-lab.git 
---

## Clone the Repository into your workspace (Folder must be selected before!!!)
```bash
git clone https://github.com/puttipong-6609650541/go-api-gin-lab.git 
 ```
---
## Change Directory to the repository that you've clone
```bash
cd go-api-gin-lab
```
## Run GO
```bash
go run main.go
```
---
## Expected Result after run the command
>`Listening and serving HTTP on :8080 `

Server will run in : http://localhost:8080 , You can test the API in `Postman`
---

# Test API Section

### GET /students
Retrieve all students

```
GET http://localhost:8080/students
```

**Response 200 OK:**
```json
[
  { "id": "66090001", "name": "John Doe", "major": "Computer Science", "gpa": 3.75 }
]
```
**OR If you have no student in database (still 200 OK):**
```json
null
```

---

### GET /students/:id
Retrieve a student by ID

```
GET http://localhost:8080/students/66090001
```

**Response 200 OK:**
```json
{ "id": "66090001", "name": "John Doe", "major": "Computer Science", "gpa": 3.75 }
```

**Response 404 Not Found:**
```json
{ "error": "Student not found" }
```

---

### POST /students
Create a new student

```
POST http://localhost:8080/students
In Postman choose raw in Body field
```

**Request Body:**
```json
{
  "id": "66090001",
  "name": "John Doe",
  "major": "Computer Science",
  "gpa": 3.75
}
```

**Response 201 Created:**
```json
{ "id": "66090001", "name": "John Doe", "major": "Computer Science", "gpa": 3.75 }
```

**Response 400 Bad Request:**
```json
{ "error": "name must not be empty" }
```

---

### PUT /students/:id
Update a student by ID

```
PUT http://localhost:8080/students/66090001
In Postman choose raw in Body field
```

**Request Body:**
```json
{
  "name": "John Updated",
  "major": "Data Science",
  "gpa": 3.90
}
```

**Response 200 OK:**
```json
{ "id": "66090001", "name": "John Updated", "major": "Data Science", "gpa": 3.90 }
```

**Response 404 Not Found:**
```json
{ "error": "Student not found" }
```

---

### DELETE /students/:id
Delete a student by ID

```
DELETE http://localhost:8080/students/66090001
```

**Response 204 No Content**  (no body)

**Response 404 Not Found:**
```json
{ "error": "Student not found" }
```

---

## Validation Rules

| Field | Rule |
|-------|------|
| `id` | Must not be empty (POST only) |
| `name` | Must not be empty |
| `gpa` | Must be between 0.00 and 4.00 |
