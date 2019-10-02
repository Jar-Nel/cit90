// Bowler Table
//  FirstName varchar(100)
//  LastName varchar(100)
//  TeamID int
//  Zipcode varchar(10)

package main

import (
	"fmt"
)

type sqldbt struct{
	bowlers []bowler
	teams []team
}

type sqlRow struct{
	rowdata interface{}
}

type mongoRow struct{
	rowdata interface{}
}


func (r sqlRow) save(db *sqldbt){
//	fmt.Printf("SQL: %v %T\n", r.rowdata, r.rowdata)
	switch r.rowdata.(type){
	case bowler:
		db.bowlers=append(db.bowlers,r.rowdata.(bowler))	
	case team:
		db.teams=append(db.teams,r.rowdata.(team))
	}
}

func (d sqldbt) retrieve(t string,row int) interface{} {
	var ret interface{}
	switch(t){
	case "bowler":
		ret= d.bowlers[row]
	case "team":
		ret= d.teams[row]
	}
	return ret
}

func (r mongoRow) save(){
	//fmt.Println(r.rowdata)
	fmt.Printf("MONGO: %v %T\n", r.rowdata, r.rowdata)
}

type row interface{}
type database interface {
	save()
}

type bowler struct {
	FirstName string
	LastName string
	TeamID int
	Zipcode string
}

type team struct {
	TeamName string
}

func main(){
	var sqldb =sqldbt{}
	b:= bowler{
		FirstName: "Jared",
		LastName:"Nelson",
		TeamID: 3,
		Zipcode: "93726",
	}
	
	var r row = sqlRow {rowdata: b}
	r.(sqlRow).save(&sqldb)

	r=mongoRow {rowdata: b}
	r.(mongoRow).save()

	r=42
	fmt.Println(r.(int))

	b2:=bowler{
		FirstName:"John",
		LastName:"Doe",
		TeamID: 1,
		Zipcode: "90210",
	}
	row:=sqlRow  {rowdata: b2}
	row.save(&sqldb)

	row=sqlRow {rowdata: team {TeamName:"Bowlerz"}}
	row.save(&sqldb)
	
	fmt.Println("Retrieve: ", sqldb.retrieve("team", 0).(team).TeamName)
	fmt.Println("Retrieve: ", sqldb.retrieve("bowler", 0).(bowler).FirstName)
	fmt.Println("Retrieve: ", sqldb.retrieve("bowler", 1))
	//writeData(bowlerRow.rowdata)
}