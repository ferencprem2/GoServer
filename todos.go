package main

var DB []TodoItem = make([]TodoItem, 0)
var PrimaryKey int = 1

func GetTodoItems() ([]TodoItem, error) {
	result := make([]TodoItem, 0)

	for _, item := range DB {
		result = append(result, item)
	}

	return result, nil
}

func GetTodoItem(id int) (*TodoEditor, error) {
	result := &TodoEditor{}

	for _, item := range DB {
		if item.TodoId == id {
			result.Done = item.Done
			result.Text = item.Text
			break
		}
	}

	return result, nil
}

func CreateTodoItem(editor TodoEditor) error {
	DB = append(DB, TodoItem{
		TodoId: PrimaryKey,
		Text:   editor.Text,
		Done:   editor.Done,
	})
	PrimaryKey++

	return nil

}

func UpdateTodoItem(editor TodoEditor, id int) error {
	//UPDATE TODOS SET 'Text'=:text, ""
	for i := range DB {
		if DB[i].TodoId == id {
			DB[i].Text = editor.Text
			DB[i].Done = editor.Done
			break
		}
	}

	return nil

}

func DeleteTodoItem(id int) error {
	//DELETE FROM TODOS WHERE 'TodoId'=:id
	for i := range DB {
		if DB[i].TodoId == id {
			DB = append(DB[:i], DB[i+1:]...)
			break
		}
	}
	return nil
}
