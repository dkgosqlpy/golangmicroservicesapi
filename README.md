# golangmicroservicesapi
golang micro services using go-gin gorm and mysql

APP name: SpringCT_Golang
-------------------------

1. Extract given springctapp.zip

2. cd springctapp

3. Creating Database
	create database demo_springct_app;

4. Need to make database and import given "demo_springct_app.sql"
	Change database credential line number 17:
	
	springctpkgs/databases/dbconnection.go
	
	_user = "root"
	_pass = "root#123PD"
	_host = "127.0.0.1"
	_port = "3306"
	_dbname = "demo_springct_app"
 
6. make sure current dir will "springctapp"
 
	go run main.go
	

7. create view
	drop  view if exists `view_students`;
	create view view_students as SELECT st.st_id, st.st_name, st.st_email, st.st_phone, GROUP_CONCAT(cs.course_name) as enrolled_courses FROM `students` as st INNER JOIN map_students_courses as msc on st.st_id=msc.st_id INNER JOIN courses as cs on cs.id = msc.cs_id group by st.st_id order by st.st_name;


8. All APIs:
------------
	GET:
		http://localhost:8088/test
		http://localhost:8088/conn
		http://localhost:8088/studentlist
		http://localhost:8088/courselist
		http://localhost:8088/viewstudentlist

	POST:
		http://localhost:8088/addstudent
		http://localhost:8088/addcourse
		
	POST:
		http://localhost:8088/addcourse
		Form Data:
		----------
			name:CIC
			desc:CIC desc
			profname:Rahul Patel
			
		Output:
		-------
		{ "data": { "ID": 9, "CourseName": "CIC", "Description": "CIC desc", "CourseProfName": "Rahul Patel",
			"CreatedDt": "2021-08-19T10:31:46.471931964+05:30" }, "message": "Register"	}


		http://localhost:8088/addstudent
		Form Data:
		----------
			name:CIC
			email:dhananjayksharma@gmail.com
			phone:9819545584
			
		Output:
		-------
		{ "data": { "StId": 11, "StName": "CIC", "StEmail": "dhananjayksharma@gmail.com", "StPhone": "9819545584", "CreatedDt": "2021-08-19T11:06:15.803217922+05:30" }, "message": "Register" }



		http://localhost:8088/enroll
		
		Form Data:
		----------
			stid:1
			csid:9
			
		Output:
		-------
		{ "data": { "MscId": 10, "StId": 1, "CsId": 9, "CreatedDt": "2021-08-19T11:50:07.519117434+05:30" }, "message": "Register" }
		
	DELETE:
		http://localhost:8088/courselist
		Form Data:
		---------- 
			csid:9
 
