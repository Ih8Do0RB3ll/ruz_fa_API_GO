# API расписания Финансового университета при Правительстве Российской Федерации
____

Простой пакет для получения расписания с https://ruz.fa.ru
В теории может быть использован для работы с https://ruz.hse.ru

Небольшой дисклеймер: Я НЕ ЗНАЛ ЧТО ДЕЛАЮ и API описано неполностью
## Установка
___
`
go get github.com/Ih8Do0RB3ll/ruz_fa_API_GO
`
## Примеры

____
###Пример работы с группами

```go
package main

import (
	fa "github.com/Ih8Do0RB3ll/ruz_fa_API_GO"
	"log"
)

func main() {
	// задаем хост к которому будем обращаться, может быть https://ruz.fa.ru либо https://ruz.hse.ru
	host := "https://ruz.fa.ru"

	//ищем группу ЗБ-ПИ19-2
	groupID, err := fa.SearchGroup(host, "ЗБ-ПИ19-2")
	if err != nil {
	}
	// получаем структуру с расписанием группы ЗБ-ПИ19-2 на сегодня
	timetable, err := fa.TimeTableGroup(host, groupID, "", "")
	if err != nil {
	}
	// выводим структуру на печать, конечно лучше с нею поработать перед выводом ¯\_(ツ)_/¯ 
	log.Println(timetable)

	//Ищем группу ЗБ-ПИ19-1
	groupID, err = fa.SearchGroup(host, "ЗБ-ПИ19-1")
	if err != nil {}
	//получаем структуру содержащую расписание группы ЗБ-ПИ19-1 с 17.01.2022 ро 23.01.2022
	timetable, err = fa.TimeTableGroup(host, groupID, "2022.01.17", "2022.01.23")
	if err != nil {}
    //выводим структуру на печать
	log.Println(timetable)
}
```
### Пример работы с преподавателями
```go
package main

import (
	fa "github.com/Ih8Do0RB3ll/ruz_fa_API_GO"
	"log"
)

func main() {
	// задаем хост к которому будем обращаться, может быть https://ruz.fa.ru либо https://ruz.hse.ru
	host := "https://ruz.fa.ru"

	//ищем ID преподавателя
	teacher,err := fa.SearchTeacher(host,"Хасаншин Ильшат")
	//Получаем структуру содержащую расписание преподавателя на сегодня
	timetable1,err := fa.TimeTableTeacher(host,teacher,"","")
	if err!= nil{}
	//выводим стуктуру на печать
	log.Println(timetable1)
}
```
### Функции
___
```go
SearchGroup(host, groupName string) -> (string, error)
TimeTableGroup(host, groupId, date, dateEnd string) -> (TimeTable, error)
SearchTeacher(host, teacherName string) -> (string, error)
TimeTableTeacher(host, teachId, date, dateEnd string) -> (TimeTable, error)
```