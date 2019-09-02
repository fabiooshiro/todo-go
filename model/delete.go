package model

func DeleteTodo(name string) error {
	insertQ, err := con.Query("delete from TODO where name = ?", name)
	defer insertQ.Close()
	if err != nil {
		return err
	}
	return nil
}
