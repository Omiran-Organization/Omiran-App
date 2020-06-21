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
