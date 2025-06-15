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
System(Kafka, "Брокер сообщений", "Агрегирует информацию со светчиков через персональный умный дом")
System(Сервис_сценариев_умного_дома, "Сервис сценариев умного дома", "Обработка тригеров для запуска сценария")
System(Управление_пользователями, "Сервис управления пользователями", "Обработка пользовательских данных")
System(Api_gateway, "Api gateway", "Дробление запросов и распределение по сервисам")


System_Ext(Датчик, "Датчик", "Датчик отслеживающий температуру помещения")
System_Ext(Датчик_х, "Датчик_х", "Отслеживание каих лиьо контроллируемых велечин")
System_Ext(База_данных, "База данных", "Хранит информацию о датчиках и пользователях")

Rel(Пользователь, Персональный_Умный_дом, "Управление отоплением/проверка температуры/CRUD операции с датчиками")

Rel(Персональный_Умный_дом, Api_gateway, "api call")



Rel_Neighbor(Персональный_Умный_дом, Датчик, "Получает данные о температуре")
Rel_Neighbor(Персональный_Умный_дом, Датчик_х, "Получает данные")
Rel(Api_gateway, Контроллер_умных_домов, "Отправляет/Получает необходимые данные")
Rel(Контроллер_умных_домов, База_данных, "Сохраняет/Обновляет/Получает информацию о датчиках/умных домах")
Rel(Api_gateway, Kafka, "Получение информации со счетчиков и обработка сценариев")
Rel(Kafka, Сервис_сценариев_умного_дома, "Подписка на данные со счетчиков")
Rel(Сервис_сценариев_умного_дома, Контроллер_умных_домов, "Активация датчика/элемента умного дома")
Rel(Api_gateway, Управление_пользователями, "Регистрация/Активация персонального умного дома")
Rel(Управление_пользователями, База_данных, "Сохраняет/Обновляет пользовательскую информацию")
@enduml
```

```markdown
[PlantUml_Container_diagram_C4](https://uml.planttext.com/plantuml/png/hLTDJnDH5DttLppPg0ao6rUkCEZAZnYDwya08qoSJaax3RAZL3G654XOkL29SJyKHqfforzuxnzvxdtpFTE6eY4fdVVUlUVSZ_FkTB7bMqs_NNThjnnlnMslscBTzpTQ3moZQMqkh3d-UdkvtRARAmtFjpr_OQLHDpPSooERUor_rNMMZQNxzrxGqkjdJuswrVBj9fRC9R9uvoziaFVRekt1j0uNmb-tnPk6wpOs7MzDk8vd2yjR5Rbjfzd2juO0j70RAvRhRi57mh7elkOxlciBUIpxCf1dQXkVG_m5GWxaM8xKZWpbXHp945jZsQlLNjZDLiEx8x_ZyLBjotDQbW4VnFEyc9kqDJTbJyY-wiYhU5DrrHVr2PmkXEg2paZjuQlyZVtXtDtQgotagauSG0js80hUPCa7JVapjZ6rZKQ-rk64vJdu4Pq3sEqGZIlEJe-Mr0vH7eFz6AGl4SC8P4Cv5B8NqPGXOX9iH94IjZRZa8SmevGY0OXrMtNdjTbf7c54jg0MqYwT_GYOWUndSl1LbubXBecAsZRZdAdTA7voGt5T-p1duP0R8fHdA0ueSXw0FrBl4VesTOdwWAqlGdL4waLal2GqdrXltbhayuYn1uGik1Y86PNM2RoO5LQEwOEo6wS6g09VG_PmBdIQorMVsbOdRDwZbZHL1s34Bi0IbRWN9GmsPt8SfomWuXaxKROJMRiuIG57_0qveDuwvPOQG7rSepvEHhgZy-WAWQuCq5DTHCJ7SYZg81F4pxHVu_OqKu7br7S08lrI0DsY0uAVuc1291DCzZc2lcwN0EwvTrB-3pSSSytoxKrh2txn9A8dXZdch9oceSM9mJ8NkqVYOtrrU3K4IfWzVgLR9ucNxZsWro9yy_4x_mvknKHpZ9e-pXMU-JA05qBuWmKADBd1AWMl4uDh8kx-JFLAwARQBG1YHIUy21i7YDphz5xYtpvrtA0W_aXTSgXLszC9OlUZ0fKZE4E2pMojMCrwFLz3slZ5NK1HJT2kba1wfGL4eJ-nmRWrQYzjTyhWcMqSLEYZyaOtibTufCsAUkuPIozVFn9ihTSH9_LWfn0dIoMDvTgCyuAm80WQxjoovCjyRZjhwykDvanEYnskSzl5Gg2ltQIkd9tAq6K84l-5GTJbqqXegFpjCDi8X62oavmGZ38aNdXuCktYypUu3PCXN60o2_X46PsWyNQrOc96k4Zrz9z2yfgwCigpEM0zbJAPJEaKAf562OoRcszHdhVP0Kn27fTdMcaaHXmYs1lFtUzSf2k-nddaZ79N4jXF81hNy6R0i_NTeHmW3teoVD33D5SdGttcQsT8u3WPJ7XVg1B0JMRzZ-WDgyC1H9oCAfeLhr0HgMaCRq3gd_GnzKM3VazKwWLa5srl5J-5_W80)
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
Container(api_gateway, "Api gateway", "")

Container_Boundary(api, "API Application персонального умного дома") {
    Component(круд_датчиков, "Круд датчиков", "MVC Rest Controller", "Позволяет пользователям выполнять круд операции с датчики")
    Component(модуль_опроса_датчиков, "Модуль опроса датчиков", "Опрашивает датчики")
    Component(клиент, "Клиент", "Обеспечивает обмен именформации между персональным умном домом и серверами компании") 
    Component(producer, "Producer", "Отправляет информцию со счетчиков")
    Component(репозиторий, "Репозиторий (локальное хранилище)", "Хранение последних данных со счетчиков") 

    Rel(круд_датчиков, клиент, "Uses")
    Rel(модуль_опроса_датчиков, клиент, "Uses")
    Rel(модуль_опроса_датчиков, репозиторий, "Uses")
    Rel(модуль_опроса_датчиков, producer, "Uses")
    Rel(producer, клиент, "Uses")
}

Rel(user_interface, круд_датчиков, "Uses", "JSON/HTTPS")
Rel(user_interface, модуль_опроса_датчиков, "Uses", "JSON/HTTPS")
Rel(модуль_опроса_датчиков, датчик, "Uses")
Rel(клиент, api_gateway, "Uses", "JSON/HTTPS")

@enduml
```

```markdown
[PlantUml_Container_diagram_C4](https://uml.planttext.com/plantuml/png/hLNBRjD05DtxAoxPa4gqsR1YrH8X0Q8GzS7MSXEtjUJObkseH0Yf3qf1L7JJ5Qg03NitZLitJT9VcFaZpXsxSPAcHH6qqcJcpdsSS-REU3u83Jzir6tjWULKxKRDfCqmz88dfP9lRCrjME5cOwqHc7xLTKBJ2UUgRhtasOR3CRD1MBEjjLBvyMo5JQkBhqfr8mXD7oQzxDOzr-4OZ_C_f8Q3u3eC56wQjExQjhjbEHjaMuv9XbEZaMtB3p1p2SN9Tgk6RJUHOu3oTW4jj4BR974i4X69CxcDiOl_Y4H7z36vnjWJiUZ9NSnQcbO69mF5_I8pr2qmzDUDglc82k84BgV8T2KF4TGNCObhf8G1YtFEWdbSODUVi5o85iOhUIHYkKkg54TUS24HeM2CzJSIYTovlIaFIUu1O4z-5PSaseosCACjDo5bNN_sBYp2ndu7gDdXaiVv4euv3yEpz0qZDBUC9jmMF8ko5GE5P-wgFtKRJitmcnp3lfKNjE1vjbKrGijrc48iju6MDOi6Cfo9VeOrMwHeHLIOeVSQuMzmAaNHGO8zqTQ7qJDdPl0ztQFnFGQw-BPCIoQEdz7ww1FJdqxe8z4bqOAQQXkw7YcHRm0Hh8eU0ZwXS49oPnH80gt6o7GHr4OqgkaSZ-W-v8aciZl9dMdOUI9TyKCvHF8ph22HCliBd0vu9MapApdpPPRo56bs5Cs3uRJSrLrsH2-c4_aHsBRv33CnbFcSyK_i0iZQpNkWUzC1QifAAc5RgRom9_drYzIb0HaQO-FvRgrHDNqWhsJJb0NEKqc3V8CZJtB831YtYl5X0D5uHDjnqH0IgvQuGDYkIf68IoxqQz865TLbxmorVqnoN-7fgPSWaLz4FAEW_ixCsSEWEXEuhc1eiqdkgoF5509YSITcqXJg9TE-xmADdVzgO0OPNHKuHRF-lqntoliFEOSwOonBlZEPmGTDOxVnb_s-Hqb5ulVbyflNfUShAvLbv9cOP9h7uCwiKvtIw0SWuvWsoJ3xqG_0nDBQlEdKyC7y0m00)
```

```aiignore
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Component.puml
' uncomment the following line and comment the first to use locally
' !include C4_Component.puml

title Диаграмма компонентов

Container(персональный_умный_дом, "Персональный умный дом", "система умного дома пользователя")
Container(сервис_сценариев_умного_дома, "Сервис сценариев умного дома", "обработка сценариев")
ContainerDb(db, "База данных", "PostgresSQL", "Хранилище данных")

Container_Boundary(api, "API Application контроллер умных домов") {
    Component(круд_датчиков, "Круд датчиков", "MVC Rest Controller", "Позволяет серверам персонального умного дома управлять датчиками")
    Component(контроллер_данных_датчиков, "Контроллер данных датчиков", "Используется для обработки данных самих датчиков")
    Component(компонент_вкл_выкл_приборов, "Вкл/выкл приборов", "Используется для вкл выкл приборов персонального умного дома")
}

Rel(персональный_умный_дом, круд_датчиков, "Uses", "JSON/HTTPS")
Rel(персональный_умный_дом, контроллер_данных_датчиков, "Uses", "JSON/HTTPS")

Rel(контроллер_данных_датчиков, db, "Read & write to", "JDBC")
Rel(круд_датчиков, db, "Read & write to", "JDBC")
Rel(сервис_сценариев_умного_дома, компонент_вкл_выкл_приборов, "Uses")
Rel(компонент_вкл_выкл_приборов, персональный_умный_дом, "Uses")


@enduml
```

```markdown
[PlantUml_Container_diagram_C4](https://uml.planttext.com/plantuml/png/dLLDRzD04BtxLomvG2hH-iA9Kzlq0AY5q0-kbXClYQMDRTaRHHL2IbiLXAXKWJWWBeW3zpHDr92Gv2-i_n5lracQ3nSrYQLujJllpNipOtipadOegnLXF72zegWwd9Mb3AAdfXdQjOsIAylLGZNYOT7t9FVaHj6lc86mFSAiHz8HRi7CFLdFqzRHtgvPiIF9GsnPERyI-1vX0k9_nAeUm1LiC5dcxAqlX5zplH8JhiUPxJbivjWD8wnyXkHC-4LRY6DmJ5Gk9Z2aAmLdwflgg8QwZklu_ukhmLHN3R0Ouh-lsgeVdsBLD8mSFDb87cPnrexhyOcEQAXUV86ep-gF5Py1c2nLYrWUiupwcHxCReDP4fn1DE8wyGbIjhMM90Q7rsemYYA5f8s8RaYPQkZmNdoPMPiI2IRanJ7uB51-82-8hSDmMpMjQMfhJ4rwVztYs28kNH99n-fAb_5A3Q2eI_8NuDCITmfPfq30Bm3TaAyM5dsgINnEZ7a_agMGHmUlTtM2tvg-3wOUr7rIxLd4sbIFh6s_wZbsU9or0nVOhVnpjXK4mYtQqlMzfCbeBHXHIbpmFEb8V3usDY35x9t1y9iCK5PrqT0prQAoeVZnHoWYlYQ9_96SiVapSh3t9iVsEIQLT8OOQHvcaX4Pe01DtTLBI3bba-wraybaQJCtQaBwa61tgB5D8eMIYrb9DEuTr6pUsc9ThEagftjEAUOqABKOwZlSJ2OPSk4R0tD9mJswpSzJPvOJQEqXZJxLrimpRQ4kNTN33K-WNWprZ5Ao-jZNLpenno5iFkHUBfe9z2wEfHiBR-yDOv-BfLv3_nlPeuX7PEN5mQkNvhF3m_m1SYoVO8N1ISsSf5w1JhzCzhdji8Ui5hgIuuEWsNUsSsD7TvVX7kWLNwch3P-kpTfjCPObMEeZDKfc69lSS_1v_0S0)
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

Rel(контроллер_умных_домов, клиент_контроллера_умных_домов, "JSON/HTTPS Response")
Rel(клиент_контроллера_умных_домов, круд_счетчиков, "Результат CreateSensor()")
Rel(круд_счетчиков, Мобильное_приложение, "Операции со счетчиками", "JSON/HTTPS")

@enduml
```

```markdown
[PlantUml_Container_diagram_C4](http://uml.planttext.com/plantuml/png/hLNBRjD05DtdAuQi4ed66rQiMY8aG42Z9gn7RZngBJbZOu-5sBMfW0K899O8KVY3u3OqJHlt5-xy4MViuAGu2Nsmi3Lpxptddlj8l1ufErHnpwlSSMNNYnt1TfKAelkM5Tgl6ZkksestuqY4NLyg8LMZw_UimBEbmQn5ol7SRQjvRwrbJ2-UFRLwTgH421DlYvUna5tH20nzfGa2svKYhD5NIkadZUXCVw09paDE5thF673vHKCOHpIyowhbIBOeicf23yjsHiT0VcBw84CDA84_y-N6SvheztHkmb9Sy5JhSqBv0p-MZXs-hj4X3VMUtaV813IP6Y1FUK5oodCIADbeFMORGU2vNLkvleJYCXZd8qgd6ehhLCS0AXXh-ht0bsXvZSQ07z0nr_lw7RZwU8zeZDZ45FubzxAojyI3yaUc0RhFimWSq9SKvZEZaGP5GVeDBxgIPvbXsM8ics5d7Ksg9N0-4tGbnH6Iz4UZUpeL6gEEZ2_9MSrunXdP1Ivch2EKzMSo_A5qQlLBoxL0s_AgViodugL4kVOItv9AeMDBU5TVxt_CzLk-E3JGRuqOqvVqxo6ZFRdC9-tDvzQZJgVLXenCnIhoMst2Pwpx2Pu4sS7Bcg6mbMWB6VbXhLwalt66QqrmLZThRROxbyGK0xa6uMqwCwTbIqJuqKNYVtHZvP9y1-a9_DWrtSUgB1t7gcsu-l_npNUoiYwaWm_1Rm00)
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

entity "Сценарии" as e04 {
  *id : number <<generated>>
  --
  *name : text
  *description : text
  *user_id: number <<FK>>
  *house_id: number <<FK>>
  *sensor_id: []numbers <<FK>>
  *logic: text
}

entity "Комната" as e05 {
  *id : number <<generated>>
  --
  *name : text
  *house_id: number <<FK>>
  *sensor_id: []numbers <<FK>>
}

e01 }|..|| e02
e01 ||..o{ e03
e01 ||..o{ e04
e01 ||..o{ e05

@enduml
```

```markdown
[PlantUml_ER_diagram](https://uml.planttext.com/plantuml/png/lLJDIYGn4BxtKnHU145HLoyYuicBZo3B43lbTDXqqYJLkk88hWlhGO-Ul8XlC8YYwzyp96zafMsP7b1Mvh2dvAklI7rVLNMlUnAEwb9dsJGKIY9GWU0hI-yuLoxNc6YnQvM4ojbjZQM7FKK52DFNA25tTi_33Y9b_eSobN2Y1AqCqdw5O1qLDilGaA9zc0gNuJOUnL_X9Jo7ONYCP-4vdeQ_47-7f_OQRfXzcWBX0USNu203cE7YAs3gSXiThAxsqQ0JX79jZSdPsPHXH8cSG_YJ4fJeSwSgKjPqehL7rrEoyzJ6PlF6Z4VZxKVSuPZubqRo7P_NOHYFmst2hTJ5hqhTKRQZgR2iwX-5prEnU1BkmqDRxDk4VMc6mf2F1BNDHUhHA65Nw1fNGDgQHztqp2WQqRmpLFiHri9JhwvagifHLI9db5MYSeSfsXFK9Rgshk8VxY0lGZnYQ_UjjQN_DFBF6Z--3LlVtrZVeRNjgxmjqVLpaHQtyNCSXgsRvGdTJAWkoU6Fvd0mDpSOfALiu82XFKXhCmwNnk5obgsZaVnBU0K0)
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


