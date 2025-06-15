# Project_template

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере
работы над решением.

# Задание 1. Анализ и планирование

<aside>

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании и
условия задания. Это нормально.

</aside>

### 1. Описание функциональности монолитного приложения

**Управление счетчиками:**

- Пользователи могут получить список всех сенсоров
- Пользователи могут получить определенный сенсор по id
- Пользователи могут удалить сенсор по id
- Пользователи могут создать сенсор
- Пользователи могут обновить аттрибуты сенсора по id
- Пользователи могут обновить статус сенсора по id

**Мониторинг температуры:**

- Пользователи могут получить значение температуры с сенсора

### 2. Анализ архитектуры монолитного приложения

- Язык программирования: Go
- База данных: PostgreSQL
- Архитектура: Монолитная
- Взаимодействие: Синхронное
- Масштабируемость: Ограничена
- Развертывание: Требует остановки всего приложения

### 3. Определение доменов и границы контекстов

1. Домен: управление умным домом
   1. Контекст: управление датчиками 
   2. Контекст: мониторинг температуры
2. Домен: датчик
   1. Контекст: показание температуры

### **4. Проблемы монолитного решения**

При развертывании пользователи не могут управлять датчиками
и получать текущие показатели температуры. 
При большом количестве датчиков количетсво запросов от пользователей по получению текущей температуре будет сильно расти.
Для решения проблемы с получением пользователями актуальной информации с большого 
количества датчиков для масштабирования придется 
увеличивать количество инстансов всего приложения.

### 5. Визуализация контекста системы — диаграмма С4

```markdown
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml
' uncomment the following line and comment the first to use locally
' !include C4_Context.puml

title Диаграмма контекса

Person(Пользователь, "Пользователь", "Пользователь использующий умный дом")
System(Умный_дом, "Умный дом", "Позволяет управлять отоплением в доме и проверять температуру")

System_Ext(Датчик, "Датчик", "Датчик отслеживающий температуру помещения")
System_Ext(База_данных, "База данных", "Хранит информацию о датчиках")

Rel(Пользователь, Умный_дом, "Управление отоплением/проверка температуры")
Rel_Neighbor(Умный_дом, Датчик, "Получает данные о температуре")
Rel(Умный_дом, База_данных, "Сохраняет/Обновляет/Получает информацию о датчиках")
@enduml
```

```markdown
[PlantUml_Context diagram_C4](https://uml.planttext.com/plantuml/png/ZPD1RzD048Nl_XL3BWIfZI-SE5LK7900LI3EaPCiYSLw7TaRjRs5c81GKIiapaMeucwgcXIQfdzXzX_nTfs449oAKzPlTkPz-tQpa-eWqOD8UdT2rPQ3Zg2UrltqWUydmNwZ6-hUe3L8HTAEbHPADzfnvFTbe6pFTgexCcpvk_UtzwpqykaJFmfIBH98pLtRSQ0RVJlzBWqKMYEC8DqJz2gMCjuFLPTag0G5gaD_bSCanIecM9ECsu6KXvYnP5mVx-bGIq4l3c4UqRn4dJ3e9a44iuGUGqgKqFGmKA-jRxdNy_P4aiQgpbzunbVc8q_mUywv6N5XlxUejgbKkwL6V6dUyCsYQ39pR3xm9Vyaa_6KhyqHbdo1-hHspojPwdosA3NBaZKuMz--S9t0puu_uSACxDWRC-GS8YGpiWWpqCm0SGMcQvWNF2K-dy_X0epacYnvOORpDdU4ATeAEmzIXb86oZbcyz61hlDdLtgFgRyiqCfdRUtRSI0CI_43Warf6KQb5xdWG8XTZjoSB4CgtJ_Xe1FERKmvDY0VCxQsSvrMTSVppMLZHurmR1JUWce8ZzoyWtWCIDUqWCPgR8_yNCXRdiQ62rjUn2BtoglmLxEtZjLf70435Crd8kpsMd5IzKhMRwC4pY3axd7yYSEnL1fnKHfLpj-KzrSSR5n6Mxv3dq_vEtRCodTOIl_W_Fy5x0ZLmH_yDm00)
```

# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся
микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите
диаграммы C4, они и так это покажут.

**Диаграмма контейнеров (Containers)**

```aiignore
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Context.puml
' uncomment the following line and comment the first to use locally
' !include C4_Context.puml

title Диаграмма контейнеров

Person(Пользователь, "Пользователь", "Пользователь использующий умный дом")
System(Персональный_Умный_дом, "Персональный_Умный_до", "Позволяет управлять отоплением в доме и проверять температуру, добавлять новые датчики")
System(Контроллер_умных_домов, "Контроллер_умных_домов", "Агрегирует информацию с умных домов")


System_Ext(Датчик, "Датчик", "Датчик отслеживающий температуру помещения")
System_Ext(Датчик_х, "Датчик_х", "Отслеживание каих лиьо контроллируемых велечин")
System_Ext(База_данных, "База данных", "Хранит информацию о датчиках и пользователях")

Rel(Пользователь, Персональный_Умный_дом, "Управление отоплением/проверка температуры/CRUD операции с датчиками")
Rel_Neighbor(Персональный_Умный_дом, Датчик, "Получает данные о температуре")
Rel_Neighbor(Персональный_Умный_дом, Датчик_х, "Получает данные")
Rel(Персональный_Умный_дом, Контроллер_умных_домов, "Отправляет/Получает необходимые данные")
Rel(Контроллер_умных_домов, База_данных, "Сохраняет/Обновляет/Получает информацию о датчиках/умных домах")
@enduml
```

```markdown
[PlantUml_Container_diagram_C4](https://uml.planttext.com/plantuml/png/hPN1Jjj048RlVefjBWK9uajFFI7HtjeAKN6sGk8cbXmRsHi1jq2QLXK8Mgcd7gWgrBiRS0aa3gym-qRzPvqOnmaHaQf8IJoxC_FDlpFhZL1M0jbgkiOpnwkwhPej6bBkXIzDCwZib-kERBHsMw4TL7rFsfuiL_sckUTMFFPP3sNDTNRDpHVhMspQUVl6R5P2QGSmMPliSI3BUnpzkMXvS6qYX90DMtpmNTVVTxowS1tF5XMl9gQMdI34FryWjN3zQiLr3n4ZOoo6DwGZNLlGT-fJH5UgZUyXFf6WEnfHeeuffbjAACRIY7g6iMK7eU-jq4yy3jGftR2P8hqHpskYzDXIQS6QeBuwell9ekgeC_K5JBT2TO2Jg1FyfMki3qkhnlj3QDLa1c313w3WPa-zqQBBYO-L-gIPdxGvexm17-ESm--OCUwrEZqsgMD67e5-1EW1QaW06zDGK6-CIJ5g4jg9A-NSgPikUGWtbXG2eDQswgobRd-cCoJi2xIOLtdlPwIvetvEWn_fCUaq0qPHRMkYcUgEw-SmNDUJDsiLlkg6YEaAXmD4hGFo9-eZIczpbwXFM3eJwaWyH16vAAk6CGQrNX_85RHPLWGdo3sM2izQNPmMQ_iN1kwLh2NcAYXq-q1txDBdeSupcMQoMwfRI0YBPhWefasFbkSXGaoKDy3FAOso2SdK7Aj4mrI67hSsS3XyKa3vXc0t5B7Yd40Bnzd7Tf6tQwpVklcugaVEW87oFH8nGbyypDJKm9qZB0vdstONJF9oytNvC22JUPWx8cP-9XXq_dcUc9lREw-4zaxjg1C5eTUcwuGdpmDAiTxPJhsnwmTBtGp5fapbwC0GfHsV7OKkQ2ujnV-BODoOYo3ILCjbMEgMm0ZaBpi6C6U9y4xW2qjryNqDz65sJnL1bqZ-w63yGYjrqncO45tWjjGNv0B89qwAENjzHNf2DcolXlVaFm00)
```

