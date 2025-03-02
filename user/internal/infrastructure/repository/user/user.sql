CREATE TABLE user (
                      id VARCHAR(36) NOT NULL COMMENT 'id',
                      name varchar(255) NULL COMMENT 'The username',
                      password varchar(255) NOT NULL DEFAULT '' COMMENT 'The user password',
                      mobile varchar(255) NOT NULL DEFAULT '' COMMENT 'The mobile phone number',
                      gender char(10) NOT NULL DEFAULT 'male' COMMENT 'gender, male|female|unknown',
                      nickname varchar(255) NULL DEFAULT '' COMMENT 'The nickname',
                      type tinyint(1) NULL DEFAULT 0 COMMENT 'The user type, 0:normal,1:vip, for test golang keyword',
                      balance INT NOT NULL DEFAULT 0.00 COMMENT 'user balance（单位分）',
                      version INT NOT NULL DEFAULT 0 COMMENT 'Optimistic concurrency version',
                      create_at timestamp NULL,
                      update_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                      UNIQUE mobile_index (mobile),
                      UNIQUE name_index (name),
                      PRIMARY KEY (id)
) ENGINE = InnoDB COLLATE utf8mb4_general_ci COMMENT 'user table';
