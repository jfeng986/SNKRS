USE product;

INSERT INTO category(id, name)
VALUES 
(1, 'Road-Running'),
(2, 'Trail-Running'),
(3, 'Walking'),
(4, 'Hiking'),
(5, 'Court'),
(6, 'Cross-Training'),
(7, 'Casual'),
(8, 'Leather');

INSERT INTO product (id, cateid, name, subtitle, images, detail, price, stock, status, create_time, update_time)
VALUES
(1, 1, 'Speed Runner', 'A shoe for road running', 'speed_runner.jpg', 'This is a fast road running shoe', 100.00, 10, '1', '2023-09-01 12:01:00', '2023-09-01 12:01:00'),
(2, 2, 'TrailBlaze', 'Master of all trails', 'trail_blaze.jpg', 'This shoe is made for trail running', 120.00, 15, '1', '2023-09-02 13:14:00', '2023-09-02 13:14:00'),
(3, 3, 'Easy Walk', 'Your daily walking companion', 'easy_walk.jpg', 'This shoe is designed for walking', 80.00, 20, '1', '2023-09-03 14:00:00', '2023-09-03 14:00:00'),
(4, 4, 'Mountain Goer', 'Built for hiking', 'mountain_goer.jpg', 'This shoe is robust for all hiking terrains', 150.00, 5, '1', '2023-09-04 09:09:00', '2023-09-04 09:09:00'),
(5, 5, 'CourtMaster', 'Rule the court', 'court_master.jpg', 'This shoe is ideal for court sports', 110.00, 8, '1', '2023-09-05 10:00:00', '2023-09-05 10:00:00'),
(6, 6, 'X-Train', 'The all-around trainer', 'x_train.jpg', 'Ideal for cross-training', 130.00, 12, '1', '2023-09-05 11:11:00', '2023-09-05 11:11:00'),
(7, 7, 'CasualFit', 'Casual yet stylish', 'casual_fit.jpg', 'Ideal for daily use', 75.00, 25, '1', '2023-09-06 15:15:00', '2023-09-06 15:15:00'),
(8, 8, 'Classic Leather', 'Elegance in leather', 'classic_leather.jpg', 'This is a leather shoe', 200.00, 6, '1', '2023-09-07 16:16:00', '2023-09-07 16:16:00'),
(9, 1, 'Urban Runner', 'Run the city streets', 'urban_runner.jpg', 'Designed for urban running', 105.00, 11, '1', '2023-09-08 10:00:00', '2023-09-08 10:00:00'),
(10, 2, 'Rock Climber', 'Conquer rocky trails', 'rock_climber.jpg', 'Built for rocky trails', 125.00, 10, '1', '2023-09-09 11:00:00', '2023-09-09 11:00:00'),
(11, 3, 'Morning Walker', 'Start your day right', 'morning_walker.jpg', 'Perfect for morning walks', 70.00, 19, '1', '2023-09-10 12:00:00', '2023-09-10 12:00:00'),
(12, 4, 'Peak Explorer', 'Reach new heights', 'peak_explorer.jpg', 'Designed for peak performance', 155.00, 7, '1', '2023-09-11 13:00:00', '2023-09-11 13:00:00'),
(13, 5, 'Match Winner', 'Win every game', 'match_winner.jpg', 'Optimized for court sports', 115.00, 10, '1', '2023-09-12 14:00:00', '2023-09-12 14:00:00'),
(14, 6, 'Fitness Pro', 'Your gym partner', 'fitness_pro.jpg', 'The ultimate gym shoe', 135.00, 9, '1', '2023-09-13 15:00:00', '2023-09-13 15:00:00'),
(15, 7, 'Easy Day', 'Relax in style', 'easy_day.jpg', 'The perfect shoe for a casual day', 70.00, 30, '1', '2023-09-14 16:00:00', '2023-09-14 16:00:00'),
(16, 8, 'Elite Leather', 'Premium leather experience', 'elite_leather.jpg', 'Pure leather luxury', 220.00, 4, '1', '2023-09-15 17:00:00', '2023-09-15 17:00:00');

USE user;

INSERT INTO `user` (`username`, `password`, `phone`, `question`, `answer`) VALUES
('user1', 'password1', '1234567890', 'Question 1', 'Answer 1'),
('user2', 'password2', '2345678901', 'Question 2', 'Answer 2'),
('user3', 'password3', '3456789012', 'Question 3', 'Answer 3'),
('user4', 'password4', '4567890123', 'Question 4', 'Answer 4'),
('user5', 'password5', '5678901234', 'Question 5', 'Answer 5'),
('user6', 'password6', '6789012345', 'Question 6', 'Answer 6'),
('user7', 'password7', '7890123456', 'Question 7', 'Answer 7'),
('user8', 'password8', '8901234567', 'Question 8', 'Answer 8'),
('user9', 'password9', '9012345678', 'Question 9', 'Answer 9'),
('user10', 'password10', '0123456789', 'Question 10', 'Answer 10'),
('user11', 'password11', '0987654321', 'Question 11', 'Answer 11'),
('user12', 'password12', '9876543210', 'Question 12', 'Answer 12'),
('user13', 'password13', '8765432109', 'Question 13', 'Answer 13'),
('user14', 'password14', '7654321098', 'Question 14', 'Answer 14'),
('user15', 'password15', '6543210987', 'Question 15', 'Answer 15'),
('user16', 'password16', '5432109876', 'Question 16', 'Answer 16'),
('user17', 'password17', '4321098765', 'Question 17', 'Answer 17'),
('user18', 'password18', '3210987654', 'Question 18', 'Answer 18'),
('user19', 'password19', '2109876543', 'Question 19', 'Answer 19'),
('user20', 'password20', '1098765432', 'Question 20', 'Answer 20');
