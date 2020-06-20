// import * as React from "react";
import React, { useState, useEffect } from 'react';
import { initializeApollo } from '../lib/apollo'
import { ProfileData } from "@/types/profile";
import gql from 'graphql-tag';
import Head from '../components/head'


type ProfileProps = {
  data: ProfileData;
  isLoggedIn?: boolean;

};
const ProfileComponent: React.FunctionComponent<ProfileProps> = (props) => {
  console.log(props)
  const initialState = props["initialApolloState"]
  const user_idx = Object.keys(initialState)[0]
  const followers_idx = Object.keys(initialState)[2]
  const followee_idx = Object.keys(initialState)[3]
  const _user = initialState[user_idx]
  const _following = initialState[followee_idx]
  const _followers = initialState[followers_idx]

  let merged = {..._user, ..._following, ..._followers}

  const [values, setValues] = useState({
    isLoggedIn: false
  })
  let followers = Object.values(_followers)

  let following = Object.values(_following)

  const { username, profilePicture } = merged;

  const listItems = followers.map((ele,idx) =>
    <li key={idx}> {ele}</li>
  );

  return (
    <div className="main flex flex-col justify-center items-center w-4/5 lg:w-1/2 mx-auto text-center">
    <Head >
      <title>Omiran</title>
    </Head>
    <div className="flex margin-top:100px flex-col border border-gray-500 rounded-lg w-full p-5">
      <div className="flex flex-row items-center w-full"> 
        <img
          className="rounded-full mr-6"
          src={profilePicture}
          alt={username}
          height={100}
          width={100}
        />
        <div className="flex flex-col">
          <h1 className="text-xl sm:text-2xl md:text-3xl text-left">
            {username}
          </h1>
          <div className="flex flex-row">
            <span className="text-sm mr-3">
              <ul>{listItems}</ul> 
            </span>
            <span className="text-sm">
              <b>{following}</b> Following
            </span>
          </div>
        </div>
      </div>
      <div className="flex-grow" />
      <button className="btn btn-orange">
        {values.isLoggedIn ? "Edit Profile" : "Follow"}
      </button>
    </div>
    </div>
  );
};
export async function getStaticProps() {
  
    const apolloClient = initializeApollo()


    await apolloClient.query({

	  query: gql`
	  query User { 
      User (uuid: "02e92cbf-6736-46d9-bc7c-549209107a48"){
        
        username,
        email,
        description,
        profile_picture
    }
    user_i_follow: Follows (follower:"f5a13066-31d3-4aef-81b7-a5613b774734"){
          username,
          email 
      }
      users_following_me: Follows (followee: "02e92cbf-6736-46d9-bc7c-549209107a48"){
        username
        email
      }
    }
	`
    }).then(result => console.log(result)
    );
    

    return {
      props: {
        initialApolloState: apolloClient.cache.extract(),
      },
      unstable_revalidate: 1,
    }
  }


export default ProfileComponent;

function getRandomInt(max) {
  return Math.floor(Math.random() * Math.floor(max));
}
