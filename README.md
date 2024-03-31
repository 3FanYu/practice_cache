# How to use
## Setup
1. Make sure docker & docker-compose are installed.
2. Clone the repo and cd into the project directoy.
3. Run `make up`
4. Ensure containers are up and running using `docker-compose ps`
5. Run `make migrate` to run migrations

## Use the APIs 
1. Register
   - POST `/v1/user/register`
  - payload ```"email": "your@email.com", "password": "Testtest!"```
3. Verify Email
   - I did not implement real Email feature, but I've made an mock email API!
   - GET `/v1/emails?email=your@email.com`
   - You should see a result with `VerifyLink` value, enter the link via a browser to verify the email.
4. Sign In
     - POST /v1/user/sign_in
     - payload ```"email": "your@email.com", "password": "Testtest!"```
     - You will find an authorization token in the response header after sign_in. Copy it for later use!
5. Finally, view item recommendations!
   - GET `/v1/items/recommendations`
   - Remember to bring a header of `Authorization: Bearer {YourToken}`
