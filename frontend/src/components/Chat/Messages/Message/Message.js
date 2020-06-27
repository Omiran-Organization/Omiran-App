import React from 'react';

import ReactEmoji from 'react-emoji';



const Message = ({message, user, name}) => {

  let isSentByCurrentUser = false;
  if(user=== name) {
    isSentByCurrentUser = true;
  }
  return (
    isSentByCurrentUser
      ? (
        <div className="messageContainer justifyEnd">
          <p className="sentText pr-10">{name}</p>
          <div className="messageBox backgroundBlue">
            <p className="messageText colorWhite">{message}</p>
          </div>
        </div>
        )
        : (
          <div className="messageContainer justifyStart">
             <div className="messageBox backgroundLight">
              <p className="messageText colorDark">{message}</p> 
            </div>
            <p className="sentText pl-10 ">{user}</p>
          </div>
        )
  );
}

export default Message;