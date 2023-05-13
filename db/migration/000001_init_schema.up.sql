CREATE TABLE `user` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL COMMENT 'admin, user',
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `created_at` timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE `invest` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `user_id` integer,
  `amount` decimal(15,2) COMMENT 'HKD',
  `type` varchar(255) NOT NULL COMMENT 'CEX, DEX',
  `invested_at` timestamp NOT NULL DEFAULT (now()),
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `created_at` timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE `invest` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
