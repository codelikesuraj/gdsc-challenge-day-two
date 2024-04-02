# Day 2: Returning a message

Now that you have your HTTP server, you will create an endpoint, /hello. This endpoint should return the Hello GDSC in bold. [Hint: Use HTML tags for this].
Another endpoint, /api/hello should return the same thing but this time in JSON form with the message property set to “Hello GDSC”.

## Setup
- Navigate to the root of this repo.
- Run the command ```go run ./main.go``` to start the server.
- Visit the url endpoints:
    - http://127.0.0.1:3000/hello
    - http://127.0.0.1:3000/api/hello