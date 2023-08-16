

CREATE DATABASE db_testMeli

use db_testMeli
-- exercicio 2
CREATE TABLE tb_estudantes(
id bigint auto_increment,
title BIGINT NOT NULL,
price BIGINT NOT NULL,
quantity int,

PRIMARY KEY (id)
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;
