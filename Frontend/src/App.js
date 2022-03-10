import React from 'react';
import './App.css';
import {BrowserRouter as Router, Route, Switch} 
      from 'react-router-dom';
import ListTodoComponent from './components/ListTodoComponent';
import CreateTodoComponent from './components/CreateTodoComponent';
import ViewTodoComponent from './components/ViewTodoComponent';

function App() {
  return (
        <Router>
                    <Switch> 
                    
                    <Route path = "/" exact component =
                              {ListTodoComponent}/>
                    <Route path = "/todo" component = 
                              {ListTodoComponent}/>
                          <Route path = "/add-todo/:id" component = 
                              {CreateTodoComponent}/>
                          <Route path = "/view-todo/:id" component = 
                              {ViewTodoComponent}/>
                               
                         </Switch>
          
        </Router>
    
  );
}

export default App;