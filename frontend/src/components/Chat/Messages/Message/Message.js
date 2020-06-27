import React from 'react';

import ReactEmoji from 'react-emoji';



const Message = ({message, name}) => {

  // let newMessage = message.data[message.length-1]
  // let newUser = message.name[user.length-1]
  // let user = message.name


  let isSentByCurrentUser = false;
  if(message.name === name) {
    isSentByCurrentUser = true;
  }
  return (
    isSentByCurrentUser
      ? (
        <div className="messageContainer justifyEnd">
          <p className="sentText pr-10">{name}</p>
          <div className="messageBox backgroundBlue">
            <p className="messageText colorWhite">{message.data}</p>
          </div>
        </div>
        )
        : (
          <div className="messageContainer justifyStart">
             <div className="messageBox backgroundLight">
              <p className="messageText colorDark">{message.data}</p> 
            </div>
            <p className="sentText pl-10 ">{message.name}</p>
          </div>
        )
  );
}

export default Message;