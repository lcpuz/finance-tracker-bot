package lang

var (
	// Welcome message in different languages
	WelcomeMessage = map[string]string{
		"uz": "👋 Assalomu alaykum. Finance Tracker Botga xush kelibsiz!\n\nMen sizning moliyangizni osonlik bilan boshqarishda yordam beraman. Xarajatlaringizni yozish, daromadlaringizni ro'yxatga olish yoki hisobotlar tuzishni istaysizmi? Mana, men sizga yordam bera olishim mumkin.\n\nFinanc eTracker Bot bilan qilishingiz mumkinlar:\n\n💸 Xarajatlarni qo'shish\n💰 Daromadlarni qo'shish\n📊 Hisobotlar tuzish\n📆 Budjet o'rnatish\n📅 Eslatmalar\n📈 Diagrammalar ko'rish\n🗂 Kategoriyalarni o'rnatish\n✅ Va boshqalar!",
		"ru": "👋 Добро пожаловать в Finance Tracker Bot!\n\nЯ здесь, чтобы помочь вам легко отслеживать ваши финансы. Хотите записать ваши расходы, зафиксировать доходы или создать отчеты для лучшего управления своими деньгами? У меня есть все необходимое для этого.\n\nВот некоторые функции, которые вы можете использовать с Finance Tracker Bot:\n\n💸 Добавление расходов\n💰 Добавление доходов\n📊 Создание отчетов\n📆 Установка бюджета\n📅 Напоминания\n📈 Диаграммы\n🗂 Категории\n✅ И многое другое!",
		"en": "👋 Welcome to Finance Tracker Bot!\n\nI'm here to help you keep track of your finances effortlessly. Whether you want to log your expenses, record your income, or generate reports to better manage your money, I've got you covered.\n\nHere are some things you can do with Finance Tracker Bot:\n\n💸 Add Expenses\n💰 Add Income\n📊 Generate Reports\n📆 Set Budgets\n📅 Reminders\n📈 Charts\n🗂 Categories\n✅ And much more!",
		"ar": "👋 مرحبًا بك في Finance Tracker Bot!\n\nأنا هنا لمساعدتك في تتبع أمورك المالية بسهولة. سواء كنت ترغب في تسجيل نفقاتك، أو تسجيل دخلك، أو إنشاء تقارير لإدارة أموالك بشكل أفضل، لدي كل ما تحتاجه.\n\nإليك بعض الأشياء التي يمكنك القيام بها باستخدام Finance Tracker Bot:\n\n💸 إضافة مصاريف\n💰 إضافة دخل\n📊 إعداد التقارير\n📆 تعيين الميزانيات\n📅 تذكير\n📈 رسوم بيانية\n🗂 تصنيفات\n✅ وأكثر!",
	}

	// Main menu buttons
	Income = map[string]string{
		"uz": "⬇️Kirim",
		"ru": "⬇️Доход",
		"en": "⬇️Income",
		"ar": "⬇️دخل",
	}

	Spendings = map[string]string{
		"uz": "⬆️Chiqim",
		"ru": "⬆️Расход",
		"en": "⬆️Spendings",
		"ar": "⬆️المصروفات",
	}

	Report = map[string]string{
		"uz": "📊Hisobot",
		"ru": "📊Отчет",
		"en": "📊Report",
		"ar": "📊تقرير",
	}

	Settings = map[string]string{
		"uz": "⚙️Sozlamalar",
		"ru": "⚙️Настройки",
		"en": "⚙️Settings",
		"ar": "⚙️الإعدادات",
	}

	// Report menu buttons
	ReportForDay = map[string]string{
		"uz": "📈Kunlik hisobot",
		"ru": "📈Отчет за день",
		"en": "📈Report for day",
		"ar": "📈تقرير ليوم",
	}

	ReportForWeek = map[string]string{
		"uz": "📈Haftalik hisobot",
		"ru": "📈Отчет за неделю",
		"en": "📈Report for week",
		"ar": "📈تقرير لأسبوع",
	}

	ReportForMonth = map[string]string{
		"uz": "📈Oylik hisobot",
		"ru": "📈Отчет за месяц",
		"en": "📈Report for month",
		"ar": "📈تقرير لشهر",
	}

	ReportForYear = map[string]string{
		"uz": "📈Yillik hisobot",
		"ru": "📈Отчет за год",
		"en": "📈Report for year",
		"ar": "📈تقرير لعام",
	}

	// Settings menu buttons
	Currencies = map[string]string{
		"uz": "💵Valyutalar",
		"ru": "💵Валюты",
		"en": "💵Currencies",
		"ar": "💵العملات",
	}

	Remainder = map[string]string{
		"uz": "🔔Eslatma",
		"ru": "🔔Напоминание",
		"en": "🔔Remainder",
		"ar": "🔔تذكير",
	}

	Language = map[string]string{
		"uz": "🌐Til",
		"ru": "🌐Язык",
		"en": "🌐Language",
		"ar": "🌐لغة",
	}

	// Currencies menu buttons
	USD = map[string]string{
		"uz": "USD",
		"ru": "USD",
		"en": "USD",
		"ar": "USD",
	}

	EUR = map[string]string{
		"uz": "EUR",
		"ru": "EUR",
		"en": "EUR",
		"ar": "EUR",
	}

	RUB = map[string]string{
		"uz": "RUB",
		"ru": "RUB",
		"en": "RUB",
		"ar": "RUB",
	}

	UZS = map[string]string{
		"uz": "UZS",
		"ru": "UZS",
		"en": "UZS",
		"ar": "UZS",
	}

	// Remainder menu buttons
	EveryHour = map[string]string{
		"uz": "Har soat",
		"ru": "Каждый час",
		"en": "Every hour",
		"ar": "كل ساعة",
	}

	EverySixHours = map[string]string{
		"uz": "Har 6 soatda",
		"ru": "Каждые 6 часов",
		"en": "Every 6 hours",
		"ar": "كل 6 ساعات",
	}

	EveryTwelveHours = map[string]string{
		"uz": "Har 12 soatda",
		"ru": "Каждые 12 часов",
		"en": "Every 12 hours",
		"ar": "كل 12 ساعة",
	}

	EveryDay = map[string]string{
		"uz": "Har kuni",
		"ru": "Каждый день",
		"en": "Every day",
		"ar": "كل يوم",
	}
)
