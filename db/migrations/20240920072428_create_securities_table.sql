-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS security
(
    ticker    VARCHAR(255) PRIMARY KEY,
    shortname VARCHAR(255),
    secname   VARCHAR(255)
);
COMMENT ON TABLE security IS 'Эмитенты';
COMMENT ON COLUMN security.ticker IS 'Тикер';
COMMENT ON COLUMN security.shortname IS 'Краткое наименование';
COMMENT ON COLUMN security.secname IS 'Полное наименование';

CREATE INDEX idx_security_ticker ON security (ticker);

CREATE TYPE currency AS ENUM ('RUB', 'USD', 'EUR', 'CYN');
COMMENT ON TYPE currency IS 'Валюты';

CREATE TABLE quotes
(
    id     SERIAL PRIMARY KEY,
    ticker VARCHAR(255) REFERENCES security (ticker),
    price  DECIMAL(10, 2),
    time   TIMESTAMP
);
COMMENT ON TABLE quotes IS 'Котировки';
comment on COLUMN quotes.ticker IS 'Тикер';
comment on COLUMN quotes.price IS 'Цена';
comment on COLUMN quotes.time IS 'Время';

CREATE TABLE "user"
(
    id       SERIAL PRIMARY KEY,
    name     VARCHAR(255),
    telegram VARCHAR(255)
);
COMMENT ON COLUMN "user".name IS 'Имя';
COMMENT ON COLUMN "user".name IS 'Телеграмм';

CREATE TABLE user_security_fulfil
(
    id               SERIAL PRIMARY KEY,
    ticker           VARCHAR(255) REFERENCES security (ticker),
    p_e_msfo_fulfil  DECIMAL(10, 2),
    p_bv_msfo_fulfil DECIMAL(10, 2)
);
COMMENT ON TABLE user_security_fulfil IS 'Цели пользователей по эмитентам';
COMMENT ON COLUMN user_security_fulfil.ticker IS 'Тикер';
COMMENT ON COLUMN user_security_fulfil.p_e_msfo_fulfil IS 'Цель по P/E (МСФО)';
COMMENT ON COLUMN user_security_fulfil.p_bv_msfo_fulfil IS 'Цель по P/BV (МСФО)';

CREATE TABLE security_financial_report_msfo
(
    id SERIAL PRIMARY KEY,
    ticker VARCHAR(255) REFERENCES security(ticker),
    year SMALLINT,
    quarter SMALLINT CHECK ( quarter >= 1 AND quarter <= 4),
    report_date DATE,
    income DECIMAL(15, 2),
    operation_income DECIMAL(15, 2),
    net_income DECIMAL(15,2),
    dividends_total DECIMAL(15, 2),
    dividends_to_income DECIMAL(2, 2) CHECK ( dividends_to_income >= 0 AND dividends_to_income <= 100  ),
    operational_expenses DECIMAL(15, 2),
    amortization DECIMAL(15, 2),
    employee_expenses DECIMAL(15, 2),
    interest_expenses DECIMAL(15, 2),
    reserves_creation DECIMAL(15, 2),
    assets DECIMAL(15, 2),
    net_assets DECIMAL(15, 2),
    debt DECIMAL(15, 2),
    cash DECIMAL(15, 2),
    net_debt DECIMAL(15, 2),
    total_shares BIGINT,
    book_value DECIMAL(15, 2)
);
COMMENT ON COLUMN security_financial_report_msfo.ticker IS 'Тикер';
COMMENT ON COLUMN security_financial_report_msfo.year IS 'Год' ;
COMMENT ON COLUMN security_financial_report_msfo.quarter IS 'Квартал' ;
COMMENT ON COLUMN security_financial_report_msfo.report_date IS 'Дата отчета' ;
COMMENT ON COLUMN security_financial_report_msfo.income IS 'Выручка';
COMMENT ON COLUMN security_financial_report_msfo.operation_income IS 'Операционная прибыль';
COMMENT ON COLUMN security_financial_report_msfo.net_income IS 'Чистая прибыль';
COMMENT ON COLUMN security_financial_report_msfo.dividends_total IS 'Дивидендные выплаты';
COMMENT ON COLUMN security_financial_report_msfo.dividends_to_income IS 'Дивиденды / прибыль %';
COMMENT ON COLUMN security_financial_report_msfo.operational_expenses IS 'Операционные расходы';
COMMENT ON COLUMN security_financial_report_msfo.amortization IS 'Амортизация';
COMMENT ON COLUMN security_financial_report_msfo.employee_expenses IS 'Расходы на персонал';
COMMENT ON COLUMN security_financial_report_msfo.interest_expenses IS 'Процентные расходы';
COMMENT ON COLUMN security_financial_report_msfo.reserves_creation IS 'Создание резервов';
COMMENT ON COLUMN security_financial_report_msfo.assets IS 'Активы';
COMMENT ON COLUMN security_financial_report_msfo.net_assets IS 'Чистые активы';
COMMENT ON COLUMN security_financial_report_msfo.debt IS 'Долг';
COMMENT ON COLUMN security_financial_report_msfo.cash IS 'Наличность';
COMMENT ON COLUMN security_financial_report_msfo.net_debt IS 'Чистый доход';
COMMENT ON COLUMN security_financial_report_msfo.total_shares  IS 'Всего акций';
COMMENT ON COLUMN security_financial_report_msfo.book_value IS 'Балансовая стоимость';


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS security_financial_report_msfo;
DROP TABLE IF EXISTS user_security_fulfil;
DROP TABLE IF EXISTS quotes;
DROP TABLE IF EXISTS security;
DROP TABLE IF EXISTS "user";
DROP TYPE currency;
-- +goose StatementEnd
