### 234. Палиндромный связанный список
Решено  
Легкая сложность

Дан указатель `head` на начало односвязного списка. Верните `true`, если связанный список является палиндромом, и `false` в противном случае.

---

### Примеры

**Пример 1:**

Ввод: `head = [1,2,2,1]`  
Вывод: `true`

Объяснение:  
Список симметричен, поэтому он является палиндромом.

---

**Пример 2:**

Ввод: `head = [1,2]`  
Вывод: `false`

Объяснение:  
Список несимметричен, поэтому он не является палиндромом.

---

### Ограничения:

- Количество узлов в списке находится в диапазоне `[1, 10^5]`.
- $ 0 \leq \text{Node.val} \leq 9 $

---

### Дополнительно:
Можете ли вы решить задачу за время $ O(n) $ и с использованием дополнительной памяти $ O(1) $?