**Диаграмма компонентов (Components)**

```aiignore
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml
' uncomment the following line and comment the first to use locally
' !include C4_Component.puml

title Диаграмма компонентов

Container(user_interface, "Мобильное приложение", "Позволяет пользователю использовать умный дом")
System_Ext(датчик, "Датчик")
Container(controller_smart_houses, "Контроллер умных домов", "Агрегирует, сохраняет данные о датчиках")

Container_Boundary(api, "API Application персонального умного дома") {
    Component(круд_датчиков, "Круд датчиков", "MVC Rest Controller", "Позволяет пользователям выполнять круд операции с датчики")
    Component(модуль_опроса_датчиков, "Модуль опроса датчиков", "Опрашивает датчики")
    Component(клиент, "Клиент", "Обеспечивает обмен именформации между персональным умном домом и серверами компании") 

    Rel(круд_датчиков, клиент, "Uses")
    Rel(модуль_опроса_датчиков, клиент, "Uses")
}

Rel(user_interface, круд_датчиков, "Uses", "JSON/HTTPS")
Rel(user_interface, модуль_опроса_датчиков, "Uses", "JSON/HTTPS")
Rel(модуль_опроса_датчиков, датчик, "Uses")
Rel(клиент, controller_smart_houses, "Uses", "JSON/HTTPS")

@enduml
```

```markdown
[PlantUml_Container_diagram_C4](https://uml.planttext.com/plantuml/png/bLJBJjj05DtxAwRPD4X0Dbrr1OqgjLLf4OzkBPCOOcdya3qHHLMbG4kfL5IwO5MXzXTC4aCQal0BCt_KSyV6Ja90IYmvCtVkuvnpxtWb4bOecYutdZXUZJVhDci84KJFAvNGsb_USqIZkTECxB3cUyBsn7BDToi1jpoAMOf4dJixbUgpfNKoRQ-zhRXM9EmG9hFgku7lKKn0-P-ofeTW5mOc6ZRRzJdtznrlZt77ivdbrTd4iHD6MFaCnHdtQnRd1yYHexnVm12Eu3QJvx8dOzbHBRm7U68c-tA4nItUGvd8eJh2gcqOLN2oK3mi4qFJ0SDmrwhPYwmaB-1oYKmtwXH18vamUOkKC61pHLcmJaha-XUMQzd6-qQToKGTCLsA8g-fa8nHCC7-9vCzTNZ_K9qozGK0X-g7_CTabz2M5epD0qZfcYy_YZ9iv7U2cdqgULviuLZme2w541Pi8XUDDHi-g4KKyflO8wX5rV4agfLNLSTPLKAaEVtI0YQo0yGjk87L8WFsaJhMmWxlc7Pfev53eH4R1uhLCT0LyCmNVjEhM-51sGeSL5bTVyrMWu0xDKiulaUo0HLLGSeubxw3l2dIR9DXZKiBx9F1yCideIpx6cxNd002Lbg2z8nDdn7ZjGzLjc5ZvAgvXFCrzqmEc6n31trC0kd6tW62E2az17n3uHxKd0JIWrPJPAWVNKIZcadnr3p84yzaTr4uit7dcNJb7-qGg--mWaJUpCVWzC6hbrwWTABoRPRo4caEDSsJyRHqamRaYFbF5-ehiBMeXvaOsdn5-6TE0MGT53CmAAOLIr9I2zjEvOMzLrpvM5zKa666fhDXyyU6P8hZDYvFfeCEdACXipDzDWnADFszUMniTIJ-tso-VrTvjRMrleayCvFCCouFPfsBu-HdAUEOIZpE_k7FqamOnehjrV59_my0)
```

