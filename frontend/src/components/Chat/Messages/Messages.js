import React from 'react';

import ScrollToBottom from 'react-scroll-to-bottom';

import Message from './Message/Message';


const Messages = ({ messages, users, name }) => {
  
  let user = messages.map(message=> message.name)
  let data = messages.map(message=> message.data)


  return (
  <ScrollToBottom className="messages">

  {messages.map((data, i) => <div key={i}><Message message={data} user={user} name={name}/></div>)}
  </ScrollToBottom>
  )};

export default Messages;