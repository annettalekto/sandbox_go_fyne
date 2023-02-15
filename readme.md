##### Заметки по GO Fyne

Часто используемые элементы Fyne. 

**Диалоговые окна**

- **Confirm** - диалоговое окно с выбором "Да/Нет".
<img src="dialog/img/dialog_confirm.png" alt="screen"/>
  ![Confirm](./dialog/img/dialog_confirm.png)

- **Custom** - диалоговое окно с возможностью добавить свой элемент CanvasObject (круг метку, цветной текст).
  ![Custom](./dialog/img/dialog_custom.png)

- **Form** - позволяет добавить свой widget (даже несколько) в диалоговое окно. Можно использовать подсказку (поле hint), но работает она совсем не так как ожидалось (просто текст ниже widget).
  ![FormItem](./dialog/img/dialog_formitem.png)

- **Information** - обычное окно с информацией и кнопкой.
  ![Information](./dialog/img/dialog_information.png)

В диалоговые окна GO Fyne переименновать кнопки по русски (изменить стандартное ok на любое другое значение) можно через методы **SetDismissText**  **SetConfirmText** (если нет этого поля при создании диалога).

моя папка: [sandbox_go_fyne/dialog at main · annettalekto/sandbox_go_fyne · GitHub](https://github.com/annettalekto/sandbox_go_fyne/tree/main/dialog)

нормальное описание: [Dialog List | Develop using Fyne](https://developer.fyne.io/explore/dialogs)
