import React from "react";
import { useApolloClient, useMutation } from "@apollo/react-hooks";
import { INCREMENT } from "../lib/mutations";

// This page has a button that increments the count variable inside apollo local state

// When you click the button a Mutation is sent to the Resolver which is loaded by Apollo.ts
// in apollo client's initialization. All of these files are in /lib

// You can see the cache in the console and the count var is under data.

// useMutation hook returns 2 variables, a function to do the mutation and an information variable about the mutation.

// Compared to Redux a mutation is an action, a resolver is a reducer and ApolloClient is the store.

export default function apolloLinkTest() {
  const client = useApolloClient();
  const [increment, { data }] = useMutation(INCREMENT);
  console.log(data);
  return (
    <div>
      <h1>This is an apollo link test</h1>
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
