package main

func (app *Config) SendMail(msg Message) {
	app.Wait.Add(1)
	app.Mailer.MailerChan <- msg
}
