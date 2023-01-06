-- +migrate Up
CREATE TABLE IF NOT EXISTS `gank`.order_detail
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `order_id`    INT          NOT NULL,
    `product_id`  VARCHAR(11)  NOT NULL,
    `quantity`    FLOAT        NOT NULL,
    `price`       FLOAT        NOT NULL,
    `total_price` FLOAT        NOT NULL,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;
-- +migrate Down
