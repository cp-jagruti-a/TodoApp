import axios from 'axios';

const TODO_API_BASE_URL = "http://localhost:8080/todo";

class TodoService {

    getTodo(){
        return axios.get(TODO_API_BASE_URL);
    }

    createTodo(todo){
        return axios.post(TODO_API_BASE_URL, todo);
    }

    getTodoById(todoId){
        return axios.get(TODO_API_BASE_URL + '/' + todoId);
    }

    updateTodo(todo, todoId){
        return axios.put(TODO_API_BASE_URL + '/' + todoId, todo);
    }

    deleteTodo(todoId){
        return axios.delete(TODO_API_BASE_URL + '/' + todoId);
    }
}

export default new TodoService()