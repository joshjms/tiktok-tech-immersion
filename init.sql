CREATE TABLE messages
(
    id SERIAL PRIMARY KEY,
    chat VARCHAR(100) NOT NULL,
    "text" TEXT NOT NULL,
    sender VARCHAR(100) NOT NULL,
    send_time INTEGER NOT NULL
);