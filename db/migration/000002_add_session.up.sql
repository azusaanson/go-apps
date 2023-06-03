CREATE TABLE `session` (
  `id` integer PRIMARY KEY AUTO_INCREMENT,
  `uuid` varchar(255) NOT NULL,
  `user_id` integer NOT NULL,
  `refresh_token` varchar(255) NOT NULL,
  `user_agent` varchar(255) NOT NULL,
  `client_ip` varchar(255) NOT NULL,
  `is_blocked` boolean NOT NULL DEFAULT false,
  `expires_at` timestamp NOT NULL,
  `updated_at` timestamp NOT NULL DEFAULT (now()),
  `created_at` timestamp NOT NULL DEFAULT (now())
);

ALTER TABLE `session` ADD FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
