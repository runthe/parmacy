package main

import (
	"encoding/xml"
)

type Response struct {
	XMLName xml.Name `xml:"response"`
	Header  Header `xml:"header"`
	Body    Body `xml:"body"`
}

type Header struct {
	ResultCode string `xml:"resultCode"`
	ResultMsg  string `xml:"resultMsg"`
}

type Body struct {
	Items      []Item `xml:"items>item"`
	NumOfRows  int `xml:"numOfRows"`
	PageNo     int `xml:"pageNo"`
	TotalCount int `xml:"totalCount"`
}

type Item struct {
	DutyName   string `xml:"dutyName"`   //기관명
	PostCdn1   string `xml:"postCdn1"`   //우편번호1
	PostCdn2   string `xml:"postCdn2"`   //우편번호2
	DutyAddr   string `xml:"dutyAddr"`   //주소
	DutyTel1   string `xml:"dutyTel1"`   //대표전화1
	Wgs84Lan   string `xml:"wgs84Lon"`   //병원경도
	Wgs84Lat   string `xml:"wgs84Lat"`   //병원위도
	DutyEtc    string `xml:"dutyEtc"`    //비고
	DutyInf    string `xml:"dutyInf"`    //기관설명
	DutyMapimg string `xml:"dutyMapimg"` //간이약도

	DutyTime1c string `xml:"dutyTime1c"` //진료시간(월요일)C
	DutyTime2c string `xml:"dutyTime2c"` //진료시간(화요일)C
	DutyTime3c string `xml:"dutyTime3c"` //진료시간(수요일)C
	DutyTime4c string `xml:"dutyTime4c"` //진료시간(목요일)C
	DutyTime5c string `xml:"dutyTime5c"` //진료시간(금요일)C
	DutyTime6c string `xml:"dutyTime6c"` //진료시간(토요일)C
	DutyTime7c string `xml:"dutyTime7c"` //진료시간(일요일)C
	DutyTime8c string `xml:"dutyTime8c"` //진료시간(공휴일)C
	DutyTime1s string `xml:"dutyTime1s"` //진료시간(월요일)S
	DutyTime2s string `xml:"dutyTime2s"` //진료시간(월\화요일)S
	DutyTime3s string `xml:"dutyTime3s"` //진료시간(수요일)S
	DutyTime4s string `xml:"dutyTime4s"` //진료시간(목요일)S
	DutyTime5s string `xml:"dutyTime5s"` //진료시간(금요일)S
	DutyTime6s string `xml:"dutyTime6s"` //진료시간(토요일)S
	DutyTime7s string `xml:"dutyTime7s"` //진료시간(일요일)S
}

