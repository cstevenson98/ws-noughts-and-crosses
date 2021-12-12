import React from 'react';
import {
    BrowserRouter as Router,
    Routes,
    Route,
    Link
} from "react-router-dom";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faHome } from '@fortawesome/free-solid-svg-icons'

import "./style.css"
import OandXGame from "./OandX";

function NavBar() {
    return (
        <nav className="fixed w-full flex items-center justify-between flex-wrap bg-gray-600 p-1">
            <div className="flex flex-grow">
                <Link to="/" className="mx-1 hover:bg-gray-400 rounded-lg hover:rounded-xxl transition-all">
                    <FontAwesomeIcon className='mx-1' key={'faHome'} icon={faHome} color='white'/>
                </Link>
                <Link to="/noughtsandcrosses" className="hover:bg-gray-400 rounded-lg hover:rounded-xxl transition-all">
                    <img className='m-1' src="OandX.svg" alt={'Noughts and Crosses Icon'}/>
                </Link>
            </div>
        </nav>
    );
}

function Home() {
    return (
        <div className='flex justify-center text-gray-400'>
            <div>Choose a game above!</div>
        </div>
    );
}

export default function Main() {
    return (
        <Router>
            <div>
                {/*Navigation*/}
                <NavBar />

                {/*Noughts and Crosses*/}
                <div className='h-screen flex justify-center items-center bg-gray-800'>
                    <Routes>
                        <Route path="/" element={ <Home /> }/>
                        <Route path="/noughtsandcrosses" element={ <OandXGame /> }/>
                    </Routes>
                </div>
            </div>
        </Router>
    );
}
