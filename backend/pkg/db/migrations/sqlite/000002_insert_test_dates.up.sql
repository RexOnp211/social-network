INSERT INTO users (nickname, email, password, firstname, lastname, dob, aboutme, public, avatar)
VALUES 
('privateuser', 'private@example.com', 'password123', 'Private', 'User', '1990-01-01', 'This is a private user.', 0, ''),
('publicuser', 'public@example.com', 'password123', 'Public', 'User', '1990-01-01', 'This is a public user.', 1, ''),
('pu2', 'pu2@example.com', 'pu2', 'Public', 'User', '1990-01-01', 'This is another public user.', 1, '');

INSERT INTO posts (nickname, subject, content, privacy) VALUES
  ('pu2', 'My First Post', 'This is the content of my first post.', 'public'),
  ('pu2', 'Another Day', 'Today was a great day! I went for a walk in the park.', 'public'),
  ('pu2', 'Cooking Recipe', 'Here is a delicious recipe for you to try out.', 'private'),
  ('pu2', 'Travel Diaries', 'I just came back from a trip to the mountains.', 'public'),
  ('publicuser', 'Book Review', 'I recently read a fantastic book that I would recommend.', 'public'),
  ('privateuser', 'Private Thoughts', 'Sometimes I just need to write down my thoughts.', 'private');

INSERT INTO groups (creator_name, title, description) VALUES
  ('pu2', 'test group', 'this is group description'),
  ('pu2', 'test group 2', 'this is group description'),
  ('pu2', 'test group 3', 'this is group description'),
  ('pu2', 'test group 4', 'this is group description'),
  ('pu2', 'test group long name', 'this is group description'),
  ('pu2', 'test group 5', 'this is group description'),
  ('publicuser', 'publicuser made', 'this is group description');


INSERT INTO group_members (title, nickname, status) VALUES
  ('test group', 'publicuser', 'invited'),
  ('test group 2', 'publicuser', 'invited'),
  ('test group 3', 'publicuser', 'invited'),
  ('test group 4', 'publicuser', 'invited'),
  ('test group long name', 'publicuser', 'invited');