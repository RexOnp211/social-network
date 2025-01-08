
----------------------- USER DATA

CREATE TABLE IF NOT EXISTS users (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    nickname TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    password TEXT DEFAULT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    dob DATE NOT NULL, -- date of birth
    aboutme TEXT NOT NULL,
    public INTEGER DEFAULT 1, -- profile public status: 1 = public, 0 = private
    avatar TEXT
);

----------------------- GROUP DATA

CREATE TABLE IF NOT EXISTS groups (
    chatId INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    creator_name TEXT NOT NULL,
    title TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    FOREIGN KEY(creator_name) REFERENCES users(nickname)
);
CREATE TABLE IF NOT EXISTS memberships (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    title TEXT NOT NULL,
    nickname TEXT NOT NULL,
    status TEXT NOT NULL,  -- status ('requested', 'invited', 'approved')
    FOREIGN KEY(title) REFERENCES groups(title),
    FOREIGN KEY(nickname) REFERENCES users(nickname),
    UNIQUE (title, nickname)
);
CREATE TABLE IF NOT EXISTS group_posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    group_title TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    nickname TEXT NOT NULL,
    subject TEXT NOT NULL,
    content TEXT NOT NULL,
    image TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(group_title) REFERENCES groups(title),
    FOREIGN KEY(user_id) REFERENCES users(user_id),
    FOREIGN KEY(nickname) REFERENCES users(nickname)
);
CREATE TABLE IF NOT EXISTS group_comments (
    comment_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    nickname TEXT NOT NULL,
    content TEXT NOT NULL,
    image TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(post_id) REFERENCES posts(post_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id),
    FOREIGN KEY(nickname) REFERENCES users(nickname)
);
CREATE TABLE IF NOT EXISTS group_events (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    group_title TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    nickname TEXT NOT NULL,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    event_date TIMESTAMP NOT NULL,
    FOREIGN KEY(group_title) REFERENCES groups(title),
    FOREIGN KEY(user_id) REFERENCES users(user_id),
    FOREIGN KEY(nickname) REFERENCES users(nickname)
);
CREATE TABLE IF NOT EXISTS user_event_status (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    nickname TEXT NOT NULL,
    event_id INTEGER NOT NULL,
    going BOOLEAN NOT NULL DEFAULT 0,
    FOREIGN KEY(nickname) REFERENCES users(nickname),
    FOREIGN KEY(event_id) REFERENCES group_events(id),
    UNIQUE (nickname, event_id)
);

----------------------- POST DATA

CREATE TABLE IF NOT EXISTS posts (
    post_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    user_id INTEGER NOT NULL,
    subject TEXT NOT NULL,
    content TEXT NOT NULL,
    image TEXT DEFAULT '',
    privacy TEXT DEFAULT 'public',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS post_privacy (
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    FOREIGN KEY(post_id) REFERENCES posts(post_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id),
    UNIQUE(post_id, user_id)
);
CREATE TABLE IF NOT EXISTS categories (
    category_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    category TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS comments (
    comment_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    post_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,
    content TEXT NOT NULL,
    image TEXT DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(post_id) REFERENCES posts(post_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS likes (
    like_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    post_id INTEGER,
    comment_id INTEGER,
    user_id INTEGER NOT NULL,
    type INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(post_id) REFERENCES posts(post_id),
    FOREIGN KEY(comment_id) REFERENCES comments(comment_id),
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);
CREATE TABLE IF NOT EXISTS post_categories (
    post_id INTEGER NOT NULL,
    category_id INTEGER NOT NULL,
    FOREIGN KEY(post_id) REFERENCES posts(post_id),
    FOREIGN KEY(category_id) REFERENCES categories(category_id)
);

----------------------- SESSION DATA

CREATE TABLE IF NOT EXISTS sessions (
    session_id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    token TEXT NOT NULL,
    nickname TEXT NOT NULL,
    user_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);

----------------------- MESSAGE DATA

CREATE TABLE IF NOT EXISTS messages (
    group_id INTEGER NOT NULL,
    message_from INTEGER NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    is_read BOOLEAN NOT NULL DEFAULT 0
);

CREATE TABLE IF NOT EXISTS chatRoom (
    group_id INTEGER NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(chatId)
);

CREATE TABLE IF NOT EXISTS chatRoomMembers (
    group_id INTEGER NOT NULL,
    user_designation TEXT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES chatRoom(group_id),
    FOREIGN KEY (user_designation) REFERENCES users(nickname)
);

CREATE TABLE IF NOT EXISTS user_status (
    user_id UUID PRIMARY KEY NOT NULL,
    is_online BOOLEAN NOT NULL DEFAULT 0,
    last_activity TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(user_id)
);

CREATE TABLE IF NOT EXISTS followers (
    follower_id INTEGER NOT NULL,
    followee_id INTEGER NOT NULL,
    accepted BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (follower_id, followee_id),
    FOREIGN KEY(follower_id) REFERENCES users(user_id),
    FOREIGN KEY(followee_id) REFERENCES users(user_id)
);
