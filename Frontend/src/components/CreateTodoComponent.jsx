import React, { Component } from 'react'
import TodoService from '../services/TodoService';

class CreateTodoComponent extends Component {
    constructor(props) {
        super(props)

        this.state = {
            // step 2
            id: this.props.match.params.id,
            title: '',
            description: '',
        }
        this.changeTitleHandler =
            this.changeTitleHandler.bind(this);
        this.changeDescriptionHandler =
            this.changeDescriptionHandler.bind(this);
        this.saveOrUpdateTodo =
            this.saveOrUpdateTodo.bind(this);
    }

    // step 3
    componentDidMount() {

        // step 4
        if (this.state.id === '_add') {
            return
        } else {
            TodoService.getTodoById(this.state.id).
            then((res) => {
                let todo = res.data;
                this.setState({
                    title: todo.title,
                    description: todo.description,
                });
            });
        }
    }
    saveOrUpdateTodo = (e) => {
        e.preventDefault();
        let todo = { title: this.state.title, description:
             this.state.description, };
        console.log('todo => ' + JSON.stringify(todo));

        // step 5
        if (this.state.id === '_add') {
            TodoService.createTodo(todo).then(res => {
                this.props.history.push('/todo');
            });
        } else {
            TodoService.
            updateTodo(todo, this.state.id).then(res => {
                this.props.history.push('/todo');
            });
        }
    }

    changeTitleHandler = (event) => {
        this.setState({ title: event.target.value });
    }

    changeDescriptionHandler = (event) => {
        this.setState({ description: event.target.value });
    }

    cancel() {
        this.props.history.push('/todo');
    }

    getTitle() {
        if (this.state.id === '_add') {
            return <h3 className="text-center">Add Todo</h3>
        } else {
            return <h3 className="text-center">Update Todo</h3>
        }
    }
    render() {
        return (
            <div>
                <br></br>
        <div className="container">
            <div className="row">
               <div className="card col-md-6 offset-md-3 offset-md-3">
                            {
                                this.getTitle()
                            }
                            <div className="card-body">
                                <form>
            <div className="form-group">
              <label> Title: </label>
                <input placeholder="Title" 
                  name="title" className="form-control"
                    value={this.state.title} 
                      onChange={this.changeTitleHandler} />
                          </div>
            <div className="form-group">
              <label> Description: </label>
                <input placeholder="Description" 
                   name="description" className="form-control"
                     value={this.state.description} 
                      onChange={this.changeDescriptionHandler} />
                                    </div>

             <button className="btn btn-success" 
                  onClick={this.saveOrUpdateTodo}>Save
                    </button>
             <button className="btn btn-danger" 
                  onClick={this.cancel.bind(this)} 
                     style={{ marginLeft: "10px" }}>Cancel
                        </button>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default CreateTodoComponent
