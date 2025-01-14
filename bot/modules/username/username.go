package username

import (
	"SiskamlingBot/bot/core/telegram"
	"SiskamlingBot/bot/models"
	"fmt"
	"regexp"
	"strconv"
	"sync"
)

const (
	unameLog = `#USERNAME
<b>User Name:</b> %s
<b>User ID:</b> <code>%v</code>
<b>Chat Name:</b> %s
<b>Chat ID:</b> <code>%v</code>
<b>Link:</b> %s`

	unameMsg = "⚠ <b>%v</b> [<code>%v</code>] telah dibisukan karena belum memasang <b>Username!</b>"
)

func (m Module) usernameScan(ctx *telegram.TgContext) {
	// if core.IsUserRestricted(ctx) {
	// 	 return
	// }

	newUsername := models.NewUsername(ctx.User.Id, ctx.User.Id, true)
	models.SaveUsername(m.App.DB, newUsername)

	if !ctx.RestrictMember(0, 0) {
		unavailable := unameMsg + "\n\n🚫 <b>Tetapi saya tidak bisa membisukannya, mohon periksa kembali perizinan saya!</b>"
		textToSend := fmt.Sprintf(unavailable, telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName), ctx.User.Id)
		ctx.SendMessage(textToSend, 0)
		return
	}

	var wg sync.WaitGroup
	wg.Add(3)

	go func() { defer wg.Done(); ctx.DeleteMessage(0) }()

	go func() {
		defer wg.Done()
		textToSend := fmt.Sprintf(unameMsg, telegram.MentionHtml(int(ctx.User.Id), ctx.User.FirstName), ctx.User.Id)
		ctx.SendMessageKeyboard(textToSend, 0, telegram.BuildKeyboardf("./data/keyboard/username.json", 1, map[string]string{"1": strconv.Itoa(int(ctx.User.Id))}))
	}()

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

func (m Module) usernameCallback(ctx *telegram.TgContext) {
	pattern, _ := regexp.Compile(`username\((.+?)\)`)
	if !(pattern.FindStringSubmatch(ctx.Callback.Data)[1] == strconv.Itoa(int(ctx.Callback.From.Id))) {
		getUsername := models.GetUsernameByID(m.App.DB, ctx.Callback.From.Id)
		if getUsername != nil && getUsername.ChatID == ctx.Callback.Message.Chat.Id {
			if ctx.User.Username == "" {
				ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
				return
			}

			models.DeleteUsernameByID(m.App.DB, ctx.Callback.From.Id)

			ctx.UnRestrictMember(0)
			ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
			return
		}

		ctx.AnswerCallback("❌ ANDA BUKAN PENGGUNA YANG DIMAKSUD!", true)
		return
	}

	if ctx.User.Username == "" {
		ctx.AnswerCallback("❌ ANDA BELUM MEMASANG USERNAME", true)
		return
	}

	models.DeleteUsernameByID(m.App.DB, ctx.Callback.From.Id)

	ctx.UnRestrictMember(0)
	ctx.AnswerCallback("✅ Terimakasih telah memasang Username", true)
	ctx.DeleteMessage(0)
}