```aiignore
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml
' uncomment the following line and comment the first to use locally
' !include C4_Component.puml

title Диаграмма компонентов

Container(персональный_умный_дом, "Персональный умный дом", "система умного дома пользователя")
ContainerDb(db, "База данных", "PostgresSQL", "Хранилище данных")

Container_Boundary(api, "API Application контроллер умных домов") {
    Component(круд_датчиков, "Круд датчиков", "MVC Rest Controller", "Позволяет серверам персонального умного дома управлять датчиками")
    Component(контроллер_данных_датчиков, "Контроллер данных датчиков", "Используется для обработки данных самих датчиков")
}

Rel(персональный_умный_дом, круд_датчиков, "Uses", "JSON/HTTPS")
Rel(персональный_умный_дом, контроллер_данных_датчиков, "Uses", "JSON/HTTPS")

Rel(контроллер_данных_датчиков, db, "Read & write to", "JDBC")
Rel(круд_датчиков, db, "Read & write to", "JDBC")


@enduml
```

```markdown
[PlantUml_Container_diagram_C4](https://uml.planttext.com/plantuml/png/dLJBRjD05DtxAovPG2BHUiEALTjq0QY5aBHi8oSU4akJiMLF55K8gGzL2BLIXCG2iM75FaqJDIGa_ONn7t6ktQP9Qq1jBDcJkO_pc6klHTe9TRibhKUUgiksAwYfTH0zj-tGwQmsFDrirzgH2Ek-qaBfrRhViWFfAAvPYRGhlPfTVBPIugtTxIsxvKHQXDYg5lrMu2kk2RZ_4seh5BUmGRefwBqlfT_nL8EafmGvogMbi1T6MFa4S99-tP5o3ptcBEy2MDhJKf3vQaQcPyxZVTn_u-gH6PiP5fUuJytGJED3hFgMLOGc1-1X7h5Xl1yV91azCubFaNLYVbNZ8nIcIpFWBayfPtva9zDDCgN9EMGZRnGV07AOS4bp43mtiwiiPiZSkD45CpEz97qITtE56vARjRnRO_Wl0BtWeW4MKmQCZncfv4Uw4OgeydQB_vgVYGLJs341XSzck5nHM32Wkk6tbUk4Utady52xNdf9wq4WlRgZFL-b3i8tT0HFN31WBZS-lXO2xha2VR08l_dfvCqORXsP0Ji8PV4dCE9-VIRvFOtHxHWhs7vNfBB0633F4FCYmbpg_mm6z1FBkg1oI30PZB0rJ8-TiWxqolJi4y3kPLBRvwPWShfCYMTf1CzkIxlhItNHvMpD6MOk5cMQORv1pNnCG1UwCLTTJfxmOsRE4W5dPeRn6PlHSazK9ngosXUiZvPL5l9Uxy6_ZdKt4X6JVbLvyzf-iRDJgW3Z_W0FC3SJEOL-GBlaXIiBnwN7r0azBV15IhflRXIl5VtTXlzKMzQQK2u-NNy0)
```

**Диаграмма кода (Code)**

```aiignore
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Sequence.puml

Container(Мобильное_приложение, "Мобильное приложение", "Приложение для управления умным домом")

Container_Boundary(Персональный_умный_дом, "API Application персонального умнного дома")
  Component(круд_счетчиков, "Круд счетчиков")
  Component(клиент_контроллера_умных_домов, "Клиент контроллера умных домов", "Клиент для подключения к серверам компании")
Boundary_End()

Container(контроллер_умных_домов, "Сервера_компании", "Контроллер умных домов")

Rel(Мобильное_приложение, круд_счетчиков, "Операции со счетчиками", "JSON/HTTPS")
Rel(круд_счетчиков, клиент_контроллера_умных_домов, "Вызывает CreateSensor()")
Rel(клиент_контроллера_умных_домов, контроллер_умных_домов, "JSON/HTTPS POST CreateSensor")

@enduml
```

