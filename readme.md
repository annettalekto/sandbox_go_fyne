## Заметки по GO Fyne

Часто используемые элементы Fyne. 

### Диалоговые окна

- **Confirm** - диалоговое окно с выбором "Да/Нет".
<img src="dialog/img/dialog_confirm.PNG" alt="screen"/>

- **Custom** - диалоговое окно с возможностью добавить свой элемент CanvasObject (круг метку, цветной текст).
<img src="dialog/img/dialog_custom.PNG" alt="screen"/>

- **Form** - позволяет добавить свой widget (даже несколько) в диалоговое окно. Можно использовать подсказку (поле hint), но работает она совсем не так как ожидалось (просто текст ниже widget).
<img src="dialog/img/dialog_formitem.PNG" alt="screen"/>

- **Information** - обычное окно с информацией и кнопкой.
<img src="dialog/img/dialog_information.PNG" alt="screen"/>


В диалоговых окнах GO Fyne изменить название кнопок (стандартное ok) на любое другое значение можно через методы **SetDismissText** и **SetConfirmText** (если нет этого поля при создании диалога).

моя папка: [sandbox_go_fyne/dialog at main · annettalekto/sandbox_go_fyne · GitHub](https://github.com/annettalekto/sandbox_go_fyne/tree/main/dialog)

нормальное описание: [Dialog List | Develop using Fyne](https://developer.fyne.io/explore/dialogs)
