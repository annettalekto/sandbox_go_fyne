## Заметки по GO Fyne

Часто используемые элементы Fyne, особенности перевода на русский, разница в похожих свойствах.

### Меню

После создания меню, автоматически добавляется подменю "Quit", которое постоянно приходится переименовывать в "Выход" (если прога на русском языке).

<img src="menu/img/menu.png" alt="screen"/>

[my menu example](https://github.com/annettalekto/sandbox_go_fyne/tree/main/menu)

### Entry – поле ввода

<img src="entry/img/entry.png" alt="screen"/>

[my entry example](https://github.com/annettalekto/sandbox_go_fyne/tree/main/entry)

Полезное свойство Entry – **Wrapping**. При создании элемента `Wrapping = TextTruncate`, это самый адекватный вариант для однострочного виджета, так что можно вообще ничего не прописывать. 

Но есть еще варианты.

**Entry.Wrapping:**

- `TextTruncate` – если виджет заполняется больше, чем его ширина, то появляется вертикальная полоса прокрутки. Прям ок, редко нужно что-то другое.

- `TextWrapOff` – если заполнить больше ширины виджета, то окно проги разъезжается! Тупо выглядит. Интересно зачем это...

- `TextWrapBreak` – для мультистрочного виджета (тут просто прокрутка).

- `TextWrapWord` – для мультистрочного виджета (тут просто прокрутка).

**MultiLineEntry.Wrapping:**

- `TextTruncate` – если виджет заполняется больше своей ширины появляется вертикальная полоса прокрутки и если строк больше чем его длинна, появляется горизонтальная полоса прокрутки.

- `TextWrapOff`  – окно проги разъезжается! и в ширину и длину при вводе длинных строк ¯\\_(ツ)_/¯.

- `TextWrapBreak`  – если виджет заполняется больше ширины, то (вместо горизонтальной полосы прокрутки) все что не уместилось в одну строчку, переносится на другую строку (часто посередине слова). Вертикальная полоса прокрутки есть.

- `TextWrapWord`  –  тоже с переносом на новую строку, но по словам, не обрывая на полуслове. Самый удобный вариант для редактора. Вертикальная полоса прокрутки есть.

### Расширение базового типа

Стандартные виджеты обеспечивают минимальную функциональность, предполагается расширение базового типа для дополнения его нужными функциями. 

За основу берется виджет widget.BaseWidget и с помощью ExtendBaseWidget получаем доступ к его полям.

```go
type numericalEntry struct {
    widget.Entry
}

// создаем функцию конструктор на основе базового типа 
func newNumericalEntry() *numericalEntry {
    entry := &numericalEntry{}
    entry.ExtendBaseWidget(entry)
    return entry
}
```

Теперь можем внести дополнения к стандартным методам.

[my entry example](https://github.com/annettalekto/sandbox_go_fyne/tree/main/entry)

Полное описание: 
[Extending Widgets](https://developer.fyne.io/extend/extending-widgets)
[Numerical-entry](https://developer.fyne.io/extend/numerical-entry)

### Диалоговые окна

- **Confirm** – диалоговое окно с выбором "Да/Нет".
  
  <img src="dialog/img/dialog_confirm.PNG" alt="screen"/>

- **Custom** – диалоговое окно с возможностью добавить свой элемент CanvasObject (круг метку, цветной текст).
  
  <img src="dialog/img/dialog_custom.PNG" alt="screen"/>

- **Form** – позволяет добавить свой widget (даже несколько) в диалоговое окно. Можно использовать подсказку (поле hint), но работает она совсем не так как ожидалось (просто текст ниже widget).
  
  <img src="dialog/img/dialog_formitem.PNG" alt="screen"/>

- **Information** – обычное окно с информацией и кнопкой.
  
  <img src="dialog/img/dialog_information.PNG" alt="screen"/>

В диалоговых окнах GO Fyne изменить название кнопок (стандартное ok) на любое другое значение можно через методы **SetDismissText** и **SetConfirmText** (если нет этого поля при создании диалога).

мои примеры: [my dialog example](https://github.com/annettalekto/sandbox_go_fyne/tree/main/dialog)

нормальное описание: [Dialog List | Develop using Fyne](https://developer.fyne.io/explore/dialogs)

### Связь объекта с переменной (bindable value)

Иногда, если состояние объекта зависит только от одной переменной, удобно привязать объект к этой переменной. При изменении переменной объект будет изменятся автоматически (bindable value доступны от версии Fyne 2.0).

Описание [fyne: bindable value](https://developer.fyne.io/binding/).

Например, если виджет label отображает только одну строковую переменную (которая изменяется где-то в других местах проги), можно привязать строку к виджету. Переменная должна быть специально типа bindable string:

```go
boundString := binding.NewString()
```

 привязать можно:

- при создании объекта используя функцию New...**WithData()**.
  
  ```go
  label := widget.NewLabelWithData(boundString)
  ```

- с помощью метода **Bind()**, если label cоздан как обычно.
  
  ```go
  label.Bind(boundString)  
  ```

предусмотрены методы связать **Bind(str)** и отвязать **Unbind()**. 

Теперь любое изменение переменной автоматически отобразится в label. Не нужно дополнительно передавать строку в виджет label.SetText(), делать обновление label.Refresh(), можно даже не хранить ссылку на объект. Для взаимодействия с переменной предусмотрены методы str.**Set()**, str.**Get()** и str.**Reload()**.

Мой пример: [bindable value example](https://github.com/annettalekto/sandbox_go_fyne/blob/main/bind/main.go)
