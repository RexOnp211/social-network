INSERT INTO users (nickname, email, password, firstname, lastname, dob, aboutme, public, avatar)
VALUES 
('privateuser', 'private@example.com', 'password123', 'Private', 'User', '1990-01-01', 'This is a private user.', 0, ''),
('publicuser', 'public@example.com', 'password123', 'Public', 'User', '1990-01-01', 'This is a public user.', 1, ''),
('pu2', 'pu2@example.com', 'pu2', 'Public', 'User', '1990-01-01', 'This is another public user.', 1, '');

INSERT INTO posts (user_id, subject, content, privacy) VALUES
  ('1', 'My First Post', 'This is the content of my first post.', 'public'),
  ('1', 'Another Day', 'Today was a great day! I went for a walk in the park.', 'public'),
  ('1', 'Cooking Recipe', 'Here is a delicious recipe for you to try out.', 'private'),
  ('2', 'Travel Diaries', 'I just came back from a trip to the mountains.', 'public'),
  ('2', 'Book Review', 'I recently read a fantastic book that I would recommend.', 'public'),
  ('2', 'Private Thoughts', 'Sometimes I just need to write down my thoughts.', 'private');
