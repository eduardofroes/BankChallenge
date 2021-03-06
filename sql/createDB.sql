

CREATE TABLE accounts (
    ID UUID NOT NULL,
    NAME VARCHAR(100) NOT NULL,
    CPF VARCHAR(15) NOT NULL,
    SECRET VARCHAR(200) NOT NULL,
    BALANCE DECIMAL NOT NULL,
    CREATED_AT TIMESTAMP NOT NULL,
    PRIMARY KEY (ID)
);

CREATE TABLE transfers (
    ID UUID NOT NULL,
    ACCOUNT_ORIGIN_ID UUID NOT NULL,
    ACCOUNT_DESTINATION_ID UUID NOT NULL,
    AMOUNT DECIMAL NOT NULL,
    CREATED_AT TIMESTAMP NOT NULL,
    PRIMARY KEY (ID)
);