import React, {useState, useEffect, useRef } from 'react'

import Head from 'next/head'

import { useRouter } from 'next/router'
import { initializeApollo, useApollo } from '@/utils/apollo'
import { ProfileDataQuery } from '@/gql'

import { UserData } from '@/types'
import { NextPageContext } from 'next'

type ProfilePageProps = {
  initialApolloState: any
}

const ProfilePage: React.FunctionComponent<ProfilePageProps> = ({ initialApolloState }) => {
  const apolloClient = useApollo(initialApolloState)
  const router = useRouter()
  // const [isPaused, setPause] = useState(false)

//   useEffect(() => {
//     ws.current = new WebSocket("ws://localhost:8080/ws");
//     ws.current.onopen = () => console.log("ws opened");
//     ws.current.onclose = () => console.log("ws closed");

//     return () => {
//       ws.current.close()
//     }
//   }, []);
//   useEffect(() => {
//     if (!ws.current) return;

//     ws.current.onmessage = e => {
//         if (isPaused) return;
//         const message = JSON.parse(e.data);
//         console.log("e", message);
//     };
// }, [isPaused]);
  const data = apolloClient.readQuery({
    query: ProfileDataQuery,
    variables: {
      uuid: router.query.uuid
    }
  })
  console.log(data)

  const userData: UserData = data.Users[0]
  const followers: UserData[] = data.followers
  const following: UserData[] = data.following
  const ws = useRef(null)
  
  return (
    <div className="main flex flex-col">
      <Head>
        <title>{ userData.username } - Omiran</title>
      </Head>
      <div className="flex-grow"/>
      <div className="flex flex-col border border-gray-500 rounded-lg w-11/12 md:w-4/5 p-5 mx-auto">
        <div className="flex flex-row items-center w-full">
          <img
            className="rounded-full mr-6"
            src={userData.profilePicture}
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
        </div>
      </div>
      <div className="flex-grow-3"/>
    </div>
  )
}

export async function getServerSideProps(context: NextPageContext) {
  const { uuid } = context.query

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
