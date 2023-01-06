-- +migrate Up
CREATE TABLE IF NOT EXISTS `gank`.order
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id`     INT          NOT NULL,
    `order_code`  VARCHAR(11)  NOT NULL,
    `final_price` FLOAT        NOT NULL,
    `status`      INT          NOT NULL,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `gank`.order;