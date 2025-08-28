# Design

## 1. Architecture

The Media Tracker will be a web-based application with a client–server model.

- **Frontend:** Plain JavaScript and HTML (simple, no build tools needed).
- **Backend:** Go using Gorilla/Mux for routing.
- **Database:** SQLite3 (lightweight and easy to set up).
- **Authentication:** JSON Web Token (JWT) for session management.

## 2. Database Model

### users

| Column        | Type     | Constraints               |
| ------------- | -------- | ------------------------- |
| id            | INTEGER  | PRIMARY KEY               |
| username      | TEXT     | UNIQUE, NOT NULL          |
| password_hash | TEXT     | NOT NULL                  |
| created_at    | DATETIME | DEFAULT CURRENT_TIMESTAMP |

### media_items

| Column      | Type    | Constraints                                      |
| ----------- | ------- | ------------------------------------------------ |
| id          | INTEGER | PRIMARY KEY                                      |
| title       | TEXT    | NOT NULL                                         |
| type        | TEXT    | CHECK(type IN ('TV_SHOW','MOVIE','BOOK','GAME')) |
| source      | TEXT    | CHECK(source IN ('USER_ADDED','EXTERNAL'))       |
| external_id | INTEGER | NULLABLE                                         |
The reason for the source column is to make future proof by allowing ability for external sources.

### user_media

| Column   | Type    | Constraints                                          |
| -------- | ------- | ---------------------------------------------------- |
| id       | INTEGER | PRIMARY KEY                                          |
| user_id  | INTEGER | NOT NULL, FOREIGN KEY REFERENCES users(id)           |
| media_id | INTEGER | NOT NULL, FOREIGN KEY REFERENCES media_items(id)     |
| status   | TEXT    | CHECK(status IN ('TO_DO','IN_PROGRESS','COMPLETED')) |
| note     | TEXT    | Optional                                             |

## 3. API Design

### Authentication

`POST /auth/register` - Create a new account.
`POST /auth/login` - Retrieve a JWT token.

### Media items

These endpoints manage the shared pool of media items that all users can reference.

`GET /media` - Retrieve a list of all media.
`GET /media/{id}` - Retrieve details of a specific media item.
`POST /media` - Create a new media entry.
`PUT /media/{id}` - Update a specific media item.
`DELETE /media/{id}` - Delete a specific media item.

### User Media items

These endpoints let a user track their relationship with a media item.

`GET /user/media` - List all relationships for current user.
`GET /user/media/{id}` - Retrieve details of a specific media relationship for current user.
`POST /user/media` - Create a new relationship.
`PUT /user/media/{id}` - Update a specific relationship.
`DELETE /user/media/{id}` - Delete a specific relationship.