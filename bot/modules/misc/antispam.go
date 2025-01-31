package user

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/utils"
	"fmt"
	"sync"
)

const unameLog = `#BAN
<b>User Name:</b> %s
<b>User ID:</b> <code>%v</code>
<b>Chat Name:</b> %s
<b>Chat ID:</b> <code>%v</code>
<b>Link:</b> %s`

func (m Module) antispam(ctx *telegram.TgContext) {
	user := ctx.User

	if !IsBan(m.App.DB, user.Id) {
		return
	}

	dataMap := map[string]string{"1": telegram.MentionHtml(int(user.Id), user.FirstName), "2": utils.IntToStr(int(user.Id))}
	text, keyb := telegram.CreateMenuf("./data/menu/spam.json", 1, dataMap)

	if _, err := ctx.Bot.BanChatMember(ctx.Message.Chat.Id, user.Id, nil); err != nil {
		text += "\n\n🚫 <b>Tetapi saya tidak bisa mengeluarkannya, mohon periksa kembali perizinan saya!</b>"
		ctx.SendMessage(text, 0)
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() { defer wg.Done(); ctx.SendMessageKeyboard(text, 0, keyb) }()
	go func() { defer wg.Done(); ctx.DeleteMessage(0) }()
	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(unameLog,
			telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName),
			ctx.User.Id,
			ctx.Chat.Title,
			ctx.Chat.Id,
			telegram.CreateLinkHtml(telegram.CreateMessageLink(ctx.Chat, ctx.Message.MessageId), "Here"))
		ctx.SendMessage(textToSend, m.App.Config.LogEvent)
	}()
	wg.Wait()
}
