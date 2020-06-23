import React from 'react';

import Head from 'next/head'

const HomePage: React.FunctionComponent = () => {
  return (
    <div className="main flex flex-col justify-center items-center w-4/5 lg:w-1/2 mx-auto text-center">
      <Head>
        <title>Omiran</title>
      </Head>

      <h1 className="text-orange-500 text-6xl font-bold m-3">Omiran</h1>
      <h3 className="text-white text-3xl mb-5">
        An open source live streaming application that is developer-oriented and removes the
        political overhead and censorship that other streaming platforms exhibit.
      </h3>

      <button className="bg-black px-8 py-3 text-xl rounded-lg text-white">Go Live</button>
    </div>
  )
}

export default HomePage
