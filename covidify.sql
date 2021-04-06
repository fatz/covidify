CREATE TABLE `visitors` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `visit_id` varchar(36) NOT NULL,
  `name` varchar(120) NOT NULL,
  `email` varchar(120) NULL,
  `phone` varchar(120) NOT NULL,
  `country` varchar(120) NOT NULL,
  `city` varchar(120) NOT NULL,
  `zip_code` varchar(120) NOT NULL,
  `street` varchar(120) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `ix_visitor_visit_id` (`visit_id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `risks` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `risk` varchar(120) NOT NULL,
  `description` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `visits` (
  `id` varchar(36) NOT NULL,
  `table_number` varchar(120) NOT NULL,
  `check_in` datetime NULL,
  `check_out` datetime NULL,
  `risk_id` bigint(20) NULL,
  PRIMARY KEY (`id`),
  KEY `risk_id` (`risk_id`)
  -- CONSTRAINT `visit_fk_1` FOREIGN KEY (`risk_id`) REFERENCES `risks` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
