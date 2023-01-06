-- +migrate Up
CREATE TABLE IF NOT EXISTS `gank`.cart
(
    `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `product_id`    INT          NOT NULL,
    `user_id`       INT          NOT NULL,
    `quantity`      FLOAT        NOT NULL,
    `is_claimed`    TINYINT      NOT NULL,
    `order_claimed` VARCHAR(11)  NOT NULL,
    `created_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;
-- +migrate Down
DROP TABLE IF EXISTS `gank`.cart;
