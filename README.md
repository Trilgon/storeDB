## Спецификация API

### Создание товара

- Метод: `POST`
- URL: `/api/goods/add`

Тело запроса (JSON):

```json
{
  "goods_id": 123,
  "name": "Ноутбук",
  "price": 50000,
  "quantity": 10
}
```

### Получение информации о товаре

- Метод: `GET`
- URL: `/api/goods/get?goods_id`

Ответ: 
```json
{
  "goods_id": 123,
  "name": "Ноутбук",
  "price": 50000,
  "quantity": 10
}
```

### Обновление информации о товаре

- Метод: `PUT`
- URL: `/api/goods/update?goods_id`

Тело запроса (JSON):

```json
{
  "name": "Ноутбук Asus",
  "price": 55000,
  "quantity": 15
}
```

Ответ: `204`

### Удаление товара

- Метод: `DELETE`
- URL: `/api/goods/delete?goods_id`

Ответ: `204`

### Создание корзины

- Метод: `POST`
- URL: `/api/carts/create`

Тело запроса (JSON): `нет`

Ответ: 
```json
{
  "cart_id": 4
}
```

### Добавление товара в корзину

- Метод: `PUT`
- URL: `/api/carts/goods/add?cart_id`

Тело запроса (JSON):

```json
{
  "goods_id": 123,
  "quantity": 2
}
```

Ответ: `204`

### Получение списка товаров в корзине

- Метод: `GET`
- URL: `/api/carts/goods/get`

Ответ:
```json
{
  "goods": [
    {
      "goods_id": 123,
      "name": "Ноутбук",
      "price": 50000,
      "quantity": 1
    },
    {
      "goods_id": 32,
      "name": "Планшет",
      "price": 8000,
      "quantity": 2
    }
  ],
  "total": 66000
}
```

### Обновление информации о товаре в корзине

- Метод: `PUT`
- URL: `/api/carts/goods/update?cart_id&goods_id`

Тело запроса (JSON):

```json
{
  "quantity": 3
}
```

Ответ: `204`

### Удаление товара из корзины

- Метод: `DELETE`
- URL: `/api/carts/goods/delete?cart_id&goods_id`

Ответ: `204`

### Удаление корзины

- Метод: `DELETE`
- URL: `/api/carts/delete?cart_id`

Ответ: `204`

### Оформление заказа на основе корзины

- Метод: `POST`
- URL: `/api/orders/create`

Тело запроса (JSON):

```json
{
  "cart_id": 12
}
```

Ответ:
```json
{
  "goods": [
    {
      "goods_id": 123,
      "name": "Ноутбук",
      "price": 50000,
      "quantity": 1
    },
    {
      "goods_id": 32,
      "name": "Планшет",
      "price": 8000,
      "quantity": 2
    }
  ],
  "total": 66000,
  "order_time": "2023-03-20T12:00:00Z",
  "finish_time": null
}
```

### Получение информации о заказе

- Метод: `GET`
- URL: `/api/orders/get?order_id`

Ответ:
```json
{
  "goods": [
    {
      "goods_id": 123,
      "name": "Ноутбук",
      "price": 50000,
      "quantity": 1
    },
    {
      "goods_id": 32,
      "name": "Планшет",
      "price": 8000,
      "quantity": 2
    }
  ],
  "total": 66000,
  "order_time": "2023-03-20T12:00:00Z",
  "finish_time": "2023-03-20T12:00:00Z" // nullable
}
```

### Удаление заказа

- Метод: `DELETE`
- URL: `/api/orders/delete?order_id`

Ответ: `204`