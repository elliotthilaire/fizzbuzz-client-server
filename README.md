# FizzBuzz Rails

Status: Compelete with some rough edges.

Code for a coding challenge. See below for notes.

## The challenge

### Task A: Ruby on Rails FizzBuzz application
- Standard FizzBuzz rules, divisible by 3 is Fizz, divisible by 5 is Buzz
- Should display values from 1 to 100 on the homepage
- Should allow viewing values up to 100,000,000,000
- Should have pagination
- Should allow changing of the page size
- User should be able to mark certain numbers as their favourites, these should indicate that they are favourites on the UI and be persisted
- Should provide a JSON API with all the above-mentioned functionality

### Task B: Client to consume the JSON API
- Create a client to consume the FizzBuzz application API
- Should be available from the command line
- Should be written in a language other than Ruby

### Notes:
- Please take your time to perform the test.
- We really value best practice and clean code.
- Test your Rails app using Rspec with a decent coverage.
- Use a public Github/Bitbucket account to publish your result.
- Commit early and often with clear messages.

## The server

Launch with
```bash
cd rails-server
bin/setup
rails server
```

Run specs with 
```
cd rails-server
rspec --format doc
```

## The client

Launch the prebuilt binary with 
```
cd go-client
bin/fizzbuzz 
```

Alternitavely run from the source code with
```
cd go-client
cd src/fizzbizz-client-server/fizzbuzz
go run main.go
```
