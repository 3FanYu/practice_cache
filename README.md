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


# Design Road Map
## Functinal Requirements
1. Registration
2. User sign in
3. Email verification
4. View recommendation list

## Non Functional Requirements
1. 300RPM == 5RPS
2. 3 sec DB query

## Caching 
As explained above, obviously 3 seconds of DB query will lead to bad UX. Therefore, cache implementaion is a must in this scenario.
After the cache is stored, it is going to greatly improve reponse time of the API. How about before the cache is stored? We've gotta initiate the cache at some point.
When the API is first used, its going to take more than 3 seconds since the DB qeury itself takes 3 seonds. And we are dealing with 5RPS, that means 15 requests within 3 seconds.
That is 15 concurrent requests all with missed cache and querying against DB!
To prevent that, I decided to implement SingleFlight(https://pkg.go.dev/golang.org/x/sync/singleflight) to have just 1 request to query against DB at once.

