-- Migration generated by the command below; DO NOT EDIT.
-- hydra:generate hydra migrations gen
CREATE TABLE `networks` (
  `id` char(36) NOT NULL,
  PRIMARY KEY(`id`),
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
);