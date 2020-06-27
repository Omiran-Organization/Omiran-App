import React from 'react';

import ScrollToBottom from 'react-scroll-to-bottom';

import Message from './Message/Message';


const Messages = ({ messages,  name }) => {
  return (
  <ScrollToBottom className="messages">

  {messages.map((message, i) => {
    let len = message.data.split(" ").length
    let user = message.data.split(" ").pop()
    let data = message.data.split(" ").slice(0,len-1).join(" ")



    return (
      <div key={i}><Message message={data} user={user} name={name}/></div>
      
      )
    })
}
  </ScrollToBottom>
  )};

export default Messages;