```markdown
[PlantUml_Container_diagram_C4](https://uml.planttext.com/plantuml/png/fLHDRzD04BttLomvEHADBvmujKH8BQ8qmk6ySkqLjUIiZRqMujOF0GSGI1nHHSK_2AO1DstSlp3xZtXhfqwBKpuFZdPdvxrvCs-KjLIvYSf6uSgjG7fXvWksgrISthNjn7tUUngetMmdIqNYHL89gNfUDBBZq9K6ivegFmnsxFwTrO49FTbwQ8_SL8a48Uw8PvcGdkZ5Xdwb3m8ta2Anw8XA-a85dUctDCDvmkb2xva0BjzeWc11azkiqyvaop8x9lMu7MTq0kHxfWyhr9Xol5TlTV2SPleDdPkq4XTydUuLeVnUb4dVJLvOT4mJlQVtaJ83JQK6o5FUa9pocWHArWURR3sEmy1pLH19A6w3SVvAvLn3Svth64C5O_re54SI8xTe2lWXdN2zhr-3wm2_1KsHcvl6FzQlhFtQua7xXHc0Fk1L1WwOIudmcT58uwOX_P8tKwchBB1iEPOji8k9vfqMkFO4KobnX2JzpkYUkq9Jz57nvJMhiMTQaLtWOcmjqDQbC_o-zAtkjULQektchZvVBSLRXMhjBRuREeMEnoByy_N-ZQ-VwiMXiNvbn9Yvb3-RZF7KCXyusu_iZU5mu416fU9NvF-r2H-mxj_nvQWENjPFXAk48sGQ9LQtAV_F5VxAmKNVRB3j3A-9COQi2UdZx-S7)
```

# Задание 3. Разработка ER-диаграммы

```aiignore
@startuml

' hide the spot
' hide circle

' avoid problems with angled crows feet
skinparam linetype ortho

entity "Персональный умный дом" as e01 {
  *id : number <<generated>>
  --
  *name : text
  *description : text
  *user_id: number <<FK>>
  *sensor_id: number <<FK>>
}

entity "Пользователь" as e02 {
  *id : number <<generated>>
  --
  *fio : text
  *house_id: number <<FK>>
}

entity "Датчик" as e03 {
  *id : number <<generated>>
  --
  *name : text
  *type : type
  *location: text
  *value: double
  *unit: text
  *status: text
  *last_updated: timestamp
  *created_at: timestamp
}

e01 }|..|| e02
e01 ||..o{ e03

@enduml
```

```markdown
[PlantUml_ER_diagram](https://uml.planttext.com/plantuml/png/dLBBQW8n5DtFLrpS50IbZvs8kEgc7o5nSdL2CycGtD6A8_InwA9TTjzVa5BfI_s6p1yr6InEeT3IrSsvvv9pxgDlYHdAKnb51v08Za09WiqqVU5Oc5XYeDb42mwPqKE9gOMfe0IO6ala41izjJ12fCYU2vKnmrAGGY7DCWHjAD5HX8e4pQ3X7jsgl2oltDPjtD9zbFTkKzwvDoXlt7htTC-UNJU0MS338vX704qltW6LfqCqqEsEKQ5XXBpNysIh5IeKIz7N45vGW1njR4H6GgjQDhTe1eBNlZezg_veMbHM_yGjlfdVLfPVV7noo_BQhGBUMJt-gzMHq3LFYVQkVX5-26BbhNjrxpknat_EfLgAXpu4A7NCmeps1HCcS-m0rxbVTJKp9MXF-vkXtEwnP9O6USQ3gi-A57r5cWKgDXYo0qPr8hJbjxeesksY25EhOE6XdeU-egYFYeUh_0G0)
```

# Задание 4. Создание и документирование API

### 1. Тип API

Для общения между сервисами будет использоваться REST API так как вызовы API планируются без сохранения
состояния (statelessness) и для использовании многоуровней системы. Способ передачи информации будет на основе JSON так он 
структурирован и понятен человеку.

### 2. Документация API

```markdown
[Swagger_personal_smart_house](./api/swagger/personalSmartHouse/swagger_personal_smart_house.yaml)
```

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1) сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе
   /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```

2) Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

3) Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска
   postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.


