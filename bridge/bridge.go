package bridge

type MessageSenderInterface interface {
	Send(msg string) error
}

type EmailMessageSender struct {
	messages []string
}

func NewEmailMessageSender(messages []string) *EmailMessageSender {
	return &EmailMessageSender{
		messages: messages,
	}
}

func (ms *EmailMessageSender) Send(msg string) error {
	// Ignore send details
	return nil
}

type MobileMessageSender struct {
	message string
}

func NewMobileMessageSender(message string) *MobileMessageSender {
	return &MobileMessageSender{
		message: message,
	}
}

func (ms *MobileMessageSender) Send(msg string) error {
	return nil
}

type Notification interface {
	Notify(msg string) error
}

type ErrorNotification struct {
	sender MessageSenderInterface
}

func NewErrorNotification(sender MessageSenderInterface) *ErrorNotification {
	return &ErrorNotification{
		sender: sender,
	}
}

func (en *ErrorNotification) Notify(msg string) error {
	return en.sender.Send(msg)
}

type NormalNotification struct {
	sender MessageSenderInterface
}

func NewNormalNotification(sender MessageSenderInterface) *NormalNotification {
	return &NormalNotification{
		sender: sender,
	}
}

func (nn *NormalNotification) Notify(msg string) error {
	return nn.sender.Send(msg)
}
