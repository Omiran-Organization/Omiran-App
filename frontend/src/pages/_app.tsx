import React from 'react'

import { ApolloProvider } from '@apollo/react-hooks'

import { useApollo } from '../lib/apollo'

import '../styles/style.css'

export default function App({ Component, pageProps }) {
  const apolloClient = useApollo(pageProps.initialApolloState)

  return (
    <ApolloProvider client={apolloClient}>
      <Component {...pageProps} />
    </ApolloProvider>
  )
}

// const App = ({ Component, pageProps }: AppProps) => {
//   const apolloClient = useApollo(pageProps.initialApolloState)

//   return (
//     <ApolloProvider client={apolloClient}>
//       <Component {...pageProps} />
//     </ApolloProvider>
//   )
// }
