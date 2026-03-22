# CutterService

## Запуск

### 1. In-memory режим

```bash
# Docker
make im

# Или вручную
docker build -t cutter-service .
docker run -dp 8080:8080 --name cutter-im cutter-service ./cutter -m=im
```

### 2. PostgreSQL режим

``` bash
# Docker (поднимает app + postgres)
ake psql

# Или вручную
docker-compose up -d --build
```

### 3. Завершение
``` bash
make down

# Или через команды docker

# Для postgres режима
docker-compose down 

# Для in-memory режима
docker stop cutter-im
```

## Использование

### Создать короткую ссылку

``` bash
POST http://localhost:8080/
Content-Type: application/json

{
  "url": "https://example.com/very/long/url"
}

# Ответ
{
  "short_url": "abc123_def"
}
```

### Получение оригинальной ссылки
``` bash 
GET http://localhost:8080/abc123_def

# Ответ
{
  "original_url": "https://example.com/very/long/url"
}
```