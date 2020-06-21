import { useDispatch } from 'react-redux'
// import { initializeStore } from '../lib/redux'
import { initializeApollo } from '../lib/apollo'


import Head from "next/head";
import gql from "graphql-tag";

const HomePage = (props) => {
  console.log(props)
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
      </div>
   

    
  );
};

  // export async function getStaticProps() {

  //   const apolloClient = initializeApollo()

  //   await apolloClient.query({

  //       query: gql`
  //       query User { 
  //         User (uuid: "02e92cbf-6736-46d9-bc7c-549209107a48"){
            
  //           username,
  //           email,
  //           description,
  //           profile_picture
  //       }
  //       user_i_follow: Follows (follower:"f5a13066-31d3-4aef-81b7-a5613b774734"){
  //             username,
  //             email 
  //         }
  //         users_following_me: Follows (followee: "02e92cbf-6736-46d9-bc7c-549209107a48"){
  //           username
  //           email
  //         }
  //       }
  //       `
  //   }).then(result => console.log(`result${JSON.stringify(result)}`)
  //   );
    

  //   return {
  //     props: {

  //       initialApolloState: apolloClient.cache.extract(),
  //     },
  //     unstable_revalidate: 1,
  //   }
  // }


export default HomePage;
