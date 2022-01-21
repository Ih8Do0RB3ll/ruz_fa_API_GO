package ruz_fa_API_GO

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var host = "https://ruz.fa.ru"

func dateNow() string {
	return time.Now().Format("2006.01.02")
}

func SearchGroup(groupName string) (string, error) {
	// поиск группы по названию, возвращает ID группы.
	url := fmt.Sprintf("%s/api/search?term=%s&type=group", host, groupName)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result Group
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	id := result[0].ID
	return id, err
}

func TimeTableGroup(groupId, date, dateEnd string) (TimeTable, error) {
	// отдает структуру timeTable(расписание) по ID группы ¯\_(ツ)_/¯
	if date == "" || dateEnd == "" {
		date = dateNow()
		dateEnd = dateNow()
	}
	var result TimeTable
	url := fmt.Sprintf("%s/api/schedule/group/%s?start=%s&finish=%s&lng=1", host, groupId, date, dateEnd)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	// log.Println(string(body))
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, err

}

func SearchTeacher(teacherName string) (string, error) {
	// поиск препода по его ФИО, отдает ID первой записи, так что лучше кормить ФИО полностью
	teacherName = strings.ReplaceAll(teacherName, " ", "%20")
	url := fmt.Sprintf("%s/api/search?term=%s&type=person", host, teacherName)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	//	log.Println(string(body))
	if err != nil {
		return "", err
	}
	var result Teacher
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	id := result[0].ID
	return id, err
}

func TimeTableTeacher(teachId, date, dateEnd string) (TimeTable, error) {
	// отдает структуру timeTable(расписание) по ID преподавателя ¯\_(ツ)_/¯
	if date == "" || dateEnd == "" {
		date = dateNow()
		dateEnd = dateNow()
	}
	var result TimeTable
	url := fmt.Sprintf("%s/api/schedule/person/%s?start=%s&finish=%s&lng=1", host, teachId, date, dateEnd)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	return result, nil
}

/*
do i even need this??
func SearchBuilding(host, buildingName string ) (string, error)  {
	url := fmt.Sprintf("%s/api/search?term=%s&type=building",host,buildingName)
	resp,err := http.Get(url)
	if err != nil{
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err!= nil{}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err!=nil{
		return "", err
	}
	var result
}
*/
