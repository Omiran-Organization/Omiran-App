import { ApolloProvider } from "@apollo/react-hooks";
import { useApollo } from "../lib/apollo";
import "../public/css/style.css";

export default function App({ Component, pageProps }) {
  const apolloClient = useApollo(pageProps.initialApolloState);

  return (
    <ApolloProvider client={apolloClient}>
      <Component {...pageProps} />
    </ApolloProvider>
  );
}
// import { ClientContext } from 'graphql-hooks'
// import { useGraphQLClient } from '../lib/graphql-client'

// export default function App({ Component, pageProps }) {
//   const graphQLClient = useGraphQLClient(pageProps.initialGraphQLState)

//   return (
//     <ClientContext.Provider value={graphQLClient}>
//       <Component {...pageProps} />
//     </ClientContext.Provider>
//   )
// }

// const App = ({ Component, pageProps }: AppProps) => {
//   const apolloClient = useApollo(pageProps.initialApolloState)

//   return (
//     <ApolloProvider client={apolloClient}>
//       <Component {...pageProps} />
//     </ApolloProvider>
//   )
// }
