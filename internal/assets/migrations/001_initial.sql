-- +migrate Up

create table transfers
(
    id            bigserial primary key,
    sender        text      not null,
    receiver      text      not null,
    amount        text      not null,
    status        text      not null,
    fee           text      not null,
    erc20         text      not null,
    r             text      not null,
    s             text      not null,
    v             int       not null,
    is_custom_fee bool      not null,
    created_at    timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- +migrate Down

drop table transfers;



