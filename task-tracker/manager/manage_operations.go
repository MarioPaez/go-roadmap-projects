package manager

import "fmt"

//Args[1] -> add, update, delete, mark-in-progress, mark-done, list
//Args[2] -> {id}, description, done, todo, in-progress

//Case Add: args[1]==add y args[2]== description entre dobles comillas

var ADD = "add"
var UPDATE = "update"
var DELETE = "delete"
var MARK_IN_PROGRESS = "mark-in-progress"
var MARK_DONE = "mark-done"
var LIST = "list"

func ManageOperations(args []string) {
	fmt.Println("Las operaciones que hemos recibido son: ", args)
	if len(args) > 1 {
		switch args[1] {
		case ADD:
			println("has puesto add")
		case UPDATE:
			println("has puesto update")
		case DELETE:
			println("has puesto delete")
		case MARK_IN_PROGRESS:
			println("has puesto MARK_IN_PROGRESS")
		case MARK_DONE:
			println("has puesto MARK_DONE")
		case LIST:
			println("has puesto LIST")
		default:
			fmt.Println("Don't support the operation ", args[1], ". Only support add, update, delete, mark-in-progress, mark-done and list")
		}
	}
}
