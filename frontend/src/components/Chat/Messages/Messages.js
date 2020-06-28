import React from 'react';

import ScrollToBottom from 'react-scroll-to-bottom';

import Message from './Message/Message';


const Messages = ({ messages,  name }) => {
  return (
  <ScrollToBottom className="messages">

  {messages.map((message, i) => {
    return (
      <div key={i}><Message message={message.message} user={message.sender} name={name}/></div>
      
      )
    })
}
  </ScrollToBottom>
  )};

export default Messages;