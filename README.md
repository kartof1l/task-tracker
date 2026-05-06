# Task Service

Сервис для управления задачами с HTTP API на Go. Модуль трекера задач медицинской информационной системы.

## Требования

- Go `1.23+`
- Docker и Docker Compose

## Быстрый запуск через Docker Compose

```bash
docker compose up --build
После запуска сервис будет доступен по адресу http://localhost:8080.

Если postgres уже запускался ранее со старой схемой, пересоздай volume:
docker compose down -v
docker compose up --build
Причина в том, что SQL-файлы из migrations/ монтируются в docker-entrypoint-initdb.d и применяются только при инициализации пустого data volume.

Swagger
Swagger UI:
http://localhost:8080/swagger/
OpenAPI JSON:
http://localhost:8080/swagger/openapi.json
API
Базовый префикс: /api/v1

Метод	Путь	Описание
POST	/api/v1/tasks	Создать задачу
GET	/api/v1/tasks	Список задач
GET	/api/v1/tasks/{id}	Получить задачу по ID
PUT	/api/v1/tasks/{id}	Обновить задачу
DELETE	/api/v1/tasks/{id}	Удалить задачу
Периодичность задач
Задача может быть разовой или периодической. Тип периодичности задаётся полем recurrence_type:
Тип	        Описание	                        Дополнительные поля
none	    Без периодичности (по умолчанию)	    —
daily	    Каждый n-й день	                    recurrence_interval — интервал (1 = каждый день)
monthly	    На определённые числа месяца	    recurrence_days — числа от 1 до 30
specific	На конкретные даты	                recurrence_dates — даты в формате YYYY-MM-DD
parity	    По чётным/нечётным дням	            recurrence_parity — "even" или "odd"

Общие поля для всех типов (кроме none):

recurrence_start — дата начала (опционально)

recurrence_end — дата окончания (опционально, null = бессрочно)

Примеры запросов
Ежедневная задача (каждый день):
POST /api/v1/tasks
{
  "title": "обзвон пациентов",
  "recurrence_type": "daily",
  "recurrence_interval": 1,
  "recurrence_start": "2026-05-01"
}
Ежемесячная задача (1 и 15 числа):
POST /api/v1/tasks
{
  "title": "формирование отчёта",
  "recurrence_type": "monthly",
  "recurrence_days": [1, 15]
}
Задача на конкретные даты:
POST /api/v1/tasks
{
  "title": "инвентаризация",
  "recurrence_type": "specific",
  "recurrence_dates": ["2026-05-10", "2026-06-20"]
}
Задача по чётным дням:
POST /api/v1/tasks
{
  "title": "отчётность",
  "recurrence_type": "parity",
  "recurrence_parity": "even",
  "recurrence_start": "2026-05-01",
  "recurrence_end": "2026-12-31"
}

Принятые решения и допущения:
Хранение дат: recurrence_dates хранится как []string для упрощения работы с PostgreSQL DATE[] через pgx без кастомных сканеров.

Числа месяца: для monthly допустимы значения 1–30. Для 29–31 — использовать specific.

Nullable поля: recurrence_interval, recurrence_parity, recurrence_start, recurrence_end — указатели. Отсутствуют в JSON-ответе (omitempty), если не заданы.

Валидация: проверка сочетания полей для каждого типа периодичности — в доменной модели (ValidateRecurrence). Нормализация и значения по умолчанию — в сервисном слое.

Обратная совместимость: существующие задачи без периодичности продолжают работать, recurrence_type по умолчанию "none".

Генерация задач: на данном этапе реализовано хранение настроек. Генерация экземпляров задач по расписанию — следующий шаг (фоновый воркер или генерация вперёд на период).
