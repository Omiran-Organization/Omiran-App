import React, { useState, useEffect} from 'react';


const Input = ({ setMessage, sendMessage, message, input, setInput }) => {
  const [value, setValue] = useState('')
  useEffect(() => {
    setInput(message)
  },[message])
  

  return (
  <form className="form_chat">
    <input
      className="input"
      type="text"
      placeholder="Type a message..."
      value={input}
      onChange={({ target: { value } }) => setMessage(value)}
      onKeyPress={event => event.key === 'Enter' ? sendMessage(event) : null}
      
    />
    <button className="sendButton" onClick={e => sendMessage(e)}>Send</button>
  </form>
)};

export default Input;