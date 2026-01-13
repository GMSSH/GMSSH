package gi18n

func init() {
	i18nAdapter = NewI18nFileAdapter()
	i18nManager = NewI18nManager()
}

func Instance() *I18nManager {
	return i18nManager
}
