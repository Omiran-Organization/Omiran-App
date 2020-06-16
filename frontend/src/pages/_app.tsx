import React from "react";

import { AppProps } from "next/app";

import "../public/css/style.css";
import { ApolloProvider } from '@apollo/react-hooks'
import { useApollo } from '../lib/apolloClient'

const App = ({ Component, pageProps }: AppProps) => {
  const apolloClient = useApollo(pageProps.initialApolloState)

  return (
    <ApolloProvider client={apolloClient}>
      <Component {...pageProps} />
    </ApolloProvider>
  )
}

export default App;

