CREATE TABLE users 
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR (255) NOT NULL
);

CREATE TABLE notebook_lists 
(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE users_lists 
(
    id SERIAL NOT NULL UNIQUE,
    user_id INT REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES notebook_lists(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE notebooks_items
(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done boolean NOT NULL default false
);

CREATE TABLE lists_items
(
    id SERIAL NOT NULL UNIQUE,
    item_id INT REFERENCES notebooks_items (id) ON DELETE CASCADE NOT NULL,
    list_id INT REFERENCES notebook_lists (id) ON DELETE CASCADE NOT NULL 
);