# Платёжная система

Приложение реализует систему обработки транзакций платёжной системы с использованием Go, PostgreSQL. Три основных эндпоинта:

1. **POST /api/send** - Перевод средств с одного кошелька на другой.
2. **GET /api/transactions?count=N** - Получение N последних транзакций.
3. **GET /api/wallet/{address}/balance** - Получение баланса кошелька.

## Описание функционала

### 1. Send (POST /api/send)

Этот эндпоинт позволяет отправить средства с одного кошелька на другой.

**Тело запроса (JSON):**

```json
{
  "from": "e240d825d255af751f5f55af8d9671beabdf2236c0a3b4e2639b3e182d994c88",
  "to": "abcd1234efgh5678ijkl9012mnop3456qrst7890uvwx1234yzab5678cdef1234",
  "amount": 3.50
}
```

Если средств недостаточно:
    Ответ: insufficient funds

    




### 2. GetLast (GET /api/transactions?count=N)

Этот эндпоинт возвращает информацию о последних N переводах.

**Пример запроса:**
    GET /api/transactions?count=4

```json
{
    [{
        "ID":7,
        "From":"9e2af27ff230d7742a93f39fe82cfa211783d6a52c0b857d6343b70d73eb7f4a","To":"55cd13a495cba0f737d0e0f4ec09cce64358dfcc32bcfea60dbf4c6c6c614d5a",
        "Amount":12,
        "CreatedAt":"2025-01-18T21:29:31.27584Z"
        },

    {
        "ID":6,
        "From":"9e2af27ff230d7742a93f39fe82cfa211783d6a52c0b857d6343b70d73eb7f4a","To":"55cd13a495cba0f737d0e0f4ec09cce64358dfcc32bcfea60dbf4c6c6c614d5a",
        "Amount":55.34,
        "CreatedAt":"2025-01-18T21:22:20.607422Z"
    },

    {
        "ID":5,
        "From":"a40e79c4ed71eac68881bc59d8fb451343eb42cdee699f11f0864d0f83a4864f","To":"55cd13a495cba0f737d0e0f4ec09cce64358dfcc32bcfea60dbf4c6c6c614d5a",
        "Amount":55.34,
        "CreatedAt":"2025-01-18T21:22:15.252042Z"
    },
    {
        "ID":4,
        "From":"c861612b8ad0012872dfed2c0ebcd3abc34c56ace9c2c98c38cb999eabf9af3f","To":"63d3ca503a36d02b0f0edb16f8008b9de1eba18dffd2f744fb1426be4629a7d2",
        "Amount":55.34,
        "CreatedAt":"2025-01-18T21:20:03.217459Z"
    }]
}
```



### 3. GetBalance (GET /api/wallet/{address}/balance)

Этот эндпоинт возвращает информацию о балансе кошелька.

**Пример запроса:** 
    GET /api/wallet/{9e2af27ff230d7742a93f39fe82cfa211783d6a52c0b857d6343b70d73eb7f4a}/balance


```json
{
    "ID":4,
    "Address":"9e2af27ff230d7742a93f39fe82cfa211783d6a52c0b857d6343b70d73eb7f4a",
    "Balance":32.66
}
```



### . Сборка и запуск контейнера

Для того чтобы собрать и запустить проект, используйте команду:

```bash
start.bat
```
