package dao

import(
	"TODOLIST/model"
	"fmt"
	"context"
)

func InsertTask(task *model.Task) error{
	Db:= OpenDbConnection()

	stmt:=fmt.Sprintf("insert into task(name,start,due,note,id_user) values('%s','%s','%s','%s',%d)",task.Name,task.Start,task.Due,task.Note,task.Id_user)
	_,err:=Db.Exec(stmt)

	Db.Close()
	return err
}

func UpdateTask(task *model.Task) error{
	Db:= OpenDbConnection()

	stmt:=fmt.Sprintf("update task set name='%s',start='%s',due='%s',note='%s',id_user=%d where id=%d ",task.Name,task.Start,task.Due,task.Note,task.Id_user,task.Id)
	_,err:=Db.Exec(stmt)

	Db.Close()
	return err
}

func DeleteTask(id int) error{
	Db:=OpenDbConnection()

	stmt:=fmt.Sprintf("delete from task where id=%d",id)
	_,err:=Db.Exec(stmt)

	Db.Close()
	return err
}

func GetAllTasks(id int)[] model.Task{
	Db:=OpenDbConnection()

	var tasks[] model.Task

	stmt:=fmt.Sprintf("select * from task where id_user=%d",id )
	rows,_:=Db.Query(stmt)
	for rows.Next(){
		var task model.Task
		_=rows.Scan(&task.Id,&task.Name,&task.Start,&task.Due,&task.Note,&task.Id_user)
		tasks=append(tasks,task)
	}

	Db.Close()
	return tasks
}

func GetTask(id int) model.Task{
	Db:=OpenDbConnection()

	var task model.Task
	stmt:=fmt.Sprintf("select * from task where id=%d",id )
	ctx := context.Background()
	
	_ = Db.QueryRowContext(ctx, stmt).Scan(&task.Id, &task.Name,&task.Start,&task.Due,&task.Note,&task.Id_user)
	defer Db.Close()
	return task
}