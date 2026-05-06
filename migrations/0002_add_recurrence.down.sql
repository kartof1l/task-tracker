ALTER TABLE tasks --второе действие - возможность отката новых столбцов
    DROP COLUMN recurrence_type,
    DROP COLUMN recurrence_interval,
    DROP COLUMN recurrence_days,
    DROP COLUMN recurrence_dates,
    DROP COLUMN recurrence_parity,
    DROP COLUMN recurrence_start,
    DROP COLUMN recurrence_end;