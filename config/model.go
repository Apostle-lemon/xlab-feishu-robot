package config

type Config struct {
	Feishu struct {
		AppId             string
		AppSecret         string
		VerificationToken string
		EncryptKey        string
		LarkHost          string
	}
	Server struct {
		Port int
	}
}
