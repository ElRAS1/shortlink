CREATE TABLE IF NOT EXISTS link (
    oldlink VARCHAR(255),
    newlink VARCHAR(255),
    PRIMARY KEY (oldlink, newlink)
);

