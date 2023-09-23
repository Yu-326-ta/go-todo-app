package models

import (
	"reflect"
	"testing"
	"time"
)

// func TestUser_CreateTodo(t *testing.T) {
// 	type fields struct {
// 		ID        int
// 		UUID      string
// 		Name      string
// 		Email     string
// 		Password  string
// 		CreatedAt time.Time
// 		Todos     []Todo
// 	}
// 	type args struct {
// 		content string
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		args    args
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			u := &User{
// 				ID:        tt.fields.ID,
// 				UUID:      tt.fields.UUID,
// 				Name:      tt.fields.Name,
// 				Email:     tt.fields.Email,
// 				Password:  tt.fields.Password,
// 				CreatedAt: tt.fields.CreatedAt,
// 				Todos:     tt.fields.Todos,
// 			}
// 			if err := u.CreateTodo(tt.args.content); (err != nil) != tt.wantErr {
// 				t.Errorf("User.CreateTodo() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

func TestGetTodo(t *testing.T) {
	type args struct {
		id int
	}
	type Todo struct {
		ID        int
		Content   string
		UserID    int
		CreatedAt time.Time
	}
	tests := []struct {
		name     string
		args     args
		wantTodo Todo
		wantErr  bool
	}{
		{
			name: "宿題",
			args: args{6},
			wantTodo: Todo{
				ID: 6,
				Content: "宿題",
				UserID: 1,
			},
			wantErr: true,
		},
		{
			name: "宿題",
			args: args{6},
			wantTodo: Todo{
				ID: 6,
				Content: "宿題",
				UserID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTodo, err := GetTodo(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTodo, tt.wantTodo) {
				t.Errorf("GetTodo() = %v, want %v", gotTodo, tt.wantTodo)
			}
		})
	}
}

// func TestGetTodos(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		wantTodos []Todo
// 		wantErr   bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			gotTodos, err := GetTodos()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("GetTodos() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotTodos, tt.wantTodos) {
// 				t.Errorf("GetTodos() = %v, want %v", gotTodos, tt.wantTodos)
// 			}
// 		})
// 	}
// }

// func TestUser_GetTodosByUser(t *testing.T) {
// 	type fields struct {
// 		ID        int
// 		UUID      string
// 		Name      string
// 		Email     string
// 		Password  string
// 		CreatedAt time.Time
// 		Todos     []Todo
// 	}
// 	tests := []struct {
// 		name      string
// 		fields    fields
// 		wantTodos []Todo
// 		wantErr   bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			u := &User{
// 				ID:        tt.fields.ID,
// 				UUID:      tt.fields.UUID,
// 				Name:      tt.fields.Name,
// 				Email:     tt.fields.Email,
// 				Password:  tt.fields.Password,
// 				CreatedAt: tt.fields.CreatedAt,
// 				Todos:     tt.fields.Todos,
// 			}
// 			gotTodos, err := u.GetTodosByUser()
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("User.GetTodosByUser() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(gotTodos, tt.wantTodos) {
// 				t.Errorf("User.GetTodosByUser() = %v, want %v", gotTodos, tt.wantTodos)
// 			}
// 		})
// 	}
// }

// func TestTodo_UpdateTodo(t *testing.T) {
// 	type fields struct {
// 		ID        int
// 		Content   string
// 		UserID    int
// 		CreatedAt time.Time
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tr := &Todo{
// 				ID:        tt.fields.ID,
// 				Content:   tt.fields.Content,
// 				UserID:    tt.fields.UserID,
// 				CreatedAt: tt.fields.CreatedAt,
// 			}
// 			if err := tr.UpdateTodo(); (err != nil) != tt.wantErr {
// 				t.Errorf("Todo.UpdateTodo() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }

// func TestTodo_DeleteTodo(t *testing.T) {
// 	type fields struct {
// 		ID        int
// 		Content   string
// 		UserID    int
// 		CreatedAt time.Time
// 	}
// 	tests := []struct {
// 		name    string
// 		fields  fields
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tr := &Todo{
// 				ID:        tt.fields.ID,
// 				Content:   tt.fields.Content,
// 				UserID:    tt.fields.UserID,
// 				CreatedAt: tt.fields.CreatedAt,
// 			}
// 			if err := tr.DeleteTodo(); (err != nil) != tt.wantErr {
// 				t.Errorf("Todo.DeleteTodo() error = %v, wantErr %v", err, tt.wantErr)
// 			}
// 		})
// 	}
// }
