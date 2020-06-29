import React from "react";
import { useApolloClient, useMutation } from "@apollo/react-hooks";
import { INCREMENT } from "../utils/mutations";
import Head from 'next/head'
// This page has a button that increments the count variable inside apollo local state

// When you click the button a Mutation is sent to the Resolver which is loaded by Apollo.ts
// in apollo client's initialization. All of these files are in /lib

// You can see the cache in the console and the count var is under data.

// useMutation hook returns 2 variables, a function to do the mutation and an information variable about the mutation.

// Compared to Redux a mutation is an action, a resolver is a reducer and ApolloClient is the store.

const apolloLinkTest :React.FunctionComponent = () => {
  const client = useApolloClient();
  const [increment, { data }] = useMutation(INCREMENT);

  return (
    <div className="main flex flex-col justify-center items-center w-4/5 lg:w-1/2 mx-auto text-center">
      <Head>
        <title>Omiran</title>
      </Head>

      <h1 className="text-orange-500 text-6xl font-bold m-3">Omiran</h1>
      <h3 className="text-white text-3xl mb-5">

      </h3>
      <button
        onClick={() => {
          increment();
          console.log(client.cache);
        }}
      >
           Click to add a number
      </button>
    </div>
   
 

  );
}
export default apolloLinkTest