import * as React from "react";

const Home: React.FunctionComponent = () => (
  <div className="flex flex-col justify-center items-center h-screen w-1/2 mx-auto text-center">
    <h1 className="text-blue text-6xl font-bold m-3">Omiran</h1>
    <h3 className="font-normal text-3xl mb-5">
      A free and open source live stream software aimed to making streams
      accessible to everyone even with slower internet connections across the
      globe.
    </h3>
    <button className="bg-black px-8 py-3 text-xl rounded-lg text-white">
      Go Live
    </button>
  </div>
);

export default Home;
