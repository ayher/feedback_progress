import React, { Component } from 'react';
import { Link,Route } from 'react-router-dom'
import Home from './components/Home'
import About from './components/About'
import 'bootstrap/dist/css/bootstrap.min.css'
import 'bootstrap/dist/js/bootstrap'
export default class App extends Component {
    render() {
        return (
            <div>
                <div className="row">
                    <div className="col-xs-offset-2 col-xs-8">
                        <div className="page-header"><h2>React Router Demo</h2></div>
                    </div>
                </div>
                <div className="row">
                    <div className="col-xs-2 col-xs-offset-2">
                        <div className="list-group">
                            {/* 原生html 要靠<a>跳转不同的页面 */}
                            {/* <a className="list-group-item" href="./about.html">About</a>
                            <a className="list-group-item active" href="./home.html">Home</a> */}

                            {/* 在React中靠路由链接实现切换组件 */}
                            <Link className="list-group-item" to="/home">Home</Link>
                            <Link className="list-group-item" to="/about">About</Link>
                        </div>
                    </div>
                    <div className="col-xs-6">
                        <div className="panel">
                            <div className="panel-body">
                                {/* 注册路由 */}
                                <Route path="/home" component={Home}/>
                                <Route path="/about" component={About}/>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}