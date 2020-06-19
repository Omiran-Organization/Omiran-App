import { useDispatch } from 'react-redux';
import { initializeStore } from '../lib/redux';
import { initializeApollo } from '../lib/apollo';
import useInterval from '../lib/useInterval';
import Layout from '../components/Layout';
import Submit from '../components/Submit';
import Head from "next/head";
import gql from 'graphql-tag';

// import PostList, {
//   ALL_POSTS_QUERY,
//   allPostsQueryVars,
// } from '../components/PostList'

const HomePage: object = (props: any) => {
  // Tick the time every second
  const dispatch = useDispatch();
  
  useInterval(() => {
    dispatch({
      type: 'TICK',
      light: true,
      lastUpdate: Date.now(),
    });
  }, 1000)
  console.log(props);
  return (

      <div className="main flex flex-col justify-center items-center w-4/5 lg:w-1/2 mx-auto text-center">
        <Head>
          <title>Omiran</title>
        </Head>
      
        <h1 className="text-orange-500 text-6xl font-bold m-3">Omiran</h1>
        <h3 className="text-white text-3xl mb-5">
          An open source live streaming application that is developer-oriented and
          removes the political overhead and censorship that other streaming
          platforms exhibit.
      </h3>

        <button className="bg-black px-8 py-3 text-xl rounded-lg text-white">
          Go Live
    </button>

    <Layout>
        {/* Redux */}
        <hr />
        {/* Apollo */}
        <Submit />
        {/* <PostList /> */}
      </Layout>

      </div>
   

    
  );
}
  

  export async function getStaticProps(): object {
    const reduxStore = initializeStore();
    const apolloClient = initializeApollo();
    const { dispatch } = reduxStore;

    dispatch({
      type: 'TICK',
      light: true,
      lastUpdate: Date.now(),
    })

    await apolloClient.query({
      // query: ALL_POSTS_QUERY,
      // variables: allPostsQueryVars,

        query: gql`
          query Users { 
            Users {
                uuid,
                username,
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


export default HomePage;
