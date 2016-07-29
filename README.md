# FizzBuzz Rails

Code for a coding challenge.

Status: Compelete with some rough edges.

I wrote the functionality of calculating pages of fizzbuzz results independant of the Rails app. These classes have been put into the services folder.

### Highlights

* Fizzbuzz is only calculated for the numbers that appear on a page
* Only favourited numbers are stored in the database
* A single query to the database for viewing a page of results
* 100% test coverage for server
* I didn't know any Go before starting this

### Places for improvement

* more explicit testing of edge cases and behaviour
* page scrolls with favouriting in web application
* numbers are stored as bigint, so there is a limit (very high) to the highest favourite
* deprecation warnings in specs (rspec matchers yet to be updated to support redirect_back from Rails 5)
* go-client has compiled to 9mb
* api can be improved so client can take advantage of it
* tests for client

## Usage

### Server

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

## Client

Launch the prebuilt binary with
```
cd go-client
bin/fizzbuzz
```

Alternitavely run from the source code with
```
cd go-client
echo GOPATH=`pwd`
cd src/fizzbizz-client-server/fizzbuzz
go run main.go
```

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
