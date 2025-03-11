# Онлайн-библиотека песен API  

## Описание  
Этот проект представляет собой REST API для управления музыкальной библиотекой.  
С помощью API можно добавлять, просматривать и удалять песни.  

## Возможности  
- Добавление новых песен  
- Получение списка всех песен  
- Удаление песни по ID  
- Документирование API с помощью Swagger  

## Установка и запуск  

## 1. Клонирование репозитория  
Склонируйте проект на свой компьютер:  
    ``` sh
      git clone https://github.com/yourusername/online-music-library.git
      cd online-music-library
    ```
## 2. Установка зависимостей
    ``` sh
      go mod tidy
    ```
## 3. Запуск сервера
    ``` sh
      go run cmd/main.go
    ```
### Сервер будет запущен на http://localhost:8080.

## API Эндпоинты
## 1. Получение списка всех песен
### Метод: GET
### URL: /songs
### Описание: Возвращает список всех доступных песен

### Ответ:
json
    ```
      [
        {
          "id": 1,
          "group": "Imagine Dragons",
          "song": "Believer",
          "releaseDate": "2017-02-01",
          "text": "First things first...",
          "link": "https://example.com/song",
          "created_at": "2025-03-11T20:24:35.0504698+05:00",
          "updated_at": "2025-03-11T20:24:35.0504698+05:00"
        }
      ]
    ```
## 2. Добавление новой песни
### Метод: POST
### URL: /songs
### Описание: Добавляет новую песню в базу данных

### Тело запроса:
  ```json
  {
    "group": "Imagine Dragons",
    "song": "Believer"
  }
  ```
### Ответ:
  ```json
  { 
    "id": 2,
    "group": "Imagine Dragons",
    "song": "Believer",
    "releaseDate": "2017-02-01",
    "text": "First things first...",
    "link": "https://example.com/song",
    "created_at": "2025-03-11T20:25:00.123456+05:00",
    "updated_at": "2025-03-11T20:25:00.123456+05:00"
  }
  ```
## 3. Удаление песни по ID
### Метод: DELETE
### URL: /songs/{id}
### Описание: Удаляет песню по ID

### Пример запроса:
  ```sh
  DELETE /songs/1
  ```
### Ответ:
  ```
  json
  {
    "message": "Песня удалена"
  }
  ```
# Документация API
## Swagger-документация доступна по адресу:
  ```
  bash
  http://localhost:8080/swagger/index.html
  ```
### Используемые технологии
### Go (Gin) — веб-фреймворк
### PostgreSQL — база данных
### Swagger — документация API
# Авторы
## Aidar Sabyrgali
