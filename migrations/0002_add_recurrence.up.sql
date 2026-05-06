ALTER TABLE tasks  --первое действие  - решил  что необходимо создать миграцию под новую функцию, новые типы 
ADD COLUMN recurrence_type TEXT NOT NULL DEFAULT 'none',
ADD COLUMN recurrence_interval INT,
ADD COLUMN recurrence_days INT[],
ADD COLUMN recurrence_dates DATE[],--тип массива дат выбран чтобы можно было конкретные даты вписывать, а не как в days -^
ADD COLUMN recurrence_parity TEXT,-- для  четных-нечетных
ADD COLUMN recurrence_start DATE,
ADD COLUMN recurrence_end DATE;