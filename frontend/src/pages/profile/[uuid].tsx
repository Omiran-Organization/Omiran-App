import React, {useState, useEffect, useRef } from 'react'

import Head from 'next/head'
import queryString from 'query-string';

import { useRouter } from 'next/router'
import { initializeApollo, useApollo } from '@/utils/apollo'
import { ProfileDataQuery } from '@/gql'

import { UserData } from '@/types'
import { NextPageContext } from 'next'
import TextContainer from '../../components/Chat/TextContainer/TextContainer';
import Messages from '../../components/Chat/Messages/Messages';
import InfoBar from '../../components/Chat/InfoBar/InfoBar';
import Input from '../../components/Chat/Input/Input';


type ProfilePageProps = {
  initialApolloState: any
}

const ProfilePage: React.FunctionComponent<ProfilePageProps> = ({ initialApolloState }) => {

  //initialApolloState is props retrieved from server side
  const apolloClient = useApollo(initialApolloState)
  const router = useRouter()
  const [isPaused, setPause] = useState(false)
  const [message, setMessage] = useState('');
  const [users, setUsers] = useState([]);
  const [messages, setMessages] = useState([]);
  const [input, setInput] = useState('');

  const ws = useRef(null)

  const data = apolloClient.readQuery({
    query: ProfileDataQuery,
    variables: {
      uuid: router.query.uuid
    }
  })

  const userData: UserData = data.User
  const followers: UserData[] = data.followers
  const following: UserData[] = data.following

  useEffect(() => {
    ws.current = new WebSocket("ws://localhost:8080/ws/" + String(99));
    ws.current.onopen = () => console.log("ws opened");
    ws.current.onclose = () => console.log("ws closed");
    
    return () => {
      ws.current.close()
    }

  }, []);

  useEffect(() => {
    if (!ws.current) return;
    ws.current.addEventListener('message', message => {

      if (isPaused) return;
      setMessages(messages => [...messages, JSON.parse(message.data)]);
      
    });
    
    }, [isPaused]);
 
  const sendMessage = (event) => {
    event.preventDefault();
    console.log(event.target)
    if (message) {
      ws.current.send(message, () => setMessage(''));
    }
    setInput("")
  }
  const sendToStream = function () {
    router.push(`/streams/${ userData.uuid }`)
  }
  return (
    <div className="main flex flex-col">
      <Head>
        <title>{ userData.username } - Omiran</title>
      </Head>
  
      <div className="flex-grow"/>
      <div className="flex-1 flex-col border border-gray-500 rounded-lg w-11/12 md:w-4/5 p-5 mx-auto">
        <div className="flex flex-row items-center w-full">
          <img
            className="rounded-full mr-6"
            src={userData.profile_picture}
            alt={userData.username}
            height={100}
            width={100}
          />
          <div className="flex flex-col">
            <h1 className="text-xl sm:text-2xl md:text-3xl text-left">{userData.username}</h1>
            <div className="flex flex-row">
              <span className="text-sm mr-3">
                <b>{followers.length}</b> Followers
              </span>
              <span className="text-sm">
                <b>{following.length}</b> Following
              </span>
            </div>
          </div>
          <div className="flex-grow" />
          <button className="btn btn-orange">{true ? 'Edit Profile' : 'Follow'}</button>
          <button className="btn btn-orange btn-outlined" onClick={sendToStream}>Streams</button>
            <div className="outerContainer flex-grow">
              <div className="container flex-grow">
                <InfoBar room={99} />
                <Messages messages={messages}  name={userData.username} />
                <Input message={message} setMessage={setMessage} sendMessage={sendMessage} input={input} setInput={setInput}/>
              </div>
              <TextContainer users={users}/>
              </div>
            </div>
          </div>
        <div className="flex-grow" />

    </div>
  )
}

export async function getServerSideProps(context: NextPageContext) {
  const { uuid } = context.query
  console.log(uuid)
  const apolloClient = initializeApollo()

  await apolloClient
    .query({
      query: ProfileDataQuery,
      variables: {
        uuid: String(uuid)
      }
    })

  return {
    props: {
      initialApolloState: apolloClient.cache.extract(),
    },
  }
}

export default ProfilePage
