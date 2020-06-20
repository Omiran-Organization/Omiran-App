import { initializeApollo } from "../lib/apollo";
import Layout from "../components/Layout";
import Submit from "../components/Submit";
import Head from "next/head";
import gql from "graphql-tag";

// import PostList, {
//   ALL_POSTS_QUERY,
//   allPostsQueryVars,
// } from '../components/PostList'

const HomePage = (props) => {
  // Tick the time every second
  return (
    <Layout>
      <hr />
      {/* Apollo */}
      <Submit />
      {/* <PostList /> */}
    </Layout>
  );
};

export async function getStaticProps() {
  const apolloClient = initializeApollo();

  await apolloClient
    .query({
      // query: ALL_POSTS_QUERY,
      // variables: allPostsQueryVars,

      query: gql`
        query Users {
          Users {
            uuid
            username
          }
        }
      `,
    })
    .then((result) => console.log(`result${JSON.stringify(result)}`));

  return {
    props: {
      initialApolloState: apolloClient.cache.extract(),
    },
    unstable_revalidate: 1,
  };
}

export default HomePage;
