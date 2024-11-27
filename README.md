# Аутентификация пользователей с помощью jwt-токена

Цель работы — получение первичных знаний в области авторизации и аутентификации в контексте веб-приложений

В рамках данной лабораторной работы предлагается ознакомиться с наиболее популярным способом аутентификации и авторизации пользователей в веб-приложениях — jwt-токеном
 
## Основные теоретические сведения

Статья по реализации аутентификации с помощью jwt-токена на golang: https://ru.hexlet.io/courses/go-web-development/lessons/auth/theory_unit

Пример реализации jwt-аутентификации в Echo: https://echo.labstack.com/docs/cookbook/jwt

Данная лабораторная работа базируется на результатах, полученных в ходе выполнения 6-10 лабораторных работ

## Порядок выполнения

Для успешного выполнения лабораторной работы необходимо проделать следующие шаги:

1. Перекопировать код, полученный в ходе выполнения предыдущей лабораторной работы
2. Ознакомиться с теоретическими сведениями
3. Реализовать сервис Auth (регистрация пользователя и авторизация с выдачей jwt-токена)
4. Добавить в сервисы `count`, `hello` и `query` валидацию jwt-токена (если токен не валидент или отсутствует — возвращаем код 401)
5. Сделать отчёт и поместить его в директорию docs
6. Защитить лабораторную работу

## Содержание отчета

1. Титульный лист
2. Цель работы и задание
3. Ход работы со скриншотами и листингами результатов 
4. Заключение
5. Список использованных источников