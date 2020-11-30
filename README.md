# Payment Service

Requirements:
* [Go v1.15](https://golang.org/dl/)
* [Good mood*](https://www.youtube.com/watch?v=rcWynJROnyI)

## Run application locally
1. Clone repository 
   ```
    git clone https://github.com/andreybutko/payment.git
   ```
2. Navigate to project folder `/payment`
3. Copy `./configuration/config.json.sample` -> `./configuration/config.json`
4. Fill out `config.json` settings
5. Run api with
   ```
    go run main.go
   ```
6. Request examples available within Postman collection at `./postaman` folder.

## Tests
1. To run all tests use
    ```
     go test ./...
    ```
    *Currently tests available only for `UseCase`.  
    To run only them use:  
    ```
     go test github.com/andreybutko/payment/usecase/payment
    ```
    **Also, keep in mind that the project under active development, so there are a lot of TODO and coming improvements.  
    All suggestions, remarks are welcomed. You can fork or create an issue for it. There is no policy for PR/Issue, so it's at your discretion.
