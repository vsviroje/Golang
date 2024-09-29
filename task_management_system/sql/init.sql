CREATE TABLE IF NOT EXISTS  `tasks` (
  `id` varchar(50) NOT NULL,
  `user_id` varchar(50) NULL,
  `status` varchar(50) NOT NULL,
  `title` varchar(50) NULL,
  `description` text NULL,
  `due_date` timestamp NULL,
  `is_deleted` tinyint(1) DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  FOREIGN KEY (`user_id`) REFERENCES users(id)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE IF NOT EXISTS  `users` (
  `id` varchar(50) NOT NULL,
  `name` varchar(50) NULL,
  `email` varchar(50) NOT NULL,
  `is_deleted` tinyint(1) DEFAULT 0,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`),
  UNIQUE (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;