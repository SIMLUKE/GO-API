### API DOC
>[!IMPORTANT]
>EVERY THING IS DOCKER BASED
>```
>docker compose up --build
>```

**The api listens on : http://localhost:8080**

## The ROUTES for the api are:
# GET
- GET &emsp; `/username`:
    + DESCRIPTION:
        + Input a valid user's uuid to get user's name
    + INPUT:
        ```json
        {
            "Id":"user's uuid"
        }
        ```
    + SUCCESS OUTPUT: -- code: 302 (found)
        ```json
        {
            "Name":"user's name"
        }
        ```
    + FAILLURE OUTPUT:
        + User not found -- code: 404 (not found)

- GET &emsp; `/article`:
    + DESCRIPTION:
        + Input an article's uuid, get it's information
    + INPUT:
        ```json
        {
            "Id":"article's unique uuid"
        }
        ```
    + SUCCESS OUTPUT: -- code: 302 (found)
        ```json
        {
            "Id": "article's unique uuid",
            "Name": "article's name",
            "Size": "article's size",
            "Price": article's price (float32),
            "State": "article's state",
            "Brand": "article's brand",
            "Color": "article's color",
            "UserId": "article's owner's uuid"
        }
        ```
    + FAILLURE OUTPUT:
        + Article not found -- code: 404 (not found)

- GET &emsp; `/articles`:
    + DESCRIPTION:
        + No input, get array of all the current articles
    + INPUT:
      ```json
        {} (no input required)
      ```
    + SUCCESS OUTPUT: -- code 200 (ok)
        ```json
        [
            {
                "Id": "article's unique uuid",
                "Name": "article's name",
                "Size": "article's size",
                "Price": article's price (float32),
                "State": "article's state",
                "Brand": "article's brand",
                "Color": "article's color",
                "UserId": "article's owner's uuid"
            },
            {
                "Id": "article's unique uuid",
                "Name": "article's name",
                "Size": "article's size",
                "Price": article's price (float32),
                "State": "article's state",
                "Brand": "article's brand",
                "Color": "article's color",
                "UserId": "article's owner's uuid"
            }
        ]
        ```
    + FAILLURE OUTPUT:
        + pure function, should not be any failures.


# POST

+ POST &emsp; `/user_register`:
    + DESCRIPTION:
      + Input a valid name, mail and password to add a new user to the database
    + INPUT:
        ```json
        {
            "Name":"user's name"
            "Email":"usermail@user.com",
            "Password":"encrypted password"
        }
        ```
    + SUCCESS OUTPUT: -- code 202 (accepted)
        ```json
        {} no output, code 
        ```
    + FAILLURE OUTPUT:
        + Email already exists -- code: 406 (not acceptable)
        + Invalid name or password -- code: 406 (not acceptable) 

+ POST &emsp; `/newarticle`:
    + DESCRIPTION:
        + Input valid article data ta add a new article to the database
    + INPUT:
        ```json
        {
            "Id": "article's unique uuid",
            "Name": "article's name",
            "Size": "article's size",
            "Price": article's price (float32),
            "State": "article's state",
            "Brand": "article's brand",
            "Color": "article's color",
            "UserId": "article's owner's uuid"
        }
        ```
    + SUCCESS OUTPUT: -- code 202 (accepted)
        ```json
        {} (no output)
        ```
    + FAILLURE OUTPUT:
        + Invalid article -- code: 406 (not acceptable)
        + User not found -- code: 406 (not acceptable)

+ POST &emsp; `/user_login`:
    + DESCRIPTION:
        + Input a valid email and valid password to get user's information
    + INPUT:
        ```json
        {
            "Email":"usermail@user.com",
            "Password":"securepassword"
        }
        ```
    + SUCCESS OUTPUT: -- code: 202 (accepted)
        ```json
        {
            "Id":"user's uuid",
            "Name":"user's name"
            "Email":"usermail@user.com",
            "Password":"encrypted password"
        }
        ```
    + FAILLURE OUTPUT:
        + Email or password is invalid -- code: 404 (not found)
        + Email not found -- code: 404 (not found)
