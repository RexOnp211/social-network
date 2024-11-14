DELETE FROM users WHERE nickname = 'privateuser';
DELETE FROM users WHERE nickname = 'publicuser';
DELETE FROM users WHERE nickname = 'pu2';
DELETE FROM posts WHERE user_id = 1;
DELETE FROM posts WHERE user_id = 2;
DELETE FROM comments WHERE user_id = 1;
DELETE FROM comments WHERE user_id = 2;
DELETE FROM comments WHERE user_id = 3;
