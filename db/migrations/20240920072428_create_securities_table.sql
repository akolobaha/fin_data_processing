-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE TYPE currency AS ENUM ('RUB', 'USD', 'EUR', 'CYN');
COMMENT ON TYPE currency IS 'Валюты';

CREATE TABLE quotes
(
    id     SERIAL PRIMARY KEY,
    ticker VARCHAR(10),
    price  DECIMAL(10, 2),
    time   TIME,
    seq_num BIGINT
);
COMMENT ON TABLE quotes IS 'Котировки';
comment on COLUMN quotes.ticker IS 'Тикер';
comment on COLUMN quotes.price IS 'Цена';
comment on COLUMN quotes.time IS 'Время';
comment on COLUMN quotes.seq_num IS 'Номер обновления';

CREATE TABLE security_financial_report_msfo
(
    id SERIAL PRIMARY KEY,
    ticker VARCHAR(10),
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
DROP TABLE IF EXISTS quotes;
DROP TYPE currency;
-- +goose StatementEnd
