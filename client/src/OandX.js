import React, {useCallback, useEffect, useState} from "react";
import useWebSocket, {ReadyState} from "react-use-websocket";

function Board({ boardStateString, messenger }) {
    var cellsTot = []
    for (let i = 0; i < 3; i++) {
        var cellsRow = []
        for (let j = 0; j < 3; j++) {
            cellsRow.push(
                <button
                    key={3*i + j}
                    className='flex font-mono m-0.5 bg-gray-400 h-10 w-10 rounded-md justify-center items-center'
                    onClick={() => {
                        messenger(i, j);
                    }}>
                    { boardStateString[3*i + j] }
                </button>
            );
        }
        cellsTot.push(<div className='flex row'>{cellsRow}</div>);
    }

    return (
        <div className='p-0.5'>
            <div className='flex-col'>
                {cellsTot}
            </div>
        </div>
    );
}

const StatusIndicator = ({ state }) => {
    let styleCol = ""
    switch (state) {
        case ReadyState.CONNECTING:
            styleCol = 'bg-green-300'
            break;
        case ReadyState.OPEN:
            styleCol = 'bg-green-500'
            break;
        case ReadyState.CLOSING:
            styleCol = 'bg-red-100'
            break;
        case ReadyState.CLOSED:
            styleCol = 'bg-red-500'
            break;
        case ReadyState.UNINSTANTIATED:
            styleCol = 'bg-black'
            break;
        default:
            styleCol = 'bg-white'
    }

    return (
        <div className={'rounded-full w-2 h-2 mr-1' + styleCol}> d </div>
    );
}

export default function OandXGame() {
    const [socketUrl, setSocketUrl] = useState("ws://" + location.hostname + ":8765/noughtsandcrosses/connect");
    const [messageHistory, setMessageHistory] = useState([]);

    const {
        sendMessage,
        lastMessage,
        readyState,
    } = useWebSocket(socketUrl);

    useEffect(() => {
        if (lastMessage !== null) {
            setMessageHistory(prev => prev.concat(lastMessage));
        }
    }, [lastMessage, setMessageHistory]);

    const handleClickSendMessage = useCallback((i, j) =>
        sendMessage('[' + i.toString() + ', ' + j.toString() + ']'), []);

    return (
        <div className='inline-block py-1 px-1 bg-gray-500 rounded-sm border border-yellow-300'>
            <Board boardStateString={lastMessage ? lastMessage.data : ""} messenger={handleClickSendMessage}/>
        </div>
    );
}
