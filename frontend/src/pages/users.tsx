// import * as React from "react";
import React, { useState, useEffect } from "react";
import { initializeApollo } from "../lib/apollo";
import { getUsers } from "../api/users";
import Head from "next/head";
import { GetStaticProps } from 'next'
import { useRouter } from "next/router";


import gql from "graphql-tag";



export default function UserPage(props) {
  
  console.log(props)

  return (
    <div className="main flex flex-col justify-center items-center w-4/5 lg:w-1/2 mx-auto text-center">
    <Head>
      <title>Omiran</title>
    </Head>
  
    <h1 className="text-orange-500 text-6xl font-bold m-3">Omiran</h1>
    <h3 className="text-white text-3xl mb-5">
      This is a standard testpage
  </h3>

    <button className="bg-black px-8 py-3 text-xl rounded-lg text-white">
      Go Fuck Yourself
</button>
  </div>
  );
}


export const getStaticProps: GetStaticProps = async (context) => {  // Call an external API endpoint to get posts.
  // You can use any data fetching library
  const response =  await fetch("http://full_app:8080/users", {
      method: 'GET',
      credentials: 'include', 
  })
  const users = await response.json()
  return {
    props: {
      users,
    },
  }
}


  // By returning { props: posts }, the Blog component
  // will receive `posts` as a prop at build time
