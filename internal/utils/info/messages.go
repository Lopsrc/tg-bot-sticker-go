package info

const support = "@Lopsrc" // Name of the support.

const(
	MsgHelp= `	
Create sticker set - create a set of stickers.
Add sticker to set - add a sticker to the set.
Get sticker set - get a set of stickers.
Delete sticker - delete a set of stickers.
Help - get background information.

For questions about the bot's operation, contact support - ` + support + `.`
MsgStart = `
Welcome.

I can create stickers, as well as add, return and delete them from the set.

Select an action:

Create sticker set - create a set of stickers.
Add sticker to set - add a sticker to the set.
Get sticker set - get a set of stickers.
Delete sticker - remove the sticker from the set.
Help - get background information.

Follow the commands carefully, if an error occurs, then try to change some data.`

	MsgCreateStickerSet = `Come up with and enter the name of the sticker in English without spaces. Only letters and numbers are allowed.

Example: stickername`
	MsgAddSticker = `Enter the name of the sticker. 

Example: stickername`
	MsgDeleteOrGet = `Enter the name of the sticker.

Example: stickername`
	MsgEmoji = `Enter the emoji for the sticker.

Example: ❤️`
	MsgAddEmoji = `The sticker was added successfully. You can add another sticker, just send the emoji and follow the instructions below.`
MsgDeleteSticker = `Send the sticker from the set.`
MsgUploadPhoto = `Send a photo with the PNG or JPEG extension. Disable photo compression when sending PNG (if you do not disable compression, the sticker will not be created on a transparent background).`
	MsgUpGet = `Press the Done button to get the stickers.`

	MsgDoneCreate = `A set of stickers has been created.`
	MsgDoneAdd = `The sticker has been added successfully. To add more stickers, send emojis.`
	MsgDoneGet = `The stickers have been successfully received.`
	MsgDoneDelete = `The sticker has been successfully removed.`

	MsgErrPayload = `You entered the data incorrectly, please read the requirements carefully and enter the data again.`
MsgErrFile = `The photo you sent is not supported, please re-read the requirements and send the correct file again.`
MsgErrStickerSetExist = `You have entered an incorrect name for a set of stickers. Such a name already exists. Try again.`
MsgInternalError = `An internal error. Something went wrong inside the bot. Write to us about it - ` + support + `.`
	MsgErrStickersNotFound = `The stickers were not found. You may have entered the name incorrectly or the stickers have been deleted. Try again.`
)

// const(
// 	MsgHelp= `	
// Create sticker set - создать набор стикеров.
// Add sticker to set - добавить стикер в набор.
// Get sticker set - получить набор стикеров.
// Delete sticker - удалить набор стикеров.
// Help - получить справочную информацию.

// По вопросам работы бота обращаться в поддержку - ` + support + `.`
// 	MsgStart = `
// Добро пожаловать.

// Я могу создавать стикеры, а также добавлять, возвращать и удалять их из набора.

// Выберите действие:

// Create sticker set - создать набор стикеров.
// Add sticker to set - добавить стикер в набор.
// Get sticker set - получить набор стикеров.
// Delete sticker - удалить стикер из набора.
// Help - получить справочную информацию.

// Внимательно следуйте командам, если возникнет ошибка, то попробуйте изменить некоторые данные.`

// 	MsgCreateStickerSet = `Придумайте и введите название стикера на английском языке без пробелов. Разрешено использовать только буквы и цифры.

// Пример: stickername`
// 	MsgAddSticker = `Введите название стикера. 

// Пример: stickername`
// 	MsgDeleteOrGet = `Введите название стикера.

// Пример: stickername`
// 	MsgEmoji = `Введите эмоджи для стикера.

// Пример: ❤️`
// 	MsgAddEmoji = `Стикер добавлен успешно. Вы можете добавить еще стикер, просто отправьте эмоджи и следуйте дальнейшим инструкциям.`
// 	MsgDeleteSticker = `Отправьте стикер из набора.`
// 	MsgUploadPhoto = `Отправьте фотографию с расширением PNG или JPEG. Отключите сжатие фотографий при отправке PNG(если не отключить сжатие, то стикер не будет создан на прозрачном фоне).`
// 	MsgUpGet = `Нажмите кнопку Done, чтобы получить стикеры.`

// 	MsgDoneCreate = `Набор стикеров создан.`
// 	MsgDoneAdd = `Стикер успешно добавлен. Чтобы добавить еще стикеры, пришлите эмоджи.`
// 	MsgDoneGet = `Стикеры успешно получены.`
// 	MsgDoneDelete = `Стикер успешно удален.`

// 	MsgErrPayload = `Вы ввели данные не правильно, пожалуйста, прочитайте внимательнее требования и введите данные заново.`
// 	MsgErrFile = `Фото, которое вы прислали, не поддерживается, пожалуйста, перечитайте требования и снова отправьте корректный файл.`
// 	MsgErrStickerSetExist = `Вы ввели некорректное имя для набора стикеров. Такое имя уже существует. Повторите попытку снова.`
// 	MsgInternalError = `Внутренняя ошибка. Внутри бота что-то пошло не так. Напишите нам об этом - ` + support + `.`
// 	MsgErrStickersNotFound = `Стикеры не найдены. Возможно вы неправильно ввели имя или стикеры удалены. Повторите попытку снова.`
// )
