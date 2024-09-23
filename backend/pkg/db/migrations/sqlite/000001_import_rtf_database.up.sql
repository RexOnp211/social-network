CREATE TABLE IF NOT EXISTS categories (
    category_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    category TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    nickname TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT DEFAULT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    dob DATE NOT NULL,
    aboutme TEXT NOT NULL       
);
CREATE TABLE IF NOT EXISTS posts (
    post_id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL,
    subject TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS comments (
    comment_id UUID PRIMARY KEY NOT NULL,
    post_id UUID NOT NULL,
    user_id UUID NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(post_id) REFERENCES posts(post_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS likes (
    like_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    post_id UUID,
    comment_id UUID,
    user_id UUID NOT NULL,
    type INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(post_id) REFERENCES posts(post_id),
    FOREIGN KEY(comment_id) REFERENCES comments(comment_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS post_categories (
    post_id UUID NOT NULL,
    category_id INTEGER NOT NULL,
    FOREIGN KEY(post_id) REFERENCES posts(post_id),
    FOREIGN KEY(category_id) REFERENCES categories(category_id)
);
CREATE TABLE IF NOT EXISTS sessions (
    session_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    token TEXT NOT NULL,
    user_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS messages (
    message_id UUID PRIMARY KEY NOT NULL,
    sender_id UUID NOT NULL,
    receiver_id UUID NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_read BOOLEAN NOT NULL DEFAULT 0,
    FOREIGN KEY(sender_id) REFERENCES users(user_id),
    FOREIGN KEY(receiver_id) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS user_status (
    user_id UUID PRIMARY KEY NOT NULL,
    is_online BOOLEAN NOT NULL DEFAULT 0,
    last_activity TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);