Golang Notes App API

DB Schema:

Notes:
    Title
    Body
    Date_created
    Date_modified
    User_id

User:
    Email
    Password
    Name

NOTES APIs

POST:  
    User creates a note by entering title and body of note. 
    /note
    Req Body: title, body

PATCH: 
    User updates title and/or body of note
    /note/:id
    Req Body: title, body (both optional)

DELETE: 
    User deletes note via note id
    /note/:id

GET: 
    User reads all notes
    /notes

GET(single): 
    User reads a single note
    /note/:id


User APIs

SIGNUP:
    Method: POST
    /singup
    Req Body: email, password, name


LOGIN:
    Method: POST
    /login
    Req Body: email, password
