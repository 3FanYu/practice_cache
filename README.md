High level design

POST /v1/user/register

POST /v1/user/sign_in

GET /v1/user/item/recommendation

SCHEMA
users
- id
- email
- encrypted_password
- verify_token

items
- id
- name
- quantity
- category

emails
- id
- target_address
- validation_link