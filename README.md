# Book Search Telegram Bot
Отправляете боту название книги, он возвращает страницы поиска книги на следующих площадках:
- Ozon
- Avito
- Wildberries
- Yandex Маркет
- МегаМаркет (сбер)
- Издательство "Питер"
- Издательство "ДМК"
- Читай-город
- Издательство Эксмо
- Издательство АСТ
- Издательство Бомбора
- Издательство Альпина
- Издательство МИФ

# Зависимости
- Docker
- Go 1.21
- [telegram-bot-api v5](https://github.com/go-telegram-bot-api/telegram-bot-api)

# Ограничения
1 сообщение в 10 секунд.

# Запуск
```zsh
make build TAG=<TAG>
make run TOKEN=<TOKEN> TAG=<TAG>
```
