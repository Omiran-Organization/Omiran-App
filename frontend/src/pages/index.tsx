import * as React from "react";

import Head from "next/head";

const HomePage: React.FunctionComponent = () => (
  <div className="flex flex-col justify-center items-center h-screen w-4/5 lg:w-1/2 mx-auto text-center">
    <Head>
      <title>Omiran</title>
    </Head>
    <h1 className="text-blue-500 text-6xl font-bold m-3">Omiran</h1>
    <h3 className="font-normal text-3xl mb-5">
      An open source live streaming application that is developer-oriented and removes the political overhead and censorship that other streaming platforms exhibit. The application's name, "Omiran", is Yoruba for, "alternative". This application, from a user's perspective can be interpreted as an alternative; an alternative to the harshly controlled and revenue focused platforms. We will also implement a 0% deduction of streamers' subscription money; funded by either github sponsors, or ads.
    </h3>
    <button className="bg-black px-8 py-3 text-xl rounded-lg text-white">
      Go Live
    </button>
  </div>
);

export default HomePage;
