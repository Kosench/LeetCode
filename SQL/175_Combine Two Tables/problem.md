### 175. Объединение двух таблиц
Решено  
Легкая сложность

#### Задача:
Даны две таблицы: `Person` и `Address`. Необходимо составить запрос, который выведет имя (`firstName`), фамилию (`lastName`), город (`city`) и штат (`state`) для каждого человека из таблицы `Person`. Если у человека нет адреса в таблице `Address`, то в полях `city` и `state` должно быть значение `NULL`.

---

#### Структура таблиц:

**Таблица Person:**

| Column Name | Type    |
|-------------|---------|
| personId    | int     |
| lastName    | varchar |
| firstName   | varchar |

- `personId` — первичный ключ (уникальное значение) для этой таблицы.
- Таблица содержит информацию об ID людей и их именах.

---

**Таблица Address:**

| Column Name | Type    |
|-------------|---------|
| addressId   | int     |
| personId    | int     |
| city        | varchar |
| state       | varchar |

- `addressId` — первичный ключ (уникальное значение) для этой таблицы.
- Каждая строка таблицы содержит информацию о городе и штате для одного человека с соответствующим `personId`.

---

#### Пример:

**Ввод:**

Таблица `Person`:

| personId | lastName | firstName |
|----------|----------|-----------|
| 1        | Wang     | Allen     |
| 2        | Alice    | Bob       |

Таблица `Address`:

| addressId | personId | city          | state      |
|-----------|----------|---------------|------------|
| 1         | 2        | New York City | New York   |
| 2         | 3        | Leetcode      | California |

**Вывод:**

| firstName | lastName | city          | state    |
|-----------|----------|---------------|----------|
| Allen     | Wang     | NULL          | NULL     |
| Bob       | Alice    | New York City | New York |

**Объяснение:**
- Для `personId = 1` в таблице `Address` нет записи, поэтому в полях `city` и `state` выводится `NULL`.
- Для `personId = 2` есть запись в таблице `Address`, поэтому выводятся соответствующие значения города и штата.

---

#### Решение (SQL):

Для решения задачи нужно использовать **LEFT JOIN**, чтобы объединить таблицы `Person` и `Address` по полю `personId`. Если для какого-то `personId` в таблице `Address` нет записи, то в результатах будут `NULL` в столбцах `city` и `state`.

```sql
SELECT 
    p.firstName,
    p.lastName,
    a.city,
    a.state
FROM 
    Person p
LEFT JOIN 
    Address a
ON 
    p.personId = a.personId;
```

---

#### Пояснение к запросу:
1. **FROM Person p**: Выбираем данные из таблицы `Person` и используем псевдоним `p`.
2. **LEFT JOIN Address a ON p.personId = a.personId**: Объединяем таблицу `Person` с таблицей `Address` по условию совпадения `personId`. Используется `LEFT JOIN`, чтобы сохранить все записи из таблицы `Person`, даже если для них нет соответствующих записей в таблице `Address`.
3. **SELECT**: Выводим необходимые столбцы: `firstName`, `lastName`, `city` и `state`.

---

#### Результат:
Запрос вернет таблицу с именами, фамилиями, городами и штатами. Если для какого-то человека нет адреса, то в полях `city` и `state` будет значение `NULL`.