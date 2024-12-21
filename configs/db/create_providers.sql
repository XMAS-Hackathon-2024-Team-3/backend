CREATE TABLE providers (
    id INT PRIMARY KEY,
    conversion FLOAT,
    avg_time FLOAT,
    min_sum FLOAT,
    max_sum FLOAT,
    limit_max FLOAT,
    limit_min FLOAT,
    commission FLOAT,
    currency VARCHAR(3)
)
