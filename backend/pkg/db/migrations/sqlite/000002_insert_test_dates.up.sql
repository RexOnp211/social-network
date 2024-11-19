INSERT INTO users (nickname, email, password, firstname, lastname, dob, aboutme, public, avatar)
VALUES
('privateuser', 'private@example.com', '$2y$10$M0DAN.jRAYbA42.3UHg7o.0Li5JYQBt0WGTEl18V2jW3LKOMx7moq', 'Private', 'User', '1990-01-01', 'This is a private user.', 0, ''),
-- password "pr"
('publicuser', 'public@example.com', '$2y$10$yXTT5adkZxaWM6Xd7NTAlOYqJvfqT/pBjIsH.cHbF1xBWdD6FlT3a', 'Public', 'User', '1990-01-01', 'This is a public user.', 1, ''),
-- password "pu"
('pu2', 'pu2@example.com', '$2y$10$sGvVL63Ve8IWC7DPe3MDPuC6sjMODPL6OZmZ/4xv4TAStxkkkcMee', 'Public', 'User', '1990-01-01', 'This is another public user.', 1, '');
-- password "pu2"

INSERT INTO posts (user_id, subject, content, privacy) VALUES
  ('1', 'My First Post', 'This is the content of my first post.', 'public'),
  ('1', 'Another Day', 'Today was a great day! I went for a walk in the park.', 'public'),
  ('1', 'Cooking Recipe', 'Here is a delicious recipe for you to try out.', 'private'),
  ('2', 'Travel Diaries', 'I just came back from a trip to the mountains.', 'public'),
  ('2', 'Book Review', 'I recently read a fantastic book that I would recommend.', 'public'),
  ('2', 'Private Thoughts', 'Sometimes I just need to write down my thoughts.', 'private');

INSERT INTO comments (post_id, user_id, content) VALUES
  ('1', '2', 'Great post!'),
  ('1', '3', 'I agree!'),
  ('2', '3', 'Sounds like a fun day!'),
  ('3', '1', 'I would love to try this recipe!'),
  ('4', '1', 'I love the mountains!'),
  ('5', '3', 'What book did you read?'),
  ('6', '1', 'I understand the need to write down your thoughts.');

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
  ('test group long name', 'publicuser', 'invited'),
  ('publicuser made', 'pu2', 'invited');