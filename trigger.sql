DELIMITER $$

USE `price_comparator`$$

DROP TRIGGER /*!50032 IF EXISTS */ `price`$$

CREATE
    /*!50017 DEFINER = 'root'@'localhost' */
    TRIGGER `price` AFTER UPDATE ON `commodity_item` 
    FOR EACH ROW BEGIN
    DECLARE commodity_item_id INT;
    DECLARE current_price DECIMAL(10, 2);
    DECLARE create_at DATETIME;
    
    SET commodity_item_id = NEW.id;
    SET current_price = NEW.price;
    SET create_at = NOW();
    INSERT INTO price_change (commodity_item_id, new_price, update_at)
	VALUES (commodity_item_id, current_price, create_at);
    IF current_price <> 0 AND current_price < OLD.price THEN
        INSERT INTO message (current_price, commodity_item_id, create_at, user_id)
          (SELECT current_price, commodity_item_id, create_at, favorite.user_id FROM favorite WHERE favorite.commodity_item_id = commodity_item_id AND favorite.price_limit >= current_price);
    END IF;
    END;
$$

DELIMITER ;