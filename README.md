# Unit Converter
# Конвертер единиц измерения

Unit Converter - это простой веб-сервис, который позволяет преобразовывать значения между различными единицами измерения. Проект написан на языке Go.
# Функциональность
Unit Converter поддерживает следующие единицы измерения:
Миллиметры (мм)
Сантиметры (см)
Дециметры (дм)
Метры (м)

# Установка
Чтобы установить Unit Converter, вы можете клонировать этот репозиторий и запустить команду `go run main.go`. Это запустит веб-сервер на порту 8080.

# Использование
Чтобы использовать Unit Converter, вы можете сделать GET-запрос на URL `http://localhost:8080/?from=мм&value=20&to=см`. В этом примере мы преобразуем 20 мм в см.
Параметры запроса
- `from` : единица измерения, из которой мы хотим преобразовать значение (мм, см, дм, м)
- `value` : значение, которое мы хотим преобразовать
- `to` : единица измерения, в которую мы хотим преобразовать значение (мм, см, дм, м)
