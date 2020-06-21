// import * as React from "react";
import React, { useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { initializeStore } from '../lib/redux';
import { initializeApollo } from '../lib/apollo';
import { ProfileData } from "@/types/profile";
import gql from 'graphql-tag';

// type ProfileComponentProps = {
//   data: ProfileData;
//   isLoggedIn?: boolean;
// };

// const ProfileComponent: React.FunctionComponent<ProfileComponentProps> = ({
//   props,
//   data,
//   isLoggedIn,
// }) => {

  const ProfileComponent: React.FunctionComponent = (props) => {

    const stuff = Object.values(props);
    const other_stuff = stuff[1];
    console.log(other_stuff);
    const more_stuff = Object.values(other_stuff);
    console.log(more_stuff);

    const idx = getRandomInt(more_stuff.length);
    const my_stuff = more_stuff[idx];
    console.log(my_stuff);

    const [values, setValues] = useState({
      isLoggedIn: false;
    })
    const { username, profilePicture, following, followers } = props;
    // console.log(props.)

    return (
      <div className="flex flex-col border border-gray-500 rounded-lg w-full p-5">
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
                <b>{followers}</b> Followers
              </span>
              <span className="text-sm">
                <b>{following}</b> Following
              </span>
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
  const reduxStore = initializeStore();
  const apolloClient = initializeApollo();

  // const typeDefs = gql`
  //   interface Person {
  //     uuid
  //     username
  //     email
  //     followers
  //     following
  //   }`
  const { dispatch } = reduxStore;

  dispatch({
    type: 'TICK',
    light: true,
    lastUpdate: Date.now(),
  })

  await apolloClient.query({
      query: gql`
        query Users {
          Users {
            uuid,
            username,
            email,
            description,
            profile_picture
          }
      }
      `
  }).then(result => console.log(`result${JSON.stringify(result)}`);
  );
  

  return {
    props: {
      initialReduxState: reduxStore.getState(),
      initialApolloState: apolloClient.cache.extract(),
    },
    unstable_revalidate: 1,
  }
}
export default ProfileComponent;


function getRandomInt(max) {
  return Math.floor(Math.random() * Math.floor(max));
}