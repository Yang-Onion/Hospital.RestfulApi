package main

import (
	c "controller"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	//体检报告 每次20条
	router.HandleFunc("/api/physical/list", c.GetReportList).Methods("GET")
	//根据体检任务Id获取体检报告
	router.HandleFunc("/api/physical/list/task/{id}", c.GetReportByTaskId).Methods("GET")
	//根据病人Id获取体检报告
	router.HandleFunc("/api/physical/list/patients/{bingRenId}", c.GetReportByPatientId).Methods("GET")
	//体检报告 状态回写
	router.HandleFunc("/api/physical/callback", c.CallBack).Methods("POST")
	//获取基础信息
	router.HandleFunc("/api/physical/base-data", c.GetBasicDataList).Methods("GET")
	//体检单位
	router.HandleFunc("/api/physical/team/list", c.GetTeamList).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
