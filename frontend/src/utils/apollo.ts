import { useMemo } from "react";
import { ApolloClient } from "apollo-client";
import { InMemoryCache, NormalizedCacheObject } from "apollo-cache-inmemory";
import { HttpLink } from "apollo-link-http";
import { resolvers } from "./resolvers";
import typeDefs from "./typeDefs"

let apolloClient;

function createApolloClient(): ApolloClient<NormalizedCacheObject> {
  return new ApolloClient({
    ssrMode: typeof window === "undefined",
    link: new HttpLink({
      uri: "http://localhost:8080/graphql", // Server URL (must be absolute)
      credentials: "same-origin", // Additional fetch() options like `credentials` or `headers`
      // uri: process.env.API_URL + '/graphql', // Server URL (must be absolute)
      // credentials: 'same-origin', // Additional fetch() options like `credentials` or `headers`
    }),
    cache: new InMemoryCache(),
    resolvers, // resolvers which contain functions that run according to the called mutation.
    typeDefs,// Compared to Redux a mutation is an action, a resolver is a reducer and ApolloClient is the store.
  });
}

export function initializeApollo(initialState: any = null): ApolloClient<NormalizedCacheObject> {
  const _apolloClient = apolloClient ?? createApolloClient()

  // If your page has Next.js data fetching methods that use Apollo Client, the initial state
  // get hydrated here
  if (initialState) {
    _apolloClient.cache.restore(initialState);
  }
  // For SSG and SSR always create a new Apollo Client
  if (typeof window === "undefined") return _apolloClient;
  // Create the Apollo Client once in the client
  if (!apolloClient) apolloClient = _apolloClient;

  return _apolloClient;
}

export function useApollo(initialState: any): ApolloClient<NormalizedCacheObject> {
  const client = useMemo(() => initializeApollo(initialState), [initialState])
  return client
}
