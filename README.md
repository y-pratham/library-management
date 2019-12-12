# library-management
Creating rest apis using go language for library management

1. Download this code
2. In command prompt within the specified folder execute `$ go run main.go` 
3. Make sure you already have Gorilla Mux, else install it in the GOPATH using `$ go get -u github.com/gorilla/mux`

### REST APIs

|API Name       | Method        | JSON                                              | DESCRIPTION |
|:-------------:|:-------------:|:-------------------------------------------------:|:------------------------------------------------:
| /             | GET           |                                                   | Welcome User|
| /addBook      | POST          |{"UserID":"1234","BookID":"2","rating": 4.1,"Title":"book 2"} | Only the admin can add books **whose user id is 1234**|
|               |               |                                                   |
| /books        | GET           |                                                   | Gives a list of all books|
|               |               |                                                   | 
| /mostPopular  | GET           |                                                   |Gives the most popular book according to user ratings|
|               |               |                                                   |
| /mostIssued   | GET           |                                                   |Gives the most issued book |
|               |               |                                                   |
| /books/{id}   | GET           |                                                   |gives the details of a particular book of given ID|
|               |               |                                                   |
| /status/{id}  | GET           |                                                   |gives the status (available/ unavailable) for a book|
|               |               |                                                   |
| /ratings/{id} | GET           |                                                   |gives the ratings of a particular book |
|               |               |                                                   |
| /issue/{id}   | POST          |   {  "Description" : "user id" }                  |allows user (of a particular user id) to issue a book, it highlights book is unavailable if the book has already been issued|
|               |               |                                                   |
| /return/{id}  | POST          |   {  "Description" : "user id" }                  |Allows user to return a book|
|               |               |                                                   |
| /rate/{id}    | POST          |   {  "Description" : "rating" }                   |Allows user to rate a particular book|
|               |               |                                                   |

